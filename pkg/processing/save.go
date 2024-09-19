package processing

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/buildbarn/bb-portal/ent/gen/ent"
	"github.com/buildbarn/bb-portal/ent/gen/ent/blob"
	"github.com/buildbarn/bb-portal/ent/gen/ent/build"
	"github.com/buildbarn/bb-portal/ent/gen/ent/missdetail"
	"github.com/buildbarn/bb-portal/pkg/summary"
	"github.com/buildbarn/bb-portal/pkg/summary/detectors"
)

type SaveActor struct {
	db           *ent.Client
	blobArchiver BlobMultiArchiver
}

func (act SaveActor) SaveSummary(ctx context.Context, summary *summary.Summary) (*ent.BazelInvocation, error) {
	eventFile, err := act.saveEventFile(ctx, summary)
	if err != nil {
		return nil, fmt.Errorf("could not save EventFile: %w", err)
	}

	buildRecord, err := act.findOrCreateBuild(ctx, summary)
	if err != nil {
		return nil, err
	}

	metrics, err := act.CreateMetrics(ctx, summary)
	if err != nil {
		return nil, err
	}

	bazelInvocation, err := act.saveBazelInvocation(ctx, summary, eventFile, buildRecord, metrics)
	if err != nil {
		return nil, fmt.Errorf("could not save BazelInvocation: %w", err)
	}

	var detectedBlobs []detectors.BlobURI

	err = act.db.BazelInvocationProblem.MapCreateBulk(summary.Problems, func(create *ent.BazelInvocationProblemCreate, i int) {
		problem := summary.Problems[i]
		detectedBlobs = append(detectedBlobs, problem.DetectedBlobs...)
		create.
			SetProblemType(string(problem.ProblemType)).
			SetLabel(problem.Label).
			SetBepEvents(problem.BEPEvents).
			SetBazelInvocation(bazelInvocation)
	}).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not save BazelInvocationProblems: %w", err)
	}

	missingBlobs, err := act.determineMissingBlobs(ctx, detectedBlobs)
	if err != nil {
		return nil, err
	}

	err = act.db.Blob.MapCreateBulk(missingBlobs, func(create *ent.BlobCreate, i int) {
		b := missingBlobs[i]
		create.SetURI(string(b))
		// Leave defaults for other fields, all updated during archiving if it is enabled:
		// 	- size_bytes: 0
		// 	- archiving_status: QUEUED
		// 	- reason: null
		// 	- archive_url: null
	}).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not save Blobs: %w", err)
	}

	var archivedBlobs []ent.Blob
	archivedBlobs, err = act.blobArchiver.ArchiveBlobs(ctx, missingBlobs)
	if err != nil {
		return nil, fmt.Errorf("failed to archive blobs: %w", err)
	}
	for _, archivedBlob := range archivedBlobs {
		act.updateBlobRecord(ctx, archivedBlob)
	}

	return bazelInvocation, nil
}

func (act SaveActor) determineMissingBlobs(ctx context.Context, detectedBlobs []detectors.BlobURI) ([]detectors.BlobURI, error) {
	detectedBlobURIs := make([]string, 0, len(detectedBlobs))
	blobMap := make(map[string]struct{}, len(detectedBlobs))
	for _, detectedBlob := range detectedBlobs {
		detectedBlobURIs = append(detectedBlobURIs, string(detectedBlob))
	}
	foundInDB, err := act.db.Blob.Query().Where(blob.URIIn(detectedBlobURIs...)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query Blobs: %w", err)
	}

	for _, foundBlob := range foundInDB {
		blobMap[foundBlob.URI] = struct{}{}
	}
	missingBlobs := make([]detectors.BlobURI, 0, len(detectedBlobs)-len(foundInDB))
	for _, detectedBlob := range detectedBlobs {
		if _, ok := blobMap[string(detectedBlob)]; ok {
			continue
		}
		missingBlobs = append(missingBlobs, detectedBlob)
	}
	return missingBlobs, nil
}

