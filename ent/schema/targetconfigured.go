package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type TargetConfigured struct {
	ent.Schema
}

// Fields of the Blob.
func (TargetConfigured) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("tag").Optional(),
		field.String("target_kind").Optional(),
		field.Int64("start_time_in_ms").Optional(),
		field.Enum("test_size").
			Values("UNKNOWN", "SMALL", "MEDIUM", "LARGE", "ENORMOUS").
			Default("UNKNOWN").Optional(),
	}
}

// Edges of the Blob.
func (TargetConfigured) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("target_pair", TargetPair.Type).
			Ref("configuration"),
	}
}
