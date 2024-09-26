package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TimingChild holds the schema definition for the TimingChild entity.
type TimingChild struct {
	ent.Schema
}

// Fields of the TimingChild.
func (TimingChild) Fields() []ent.Field {
	return []ent.Field{

		//name of the activity
		field.String("name").Optional(),

		//time spent performing the activity (duration)
		field.Int64("time").Optional(),
	}
}

// Edges of TimingChild
func (TimingChild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("timing_breakdown", TimingBreakdown.Type).Ref("child"),
	}
}
