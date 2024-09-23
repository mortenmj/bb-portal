package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TestResult holds the schema definition for the TestResult entity.
type TimingChild struct {
	ent.Schema
}

// Fields of the TestResult.
func (TimingChild) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").Optional(),
		field.String("time").Optional(),
	}
}

// Edges of TestResult
func (TimingChild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("timing_breakdown", TimingBreakdown.Type).Ref("child"),
	}
}
