package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blob holds the schema definition for the Blob entity.
type ActionSummary struct {
	ent.Schema
}

// Fields of the Blob.
func (ActionSummary) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("actions_created").Optional(),
		field.Int64("actions_created_not_including_aspects").Optional(),
		field.Int64("actions_executed").Optional(),
		field.Int64("remote_cache_hits").Optional(), //DEPRECATED
	}
}

// Edges of the Blob.
func (ActionSummary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("action_data", ActionData.Type),
		edge.To("runner_count", RunnerCount.Type),
		edge.To("action_cache_statistics", ActionCacheStatistics.Type),
		edge.From("metrics", Metrics.Type).Ref("action_summary").Unique(),
	}
}
