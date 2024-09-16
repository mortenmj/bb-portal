package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CumulativeMetrics holds the schema definition for the CumulativeMetrics entity.
type CumulativeMetrics struct {
	ent.Schema
}

// Fields of the CumulativeMetrics.
func (CumulativeMetrics) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("num_analyses").Optional(),
		field.Int32("num_builds").Optional(),
	}
}

// Edges of the TimingMetrics.
func (CumulativeMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metrics", Metrics.Type).Ref("cumulative_metrics"),
	}
}
