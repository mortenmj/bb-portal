package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TimingMetrics holds the schema definition for the TimingMetrics entity.
type TimingMetrics struct {
	ent.Schema
}

// Fields of the TimingMetrics.
func (TimingMetrics) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("cpu_time_in_ms").Optional(),
		field.Int64("wall_time_in_ms").Optional(),
		field.Int64("analysis_phase_time_in_ms").Optional(),
		field.Int64("execution_phase_time_in_ms").Optional(),
		field.Int64("actions_execution_start_in_ms").Optional(),
	}
}

// Edges of the TimingMetrics.
func (TimingMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metrics", Metrics.Type).Ref("timing_metrics"),
	}
}
