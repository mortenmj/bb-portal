// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetconfigured"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetpair"
)

// TargetConfiguredUpdate is the builder for updating TargetConfigured entities.
type TargetConfiguredUpdate struct {
	config
	hooks    []Hook
	mutation *TargetConfiguredMutation
}

// Where appends a list predicates to the TargetConfiguredUpdate builder.
func (tcu *TargetConfiguredUpdate) Where(ps ...predicate.TargetConfigured) *TargetConfiguredUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetTag sets the "tag" field.
func (tcu *TargetConfiguredUpdate) SetTag(s []string) *TargetConfiguredUpdate {
	tcu.mutation.SetTag(s)
	return tcu
}

// AppendTag appends s to the "tag" field.
func (tcu *TargetConfiguredUpdate) AppendTag(s []string) *TargetConfiguredUpdate {
	tcu.mutation.AppendTag(s)
	return tcu
}

// ClearTag clears the value of the "tag" field.
func (tcu *TargetConfiguredUpdate) ClearTag() *TargetConfiguredUpdate {
	tcu.mutation.ClearTag()
	return tcu
}

// SetTargetKind sets the "target_kind" field.
func (tcu *TargetConfiguredUpdate) SetTargetKind(s string) *TargetConfiguredUpdate {
	tcu.mutation.SetTargetKind(s)
	return tcu
}

// SetNillableTargetKind sets the "target_kind" field if the given value is not nil.
func (tcu *TargetConfiguredUpdate) SetNillableTargetKind(s *string) *TargetConfiguredUpdate {
	if s != nil {
		tcu.SetTargetKind(*s)
	}
	return tcu
}

// ClearTargetKind clears the value of the "target_kind" field.
func (tcu *TargetConfiguredUpdate) ClearTargetKind() *TargetConfiguredUpdate {
	tcu.mutation.ClearTargetKind()
	return tcu
}

// SetStartTimeInMs sets the "start_time_in_ms" field.
func (tcu *TargetConfiguredUpdate) SetStartTimeInMs(i int64) *TargetConfiguredUpdate {
	tcu.mutation.ResetStartTimeInMs()
	tcu.mutation.SetStartTimeInMs(i)
	return tcu
}

// SetNillableStartTimeInMs sets the "start_time_in_ms" field if the given value is not nil.
func (tcu *TargetConfiguredUpdate) SetNillableStartTimeInMs(i *int64) *TargetConfiguredUpdate {
	if i != nil {
		tcu.SetStartTimeInMs(*i)
	}
	return tcu
}

// AddStartTimeInMs adds i to the "start_time_in_ms" field.
func (tcu *TargetConfiguredUpdate) AddStartTimeInMs(i int64) *TargetConfiguredUpdate {
	tcu.mutation.AddStartTimeInMs(i)
	return tcu
}

// ClearStartTimeInMs clears the value of the "start_time_in_ms" field.
func (tcu *TargetConfiguredUpdate) ClearStartTimeInMs() *TargetConfiguredUpdate {
	tcu.mutation.ClearStartTimeInMs()
	return tcu
}

// SetTestSize sets the "test_size" field.
func (tcu *TargetConfiguredUpdate) SetTestSize(ts targetconfigured.TestSize) *TargetConfiguredUpdate {
	tcu.mutation.SetTestSize(ts)
	return tcu
}

// SetNillableTestSize sets the "test_size" field if the given value is not nil.
func (tcu *TargetConfiguredUpdate) SetNillableTestSize(ts *targetconfigured.TestSize) *TargetConfiguredUpdate {
	if ts != nil {
		tcu.SetTestSize(*ts)
	}
	return tcu
}

// ClearTestSize clears the value of the "test_size" field.
func (tcu *TargetConfiguredUpdate) ClearTestSize() *TargetConfiguredUpdate {
	tcu.mutation.ClearTestSize()
	return tcu
}

// AddTargetPairIDs adds the "target_pair" edge to the TargetPair entity by IDs.
func (tcu *TargetConfiguredUpdate) AddTargetPairIDs(ids ...int) *TargetConfiguredUpdate {
	tcu.mutation.AddTargetPairIDs(ids...)
	return tcu
}

