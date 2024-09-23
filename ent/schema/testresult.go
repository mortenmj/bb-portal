package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TestResult holds the schema definition for the TestResult entity.
type TestResult struct {
	ent.Schema
}

// Fields of the TestResult.
func (TestResult) Fields() []ent.Field {
	return []ent.Field{

		field.Enum("test_status").Optional().
			Values("NO_STATUS", "PASSED", "FLAKY", "TIMEOUT", "FAILED", "INCOMPLETE", "REMOTE_FAILURE", "FAILED_TO_BUILD", "TOOL_HALTED_BEFORE_TESTING").
			Default("NO_STATUS"),
		field.String("status_details").Optional(),
		field.Bool("cached_locally").Optional(),
		field.Int64("test_attempt_start_millis_epoch").Optional(),
		field.Int64("test_attempt_duration_millis").Optional(),
		field.Int64("targets_configured_not_including_aspects").Optional(),
		field.Strings("warning").Optional(),
		//is this a thing?
		field.Int("run").Optional(),
		field.Int("shard").Optional(),
		field.Int("attempt").Optional(),
		//is this a thing?
		field.String("label").Optional(),
	}
}

// Edges of TestResult
func (TestResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bazel_invocation", Metrics.Type).Ref("test_result").Unique(),
		edge.To("test_action_output", TestFile.Type),
		edge.To("execution_info", ExectionInfo.Type).Unique(),
	}
}
