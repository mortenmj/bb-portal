package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MissDetail holds the schema definition for the MissDetailMissDetail entity.
type MissDetail struct {
	ent.Schema
}

// Fields of the MissDetail.
func (MissDetail) Fields() []ent.Field {
	return []ent.Field{

		field.Enum("reason").
			Values("DIFFERENT_ACTION_KEY", "DIFFERENT_DEPS", "DIFFERENT_ENVIRONMENT", "DIFFERENT_FILES", "CORRUPTED_CACHE_ENTRY", "NOT_CACHED", "UNCONDITIONAL_EXECUTION", "UNKNOWN").
			Default("UNKNOWN").Optional(),

		field.Int32("count").Optional(),
	}
}

// Edges of the MissDetail.
func (MissDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("action_cache_statistics", ActionCacheStatistics.Type).Ref("miss_details"),
	}
}
