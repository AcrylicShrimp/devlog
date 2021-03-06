// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postvideo"
	"devlog/ent/predicate"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostVideoQuery is the builder for querying PostVideo entities.
type PostVideoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PostVideo
	// eager-loading edges.
	withPost *PostQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PostVideoQuery builder.
func (pvq *PostVideoQuery) Where(ps ...predicate.PostVideo) *PostVideoQuery {
	pvq.predicates = append(pvq.predicates, ps...)
	return pvq
}

// Limit adds a limit step to the query.
func (pvq *PostVideoQuery) Limit(limit int) *PostVideoQuery {
	pvq.limit = &limit
	return pvq
}

// Offset adds an offset step to the query.
func (pvq *PostVideoQuery) Offset(offset int) *PostVideoQuery {
	pvq.offset = &offset
	return pvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pvq *PostVideoQuery) Unique(unique bool) *PostVideoQuery {
	pvq.unique = &unique
	return pvq
}

// Order adds an order step to the query.
func (pvq *PostVideoQuery) Order(o ...OrderFunc) *PostVideoQuery {
	pvq.order = append(pvq.order, o...)
	return pvq
}

// QueryPost chains the current query on the "post" edge.
func (pvq *PostVideoQuery) QueryPost() *PostQuery {
	query := &PostQuery{config: pvq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(postvideo.Table, postvideo.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, postvideo.PostTable, postvideo.PostColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PostVideo entity from the query.
// Returns a *NotFoundError when no PostVideo was found.
func (pvq *PostVideoQuery) First(ctx context.Context) (*PostVideo, error) {
	nodes, err := pvq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{postvideo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pvq *PostVideoQuery) FirstX(ctx context.Context) *PostVideo {
	node, err := pvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PostVideo ID from the query.
// Returns a *NotFoundError when no PostVideo ID was found.
func (pvq *PostVideoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{postvideo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pvq *PostVideoQuery) FirstIDX(ctx context.Context) int {
	id, err := pvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PostVideo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PostVideo entity is not found.
// Returns a *NotFoundError when no PostVideo entities are found.
func (pvq *PostVideoQuery) Only(ctx context.Context) (*PostVideo, error) {
	nodes, err := pvq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{postvideo.Label}
	default:
		return nil, &NotSingularError{postvideo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pvq *PostVideoQuery) OnlyX(ctx context.Context) *PostVideo {
	node, err := pvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PostVideo ID in the query.
// Returns a *NotSingularError when exactly one PostVideo ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pvq *PostVideoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = &NotSingularError{postvideo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pvq *PostVideoQuery) OnlyIDX(ctx context.Context) int {
	id, err := pvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PostVideos.
func (pvq *PostVideoQuery) All(ctx context.Context) ([]*PostVideo, error) {
	if err := pvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pvq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pvq *PostVideoQuery) AllX(ctx context.Context) []*PostVideo {
	nodes, err := pvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PostVideo IDs.
func (pvq *PostVideoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pvq.Select(postvideo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pvq *PostVideoQuery) IDsX(ctx context.Context) []int {
	ids, err := pvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pvq *PostVideoQuery) Count(ctx context.Context) (int, error) {
	if err := pvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pvq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pvq *PostVideoQuery) CountX(ctx context.Context) int {
	count, err := pvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pvq *PostVideoQuery) Exist(ctx context.Context) (bool, error) {
	if err := pvq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pvq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pvq *PostVideoQuery) ExistX(ctx context.Context) bool {
	exist, err := pvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PostVideoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pvq *PostVideoQuery) Clone() *PostVideoQuery {
	if pvq == nil {
		return nil
	}
	return &PostVideoQuery{
		config:     pvq.config,
		limit:      pvq.limit,
		offset:     pvq.offset,
		order:      append([]OrderFunc{}, pvq.order...),
		predicates: append([]predicate.PostVideo{}, pvq.predicates...),
		withPost:   pvq.withPost.Clone(),
		// clone intermediate query.
		sql:  pvq.sql.Clone(),
		path: pvq.path,
	}
}

// WithPost tells the query-builder to eager-load the nodes that are connected to
// the "post" edge. The optional arguments are used to configure the query builder of the edge.
func (pvq *PostVideoQuery) WithPost(opts ...func(*PostQuery)) *PostVideoQuery {
	query := &PostQuery{config: pvq.config}
	for _, opt := range opts {
		opt(query)
	}
	pvq.withPost = query
	return pvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PostVideo.Query().
//		GroupBy(postvideo.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pvq *PostVideoQuery) GroupBy(field string, fields ...string) *PostVideoGroupBy {
	group := &PostVideoGroupBy{config: pvq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pvq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//	}
//
//	client.PostVideo.Query().
//		Select(postvideo.FieldUUID).
//		Scan(ctx, &v)
//
func (pvq *PostVideoQuery) Select(field string, fields ...string) *PostVideoSelect {
	pvq.fields = append([]string{field}, fields...)
	return &PostVideoSelect{PostVideoQuery: pvq}
}

func (pvq *PostVideoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pvq.fields {
		if !postvideo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pvq.path != nil {
		prev, err := pvq.path(ctx)
		if err != nil {
			return err
		}
		pvq.sql = prev
	}
	return nil
}

func (pvq *PostVideoQuery) sqlAll(ctx context.Context) ([]*PostVideo, error) {
	var (
		nodes       = []*PostVideo{}
		withFKs     = pvq.withFKs
		_spec       = pvq.querySpec()
		loadedTypes = [1]bool{
			pvq.withPost != nil,
		}
	)
	if pvq.withPost != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, postvideo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PostVideo{config: pvq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pvq.withPost; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PostVideo)
		for i := range nodes {
			if nodes[i].post_videos == nil {
				continue
			}
			fk := *nodes[i].post_videos
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(post.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "post_videos" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Post = n
			}
		}
	}

	return nodes, nil
}

func (pvq *PostVideoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pvq.querySpec()
	return sqlgraph.CountNodes(ctx, pvq.driver, _spec)
}

func (pvq *PostVideoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pvq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pvq *PostVideoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postvideo.Table,
			Columns: postvideo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postvideo.FieldID,
			},
		},
		From:   pvq.sql,
		Unique: true,
	}
	if unique := pvq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pvq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postvideo.FieldID)
		for i := range fields {
			if fields[i] != postvideo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pvq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pvq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pvq *PostVideoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pvq.driver.Dialect())
	t1 := builder.Table(postvideo.Table)
	selector := builder.Select(t1.Columns(postvideo.Columns...)...).From(t1)
	if pvq.sql != nil {
		selector = pvq.sql
		selector.Select(selector.Columns(postvideo.Columns...)...)
	}
	for _, p := range pvq.predicates {
		p(selector)
	}
	for _, p := range pvq.order {
		p(selector)
	}
	if offset := pvq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pvq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PostVideoGroupBy is the group-by builder for PostVideo entities.
type PostVideoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pvgb *PostVideoGroupBy) Aggregate(fns ...AggregateFunc) *PostVideoGroupBy {
	pvgb.fns = append(pvgb.fns, fns...)
	return pvgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pvgb *PostVideoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pvgb.path(ctx)
	if err != nil {
		return err
	}
	pvgb.sql = query
	return pvgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pvgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pvgb *PostVideoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pvgb.fields) > 1 {
		return nil, errors.New("ent: PostVideoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) StringsX(ctx context.Context) []string {
	v, err := pvgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pvgb *PostVideoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pvgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = fmt.Errorf("ent: PostVideoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) StringX(ctx context.Context) string {
	v, err := pvgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pvgb *PostVideoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pvgb.fields) > 1 {
		return nil, errors.New("ent: PostVideoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) IntsX(ctx context.Context) []int {
	v, err := pvgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pvgb *PostVideoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pvgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = fmt.Errorf("ent: PostVideoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) IntX(ctx context.Context) int {
	v, err := pvgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pvgb *PostVideoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pvgb.fields) > 1 {
		return nil, errors.New("ent: PostVideoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pvgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pvgb *PostVideoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pvgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = fmt.Errorf("ent: PostVideoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pvgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pvgb *PostVideoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pvgb.fields) > 1 {
		return nil, errors.New("ent: PostVideoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pvgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pvgb *PostVideoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pvgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = fmt.Errorf("ent: PostVideoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pvgb *PostVideoGroupBy) BoolX(ctx context.Context) bool {
	v, err := pvgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvgb *PostVideoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pvgb.fields {
		if !postvideo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pvgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pvgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pvgb *PostVideoGroupBy) sqlQuery() *sql.Selector {
	selector := pvgb.sql
	columns := make([]string, 0, len(pvgb.fields)+len(pvgb.fns))
	columns = append(columns, pvgb.fields...)
	for _, fn := range pvgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pvgb.fields...)
}

// PostVideoSelect is the builder for selecting fields of PostVideo entities.
type PostVideoSelect struct {
	*PostVideoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pvs *PostVideoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pvs.prepareQuery(ctx); err != nil {
		return err
	}
	pvs.sql = pvs.PostVideoQuery.sqlQuery(ctx)
	return pvs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pvs *PostVideoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pvs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pvs *PostVideoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pvs.fields) > 1 {
		return nil, errors.New("ent: PostVideoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pvs *PostVideoSelect) StringsX(ctx context.Context) []string {
	v, err := pvs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pvs *PostVideoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pvs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = fmt.Errorf("ent: PostVideoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pvs *PostVideoSelect) StringX(ctx context.Context) string {
	v, err := pvs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pvs *PostVideoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pvs.fields) > 1 {
		return nil, errors.New("ent: PostVideoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pvs *PostVideoSelect) IntsX(ctx context.Context) []int {
	v, err := pvs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pvs *PostVideoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pvs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = fmt.Errorf("ent: PostVideoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pvs *PostVideoSelect) IntX(ctx context.Context) int {
	v, err := pvs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pvs *PostVideoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pvs.fields) > 1 {
		return nil, errors.New("ent: PostVideoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pvs *PostVideoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pvs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pvs *PostVideoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pvs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = fmt.Errorf("ent: PostVideoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pvs *PostVideoSelect) Float64X(ctx context.Context) float64 {
	v, err := pvs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pvs *PostVideoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pvs.fields) > 1 {
		return nil, errors.New("ent: PostVideoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pvs *PostVideoSelect) BoolsX(ctx context.Context) []bool {
	v, err := pvs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pvs *PostVideoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pvs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postvideo.Label}
	default:
		err = fmt.Errorf("ent: PostVideoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pvs *PostVideoSelect) BoolX(ctx context.Context) bool {
	v, err := pvs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvs *PostVideoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pvs.sqlQuery().Query()
	if err := pvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pvs *PostVideoSelect) sqlQuery() sql.Querier {
	selector := pvs.sql
	selector.Select(selector.Columns(pvs.fields...)...)
	return selector
}
