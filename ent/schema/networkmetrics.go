package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Blob holds the schema definition for the Blob entity.
type NetworkMetrics struct {
	ent.Schema
}

// Fields of the Blob.
func (NetworkMetrics) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Blob.
func (NetworkMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metrics", Metrics.Type).Ref("network_metrics"),
		edge.To("system_network_stats", SystemNetworkStats.Type),
	}
}
