// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/namedsetoffiles"
	"github.com/buildbarn/bb-portal/ent/gen/ent/outputgroup"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testfile"
)

// NamedSetOfFilesUpdate is the builder for updating NamedSetOfFiles entities.
type NamedSetOfFilesUpdate struct {
	config
	hooks    []Hook
	mutation *NamedSetOfFilesMutation
}

// Where appends a list predicates to the NamedSetOfFilesUpdate builder.
func (nsofu *NamedSetOfFilesUpdate) Where(ps ...predicate.NamedSetOfFiles) *NamedSetOfFilesUpdate {
	nsofu.mutation.Where(ps...)
	return nsofu
}

// AddOutputGroupIDs adds the "output_group" edge to the OutputGroup entity by IDs.
func (nsofu *NamedSetOfFilesUpdate) AddOutputGroupIDs(ids ...int) *NamedSetOfFilesUpdate {
	nsofu.mutation.AddOutputGroupIDs(ids...)
	return nsofu
}

// AddOutputGroup adds the "output_group" edges to the OutputGroup entity.
func (nsofu *NamedSetOfFilesUpdate) AddOutputGroup(o ...*OutputGroup) *NamedSetOfFilesUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nsofu.AddOutputGroupIDs(ids...)
}

// AddFileIDs adds the "files" edge to the TestFile entity by IDs.
func (nsofu *NamedSetOfFilesUpdate) AddFileIDs(ids ...int) *NamedSetOfFilesUpdate {
	nsofu.mutation.AddFileIDs(ids...)
	return nsofu
}

// AddFiles adds the "files" edges to the TestFile entity.
func (nsofu *NamedSetOfFilesUpdate) AddFiles(t ...*TestFile) *NamedSetOfFilesUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nsofu.AddFileIDs(ids...)
}

// SetFileSetsID sets the "file_sets" edge to the NamedSetOfFiles entity by ID.
func (nsofu *NamedSetOfFilesUpdate) SetFileSetsID(id int) *NamedSetOfFilesUpdate {
	nsofu.mutation.SetFileSetsID(id)
	return nsofu
}

// SetNillableFileSetsID sets the "file_sets" edge to the NamedSetOfFiles entity by ID if the given value is not nil.
func (nsofu *NamedSetOfFilesUpdate) SetNillableFileSetsID(id *int) *NamedSetOfFilesUpdate {
	if id != nil {
		nsofu = nsofu.SetFileSetsID(*id)
	}
	return nsofu
}

// SetFileSets sets the "file_sets" edge to the NamedSetOfFiles entity.
func (nsofu *NamedSetOfFilesUpdate) SetFileSets(n *NamedSetOfFiles) *NamedSetOfFilesUpdate {
	return nsofu.SetFileSetsID(n.ID)
}

// Mutation returns the NamedSetOfFilesMutation object of the builder.
func (nsofu *NamedSetOfFilesUpdate) Mutation() *NamedSetOfFilesMutation {
	return nsofu.mutation
}

// ClearOutputGroup clears all "output_group" edges to the OutputGroup entity.
func (nsofu *NamedSetOfFilesUpdate) ClearOutputGroup() *NamedSetOfFilesUpdate {
	nsofu.mutation.ClearOutputGroup()
	return nsofu
}

// RemoveOutputGroupIDs removes the "output_group" edge to OutputGroup entities by IDs.
func (nsofu *NamedSetOfFilesUpdate) RemoveOutputGroupIDs(ids ...int) *NamedSetOfFilesUpdate {
	nsofu.mutation.RemoveOutputGroupIDs(ids...)
	return nsofu
}

// RemoveOutputGroup removes "output_group" edges to OutputGroup entities.
func (nsofu *NamedSetOfFilesUpdate) RemoveOutputGroup(o ...*OutputGroup) *NamedSetOfFilesUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nsofu.RemoveOutputGroupIDs(ids...)
}

// ClearFiles clears all "files" edges to the TestFile entity.
func (nsofu *NamedSetOfFilesUpdate) ClearFiles() *NamedSetOfFilesUpdate {
	nsofu.mutation.ClearFiles()
	return nsofu
}

// RemoveFileIDs removes the "files" edge to TestFile entities by IDs.
func (nsofu *NamedSetOfFilesUpdate) RemoveFileIDs(ids ...int) *NamedSetOfFilesUpdate {
	nsofu.mutation.RemoveFileIDs(ids...)
	return nsofu
}

// RemoveFiles removes "files" edges to TestFile entities.
func (nsofu *NamedSetOfFilesUpdate) RemoveFiles(t ...*TestFile) *NamedSetOfFilesUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nsofu.RemoveFileIDs(ids...)
}

