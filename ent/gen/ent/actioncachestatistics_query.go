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
	"github.com/buildbarn/bb-portal/ent/gen/ent/actioncachestatistics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/actionsummary"
	"github.com/buildbarn/bb-portal/ent/gen/ent/missdetail"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// ActionCacheStatisticsQuery is the builder for querying ActionCacheStatistics entities.
type ActionCacheStatisticsQuery struct {
	config
	ctx                    *QueryContext
	order                  []actioncachestatistics.OrderOption
	inters                 []Interceptor
	predicates             []predicate.ActionCacheStatistics
	withActionSummary      *ActionSummaryQuery
	withMissDetails        *MissDetailQuery
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*ActionCacheStatistics) error
	withNamedActionSummary map[string]*ActionSummaryQuery
	withNamedMissDetails   map[string]*MissDetailQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ActionCacheStatisticsQuery builder.
func (acsq *ActionCacheStatisticsQuery) Where(ps ...predicate.ActionCacheStatistics) *ActionCacheStatisticsQuery {
	acsq.predicates = append(acsq.predicates, ps...)
	return acsq
}

// Limit the number of records to be returned by this query.
func (acsq *ActionCacheStatisticsQuery) Limit(limit int) *ActionCacheStatisticsQuery {
	acsq.ctx.Limit = &limit
	return acsq
}

// Offset to start from.
func (acsq *ActionCacheStatisticsQuery) Offset(offset int) *ActionCacheStatisticsQuery {
	acsq.ctx.Offset = &offset
	return acsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acsq *ActionCacheStatisticsQuery) Unique(unique bool) *ActionCacheStatisticsQuery {
	acsq.ctx.Unique = &unique
	return acsq
}

// Order specifies how the records should be ordered.
func (acsq *ActionCacheStatisticsQuery) Order(o ...actioncachestatistics.OrderOption) *ActionCacheStatisticsQuery {
	acsq.order = append(acsq.order, o...)
	return acsq
}

