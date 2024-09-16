// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/networkmetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/systemnetworkstats"
)

// SystemNetworkStatsQuery is the builder for querying SystemNetworkStats entities.
type SystemNetworkStatsQuery struct {
	config
	ctx                *QueryContext
	order              []systemnetworkstats.OrderOption
	inters             []Interceptor
	predicates         []predicate.SystemNetworkStats
	withNetworkMetrics *NetworkMetricsQuery
	withFKs            bool
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*SystemNetworkStats) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SystemNetworkStatsQuery builder.
func (snsq *SystemNetworkStatsQuery) Where(ps ...predicate.SystemNetworkStats) *SystemNetworkStatsQuery {
	snsq.predicates = append(snsq.predicates, ps...)
	return snsq
}

// Limit the number of records to be returned by this query.
func (snsq *SystemNetworkStatsQuery) Limit(limit int) *SystemNetworkStatsQuery {
	snsq.ctx.Limit = &limit
	return snsq
}

// Offset to start from.
func (snsq *SystemNetworkStatsQuery) Offset(offset int) *SystemNetworkStatsQuery {
	snsq.ctx.Offset = &offset
	return snsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (snsq *SystemNetworkStatsQuery) Unique(unique bool) *SystemNetworkStatsQuery {
	snsq.ctx.Unique = &unique
	return snsq
}

// Order specifies how the records should be ordered.
func (snsq *SystemNetworkStatsQuery) Order(o ...systemnetworkstats.OrderOption) *SystemNetworkStatsQuery {
	snsq.order = append(snsq.order, o...)
	return snsq
}

// QueryNetworkMetrics chains the current query on the "network_metrics" edge.
func (snsq *SystemNetworkStatsQuery) QueryNetworkMetrics() *NetworkMetricsQuery {
	query := (&NetworkMetricsClient{config: snsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := snsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := snsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(systemnetworkstats.Table, systemnetworkstats.FieldID, selector),
			sqlgraph.To(networkmetrics.Table, networkmetrics.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, systemnetworkstats.NetworkMetricsTable, systemnetworkstats.NetworkMetricsColumn),
		)
		fromU = sqlgraph.SetNeighbors(snsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SystemNetworkStats entity from the query.
// Returns a *NotFoundError when no SystemNetworkStats was found.
func (snsq *SystemNetworkStatsQuery) First(ctx context.Context) (*SystemNetworkStats, error) {
	nodes, err := snsq.Limit(1).All(setContextOp(ctx, snsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{systemnetworkstats.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (snsq *SystemNetworkStatsQuery) FirstX(ctx context.Context) *SystemNetworkStats {
	node, err := snsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SystemNetworkStats ID from the query.
// Returns a *NotFoundError when no SystemNetworkStats ID was found.
func (snsq *SystemNetworkStatsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = snsq.Limit(1).IDs(setContextOp(ctx, snsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{systemnetworkstats.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (snsq *SystemNetworkStatsQuery) FirstIDX(ctx context.Context) int {
	id, err := snsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SystemNetworkStats entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SystemNetworkStats entity is found.
// Returns a *NotFoundError when no SystemNetworkStats entities are found.
func (snsq *SystemNetworkStatsQuery) Only(ctx context.Context) (*SystemNetworkStats, error) {
	nodes, err := snsq.Limit(2).All(setContextOp(ctx, snsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{systemnetworkstats.Label}
	default:
		return nil, &NotSingularError{systemnetworkstats.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (snsq *SystemNetworkStatsQuery) OnlyX(ctx context.Context) *SystemNetworkStats {
	node, err := snsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SystemNetworkStats ID in the query.
// Returns a *NotSingularError when more than one SystemNetworkStats ID is found.
// Returns a *NotFoundError when no entities are found.
func (snsq *SystemNetworkStatsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = snsq.Limit(2).IDs(setContextOp(ctx, snsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{systemnetworkstats.Label}
	default:
		err = &NotSingularError{systemnetworkstats.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (snsq *SystemNetworkStatsQuery) OnlyIDX(ctx context.Context) int {
	id, err := snsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SystemNetworkStatsSlice.
func (snsq *SystemNetworkStatsQuery) All(ctx context.Context) ([]*SystemNetworkStats, error) {
	ctx = setContextOp(ctx, snsq.ctx, "All")
	if err := snsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SystemNetworkStats, *SystemNetworkStatsQuery]()
	return withInterceptors[[]*SystemNetworkStats](ctx, snsq, qr, snsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (snsq *SystemNetworkStatsQuery) AllX(ctx context.Context) []*SystemNetworkStats {
	nodes, err := snsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SystemNetworkStats IDs.
func (snsq *SystemNetworkStatsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if snsq.ctx.Unique == nil && snsq.path != nil {
		snsq.Unique(true)
	}
	ctx = setContextOp(ctx, snsq.ctx, "IDs")
	if err = snsq.Select(systemnetworkstats.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (snsq *SystemNetworkStatsQuery) IDsX(ctx context.Context) []int {
	ids, err := snsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (snsq *SystemNetworkStatsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, snsq.ctx, "Count")
	if err := snsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, snsq, querierCount[*SystemNetworkStatsQuery](), snsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (snsq *SystemNetworkStatsQuery) CountX(ctx context.Context) int {
	count, err := snsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (snsq *SystemNetworkStatsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, snsq.ctx, "Exist")
	switch _, err := snsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (snsq *SystemNetworkStatsQuery) ExistX(ctx context.Context) bool {
	exist, err := snsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SystemNetworkStatsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (snsq *SystemNetworkStatsQuery) Clone() *SystemNetworkStatsQuery {
	if snsq == nil {
		return nil
	}
	return &SystemNetworkStatsQuery{
		config:             snsq.config,
		ctx:                snsq.ctx.Clone(),
		order:              append([]systemnetworkstats.OrderOption{}, snsq.order...),
		inters:             append([]Interceptor{}, snsq.inters...),
		predicates:         append([]predicate.SystemNetworkStats{}, snsq.predicates...),
		withNetworkMetrics: snsq.withNetworkMetrics.Clone(),
		// clone intermediate query.
		sql:  snsq.sql.Clone(),
		path: snsq.path,
	}
}

// WithNetworkMetrics tells the query-builder to eager-load the nodes that are connected to
// the "network_metrics" edge. The optional arguments are used to configure the query builder of the edge.
func (snsq *SystemNetworkStatsQuery) WithNetworkMetrics(opts ...func(*NetworkMetricsQuery)) *SystemNetworkStatsQuery {
	query := (&NetworkMetricsClient{config: snsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	snsq.withNetworkMetrics = query
	return snsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BytesSent int64 `json:"bytes_sent,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SystemNetworkStats.Query().
//		GroupBy(systemnetworkstats.FieldBytesSent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (snsq *SystemNetworkStatsQuery) GroupBy(field string, fields ...string) *SystemNetworkStatsGroupBy {
	snsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SystemNetworkStatsGroupBy{build: snsq}
	grbuild.flds = &snsq.ctx.Fields
	grbuild.label = systemnetworkstats.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BytesSent int64 `json:"bytes_sent,omitempty"`
//	}
//
//	client.SystemNetworkStats.Query().
//		Select(systemnetworkstats.FieldBytesSent).
//		Scan(ctx, &v)
func (snsq *SystemNetworkStatsQuery) Select(fields ...string) *SystemNetworkStatsSelect {
	snsq.ctx.Fields = append(snsq.ctx.Fields, fields...)
	sbuild := &SystemNetworkStatsSelect{SystemNetworkStatsQuery: snsq}
	sbuild.label = systemnetworkstats.Label
	sbuild.flds, sbuild.scan = &snsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SystemNetworkStatsSelect configured with the given aggregations.
func (snsq *SystemNetworkStatsQuery) Aggregate(fns ...AggregateFunc) *SystemNetworkStatsSelect {
	return snsq.Select().Aggregate(fns...)
}

func (snsq *SystemNetworkStatsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range snsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, snsq); err != nil {
				return err
			}
		}
	}
	for _, f := range snsq.ctx.Fields {
		if !systemnetworkstats.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if snsq.path != nil {
		prev, err := snsq.path(ctx)
		if err != nil {
			return err
		}
		snsq.sql = prev
	}
	return nil
}

func (snsq *SystemNetworkStatsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SystemNetworkStats, error) {
	var (
		nodes       = []*SystemNetworkStats{}
		withFKs     = snsq.withFKs
		_spec       = snsq.querySpec()
		loadedTypes = [1]bool{
			snsq.withNetworkMetrics != nil,
		}
	)
	if snsq.withNetworkMetrics != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, systemnetworkstats.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SystemNetworkStats).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SystemNetworkStats{config: snsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(snsq.modifiers) > 0 {
		_spec.Modifiers = snsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, snsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := snsq.withNetworkMetrics; query != nil {
		if err := snsq.loadNetworkMetrics(ctx, query, nodes, nil,
			func(n *SystemNetworkStats, e *NetworkMetrics) { n.Edges.NetworkMetrics = e }); err != nil {
			return nil, err
		}
	}
	for i := range snsq.loadTotal {
		if err := snsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (snsq *SystemNetworkStatsQuery) loadNetworkMetrics(ctx context.Context, query *NetworkMetricsQuery, nodes []*SystemNetworkStats, init func(*SystemNetworkStats), assign func(*SystemNetworkStats, *NetworkMetrics)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SystemNetworkStats)
	for i := range nodes {
		if nodes[i].network_metrics_system_network_stats == nil {
			continue
		}
		fk := *nodes[i].network_metrics_system_network_stats
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(networkmetrics.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "network_metrics_system_network_stats" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (snsq *SystemNetworkStatsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := snsq.querySpec()
	if len(snsq.modifiers) > 0 {
		_spec.Modifiers = snsq.modifiers
	}
	_spec.Node.Columns = snsq.ctx.Fields
	if len(snsq.ctx.Fields) > 0 {
		_spec.Unique = snsq.ctx.Unique != nil && *snsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, snsq.driver, _spec)
}

func (snsq *SystemNetworkStatsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(systemnetworkstats.Table, systemnetworkstats.Columns, sqlgraph.NewFieldSpec(systemnetworkstats.FieldID, field.TypeInt))
	_spec.From = snsq.sql
	if unique := snsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if snsq.path != nil {
		_spec.Unique = true
	}
	if fields := snsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemnetworkstats.FieldID)
		for i := range fields {
			if fields[i] != systemnetworkstats.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := snsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := snsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := snsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := snsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (snsq *SystemNetworkStatsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(snsq.driver.Dialect())
	t1 := builder.Table(systemnetworkstats.Table)
	columns := snsq.ctx.Fields
	if len(columns) == 0 {
		columns = systemnetworkstats.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if snsq.sql != nil {
		selector = snsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if snsq.ctx.Unique != nil && *snsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range snsq.predicates {
		p(selector)
	}
	for _, p := range snsq.order {
		p(selector)
	}
	if offset := snsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := snsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SystemNetworkStatsGroupBy is the group-by builder for SystemNetworkStats entities.
type SystemNetworkStatsGroupBy struct {
	selector
	build *SystemNetworkStatsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (snsgb *SystemNetworkStatsGroupBy) Aggregate(fns ...AggregateFunc) *SystemNetworkStatsGroupBy {
	snsgb.fns = append(snsgb.fns, fns...)
	return snsgb
}

// Scan applies the selector query and scans the result into the given value.
func (snsgb *SystemNetworkStatsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, snsgb.build.ctx, "GroupBy")
	if err := snsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SystemNetworkStatsQuery, *SystemNetworkStatsGroupBy](ctx, snsgb.build, snsgb, snsgb.build.inters, v)
}

func (snsgb *SystemNetworkStatsGroupBy) sqlScan(ctx context.Context, root *SystemNetworkStatsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(snsgb.fns))
	for _, fn := range snsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*snsgb.flds)+len(snsgb.fns))
		for _, f := range *snsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*snsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := snsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SystemNetworkStatsSelect is the builder for selecting fields of SystemNetworkStats entities.
type SystemNetworkStatsSelect struct {
	*SystemNetworkStatsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (snss *SystemNetworkStatsSelect) Aggregate(fns ...AggregateFunc) *SystemNetworkStatsSelect {
	snss.fns = append(snss.fns, fns...)
	return snss
}

// Scan applies the selector query and scans the result into the given value.
func (snss *SystemNetworkStatsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, snss.ctx, "Select")
	if err := snss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SystemNetworkStatsQuery, *SystemNetworkStatsSelect](ctx, snss.SystemNetworkStatsQuery, snss, snss.inters, v)
}

func (snss *SystemNetworkStatsSelect) sqlScan(ctx context.Context, root *SystemNetworkStatsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(snss.fns))
	for _, fn := range snss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*snss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := snss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}