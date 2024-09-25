package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TargetPair holds the schema definition for the TargetPair entity.
type TargetPair struct {
	ent.Schema
}

// Fields of the TargetPair.
func (TargetPair) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").Optional(),
		field.Int64("duration_in_ms").Optional(),
		//duplicates data from the edges to this target pair object to try to speed up the queries
		field.Bool("success").Optional().Default(false),
		field.String("target_kind").Optional(),
		field.Enum("test_size").
			Values("UNKNOWN",
				"SMALL",
				"MEDIUM",
				"LARGE",
				"ENORMOUS").
			Default("UNKNOWN").
			Optional(),
		field.Enum("abort_reason").
			Values("UNKNOWN",
				"USER_INTERRUPTED",
				"NO_ANALYZE",
				"NO_BUILD",
				"TIME_OUT",
				"REMOTE_ENVIRONMENT_FAILURE",
				"INTERNAL",
				"LOADING_FAILURE",
				"ANALYSIS_FAILURE",
				"SKIPPED",
				"INCOMPLETE",
				"OUT_OF_MEMORY").
			Optional(),
	}
}

// Edges of the TargetPair.
func (TargetPair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bazel_invocation", BazelInvocation.Type).
			Ref("targets"),
		edge.To("configuration", TargetConfigured.Type).Unique(),
		edge.To("completion", TargetComplete.Type).Unique(),
	}
}
