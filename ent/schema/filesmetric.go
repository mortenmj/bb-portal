package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FilesMetric holds the schema definition for the FilesMetric entity.
type FilesMetric struct {
	ent.Schema
}

// Fields of the FilesMetric.
func (FilesMetric) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("size_in_bytes").Optional(),
		field.Int32("count").Optional(),
	}
}

// Edges of the FilesMetric.
func (FilesMetric) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("artifact_metrics", ArtifactMetrics.Type).
			Ref("source_artifacts_read").
			Ref("output_artifacts_seen").
			Ref("output_artifacts_from_action_cache").
			Ref("top_level_artifacts"),
	}
}
