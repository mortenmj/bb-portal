package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ResourceUsage holds the schema definition for the ResourceUsage entity.
type ResourceUsage struct {
	ent.Schema
}

// Fields of the ResourceUsage.
func (ResourceUsage) Fields() []ent.Field {
	return []ent.Field{

		//NOTE: not currently implemented on the proto but included here for completeness

		//the name
		field.String("name").Optional(),

		//the value?
		field.String("value").Optional(),
	}
}

// Edges of the ResourceUsage.
func (ResourceUsage) Edges() []ent.Edge {
	return []ent.Edge{

		//edge back to the execution info
		edge.From("execution_info", ExectionInfo.Type).
			Ref("resource_usage"),
	}
}
