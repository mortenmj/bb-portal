package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MemoryMetrics holds the schema definition for the Blob entity.
type GarbageMetrics struct {
	ent.Schema
}

// Fields of the MemoryMetrics.
func (GarbageMetrics) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Optional(),
		field.Int64("garbage_collected").Optional(),
	}
}

// Edges of MemoryMetrics
func (GarbageMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("memory_metrics", MemoryMetrics.Type).Ref("garbage_metrics"),
	}
}
