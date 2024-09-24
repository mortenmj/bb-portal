package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Blob holds the schema definition for the Blob entity.
type NamedSetOfFiles struct {
	ent.Schema
}

// Fields of the Blob.
func (NamedSetOfFiles) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Blob.
func (NamedSetOfFiles) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("output_group", OutputGroup.Type).
			Ref("file_sets"),
		edge.To("files", TestFile.Type),
		edge.To("file_sets", NamedSetOfFiles.Type).Unique(),
	}
}