// QueryActionSummary chains the current query on the "action_summary" edge.
func (acsq *ActionCacheStatisticsQuery) QueryActionSummary() *ActionSummaryQuery {
	query := (&ActionSummaryClient{config: acsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(actioncachestatistics.Table, actioncachestatistics.FieldID, selector),
			sqlgraph.To(actionsummary.Table, actionsummary.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, actioncachestatistics.ActionSummaryTable, actioncachestatistics.ActionSummaryPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(acsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissDetails chains the current query on the "miss_details" edge.
func (acsq *ActionCacheStatisticsQuery) QueryMissDetails() *MissDetailQuery {
	query := (&MissDetailClient{config: acsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(actioncachestatistics.Table, actioncachestatistics.FieldID, selector),
			sqlgraph.To(missdetail.Table, missdetail.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, actioncachestatistics.MissDetailsTable, actioncachestatistics.MissDetailsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(acsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ActionCacheStatistics entity from the query.
// Returns a *NotFoundError when no ActionCacheStatistics was found.
func (acsq *ActionCacheStatisticsQuery) First(ctx context.Context) (*ActionCacheStatistics, error) {
	nodes, err := acsq.Limit(1).All(setContextOp(ctx, acsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{actioncachestatistics.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acsq *ActionCacheStatisticsQuery) FirstX(ctx context.Context) *ActionCacheStatistics {
	node, err := acsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ActionCacheStatistics ID from the query.
// Returns a *NotFoundError when no ActionCacheStatistics ID was found.
func (acsq *ActionCacheStatisticsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = acsq.Limit(1).IDs(setContextOp(ctx, acsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{actioncachestatistics.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acsq *ActionCacheStatisticsQuery) FirstIDX(ctx context.Context) int {
	id, err := acsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ActionCacheStatistics entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ActionCacheStatistics entity is found.
// Returns a *NotFoundError when no ActionCacheStatistics entities are found.
func (acsq *ActionCacheStatisticsQuery) Only(ctx context.Context) (*ActionCacheStatistics, error) {
	nodes, err := acsq.Limit(2).All(setContextOp(ctx, acsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{actioncachestatistics.Label}
	default:
		return nil, &NotSingularError{actioncachestatistics.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acsq *ActionCacheStatisticsQuery) OnlyX(ctx context.Context) *ActionCacheStatistics {
	node, err := acsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ActionCacheStatistics ID in the query.
// Returns a *NotSingularError when more than one ActionCacheStatistics ID is found.
// Returns a *NotFoundError when no entities are found.
func (acsq *ActionCacheStatisticsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = acsq.Limit(2).IDs(setContextOp(ctx, acsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{actioncachestatistics.Label}
	default:
		err = &NotSingularError{actioncachestatistics.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acsq *ActionCacheStatisticsQuery) OnlyIDX(ctx context.Context) int {
	id, err := acsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ActionCacheStatisticsSlice.
func (acsq *ActionCacheStatisticsQuery) All(ctx context.Context) ([]*ActionCacheStatistics, error) {
	ctx = setContextOp(ctx, acsq.ctx, "All")
	if err := acsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ActionCacheStatistics, *ActionCacheStatisticsQuery]()
	return withInterceptors[[]*ActionCacheStatistics](ctx, acsq, qr, acsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acsq *ActionCacheStatisticsQuery) AllX(ctx context.Context) []*ActionCacheStatistics {
	nodes, err := acsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ActionCacheStatistics IDs.
func (acsq *ActionCacheStatisticsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if acsq.ctx.Unique == nil && acsq.path != nil {
		acsq.Unique(true)
	}
	ctx = setContextOp(ctx, acsq.ctx, "IDs")
	if err = acsq.Select(actioncachestatistics.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acsq *ActionCacheStatisticsQuery) IDsX(ctx context.Context) []int {
	ids, err := acsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acsq *ActionCacheStatisticsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acsq.ctx, "Count")
	if err := acsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acsq, querierCount[*ActionCacheStatisticsQuery](), acsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acsq *ActionCacheStatisticsQuery) CountX(ctx context.Context) int {
	count, err := acsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acsq *ActionCacheStatisticsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acsq.ctx, "Exist")
	switch _, err := acsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acsq *ActionCacheStatisticsQuery) ExistX(ctx context.Context) bool {
	exist, err := acsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ActionCacheStatisticsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acsq *ActionCacheStatisticsQuery) Clone() *ActionCacheStatisticsQuery {
	if acsq == nil {
		return nil
	}
	return &ActionCacheStatisticsQuery{
		config:            acsq.config,
		ctx:               acsq.ctx.Clone(),
		order:             append([]actioncachestatistics.OrderOption{}, acsq.order...),
		inters:            append([]Interceptor{}, acsq.inters...),
		predicates:        append([]predicate.ActionCacheStatistics{}, acsq.predicates...),
		withActionSummary: acsq.withActionSummary.Clone(),
		withMissDetails:   acsq.withMissDetails.Clone(),
		// clone intermediate query.
		sql:  acsq.sql.Clone(),
		path: acsq.path,
	}
}

// WithActionSummary tells the query-builder to eager-load the nodes that are connected to
// the "action_summary" edge. The optional arguments are used to configure the query builder of the edge.
func (acsq *ActionCacheStatisticsQuery) WithActionSummary(opts ...func(*ActionSummaryQuery)) *ActionCacheStatisticsQuery {
	query := (&ActionSummaryClient{config: acsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acsq.withActionSummary = query
	return acsq
}

// WithMissDetails tells the query-builder to eager-load the nodes that are connected to
// the "miss_details" edge. The optional arguments are used to configure the query builder of the edge.
func (acsq *ActionCacheStatisticsQuery) WithMissDetails(opts ...func(*MissDetailQuery)) *ActionCacheStatisticsQuery {
	query := (&MissDetailClient{config: acsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acsq.withMissDetails = query
	return acsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SizeInBytes int64 `json:"size_in_bytes,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ActionCacheStatistics.Query().
//		GroupBy(actioncachestatistics.FieldSizeInBytes).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (acsq *ActionCacheStatisticsQuery) GroupBy(field string, fields ...string) *ActionCacheStatisticsGroupBy {
	acsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ActionCacheStatisticsGroupBy{build: acsq}
	grbuild.flds = &acsq.ctx.Fields
	grbuild.label = actioncachestatistics.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SizeInBytes int64 `json:"size_in_bytes,omitempty"`
//	}
//
//	client.ActionCacheStatistics.Query().
//		Select(actioncachestatistics.FieldSizeInBytes).
//		Scan(ctx, &v)
func (acsq *ActionCacheStatisticsQuery) Select(fields ...string) *ActionCacheStatisticsSelect {
	acsq.ctx.Fields = append(acsq.ctx.Fields, fields...)
	sbuild := &ActionCacheStatisticsSelect{ActionCacheStatisticsQuery: acsq}
	sbuild.label = actioncachestatistics.Label
	sbuild.flds, sbuild.scan = &acsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ActionCacheStatisticsSelect configured with the given aggregations.
func (acsq *ActionCacheStatisticsQuery) Aggregate(fns ...AggregateFunc) *ActionCacheStatisticsSelect {
	return acsq.Select().Aggregate(fns...)
}

func (acsq *ActionCacheStatisticsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acsq); err != nil {
				return err
			}
		}
	}
	for _, f := range acsq.ctx.Fields {
		if !actioncachestatistics.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acsq.path != nil {
		prev, err := acsq.path(ctx)
		if err != nil {
			return err
		}
		acsq.sql = prev
	}
	return nil
}

func (acsq *ActionCacheStatisticsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ActionCacheStatistics, error) {
	var (
		nodes       = []*ActionCacheStatistics{}
		_spec       = acsq.querySpec()
		loadedTypes = [2]bool{
			acsq.withActionSummary != nil,
			acsq.withMissDetails != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ActionCacheStatistics).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ActionCacheStatistics{config: acsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(acsq.modifiers) > 0 {
		_spec.Modifiers = acsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := acsq.withActionSummary; query != nil {
		if err := acsq.loadActionSummary(ctx, query, nodes,
			func(n *ActionCacheStatistics) { n.Edges.ActionSummary = []*ActionSummary{} },
			func(n *ActionCacheStatistics, e *ActionSummary) {
				n.Edges.ActionSummary = append(n.Edges.ActionSummary, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := acsq.withMissDetails; query != nil {
		if err := acsq.loadMissDetails(ctx, query, nodes,
			func(n *ActionCacheStatistics) { n.Edges.MissDetails = []*MissDetail{} },
			func(n *ActionCacheStatistics, e *MissDetail) { n.Edges.MissDetails = append(n.Edges.MissDetails, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range acsq.withNamedActionSummary {
		if err := acsq.loadActionSummary(ctx, query, nodes,
			func(n *ActionCacheStatistics) { n.appendNamedActionSummary(name) },
			func(n *ActionCacheStatistics, e *ActionSummary) { n.appendNamedActionSummary(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range acsq.withNamedMissDetails {
		if err := acsq.loadMissDetails(ctx, query, nodes,
			func(n *ActionCacheStatistics) { n.appendNamedMissDetails(name) },
			func(n *ActionCacheStatistics, e *MissDetail) { n.appendNamedMissDetails(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range acsq.loadTotal {
		if err := acsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (acsq *ActionCacheStatisticsQuery) loadActionSummary(ctx context.Context, query *ActionSummaryQuery, nodes []*ActionCacheStatistics, init func(*ActionCacheStatistics), assign func(*ActionCacheStatistics, *ActionSummary)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ActionCacheStatistics)
	nids := make(map[int]map[*ActionCacheStatistics]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(actioncachestatistics.ActionSummaryTable)
		s.Join(joinT).On(s.C(actionsummary.FieldID), joinT.C(actioncachestatistics.ActionSummaryPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(actioncachestatistics.ActionSummaryPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(actioncachestatistics.ActionSummaryPrimaryKey[1]))
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
					nids[inValue] = map[*ActionCacheStatistics]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*ActionSummary](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "action_summary" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (acsq *ActionCacheStatisticsQuery) loadMissDetails(ctx context.Context, query *MissDetailQuery, nodes []*ActionCacheStatistics, init func(*ActionCacheStatistics), assign func(*ActionCacheStatistics, *MissDetail)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ActionCacheStatistics)
	nids := make(map[int]map[*ActionCacheStatistics]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(actioncachestatistics.MissDetailsTable)
		s.Join(joinT).On(s.C(missdetail.FieldID), joinT.C(actioncachestatistics.MissDetailsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(actioncachestatistics.MissDetailsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(actioncachestatistics.MissDetailsPrimaryKey[0]))
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
					nids[inValue] = map[*ActionCacheStatistics]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*MissDetail](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "miss_details" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (acsq *ActionCacheStatisticsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acsq.querySpec()
	if len(acsq.modifiers) > 0 {
		_spec.Modifiers = acsq.modifiers
	}
	_spec.Node.Columns = acsq.ctx.Fields
	if len(acsq.ctx.Fields) > 0 {
		_spec.Unique = acsq.ctx.Unique != nil && *acsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acsq.driver, _spec)
}

func (acsq *ActionCacheStatisticsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(actioncachestatistics.Table, actioncachestatistics.Columns, sqlgraph.NewFieldSpec(actioncachestatistics.FieldID, field.TypeInt))
	_spec.From = acsq.sql
	if unique := acsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acsq.path != nil {
		_spec.Unique = true
	}
	if fields := acsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, actioncachestatistics.FieldID)
		for i := range fields {
			if fields[i] != actioncachestatistics.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acsq *ActionCacheStatisticsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acsq.driver.Dialect())
	t1 := builder.Table(actioncachestatistics.Table)
	columns := acsq.ctx.Fields
	if len(columns) == 0 {
		columns = actioncachestatistics.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acsq.sql != nil {
		selector = acsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acsq.ctx.Unique != nil && *acsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range acsq.predicates {
		p(selector)
	}
	for _, p := range acsq.order {
		p(selector)
	}
	if offset := acsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedActionSummary tells the query-builder to eager-load the nodes that are connected to the "action_summary"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (acsq *ActionCacheStatisticsQuery) WithNamedActionSummary(name string, opts ...func(*ActionSummaryQuery)) *ActionCacheStatisticsQuery {
	query := (&ActionSummaryClient{config: acsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if acsq.withNamedActionSummary == nil {
		acsq.withNamedActionSummary = make(map[string]*ActionSummaryQuery)
	}
	acsq.withNamedActionSummary[name] = query
	return acsq
}

// WithNamedMissDetails tells the query-builder to eager-load the nodes that are connected to the "miss_details"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (acsq *ActionCacheStatisticsQuery) WithNamedMissDetails(name string, opts ...func(*MissDetailQuery)) *ActionCacheStatisticsQuery {
	query := (&MissDetailClient{config: acsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if acsq.withNamedMissDetails == nil {
		acsq.withNamedMissDetails = make(map[string]*MissDetailQuery)
	}
	acsq.withNamedMissDetails[name] = query
	return acsq
}

// ActionCacheStatisticsGroupBy is the group-by builder for ActionCacheStatistics entities.
type ActionCacheStatisticsGroupBy struct {
	selector
	build *ActionCacheStatisticsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acsgb *ActionCacheStatisticsGroupBy) Aggregate(fns ...AggregateFunc) *ActionCacheStatisticsGroupBy {
	acsgb.fns = append(acsgb.fns, fns...)
	return acsgb
}

// Scan applies the selector query and scans the result into the given value.
func (acsgb *ActionCacheStatisticsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acsgb.build.ctx, "GroupBy")
	if err := acsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActionCacheStatisticsQuery, *ActionCacheStatisticsGroupBy](ctx, acsgb.build, acsgb, acsgb.build.inters, v)
}

func (acsgb *ActionCacheStatisticsGroupBy) sqlScan(ctx context.Context, root *ActionCacheStatisticsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acsgb.fns))
	for _, fn := range acsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acsgb.flds)+len(acsgb.fns))
		for _, f := range *acsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ActionCacheStatisticsSelect is the builder for selecting fields of ActionCacheStatistics entities.
type ActionCacheStatisticsSelect struct {
	*ActionCacheStatisticsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acss *ActionCacheStatisticsSelect) Aggregate(fns ...AggregateFunc) *ActionCacheStatisticsSelect {
	acss.fns = append(acss.fns, fns...)
	return acss
}

// Scan applies the selector query and scans the result into the given value.
func (acss *ActionCacheStatisticsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acss.ctx, "Select")
	if err := acss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActionCacheStatisticsQuery, *ActionCacheStatisticsSelect](ctx, acss.ActionCacheStatisticsQuery, acss, acss.inters, v)
}

func (acss *ActionCacheStatisticsSelect) sqlScan(ctx context.Context, root *ActionCacheStatisticsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acss.fns))
	for _, fn := range acss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
