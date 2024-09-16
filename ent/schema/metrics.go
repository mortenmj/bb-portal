package schema

import (
	"entgo.io/ent"
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
		edge.To("action_summary", ActionSummary.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
		edge.To("memory_metrics", MemoryMetrics.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
		edge.To("target_metrics", TargetMetrics.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
		edge.To("package_metrics", PackageMetrics.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
		edge.To("timing_metrics", TimingMetrics.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
		edge.To("cumulative_metrics", CumulativeMetrics.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
		edge.To("artifact_metrics", ArtifactMetrics.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
		edge.To("network_metrics", NetworkMetrics.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
		edge.To("dynamic_execution_metrics", DynamicExecutionMetrics.Type),
		//Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.
	}
}
