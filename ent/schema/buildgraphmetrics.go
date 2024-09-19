package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type BuildGraphMetrics struct {
	ent.Schema
}

// Fields of the Blob.
func (BuildGraphMetrics) Fields() []ent.Field {
	return []ent.Field{

		field.Int32("action_lookup_value_count").Optional(),
		field.Int32("action_lookup_value_count_not_including_aspects").Optional(),
		field.Int32("action_count").Optional(),
		field.Int32("input_file_configured_target_count").Optional(),
		field.Int32("output_file_configured_target_count").Optional(),
		field.Int32("other_configured_target_count").Optional(),
		field.Int32("output_artifact_count").Optional(),
		field.Int32("post_invocation_skyframe_node_count").Optional(),
	}
}

// Edges of the Blob.
func (BuildGraphMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metrics", Metrics.Type).
			Ref("build_graph_metrics"),
		//there are all missing from the proto, but i'm including them here for now for completeness
		edge.To("dirtied_values", EvaluationStat.Type),
		edge.To("changed_values", EvaluationStat.Type),
		edge.To("built_values", EvaluationStat.Type),
		edge.To("cleaned_values", EvaluationStat.Type),
		edge.To("evaluated_values", EvaluationStat.Type),
	}
}
