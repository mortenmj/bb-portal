package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// DynamicExecutionMetrics holds the schema definition for the DynamicExecutionMetrics entity.
type DynamicExecutionMetrics struct {
	ent.Schema
}

// Fields of the DynamicExecutionMetrics.
func (DynamicExecutionMetrics) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the DynamicExecutionMetrics.
func (DynamicExecutionMetrics) Edges() []ent.Edge {
	return []ent.Edge{

		//edge back to the metrics object
		edge.From("metrics", Metrics.Type).Ref("dynamic_execution_metrics"),

		// Race statistics grouped by mnemonic, local_name, remote_name.
		edge.To("race_statistics", RaceStatistics.Type),
	}
}
