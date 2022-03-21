// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/predicate"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/voucher"
)

// VoucherQuery is the builder for querying Voucher entities.
type VoucherQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Voucher
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoucherQuery builder.
func (vq *VoucherQuery) Where(ps ...predicate.Voucher) *VoucherQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *VoucherQuery) Limit(limit int) *VoucherQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *VoucherQuery) Offset(offset int) *VoucherQuery {
	vq.offset = &offset
	return vq
}

// Order adds an order step to the query.
func (vq *VoucherQuery) Order(o ...OrderFunc) *VoucherQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// First returns the first Voucher entity from the query.
// Returns a *NotFoundError when no Voucher was found.
func (vq *VoucherQuery) First(ctx context.Context) (*Voucher, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{voucher.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VoucherQuery) FirstX(ctx context.Context) *Voucher {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Voucher ID from the query.
// Returns a *NotFoundError when no Voucher ID was found.
func (vq *VoucherQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{voucher.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VoucherQuery) FirstIDX(ctx context.Context) int64 {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Voucher entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Voucher entity is not found.
// Returns a *NotFoundError when no Voucher entities are found.
func (vq *VoucherQuery) Only(ctx context.Context) (*Voucher, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{voucher.Label}
	default:
		return nil, &NotSingularError{voucher.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VoucherQuery) OnlyX(ctx context.Context) *Voucher {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Voucher ID in the query.
// Returns a *NotSingularError when exactly one Voucher ID is not found.
// Returns a *NotFoundError when no entities are found.
func (vq *VoucherQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = &NotSingularError{voucher.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VoucherQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vouchers.
func (vq *VoucherQuery) All(ctx context.Context) ([]*Voucher, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *VoucherQuery) AllX(ctx context.Context) []*Voucher {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Voucher IDs.
func (vq *VoucherQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := vq.Select(voucher.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VoucherQuery) IDsX(ctx context.Context) []int64 {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VoucherQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VoucherQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VoucherQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VoucherQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoucherQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VoucherQuery) Clone() *VoucherQuery {
	if vq == nil {
		return nil
	}
	return &VoucherQuery{
		config:     vq.config,
		limit:      vq.limit,
		offset:     vq.offset,
		order:      append([]OrderFunc{}, vq.order...),
		predicates: append([]predicate.Voucher{}, vq.predicates...),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MerchantID int64 `json:"merchant_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Voucher.Query().
//		GroupBy(voucher.FieldMerchantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vq *VoucherQuery) GroupBy(field string, fields ...string) *VoucherGroupBy {
	group := &VoucherGroupBy{config: vq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MerchantID int64 `json:"merchant_id,omitempty"`
//	}
//
//	client.Voucher.Query().
//		Select(voucher.FieldMerchantID).
//		Scan(ctx, &v)
//
func (vq *VoucherQuery) Select(field string, fields ...string) *VoucherSelect {
	vq.fields = append([]string{field}, fields...)
	return &VoucherSelect{VoucherQuery: vq}
}

func (vq *VoucherQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !voucher.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VoucherQuery) sqlAll(ctx context.Context) ([]*Voucher, error) {
	var (
		nodes = []*Voucher{}
		_spec = vq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Voucher{config: vq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vq *VoucherQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VoucherQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vq *VoucherQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   voucher.Table,
			Columns: voucher.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: voucher.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voucher.FieldID)
		for i := range fields {
			if fields[i] != voucher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, voucher.ValidColumn)
			}
		}
	}
	return _spec
}

func (vq *VoucherQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(voucher.Table)
	selector := builder.Select(t1.Columns(voucher.Columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(voucher.Columns...)...)
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector, voucher.ValidColumn)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VoucherGroupBy is the group-by builder for Voucher entities.
type VoucherGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VoucherGroupBy) Aggregate(fns ...AggregateFunc) *VoucherGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *VoucherGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vgb *VoucherGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoucherGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VoucherGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vgb *VoucherGroupBy) StringsX(ctx context.Context) []string {
	v, err := vgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoucherGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = fmt.Errorf("ent: VoucherGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vgb *VoucherGroupBy) StringX(ctx context.Context) string {
	v, err := vgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoucherGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VoucherGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vgb *VoucherGroupBy) IntsX(ctx context.Context) []int {
	v, err := vgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoucherGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = fmt.Errorf("ent: VoucherGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vgb *VoucherGroupBy) IntX(ctx context.Context) int {
	v, err := vgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoucherGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VoucherGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vgb *VoucherGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoucherGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = fmt.Errorf("ent: VoucherGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vgb *VoucherGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoucherGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: VoucherGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vgb *VoucherGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *VoucherGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = fmt.Errorf("ent: VoucherGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vgb *VoucherGroupBy) BoolX(ctx context.Context) bool {
	v, err := vgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vgb *VoucherGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vgb.fields {
		if !voucher.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vgb *VoucherGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql
	columns := make([]string, 0, len(vgb.fields)+len(vgb.fns))
	columns = append(columns, vgb.fields...)
	for _, fn := range vgb.fns {
		columns = append(columns, fn(selector, voucher.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(vgb.fields...)
}

// VoucherSelect is the builder for selecting fields of Voucher entities.
type VoucherSelect struct {
	*VoucherQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VoucherSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.VoucherQuery.sqlQuery(ctx)
	return vs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vs *VoucherSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vs *VoucherSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VoucherSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vs *VoucherSelect) StringsX(ctx context.Context) []string {
	v, err := vs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vs *VoucherSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = fmt.Errorf("ent: VoucherSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vs *VoucherSelect) StringX(ctx context.Context) string {
	v, err := vs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vs *VoucherSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VoucherSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vs *VoucherSelect) IntsX(ctx context.Context) []int {
	v, err := vs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vs *VoucherSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = fmt.Errorf("ent: VoucherSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vs *VoucherSelect) IntX(ctx context.Context) int {
	v, err := vs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vs *VoucherSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VoucherSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vs *VoucherSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vs *VoucherSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = fmt.Errorf("ent: VoucherSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vs *VoucherSelect) Float64X(ctx context.Context) float64 {
	v, err := vs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vs *VoucherSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: VoucherSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vs *VoucherSelect) BoolsX(ctx context.Context) []bool {
	v, err := vs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vs *VoucherSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = fmt.Errorf("ent: VoucherSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vs *VoucherSelect) BoolX(ctx context.Context) bool {
	v, err := vs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vs *VoucherSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vs.sqlQuery().Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vs *VoucherSelect) sqlQuery() sql.Querier {
	selector := vs.sql
	selector.Select(selector.Columns(vs.fields...)...)
	return selector
}
