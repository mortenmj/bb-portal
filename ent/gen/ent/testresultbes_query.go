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
	"github.com/buildbarn/bb-portal/ent/gen/ent/exectioninfo"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testcollection"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testfile"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testresultbes"
)

// TestResultBESQuery is the builder for querying TestResultBES entities.
type TestResultBESQuery struct {
	config
	ctx                       *QueryContext
	order                     []testresultbes.OrderOption
	inters                    []Interceptor
	predicates                []predicate.TestResultBES
	withTestCollection        *TestCollectionQuery
	withTestActionOutput      *TestFileQuery
	withExecutionInfo         *ExectionInfoQuery
	withFKs                   bool
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*TestResultBES) error
	withNamedTestActionOutput map[string]*TestFileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestResultBESQuery builder.
func (trbq *TestResultBESQuery) Where(ps ...predicate.TestResultBES) *TestResultBESQuery {
	trbq.predicates = append(trbq.predicates, ps...)
	return trbq
}

// Limit the number of records to be returned by this query.
func (trbq *TestResultBESQuery) Limit(limit int) *TestResultBESQuery {
	trbq.ctx.Limit = &limit
	return trbq
}

// Offset to start from.
func (trbq *TestResultBESQuery) Offset(offset int) *TestResultBESQuery {
	trbq.ctx.Offset = &offset
	return trbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (trbq *TestResultBESQuery) Unique(unique bool) *TestResultBESQuery {
	trbq.ctx.Unique = &unique
	return trbq
}

// Order specifies how the records should be ordered.
func (trbq *TestResultBESQuery) Order(o ...testresultbes.OrderOption) *TestResultBESQuery {
	trbq.order = append(trbq.order, o...)
	return trbq
}

// QueryTestCollection chains the current query on the "test_collection" edge.
func (trbq *TestResultBESQuery) QueryTestCollection() *TestCollectionQuery {
	query := (&TestCollectionClient{config: trbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := trbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := trbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testresultbes.Table, testresultbes.FieldID, selector),
			sqlgraph.To(testcollection.Table, testcollection.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, testresultbes.TestCollectionTable, testresultbes.TestCollectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(trbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTestActionOutput chains the current query on the "test_action_output" edge.
func (trbq *TestResultBESQuery) QueryTestActionOutput() *TestFileQuery {
	query := (&TestFileClient{config: trbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := trbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := trbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testresultbes.Table, testresultbes.FieldID, selector),
			sqlgraph.To(testfile.Table, testfile.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, testresultbes.TestActionOutputTable, testresultbes.TestActionOutputPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(trbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExecutionInfo chains the current query on the "execution_info" edge.
func (trbq *TestResultBESQuery) QueryExecutionInfo() *ExectionInfoQuery {
	query := (&ExectionInfoClient{config: trbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := trbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := trbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testresultbes.Table, testresultbes.FieldID, selector),
			sqlgraph.To(exectioninfo.Table, exectioninfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, testresultbes.ExecutionInfoTable, testresultbes.ExecutionInfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(trbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TestResultBES entity from the query.
// Returns a *NotFoundError when no TestResultBES was found.
func (trbq *TestResultBESQuery) First(ctx context.Context) (*TestResultBES, error) {
	nodes, err := trbq.Limit(1).All(setContextOp(ctx, trbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testresultbes.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (trbq *TestResultBESQuery) FirstX(ctx context.Context) *TestResultBES {
	node, err := trbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestResultBES ID from the query.
// Returns a *NotFoundError when no TestResultBES ID was found.
func (trbq *TestResultBESQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trbq.Limit(1).IDs(setContextOp(ctx, trbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testresultbes.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (trbq *TestResultBESQuery) FirstIDX(ctx context.Context) int {
	id, err := trbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestResultBES entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestResultBES entity is found.
// Returns a *NotFoundError when no TestResultBES entities are found.
func (trbq *TestResultBESQuery) Only(ctx context.Context) (*TestResultBES, error) {
	nodes, err := trbq.Limit(2).All(setContextOp(ctx, trbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testresultbes.Label}
	default:
		return nil, &NotSingularError{testresultbes.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (trbq *TestResultBESQuery) OnlyX(ctx context.Context) *TestResultBES {
	node, err := trbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestResultBES ID in the query.
// Returns a *NotSingularError when more than one TestResultBES ID is found.
// Returns a *NotFoundError when no entities are found.
func (trbq *TestResultBESQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trbq.Limit(2).IDs(setContextOp(ctx, trbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testresultbes.Label}
	default:
		err = &NotSingularError{testresultbes.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (trbq *TestResultBESQuery) OnlyIDX(ctx context.Context) int {
	id, err := trbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestResultBESs.
func (trbq *TestResultBESQuery) All(ctx context.Context) ([]*TestResultBES, error) {
	ctx = setContextOp(ctx, trbq.ctx, "All")
	if err := trbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TestResultBES, *TestResultBESQuery]()
	return withInterceptors[[]*TestResultBES](ctx, trbq, qr, trbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (trbq *TestResultBESQuery) AllX(ctx context.Context) []*TestResultBES {
	nodes, err := trbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestResultBES IDs.
func (trbq *TestResultBESQuery) IDs(ctx context.Context) (ids []int, err error) {
	if trbq.ctx.Unique == nil && trbq.path != nil {
		trbq.Unique(true)
	}
	ctx = setContextOp(ctx, trbq.ctx, "IDs")
	if err = trbq.Select(testresultbes.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (trbq *TestResultBESQuery) IDsX(ctx context.Context) []int {
	ids, err := trbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (trbq *TestResultBESQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, trbq.ctx, "Count")
	if err := trbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, trbq, querierCount[*TestResultBESQuery](), trbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (trbq *TestResultBESQuery) CountX(ctx context.Context) int {
	count, err := trbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (trbq *TestResultBESQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, trbq.ctx, "Exist")
	switch _, err := trbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (trbq *TestResultBESQuery) ExistX(ctx context.Context) bool {
	exist, err := trbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestResultBESQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (trbq *TestResultBESQuery) Clone() *TestResultBESQuery {
	if trbq == nil {
		return nil
	}
	return &TestResultBESQuery{
		config:               trbq.config,
		ctx:                  trbq.ctx.Clone(),
		order:                append([]testresultbes.OrderOption{}, trbq.order...),
		inters:               append([]Interceptor{}, trbq.inters...),
		predicates:           append([]predicate.TestResultBES{}, trbq.predicates...),
		withTestCollection:   trbq.withTestCollection.Clone(),
		withTestActionOutput: trbq.withTestActionOutput.Clone(),
		withExecutionInfo:    trbq.withExecutionInfo.Clone(),
		// clone intermediate query.
		sql:  trbq.sql.Clone(),
		path: trbq.path,
	}
}

// WithTestCollection tells the query-builder to eager-load the nodes that are connected to
// the "test_collection" edge. The optional arguments are used to configure the query builder of the edge.
func (trbq *TestResultBESQuery) WithTestCollection(opts ...func(*TestCollectionQuery)) *TestResultBESQuery {
	query := (&TestCollectionClient{config: trbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	trbq.withTestCollection = query
	return trbq
}

// WithTestActionOutput tells the query-builder to eager-load the nodes that are connected to
// the "test_action_output" edge. The optional arguments are used to configure the query builder of the edge.
func (trbq *TestResultBESQuery) WithTestActionOutput(opts ...func(*TestFileQuery)) *TestResultBESQuery {
	query := (&TestFileClient{config: trbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	trbq.withTestActionOutput = query
	return trbq
}

// WithExecutionInfo tells the query-builder to eager-load the nodes that are connected to
// the "execution_info" edge. The optional arguments are used to configure the query builder of the edge.
func (trbq *TestResultBESQuery) WithExecutionInfo(opts ...func(*ExectionInfoQuery)) *TestResultBESQuery {
	query := (&ExectionInfoClient{config: trbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	trbq.withExecutionInfo = query
	return trbq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TestStatus testresultbes.TestStatus `json:"test_status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TestResultBES.Query().
//		GroupBy(testresultbes.FieldTestStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (trbq *TestResultBESQuery) GroupBy(field string, fields ...string) *TestResultBESGroupBy {
	trbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TestResultBESGroupBy{build: trbq}
	grbuild.flds = &trbq.ctx.Fields
	grbuild.label = testresultbes.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TestStatus testresultbes.TestStatus `json:"test_status,omitempty"`
//	}
//
//	client.TestResultBES.Query().
//		Select(testresultbes.FieldTestStatus).
//		Scan(ctx, &v)
func (trbq *TestResultBESQuery) Select(fields ...string) *TestResultBESSelect {
	trbq.ctx.Fields = append(trbq.ctx.Fields, fields...)
	sbuild := &TestResultBESSelect{TestResultBESQuery: trbq}
	sbuild.label = testresultbes.Label
	sbuild.flds, sbuild.scan = &trbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TestResultBESSelect configured with the given aggregations.
func (trbq *TestResultBESQuery) Aggregate(fns ...AggregateFunc) *TestResultBESSelect {
	return trbq.Select().Aggregate(fns...)
}

func (trbq *TestResultBESQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range trbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, trbq); err != nil {
				return err
			}
		}
	}
	for _, f := range trbq.ctx.Fields {
		if !testresultbes.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if trbq.path != nil {
		prev, err := trbq.path(ctx)
		if err != nil {
			return err
		}
		trbq.sql = prev
	}
	return nil
}

func (trbq *TestResultBESQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestResultBES, error) {
	var (
		nodes       = []*TestResultBES{}
		withFKs     = trbq.withFKs
		_spec       = trbq.querySpec()
		loadedTypes = [3]bool{
			trbq.withTestCollection != nil,
			trbq.withTestActionOutput != nil,
			trbq.withExecutionInfo != nil,
		}
	)
	if trbq.withTestCollection != nil || trbq.withExecutionInfo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, testresultbes.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TestResultBES).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TestResultBES{config: trbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(trbq.modifiers) > 0 {
		_spec.Modifiers = trbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, trbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := trbq.withTestCollection; query != nil {
		if err := trbq.loadTestCollection(ctx, query, nodes, nil,
			func(n *TestResultBES, e *TestCollection) { n.Edges.TestCollection = e }); err != nil {
			return nil, err
		}
	}
	if query := trbq.withTestActionOutput; query != nil {
		if err := trbq.loadTestActionOutput(ctx, query, nodes,
			func(n *TestResultBES) { n.Edges.TestActionOutput = []*TestFile{} },
			func(n *TestResultBES, e *TestFile) { n.Edges.TestActionOutput = append(n.Edges.TestActionOutput, e) }); err != nil {
			return nil, err
		}
	}
	if query := trbq.withExecutionInfo; query != nil {
		if err := trbq.loadExecutionInfo(ctx, query, nodes, nil,
			func(n *TestResultBES, e *ExectionInfo) { n.Edges.ExecutionInfo = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range trbq.withNamedTestActionOutput {
		if err := trbq.loadTestActionOutput(ctx, query, nodes,
			func(n *TestResultBES) { n.appendNamedTestActionOutput(name) },
			func(n *TestResultBES, e *TestFile) { n.appendNamedTestActionOutput(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range trbq.loadTotal {
		if err := trbq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (trbq *TestResultBESQuery) loadTestCollection(ctx context.Context, query *TestCollectionQuery, nodes []*TestResultBES, init func(*TestResultBES), assign func(*TestResultBES, *TestCollection)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TestResultBES)
	for i := range nodes {
		if nodes[i].test_collection_test_results == nil {
			continue
		}
		fk := *nodes[i].test_collection_test_results
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(testcollection.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "test_collection_test_results" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (trbq *TestResultBESQuery) loadTestActionOutput(ctx context.Context, query *TestFileQuery, nodes []*TestResultBES, init func(*TestResultBES), assign func(*TestResultBES, *TestFile)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*TestResultBES)
	nids := make(map[int]map[*TestResultBES]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(testresultbes.TestActionOutputTable)
		s.Join(joinT).On(s.C(testfile.FieldID), joinT.C(testresultbes.TestActionOutputPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(testresultbes.TestActionOutputPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(testresultbes.TestActionOutputPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*TestResultBES]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*TestFile](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "test_action_output" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (trbq *TestResultBESQuery) loadExecutionInfo(ctx context.Context, query *ExectionInfoQuery, nodes []*TestResultBES, init func(*TestResultBES), assign func(*TestResultBES, *ExectionInfo)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TestResultBES)
	for i := range nodes {
		if nodes[i].test_result_bes_execution_info == nil {
			continue
		}
		fk := *nodes[i].test_result_bes_execution_info
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(exectioninfo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "test_result_bes_execution_info" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (trbq *TestResultBESQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := trbq.querySpec()
	if len(trbq.modifiers) > 0 {
		_spec.Modifiers = trbq.modifiers
	}
	_spec.Node.Columns = trbq.ctx.Fields
	if len(trbq.ctx.Fields) > 0 {
		_spec.Unique = trbq.ctx.Unique != nil && *trbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, trbq.driver, _spec)
}

func (trbq *TestResultBESQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(testresultbes.Table, testresultbes.Columns, sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt))
	_spec.From = trbq.sql
	if unique := trbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if trbq.path != nil {
		_spec.Unique = true
	}
	if fields := trbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testresultbes.FieldID)
		for i := range fields {
			if fields[i] != testresultbes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := trbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := trbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := trbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := trbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (trbq *TestResultBESQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(trbq.driver.Dialect())
	t1 := builder.Table(testresultbes.Table)
	columns := trbq.ctx.Fields
	if len(columns) == 0 {
		columns = testresultbes.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if trbq.sql != nil {
		selector = trbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if trbq.ctx.Unique != nil && *trbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range trbq.predicates {
		p(selector)
	}
	for _, p := range trbq.order {
		p(selector)
	}
	if offset := trbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := trbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTestActionOutput tells the query-builder to eager-load the nodes that are connected to the "test_action_output"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (trbq *TestResultBESQuery) WithNamedTestActionOutput(name string, opts ...func(*TestFileQuery)) *TestResultBESQuery {
	query := (&TestFileClient{config: trbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if trbq.withNamedTestActionOutput == nil {
		trbq.withNamedTestActionOutput = make(map[string]*TestFileQuery)
	}
	trbq.withNamedTestActionOutput[name] = query
	return trbq
}

// TestResultBESGroupBy is the group-by builder for TestResultBES entities.
type TestResultBESGroupBy struct {
	selector
	build *TestResultBESQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (trbgb *TestResultBESGroupBy) Aggregate(fns ...AggregateFunc) *TestResultBESGroupBy {
	trbgb.fns = append(trbgb.fns, fns...)
	return trbgb
}

// Scan applies the selector query and scans the result into the given value.
func (trbgb *TestResultBESGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, trbgb.build.ctx, "GroupBy")
	if err := trbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestResultBESQuery, *TestResultBESGroupBy](ctx, trbgb.build, trbgb, trbgb.build.inters, v)
}

func (trbgb *TestResultBESGroupBy) sqlScan(ctx context.Context, root *TestResultBESQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(trbgb.fns))
	for _, fn := range trbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*trbgb.flds)+len(trbgb.fns))
		for _, f := range *trbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*trbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TestResultBESSelect is the builder for selecting fields of TestResultBES entities.
type TestResultBESSelect struct {
	*TestResultBESQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (trbs *TestResultBESSelect) Aggregate(fns ...AggregateFunc) *TestResultBESSelect {
	trbs.fns = append(trbs.fns, fns...)
	return trbs
}

// Scan applies the selector query and scans the result into the given value.
func (trbs *TestResultBESSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, trbs.ctx, "Select")
	if err := trbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestResultBESQuery, *TestResultBESSelect](ctx, trbs.TestResultBESQuery, trbs, trbs.inters, v)
}

func (trbs *TestResultBESSelect) sqlScan(ctx context.Context, root *TestResultBESQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(trbs.fns))
	for _, fn := range trbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*trbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
