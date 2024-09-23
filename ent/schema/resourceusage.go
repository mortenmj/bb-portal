package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type ResourceUsage struct {
	ent.Schema
}

// Fields of the Blob.
func (ResourceUsage) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("value").Optional(),
	}
}

// Edges of the Blob.
func (ResourceUsage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("execution_info", ExectionInfo.Type).
			Ref("resource_usage"),
	}
}
