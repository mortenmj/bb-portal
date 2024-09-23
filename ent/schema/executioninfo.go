package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TestResult holds the schema definition for the TestResult entity.
type ExectionInfo struct {
	ent.Schema
}

// Fields of the TestResult.
func (ExectionInfo) Fields() []ent.Field {
	return []ent.Field{

		field.String("strategy").Optional(),
		field.Bool("cached_remotely").Optional(),
		field.Int32("exit_code").Optional(),
		field.String("hostname").Optional(),
	}
}

// Edges of TestResult
func (ExectionInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("test_result", TestResultBES.Type).Ref("execution_info"),
		edge.To("timing_breakdown", TimingBreakdown.Type).Unique(),
		edge.To("resource_usage", ResourceUsage.Type),
	}
}
