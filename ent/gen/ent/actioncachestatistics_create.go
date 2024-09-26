// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/actioncachestatistics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/actionsummary"
	"github.com/buildbarn/bb-portal/ent/gen/ent/missdetail"
)

// ActionCacheStatisticsCreate is the builder for creating a ActionCacheStatistics entity.
type ActionCacheStatisticsCreate struct {
	config
	mutation *ActionCacheStatisticsMutation
	hooks    []Hook
}

// SetSizeInBytes sets the "size_in_bytes" field.
func (acsc *ActionCacheStatisticsCreate) SetSizeInBytes(i int64) *ActionCacheStatisticsCreate {
	acsc.mutation.SetSizeInBytes(i)
	return acsc
}

// SetNillableSizeInBytes sets the "size_in_bytes" field if the given value is not nil.
func (acsc *ActionCacheStatisticsCreate) SetNillableSizeInBytes(i *int64) *ActionCacheStatisticsCreate {
	if i != nil {
		acsc.SetSizeInBytes(*i)
	}
	return acsc
}

// SetSaveTimeInMs sets the "save_time_in_ms" field.
func (acsc *ActionCacheStatisticsCreate) SetSaveTimeInMs(i int64) *ActionCacheStatisticsCreate {
	acsc.mutation.SetSaveTimeInMs(i)
	return acsc
}

// SetNillableSaveTimeInMs sets the "save_time_in_ms" field if the given value is not nil.
func (acsc *ActionCacheStatisticsCreate) SetNillableSaveTimeInMs(i *int64) *ActionCacheStatisticsCreate {
	if i != nil {
		acsc.SetSaveTimeInMs(*i)
	}
	return acsc
}

// SetLoadTimeInMs sets the "load_time_in_ms" field.
func (acsc *ActionCacheStatisticsCreate) SetLoadTimeInMs(i int64) *ActionCacheStatisticsCreate {
	acsc.mutation.SetLoadTimeInMs(i)
	return acsc
}

// SetNillableLoadTimeInMs sets the "load_time_in_ms" field if the given value is not nil.
func (acsc *ActionCacheStatisticsCreate) SetNillableLoadTimeInMs(i *int64) *ActionCacheStatisticsCreate {
	if i != nil {
		acsc.SetLoadTimeInMs(*i)
	}
	return acsc
}

// SetHits sets the "hits" field.
func (acsc *ActionCacheStatisticsCreate) SetHits(i int32) *ActionCacheStatisticsCreate {
	acsc.mutation.SetHits(i)
	return acsc
}

// SetNillableHits sets the "hits" field if the given value is not nil.
func (acsc *ActionCacheStatisticsCreate) SetNillableHits(i *int32) *ActionCacheStatisticsCreate {
	if i != nil {
		acsc.SetHits(*i)
	}
	return acsc
}

// SetMisses sets the "misses" field.
func (acsc *ActionCacheStatisticsCreate) SetMisses(i int32) *ActionCacheStatisticsCreate {
	acsc.mutation.SetMisses(i)
	return acsc
}

// SetNillableMisses sets the "misses" field if the given value is not nil.
func (acsc *ActionCacheStatisticsCreate) SetNillableMisses(i *int32) *ActionCacheStatisticsCreate {
	if i != nil {
		acsc.SetMisses(*i)
	}
	return acsc
}

// AddActionSummaryIDs adds the "action_summary" edge to the ActionSummary entity by IDs.
func (acsc *ActionCacheStatisticsCreate) AddActionSummaryIDs(ids ...int) *ActionCacheStatisticsCreate {
	acsc.mutation.AddActionSummaryIDs(ids...)
	return acsc
}

// AddActionSummary adds the "action_summary" edges to the ActionSummary entity.
func (acsc *ActionCacheStatisticsCreate) AddActionSummary(a ...*ActionSummary) *ActionCacheStatisticsCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return acsc.AddActionSummaryIDs(ids...)
}

// AddMissDetailIDs adds the "miss_details" edge to the MissDetail entity by IDs.
func (acsc *ActionCacheStatisticsCreate) AddMissDetailIDs(ids ...int) *ActionCacheStatisticsCreate {
	acsc.mutation.AddMissDetailIDs(ids...)
	return acsc
}

