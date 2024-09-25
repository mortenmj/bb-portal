// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/namedsetoffiles"
	"github.com/buildbarn/bb-portal/ent/gen/ent/outputgroup"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testfile"
)

// NamedSetOfFilesQuery is the builder for querying NamedSetOfFiles entities.
type NamedSetOfFilesQuery struct {
	config
	ctx                  *QueryContext
	order                []namedsetoffiles.OrderOption
	inters               []Interceptor
	predicates           []predicate.NamedSetOfFiles
	withOutputGroup      *OutputGroupQuery
	withFiles            *TestFileQuery
	withFileSets         *NamedSetOfFilesQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*NamedSetOfFiles) error
	withNamedOutputGroup map[string]*OutputGroupQuery
	withNamedFiles       map[string]*TestFileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NamedSetOfFilesQuery builder.
func (nsofq *NamedSetOfFilesQuery) Where(ps ...predicate.NamedSetOfFiles) *NamedSetOfFilesQuery {
	nsofq.predicates = append(nsofq.predicates, ps...)
	return nsofq
}

// Limit the number of records to be returned by this query.
func (nsofq *NamedSetOfFilesQuery) Limit(limit int) *NamedSetOfFilesQuery {
	nsofq.ctx.Limit = &limit
	return nsofq
}

// Offset to start from.
func (nsofq *NamedSetOfFilesQuery) Offset(offset int) *NamedSetOfFilesQuery {
	nsofq.ctx.Offset = &offset
	return nsofq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nsofq *NamedSetOfFilesQuery) Unique(unique bool) *NamedSetOfFilesQuery {
	nsofq.ctx.Unique = &unique
	return nsofq
}

// Order specifies how the records should be ordered.
func (nsofq *NamedSetOfFilesQuery) Order(o ...namedsetoffiles.OrderOption) *NamedSetOfFilesQuery {
	nsofq.order = append(nsofq.order, o...)
	return nsofq
}

