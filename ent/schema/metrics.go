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

// Edges of the Metrics.
func (Metrics) Edges() []ent.Edge {
	return []ent.Edge{

		// edge back to the bazel invocation
		edge.From("bazel_invocation", BazelInvocation.Type).
			Ref("metrics").
			Unique(),

		// the action summmary with details about actions executed
		edge.To("action_summary", ActionSummary.Type),

		// details about memory usage and garbage collections
		edge.To("memory_metrics", MemoryMetrics.Type),

		// target metrics
		edge.To("target_metrics", TargetMetrics.Type),

		// package metrics
		edge.To("package_metrics", PackageMetrics.Type),

		// timing metrics
		edge.To("timing_metrics", TimingMetrics.Type),

		//cumulative metrics
		edge.To("cumulative_metrics", CumulativeMetrics.Type),

		//artifact metrics
		edge.To("artifact_metrics", ArtifactMetrics.Type),

		//network metrics if available
		edge.To("network_metrics", NetworkMetrics.Type),

		//dynamic execution metrics if available
		edge.To("dynamic_execution_metrics", DynamicExecutionMetrics.Type),

		//build graph metrics
		edge.To("build_graph_metrics", BuildGraphMetrics.Type),
	}
}

func (Metrics) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("findMetrics"),
	}
}
