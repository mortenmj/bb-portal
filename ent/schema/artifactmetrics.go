package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// ArtifactMetrics holds the schema definition for the ArtifactMetrics entity.
type ArtifactMetrics struct {
	ent.Schema
}

// Fields of the ArtifactMetrics.
func (ArtifactMetrics) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the ArtifactMetrics.
func (ArtifactMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metrics", Metrics.Type).Ref("artifact_metrics"),
		edge.To("source_artifacts_read", FilesMetric.Type),
		edge.To("output_artifacts_seen", FilesMetric.Type),
		edge.To("output_artifacts_from_action_cache", FilesMetric.Type),
		edge.To("top_level_artifacts", FilesMetric.Type),
	}
}
