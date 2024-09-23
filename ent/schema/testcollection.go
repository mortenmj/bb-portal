package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type TestCollection struct {
	ent.Schema
}

// Fields of the Blob.
func (TestCollection) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").Optional(),
	}
}

// Edges of the Blob.
func (TestCollection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bazel_invocation", BazelInvocation.Type).
			Ref("test_collection"),
		edge.To("test_summary", TestSummary.Type).Unique(),
		edge.To("test_results", TestResultBES.Type),
	}
}