// AddTargetPair adds the "target_pair" edges to the TargetPair entity.
func (tcu *TargetConfiguredUpdate) AddTargetPair(t ...*TargetPair) *TargetConfiguredUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tcu.AddTargetPairIDs(ids...)
}

// Mutation returns the TargetConfiguredMutation object of the builder.
func (tcu *TargetConfiguredUpdate) Mutation() *TargetConfiguredMutation {
	return tcu.mutation
}

// ClearTargetPair clears all "target_pair" edges to the TargetPair entity.
func (tcu *TargetConfiguredUpdate) ClearTargetPair() *TargetConfiguredUpdate {
	tcu.mutation.ClearTargetPair()
	return tcu
}

// RemoveTargetPairIDs removes the "target_pair" edge to TargetPair entities by IDs.
func (tcu *TargetConfiguredUpdate) RemoveTargetPairIDs(ids ...int) *TargetConfiguredUpdate {
	tcu.mutation.RemoveTargetPairIDs(ids...)
	return tcu
}

// RemoveTargetPair removes "target_pair" edges to TargetPair entities.
func (tcu *TargetConfiguredUpdate) RemoveTargetPair(t ...*TargetPair) *TargetConfiguredUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tcu.RemoveTargetPairIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TargetConfiguredUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tcu.sqlSave, tcu.mutation, tcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TargetConfiguredUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TargetConfiguredUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TargetConfiguredUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcu *TargetConfiguredUpdate) check() error {
	if v, ok := tcu.mutation.TestSize(); ok {
		if err := targetconfigured.TestSizeValidator(v); err != nil {
			return &ValidationError{Name: "test_size", err: fmt.Errorf(`ent: validator failed for field "TargetConfigured.test_size": %w`, err)}
		}
	}
	return nil
}

func (tcu *TargetConfiguredUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(targetconfigured.Table, targetconfigured.Columns, sqlgraph.NewFieldSpec(targetconfigured.FieldID, field.TypeInt))
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.Tag(); ok {
		_spec.SetField(targetconfigured.FieldTag, field.TypeJSON, value)
	}
	if value, ok := tcu.mutation.AppendedTag(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, targetconfigured.FieldTag, value)
		})
	}
	if tcu.mutation.TagCleared() {
		_spec.ClearField(targetconfigured.FieldTag, field.TypeJSON)
	}
	if value, ok := tcu.mutation.TargetKind(); ok {
		_spec.SetField(targetconfigured.FieldTargetKind, field.TypeString, value)
	}
	if tcu.mutation.TargetKindCleared() {
		_spec.ClearField(targetconfigured.FieldTargetKind, field.TypeString)
	}
	if value, ok := tcu.mutation.StartTimeInMs(); ok {
		_spec.SetField(targetconfigured.FieldStartTimeInMs, field.TypeInt64, value)
	}
	if value, ok := tcu.mutation.AddedStartTimeInMs(); ok {
		_spec.AddField(targetconfigured.FieldStartTimeInMs, field.TypeInt64, value)
	}
	if tcu.mutation.StartTimeInMsCleared() {
		_spec.ClearField(targetconfigured.FieldStartTimeInMs, field.TypeInt64)
	}
	if value, ok := tcu.mutation.TestSize(); ok {
		_spec.SetField(targetconfigured.FieldTestSize, field.TypeEnum, value)
	}
	if tcu.mutation.TestSizeCleared() {
		_spec.ClearField(targetconfigured.FieldTestSize, field.TypeEnum)
	}
	if tcu.mutation.TargetPairCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.RemovedTargetPairIDs(); len(nodes) > 0 && !tcu.mutation.TargetPairCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.TargetPairIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{targetconfigured.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tcu.mutation.done = true
	return n, nil
}

// TargetConfiguredUpdateOne is the builder for updating a single TargetConfigured entity.
type TargetConfiguredUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TargetConfiguredMutation
}

// SetTag sets the "tag" field.
func (tcuo *TargetConfiguredUpdateOne) SetTag(s []string) *TargetConfiguredUpdateOne {
	tcuo.mutation.SetTag(s)
	return tcuo
}