func (act SaveActor) saveBazelInvocation(ctx context.Context, summary *summary.Summary, eventFile *ent.EventFile, buildRecord *ent.Build, metrics *ent.Metrics) (*ent.BazelInvocation, error) {
	create := act.db.BazelInvocation.Create().
		SetInvocationID(uuid.MustParse(summary.InvocationID)).
		SetStartedAt(summary.StartedAt).
		SetNillableEndedAt(summary.EndedAt).
		SetChangeNumber(int32(summary.ChangeNumber)).
		SetPatchsetNumber(int32(summary.PatchsetNumber)).
		SetSummary(*summary.InvocationSummary).
		SetBepCompleted(summary.BEPCompleted).
		SetStepLabel(summary.StepLabel).
		SetUserEmail(summary.UserEmail).
		SetBuildLogs(summary.BuildLogs.String()).
		SetUserLdap(summary.UserLDAP).
		SetRelatedFiles(summary.RelatedFiles).
		SetEventFile(eventFile).
		//metrics
		SetMetrics(metrics)

	if buildRecord != nil {
		create = create.SetBuild(buildRecord)
	}

	return create.
		Save(ctx)
}

func (act SaveActor) saveEventFile(ctx context.Context, summary *summary.Summary) (*ent.EventFile, error) {
	eventFile, err := act.db.EventFile.Create().
		SetURL(summary.EventFileURL).
		SetModTime(time.Now()).              // TODO: Save modTime in summary?
		SetProtocol("BEP").                  // Legacy: used to detect other protocols, e.g. for codechecks.
		SetMimeType("application/x-ndjson"). // NOTE: Only ndjson supported right now, but we should be able to add binary support.
		SetStatus("SUCCESS").                // TODO: Keep workflow of DETECTED->IMPORTING->...?
		Save(ctx)
	return eventFile, err
}

