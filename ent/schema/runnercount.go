package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type RunnerCount struct {
	ent.Schema
}

// Fields of the Blob.
func (RunnerCount) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("exec_kind").Optional(),
		field.Int64("actions_executed").Optional(),
	}
}

// Edges of the Blob.
func (RunnerCount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("action_summary", ActionSummary.Type).
			Ref("runner_count"),
	}
}

func (RunnerCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("findRunnerCounts"),
	}
}