// ClearFileSets clears the "file_sets" edge to the NamedSetOfFiles entity.
func (nsofu *NamedSetOfFilesUpdate) ClearFileSets() *NamedSetOfFilesUpdate {
	nsofu.mutation.ClearFileSets()
	return nsofu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nsofu *NamedSetOfFilesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nsofu.sqlSave, nsofu.mutation, nsofu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nsofu *NamedSetOfFilesUpdate) SaveX(ctx context.Context) int {
	affected, err := nsofu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nsofu *NamedSetOfFilesUpdate) Exec(ctx context.Context) error {
	_, err := nsofu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsofu *NamedSetOfFilesUpdate) ExecX(ctx context.Context) {
	if err := nsofu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nsofu *NamedSetOfFilesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(namedsetoffiles.Table, namedsetoffiles.Columns, sqlgraph.NewFieldSpec(namedsetoffiles.FieldID, field.TypeInt))
	if ps := nsofu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nsofu.mutation.OutputGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   namedsetoffiles.OutputGroupTable,
			Columns: []string{namedsetoffiles.OutputGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outputgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofu.mutation.RemovedOutputGroupIDs(); len(nodes) > 0 && !nsofu.mutation.OutputGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   namedsetoffiles.OutputGroupTable,
			Columns: []string{namedsetoffiles.OutputGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outputgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofu.mutation.OutputGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   namedsetoffiles.OutputGroupTable,
			Columns: []string{namedsetoffiles.OutputGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outputgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nsofu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namedsetoffiles.FilesTable,
			Columns: []string{namedsetoffiles.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !nsofu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namedsetoffiles.FilesTable,
			Columns: []string{namedsetoffiles.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namedsetoffiles.FilesTable,
			Columns: []string{namedsetoffiles.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nsofu.mutation.FileSetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   namedsetoffiles.FileSetsTable,
			Columns: []string{namedsetoffiles.FileSetsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namedsetoffiles.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofu.mutation.FileSetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   namedsetoffiles.FileSetsTable,
			Columns: []string{namedsetoffiles.FileSetsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namedsetoffiles.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nsofu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namedsetoffiles.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nsofu.mutation.done = true
	return n, nil
}

// NamedSetOfFilesUpdateOne is the builder for updating a single NamedSetOfFiles entity.
type NamedSetOfFilesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NamedSetOfFilesMutation
}

// AddOutputGroupIDs adds the "output_group" edge to the OutputGroup entity by IDs.
func (nsofuo *NamedSetOfFilesUpdateOne) AddOutputGroupIDs(ids ...int) *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.AddOutputGroupIDs(ids...)
	return nsofuo
}

// AddOutputGroup adds the "output_group" edges to the OutputGroup entity.
func (nsofuo *NamedSetOfFilesUpdateOne) AddOutputGroup(o ...*OutputGroup) *NamedSetOfFilesUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nsofuo.AddOutputGroupIDs(ids...)
}

// AddFileIDs adds the "files" edge to the TestFile entity by IDs.
func (nsofuo *NamedSetOfFilesUpdateOne) AddFileIDs(ids ...int) *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.AddFileIDs(ids...)
	return nsofuo
}

// AddFiles adds the "files" edges to the TestFile entity.
func (nsofuo *NamedSetOfFilesUpdateOne) AddFiles(t ...*TestFile) *NamedSetOfFilesUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nsofuo.AddFileIDs(ids...)
}

// SetFileSetsID sets the "file_sets" edge to the NamedSetOfFiles entity by ID.
func (nsofuo *NamedSetOfFilesUpdateOne) SetFileSetsID(id int) *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.SetFileSetsID(id)
	return nsofuo
}

// SetNillableFileSetsID sets the "file_sets" edge to the NamedSetOfFiles entity by ID if the given value is not nil.
func (nsofuo *NamedSetOfFilesUpdateOne) SetNillableFileSetsID(id *int) *NamedSetOfFilesUpdateOne {
	if id != nil {
		nsofuo = nsofuo.SetFileSetsID(*id)
	}
	return nsofuo
}

// SetFileSets sets the "file_sets" edge to the NamedSetOfFiles entity.
func (nsofuo *NamedSetOfFilesUpdateOne) SetFileSets(n *NamedSetOfFiles) *NamedSetOfFilesUpdateOne {
	return nsofuo.SetFileSetsID(n.ID)
}

// Mutation returns the NamedSetOfFilesMutation object of the builder.
func (nsofuo *NamedSetOfFilesUpdateOne) Mutation() *NamedSetOfFilesMutation {
	return nsofuo.mutation
}

// ClearOutputGroup clears all "output_group" edges to the OutputGroup entity.
func (nsofuo *NamedSetOfFilesUpdateOne) ClearOutputGroup() *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.ClearOutputGroup()
	return nsofuo
}

// RemoveOutputGroupIDs removes the "output_group" edge to OutputGroup entities by IDs.
func (nsofuo *NamedSetOfFilesUpdateOne) RemoveOutputGroupIDs(ids ...int) *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.RemoveOutputGroupIDs(ids...)
	return nsofuo
}

