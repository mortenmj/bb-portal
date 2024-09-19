package summary

import (
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/buildbarn/bb-portal/pkg/summary/detectors"
)

const (
	// stepLabelKey is used in buildMetadata events to provide a human-readable label for build steps.
	stepLabelKey = "BUILD_STEP_LABEL"
	userEmailKey = "user_email"
	userLdapKey  = "user_ldap"
)

const (
	ExitCodeSuccess     = 0
	ExitCodeInterrupted = 8
)

type Summary struct {
	*InvocationSummary
	Problems       []detectors.Problem
	RelatedFiles   map[string]string
	EventFileURL   string
	BEPCompleted   bool
	StartedAt      time.Time
	InvocationID   string
	StepLabel      string
	EndedAt        *time.Time
	ChangeNumber   int
	PatchsetNumber int
	BuildURL       string
	BuildUUID      uuid.UUID
	UserLDAP       string
	UserEmail      string
	BuildLogs      strings.Builder
	Metrics        Metrics
}

type Metrics struct {
	ActionSummary     ActionSummary
	MemoryMetrics     MemoryMetrics
	TargetMetrics     TargetMetrics
	PackageMetrics    PackageMetrics
	TimingMetrics     TimingMetrics
	CumulativeMetrics CumulativeMetrics
	ArtifactMetrics   ArtifactMetrics
	BuildGraphMetrics BuildGraphMetrics
	//below types are missing from the proto, but i want em
	NetworkMetrics          NetworkMetrics
	DynamicExecutionMetrics DynamicExecutionMetrics
}

type InvocationSummary struct {
	EnvVars          map[string]string
	ExitCode         *ExitCode
	BazelVersion     string
	BazelCommandLine BazelCommandLine
}

type ExitCode struct {
	Code int
	Name string
}

type BazelCommandLine struct {
	Executable string
	Command    string
	Residual   string
	Options    []string
}

// Blob holds information about a blob in the CAS. Should be easily converted to/from the one in the
// cas package. Copied into here so this package does not have *any* dependencies except standard
// libraries.
type Blob struct {
	BlobURI  url.URL
	Size     int
	Contents string
	Name     string
}

type ActionSummary struct {
	ActionsCreated                    int64
	ActionsCreatedNotIncludingAspects int64
	ActionsExecuted                   int64
	ActionData                        []ActionData

	RemoteCacheHits       int64
	RunnerCount           []RunnerCount
	ActionCacheStatistics ActionCacheStatistics
}

type ActionData struct {
	Mnemonic        string
	ActionsExecuted int64
	FirstStartedMs  int64
	LastEndedMs     int64
	SystemTime      time.Duration
	UserTime        time.Duration
}

type RunnerCount struct {
	Name     string
	Count    int32
	ExecKind string
}

type GarbageMetrics struct {
	Type             string
	GarbageCollected int64
}

type MemoryMetrics struct {
	UsedHeapSizePostBuild          int64
	PeakPostGcHeapSize             int64
	PeakPostGcTenuredSpaceHeapSize int64
	GarbageMetrics                 []GarbageMetrics
}

type TargetMetrics struct {
	TargetsLoaded                        int64
	TargetsConfigured                    int64
	TargetsConfiguredNotIncludingAspects int64
}

type PackageMetrics struct {
	PackagesLoaded     int64
	PackageLoadMetrics []PackageLoadMetrics
}

type TimingMetrics struct {
	CpuTimeInMs            int64
	WallTimeInMs           int64
	AnalysisPhaseTimeInMs  int64
	ExecutionPhaseTimeInMs int64
}

type CumulativeMetrics struct {
	NumAnalyses int32
	NumBuilds   int32
}

type ArtifactMetrics struct {
	SourceArtifactsRead            FilesMetric
	OutputArtifactsSeen            FilesMetric
	OutputArtifactsFromActionCache FilesMetric
	TopLevelArtifacts              FilesMetric
}

type FilesMetric struct {
	SizeInBytes int64
	Count       int32
}

type SystemNetworkStats struct {
	BytesSent             uint64
	BytesRecv             uint64
	PacketsSent           uint64
	PacketsRecv           uint64
	PeakBytesSentPerSec   uint64
	PeakBytesRecvPerSec   uint64
	PeakPacketsSentPerSec uint64
	PeakPacketsRecvPerSec uint64
}

type NetworkMetrics struct {
	SystemNetworkStats *SystemNetworkStats
}
type ActionCacheStatistics struct {
	SizeInBytes  uint64
	SaveTimeInMs uint64
	LoadTimeInMs uint64
	Hits         int32
	Misses       int32
	MissDetails  []MissDetail
}
type MissDetail struct {
	Reason MissReason
	Count  int32
}

func (r MissReason) EnumIndex() int {
	return int(r)
}

func (r MissReason) String() string {
	return [...]string{
		"UNKNOWN",
		"DIFFERENT_ACTION_KEY",
		"DIFFERENT_DEPS",
		"DIFFERENT_ENVIRONMENT",
		"DIFFERENT_FILES",
		"CORRUPTED_CACHE_ENTRY",
		"NOT_CACHED",
		"UNCONDITIONAL_EXECUTION",
	}[r]
}

type MissReason int32

const (
	DIFFERENT_ACTION_KEY MissReason = iota + 1
	DIFFERENT_DEPS
	DIFFERENT_ENVIRONMENT
	DIFFERENT_FILES
	CORRUPTED_CACHE_ENTRY
	NOT_CACHED
	UNCONDITIONAL_EXECUTION
)

type PackageLoadMetrics struct {
	Name               string
	LoadDuration       time.Duration
	NumTargets         uint64
	ComputationSteps   uint64
	NumTransitiveLoads uint64
	PackageOverhead    uint64
}

type DynamicExecutionMetrics struct {
	RaceStatistics []RaceStatistics
}

type RaceStatistics struct {
	Mnemonic     string
	LocalRunner  string
	RemoteRunner string
	LocalWins    int64
	RemoteWins   int64
}

type EvaluationStat struct {
	SkyfunctionName string
	Count           int64
}

type BuildGraphMetrics struct {
	ActionLookupValueCount                    int32
	ActionLookupValueCountNotIncludingAspects int32
	ActionCount                               int32
	InputFileConfiguredTargetCount            int32
	OutputFileConfiguredTargetCount           int32
	OtherConfiguredTargetCount                int32
	OutputArtifactCount                       int32
	PostInvocationSkyframeNodeCount           int32
	DirtiedValues                             []EvaluationStat
	ChangedValues                             []EvaluationStat
	BuiltValues                               []EvaluationStat
	CleanedValues                             []EvaluationStat
	EvaluatedValues                           []EvaluationStat
}
