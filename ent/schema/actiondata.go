package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ActionData holds the schema definition for the ActionData entity.
type ActionData struct {
	ent.Schema
}

// Fields of the ActionData.
func (ActionData) Fields() []ent.Field {
	return []ent.Field{

		//the action name
		field.String("mnemonic").Optional(),

		// The total number of actions of this type executed during the build. As
		// above, includes remote cache hits but excludes local action cache hits.
		field.Int64("actions_executed").Optional(),

		// The total number of actions of this type registered during the build.
		field.Int64("actions_created").Optional(),

		// When the first action of this type started being executed, in
		// milliseconds from the epoch.
		field.Int64("first_started_ms").Optional(),

		// When the last action of this type ended being executed, in
		// milliseconds from the epoch.
		field.Int64("last_ended_ms").Optional(),

		// Accumulated CPU time of all spawned actions of this type.
		// This is only set if all the actions reported a time
		field.Int64("system_time").
			//GoType(time.Duration(0)).
			Optional(),
		field.Int64("user_time").
			//GoType(time.Duration(0)).
			Optional(),
	}
}

// Edges of the ActionData.
func (ActionData) Edges() []ent.Edge {
	return []ent.Edge{
		//edge back to the associated action summary
		edge.From("action_summary", ActionSummary.Type).
			Ref("action_data"),
	}
}
