// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postimage"
	"devlog/ent/predicate"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostImageQuery is the builder for querying PostImage entities.
type PostImageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PostImage
	// eager-loading edges.
	withPost *PostQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PostImageQuery builder.
func (piq *PostImageQuery) Where(ps ...predicate.PostImage) *PostImageQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit adds a limit step to the query.
func (piq *PostImageQuery) Limit(limit int) *PostImageQuery {
	piq.limit = &limit
	return piq
}

// Offset adds an offset step to the query.
func (piq *PostImageQuery) Offset(offset int) *PostImageQuery {
	piq.offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *PostImageQuery) Unique(unique bool) *PostImageQuery {
	piq.unique = &unique
	return piq
}

// Order adds an order step to the query.
func (piq *PostImageQuery) Order(o ...OrderFunc) *PostImageQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// QueryPost chains the current query on the "post" edge.
func (piq *PostImageQuery) QueryPost() *PostQuery {
	query := &PostQuery{config: piq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(postimage.Table, postimage.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, postimage.PostTable, postimage.PostColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PostImage entity from the query.
// Returns a *NotFoundError when no PostImage was found.
func (piq *PostImageQuery) First(ctx context.Context) (*PostImage, error) {
	nodes, err := piq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{postimage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *PostImageQuery) FirstX(ctx context.Context) *PostImage {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PostImage ID from the query.
// Returns a *NotFoundError when no PostImage ID was found.
func (piq *PostImageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{postimage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *PostImageQuery) FirstIDX(ctx context.Context) int {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PostImage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PostImage entity is not found.
// Returns a *NotFoundError when no PostImage entities are found.
func (piq *PostImageQuery) Only(ctx context.Context) (*PostImage, error) {
	nodes, err := piq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{postimage.Label}
	default:
		return nil, &NotSingularError{postimage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *PostImageQuery) OnlyX(ctx context.Context) *PostImage {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PostImage ID in the query.
// Returns a *NotSingularError when exactly one PostImage ID is not found.
// Returns a *NotFoundError when no entities are found.
func (piq *PostImageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = &NotSingularError{postimage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *PostImageQuery) OnlyIDX(ctx context.Context) int {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PostImages.
func (piq *PostImageQuery) All(ctx context.Context) ([]*PostImage, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return piq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (piq *PostImageQuery) AllX(ctx context.Context) []*PostImage {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PostImage IDs.
func (piq *PostImageQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := piq.Select(postimage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *PostImageQuery) IDsX(ctx context.Context) []int {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *PostImageQuery) Count(ctx context.Context) (int, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return piq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (piq *PostImageQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *PostImageQuery) Exist(ctx context.Context) (bool, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return piq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *PostImageQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PostImageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *PostImageQuery) Clone() *PostImageQuery {
	if piq == nil {
		return nil
	}
	return &PostImageQuery{
		config:     piq.config,
		limit:      piq.limit,
		offset:     piq.offset,
		order:      append([]OrderFunc{}, piq.order...),
		predicates: append([]predicate.PostImage{}, piq.predicates...),
		withPost:   piq.withPost.Clone(),
		// clone intermediate query.
		sql:  piq.sql.Clone(),
		path: piq.path,
	}
}

// WithPost tells the query-builder to eager-load the nodes that are connected to
// the "post" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *PostImageQuery) WithPost(opts ...func(*PostQuery)) *PostImageQuery {
	query := &PostQuery{config: piq.config}
	for _, opt := range opts {
		opt(query)
	}
	piq.withPost = query
	return piq
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
//	client.PostImage.Query().
//		GroupBy(postimage.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (piq *PostImageQuery) GroupBy(field string, fields ...string) *PostImageGroupBy {
	group := &PostImageGroupBy{config: piq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return piq.sqlQuery(ctx), nil
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
//	client.PostImage.Query().
//		Select(postimage.FieldUUID).
//		Scan(ctx, &v)
//
func (piq *PostImageQuery) Select(field string, fields ...string) *PostImageSelect {
	piq.fields = append([]string{field}, fields...)
	return &PostImageSelect{PostImageQuery: piq}
}

func (piq *PostImageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range piq.fields {
		if !postimage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *PostImageQuery) sqlAll(ctx context.Context) ([]*PostImage, error) {
	var (
		nodes       = []*PostImage{}
		withFKs     = piq.withFKs
		_spec       = piq.querySpec()
		loadedTypes = [1]bool{
			piq.withPost != nil,
		}
	)
	if piq.withPost != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, postimage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PostImage{config: piq.config}
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
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := piq.withPost; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PostImage)
		for i := range nodes {
			if nodes[i].post_images == nil {
				continue
			}
			fk := *nodes[i].post_images
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
				return nil, fmt.Errorf(`unexpected foreign-key "post_images" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Post = n
			}
		}
	}

	return nodes, nil
}

func (piq *PostImageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *PostImageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := piq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (piq *PostImageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postimage.Table,
			Columns: postimage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postimage.FieldID,
			},
		},
		From:   piq.sql,
		Unique: true,
	}
	if unique := piq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := piq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postimage.FieldID)
		for i := range fields {
			if fields[i] != postimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *PostImageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(postimage.Table)
	selector := builder.Select(t1.Columns(postimage.Columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(postimage.Columns...)...)
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PostImageGroupBy is the group-by builder for PostImage entities.
type PostImageGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *PostImageGroupBy) Aggregate(fns ...AggregateFunc) *PostImageGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the group-by query and scans the result into the given value.
func (pigb *PostImageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pigb.path(ctx)
	if err != nil {
		return err
	}
	pigb.sql = query
	return pigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pigb *PostImageGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pigb *PostImageGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pigb.fields) > 1 {
		return nil, errors.New("ent: PostImageGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pigb *PostImageGroupBy) StringsX(ctx context.Context) []string {
	v, err := pigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pigb *PostImageGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = fmt.Errorf("ent: PostImageGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pigb *PostImageGroupBy) StringX(ctx context.Context) string {
	v, err := pigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pigb *PostImageGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pigb.fields) > 1 {
		return nil, errors.New("ent: PostImageGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pigb *PostImageGroupBy) IntsX(ctx context.Context) []int {
	v, err := pigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pigb *PostImageGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = fmt.Errorf("ent: PostImageGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pigb *PostImageGroupBy) IntX(ctx context.Context) int {
	v, err := pigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pigb *PostImageGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pigb.fields) > 1 {
		return nil, errors.New("ent: PostImageGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pigb *PostImageGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pigb *PostImageGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = fmt.Errorf("ent: PostImageGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pigb *PostImageGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pigb *PostImageGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pigb.fields) > 1 {
		return nil, errors.New("ent: PostImageGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pigb *PostImageGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pigb *PostImageGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = fmt.Errorf("ent: PostImageGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pigb *PostImageGroupBy) BoolX(ctx context.Context) bool {
	v, err := pigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pigb *PostImageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pigb.fields {
		if !postimage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pigb *PostImageGroupBy) sqlQuery() *sql.Selector {
	selector := pigb.sql
	columns := make([]string, 0, len(pigb.fields)+len(pigb.fns))
	columns = append(columns, pigb.fields...)
	for _, fn := range pigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pigb.fields...)
}

// PostImageSelect is the builder for selecting fields of PostImage entities.
type PostImageSelect struct {
	*PostImageQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pis *PostImageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	pis.sql = pis.PostImageQuery.sqlQuery(ctx)
	return pis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pis *PostImageSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pis *PostImageSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pis.fields) > 1 {
		return nil, errors.New("ent: PostImageSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pis *PostImageSelect) StringsX(ctx context.Context) []string {
	v, err := pis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pis *PostImageSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = fmt.Errorf("ent: PostImageSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pis *PostImageSelect) StringX(ctx context.Context) string {
	v, err := pis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pis *PostImageSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pis.fields) > 1 {
		return nil, errors.New("ent: PostImageSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pis *PostImageSelect) IntsX(ctx context.Context) []int {
	v, err := pis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pis *PostImageSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = fmt.Errorf("ent: PostImageSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pis *PostImageSelect) IntX(ctx context.Context) int {
	v, err := pis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pis *PostImageSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pis.fields) > 1 {
		return nil, errors.New("ent: PostImageSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pis *PostImageSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pis *PostImageSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = fmt.Errorf("ent: PostImageSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pis *PostImageSelect) Float64X(ctx context.Context) float64 {
	v, err := pis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pis *PostImageSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pis.fields) > 1 {
		return nil, errors.New("ent: PostImageSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pis *PostImageSelect) BoolsX(ctx context.Context) []bool {
	v, err := pis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pis *PostImageSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postimage.Label}
	default:
		err = fmt.Errorf("ent: PostImageSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pis *PostImageSelect) BoolX(ctx context.Context) bool {
	v, err := pis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pis *PostImageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pis.sqlQuery().Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pis *PostImageSelect) sqlQuery() sql.Querier {
	selector := pis.sql
	selector.Select(selector.Columns(pis.fields...)...)
	return selector
}
