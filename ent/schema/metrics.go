package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// Metrics holds the schema definition for the Metrics entity.
type Metrics struct {
	ent.Schema
}

// Fields of the Metrics struct
func (Metrics) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the BazelInvocation.
func (Metrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bazel_invocation", BazelInvocation.Type).
			Ref("metrics").
			Unique(),
		edge.To("action_summary", ActionSummary.Type),
		edge.To("memory_metrics", MemoryMetrics.Type),
		edge.To("target_metrics", TargetMetrics.Type),
		edge.To("package_metrics", PackageMetrics.Type),
		edge.To("timing_metrics", TimingMetrics.Type),
		edge.To("cumulative_metrics", CumulativeMetrics.Type),
		edge.To("artifact_metrics", ArtifactMetrics.Type),
		edge.To("network_metrics", NetworkMetrics.Type),
		edge.To("dynamic_execution_metrics", DynamicExecutionMetrics.Type),
		edge.To("build_graph_metrics", BuildGraphMetrics.Type),
		edge.To("test_results", TestResultBES.Type),
		edge.To("test_summary", TestSummary.Type),
	}
}

func (Metrics) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("findMetrics"),
	}
}
