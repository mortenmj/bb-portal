package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type SystemNetworkStats struct {
	ent.Schema
}

// Fields of the Blob.
func (SystemNetworkStats) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("bytes_sent").Optional(),
		field.Int64("bytes_recv").Optional(),
		field.Int64("packets_sent").Optional(),
		field.Int64("packets_recv").Optional(),
		field.Int64("peak_bytes_sent_per_sec").Optional(),
		field.Int64("peak_bytes_recv_per_sec").Optional(),
		field.Int64("peak_packets_sent_per_sec").Optional(),
		field.Int64("peak_packets_recv_per_sec").Optional(),
	}
}

// Edges of the Blob.
func (SystemNetworkStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("network_metrics", NetworkMetrics.Type).Ref("system_network_stats").Unique(),
	}
}
