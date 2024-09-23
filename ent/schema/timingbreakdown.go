package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TestResult holds the schema definition for the TestResult entity.
type TimingBreakdown struct {
	ent.Schema
}

// Fields of the TestResult.
func (TimingBreakdown) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").Optional(),
		field.String("time").Optional(),
	}
}

// Edges of TestResult
func (TimingBreakdown) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exection_info", ExectionInfo.Type).Ref("timing_breakdown"),
		edge.To("child", TimingChild.Type),
	}
}
