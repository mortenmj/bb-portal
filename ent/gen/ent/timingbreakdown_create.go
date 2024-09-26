// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/exectioninfo"
	"github.com/buildbarn/bb-portal/ent/gen/ent/timingbreakdown"
	"github.com/buildbarn/bb-portal/ent/gen/ent/timingchild"
)

// TimingBreakdownCreate is the builder for creating a TimingBreakdown entity.
type TimingBreakdownCreate struct {
	config
	mutation *TimingBreakdownMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tbc *TimingBreakdownCreate) SetName(s string) *TimingBreakdownCreate {
	tbc.mutation.SetName(s)
	return tbc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tbc *TimingBreakdownCreate) SetNillableName(s *string) *TimingBreakdownCreate {
	if s != nil {
		tbc.SetName(*s)
	}
	return tbc
}

// SetTime sets the "time" field.
func (tbc *TimingBreakdownCreate) SetTime(s string) *TimingBreakdownCreate {
	tbc.mutation.SetTime(s)
	return tbc
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (tbc *TimingBreakdownCreate) SetNillableTime(s *string) *TimingBreakdownCreate {
	if s != nil {
		tbc.SetTime(*s)
	}
	return tbc
}

// AddExectionInfoIDs adds the "exection_info" edge to the ExectionInfo entity by IDs.
func (tbc *TimingBreakdownCreate) AddExectionInfoIDs(ids ...int) *TimingBreakdownCreate {
	tbc.mutation.AddExectionInfoIDs(ids...)
	return tbc
}

// AddExectionInfo adds the "exection_info" edges to the ExectionInfo entity.
func (tbc *TimingBreakdownCreate) AddExectionInfo(e ...*ExectionInfo) *TimingBreakdownCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tbc.AddExectionInfoIDs(ids...)
}

// AddChildIDs adds the "child" edge to the TimingChild entity by IDs.
func (tbc *TimingBreakdownCreate) AddChildIDs(ids ...int) *TimingBreakdownCreate {
	tbc.mutation.AddChildIDs(ids...)
	return tbc
}

// AddChild adds the "child" edges to the TimingChild entity.
func (tbc *TimingBreakdownCreate) AddChild(t ...*TimingChild) *TimingBreakdownCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tbc.AddChildIDs(ids...)
}

// Mutation returns the TimingBreakdownMutation object of the builder.
func (tbc *TimingBreakdownCreate) Mutation() *TimingBreakdownMutation {
	return tbc.mutation
}

// Save creates the TimingBreakdown in the database.
func (tbc *TimingBreakdownCreate) Save(ctx context.Context) (*TimingBreakdown, error) {
	return withHooks(ctx, tbc.sqlSave, tbc.mutation, tbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tbc *TimingBreakdownCreate) SaveX(ctx context.Context) *TimingBreakdown {
	v, err := tbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tbc *TimingBreakdownCreate) Exec(ctx context.Context) error {
	_, err := tbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tbc *TimingBreakdownCreate) ExecX(ctx context.Context) {
	if err := tbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tbc *TimingBreakdownCreate) check() error {
	return nil
}

func (tbc *TimingBreakdownCreate) sqlSave(ctx context.Context) (*TimingBreakdown, error) {
	if err := tbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tbc.mutation.id = &_node.ID
	tbc.mutation.done = true
	return _node, nil
}

func (tbc *TimingBreakdownCreate) createSpec() (*TimingBreakdown, *sqlgraph.CreateSpec) {
	var (
		_node = &TimingBreakdown{config: tbc.config}
		_spec = sqlgraph.NewCreateSpec(timingbreakdown.Table, sqlgraph.NewFieldSpec(timingbreakdown.FieldID, field.TypeInt))
	)
	if value, ok := tbc.mutation.Name(); ok {
		_spec.SetField(timingbreakdown.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tbc.mutation.Time(); ok {
		_spec.SetField(timingbreakdown.FieldTime, field.TypeString, value)
		_node.Time = value
	}
	if nodes := tbc.mutation.ExectionInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   timingbreakdown.ExectionInfoTable,
			Columns: []string{timingbreakdown.ExectionInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exectioninfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tbc.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timingbreakdown.ChildTable,
			Columns: timingbreakdown.ChildPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timingchild.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TimingBreakdownCreateBulk is the builder for creating many TimingBreakdown entities in bulk.
type TimingBreakdownCreateBulk struct {
	config
	err      error
	builders []*TimingBreakdownCreate
}

// Save creates the TimingBreakdown entities in the database.
func (tbcb *TimingBreakdownCreateBulk) Save(ctx context.Context) ([]*TimingBreakdown, error) {
	if tbcb.err != nil {
		return nil, tbcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tbcb.builders))
	nodes := make([]*TimingBreakdown, len(tbcb.builders))
	mutators := make([]Mutator, len(tbcb.builders))
	for i := range tbcb.builders {
		func(i int, root context.Context) {
			builder := tbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimingBreakdownMutation)
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
					_, err = mutators[i+1].Mutate(root, tbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tbcb *TimingBreakdownCreateBulk) SaveX(ctx context.Context) []*TimingBreakdown {
	v, err := tbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tbcb *TimingBreakdownCreateBulk) Exec(ctx context.Context) error {
	_, err := tbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tbcb *TimingBreakdownCreateBulk) ExecX(ctx context.Context) {
	if err := tbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
