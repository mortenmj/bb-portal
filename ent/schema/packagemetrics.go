package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PackageMetrics holds the schema definition for the Blob entity.
type PackageMetrics struct {
	ent.Schema
}

// Fields of the PackageMetrics.
func (PackageMetrics) Fields() []ent.Field {
	return []ent.Field{

		// Size of the JVM heap post build in bytes. This is only collected if
		// --memory_profile is set, since it forces a full GC.
		field.Int64("packages_loaded").Optional(),
	}
}

// Edges of PackageMetrics
func (PackageMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("package_load_metrics", PackageLoadMetrics.Type),
		edge.From("metrics", Metrics.Type).Ref("package_metrics"),
	}
}
