package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type TargetComplete struct {
	ent.Schema
}

// Fields of the Blob.
func (TargetComplete) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("success").Optional(),
		field.Strings("tag").Optional(),
		field.String("target_kind").Optional(),
		field.Int64("end_time_in_ms").Optional(),
		field.Int64("test_timeout_seconds").Optional(),
		field.Int64("test_timeout").Optional(),
		field.Enum("test_size").
			Values("UNKNOWN", "SMALL", "MEDIUM", "LARGE", "ENORMOUS").
			Default("UNKNOWN").Optional(),
		//TODO: implement OutputGroup, ImportantOutput, DirectoryOutput
	}
}

// Edges of the Blob.
func (TargetComplete) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("target_pair", TargetPair.Type).
			Ref("completion"),
		edge.To("important_output", TestFile.Type),
		edge.To("directory_output", TestFile.Type),
		edge.To("output_group", OutputGroup.Type).Unique(),
	}
}
