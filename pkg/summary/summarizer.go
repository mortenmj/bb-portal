package summary

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/iterator"

	"github.com/buildbarn/bb-portal/pkg/events"
	"github.com/buildbarn/bb-portal/pkg/summary/detectors"
	"github.com/buildbarn/bb-portal/third_party/bazel/gen/bes"
	"github.com/buildbarn/bb-portal/third_party/bazel/gen/bescore"
)

type Summarizer struct {
	summary         *Summary
	problemDetector detectors.ProblemDetector
}

func Summarize(ctx context.Context, eventFileURL string) (*Summary, error) {
	reader, err := os.Open(eventFileURL)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %w", eventFileURL, err)
	}
	defer reader.Close()

	problemDetector := detectors.NewProblemDetector()
	summarizer := newSummarizer(eventFileURL, problemDetector)
	it := events.NewBuildEventIterator(ctx, reader)
	return summarizer.summarize(it)
}

func NewSummarizer() *Summarizer {
	return newSummarizer("", detectors.NewProblemDetector())
}

func newSummarizer(eventFileURL string, problemDetector detectors.ProblemDetector) *Summarizer {
	return &Summarizer{
		summary: &Summary{
			InvocationSummary: &InvocationSummary{},
			EventFileURL:      eventFileURL,
			RelatedFiles: map[string]string{
				filepath.Base(eventFileURL): eventFileURL,
			},
		},
		problemDetector: problemDetector,
	}
}

func (s Summarizer) summarize(it *events.BuildEventIterator) (*Summary, error) {
	for {
		buildEvent, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get build event: %w", err)
		}

		err = s.ProcessEvent(buildEvent)
		if err != nil {
			return nil, fmt.Errorf("failed to process event (with id: %s): %w", buildEvent.Id.String(), err)
		}
	}

	return s.FinishProcessing()
}

func (s Summarizer) FinishProcessing() (*Summary, error) {
	// If problems are ignored for the exit code, return immediately.
	if !shouldIgnoreProblems(s.summary.ExitCode) {
		// Add any detected test problems.
		problems, problemsErr := s.problemDetector.Problems()
		if problemsErr != nil {
			return nil, problemsErr
		}
		s.summary.Problems = append(s.summary.Problems, problems...)
	}

	return s.summary, nil
}

func (s Summarizer) ProcessEvent(buildEvent *events.BuildEvent) error {
	// Let problem detector process every event.
	s.problemDetector.ProcessBEPEvent(buildEvent)

	switch buildEvent.GetId().GetId().(type) {
	case *bes.BuildEventId_Started:
		s.handleStarted(buildEvent.GetStarted())

	case *bes.BuildEventId_BuildMetadata:
		s.handleBuildMetadata(buildEvent.GetBuildMetadata())

	case *bes.BuildEventId_BuildFinished:
		s.handleBuildFinished(buildEvent.GetFinished())

	case *bes.BuildEventId_BuildMetrics:
		s.handleBuildMetrics(buildEvent.GetBuildMetrics())

	case *bes.BuildEventId_StructuredCommandLine:
		err := s.handleStructuredCommandLine(buildEvent.GetStructuredCommandLine())
		if err != nil {
			return err
		}

	case *bes.BuildEventId_OptionsParsed:
		s.handleOptionsParsed(buildEvent.GetOptionsParsed())

	case *bes.BuildEventId_BuildToolLogs:
		err := s.handleBuildToolLogs(buildEvent.GetBuildToolLogs())
		if err != nil {
			return err
		}
	case *bes.BuildEventId_Progress:
		s.handleProgress(buildEvent.GetProgress())
	}

	s.summary.BEPCompleted = buildEvent.GetLastMessage()
	return nil
}

func (s Summarizer) handleStarted(started *bes.BuildStarted) {
	var startedAt time.Time
	if started.GetStartTime() != nil {
		startedAt = started.GetStartTime().AsTime()
	} else {
		//nolint:staticcheck // Keep backwards compatibility until the field is removed.
		startedAt = time.UnixMilli(started.GetStartTimeMillis())
	}
	s.summary.StartedAt = startedAt
	s.summary.InvocationID = started.GetUuid()
	s.summary.BazelVersion = started.GetBuildToolVersion()
}

