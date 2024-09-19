// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/packageloadmetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/packagemetrics"
)

// PackageLoadMetricsCreate is the builder for creating a PackageLoadMetrics entity.
type PackageLoadMetricsCreate struct {
	config
	mutation *PackageLoadMetricsMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (plmc *PackageLoadMetricsCreate) SetName(s string) *PackageLoadMetricsCreate {
	plmc.mutation.SetName(s)
	return plmc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (plmc *PackageLoadMetricsCreate) SetNillableName(s *string) *PackageLoadMetricsCreate {
	if s != nil {
		plmc.SetName(*s)
	}
	return plmc
}

// SetLoadDuration sets the "load_duration" field.
func (plmc *PackageLoadMetricsCreate) SetLoadDuration(i int64) *PackageLoadMetricsCreate {
	plmc.mutation.SetLoadDuration(i)
	return plmc
}

// SetNillableLoadDuration sets the "load_duration" field if the given value is not nil.
func (plmc *PackageLoadMetricsCreate) SetNillableLoadDuration(i *int64) *PackageLoadMetricsCreate {
	if i != nil {
		plmc.SetLoadDuration(*i)
	}
	return plmc
}

// SetNumTargets sets the "num_targets" field.
func (plmc *PackageLoadMetricsCreate) SetNumTargets(i int64) *PackageLoadMetricsCreate {
	plmc.mutation.SetNumTargets(i)
	return plmc
}

// SetNillableNumTargets sets the "num_targets" field if the given value is not nil.
func (plmc *PackageLoadMetricsCreate) SetNillableNumTargets(i *int64) *PackageLoadMetricsCreate {
	if i != nil {
		plmc.SetNumTargets(*i)
	}
	return plmc
}

// SetComputationSteps sets the "computation_steps" field.
func (plmc *PackageLoadMetricsCreate) SetComputationSteps(i int64) *PackageLoadMetricsCreate {
	plmc.mutation.SetComputationSteps(i)
	return plmc
}

// SetNillableComputationSteps sets the "computation_steps" field if the given value is not nil.
func (plmc *PackageLoadMetricsCreate) SetNillableComputationSteps(i *int64) *PackageLoadMetricsCreate {
	if i != nil {
		plmc.SetComputationSteps(*i)
	}
	return plmc
}

// SetNumTransitiveLoads sets the "num_transitive_loads" field.
func (plmc *PackageLoadMetricsCreate) SetNumTransitiveLoads(i int64) *PackageLoadMetricsCreate {
	plmc.mutation.SetNumTransitiveLoads(i)
	return plmc
}

// SetNillableNumTransitiveLoads sets the "num_transitive_loads" field if the given value is not nil.
func (plmc *PackageLoadMetricsCreate) SetNillableNumTransitiveLoads(i *int64) *PackageLoadMetricsCreate {
	if i != nil {
		plmc.SetNumTransitiveLoads(*i)
	}
	return plmc
}

// SetPackageOverhead sets the "package_overhead" field.
func (plmc *PackageLoadMetricsCreate) SetPackageOverhead(i int64) *PackageLoadMetricsCreate {
	plmc.mutation.SetPackageOverhead(i)
	return plmc
}

// SetNillablePackageOverhead sets the "package_overhead" field if the given value is not nil.
func (plmc *PackageLoadMetricsCreate) SetNillablePackageOverhead(i *int64) *PackageLoadMetricsCreate {
	if i != nil {
		plmc.SetPackageOverhead(*i)
	}
	return plmc
}

// AddPackageMetricIDs adds the "package_metrics" edge to the PackageMetrics entity by IDs.
func (plmc *PackageLoadMetricsCreate) AddPackageMetricIDs(ids ...int) *PackageLoadMetricsCreate {
	plmc.mutation.AddPackageMetricIDs(ids...)
	return plmc
}

// AddPackageMetrics adds the "package_metrics" edges to the PackageMetrics entity.
func (plmc *PackageLoadMetricsCreate) AddPackageMetrics(p ...*PackageMetrics) *PackageLoadMetricsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plmc.AddPackageMetricIDs(ids...)
}

// Mutation returns the PackageLoadMetricsMutation object of the builder.
func (plmc *PackageLoadMetricsCreate) Mutation() *PackageLoadMetricsMutation {
	return plmc.mutation
}

// Save creates the PackageLoadMetrics in the database.
func (plmc *PackageLoadMetricsCreate) Save(ctx context.Context) (*PackageLoadMetrics, error) {
	return withHooks(ctx, plmc.sqlSave, plmc.mutation, plmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (plmc *PackageLoadMetricsCreate) SaveX(ctx context.Context) *PackageLoadMetrics {
	v, err := plmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (plmc *PackageLoadMetricsCreate) Exec(ctx context.Context) error {
	_, err := plmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plmc *PackageLoadMetricsCreate) ExecX(ctx context.Context) {
	if err := plmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plmc *PackageLoadMetricsCreate) check() error {
	return nil
}

func (plmc *PackageLoadMetricsCreate) sqlSave(ctx context.Context) (*PackageLoadMetrics, error) {
	if err := plmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := plmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, plmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	plmc.mutation.id = &_node.ID
	plmc.mutation.done = true
	return _node, nil
}

func (plmc *PackageLoadMetricsCreate) createSpec() (*PackageLoadMetrics, *sqlgraph.CreateSpec) {
	var (
		_node = &PackageLoadMetrics{config: plmc.config}
		_spec = sqlgraph.NewCreateSpec(packageloadmetrics.Table, sqlgraph.NewFieldSpec(packageloadmetrics.FieldID, field.TypeInt))
	)
	if value, ok := plmc.mutation.Name(); ok {
		_spec.SetField(packageloadmetrics.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := plmc.mutation.LoadDuration(); ok {
		_spec.SetField(packageloadmetrics.FieldLoadDuration, field.TypeInt64, value)
		_node.LoadDuration = value
	}
	if value, ok := plmc.mutation.NumTargets(); ok {
		_spec.SetField(packageloadmetrics.FieldNumTargets, field.TypeInt64, value)
		_node.NumTargets = value
	}
	if value, ok := plmc.mutation.ComputationSteps(); ok {
		_spec.SetField(packageloadmetrics.FieldComputationSteps, field.TypeInt64, value)
		_node.ComputationSteps = value
	}
	if value, ok := plmc.mutation.NumTransitiveLoads(); ok {
		_spec.SetField(packageloadmetrics.FieldNumTransitiveLoads, field.TypeInt64, value)
		_node.NumTransitiveLoads = value
	}
	if value, ok := plmc.mutation.PackageOverhead(); ok {
		_spec.SetField(packageloadmetrics.FieldPackageOverhead, field.TypeInt64, value)
		_node.PackageOverhead = value
	}
	if nodes := plmc.mutation.PackageMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   packageloadmetrics.PackageMetricsTable,
			Columns: packageloadmetrics.PackageMetricsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagemetrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PackageLoadMetricsCreateBulk is the builder for creating many PackageLoadMetrics entities in bulk.
type PackageLoadMetricsCreateBulk struct {
	config
	err      error
	builders []*PackageLoadMetricsCreate
}

// Save creates the PackageLoadMetrics entities in the database.
func (plmcb *PackageLoadMetricsCreateBulk) Save(ctx context.Context) ([]*PackageLoadMetrics, error) {
	if plmcb.err != nil {
		return nil, plmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(plmcb.builders))
	nodes := make([]*PackageLoadMetrics, len(plmcb.builders))
	mutators := make([]Mutator, len(plmcb.builders))
	for i := range plmcb.builders {
		func(i int, root context.Context) {
			builder := plmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PackageLoadMetricsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, plmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, plmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, plmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (plmcb *PackageLoadMetricsCreateBulk) SaveX(ctx context.Context) []*PackageLoadMetrics {
	v, err := plmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (plmcb *PackageLoadMetricsCreateBulk) Exec(ctx context.Context) error {
	_, err := plmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plmcb *PackageLoadMetricsCreateBulk) ExecX(ctx context.Context) {
	if err := plmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