func (act SaveActor) CreateMetrics(ctx context.Context, summary *summary.Summary) (*ent.Metrics, error) {
	var err error
	var metrics *ent.Metrics

	//create the miss details
	slog.Debug("creating miss details")
	var missDetails []*ent.MissDetail
	for _, md := range summary.Metrics.ActionSummary.ActionCacheStatistics.MissDetails {

		var missDetail *ent.MissDetail
		switch md.Reason.String() {
		case "UNKNOWN":

			missDetail, err = act.db.MissDetail.Create().
				SetCount(md.Count).
				SetReason(missdetail.DefaultReason).
				Save(ctx)

		case "DIFFERENT_ACTION_KEY":

			missDetail, err = act.db.MissDetail.Create().
				SetCount(md.Count).
				SetReason(missdetail.ReasonDIFFERENT_ACTION_KEY).
				Save(ctx)

		case "DIFFERENT_DEPS":

			missDetail, err = act.db.MissDetail.Create().
				SetCount(md.Count).
				SetReason(missdetail.ReasonDIFFERENT_DEPS).
				Save(ctx)

		case "DIFFERENT_ENVIRONMENT":

			missDetail, err = act.db.MissDetail.Create().
				SetCount(md.Count).
				SetReason(missdetail.ReasonDIFFERENT_ENVIRONMENT).
				Save(ctx)

		case "DIFFERENT_FILES":

			missDetail, err = act.db.MissDetail.Create().
				SetCount(md.Count).
				SetReason(missdetail.ReasonDIFFERENT_FILES).
				Save(ctx)

		case "CORRUPTED_CACHE_ENTRY":

			missDetail, err = act.db.MissDetail.Create().
				SetCount(md.Count).
				SetReason(missdetail.ReasonCORRUPTED_CACHE_ENTRY).
				Save(ctx)

		case "NOT_CACHED":

			missDetail, err = act.db.MissDetail.Create().
				SetCount(md.Count).
				SetReason(missdetail.ReasonNOT_CACHED).
				Save(ctx)

		case "UNCONDITIONAL_EXECUTION":

			missDetail, err = act.db.MissDetail.Create().
				SetCount(md.Count).
				SetReason(missdetail.ReasonUNCONDITIONAL_EXECUTION).
				Save(ctx)

		}
		if err != nil {
			slog.Error("unable to create miss detail %w", err)
			err = nil
		}
		missDetails = append(missDetails, missDetail)
	}

	//create the action cache statistics
	slog.Debug("creating action cache statistics")
	var actionCacheStatistics *ent.ActionCacheStatistics
	actionCacheStatistics, err = act.db.ActionCacheStatistics.Create().
		SetSizeInBytes(int64(summary.Metrics.ActionSummary.ActionCacheStatistics.SizeInBytes)).
		SetSaveTimeInMs(int64(summary.Metrics.ActionSummary.ActionCacheStatistics.SaveTimeInMs)).
		SetHits(summary.Metrics.ActionSummary.ActionCacheStatistics.Hits).
		SetMisses(summary.Metrics.ActionSummary.ActionCacheStatistics.Misses).
		AddMissDetails(missDetails...).
		Save(ctx)

	if err != nil {
		slog.Error("error creating action cache statistics. %w", err)
		err = nil
	}

	//create runner counters
	slog.Debug("creating runner counts ")
	var runnerCounts []*ent.RunnerCount
	for _, rc := range summary.Metrics.ActionSummary.RunnerCount {
		var runnerCount *ent.RunnerCount
		runnerCount, err = act.db.RunnerCount.Create().
			SetActionsExecuted(int64(rc.Count)).
			SetName(rc.Name).
			SetExecKind(rc.ExecKind).
			Save(ctx)

		if err != nil {
			slog.Error("error creating runner count. %w", err)
			err = nil
		}

		runnerCounts = append(runnerCounts, runnerCount)

	}

	//create action datas
	slog.Debug("creating action datas")
	var actionDatas []*ent.ActionData
	for _, ad := range summary.Metrics.ActionSummary.ActionData {
		var actionData *ent.ActionData
		actionData, err = act.db.ActionData.Create().
			SetActionsExecuted(ad.ActionsExecuted).
			SetMnemonic(ad.Mnemonic).
			SetFirstStartedMs(ad.FirstStartedMs).
			SetLastEndedMs(ad.LastEndedMs).
			SetSystemTime(ad.SystemTime.Milliseconds()).
			SetUserTime(ad.UserTime.Milliseconds()).
			Save(ctx)

		if err != nil {
			slog.Error("error creating action data. %w", err)
			err = nil
		}

		actionDatas = append(actionDatas, actionData)

	}

	//create the action summary
	slog.Debug("creating acton summary")
	var actionSummary *ent.ActionSummary
	actionSummary, err = act.db.ActionSummary.Create().
		SetActionsCreated(summary.Metrics.ActionSummary.ActionsCreated).
		SetActionsCreatedNotIncludingAspects(summary.Metrics.ActionSummary.ActionsCreatedNotIncludingAspects).
		SetActionsExecuted(summary.Metrics.ActionSummary.ActionsExecuted).
		SetRemoteCacheHits(summary.Metrics.ActionSummary.RemoteCacheHits).
		AddActionCacheStatistics(actionCacheStatistics).
		AddRunnerCount(runnerCounts...).
		AddActionData(actionDatas...).
		Save(ctx)

	if err != nil {
		slog.Error("error creating action summary. %w", err)
		err = nil
	}

	//TODO:implement EvalutionStats once they exist on the proto
	//create the build graph metrics
	slog.Debug("creating memory metrics")
	var buildGraphMetrics *ent.BuildGraphMetrics
	buildGraphMetrics, err = act.db.BuildGraphMetrics.Create().
		SetActionLookupValueCount(summary.Metrics.BuildGraphMetrics.ActionLookupValueCount).
		SetActionLookupValueCountNotIncludingAspects(summary.Metrics.BuildGraphMetrics.ActionLookupValueCountNotIncludingAspects).
		SetActionCount(summary.Metrics.BuildGraphMetrics.ActionCount).
		SetInputFileConfiguredTargetCount(summary.Metrics.BuildGraphMetrics.InputFileConfiguredTargetCount).
		SetOutputFileConfiguredTargetCount(summary.Metrics.BuildGraphMetrics.OutputFileConfiguredTargetCount).
		SetOtherConfiguredTargetCount(summary.Metrics.BuildGraphMetrics.OtherConfiguredTargetCount).
		SetOutputArtifactCount(summary.Metrics.BuildGraphMetrics.OutputArtifactCount).
		SetPostInvocationSkyframeNodeCount(summary.Metrics.BuildGraphMetrics.PostInvocationSkyframeNodeCount).
		Save(ctx)
	if err != nil {
		slog.Error("error creating buildgraph metrics. %w", err)
		err = nil
	}

	//create garbage metrics
	slog.Debug("creating garbage metrics")
	var garbageMetrics []*ent.GarbageMetrics
	for _, gm := range summary.Metrics.MemoryMetrics.GarbageMetrics {
		var garbageMetric *ent.GarbageMetrics
		garbageMetric, err = act.db.GarbageMetrics.Create().
			SetGarbageCollected(gm.GarbageCollected).
			SetType(gm.Type).
			Save(ctx)

		if err != nil {
			slog.Error("error creating garbage metrics. %w", err)
			err = nil
		}

		garbageMetrics = append(garbageMetrics, garbageMetric)
	}

	//create memory metrics
	slog.Debug("creating memory metrics")
	var memoryMetrics *ent.MemoryMetrics
	memoryMetrics, err = act.db.MemoryMetrics.Create().
		SetPeakPostGcHeapSize(summary.Metrics.MemoryMetrics.PeakPostGcHeapSize).
		SetPeakPostGcTenuredSpaceHeapSize(summary.Metrics.MemoryMetrics.PeakPostGcTenuredSpaceHeapSize).
		SetUsedHeapSizePostBuild(summary.Metrics.MemoryMetrics.UsedHeapSizePostBuild).
		AddGarbageMetrics(garbageMetrics...).
		Save(ctx)
	if err != nil {
		slog.Error("error creating memory metrics. %w", err)
		err = nil
	}

	//create target metrics
	slog.Debug("creating target metrics")
	var targetMetrics *ent.TargetMetrics
	targetMetrics, err = act.db.TargetMetrics.Create().
		SetTargetsConfigured(summary.Metrics.TargetMetrics.TargetsConfigured).
		SetTargetsConfiguredNotIncludingAspects(summary.Metrics.TargetMetrics.TargetsConfiguredNotIncludingAspects).
		SetTargetsLoaded(summary.Metrics.TargetMetrics.TargetsLoaded).
		Save(ctx)
	if err != nil {
		slog.Error("error creating target metrics. %w", err)
		err = nil
	}

	//create the package load metrics
	slog.Debug("creating package load metrics")
	var packageLoadMetrics []*ent.PackageLoadMetrics

	for _, plm := range summary.Metrics.PackageMetrics.PackageLoadMetrics {
		var packageLoadMetric *ent.PackageLoadMetrics
		packageLoadMetric, err = act.db.PackageLoadMetrics.Create().
			SetName(plm.Name).
			SetLoadDuration(plm.LoadDuration.Milliseconds()).
			SetNumTargets(int64(plm.NumTargets)).
			SetComputationSteps(int64(plm.ComputationSteps)).
			SetNumTransitiveLoads(int64(plm.NumTransitiveLoads)).
			SetPackageOverhead(int64(plm.PackageOverhead)).
			Save(ctx)
		if err != nil {
			slog.Error("error creating package metrics. %w", err)
			err = nil
		}
		packageLoadMetrics = append(packageLoadMetrics, packageLoadMetric)
	}

	//create the package metrics
	slog.Debug("creating package metrics")
	var packageMetrics *ent.PackageMetrics
	packageMetrics, err = act.db.PackageMetrics.Create().
		SetPackagesLoaded(summary.Metrics.PackageMetrics.PackagesLoaded).
		AddPackageLoadMetrics(packageLoadMetrics...).
		Save(ctx)
	if err != nil {
		slog.Error("error creating package metrics. %w", err)
		err = nil
	}

	//create the cumulative metrics
	var cumulativeMetrics *ent.CumulativeMetrics
	slog.Debug("creating cumulative  metrics")
	cumulativeMetrics, err = act.db.CumulativeMetrics.Create().
		SetNumAnalyses(summary.Metrics.CumulativeMetrics.NumAnalyses).
		SetNumBuilds(summary.Metrics.CumulativeMetrics.NumBuilds).
		Save(ctx)
	if err != nil {
		slog.Error("error creating cumulative metrics. %w", err)
		err = nil
	}

	//create the timing metrics
	var timingMetrics *ent.TimingMetrics
	slog.Debug("creating timing metrics")
	timingMetrics, err = act.db.TimingMetrics.Create().
		SetAnalysisPhaseTimeInMs(summary.Metrics.TimingMetrics.AnalysisPhaseTimeInMs).
		SetCPUTimeInMs(summary.Metrics.TimingMetrics.CpuTimeInMs).
		SetExecutionPhaseTimeInMs(summary.Metrics.TimingMetrics.ExecutionPhaseTimeInMs).
		SetWallTimeInMs(summary.Metrics.TimingMetrics.WallTimeInMs).
		//TODO:
		//SetActionsExecutionStartInMs()
		Save(ctx)
	if err != nil {
		slog.Error("error creating timing metrics. %w", err)
		err = nil
	}

	//create source artifacts read
	slog.Debug("creating artifact metrics")
	var soureArtifactsRead *ent.FilesMetric
	soureArtifactsRead, err = act.db.FilesMetric.Create().
		SetCount(summary.Metrics.ArtifactMetrics.SourceArtifactsRead.Count).
		SetSizeInBytes(summary.Metrics.ArtifactMetrics.SourceArtifactsRead.SizeInBytes).
		Save(ctx)
	if err != nil {
		slog.Error("error creating source artifacts read metrics. %w", err)
		err = nil
	}

	//create output artifacts seen
	var outputArtifactsSeen *ent.FilesMetric
	outputArtifactsSeen, err = act.db.FilesMetric.Create().
		SetCount(summary.Metrics.ArtifactMetrics.OutputArtifactsSeen.Count).
		SetSizeInBytes(summary.Metrics.ArtifactMetrics.OutputArtifactsSeen.SizeInBytes).
		Save(ctx)
	if err != nil {
		slog.Error("error creating output artifacts seen metrics. %w", err)
		err = nil
	}

	//create output artifacts from action cache
	var outputArtifactsFromActionCache *ent.FilesMetric
	outputArtifactsFromActionCache, err = act.db.FilesMetric.Create().
		SetCount(summary.Metrics.ArtifactMetrics.OutputArtifactsFromActionCache.Count).
		SetSizeInBytes(summary.Metrics.ArtifactMetrics.OutputArtifactsFromActionCache.SizeInBytes).
		Save(ctx)
	if err != nil {
		slog.Error("error creating output artifacts from action cache metrics. %w", err)
		err = nil
	}

	//create top level artifacts
	var topLevelArtifacts *ent.FilesMetric
	topLevelArtifacts, err = act.db.FilesMetric.Create().
		SetCount(summary.Metrics.ArtifactMetrics.TopLevelArtifacts.Count).
		SetSizeInBytes(summary.Metrics.ArtifactMetrics.TopLevelArtifacts.SizeInBytes).
		Save(ctx)
	if err != nil {
		slog.Error("error creating top level artifacts metrics. %w", err)
		err = nil
	}

	//create the artifact metrics
	var artifactMetrics *ent.ArtifactMetrics
	slog.Debug("creating artifact metrics")
	artifactMetrics, err = act.db.ArtifactMetrics.Create().
		AddSourceArtifactsRead(soureArtifactsRead).
		AddOutputArtifactsSeen(outputArtifactsSeen).
		AddOutputArtifactsFromActionCache(outputArtifactsFromActionCache).
		AddTopLevelArtifacts(topLevelArtifacts).
		Save(ctx)
	if err != nil {
		slog.Error("error creating artifact metrics. %w", err)
		err = nil
	}

	slog.Debug("creating network metrics")
	//create the system network stats
	var systemNetworkStats *ent.SystemNetworkStats
	systemNetworkStats, err = act.db.SystemNetworkStats.Create().
		SetBytesRecv(int64(summary.Metrics.NetworkMetrics.SystemNetworkStats.BytesRecv)).
		SetBytesSent(int64(summary.Metrics.NetworkMetrics.SystemNetworkStats.BytesSent)).
		SetPacketsRecv(int64(summary.Metrics.NetworkMetrics.SystemNetworkStats.PacketsRecv)).
		SetPacketsSent(int64(summary.Metrics.NetworkMetrics.SystemNetworkStats.PacketsSent)).
		SetPeakBytesRecvPerSec(int64(summary.Metrics.NetworkMetrics.SystemNetworkStats.PeakBytesRecvPerSec)).
		SetPeakBytesSentPerSec(int64(summary.Metrics.NetworkMetrics.SystemNetworkStats.PeakBytesSentPerSec)).
		SetPeakPacketsRecvPerSec(int64(summary.Metrics.NetworkMetrics.SystemNetworkStats.PeakPacketsRecvPerSec)).
		SetPeakBytesSentPerSec(int64(summary.Metrics.NetworkMetrics.SystemNetworkStats.PeakPacketsSentPerSec)).
		Save(ctx)
	if err != nil {
		slog.Error("error creating system network stats metrics. %w", err)
		err = nil
	}

	//create the network metrics
	var networkMetrics *ent.NetworkMetrics
	networkMetrics, err = act.db.NetworkMetrics.Create().
		AddSystemNetworkStats(systemNetworkStats).
		Save(ctx)
	if err != nil {
		slog.Error("error creating network metrics. %w", err)
		err = nil
	}

	//create the metrics object
	slog.Debug("creating metrics object")
	metrics, err = act.db.Metrics.Create().
		AddActionSummary(actionSummary).
		AddBuildGraphMetrics(buildGraphMetrics).
		AddMemoryMetrics(memoryMetrics).
		AddTargetMetrics(targetMetrics).
		AddPackageMetrics(packageMetrics).
		AddCumulativeMetrics(cumulativeMetrics).
		AddTimingMetrics(timingMetrics).
		AddArtifactMetrics(artifactMetrics).
		AddNetworkMetrics(networkMetrics).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("unable to create metrics %w", err)
	}

	return metrics, nil

}