func (s Summarizer) handleBuildMetadata(metadataProto *bes.BuildMetadata) {
	metadataMap := metadataProto.GetMetadata()
	//extract user data
	if metadataMap == nil {
		return
	}
	stepLabel, stepLabelOk := metadataMap[stepLabelKey]
	if !stepLabelOk {
		slog.Debug("No step label found in build metadata")
	}
	userEmail, userEmailOk := metadataMap[userEmailKey]
	if !userEmailOk {
		slog.Debug("No user email found in build metadata")
	}
	userLdap, userLdapOk := metadataMap[userLdapKey]
	if !userLdapOk {
		slog.Debug("No user ldap information found in build metadata")
	}
	s.summary.StepLabel = stepLabel
	s.summary.UserEmail = userEmail
	s.summary.UserLDAP = userLdap
}

func (s Summarizer) handleBuildMetrics(metrics *bes.BuildMetrics) {

	//action metrics

	var miss_details []MissDetail = make([]MissDetail, 0)

	for _, md := range metrics.ActionSummary.ActionCacheStatistics.MissDetails {
		miss_detail := MissDetail{
			Count:  md.Count,
			Reason: MissReason(*md.Reason.Enum()),
		}
		miss_details = append(miss_details, miss_detail)
	}

	action_cache_statistics := ActionCacheStatistics{
		SizeInBytes:  metrics.ActionSummary.ActionCacheStatistics.SizeInBytes,
		SaveTimeInMs: metrics.ActionSummary.ActionCacheStatistics.SaveTimeInMs,
		//TODO: investigate why load time in ms is not available on bes object
		//LoadTimeInMs: metrics.ActionSummary.ActionCacheStatistics
		Hits:        metrics.ActionSummary.ActionCacheStatistics.Hits,
		Misses:      metrics.ActionSummary.ActionCacheStatistics.Misses,
		MissDetails: miss_details,
	}

	var runner_counts []RunnerCount = make([]RunnerCount, 0)
	for _, rc := range metrics.ActionSummary.RunnerCount {
		runner_count := RunnerCount{
			ExecKind: rc.ExecKind,
			Count:    rc.Count,
			Name:     rc.Name,
		}
		runner_counts = append(runner_counts, runner_count)
	}

	var action_datas []ActionData = make([]ActionData, 0)
	for _, ad := range metrics.ActionSummary.ActionData {
		action_data := ActionData{
			Mnemonic:        ad.Mnemonic,
			UserTime:        ad.UserTime.AsDuration(),
			SystemTime:      ad.SystemTime.AsDuration(),
			ActionsExecuted: ad.ActionsExecuted,
			FirstStartedMs:  ad.FirstStartedMs,
			LastEndedMs:     ad.LastEndedMs,
		}
		action_datas = append(action_datas, action_data)
	}

	action_summary := ActionSummary{
		ActionsCreated:                    metrics.ActionSummary.ActionsCreated,
		ActionsExecuted:                   metrics.ActionSummary.ActionsExecuted,
		ActionsCreatedNotIncludingAspects: metrics.ActionSummary.ActionsCreatedNotIncludingAspects,
		ActionCacheStatistics:             action_cache_statistics,
		RunnerCount:                       runner_counts,
		ActionData:                        action_datas,
	}

	//memory metrics
	var garbage_metrics []GarbageMetrics = make([]GarbageMetrics, 0)

	for _, gm := range metrics.MemoryMetrics.GarbageMetrics {
		garbage_metric := GarbageMetrics{
			Type:             gm.Type,
			GarbageCollected: gm.GarbageCollected,
		}
		garbage_metrics = append(garbage_metrics, garbage_metric)
	}

	memory_metrics := MemoryMetrics{
		PeakPostGcHeapSize:             metrics.MemoryMetrics.PeakPostGcHeapSize,
		PeakPostGcTenuredSpaceHeapSize: metrics.MemoryMetrics.PeakPostGcTenuredSpaceHeapSize,
		UsedHeapSizePostBuild:          metrics.MemoryMetrics.UsedHeapSizePostBuild,
		GarbageMetrics:                 garbage_metrics,
	}

	//target metrics
	target_metrics := TargetMetrics{
		TargetsConfigured:                    metrics.TargetMetrics.TargetsConfigured,
		TargetsConfiguredNotIncludingAspects: metrics.TargetMetrics.TargetsConfiguredNotIncludingAspects,
		TargetsLoaded:                        metrics.TargetMetrics.TargetsLoaded,
	}

	//package metrics
	var package_load_metrics []PackageLoadMetrics = make([]PackageLoadMetrics, 0)

	for _, plm := range metrics.PackageMetrics.PackageLoadMetrics {
		package_load_metric := PackageLoadMetrics{
			Name:               *plm.Name,
			NumTargets:         *plm.NumTargets,
			LoadDuration:       plm.LoadDuration.AsDuration(),
			ComputationSteps:   *plm.ComputationSteps,
			NumTransitiveLoads: *plm.NumTransitiveLoads,
			PackageOverhead:    *plm.PackageOverhead,
		}
		package_load_metrics = append(package_load_metrics, package_load_metric)
	}

	package_metrics := PackageMetrics{
		PackagesLoaded:     metrics.PackageMetrics.PackagesLoaded,
		PackageLoadMetrics: package_load_metrics,
	}

	//timing metrics

	timing_metrics := TimingMetrics{
		CpuTimeInMs:            metrics.TimingMetrics.CpuTimeInMs,
		WallTimeInMs:           metrics.TimingMetrics.WallTimeInMs,
		ExecutionPhaseTimeInMs: metrics.TimingMetrics.ExecutionPhaseTimeInMs,
		AnalysisPhaseTimeInMs:  metrics.TimingMetrics.AnalysisPhaseTimeInMs,
		//TODO: why isn't this on the proto
		//ActionsExecutionStartInMs: metrics.TimingMetrics.ActionsExecutionStartInMs,
	}

	//artifact metrics

	source_artifacts_read := FilesMetric{
		SizeInBytes: metrics.ArtifactMetrics.SourceArtifactsRead.SizeInBytes,
		Count:       metrics.ArtifactMetrics.SourceArtifactsRead.Count,
	}

	output_artifacts_seen := FilesMetric{
		SizeInBytes: metrics.ArtifactMetrics.OutputArtifactsSeen.SizeInBytes,
		Count:       metrics.ArtifactMetrics.OutputArtifactsSeen.Count,
	}

	output_artifacts_from_action_cache := FilesMetric{
		SizeInBytes: metrics.ArtifactMetrics.OutputArtifactsFromActionCache.SizeInBytes,
		Count:       metrics.ArtifactMetrics.OutputArtifactsFromActionCache.Count,
	}

	top_level_artifacts := FilesMetric{
		SizeInBytes: metrics.ArtifactMetrics.TopLevelArtifacts.SizeInBytes,
		Count:       metrics.ArtifactMetrics.TopLevelArtifacts.Count,
	}

	artifact_metrics := ArtifactMetrics{
		SourceArtifactsRead:            source_artifacts_read,
		OutputArtifactsSeen:            output_artifacts_seen,
		OutputArtifactsFromActionCache: output_artifacts_from_action_cache,
		TopLevelArtifacts:              top_level_artifacts,
	}

	//cumulative metrics

	cumulative_metrics := CumulativeMetrics{
		NumAnalyses: metrics.CumulativeMetrics.NumAnalyses,
		NumBuilds:   metrics.CumulativeMetrics.NumBuilds,
	}

	//TODO: dynamic metrics are not on the proto
	var race_statistics []RaceStatistics = make([]RaceStatistics, 0)
	// for _,rc := range metrics.
	dynamic_metrics := DynamicExecutionMetrics{
		RaceStatistics: race_statistics,
	}

	//network metrics are currently empty...not sure why

	var system_network_stats SystemNetworkStats

	if metrics.NetworkMetrics != nil {

		system_network_stats = SystemNetworkStats{
			BytesSent:             metrics.NetworkMetrics.SystemNetworkStats.BytesSent,
			BytesRecv:             metrics.NetworkMetrics.SystemNetworkStats.BytesRecv,
			PacketsSent:           metrics.NetworkMetrics.SystemNetworkStats.PacketsSent,
			PacketsRecv:           metrics.NetworkMetrics.SystemNetworkStats.PacketsRecv,
			PeakBytesSentPerSec:   metrics.NetworkMetrics.SystemNetworkStats.PeakBytesSentPerSec,
			PeakBytesRecvPerSec:   metrics.NetworkMetrics.SystemNetworkStats.PeakBytesRecvPerSec,
			PeakPacketsSentPerSec: metrics.NetworkMetrics.SystemNetworkStats.PeakPacketsSentPerSec,
			PeakPacketsRecvPerSec: metrics.NetworkMetrics.SystemNetworkStats.PeakPacketsRecvPerSec,
		}
	}

	network_metrics := NetworkMetrics{
		SystemNetworkStats: &system_network_stats,
	}

	//TODO: these vaues are not on the proto.
	var dirtied_values []EvaluationStat = make([]EvaluationStat, 0)
	var changed_values []EvaluationStat = make([]EvaluationStat, 0)
	var built_values []EvaluationStat = make([]EvaluationStat, 0)
	var cleaned_values []EvaluationStat = make([]EvaluationStat, 0)
	var evaluated_values []EvaluationStat = make([]EvaluationStat, 0)

	buildgraph_metrics := BuildGraphMetrics{
		ActionLookupValueCount:                    metrics.BuildGraphMetrics.ActionLookupValueCount,
		ActionLookupValueCountNotIncludingAspects: metrics.BuildGraphMetrics.ActionLookupValueCountNotIncludingAspects,
		ActionCount:                     metrics.BuildGraphMetrics.ActionCount,
		InputFileConfiguredTargetCount:  metrics.BuildGraphMetrics.InputFileConfiguredTargetCount,
		OutputFileConfiguredTargetCount: metrics.BuildGraphMetrics.OutputFileConfiguredTargetCount,
		OtherConfiguredTargetCount:      metrics.BuildGraphMetrics.OtherConfiguredTargetCount,
		OutputArtifactCount:             metrics.BuildGraphMetrics.OutputArtifactCount,
		PostInvocationSkyframeNodeCount: metrics.BuildGraphMetrics.PostInvocationSkyframeNodeCount,
		DirtiedValues:                   dirtied_values,
		ChangedValues:                   changed_values,
		BuiltValues:                     built_values,
		CleanedValues:                   cleaned_values,
		EvaluatedValues:                 evaluated_values,
	}

	summary_metrics := Metrics{
		ActionSummary:           action_summary,
		MemoryMetrics:           memory_metrics,
		TargetMetrics:           target_metrics,
		PackageMetrics:          package_metrics,
		TimingMetrics:           timing_metrics,
		ArtifactMetrics:         artifact_metrics,
		CumulativeMetrics:       cumulative_metrics,
		NetworkMetrics:          network_metrics,
		BuildGraphMetrics:       buildgraph_metrics,
		DynamicExecutionMetrics: dynamic_metrics,
	}

	s.summary.Metrics = summary_metrics
}

