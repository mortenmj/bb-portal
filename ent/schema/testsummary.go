package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type TestSummary struct {
	ent.Schema
}

// Fields of the Blob.
func (TestSummary) Fields() []ent.Field {
	return []ent.Field{

		field.Enum("overall_status").Optional().
			Values("NO_STATUS", "PASSED", "FLAKY", "TIMEOUT", "FAILED", "INCOMPLETE", "REMOTE_FAILURE", "FAILED_TO_BUILD", "TOOL_HALTED_BEFORE_TESTING").
			Default("NO_STATUS"),

		field.Int32("total_run_count").Optional(),
		field.Int32("run_count").Optional(),
		field.Int32("attempt_count").Optional(),
		field.Int32("shard_count").Optional(),
		field.Int32("total_num_cached").Optional(),
		field.Int64("first_start_time").Optional(),
		field.Int64("last_stop_time").Optional(),
		field.Int64("total_run_duration").Optional(),
		field.String("label").Optional(),
	}
}

// Edges of the Blob.
func (TestSummary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("test_collection", TestCollection.Type).
			Ref("test_summary"),
		edge.To("passed", TestFile.Type),
		edge.To("failed", TestFile.Type),
	}
}