// RemoveOutputGroup removes "output_group" edges to OutputGroup entities.
func (nsofuo *NamedSetOfFilesUpdateOne) RemoveOutputGroup(o ...*OutputGroup) *NamedSetOfFilesUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return nsofuo.RemoveOutputGroupIDs(ids...)
}

// ClearFiles clears all "files" edges to the TestFile entity.
func (nsofuo *NamedSetOfFilesUpdateOne) ClearFiles() *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.ClearFiles()
	return nsofuo
}

// RemoveFileIDs removes the "files" edge to TestFile entities by IDs.
func (nsofuo *NamedSetOfFilesUpdateOne) RemoveFileIDs(ids ...int) *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.RemoveFileIDs(ids...)
	return nsofuo
}

// RemoveFiles removes "files" edges to TestFile entities.
func (nsofuo *NamedSetOfFilesUpdateOne) RemoveFiles(t ...*TestFile) *NamedSetOfFilesUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nsofuo.RemoveFileIDs(ids...)
}

// ClearFileSets clears the "file_sets" edge to the NamedSetOfFiles entity.
func (nsofuo *NamedSetOfFilesUpdateOne) ClearFileSets() *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.ClearFileSets()
	return nsofuo
}

// Where appends a list predicates to the NamedSetOfFilesUpdate builder.
func (nsofuo *NamedSetOfFilesUpdateOne) Where(ps ...predicate.NamedSetOfFiles) *NamedSetOfFilesUpdateOne {
	nsofuo.mutation.Where(ps...)
	return nsofuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nsofuo *NamedSetOfFilesUpdateOne) Select(field string, fields ...string) *NamedSetOfFilesUpdateOne {
	nsofuo.fields = append([]string{field}, fields...)
	return nsofuo
}

// Save executes the query and returns the updated NamedSetOfFiles entity.
func (nsofuo *NamedSetOfFilesUpdateOne) Save(ctx context.Context) (*NamedSetOfFiles, error) {
	return withHooks(ctx, nsofuo.sqlSave, nsofuo.mutation, nsofuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nsofuo *NamedSetOfFilesUpdateOne) SaveX(ctx context.Context) *NamedSetOfFiles {
	node, err := nsofuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nsofuo *NamedSetOfFilesUpdateOne) Exec(ctx context.Context) error {
	_, err := nsofuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsofuo *NamedSetOfFilesUpdateOne) ExecX(ctx context.Context) {
	if err := nsofuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nsofuo *NamedSetOfFilesUpdateOne) sqlSave(ctx context.Context) (_node *NamedSetOfFiles, err error) {
	_spec := sqlgraph.NewUpdateSpec(namedsetoffiles.Table, namedsetoffiles.Columns, sqlgraph.NewFieldSpec(namedsetoffiles.FieldID, field.TypeInt))
	id, ok := nsofuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NamedSetOfFiles.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nsofuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, namedsetoffiles.FieldID)
		for _, f := range fields {
			if !namedsetoffiles.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != namedsetoffiles.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nsofuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nsofuo.mutation.OutputGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   namedsetoffiles.OutputGroupTable,
			Columns: []string{namedsetoffiles.OutputGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outputgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofuo.mutation.RemovedOutputGroupIDs(); len(nodes) > 0 && !nsofuo.mutation.OutputGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   namedsetoffiles.OutputGroupTable,
			Columns: []string{namedsetoffiles.OutputGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outputgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofuo.mutation.OutputGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   namedsetoffiles.OutputGroupTable,
			Columns: []string{namedsetoffiles.OutputGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outputgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nsofuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namedsetoffiles.FilesTable,
			Columns: []string{namedsetoffiles.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !nsofuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namedsetoffiles.FilesTable,
			Columns: []string{namedsetoffiles.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namedsetoffiles.FilesTable,
			Columns: []string{namedsetoffiles.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nsofuo.mutation.FileSetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   namedsetoffiles.FileSetsTable,
			Columns: []string{namedsetoffiles.FileSetsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namedsetoffiles.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsofuo.mutation.FileSetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   namedsetoffiles.FileSetsTable,
			Columns: []string{namedsetoffiles.FileSetsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namedsetoffiles.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NamedSetOfFiles{config: nsofuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nsofuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namedsetoffiles.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nsofuo.mutation.done = true
	return _node, nil
}