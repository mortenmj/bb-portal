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
		field.Enum("overall_status").Optional().
			Values("NO_STATUS", "PASSED", "FLAKY", "TIMEOUT", "FAILED", "INCOMPLETE", "REMOTE_FAILURE", "FAILED_TO_BUILD", "TOOL_HALTED_BEFORE_TESTING").
			Default("NO_STATUS"),
		field.String("strategy").Optional(),
		field.Bool("cached_locally").Optional(),
		field.Bool("cached_remotely").Optional(),
		field.Int64("duration_ms").Optional(),
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
