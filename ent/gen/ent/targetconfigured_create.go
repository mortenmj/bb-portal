// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetconfigured"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetpair"
)

// TargetConfiguredCreate is the builder for creating a TargetConfigured entity.
type TargetConfiguredCreate struct {
	config
	mutation *TargetConfiguredMutation
	hooks    []Hook
}

// SetTag sets the "tag" field.
func (tcc *TargetConfiguredCreate) SetTag(s []string) *TargetConfiguredCreate {
	tcc.mutation.SetTag(s)
	return tcc
}

// SetTargetKind sets the "target_kind" field.
func (tcc *TargetConfiguredCreate) SetTargetKind(s string) *TargetConfiguredCreate {
	tcc.mutation.SetTargetKind(s)
	return tcc
}

// SetNillableTargetKind sets the "target_kind" field if the given value is not nil.
func (tcc *TargetConfiguredCreate) SetNillableTargetKind(s *string) *TargetConfiguredCreate {
	if s != nil {
		tcc.SetTargetKind(*s)
	}
	return tcc
}

// SetStartTimeInMs sets the "start_time_in_ms" field.
func (tcc *TargetConfiguredCreate) SetStartTimeInMs(i int64) *TargetConfiguredCreate {
	tcc.mutation.SetStartTimeInMs(i)
	return tcc
}

// SetNillableStartTimeInMs sets the "start_time_in_ms" field if the given value is not nil.
func (tcc *TargetConfiguredCreate) SetNillableStartTimeInMs(i *int64) *TargetConfiguredCreate {
	if i != nil {
		tcc.SetStartTimeInMs(*i)
	}
	return tcc
}

// SetTestSize sets the "test_size" field.
func (tcc *TargetConfiguredCreate) SetTestSize(ts targetconfigured.TestSize) *TargetConfiguredCreate {
	tcc.mutation.SetTestSize(ts)
	return tcc
}

// SetNillableTestSize sets the "test_size" field if the given value is not nil.
func (tcc *TargetConfiguredCreate) SetNillableTestSize(ts *targetconfigured.TestSize) *TargetConfiguredCreate {
	if ts != nil {
		tcc.SetTestSize(*ts)
	}
	return tcc
}

// AddTargetPairIDs adds the "target_pair" edge to the TargetPair entity by IDs.
func (tcc *TargetConfiguredCreate) AddTargetPairIDs(ids ...int) *TargetConfiguredCreate {
	tcc.mutation.AddTargetPairIDs(ids...)
	return tcc
}

// AddTargetPair adds the "target_pair" edges to the TargetPair entity.
func (tcc *TargetConfiguredCreate) AddTargetPair(t ...*TargetPair) *TargetConfiguredCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tcc.AddTargetPairIDs(ids...)
}

// Mutation returns the TargetConfiguredMutation object of the builder.
func (tcc *TargetConfiguredCreate) Mutation() *TargetConfiguredMutation {
	return tcc.mutation
}

// Save creates the TargetConfigured in the database.
func (tcc *TargetConfiguredCreate) Save(ctx context.Context) (*TargetConfigured, error) {
	tcc.defaults()
	return withHooks(ctx, tcc.sqlSave, tcc.mutation, tcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tcc *TargetConfiguredCreate) SaveX(ctx context.Context) *TargetConfigured {
	v, err := tcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcc *TargetConfiguredCreate) Exec(ctx context.Context) error {
	_, err := tcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcc *TargetConfiguredCreate) ExecX(ctx context.Context) {
	if err := tcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcc *TargetConfiguredCreate) defaults() {
	if _, ok := tcc.mutation.TestSize(); !ok {
		v := targetconfigured.DefaultTestSize
		tcc.mutation.SetTestSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcc *TargetConfiguredCreate) check() error {
	if v, ok := tcc.mutation.TestSize(); ok {
		if err := targetconfigured.TestSizeValidator(v); err != nil {
			return &ValidationError{Name: "test_size", err: fmt.Errorf(`ent: validator failed for field "TargetConfigured.test_size": %w`, err)}
		}
	}
	return nil
}

func (tcc *TargetConfiguredCreate) sqlSave(ctx context.Context) (*TargetConfigured, error) {
	if err := tcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tcc.mutation.id = &_node.ID
	tcc.mutation.done = true
	return _node, nil
}

func (tcc *TargetConfiguredCreate) createSpec() (*TargetConfigured, *sqlgraph.CreateSpec) {
	var (
		_node = &TargetConfigured{config: tcc.config}
		_spec = sqlgraph.NewCreateSpec(targetconfigured.Table, sqlgraph.NewFieldSpec(targetconfigured.FieldID, field.TypeInt))
	)
	if value, ok := tcc.mutation.Tag(); ok {
		_spec.SetField(targetconfigured.FieldTag, field.TypeJSON, value)
		_node.Tag = value
	}
	if value, ok := tcc.mutation.TargetKind(); ok {
		_spec.SetField(targetconfigured.FieldTargetKind, field.TypeString, value)
		_node.TargetKind = value
	}
	if value, ok := tcc.mutation.StartTimeInMs(); ok {
		_spec.SetField(targetconfigured.FieldStartTimeInMs, field.TypeInt64, value)
		_node.StartTimeInMs = value
	}
	if value, ok := tcc.mutation.TestSize(); ok {
		_spec.SetField(targetconfigured.FieldTestSize, field.TypeEnum, value)
		_node.TestSize = value
	}
	if nodes := tcc.mutation.TargetPairIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   targetconfigured.TargetPairTable,
			Columns: []string{targetconfigured.TargetPairColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetpair.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TargetConfiguredCreateBulk is the builder for creating many TargetConfigured entities in bulk.
type TargetConfiguredCreateBulk struct {
	config
	err      error
	builders []*TargetConfiguredCreate
}

// Save creates the TargetConfigured entities in the database.
func (tccb *TargetConfiguredCreateBulk) Save(ctx context.Context) ([]*TargetConfigured, error) {
	if tccb.err != nil {
		return nil, tccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tccb.builders))
	nodes := make([]*TargetConfigured, len(tccb.builders))
	mutators := make([]Mutator, len(tccb.builders))
	for i := range tccb.builders {
		func(i int, root context.Context) {
			builder := tccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TargetConfiguredMutation)
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
					_, err = mutators[i+1].Mutate(root, tccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tccb *TargetConfiguredCreateBulk) SaveX(ctx context.Context) []*TargetConfigured {
	v, err := tccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tccb *TargetConfiguredCreateBulk) Exec(ctx context.Context) error {
	_, err := tccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tccb *TargetConfiguredCreateBulk) ExecX(ctx context.Context) {
	if err := tccb.Exec(ctx); err != nil {
		panic(err)
	}
}