// QueryOutputGroup chains the current query on the "output_group" edge.
func (nsofq *NamedSetOfFilesQuery) QueryOutputGroup() *OutputGroupQuery {
	query := (&OutputGroupClient{config: nsofq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nsofq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nsofq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(namedsetoffiles.Table, namedsetoffiles.FieldID, selector),
			sqlgraph.To(outputgroup.Table, outputgroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, namedsetoffiles.OutputGroupTable, namedsetoffiles.OutputGroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(nsofq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFiles chains the current query on the "files" edge.
func (nsofq *NamedSetOfFilesQuery) QueryFiles() *TestFileQuery {
	query := (&TestFileClient{config: nsofq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nsofq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nsofq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(namedsetoffiles.Table, namedsetoffiles.FieldID, selector),
			sqlgraph.To(testfile.Table, testfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, namedsetoffiles.FilesTable, namedsetoffiles.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(nsofq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFileSets chains the current query on the "file_sets" edge.
func (nsofq *NamedSetOfFilesQuery) QueryFileSets() *NamedSetOfFilesQuery {
	query := (&NamedSetOfFilesClient{config: nsofq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nsofq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nsofq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(namedsetoffiles.Table, namedsetoffiles.FieldID, selector),
			sqlgraph.To(namedsetoffiles.Table, namedsetoffiles.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, namedsetoffiles.FileSetsTable, namedsetoffiles.FileSetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(nsofq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NamedSetOfFiles entity from the query.
// Returns a *NotFoundError when no NamedSetOfFiles was found.
func (nsofq *NamedSetOfFilesQuery) First(ctx context.Context) (*NamedSetOfFiles, error) {
	nodes, err := nsofq.Limit(1).All(setContextOp(ctx, nsofq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{namedsetoffiles.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nsofq *NamedSetOfFilesQuery) FirstX(ctx context.Context) *NamedSetOfFiles {
	node, err := nsofq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NamedSetOfFiles ID from the query.
// Returns a *NotFoundError when no NamedSetOfFiles ID was found.
func (nsofq *NamedSetOfFilesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nsofq.Limit(1).IDs(setContextOp(ctx, nsofq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{namedsetoffiles.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nsofq *NamedSetOfFilesQuery) FirstIDX(ctx context.Context) int {
	id, err := nsofq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NamedSetOfFiles entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NamedSetOfFiles entity is found.
// Returns a *NotFoundError when no NamedSetOfFiles entities are found.
func (nsofq *NamedSetOfFilesQuery) Only(ctx context.Context) (*NamedSetOfFiles, error) {
	nodes, err := nsofq.Limit(2).All(setContextOp(ctx, nsofq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{namedsetoffiles.Label}
	default:
		return nil, &NotSingularError{namedsetoffiles.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nsofq *NamedSetOfFilesQuery) OnlyX(ctx context.Context) *NamedSetOfFiles {
	node, err := nsofq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NamedSetOfFiles ID in the query.
// Returns a *NotSingularError when more than one NamedSetOfFiles ID is found.
// Returns a *NotFoundError when no entities are found.
func (nsofq *NamedSetOfFilesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nsofq.Limit(2).IDs(setContextOp(ctx, nsofq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{namedsetoffiles.Label}
	default:
		err = &NotSingularError{namedsetoffiles.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nsofq *NamedSetOfFilesQuery) OnlyIDX(ctx context.Context) int {
	id, err := nsofq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NamedSetOfFilesSlice.
func (nsofq *NamedSetOfFilesQuery) All(ctx context.Context) ([]*NamedSetOfFiles, error) {
	ctx = setContextOp(ctx, nsofq.ctx, "All")
	if err := nsofq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NamedSetOfFiles, *NamedSetOfFilesQuery]()
	return withInterceptors[[]*NamedSetOfFiles](ctx, nsofq, qr, nsofq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nsofq *NamedSetOfFilesQuery) AllX(ctx context.Context) []*NamedSetOfFiles {
	nodes, err := nsofq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NamedSetOfFiles IDs.
func (nsofq *NamedSetOfFilesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nsofq.ctx.Unique == nil && nsofq.path != nil {
		nsofq.Unique(true)
	}
	ctx = setContextOp(ctx, nsofq.ctx, "IDs")
	if err = nsofq.Select(namedsetoffiles.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nsofq *NamedSetOfFilesQuery) IDsX(ctx context.Context) []int {
	ids, err := nsofq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nsofq *NamedSetOfFilesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nsofq.ctx, "Count")
	if err := nsofq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nsofq, querierCount[*NamedSetOfFilesQuery](), nsofq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nsofq *NamedSetOfFilesQuery) CountX(ctx context.Context) int {
	count, err := nsofq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nsofq *NamedSetOfFilesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nsofq.ctx, "Exist")
	switch _, err := nsofq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nsofq *NamedSetOfFilesQuery) ExistX(ctx context.Context) bool {
	exist, err := nsofq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NamedSetOfFilesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nsofq *NamedSetOfFilesQuery) Clone() *NamedSetOfFilesQuery {
	if nsofq == nil {
		return nil
	}
	return &NamedSetOfFilesQuery{
		config:          nsofq.config,
		ctx:             nsofq.ctx.Clone(),
		order:           append([]namedsetoffiles.OrderOption{}, nsofq.order...),
		inters:          append([]Interceptor{}, nsofq.inters...),
		predicates:      append([]predicate.NamedSetOfFiles{}, nsofq.predicates...),
		withOutputGroup: nsofq.withOutputGroup.Clone(),
		withFiles:       nsofq.withFiles.Clone(),
		withFileSets:    nsofq.withFileSets.Clone(),
		// clone intermediate query.
		sql:  nsofq.sql.Clone(),
		path: nsofq.path,
	}
}

// WithOutputGroup tells the query-builder to eager-load the nodes that are connected to
// the "output_group" edge. The optional arguments are used to configure the query builder of the edge.
func (nsofq *NamedSetOfFilesQuery) WithOutputGroup(opts ...func(*OutputGroupQuery)) *NamedSetOfFilesQuery {
	query := (&OutputGroupClient{config: nsofq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nsofq.withOutputGroup = query
	return nsofq
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (nsofq *NamedSetOfFilesQuery) WithFiles(opts ...func(*TestFileQuery)) *NamedSetOfFilesQuery {
	query := (&TestFileClient{config: nsofq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nsofq.withFiles = query
	return nsofq
}

// WithFileSets tells the query-builder to eager-load the nodes that are connected to
// the "file_sets" edge. The optional arguments are used to configure the query builder of the edge.
func (nsofq *NamedSetOfFilesQuery) WithFileSets(opts ...func(*NamedSetOfFilesQuery)) *NamedSetOfFilesQuery {
	query := (&NamedSetOfFilesClient{config: nsofq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nsofq.withFileSets = query
	return nsofq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (nsofq *NamedSetOfFilesQuery) GroupBy(field string, fields ...string) *NamedSetOfFilesGroupBy {
	nsofq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NamedSetOfFilesGroupBy{build: nsofq}
	grbuild.flds = &nsofq.ctx.Fields
	grbuild.label = namedsetoffiles.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (nsofq *NamedSetOfFilesQuery) Select(fields ...string) *NamedSetOfFilesSelect {
	nsofq.ctx.Fields = append(nsofq.ctx.Fields, fields...)
	sbuild := &NamedSetOfFilesSelect{NamedSetOfFilesQuery: nsofq}
	sbuild.label = namedsetoffiles.Label
	sbuild.flds, sbuild.scan = &nsofq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NamedSetOfFilesSelect configured with the given aggregations.
func (nsofq *NamedSetOfFilesQuery) Aggregate(fns ...AggregateFunc) *NamedSetOfFilesSelect {
	return nsofq.Select().Aggregate(fns...)
}

func (nsofq *NamedSetOfFilesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nsofq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nsofq); err != nil {
				return err
			}
		}
	}
	for _, f := range nsofq.ctx.Fields {
		if !namedsetoffiles.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nsofq.path != nil {
		prev, err := nsofq.path(ctx)
		if err != nil {
			return err
		}
		nsofq.sql = prev
	}
	return nil
}

func (nsofq *NamedSetOfFilesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NamedSetOfFiles, error) {
	var (
		nodes       = []*NamedSetOfFiles{}
		withFKs     = nsofq.withFKs
		_spec       = nsofq.querySpec()
		loadedTypes = [3]bool{
			nsofq.withOutputGroup != nil,
			nsofq.withFiles != nil,
			nsofq.withFileSets != nil,
		}
	)
	if nsofq.withFileSets != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, namedsetoffiles.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NamedSetOfFiles).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NamedSetOfFiles{config: nsofq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(nsofq.modifiers) > 0 {
		_spec.Modifiers = nsofq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nsofq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nsofq.withOutputGroup; query != nil {
		if err := nsofq.loadOutputGroup(ctx, query, nodes,
			func(n *NamedSetOfFiles) { n.Edges.OutputGroup = []*OutputGroup{} },
			func(n *NamedSetOfFiles, e *OutputGroup) { n.Edges.OutputGroup = append(n.Edges.OutputGroup, e) }); err != nil {
			return nil, err
		}
	}
	if query := nsofq.withFiles; query != nil {
		if err := nsofq.loadFiles(ctx, query, nodes,
			func(n *NamedSetOfFiles) { n.Edges.Files = []*TestFile{} },
			func(n *NamedSetOfFiles, e *TestFile) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	if query := nsofq.withFileSets; query != nil {
		if err := nsofq.loadFileSets(ctx, query, nodes, nil,
			func(n *NamedSetOfFiles, e *NamedSetOfFiles) { n.Edges.FileSets = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range nsofq.withNamedOutputGroup {
		if err := nsofq.loadOutputGroup(ctx, query, nodes,
			func(n *NamedSetOfFiles) { n.appendNamedOutputGroup(name) },
			func(n *NamedSetOfFiles, e *OutputGroup) { n.appendNamedOutputGroup(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range nsofq.withNamedFiles {
		if err := nsofq.loadFiles(ctx, query, nodes,
			func(n *NamedSetOfFiles) { n.appendNamedFiles(name) },
			func(n *NamedSetOfFiles, e *TestFile) { n.appendNamedFiles(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range nsofq.loadTotal {
		if err := nsofq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nsofq *NamedSetOfFilesQuery) loadOutputGroup(ctx context.Context, query *OutputGroupQuery, nodes []*NamedSetOfFiles, init func(*NamedSetOfFiles), assign func(*NamedSetOfFiles, *OutputGroup)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*NamedSetOfFiles)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.OutputGroup(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(namedsetoffiles.OutputGroupColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.output_group_file_sets
		if fk == nil {
			return fmt.Errorf(`foreign-key "output_group_file_sets" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "output_group_file_sets" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nsofq *NamedSetOfFilesQuery) loadFiles(ctx context.Context, query *TestFileQuery, nodes []*NamedSetOfFiles, init func(*NamedSetOfFiles), assign func(*NamedSetOfFiles, *TestFile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*NamedSetOfFiles)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TestFile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(namedsetoffiles.FilesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.named_set_of_files_files
		if fk == nil {
			return fmt.Errorf(`foreign-key "named_set_of_files_files" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "named_set_of_files_files" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nsofq *NamedSetOfFilesQuery) loadFileSets(ctx context.Context, query *NamedSetOfFilesQuery, nodes []*NamedSetOfFiles, init func(*NamedSetOfFiles), assign func(*NamedSetOfFiles, *NamedSetOfFiles)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*NamedSetOfFiles)
	for i := range nodes {
		if nodes[i].named_set_of_files_file_sets == nil {
			continue
		}
		fk := *nodes[i].named_set_of_files_file_sets
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(namedsetoffiles.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "named_set_of_files_file_sets" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (nsofq *NamedSetOfFilesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nsofq.querySpec()
	if len(nsofq.modifiers) > 0 {
		_spec.Modifiers = nsofq.modifiers
	}
	_spec.Node.Columns = nsofq.ctx.Fields
	if len(nsofq.ctx.Fields) > 0 {
		_spec.Unique = nsofq.ctx.Unique != nil && *nsofq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nsofq.driver, _spec)
}

func (nsofq *NamedSetOfFilesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(namedsetoffiles.Table, namedsetoffiles.Columns, sqlgraph.NewFieldSpec(namedsetoffiles.FieldID, field.TypeInt))
	_spec.From = nsofq.sql
	if unique := nsofq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nsofq.path != nil {
		_spec.Unique = true
	}
	if fields := nsofq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, namedsetoffiles.FieldID)
		for i := range fields {
			if fields[i] != namedsetoffiles.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nsofq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nsofq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nsofq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nsofq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nsofq *NamedSetOfFilesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nsofq.driver.Dialect())
	t1 := builder.Table(namedsetoffiles.Table)
	columns := nsofq.ctx.Fields
	if len(columns) == 0 {
		columns = namedsetoffiles.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nsofq.sql != nil {
		selector = nsofq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nsofq.ctx.Unique != nil && *nsofq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nsofq.predicates {
		p(selector)
	}
	for _, p := range nsofq.order {
		p(selector)
	}
	if offset := nsofq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nsofq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOutputGroup tells the query-builder to eager-load the nodes that are connected to the "output_group"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (nsofq *NamedSetOfFilesQuery) WithNamedOutputGroup(name string, opts ...func(*OutputGroupQuery)) *NamedSetOfFilesQuery {
	query := (&OutputGroupClient{config: nsofq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if nsofq.withNamedOutputGroup == nil {
		nsofq.withNamedOutputGroup = make(map[string]*OutputGroupQuery)
	}
	nsofq.withNamedOutputGroup[name] = query
	return nsofq
}

// WithNamedFiles tells the query-builder to eager-load the nodes that are connected to the "files"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (nsofq *NamedSetOfFilesQuery) WithNamedFiles(name string, opts ...func(*TestFileQuery)) *NamedSetOfFilesQuery {
	query := (&TestFileClient{config: nsofq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if nsofq.withNamedFiles == nil {
		nsofq.withNamedFiles = make(map[string]*TestFileQuery)
	}
	nsofq.withNamedFiles[name] = query
	return nsofq
}

// NamedSetOfFilesGroupBy is the group-by builder for NamedSetOfFiles entities.
type NamedSetOfFilesGroupBy struct {
	selector
	build *NamedSetOfFilesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nsofgb *NamedSetOfFilesGroupBy) Aggregate(fns ...AggregateFunc) *NamedSetOfFilesGroupBy {
	nsofgb.fns = append(nsofgb.fns, fns...)
	return nsofgb
}

// Scan applies the selector query and scans the result into the given value.
func (nsofgb *NamedSetOfFilesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nsofgb.build.ctx, "GroupBy")
	if err := nsofgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NamedSetOfFilesQuery, *NamedSetOfFilesGroupBy](ctx, nsofgb.build, nsofgb, nsofgb.build.inters, v)
}

func (nsofgb *NamedSetOfFilesGroupBy) sqlScan(ctx context.Context, root *NamedSetOfFilesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nsofgb.fns))
	for _, fn := range nsofgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nsofgb.flds)+len(nsofgb.fns))
		for _, f := range *nsofgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nsofgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nsofgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NamedSetOfFilesSelect is the builder for selecting fields of NamedSetOfFiles entities.
type NamedSetOfFilesSelect struct {
	*NamedSetOfFilesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nsofs *NamedSetOfFilesSelect) Aggregate(fns ...AggregateFunc) *NamedSetOfFilesSelect {
	nsofs.fns = append(nsofs.fns, fns...)
	return nsofs
}

// Scan applies the selector query and scans the result into the given value.
func (nsofs *NamedSetOfFilesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nsofs.ctx, "Select")
	if err := nsofs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NamedSetOfFilesQuery, *NamedSetOfFilesSelect](ctx, nsofs.NamedSetOfFilesQuery, nsofs, nsofs.inters, v)
}

func (nsofs *NamedSetOfFilesSelect) sqlScan(ctx context.Context, root *NamedSetOfFilesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nsofs.fns))
	for _, fn := range nsofs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nsofs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nsofs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}