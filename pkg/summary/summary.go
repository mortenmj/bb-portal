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

// MISS REASON ENUM
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

func (r MissReason) EnumIndex() int32 {
	return int32(r)
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

// TEST STATUS ENUM
type TestStatus int32

const (
	NO_STATUS TestStatus = iota + 1
	PASSED
	FLAKY
	TIMEOUT
	FAILED
	INCOMPLETE
	REMOTE_FAILURE
	FAILED_TO_BUILD
	TOOL_HALTED_BEFORE_TESTING
)

func (r TestStatus) EnumIndex() int32 {
	return int32(r)
}

func (r TestStatus) String() string {
	return [...]string{
		"NO_STATUS",
		"PASSED",
		"FLAKY",
		"TIMEOUT",
		"FAILED",
		"INCOMPLETE",
		"REMOTE_FAILURE",
		"FAILED_TO_BUILD",
		"TOOL_HALTED_BEFORE_TESTING",
	}[r]
}

// TEST SIZE ENUM
type TestSize int32

const (
	UNKNOWN TestSize = iota + 1
	SMALL
	MEDIUM
	LARGE
	ENORMOUS
)

func (r TestSize) EnumIndex() int32 {
	return int32(r)
}
func (r TestSize) String() string {
	return [...]string{
		"UNKNOWN",
		"SMALL",
		"MEDIUM",
		"LARGE",
		"ENORMOUS",
	}[r]
}

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
	Tests          map[string]TestsCollection
	Targets        map[string]TargetPair
}

type Metrics struct {
	ActionSummary           ActionSummary
	MemoryMetrics           MemoryMetrics
	TargetMetrics           TargetMetrics
	PackageMetrics          PackageMetrics
	TimingMetrics           TimingMetrics
	CumulativeMetrics       CumulativeMetrics
	ArtifactMetrics         ArtifactMetrics
	BuildGraphMetrics       BuildGraphMetrics
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
	SystemTime      int64
	UserTime        int64
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

type ExecutionInfo struct {
	Strategy        string
	CachedRemotely  bool
	ExitCode        int32
	Hostname        string
	TimingBreakdown TimingBreakdown
	ResourceUsage   []ResourceUsage
}

type TestResult struct {
	Status                      TestStatus
	StatusDetails               string
	CachedLocally               bool
	TestAttemptDurationMillis   int64
	TestAttemptStartMillisEpoch int64
	Warning                     []string
	Run                         int
	Shard                       int
	Attempt                     int
	Label                       string
	TestActionOutput            []TestFile
	ExecutionInfo               ExecutionInfo
}

type TimingBreakdown struct {
	Name  string
	Time  string
	Child []TimingChild
}

type TimingChild struct {
	Name string
	Time string
}

type ResourceUsage struct {
	Name  string
	Value string
}

type TestFile struct {
	Digest string
	File   string
	Length int64
	Name   string
	Prefix []string
}

type TestSummary struct {
	Label            string
	Status           TestStatus
	TotalRunCount    int32
	RunCount         int32
	AttemptCount     int32
	ShardCount       int32
	TotalNumCached   int32
	FirstStartTime   int64
	LastStopTime     int64
	TotalRunDuration int64
	Passed           []TestFile
	Failed           []TestFile
}

type TargetConfigured struct {
	Tag        []string
	TargetKind string
	TestSize   TestSize
	//adding this to track time
	StartTimeInMs int64
}

type TargetComplete struct {
	Success            bool
	TargetKind         string
	TestSize           TestSize
	OutputGroup        OutputGroup
	ImportantOutput    []TestFile
	DirectoryOutput    []TestFile
	Tag                []string
	TestTimeoutSeconds int64
	TestTimeout        int64
	//adding this to track time
	EndTimeInMs int64
	//to lazy to implement this...maybe if i ever figure out how to generate that crap
	//FailureDetail FailureDetail
}

type OutputGroup struct {
	Name        string
	Incomplete  bool
	InlineFiles []TestFile
	FileSets    NamedSetOfFiles
}

type NamedSetOfFiles struct {
	Files    []TestFile
	FileSets *NamedSetOfFiles
}

// summary objects
type TestsCollection struct {
	TestSummary    TestSummary
	TestResults    []TestResult
	OverallStatus  TestStatus
	Strategy       string
	CachedLocally  bool
	CachedRemotely bool
	DurationMs     int64
}
type TargetPair struct {
	Configuration TargetConfigured
	Completion    TargetComplete
	DurationInMs  int64
	Success       bool
	TargetKind    string
	TestSize      TestSize
}
