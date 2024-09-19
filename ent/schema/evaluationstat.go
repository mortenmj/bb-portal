package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type EvaluationStat struct {
	ent.Schema
}

// Fields of the Blob.
func (EvaluationStat) Fields() []ent.Field {
	return []ent.Field{

		field.String("skyfunction_name").Optional(),
		field.Int64("count").Optional(),
	}
}

// Edges of the Blob.
func (EvaluationStat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("build_graph_metrics", BuildGraphMetrics.Type).
			Ref("dirtied_values").
			Ref("changed_values").
			Ref("built_values").
			Ref("cleaned_values").
			Ref("evaluated_values"),
	}
}
