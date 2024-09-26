package schema

import (
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

		// Package Name
		field.String("name").Optional(),

		//How long it took to load this package
		field.Int64("load_duration").
			//GoType(time.Duration(0)).
			Optional(),

		//number of targets using the package
		field.Int64("num_targets").Optional(),

		//computation steps for the package
		field.Int64("computation_steps").Optional(),

		//transitive loads
		field.Int64("num_transitive_loads").Optional(),

		//package overhead
		field.Int64("package_overhead").Optional(),
	}
}

// Edges of PackageLoadMetrics
func (PackageLoadMetrics) Edges() []ent.Edge {
	return []ent.Edge{
		//edge back to the package metrics
		edge.From("package_metrics", PackageMetrics.Type).Ref("package_load_metrics"),
	}
}