func (act SaveActor) findOrCreateBuild(ctx context.Context, summary *summary.Summary) (*ent.Build, error) {
	var err error
	var buildRecord *ent.Build

	if summary.BuildURL == "" {
		return nil, nil
	}

	slog.Info("Querying for build", "url", summary.BuildURL, "uuid", summary.BuildUUID)
	buildRecord, err = act.db.Build.Query().
		Where(build.BuildUUID(summary.BuildUUID)).First(ctx)

	if ent.IsNotFound(err) {
		slog.Info("Creating build", "url", summary.BuildURL, "uuid", summary.BuildUUID)
		buildRecord, err = act.db.Build.Create().
			SetBuildURL(summary.BuildURL).
			SetBuildUUID(summary.BuildUUID).
			SetEnv(buildEnvVars(summary.EnvVars)).
			Save(ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("could not find or create build: %w", err)
	}
	return buildRecord, nil
}

func (act SaveActor) updateBlobRecord(ctx context.Context, b ent.Blob) {
	update := act.db.Blob.Update().Where(blob.URI(b.URI)).SetArchivingStatus(b.ArchivingStatus)
	if b.ArchiveURL != "" {
		update = update.SetArchiveURL(b.ArchiveURL)
	}
	if b.Reason != "" {
		update = update.SetReason(b.Reason)
	}
	if b.SizeBytes != 0 {
		update = update.SetSizeBytes(b.SizeBytes)
	}
	if _, err := update.Save(ctx); err != nil {
		slog.Error("failed to save archived blob", "uri", b.URI, "err", err)
	}
}

// buildEnvVars filters the input so it only contains well known environment
// variables injected into a CI build (e.g. a Jenkins build). These are well-known
// Jenkins, etc. environment variables and/or environment variables associated
// with plugins for GitHub, Gerrit, etc.
func buildEnvVars(env map[string]string) map[string]string {
	buildEnv := make(map[string]string)
	for k, v := range env {
		if !summary.IsBuildEnvKey(k) {
			continue
		}
		buildEnv[k] = v
	}

	return buildEnv
}
