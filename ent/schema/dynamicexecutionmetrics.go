package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Blob holds the schema definition for the Blob entity.
type DynamicExecutionMetrics struct {
	ent.Schema
}

// Fields of the Blob.
func (DynamicExecutionMetrics) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Blob.
func (DynamicExecutionMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metrics", Metrics.Type).Ref("dynamic_execution_metrics"),
		edge.To("race_statistics", RaceStatistics.Type),
	}
}