func (s Summarizer) handleBuildFinished(finished *bes.BuildFinished) {
	var endedAt time.Time
	if finished.GetFinishTime() != nil {
		endedAt = finished.GetFinishTime().AsTime()
	} else {
		//nolint:staticcheck // Keep backwards compatibility until the field is removed.
		endedAt = time.UnixMilli(finished.GetFinishTimeMillis())
	}
	s.summary.EndedAt = &endedAt
	s.summary.InvocationSummary.ExitCode = &ExitCode{
		Code: int(finished.GetExitCode().GetCode()),
		Name: finished.GetExitCode().GetName(),
	}
}

func (s Summarizer) handleStructuredCommandLine(structuredCommandLine *bescore.CommandLine) error {
	if structuredCommandLine.GetCommandLineLabel() != "original" {
		return nil
	}

	s.updateEnvVarsAndCommandFromStructuredCommandLine(structuredCommandLine)

	// Parse Gerrit change number if available.
	if changeNumberStr, ok := s.summary.InvocationSummary.EnvVars["GERRIT_CHANGE_NUMBER"]; ok && changeNumberStr != "" {
		changeNumber, err := envToI(s.summary.InvocationSummary.EnvVars, "GERRIT_CHANGE_NUMBER")
		if err != nil {
			return err
		}
		s.summary.ChangeNumber = changeNumber
	}

	// Parse Gerrit patchset number if available.
	if patchsetNumberStr, ok := s.summary.InvocationSummary.EnvVars["GERRIT_PATCHSET_NUMBER"]; ok && patchsetNumberStr != "" {
		patchsetNumber, err := envToI(s.summary.InvocationSummary.EnvVars, "GERRIT_PATCHSET_NUMBER")
		if err != nil {
			return err
		}
		s.summary.PatchsetNumber = patchsetNumber
	}

	// Decode commit message, so that client doesn't have to.
	commitMessage := s.summary.InvocationSummary.EnvVars["GERRIT_CHANGE_COMMIT_MESSAGE"]
	if commitMessage != "" {
		decodedCommitMessage, err := base64.StdEncoding.DecodeString(commitMessage)
		if err == nil {
			s.summary.InvocationSummary.EnvVars["GERRIT_CHANGE_COMMIT_MESSAGE"] = string(decodedCommitMessage)
		} else {
			slog.Debug("GERRIT_CHANGE_COMMIT_MESSAGE was not base64 encoded, assuming it is normal string")
		}
	}

	// Set build URL and UUID
	s.summary.BuildURL = s.summary.InvocationSummary.EnvVars["BUILD_URL"]
	s.summary.BuildUUID = uuid.NewSHA1(uuid.NameSpaceURL, []byte(s.summary.BuildURL))

	return nil
}