// AddMissDetails adds the "miss_details" edges to the MissDetail entity.
func (acsc *ActionCacheStatisticsCreate) AddMissDetails(m ...*MissDetail) *ActionCacheStatisticsCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return acsc.AddMissDetailIDs(ids...)
}

// Mutation returns the ActionCacheStatisticsMutation object of the builder.
func (acsc *ActionCacheStatisticsCreate) Mutation() *ActionCacheStatisticsMutation {
	return acsc.mutation
}

// Save creates the ActionCacheStatistics in the database.
func (acsc *ActionCacheStatisticsCreate) Save(ctx context.Context) (*ActionCacheStatistics, error) {
	return withHooks(ctx, acsc.sqlSave, acsc.mutation, acsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acsc *ActionCacheStatisticsCreate) SaveX(ctx context.Context) *ActionCacheStatistics {
	v, err := acsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acsc *ActionCacheStatisticsCreate) Exec(ctx context.Context) error {
	_, err := acsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acsc *ActionCacheStatisticsCreate) ExecX(ctx context.Context) {
	if err := acsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acsc *ActionCacheStatisticsCreate) check() error {
	return nil
}

func (acsc *ActionCacheStatisticsCreate) sqlSave(ctx context.Context) (*ActionCacheStatistics, error) {
	if err := acsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := acsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	acsc.mutation.id = &_node.ID
	acsc.mutation.done = true
	return _node, nil
}

func (acsc *ActionCacheStatisticsCreate) createSpec() (*ActionCacheStatistics, *sqlgraph.CreateSpec) {
	var (
		_node = &ActionCacheStatistics{config: acsc.config}
		_spec = sqlgraph.NewCreateSpec(actioncachestatistics.Table, sqlgraph.NewFieldSpec(actioncachestatistics.FieldID, field.TypeInt))
	)
	if value, ok := acsc.mutation.SizeInBytes(); ok {
		_spec.SetField(actioncachestatistics.FieldSizeInBytes, field.TypeInt64, value)
		_node.SizeInBytes = value
	}
	if value, ok := acsc.mutation.SaveTimeInMs(); ok {
		_spec.SetField(actioncachestatistics.FieldSaveTimeInMs, field.TypeInt64, value)
		_node.SaveTimeInMs = value
	}
	if value, ok := acsc.mutation.LoadTimeInMs(); ok {
		_spec.SetField(actioncachestatistics.FieldLoadTimeInMs, field.TypeInt64, value)
		_node.LoadTimeInMs = value
	}
	if value, ok := acsc.mutation.Hits(); ok {
		_spec.SetField(actioncachestatistics.FieldHits, field.TypeInt32, value)
		_node.Hits = value
	}
	if value, ok := acsc.mutation.Misses(); ok {
		_spec.SetField(actioncachestatistics.FieldMisses, field.TypeInt32, value)
		_node.Misses = value
	}
	if nodes := acsc.mutation.ActionSummaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   actioncachestatistics.ActionSummaryTable,
			Columns: actioncachestatistics.ActionSummaryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(actionsummary.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := acsc.mutation.MissDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   actioncachestatistics.MissDetailsTable,
			Columns: actioncachestatistics.MissDetailsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActionCacheStatisticsCreateBulk is the builder for creating many ActionCacheStatistics entities in bulk.
type ActionCacheStatisticsCreateBulk struct {
	config
	err      error
	builders []*ActionCacheStatisticsCreate
}

// Save creates the ActionCacheStatistics entities in the database.
func (acscb *ActionCacheStatisticsCreateBulk) Save(ctx context.Context) ([]*ActionCacheStatistics, error) {
	if acscb.err != nil {
		return nil, acscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acscb.builders))
	nodes := make([]*ActionCacheStatistics, len(acscb.builders))
	mutators := make([]Mutator, len(acscb.builders))
	for i := range acscb.builders {
		func(i int, root context.Context) {
			builder := acscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActionCacheStatisticsMutation)
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
					_, err = mutators[i+1].Mutate(root, acscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acscb *ActionCacheStatisticsCreateBulk) SaveX(ctx context.Context) []*ActionCacheStatistics {
	v, err := acscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acscb *ActionCacheStatisticsCreateBulk) Exec(ctx context.Context) error {
	_, err := acscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acscb *ActionCacheStatisticsCreateBulk) ExecX(ctx context.Context) {
	if err := acscb.Exec(ctx); err != nil {
		panic(err)
	}
}
