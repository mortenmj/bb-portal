package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RunnerCount holds the schema definition for the RunnerCount entity.
type RunnerCount struct {
	ent.Schema
}

// Fields of the RunnerCount.
func (RunnerCount) Fields() []ent.Field {
	return []ent.Field{

		//the name of the runner
		field.String("name").Optional(),

		// the execition kind (local, remote, etc)
		field.String("exec_kind").Optional(),

		//count of actions of this type executed
		field.Int64("actions_executed").Optional(),
	}
}

// Edges of the RunnerCount.
func (RunnerCount) Edges() []ent.Edge {
	return []ent.Edge{

		//edge back to the action summary
		edge.From("action_summary", ActionSummary.Type).
			Ref("runner_count"),
	}
}

// NOTE: not implemented
func (RunnerCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("findRunnerCounts"),
	}
}
