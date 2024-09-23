package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TestResult holds the schema definition for the TestResult entity.
type TestFile struct {
	ent.Schema
}

// Fields of the TestResult.
func (TestFile) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").Optional(),
		field.String("uri").Optional(),
	}
}

// Edges of TestResult
func (TestFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("test_result", TestResult.Type).Ref("test_action_output"),
	}
}
