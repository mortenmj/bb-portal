package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	"github.com/buildbarn/bb-portal/pkg/summary"
)

// BazelInvocation holds the schema definition for the BazelInvocation entity.
type BazelInvocation struct {
	ent.Schema
}

// Fields of the BazelInvocation.
func (BazelInvocation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("invocation_id", uuid.UUID{}).Unique().Immutable(),
		field.Time("started_at"),
		field.Time("ended_at").Optional(),
		// Rethink? Keep for now to capture existing processing.
		field.Int32("change_number").Optional(),
		field.Int32("patchset_number").Optional(),
		field.JSON("summary", summary.InvocationSummary{}).Annotations(entgql.Skip()), // NOTE: Internal model, not exposed to API.
		field.Bool("bep_completed").Optional(),
		field.String("step_label"),
		field.JSON("related_files", map[string]string{}).Annotations(entgql.Skip()), // NOTE: Uses custom resolver.

		//email address of the user who launched the invocation if provided
		field.String("user_email").Optional(),

		//ldap (username) of the user who launched the invocation if provided//The user who
		field.String("user_ldap").Optional(),

		//the full logs from the build
		field.String("build_logs").Optional(),

		//the cpu type from the configuration event(s)
		field.String("cpu").Optional(),

		//the platform name from the configuration event(s)
		field.String("platform_name").Optional(),

		//the name from the configuration event(s)
		field.String("configuration_mnemonic").Optional(),

		//the number of successful fetch events seen
		field.Int64("num_fetches").Optional(),
	}
}

// Edges of the BazelInvocation.
func (BazelInvocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event_file", EventFile.Type).
			Ref("bazel_invocation").
			Unique().
			Required(),
		edge.From("build", Build.Type).
			Ref("invocations").
			Unique(),

		edge.To("problems", BazelInvocationProblem.Type).
			Annotations(entgql.Skip(entgql.SkipType)), // NOTE: Uses custom resolver / types.

		//Build Metrics for the Completed Invocation
		edge.To("metrics", Metrics.Type).
			Unique(),
		//Test Data for the completed Invocation
		edge.To("test_collection", TestCollection.Type),

		//Target Data for the completed Invocation
		edge.To("targets", TargetPair.Type),
	}
}

func (BazelInvocation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("change_number", "patchset_number"),
	}
}

func (BazelInvocation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("findBazelInvocations"),
	}
}
