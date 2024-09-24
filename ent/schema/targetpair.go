package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type TargetPair struct {
	ent.Schema
}

// Fields of the Blob.
func (TargetPair) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").Optional(),
		field.Int64("duration_in_ms").Optional(),
	}
}

// Edges of the Blob.
func (TargetPair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bazel_invocation", BazelInvocation.Type).
			Ref("test_collection"),
		edge.To("configuration", TargetConfigured.Type).Unique(),
		edge.To("completion", TargetComplete.Type).Unique(),
	}
}