// AppendTag appends s to the "tag" field.
func (tcuo *TargetConfiguredUpdateOne) AppendTag(s []string) *TargetConfiguredUpdateOne {
	tcuo.mutation.AppendTag(s)
	return tcuo
}

// ClearTag clears the value of the "tag" field.
func (tcuo *TargetConfiguredUpdateOne) ClearTag() *TargetConfiguredUpdateOne {
	tcuo.mutation.ClearTag()
	return tcuo
}

// SetTargetKind sets the "target_kind" field.
func (tcuo *TargetConfiguredUpdateOne) SetTargetKind(s string) *TargetConfiguredUpdateOne {
	tcuo.mutation.SetTargetKind(s)
	return tcuo
}

// SetNillableTargetKind sets the "target_kind" field if the given value is not nil.
func (tcuo *TargetConfiguredUpdateOne) SetNillableTargetKind(s *string) *TargetConfiguredUpdateOne {
	if s != nil {
		tcuo.SetTargetKind(*s)
	}
	return tcuo
}

// ClearTargetKind clears the value of the "target_kind" field.
func (tcuo *TargetConfiguredUpdateOne) ClearTargetKind() *TargetConfiguredUpdateOne {
	tcuo.mutation.ClearTargetKind()
	return tcuo
}

// SetStartTimeInMs sets the "start_time_in_ms" field.
func (tcuo *TargetConfiguredUpdateOne) SetStartTimeInMs(i int64) *TargetConfiguredUpdateOne {
	tcuo.mutation.ResetStartTimeInMs()
	tcuo.mutation.SetStartTimeInMs(i)
	return tcuo
}

// SetNillableStartTimeInMs sets the "start_time_in_ms" field if the given value is not nil.
func (tcuo *TargetConfiguredUpdateOne) SetNillableStartTimeInMs(i *int64) *TargetConfiguredUpdateOne {
	if i != nil {
		tcuo.SetStartTimeInMs(*i)
	}
	return tcuo
}

// AddStartTimeInMs adds i to the "start_time_in_ms" field.
func (tcuo *TargetConfiguredUpdateOne) AddStartTimeInMs(i int64) *TargetConfiguredUpdateOne {
	tcuo.mutation.AddStartTimeInMs(i)
	return tcuo
}

// ClearStartTimeInMs clears the value of the "start_time_in_ms" field.
func (tcuo *TargetConfiguredUpdateOne) ClearStartTimeInMs() *TargetConfiguredUpdateOne {
	tcuo.mutation.ClearStartTimeInMs()
	return tcuo
}

// SetTestSize sets the "test_size" field.
func (tcuo *TargetConfiguredUpdateOne) SetTestSize(ts targetconfigured.TestSize) *TargetConfiguredUpdateOne {
	tcuo.mutation.SetTestSize(ts)
	return tcuo
}

// SetNillableTestSize sets the "test_size" field if the given value is not nil.
func (tcuo *TargetConfiguredUpdateOne) SetNillableTestSize(ts *targetconfigured.TestSize) *TargetConfiguredUpdateOne {
	if ts != nil {
		tcuo.SetTestSize(*ts)
	}
	return tcuo
}

// ClearTestSize clears the value of the "test_size" field.
func (tcuo *TargetConfiguredUpdateOne) ClearTestSize() *TargetConfiguredUpdateOne {
	tcuo.mutation.ClearTestSize()
	return tcuo
}

// AddTargetPairIDs adds the "target_pair" edge to the TargetPair entity by IDs.
func (tcuo *TargetConfiguredUpdateOne) AddTargetPairIDs(ids ...int) *TargetConfiguredUpdateOne {
	tcuo.mutation.AddTargetPairIDs(ids...)
	return tcuo
}

// AddTargetPair adds the "target_pair" edges to the TargetPair entity.
func (tcuo *TargetConfiguredUpdateOne) AddTargetPair(t ...*TargetPair) *TargetConfiguredUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tcuo.AddTargetPairIDs(ids...)
}

// Mutation returns the TargetConfiguredMutation object of the builder.
func (tcuo *TargetConfiguredUpdateOne) Mutation() *TargetConfiguredMutation {
	return tcuo.mutation
}

