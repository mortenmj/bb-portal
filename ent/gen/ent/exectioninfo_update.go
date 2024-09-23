// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/exectioninfo"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/resourceusage"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testresultbes"
	"github.com/buildbarn/bb-portal/ent/gen/ent/timingbreakdown"
)

// ExectionInfoUpdate is the builder for updating ExectionInfo entities.
type ExectionInfoUpdate struct {
	config
	hooks    []Hook
	mutation *ExectionInfoMutation
}

// Where appends a list predicates to the ExectionInfoUpdate builder.
func (eiu *ExectionInfoUpdate) Where(ps ...predicate.ExectionInfo) *ExectionInfoUpdate {
	eiu.mutation.Where(ps...)
	return eiu
}

// SetStrategy sets the "strategy" field.
func (eiu *ExectionInfoUpdate) SetStrategy(s string) *ExectionInfoUpdate {
	eiu.mutation.SetStrategy(s)
	return eiu
}

// SetNillableStrategy sets the "strategy" field if the given value is not nil.
func (eiu *ExectionInfoUpdate) SetNillableStrategy(s *string) *ExectionInfoUpdate {
	if s != nil {
		eiu.SetStrategy(*s)
	}
	return eiu
}

// ClearStrategy clears the value of the "strategy" field.
func (eiu *ExectionInfoUpdate) ClearStrategy() *ExectionInfoUpdate {
	eiu.mutation.ClearStrategy()
	return eiu
}

// SetCachedRemotely sets the "cached_remotely" field.
func (eiu *ExectionInfoUpdate) SetCachedRemotely(b bool) *ExectionInfoUpdate {
	eiu.mutation.SetCachedRemotely(b)
	return eiu
}

// SetNillableCachedRemotely sets the "cached_remotely" field if the given value is not nil.
func (eiu *ExectionInfoUpdate) SetNillableCachedRemotely(b *bool) *ExectionInfoUpdate {
	if b != nil {
		eiu.SetCachedRemotely(*b)
	}
	return eiu
}

// ClearCachedRemotely clears the value of the "cached_remotely" field.
func (eiu *ExectionInfoUpdate) ClearCachedRemotely() *ExectionInfoUpdate {
	eiu.mutation.ClearCachedRemotely()
	return eiu
}

// SetExitCode sets the "exit_code" field.
func (eiu *ExectionInfoUpdate) SetExitCode(i int32) *ExectionInfoUpdate {
	eiu.mutation.ResetExitCode()
	eiu.mutation.SetExitCode(i)
	return eiu
}

// SetNillableExitCode sets the "exit_code" field if the given value is not nil.
func (eiu *ExectionInfoUpdate) SetNillableExitCode(i *int32) *ExectionInfoUpdate {
	if i != nil {
		eiu.SetExitCode(*i)
	}
	return eiu
}

// AddExitCode adds i to the "exit_code" field.
func (eiu *ExectionInfoUpdate) AddExitCode(i int32) *ExectionInfoUpdate {
	eiu.mutation.AddExitCode(i)
	return eiu
}

// ClearExitCode clears the value of the "exit_code" field.
func (eiu *ExectionInfoUpdate) ClearExitCode() *ExectionInfoUpdate {
	eiu.mutation.ClearExitCode()
	return eiu
}

// SetHostname sets the "hostname" field.
func (eiu *ExectionInfoUpdate) SetHostname(s string) *ExectionInfoUpdate {
	eiu.mutation.SetHostname(s)
	return eiu
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (eiu *ExectionInfoUpdate) SetNillableHostname(s *string) *ExectionInfoUpdate {
	if s != nil {
		eiu.SetHostname(*s)
	}
	return eiu
}

// ClearHostname clears the value of the "hostname" field.
func (eiu *ExectionInfoUpdate) ClearHostname() *ExectionInfoUpdate {
	eiu.mutation.ClearHostname()
	return eiu
}

// AddTestResultIDs adds the "test_result" edge to the TestResultBES entity by IDs.
func (eiu *ExectionInfoUpdate) AddTestResultIDs(ids ...int) *ExectionInfoUpdate {
	eiu.mutation.AddTestResultIDs(ids...)
	return eiu
}

// AddTestResult adds the "test_result" edges to the TestResultBES entity.
func (eiu *ExectionInfoUpdate) AddTestResult(t ...*TestResultBES) *ExectionInfoUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eiu.AddTestResultIDs(ids...)
}

