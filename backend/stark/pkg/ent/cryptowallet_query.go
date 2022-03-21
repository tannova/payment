// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/cryptowallet"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

// CryptoWalletQuery is the builder for querying CryptoWallet entities.
type CryptoWalletQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CryptoWallet
	modifiers  []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CryptoWalletQuery builder.
func (cwq *CryptoWalletQuery) Where(ps ...predicate.CryptoWallet) *CryptoWalletQuery {
	cwq.predicates = append(cwq.predicates, ps...)
	return cwq
}

// Limit adds a limit step to the query.
func (cwq *CryptoWalletQuery) Limit(limit int) *CryptoWalletQuery {
	cwq.limit = &limit
	return cwq
}

// Offset adds an offset step to the query.
func (cwq *CryptoWalletQuery) Offset(offset int) *CryptoWalletQuery {
	cwq.offset = &offset
	return cwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cwq *CryptoWalletQuery) Unique(unique bool) *CryptoWalletQuery {
	cwq.unique = &unique
	return cwq
}

// Order adds an order step to the query.
func (cwq *CryptoWalletQuery) Order(o ...OrderFunc) *CryptoWalletQuery {
	cwq.order = append(cwq.order, o...)
	return cwq
}

// First returns the first CryptoWallet entity from the query.
// Returns a *NotFoundError when no CryptoWallet was found.
func (cwq *CryptoWalletQuery) First(ctx context.Context) (*CryptoWallet, error) {
	nodes, err := cwq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cryptowallet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cwq *CryptoWalletQuery) FirstX(ctx context.Context) *CryptoWallet {
	node, err := cwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CryptoWallet ID from the query.
// Returns a *NotFoundError when no CryptoWallet ID was found.
func (cwq *CryptoWalletQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cwq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cryptowallet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cwq *CryptoWalletQuery) FirstIDX(ctx context.Context) int64 {
	id, err := cwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CryptoWallet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CryptoWallet entity is not found.
// Returns a *NotFoundError when no CryptoWallet entities are found.
func (cwq *CryptoWalletQuery) Only(ctx context.Context) (*CryptoWallet, error) {
	nodes, err := cwq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cryptowallet.Label}
	default:
		return nil, &NotSingularError{cryptowallet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cwq *CryptoWalletQuery) OnlyX(ctx context.Context) *CryptoWallet {
	node, err := cwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CryptoWallet ID in the query.
// Returns a *NotSingularError when exactly one CryptoWallet ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cwq *CryptoWalletQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cwq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = &NotSingularError{cryptowallet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cwq *CryptoWalletQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := cwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CryptoWallets.
func (cwq *CryptoWalletQuery) All(ctx context.Context) ([]*CryptoWallet, error) {
	if err := cwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cwq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cwq *CryptoWalletQuery) AllX(ctx context.Context) []*CryptoWallet {
	nodes, err := cwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CryptoWallet IDs.
func (cwq *CryptoWalletQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := cwq.Select(cryptowallet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cwq *CryptoWalletQuery) IDsX(ctx context.Context) []int64 {
	ids, err := cwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cwq *CryptoWalletQuery) Count(ctx context.Context) (int, error) {
	if err := cwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cwq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cwq *CryptoWalletQuery) CountX(ctx context.Context) int {
	count, err := cwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cwq *CryptoWalletQuery) Exist(ctx context.Context) (bool, error) {
	if err := cwq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cwq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cwq *CryptoWalletQuery) ExistX(ctx context.Context) bool {
	exist, err := cwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CryptoWalletQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cwq *CryptoWalletQuery) Clone() *CryptoWalletQuery {
	if cwq == nil {
		return nil
	}
	return &CryptoWalletQuery{
		config:     cwq.config,
		limit:      cwq.limit,
		offset:     cwq.offset,
		order:      append([]OrderFunc{}, cwq.order...),
		predicates: append([]predicate.CryptoWallet{}, cwq.predicates...),
		// clone intermediate query.
		sql:  cwq.sql.Clone(),
		path: cwq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CryptoWallet.Query().
//		GroupBy(cryptowallet.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cwq *CryptoWalletQuery) GroupBy(field string, fields ...string) *CryptoWalletGroupBy {
	group := &CryptoWalletGroupBy{config: cwq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cwq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.CryptoWallet.Query().
//		Select(cryptowallet.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (cwq *CryptoWalletQuery) Select(fields ...string) *CryptoWalletSelect {
	cwq.fields = append(cwq.fields, fields...)
	return &CryptoWalletSelect{CryptoWalletQuery: cwq}
}

func (cwq *CryptoWalletQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cwq.fields {
		if !cryptowallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cwq.path != nil {
		prev, err := cwq.path(ctx)
		if err != nil {
			return err
		}
		cwq.sql = prev
	}
	return nil
}

func (cwq *CryptoWalletQuery) sqlAll(ctx context.Context) ([]*CryptoWallet, error) {
	var (
		nodes = []*CryptoWallet{}
		_spec = cwq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CryptoWallet{config: cwq.config}
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
	if len(cwq.modifiers) > 0 {
		_spec.Modifiers = cwq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, cwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cwq *CryptoWalletQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cwq.querySpec()
	if len(cwq.modifiers) > 0 {
		_spec.Modifiers = cwq.modifiers
	}
	return sqlgraph.CountNodes(ctx, cwq.driver, _spec)
}

func (cwq *CryptoWalletQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cwq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cwq *CryptoWalletQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cryptowallet.Table,
			Columns: cryptowallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: cryptowallet.FieldID,
			},
		},
		From:   cwq.sql,
		Unique: true,
	}
	if unique := cwq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cwq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cryptowallet.FieldID)
		for i := range fields {
			if fields[i] != cryptowallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cwq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cwq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cwq *CryptoWalletQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cwq.driver.Dialect())
	t1 := builder.Table(cryptowallet.Table)
	columns := cwq.fields
	if len(columns) == 0 {
		columns = cryptowallet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cwq.sql != nil {
		selector = cwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cwq.unique != nil && *cwq.unique {
		selector.Distinct()
	}
	for _, m := range cwq.modifiers {
		m(selector)
	}
	for _, p := range cwq.predicates {
		p(selector)
	}
	for _, p := range cwq.order {
		p(selector)
	}
	if offset := cwq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cwq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cwq *CryptoWalletQuery) ForUpdate(opts ...sql.LockOption) *CryptoWalletQuery {
	if cwq.driver.Dialect() == dialect.Postgres {
		cwq.Unique(false)
	}
	cwq.modifiers = append(cwq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cwq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cwq *CryptoWalletQuery) ForShare(opts ...sql.LockOption) *CryptoWalletQuery {
	if cwq.driver.Dialect() == dialect.Postgres {
		cwq.Unique(false)
	}
	cwq.modifiers = append(cwq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cwq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cwq *CryptoWalletQuery) Modify(modifiers ...func(s *sql.Selector)) *CryptoWalletSelect {
	cwq.modifiers = append(cwq.modifiers, modifiers...)
	return cwq.Select()
}

// CryptoWalletGroupBy is the group-by builder for CryptoWallet entities.
type CryptoWalletGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cwgb *CryptoWalletGroupBy) Aggregate(fns ...AggregateFunc) *CryptoWalletGroupBy {
	cwgb.fns = append(cwgb.fns, fns...)
	return cwgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cwgb *CryptoWalletGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cwgb.path(ctx)
	if err != nil {
		return err
	}
	cwgb.sql = query
	return cwgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cwgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cwgb *CryptoWalletGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cwgb.fields) > 1 {
		return nil, errors.New("ent: CryptoWalletGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) StringsX(ctx context.Context) []string {
	v, err := cwgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cwgb *CryptoWalletGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cwgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = fmt.Errorf("ent: CryptoWalletGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) StringX(ctx context.Context) string {
	v, err := cwgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cwgb *CryptoWalletGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cwgb.fields) > 1 {
		return nil, errors.New("ent: CryptoWalletGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) IntsX(ctx context.Context) []int {
	v, err := cwgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cwgb *CryptoWalletGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cwgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = fmt.Errorf("ent: CryptoWalletGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) IntX(ctx context.Context) int {
	v, err := cwgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cwgb *CryptoWalletGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cwgb.fields) > 1 {
		return nil, errors.New("ent: CryptoWalletGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cwgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cwgb *CryptoWalletGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cwgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = fmt.Errorf("ent: CryptoWalletGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cwgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cwgb *CryptoWalletGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cwgb.fields) > 1 {
		return nil, errors.New("ent: CryptoWalletGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cwgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cwgb *CryptoWalletGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cwgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = fmt.Errorf("ent: CryptoWalletGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cwgb *CryptoWalletGroupBy) BoolX(ctx context.Context) bool {
	v, err := cwgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cwgb *CryptoWalletGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cwgb.fields {
		if !cryptowallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cwgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cwgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cwgb *CryptoWalletGroupBy) sqlQuery() *sql.Selector {
	selector := cwgb.sql.Select()
	aggregation := make([]string, 0, len(cwgb.fns))
	for _, fn := range cwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cwgb.fields)+len(cwgb.fns))
		for _, f := range cwgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cwgb.fields...)...)
}

// CryptoWalletSelect is the builder for selecting fields of CryptoWallet entities.
type CryptoWalletSelect struct {
	*CryptoWalletQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cws *CryptoWalletSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cws.prepareQuery(ctx); err != nil {
		return err
	}
	cws.sql = cws.CryptoWalletQuery.sqlQuery(ctx)
	return cws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cws *CryptoWalletSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cws *CryptoWalletSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cws.fields) > 1 {
		return nil, errors.New("ent: CryptoWalletSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cws *CryptoWalletSelect) StringsX(ctx context.Context) []string {
	v, err := cws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cws *CryptoWalletSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = fmt.Errorf("ent: CryptoWalletSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cws *CryptoWalletSelect) StringX(ctx context.Context) string {
	v, err := cws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cws *CryptoWalletSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cws.fields) > 1 {
		return nil, errors.New("ent: CryptoWalletSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cws *CryptoWalletSelect) IntsX(ctx context.Context) []int {
	v, err := cws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cws *CryptoWalletSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = fmt.Errorf("ent: CryptoWalletSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cws *CryptoWalletSelect) IntX(ctx context.Context) int {
	v, err := cws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cws *CryptoWalletSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cws.fields) > 1 {
		return nil, errors.New("ent: CryptoWalletSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cws *CryptoWalletSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cws *CryptoWalletSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = fmt.Errorf("ent: CryptoWalletSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cws *CryptoWalletSelect) Float64X(ctx context.Context) float64 {
	v, err := cws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cws *CryptoWalletSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cws.fields) > 1 {
		return nil, errors.New("ent: CryptoWalletSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cws *CryptoWalletSelect) BoolsX(ctx context.Context) []bool {
	v, err := cws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cws *CryptoWalletSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cryptowallet.Label}
	default:
		err = fmt.Errorf("ent: CryptoWalletSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cws *CryptoWalletSelect) BoolX(ctx context.Context) bool {
	v, err := cws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cws *CryptoWalletSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cws.sql.Query()
	if err := cws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cws *CryptoWalletSelect) Modify(modifiers ...func(s *sql.Selector)) *CryptoWalletSelect {
	cws.modifiers = append(cws.modifiers, modifiers...)
	return cws
}
