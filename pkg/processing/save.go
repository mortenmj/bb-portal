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
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetcomplete"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetconfigured"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetpair"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testcollection"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testresultbes"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testsummary"
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

	metrics, err := act.createMetrics(ctx, summary)
	if err != nil {
		return nil, err
	}

	targets, err := act.createTargets(ctx, summary)
	if err != nil {
		return nil, err
	}

	tests, err := act.createTestResults(ctx, summary)
	if err != nil {
		return nil, fmt.Errorf("could not save test results: %w", err)
	}

	bazelInvocation, err := act.saveBazelInvocation(ctx, summary, eventFile, buildRecord, metrics, tests, targets)
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

func (act SaveActor) saveBazelInvocation(
	ctx context.Context,
	summary *summary.Summary,
	eventFile *ent.EventFile,
	buildRecord *ent.Build,
	metrics *ent.Metrics,
	tests []*ent.TestCollection,
	targets []*ent.TargetPair) (*ent.BazelInvocation, error) {
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
		SetMetrics(metrics).
		AddTestCollection(tests...).
		AddTargets(targets...)

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

func (act SaveActor) createTargets(ctx context.Context, summary *summary.Summary) ([]*ent.TargetPair, error) {
	var err error = nil
	var result []*ent.TargetPair

	for targetLabel, targetPair := range summary.Targets {
		slog.Debug("processing target pair for %s", targetLabel)
		var configuration = targetPair.Configuration
		var completion = targetPair.Completion

		//configuration
		slog.Debug("processing configuration for %s", targetLabel)
		var target_configuration *ent.TargetConfigured
		target_configuration, err = act.db.TargetConfigured.Create().
			SetTag(configuration.Tag).
			SetStartTimeInMs(configuration.StartTimeInMs).
			SetTargetKind(configuration.TargetKind).
			SetTestSize(targetconfigured.TestSize(configuration.TestSize.String())).
			Save(ctx)
		if err != nil {
			slog.Error("problem saving target configuratiton object: %w", err)
			err = nil
		}

		//target completion
		//output group
		slog.Debug("processing output group for label %s on invocation %s", targetLabel, summary.InvocationID)
		//inline files
		slog.Debug("processing inline files for label %s on invocation %s", targetLabel, summary.InvocationID)
		var inline_files []*ent.TestFile
		for _, inlineFile := range completion.OutputGroup.InlineFiles {
			var inline_file *ent.TestFile
			inline_file, err = act.db.TestFile.Create().
				SetDigest(inlineFile.Digest).
				SetFile(inlineFile.File).
				SetName(inlineFile.Name).
				SetLength(inlineFile.Length).
				SetPrefix(inlineFile.Prefix).
				Save(ctx)
			if err != nil {
				slog.Error("problem saving inline file object: %w", err)
				err = nil
			}
			inline_files = append(inline_files, inline_file)
		}

		var output_group *ent.OutputGroup
		output_group, err = act.db.OutputGroup.Create().
			SetName(completion.OutputGroup.Name).
			SetIncomplete(completion.OutputGroup.Incomplete).
			AddInlineFiles(inline_files...).
			//TODO: implement named set of files logic to recursively add files to this collection
			Save(ctx)
		if err != nil {
			slog.Error("problem saving output group object: %w", err)
			err = nil
		}

		//important output
		slog.Debug("processing important output for label %s on invocation %s", targetLabel, summary.InvocationID)

		var important_output []*ent.TestFile
		for _, importantFile := range completion.ImportantOutput {

			var important_file *ent.TestFile
			important_file, err = act.db.TestFile.Create().
				SetDigest(importantFile.Digest).
				SetFile(importantFile.File).
				SetName(importantFile.Name).
				SetLength(importantFile.Length).
				SetPrefix(importantFile.Prefix).
				Save(ctx)
			if err != nil {
				slog.Error("problem saving important output object: %w", err)
				err = nil
			}
			important_output = append(important_output, important_file)
		}

		//directory output
		slog.Debug("processing directory output for label %s on invocation %s", targetLabel, summary.InvocationID)

		var directory_output []*ent.TestFile
		for _, directoryFile := range completion.DirectoryOutput {

			var directory_file *ent.TestFile
			directory_file, err = act.db.TestFile.Create().
				SetDigest(directoryFile.Digest).
				SetFile(directoryFile.File).
				SetName(directoryFile.Name).
				SetLength(directoryFile.Length).
				SetPrefix(directoryFile.Prefix).
				Save(ctx)
			if err != nil {
				slog.Error("problem saving directory output object: %w", err)
				err = nil
			}
			directory_output = append(directory_output, directory_file)
		}

		//target complete
		slog.Debug("processing target omplete for label %s on invocation %s", targetLabel, summary.InvocationID)
		var target_completion *ent.TargetComplete
		target_completion, err = act.db.TargetComplete.Create().
			SetSuccess(completion.Success).
			SetTargetKind(completion.TargetKind).
			SetTestSize(targetcomplete.TestSize(completion.TestSize.String())).
			SetTag(completion.Tag).
			SetEndTimeInMs(completion.EndTimeInMs).
			SetTestTimeout(completion.TestTimeout).
			SetTestTimeoutSeconds(completion.TestTimeoutSeconds).
			SetOutputGroup(output_group).
			AddImportantOutput(important_output...).
			AddDirectoryOutput(directory_output...).
			Save(ctx)
		if err != nil {
			slog.Error("problem saving target configuratiton object: %w", err)
			err = nil
		}

		//process the target pair
		slog.Debug("processing target pair for label %s on invocation %s", targetLabel, summary.InvocationID)
		var target_pair *ent.TargetPair
		target_pair, err = act.db.TargetPair.Create().
			SetCompletion(target_completion).
			SetConfiguration(target_configuration).
			SetLabel(targetLabel).
			SetDurationInMs(targetPair.DurationInMs).
			SetSuccess(targetPair.Success).
			SetTargetKind(targetPair.TargetKind).
			SetTestSize(targetpair.TestSize(targetPair.TestSize.String())).
			Save(ctx)
		if err != nil {
			slog.Error("problem saving target pair object: %w", err)
			err = nil
		}

		result = append(result, target_pair)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (act SaveActor) createTestResults(ctx context.Context, summary *summary.Summary) ([]*ent.TestCollection, error) {

	var err error = nil

	var result []*ent.TestCollection

	for testLabel, testCollection := range summary.Tests {

		//test summary
		slog.Debug("processing test summary for %s", testLabel)
		var ts = testCollection.TestSummary
		var db_test_summary *ent.TestSummary
		db_test_summary, err = act.db.TestSummary.Create().
			SetOverallStatus(testsummary.OverallStatus(ts.Status.String())).
			SetAttemptCount(ts.AttemptCount).
			SetRunCount(ts.RunCount).
			SetShardCount(ts.ShardCount).
			SetFirstStartTime(ts.FirstStartTime).
			SetLastStopTime(ts.LastStopTime).
			SetTotalRunCount(ts.TotalRunCount).
			SetTotalNumCached(ts.TotalNumCached).
			SetTotalRunDuration(ts.TotalRunDuration).
			SetLabel(testLabel).
			AddPassed().
			AddFailed().
			Save(ctx)

		if err != nil {
			slog.Error("problem saving test summary object: %w", err)
			err = nil
		}

		//test results
		slog.Debug("processing test results")
		var test_results []*ent.TestResultBES
		for _, tr := range testCollection.TestResults {

			//create the timing children
			var timing_children []*ent.TimingChild

			for _, tc := range tr.ExecutionInfo.TimingBreakdown.Child {
				var timing_child *ent.TimingChild
				timing_child, err = act.db.TimingChild.Create().
					SetName(tc.Name).
					SetTime(tc.Time).
					Save(ctx)

				if err != nil {
					slog.Error("problem saving timing child object: %w", err)
					err = nil
				}

				timing_children = append(timing_children, timing_child)
			}

			var timing_breakdown *ent.TimingBreakdown

			timing_breakdown, err = act.db.TimingBreakdown.Create().
				SetName(tr.ExecutionInfo.TimingBreakdown.Name).
				SetTime(tr.ExecutionInfo.TimingBreakdown.Time).
				AddChild(timing_children...).
				Save(ctx)

			if err != nil {
				slog.Error("problem saving timing breakdown object: %w", err)
				err = nil
			}

			var resource_usages []*ent.ResourceUsage

			for _, ru := range tr.ExecutionInfo.ResourceUsage {

				var resource_usage *ent.ResourceUsage

				resource_usage, err = act.db.ResourceUsage.Create().
					SetName(ru.Name).
					SetValue(ru.Value).
					Save(ctx)

				if err != nil {
					slog.Error("problem saving resource usage object: %w", err)
					err = nil
				}

				resource_usages = append(resource_usages, resource_usage)
			}

			var exection_info *ent.ExectionInfo

			exection_info, err = act.db.ExectionInfo.Create().
				SetStrategy(tr.ExecutionInfo.Strategy).
				SetCachedRemotely(tr.ExecutionInfo.CachedRemotely).
				SetExitCode(tr.ExecutionInfo.ExitCode).
				SetHostname(tr.ExecutionInfo.Hostname).
				SetTimingBreakdown(timing_breakdown).
				AddResourceUsage(resource_usages...).
				Save(ctx)

			if err != nil {
				slog.Error("problem saving execution info object: %w", err)
				err = nil
			}

			var test_result *ent.TestResultBES

			test_result, err = act.db.TestResultBES.Create().
				SetTestStatus(testresultbes.TestStatus(tr.Status.String())).
				SetAttempt(tr.Attempt).
				SetCachedLocally(tr.CachedLocally).
				SetLabel(tr.Label).
				SetWarning(tr.Warning).
				SetTestAttemptDurationMillis(tr.TestAttemptDurationMillis).
				SetTestAttemptStartMillisEpoch(tr.TestAttemptStartMillisEpoch).
				SetTestStatus(testresultbes.DefaultTestStatus).
				SetExecutionInfo(exection_info).
				Save(ctx)

			if err != nil {
				slog.Error("problem saving test result object: %w", err)
				err = nil
			}

			test_results = append(test_results, test_result)
		}

		var test_collection *ent.TestCollection
		test_collection, err = act.db.TestCollection.Create().
			SetLabel(testLabel).
			SetTestSummary(db_test_summary).
			AddTestResults(test_results...).
			SetOverallStatus(testcollection.OverallStatus(testCollection.OverallStatus.String())).
			SetStrategy(testCollection.Strategy).
			SetCachedLocally(testCollection.CachedLocally).
			SetCachedRemotely(testCollection.CachedRemotely).
			SetDurationMs(testCollection.DurationMs).
			Save(ctx)
		if err != nil {
			slog.Error("problem saving test collection object: %w", err)
			err = nil
		}
		result = append(result, test_collection)

	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (act SaveActor) createMetrics(ctx context.Context, summary *summary.Summary) (*ent.Metrics, error) {
	var err error
	var metrics *ent.Metrics

	//create the miss details
	slog.Debug("creating miss details")
	var miss_details []*ent.MissDetail
	for _, md := range summary.Metrics.ActionSummary.ActionCacheStatistics.MissDetails {

		var miss_detail *ent.MissDetail

		miss_detail, err = act.db.MissDetail.Create().
			SetCount(md.Count).
			SetReason(missdetail.Reason(md.Reason.String())).
			Save(ctx)

		// switch md.Reason.String() {
		// case "UNKNOWN":

		// 	miss_detail, err = act.db.MissDetail.Create().
		// 		SetCount(md.Count).
		// 		SetReason(missdetail.DefaultReason).
		// 		Save(ctx)

		// case "DIFFERENT_ACTION_KEY":

		// 	miss_detail, err = act.db.MissDetail.Create().
		// 		SetCount(md.Count).
		// 		SetReason(missdetail.ReasonDIFFERENT_ACTION_KEY).
		// 		Save(ctx)

		// case "DIFFERENT_DEPS":

		// 	miss_detail, err = act.db.MissDetail.Create().
		// 		SetCount(md.Count).
		// 		SetReason(missdetail.ReasonDIFFERENT_DEPS).
		// 		Save(ctx)

		// case "DIFFERENT_ENVIRONMENT":

		// 	miss_detail, err = act.db.MissDetail.Create().
		// 		SetCount(md.Count).
		// 		SetReason(missdetail.ReasonDIFFERENT_ENVIRONMENT).
		// 		Save(ctx)

		// case "DIFFERENT_FILES":

		// 	miss_detail, err = act.db.MissDetail.Create().
		// 		SetCount(md.Count).
		// 		SetReason(missdetail.ReasonDIFFERENT_FILES).
		// 		Save(ctx)

		// case "CORRUPTED_CACHE_ENTRY":

		// 	miss_detail, err = act.db.MissDetail.Create().
		// 		SetCount(md.Count).
		// 		SetReason(missdetail.ReasonCORRUPTED_CACHE_ENTRY).
		// 		Save(ctx)

		// case "NOT_CACHED":

		// 	miss_detail, err = act.db.MissDetail.Create().
		// 		SetCount(md.Count).
		// 		SetReason(missdetail.ReasonNOT_CACHED).
		// 		Save(ctx)

		// case "UNCONDITIONAL_EXECUTION":

		// 	miss_detail, err = act.db.MissDetail.Create().
		// 		SetCount(md.Count).
		// 		SetReason(missdetail.ReasonUNCONDITIONAL_EXECUTION).
		// 		Save(ctx)

		// }
		if err != nil {
			slog.Error("unable to create miss detail %w", err)
			err = nil
		}
		miss_details = append(miss_details, miss_detail)
	}

	//create the action cache statistics
	slog.Debug("creating action cache statistics")
	var actionCacheStatistics *ent.ActionCacheStatistics
	actionCacheStatistics, err = act.db.ActionCacheStatistics.Create().
		SetSizeInBytes(int64(summary.Metrics.ActionSummary.ActionCacheStatistics.SizeInBytes)).
		SetSaveTimeInMs(int64(summary.Metrics.ActionSummary.ActionCacheStatistics.SaveTimeInMs)).
		SetHits(summary.Metrics.ActionSummary.ActionCacheStatistics.Hits).
		SetMisses(summary.Metrics.ActionSummary.ActionCacheStatistics.Misses).
		AddMissDetails(miss_details...).
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
			SetSystemTime(ad.SystemTime).
			SetUserTime(ad.UserTime).
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

	var networkMetrics *ent.NetworkMetrics
	if summary.Metrics.NetworkMetrics.SystemNetworkStats != nil {
		var systemNetworkStats *ent.SystemNetworkStats
		slog.Debug("creating network metrics")
		//create the system network stats
		//TODO: need null checking for when the bazel command crashes
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
		networkMetrics, err = act.db.NetworkMetrics.Create().
			AddSystemNetworkStats(systemNetworkStats).
			Save(ctx)
		if err != nil {
			slog.Error("error creating network metrics. %w", err)
			err = nil
		}
	}

	//create the metrics object
	slog.Debug("creating metrics object")

	//there has to be a more elegant way to do this...
	if networkMetrics != nil {
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
	} else {
		metrics, err = act.db.Metrics.Create().
			AddActionSummary(actionSummary).
			AddBuildGraphMetrics(buildGraphMetrics).
			AddMemoryMetrics(memoryMetrics).
			AddTargetMetrics(targetMetrics).
			AddPackageMetrics(packageMetrics).
			AddCumulativeMetrics(cumulativeMetrics).
			AddTimingMetrics(timingMetrics).
			AddArtifactMetrics(artifactMetrics).
			Save(ctx)
	}

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
