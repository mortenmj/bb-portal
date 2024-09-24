package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type OutputGroup struct {
	ent.Schema
}

// Fields of the Blob.
func (OutputGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.Bool("incomplete").Optional(),
	}
}

// Edges of the Blob.
func (OutputGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("target_complete", TargetComplete.Type).
			Ref("output_group"),
		edge.To("inline_files", TestFile.Type),
		edge.To("file_sets", NamedSetOfFiles.Type).Unique(),
	}
}