func (s Summarizer) handleOptionsParsed(optionsParsed *bes.OptionsParsed) {
	s.summary.InvocationSummary.BazelCommandLine.Options = optionsParsed.GetExplicitCmdLine()
}

func (s Summarizer) handleProgress(progressMsg *bes.Progress) {
	s.summary.BuildLogs.WriteString(progressMsg.GetStderr())
	s.summary.BuildLogs.WriteString(progressMsg.GetStdout())
}

func (s Summarizer) handleBuildToolLogs(buildToolLogs *bes.BuildToolLogs) error {
	for _, logs := range buildToolLogs.GetLog() {
		uri := logs.GetUri()
		blobURI := detectors.BlobURI(uri)

		if s.summary.RelatedFiles == nil {
			s.summary.RelatedFiles = map[string]string{}
		}
		if logs.GetUri() != "" {
			s.summary.RelatedFiles[logs.GetName()] = string(blobURI)
		}
	}
	return nil
}

func (s Summarizer) updateEnvVarsAndCommandFromStructuredCommandLine(structuredCommandLine *bescore.CommandLine) {
	sections := structuredCommandLine.GetSections()
	for _, section := range sections {
		label := section.GetSectionLabel()
		if label == "command options" {
			s.summary.InvocationSummary.EnvVars = map[string]string{}
			ParseEnvVarsFromSectionOptions(section, &s.summary.InvocationSummary.EnvVars)
		} else if section.GetChunkList() != nil {
			sectionChunksStr := strings.Join(section.GetChunkList().GetChunk(), " ")
			switch label {
			case "executable":
				s.summary.InvocationSummary.BazelCommandLine.Executable = sectionChunksStr
			case "command":
				s.summary.InvocationSummary.BazelCommandLine.Command = sectionChunksStr
			case "residual":
				s.summary.InvocationSummary.BazelCommandLine.Residual = sectionChunksStr
			}
		}
	}
}

func shouldIgnoreProblems(exitCode *ExitCode) bool {
	return exitCode != nil && (exitCode.Code == ExitCodeSuccess || exitCode.Code == ExitCodeInterrupted)
}

func envToI(envVars map[string]string, name string) (int, error) {
	res, err := strconv.Atoi(envVars[name])
	if err != nil {
		slog.Error("failed to parse env var to int", "envKey", name, "envValue", envVars[name], "err", err)
		return 0, fmt.Errorf("failed to parse %s (value: %s) as an int: %w", name, envVars[name], err)
	}
	return res, nil
}

func ParseEnvVarsFromSectionOptions(section *bescore.CommandLineSection, destMap *map[string]string) {
	if section.GetOptionList() == nil {
		return
	}
	options := section.GetOptionList().GetOption()
	for _, option := range options {
		if option.GetOptionName() != "client_env" {
			// Only looking for env vars from the client env
			continue
		}
		envPair := option.GetOptionValue()
		equalIndex := strings.Index(envPair, "=")
		if equalIndex <= 0 {
			// Skip anything missing an equals sign. The env vars come in the format key=value
			continue
		}
		envName := envPair[:equalIndex]
		envValue := envPair[equalIndex+1:]
		(*destMap)[envName] = envValue
	}
}