// SetTimingBreakdownID sets the "timing_breakdown" edge to the TimingBreakdown entity by ID.
func (eiu *ExectionInfoUpdate) SetTimingBreakdownID(id int) *ExectionInfoUpdate {
	eiu.mutation.SetTimingBreakdownID(id)
	return eiu
}

// SetNillableTimingBreakdownID sets the "timing_breakdown" edge to the TimingBreakdown entity by ID if the given value is not nil.
func (eiu *ExectionInfoUpdate) SetNillableTimingBreakdownID(id *int) *ExectionInfoUpdate {
	if id != nil {
		eiu = eiu.SetTimingBreakdownID(*id)
	}
	return eiu
}

// SetTimingBreakdown sets the "timing_breakdown" edge to the TimingBreakdown entity.
func (eiu *ExectionInfoUpdate) SetTimingBreakdown(t *TimingBreakdown) *ExectionInfoUpdate {
	return eiu.SetTimingBreakdownID(t.ID)
}

// AddResourceUsageIDs adds the "resource_usage" edge to the ResourceUsage entity by IDs.
func (eiu *ExectionInfoUpdate) AddResourceUsageIDs(ids ...int) *ExectionInfoUpdate {
	eiu.mutation.AddResourceUsageIDs(ids...)
	return eiu
}

// AddResourceUsage adds the "resource_usage" edges to the ResourceUsage entity.
func (eiu *ExectionInfoUpdate) AddResourceUsage(r ...*ResourceUsage) *ExectionInfoUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return eiu.AddResourceUsageIDs(ids...)
}

// Mutation returns the ExectionInfoMutation object of the builder.
func (eiu *ExectionInfoUpdate) Mutation() *ExectionInfoMutation {
	return eiu.mutation
}

// ClearTestResult clears all "test_result" edges to the TestResultBES entity.
func (eiu *ExectionInfoUpdate) ClearTestResult() *ExectionInfoUpdate {
	eiu.mutation.ClearTestResult()
	return eiu
}

// RemoveTestResultIDs removes the "test_result" edge to TestResultBES entities by IDs.
func (eiu *ExectionInfoUpdate) RemoveTestResultIDs(ids ...int) *ExectionInfoUpdate {
	eiu.mutation.RemoveTestResultIDs(ids...)
	return eiu
}

// RemoveTestResult removes "test_result" edges to TestResultBES entities.
func (eiu *ExectionInfoUpdate) RemoveTestResult(t ...*TestResultBES) *ExectionInfoUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eiu.RemoveTestResultIDs(ids...)
}

// ClearTimingBreakdown clears the "timing_breakdown" edge to the TimingBreakdown entity.
func (eiu *ExectionInfoUpdate) ClearTimingBreakdown() *ExectionInfoUpdate {
	eiu.mutation.ClearTimingBreakdown()
	return eiu
}

// ClearResourceUsage clears all "resource_usage" edges to the ResourceUsage entity.
func (eiu *ExectionInfoUpdate) ClearResourceUsage() *ExectionInfoUpdate {
	eiu.mutation.ClearResourceUsage()
	return eiu
}

// RemoveResourceUsageIDs removes the "resource_usage" edge to ResourceUsage entities by IDs.
func (eiu *ExectionInfoUpdate) RemoveResourceUsageIDs(ids ...int) *ExectionInfoUpdate {
	eiu.mutation.RemoveResourceUsageIDs(ids...)
	return eiu
}

