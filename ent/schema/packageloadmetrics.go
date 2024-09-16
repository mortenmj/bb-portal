package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PackageLoadMetrics holds the schema definition for the Blob entity.
type PackageLoadMetrics struct {
	ent.Schema
}

// Fields of the PackageLoadMetrics.
func (PackageLoadMetrics) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.Int64("load_duration").
			GoType(time.Duration(0)).
			Optional(),
		field.Int64("num_targets").Optional(),
		field.Int64("computation_steps").Optional(),
		field.Int64("num_transitive_loads").Optional(),
		field.Int64("package_overhead").Optional(),
	}
}

// Edges of PackageLoadMetrics
func (PackageLoadMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("package_metrics", PackageMetrics.Type).Ref("package_load_metrics"),
	}
}
