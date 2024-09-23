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
		field.String("digest").Optional(),
		field.String("file").Optional(),
		field.Int64("length").Optional(),
		field.String("name").Optional(),
		field.Strings("prefix").Optional(),
	}
}

// Edges of TestResult
func (TestFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("test_result", TestResultBES.Type).Ref("test_action_output"),
	}
}