// RemoveResourceUsage removes "resource_usage" edges to ResourceUsage entities.
func (eiu *ExectionInfoUpdate) RemoveResourceUsage(r ...*ResourceUsage) *ExectionInfoUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return eiu.RemoveResourceUsageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eiu *ExectionInfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eiu.sqlSave, eiu.mutation, eiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eiu *ExectionInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := eiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eiu *ExectionInfoUpdate) Exec(ctx context.Context) error {
	_, err := eiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiu *ExectionInfoUpdate) ExecX(ctx context.Context) {
	if err := eiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eiu *ExectionInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(exectioninfo.Table, exectioninfo.Columns, sqlgraph.NewFieldSpec(exectioninfo.FieldID, field.TypeInt))
	if ps := eiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiu.mutation.Strategy(); ok {
		_spec.SetField(exectioninfo.FieldStrategy, field.TypeString, value)
	}
	if eiu.mutation.StrategyCleared() {
		_spec.ClearField(exectioninfo.FieldStrategy, field.TypeString)
	}
	if value, ok := eiu.mutation.CachedRemotely(); ok {
		_spec.SetField(exectioninfo.FieldCachedRemotely, field.TypeBool, value)
	}
	if eiu.mutation.CachedRemotelyCleared() {
		_spec.ClearField(exectioninfo.FieldCachedRemotely, field.TypeBool)
	}
	if value, ok := eiu.mutation.ExitCode(); ok {
		_spec.SetField(exectioninfo.FieldExitCode, field.TypeInt32, value)
	}
	if value, ok := eiu.mutation.AddedExitCode(); ok {
		_spec.AddField(exectioninfo.FieldExitCode, field.TypeInt32, value)
	}
	if eiu.mutation.ExitCodeCleared() {
		_spec.ClearField(exectioninfo.FieldExitCode, field.TypeInt32)
	}
	if value, ok := eiu.mutation.Hostname(); ok {
		_spec.SetField(exectioninfo.FieldHostname, field.TypeString, value)
	}
	if eiu.mutation.HostnameCleared() {
		_spec.ClearField(exectioninfo.FieldHostname, field.TypeString)
	}
	if eiu.mutation.TestResultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   exectioninfo.TestResultTable,
			Columns: []string{exectioninfo.TestResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiu.mutation.RemovedTestResultIDs(); len(nodes) > 0 && !eiu.mutation.TestResultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   exectioninfo.TestResultTable,
			Columns: []string{exectioninfo.TestResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiu.mutation.TestResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   exectioninfo.TestResultTable,
			Columns: []string{exectioninfo.TestResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eiu.mutation.TimingBreakdownCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   exectioninfo.TimingBreakdownTable,
			Columns: []string{exectioninfo.TimingBreakdownColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timingbreakdown.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiu.mutation.TimingBreakdownIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   exectioninfo.TimingBreakdownTable,
			Columns: []string{exectioninfo.TimingBreakdownColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timingbreakdown.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eiu.mutation.ResourceUsageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exectioninfo.ResourceUsageTable,
			Columns: exectioninfo.ResourceUsagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourceusage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiu.mutation.RemovedResourceUsageIDs(); len(nodes) > 0 && !eiu.mutation.ResourceUsageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exectioninfo.ResourceUsageTable,
			Columns: exectioninfo.ResourceUsagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourceusage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiu.mutation.ResourceUsageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exectioninfo.ResourceUsageTable,
			Columns: exectioninfo.ResourceUsagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourceusage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exectioninfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eiu.mutation.done = true
	return n, nil
}

// ExectionInfoUpdateOne is the builder for updating a single ExectionInfo entity.
type ExectionInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExectionInfoMutation
}

// SetStrategy sets the "strategy" field.
func (eiuo *ExectionInfoUpdateOne) SetStrategy(s string) *ExectionInfoUpdateOne {
	eiuo.mutation.SetStrategy(s)
	return eiuo
}

// SetNillableStrategy sets the "strategy" field if the given value is not nil.
func (eiuo *ExectionInfoUpdateOne) SetNillableStrategy(s *string) *ExectionInfoUpdateOne {
	if s != nil {
		eiuo.SetStrategy(*s)
	}
	return eiuo
}

// ClearStrategy clears the value of the "strategy" field.
func (eiuo *ExectionInfoUpdateOne) ClearStrategy() *ExectionInfoUpdateOne {
	eiuo.mutation.ClearStrategy()
	return eiuo
}

// SetCachedRemotely sets the "cached_remotely" field.
func (eiuo *ExectionInfoUpdateOne) SetCachedRemotely(b bool) *ExectionInfoUpdateOne {
	eiuo.mutation.SetCachedRemotely(b)
	return eiuo
}

// SetNillableCachedRemotely sets the "cached_remotely" field if the given value is not nil.
func (eiuo *ExectionInfoUpdateOne) SetNillableCachedRemotely(b *bool) *ExectionInfoUpdateOne {
	if b != nil {
		eiuo.SetCachedRemotely(*b)
	}
	return eiuo
}

// ClearCachedRemotely clears the value of the "cached_remotely" field.
func (eiuo *ExectionInfoUpdateOne) ClearCachedRemotely() *ExectionInfoUpdateOne {
	eiuo.mutation.ClearCachedRemotely()
	return eiuo
}

// SetExitCode sets the "exit_code" field.
func (eiuo *ExectionInfoUpdateOne) SetExitCode(i int32) *ExectionInfoUpdateOne {
	eiuo.mutation.ResetExitCode()
	eiuo.mutation.SetExitCode(i)
	return eiuo
}

// SetNillableExitCode sets the "exit_code" field if the given value is not nil.
func (eiuo *ExectionInfoUpdateOne) SetNillableExitCode(i *int32) *ExectionInfoUpdateOne {
	if i != nil {
		eiuo.SetExitCode(*i)
	}
	return eiuo
}

// AddExitCode adds i to the "exit_code" field.
func (eiuo *ExectionInfoUpdateOne) AddExitCode(i int32) *ExectionInfoUpdateOne {
	eiuo.mutation.AddExitCode(i)
	return eiuo
}

// ClearExitCode clears the value of the "exit_code" field.
func (eiuo *ExectionInfoUpdateOne) ClearExitCode() *ExectionInfoUpdateOne {
	eiuo.mutation.ClearExitCode()
	return eiuo
}

// SetHostname sets the "hostname" field.
func (eiuo *ExectionInfoUpdateOne) SetHostname(s string) *ExectionInfoUpdateOne {
	eiuo.mutation.SetHostname(s)
	return eiuo
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (eiuo *ExectionInfoUpdateOne) SetNillableHostname(s *string) *ExectionInfoUpdateOne {
	if s != nil {
		eiuo.SetHostname(*s)
	}
	return eiuo
}

// ClearHostname clears the value of the "hostname" field.
func (eiuo *ExectionInfoUpdateOne) ClearHostname() *ExectionInfoUpdateOne {
	eiuo.mutation.ClearHostname()
	return eiuo
}

// AddTestResultIDs adds the "test_result" edge to the TestResultBES entity by IDs.
func (eiuo *ExectionInfoUpdateOne) AddTestResultIDs(ids ...int) *ExectionInfoUpdateOne {
	eiuo.mutation.AddTestResultIDs(ids...)
	return eiuo
}

// AddTestResult adds the "test_result" edges to the TestResultBES entity.
func (eiuo *ExectionInfoUpdateOne) AddTestResult(t ...*TestResultBES) *ExectionInfoUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eiuo.AddTestResultIDs(ids...)
}

// SetTimingBreakdownID sets the "timing_breakdown" edge to the TimingBreakdown entity by ID.
func (eiuo *ExectionInfoUpdateOne) SetTimingBreakdownID(id int) *ExectionInfoUpdateOne {
	eiuo.mutation.SetTimingBreakdownID(id)
	return eiuo
}

// SetNillableTimingBreakdownID sets the "timing_breakdown" edge to the TimingBreakdown entity by ID if the given value is not nil.
func (eiuo *ExectionInfoUpdateOne) SetNillableTimingBreakdownID(id *int) *ExectionInfoUpdateOne {
	if id != nil {
		eiuo = eiuo.SetTimingBreakdownID(*id)
	}
	return eiuo
}

// SetTimingBreakdown sets the "timing_breakdown" edge to the TimingBreakdown entity.
func (eiuo *ExectionInfoUpdateOne) SetTimingBreakdown(t *TimingBreakdown) *ExectionInfoUpdateOne {
	return eiuo.SetTimingBreakdownID(t.ID)
}

// AddResourceUsageIDs adds the "resource_usage" edge to the ResourceUsage entity by IDs.
func (eiuo *ExectionInfoUpdateOne) AddResourceUsageIDs(ids ...int) *ExectionInfoUpdateOne {
	eiuo.mutation.AddResourceUsageIDs(ids...)
	return eiuo
}

// AddResourceUsage adds the "resource_usage" edges to the ResourceUsage entity.
func (eiuo *ExectionInfoUpdateOne) AddResourceUsage(r ...*ResourceUsage) *ExectionInfoUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return eiuo.AddResourceUsageIDs(ids...)
}

// Mutation returns the ExectionInfoMutation object of the builder.
func (eiuo *ExectionInfoUpdateOne) Mutation() *ExectionInfoMutation {
	return eiuo.mutation
}

// ClearTestResult clears all "test_result" edges to the TestResultBES entity.
func (eiuo *ExectionInfoUpdateOne) ClearTestResult() *ExectionInfoUpdateOne {
	eiuo.mutation.ClearTestResult()
	return eiuo
}

// RemoveTestResultIDs removes the "test_result" edge to TestResultBES entities by IDs.
func (eiuo *ExectionInfoUpdateOne) RemoveTestResultIDs(ids ...int) *ExectionInfoUpdateOne {
	eiuo.mutation.RemoveTestResultIDs(ids...)
	return eiuo
}

// RemoveTestResult removes "test_result" edges to TestResultBES entities.
func (eiuo *ExectionInfoUpdateOne) RemoveTestResult(t ...*TestResultBES) *ExectionInfoUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eiuo.RemoveTestResultIDs(ids...)
}

// ClearTimingBreakdown clears the "timing_breakdown" edge to the TimingBreakdown entity.
func (eiuo *ExectionInfoUpdateOne) ClearTimingBreakdown() *ExectionInfoUpdateOne {
	eiuo.mutation.ClearTimingBreakdown()
	return eiuo
}

// ClearResourceUsage clears all "resource_usage" edges to the ResourceUsage entity.
func (eiuo *ExectionInfoUpdateOne) ClearResourceUsage() *ExectionInfoUpdateOne {
	eiuo.mutation.ClearResourceUsage()
	return eiuo
}

// RemoveResourceUsageIDs removes the "resource_usage" edge to ResourceUsage entities by IDs.
func (eiuo *ExectionInfoUpdateOne) RemoveResourceUsageIDs(ids ...int) *ExectionInfoUpdateOne {
	eiuo.mutation.RemoveResourceUsageIDs(ids...)
	return eiuo
}

// RemoveResourceUsage removes "resource_usage" edges to ResourceUsage entities.
func (eiuo *ExectionInfoUpdateOne) RemoveResourceUsage(r ...*ResourceUsage) *ExectionInfoUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return eiuo.RemoveResourceUsageIDs(ids...)
}

// Where appends a list predicates to the ExectionInfoUpdate builder.
func (eiuo *ExectionInfoUpdateOne) Where(ps ...predicate.ExectionInfo) *ExectionInfoUpdateOne {
	eiuo.mutation.Where(ps...)
	return eiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eiuo *ExectionInfoUpdateOne) Select(field string, fields ...string) *ExectionInfoUpdateOne {
	eiuo.fields = append([]string{field}, fields...)
	return eiuo
}

// Save executes the query and returns the updated ExectionInfo entity.
func (eiuo *ExectionInfoUpdateOne) Save(ctx context.Context) (*ExectionInfo, error) {
	return withHooks(ctx, eiuo.sqlSave, eiuo.mutation, eiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eiuo *ExectionInfoUpdateOne) SaveX(ctx context.Context) *ExectionInfo {
	node, err := eiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eiuo *ExectionInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := eiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiuo *ExectionInfoUpdateOne) ExecX(ctx context.Context) {
	if err := eiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eiuo *ExectionInfoUpdateOne) sqlSave(ctx context.Context) (_node *ExectionInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(exectioninfo.Table, exectioninfo.Columns, sqlgraph.NewFieldSpec(exectioninfo.FieldID, field.TypeInt))
	id, ok := eiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExectionInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exectioninfo.FieldID)
		for _, f := range fields {
			if !exectioninfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != exectioninfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiuo.mutation.Strategy(); ok {
		_spec.SetField(exectioninfo.FieldStrategy, field.TypeString, value)
	}
	if eiuo.mutation.StrategyCleared() {
		_spec.ClearField(exectioninfo.FieldStrategy, field.TypeString)
	}
	if value, ok := eiuo.mutation.CachedRemotely(); ok {
		_spec.SetField(exectioninfo.FieldCachedRemotely, field.TypeBool, value)
	}
	if eiuo.mutation.CachedRemotelyCleared() {
		_spec.ClearField(exectioninfo.FieldCachedRemotely, field.TypeBool)
	}
	if value, ok := eiuo.mutation.ExitCode(); ok {
		_spec.SetField(exectioninfo.FieldExitCode, field.TypeInt32, value)
	}
	if value, ok := eiuo.mutation.AddedExitCode(); ok {
		_spec.AddField(exectioninfo.FieldExitCode, field.TypeInt32, value)
	}
	if eiuo.mutation.ExitCodeCleared() {
		_spec.ClearField(exectioninfo.FieldExitCode, field.TypeInt32)
	}
	if value, ok := eiuo.mutation.Hostname(); ok {
		_spec.SetField(exectioninfo.FieldHostname, field.TypeString, value)
	}
	if eiuo.mutation.HostnameCleared() {
		_spec.ClearField(exectioninfo.FieldHostname, field.TypeString)
	}
	if eiuo.mutation.TestResultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   exectioninfo.TestResultTable,
			Columns: []string{exectioninfo.TestResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiuo.mutation.RemovedTestResultIDs(); len(nodes) > 0 && !eiuo.mutation.TestResultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   exectioninfo.TestResultTable,
			Columns: []string{exectioninfo.TestResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiuo.mutation.TestResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   exectioninfo.TestResultTable,
			Columns: []string{exectioninfo.TestResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eiuo.mutation.TimingBreakdownCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   exectioninfo.TimingBreakdownTable,
			Columns: []string{exectioninfo.TimingBreakdownColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timingbreakdown.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiuo.mutation.TimingBreakdownIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   exectioninfo.TimingBreakdownTable,
			Columns: []string{exectioninfo.TimingBreakdownColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timingbreakdown.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eiuo.mutation.ResourceUsageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exectioninfo.ResourceUsageTable,
			Columns: exectioninfo.ResourceUsagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourceusage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiuo.mutation.RemovedResourceUsageIDs(); len(nodes) > 0 && !eiuo.mutation.ResourceUsageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exectioninfo.ResourceUsageTable,
			Columns: exectioninfo.ResourceUsagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourceusage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiuo.mutation.ResourceUsageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exectioninfo.ResourceUsageTable,
			Columns: exectioninfo.ResourceUsagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resourceusage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ExectionInfo{config: eiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exectioninfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eiuo.mutation.done = true
	return _node, nil
}