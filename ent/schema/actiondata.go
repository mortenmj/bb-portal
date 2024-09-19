package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type ActionData struct {
	ent.Schema
}

// Fields of the Blob.
func (ActionData) Fields() []ent.Field {
	return []ent.Field{

		field.String("mnemonic").Optional(),
		field.Int64("actions_executed").Optional(),
		field.Int64("actions_created").Optional(),
		field.Int64("first_started_ms").Optional(),
		field.Int64("last_ended_ms").Optional(),
		field.Int64("system_time").
			//GoType(time.Duration(0)).
			Optional(),
		field.Int64("user_time").
			//GoType(time.Duration(0)).
			Optional(),
	}
}

// Edges of the Blob.
func (ActionData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("action_summary", ActionSummary.Type).
			Ref("action_data"),
	}
}
