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
	"github.com/buildbarn/bb-portal/ent/gen/ent/bazelinvocation"
	"github.com/buildbarn/bb-portal/ent/gen/ent/build"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// BuildQuery is the builder for querying Build entities.
type BuildQuery struct {
	config
	ctx                  *QueryContext
	order                []build.OrderOption
	inters               []Interceptor
	predicates           []predicate.Build
	withInvocations      *BazelInvocationQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*Build) error
	withNamedInvocations map[string]*BazelInvocationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BuildQuery builder.
func (bq *BuildQuery) Where(ps ...predicate.Build) *BuildQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BuildQuery) Limit(limit int) *BuildQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BuildQuery) Offset(offset int) *BuildQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BuildQuery) Unique(unique bool) *BuildQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BuildQuery) Order(o ...build.OrderOption) *BuildQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryInvocations chains the current query on the "invocations" edge.
func (bq *BuildQuery) QueryInvocations() *BazelInvocationQuery {
	query := (&BazelInvocationClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(build.Table, build.FieldID, selector),
			sqlgraph.To(bazelinvocation.Table, bazelinvocation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, build.InvocationsTable, build.InvocationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Build entity from the query.
// Returns a *NotFoundError when no Build was found.
func (bq *BuildQuery) First(ctx context.Context) (*Build, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{build.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BuildQuery) FirstX(ctx context.Context) *Build {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Build ID from the query.
// Returns a *NotFoundError when no Build ID was found.
func (bq *BuildQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{build.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BuildQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Build entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Build entity is found.
// Returns a *NotFoundError when no Build entities are found.
func (bq *BuildQuery) Only(ctx context.Context) (*Build, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{build.Label}
	default:
		return nil, &NotSingularError{build.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BuildQuery) OnlyX(ctx context.Context) *Build {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Build ID in the query.
// Returns a *NotSingularError when more than one Build ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BuildQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{build.Label}
	default:
		err = &NotSingularError{build.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BuildQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Builds.
func (bq *BuildQuery) All(ctx context.Context) ([]*Build, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Build, *BuildQuery]()
	return withInterceptors[[]*Build](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BuildQuery) AllX(ctx context.Context) []*Build {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Build IDs.
func (bq *BuildQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(build.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BuildQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BuildQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BuildQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BuildQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BuildQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BuildQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BuildQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BuildQuery) Clone() *BuildQuery {
	if bq == nil {
		return nil
	}
	return &BuildQuery{
		config:          bq.config,
		ctx:             bq.ctx.Clone(),
		order:           append([]build.OrderOption{}, bq.order...),
		inters:          append([]Interceptor{}, bq.inters...),
		predicates:      append([]predicate.Build{}, bq.predicates...),
		withInvocations: bq.withInvocations.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithInvocations tells the query-builder to eager-load the nodes that are connected to
// the "invocations" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BuildQuery) WithInvocations(opts ...func(*BazelInvocationQuery)) *BuildQuery {
	query := (&BazelInvocationClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withInvocations = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BuildURL string `json:"build_url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Build.Query().
//		GroupBy(build.FieldBuildURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BuildQuery) GroupBy(field string, fields ...string) *BuildGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BuildGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = build.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BuildURL string `json:"build_url,omitempty"`
//	}
//
//	client.Build.Query().
//		Select(build.FieldBuildURL).
//		Scan(ctx, &v)
func (bq *BuildQuery) Select(fields ...string) *BuildSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BuildSelect{BuildQuery: bq}
	sbuild.label = build.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BuildSelect configured with the given aggregations.
func (bq *BuildQuery) Aggregate(fns ...AggregateFunc) *BuildSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BuildQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !build.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BuildQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Build, error) {
	var (
		nodes       = []*Build{}
		_spec       = bq.querySpec()
		loadedTypes = [1]bool{
			bq.withInvocations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Build).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Build{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withInvocations; query != nil {
		if err := bq.loadInvocations(ctx, query, nodes,
			func(n *Build) { n.Edges.Invocations = []*BazelInvocation{} },
			func(n *Build, e *BazelInvocation) { n.Edges.Invocations = append(n.Edges.Invocations, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedInvocations {
		if err := bq.loadInvocations(ctx, query, nodes,
			func(n *Build) { n.appendNamedInvocations(name) },
			func(n *Build, e *BazelInvocation) { n.appendNamedInvocations(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range bq.loadTotal {
		if err := bq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BuildQuery) loadInvocations(ctx context.Context, query *BazelInvocationQuery, nodes []*Build, init func(*Build), assign func(*Build, *BazelInvocation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Build)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.BazelInvocation(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(build.InvocationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.build_invocations
		if fk == nil {
			return fmt.Errorf(`foreign-key "build_invocations" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "build_invocations" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BuildQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BuildQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(build.Table, build.Columns, sqlgraph.NewFieldSpec(build.FieldID, field.TypeInt))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, build.FieldID)
		for i := range fields {
			if fields[i] != build.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BuildQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(build.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = build.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedInvocations tells the query-builder to eager-load the nodes that are connected to the "invocations"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BuildQuery) WithNamedInvocations(name string, opts ...func(*BazelInvocationQuery)) *BuildQuery {
	query := (&BazelInvocationClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedInvocations == nil {
		bq.withNamedInvocations = make(map[string]*BazelInvocationQuery)
	}
	bq.withNamedInvocations[name] = query
	return bq
}

// BuildGroupBy is the group-by builder for Build entities.
type BuildGroupBy struct {
	selector
	build *BuildQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BuildGroupBy) Aggregate(fns ...AggregateFunc) *BuildGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BuildGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BuildQuery, *BuildGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BuildGroupBy) sqlScan(ctx context.Context, root *BuildQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BuildSelect is the builder for selecting fields of Build entities.
type BuildSelect struct {
	*BuildQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BuildSelect) Aggregate(fns ...AggregateFunc) *BuildSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BuildSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BuildQuery, *BuildSelect](ctx, bs.BuildQuery, bs, bs.inters, v)
}

func (bs *BuildSelect) sqlScan(ctx context.Context, root *BuildQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