// ClearTargetPair clears all "target_pair" edges to the TargetPair entity.
func (tcuo *TargetConfiguredUpdateOne) ClearTargetPair() *TargetConfiguredUpdateOne {
	tcuo.mutation.ClearTargetPair()
	return tcuo
}

// RemoveTargetPairIDs removes the "target_pair" edge to TargetPair entities by IDs.
func (tcuo *TargetConfiguredUpdateOne) RemoveTargetPairIDs(ids ...int) *TargetConfiguredUpdateOne {
	tcuo.mutation.RemoveTargetPairIDs(ids...)
	return tcuo
}

// RemoveTargetPair removes "target_pair" edges to TargetPair entities.
func (tcuo *TargetConfiguredUpdateOne) RemoveTargetPair(t ...*TargetPair) *TargetConfiguredUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tcuo.RemoveTargetPairIDs(ids...)
}

// Where appends a list predicates to the TargetConfiguredUpdate builder.
func (tcuo *TargetConfiguredUpdateOne) Where(ps ...predicate.TargetConfigured) *TargetConfiguredUpdateOne {
	tcuo.mutation.Where(ps...)
	return tcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TargetConfiguredUpdateOne) Select(field string, fields ...string) *TargetConfiguredUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TargetConfigured entity.
func (tcuo *TargetConfiguredUpdateOne) Save(ctx context.Context) (*TargetConfigured, error) {
	return withHooks(ctx, tcuo.sqlSave, tcuo.mutation, tcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TargetConfiguredUpdateOne) SaveX(ctx context.Context) *TargetConfigured {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TargetConfiguredUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TargetConfiguredUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcuo *TargetConfiguredUpdateOne) check() error {
	if v, ok := tcuo.mutation.TestSize(); ok {
		if err := targetconfigured.TestSizeValidator(v); err != nil {
			return &ValidationError{Name: "test_size", err: fmt.Errorf(`ent: validator failed for field "TargetConfigured.test_size": %w`, err)}
		}
	}
	return nil
}

func (tcuo *TargetConfiguredUpdateOne) sqlSave(ctx context.Context) (_node *TargetConfigured, err error) {
	if err := tcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(targetconfigured.Table, targetconfigured.Columns, sqlgraph.NewFieldSpec(targetconfigured.FieldID, field.TypeInt))
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TargetConfigured.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, targetconfigured.FieldID)
		for _, f := range fields {
			if !targetconfigured.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != targetconfigured.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.Tag(); ok {
		_spec.SetField(targetconfigured.FieldTag, field.TypeJSON, value)
	}
	if value, ok := tcuo.mutation.AppendedTag(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, targetconfigured.FieldTag, value)
		})
	}
	if tcuo.mutation.TagCleared() {
		_spec.ClearField(targetconfigured.FieldTag, field.TypeJSON)
	}
	if value, ok := tcuo.mutation.TargetKind(); ok {
		_spec.SetField(targetconfigured.FieldTargetKind, field.TypeString, value)
	}
	if tcuo.mutation.TargetKindCleared() {
		_spec.ClearField(targetconfigured.FieldTargetKind, field.TypeString)
	}
	if value, ok := tcuo.mutation.StartTimeInMs(); ok {
		_spec.SetField(targetconfigured.FieldStartTimeInMs, field.TypeInt64, value)
	}
	if value, ok := tcuo.mutation.AddedStartTimeInMs(); ok {
		_spec.AddField(targetconfigured.FieldStartTimeInMs, field.TypeInt64, value)
	}
	if tcuo.mutation.StartTimeInMsCleared() {
		_spec.ClearField(targetconfigured.FieldStartTimeInMs, field.TypeInt64)
	}
	if value, ok := tcuo.mutation.TestSize(); ok {
		_spec.SetField(targetconfigured.FieldTestSize, field.TypeEnum, value)
	}
	if tcuo.mutation.TestSizeCleared() {
		_spec.ClearField(targetconfigured.FieldTestSize, field.TypeEnum)
	}
	if tcuo.mutation.TargetPairCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.RemovedTargetPairIDs(); len(nodes) > 0 && !tcuo.mutation.TargetPairCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.TargetPairIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TargetConfigured{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{targetconfigured.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tcuo.mutation.done = true
	return _node, nil
}