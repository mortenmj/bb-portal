package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type RaceStatistics struct {
	ent.Schema
}

// Fields of the Blob.
func (RaceStatistics) Fields() []ent.Field {
	return []ent.Field{

		field.String("mnemonic").Optional(),
		field.String("local_runner").Optional(),
		field.String("remote_runner").Optional(),
		field.Int32("local_wins").Optional(),
		field.Int64("renote_wins").Optional(),
	}
}

// Edges of the Blob.
func (RaceStatistics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dynamic_execution_metrics", DynamicExecutionMetrics.Type).Ref("race_statistics"),
	}
}
