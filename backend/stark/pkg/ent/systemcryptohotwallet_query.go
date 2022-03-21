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
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemcryptohotwallet"
)

// SystemCryptoHotWalletQuery is the builder for querying SystemCryptoHotWallet entities.
type SystemCryptoHotWalletQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SystemCryptoHotWallet
	modifiers  []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SystemCryptoHotWalletQuery builder.
func (schwq *SystemCryptoHotWalletQuery) Where(ps ...predicate.SystemCryptoHotWallet) *SystemCryptoHotWalletQuery {
	schwq.predicates = append(schwq.predicates, ps...)
	return schwq
}

// Limit adds a limit step to the query.
func (schwq *SystemCryptoHotWalletQuery) Limit(limit int) *SystemCryptoHotWalletQuery {
	schwq.limit = &limit
	return schwq
}

// Offset adds an offset step to the query.
func (schwq *SystemCryptoHotWalletQuery) Offset(offset int) *SystemCryptoHotWalletQuery {
	schwq.offset = &offset
	return schwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (schwq *SystemCryptoHotWalletQuery) Unique(unique bool) *SystemCryptoHotWalletQuery {
	schwq.unique = &unique
	return schwq
}

// Order adds an order step to the query.
func (schwq *SystemCryptoHotWalletQuery) Order(o ...OrderFunc) *SystemCryptoHotWalletQuery {
	schwq.order = append(schwq.order, o...)
	return schwq
}

// First returns the first SystemCryptoHotWallet entity from the query.
// Returns a *NotFoundError when no SystemCryptoHotWallet was found.
func (schwq *SystemCryptoHotWalletQuery) First(ctx context.Context) (*SystemCryptoHotWallet, error) {
	nodes, err := schwq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{systemcryptohotwallet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (schwq *SystemCryptoHotWalletQuery) FirstX(ctx context.Context) *SystemCryptoHotWallet {
	node, err := schwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SystemCryptoHotWallet ID from the query.
// Returns a *NotFoundError when no SystemCryptoHotWallet ID was found.
func (schwq *SystemCryptoHotWalletQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = schwq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{systemcryptohotwallet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (schwq *SystemCryptoHotWalletQuery) FirstIDX(ctx context.Context) int64 {
	id, err := schwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SystemCryptoHotWallet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SystemCryptoHotWallet entity is not found.
// Returns a *NotFoundError when no SystemCryptoHotWallet entities are found.
func (schwq *SystemCryptoHotWalletQuery) Only(ctx context.Context) (*SystemCryptoHotWallet, error) {
	nodes, err := schwq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{systemcryptohotwallet.Label}
	default:
		return nil, &NotSingularError{systemcryptohotwallet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (schwq *SystemCryptoHotWalletQuery) OnlyX(ctx context.Context) *SystemCryptoHotWallet {
	node, err := schwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SystemCryptoHotWallet ID in the query.
// Returns a *NotSingularError when exactly one SystemCryptoHotWallet ID is not found.
// Returns a *NotFoundError when no entities are found.
func (schwq *SystemCryptoHotWalletQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = schwq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = &NotSingularError{systemcryptohotwallet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (schwq *SystemCryptoHotWalletQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := schwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SystemCryptoHotWallets.
func (schwq *SystemCryptoHotWalletQuery) All(ctx context.Context) ([]*SystemCryptoHotWallet, error) {
	if err := schwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return schwq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (schwq *SystemCryptoHotWalletQuery) AllX(ctx context.Context) []*SystemCryptoHotWallet {
	nodes, err := schwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SystemCryptoHotWallet IDs.
func (schwq *SystemCryptoHotWalletQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := schwq.Select(systemcryptohotwallet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (schwq *SystemCryptoHotWalletQuery) IDsX(ctx context.Context) []int64 {
	ids, err := schwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (schwq *SystemCryptoHotWalletQuery) Count(ctx context.Context) (int, error) {
	if err := schwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return schwq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (schwq *SystemCryptoHotWalletQuery) CountX(ctx context.Context) int {
	count, err := schwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (schwq *SystemCryptoHotWalletQuery) Exist(ctx context.Context) (bool, error) {
	if err := schwq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return schwq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (schwq *SystemCryptoHotWalletQuery) ExistX(ctx context.Context) bool {
	exist, err := schwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SystemCryptoHotWalletQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (schwq *SystemCryptoHotWalletQuery) Clone() *SystemCryptoHotWalletQuery {
	if schwq == nil {
		return nil
	}
	return &SystemCryptoHotWalletQuery{
		config:     schwq.config,
		limit:      schwq.limit,
		offset:     schwq.offset,
		order:      append([]OrderFunc{}, schwq.order...),
		predicates: append([]predicate.SystemCryptoHotWallet{}, schwq.predicates...),
		// clone intermediate query.
		sql:  schwq.sql.Clone(),
		path: schwq.path,
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
//	client.SystemCryptoHotWallet.Query().
//		GroupBy(systemcryptohotwallet.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (schwq *SystemCryptoHotWalletQuery) GroupBy(field string, fields ...string) *SystemCryptoHotWalletGroupBy {
	group := &SystemCryptoHotWalletGroupBy{config: schwq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := schwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return schwq.sqlQuery(ctx), nil
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
//	client.SystemCryptoHotWallet.Query().
//		Select(systemcryptohotwallet.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (schwq *SystemCryptoHotWalletQuery) Select(fields ...string) *SystemCryptoHotWalletSelect {
	schwq.fields = append(schwq.fields, fields...)
	return &SystemCryptoHotWalletSelect{SystemCryptoHotWalletQuery: schwq}
}

func (schwq *SystemCryptoHotWalletQuery) prepareQuery(ctx context.Context) error {
	for _, f := range schwq.fields {
		if !systemcryptohotwallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if schwq.path != nil {
		prev, err := schwq.path(ctx)
		if err != nil {
			return err
		}
		schwq.sql = prev
	}
	return nil
}

func (schwq *SystemCryptoHotWalletQuery) sqlAll(ctx context.Context) ([]*SystemCryptoHotWallet, error) {
	var (
		nodes = []*SystemCryptoHotWallet{}
		_spec = schwq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SystemCryptoHotWallet{config: schwq.config}
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
	if len(schwq.modifiers) > 0 {
		_spec.Modifiers = schwq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, schwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (schwq *SystemCryptoHotWalletQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := schwq.querySpec()
	if len(schwq.modifiers) > 0 {
		_spec.Modifiers = schwq.modifiers
	}
	return sqlgraph.CountNodes(ctx, schwq.driver, _spec)
}

func (schwq *SystemCryptoHotWalletQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := schwq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (schwq *SystemCryptoHotWalletQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systemcryptohotwallet.Table,
			Columns: systemcryptohotwallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: systemcryptohotwallet.FieldID,
			},
		},
		From:   schwq.sql,
		Unique: true,
	}
	if unique := schwq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := schwq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemcryptohotwallet.FieldID)
		for i := range fields {
			if fields[i] != systemcryptohotwallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := schwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := schwq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := schwq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := schwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (schwq *SystemCryptoHotWalletQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(schwq.driver.Dialect())
	t1 := builder.Table(systemcryptohotwallet.Table)
	columns := schwq.fields
	if len(columns) == 0 {
		columns = systemcryptohotwallet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if schwq.sql != nil {
		selector = schwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if schwq.unique != nil && *schwq.unique {
		selector.Distinct()
	}
	for _, m := range schwq.modifiers {
		m(selector)
	}
	for _, p := range schwq.predicates {
		p(selector)
	}
	for _, p := range schwq.order {
		p(selector)
	}
	if offset := schwq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := schwq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (schwq *SystemCryptoHotWalletQuery) ForUpdate(opts ...sql.LockOption) *SystemCryptoHotWalletQuery {
	if schwq.driver.Dialect() == dialect.Postgres {
		schwq.Unique(false)
	}
	schwq.modifiers = append(schwq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return schwq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (schwq *SystemCryptoHotWalletQuery) ForShare(opts ...sql.LockOption) *SystemCryptoHotWalletQuery {
	if schwq.driver.Dialect() == dialect.Postgres {
		schwq.Unique(false)
	}
	schwq.modifiers = append(schwq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return schwq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (schwq *SystemCryptoHotWalletQuery) Modify(modifiers ...func(s *sql.Selector)) *SystemCryptoHotWalletSelect {
	schwq.modifiers = append(schwq.modifiers, modifiers...)
	return schwq.Select()
}

// SystemCryptoHotWalletGroupBy is the group-by builder for SystemCryptoHotWallet entities.
type SystemCryptoHotWalletGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (schwgb *SystemCryptoHotWalletGroupBy) Aggregate(fns ...AggregateFunc) *SystemCryptoHotWalletGroupBy {
	schwgb.fns = append(schwgb.fns, fns...)
	return schwgb
}

// Scan applies the group-by query and scans the result into the given value.
func (schwgb *SystemCryptoHotWalletGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := schwgb.path(ctx)
	if err != nil {
		return err
	}
	schwgb.sql = query
	return schwgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := schwgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (schwgb *SystemCryptoHotWalletGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(schwgb.fields) > 1 {
		return nil, errors.New("ent: SystemCryptoHotWalletGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := schwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) StringsX(ctx context.Context) []string {
	v, err := schwgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (schwgb *SystemCryptoHotWalletGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = schwgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = fmt.Errorf("ent: SystemCryptoHotWalletGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) StringX(ctx context.Context) string {
	v, err := schwgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (schwgb *SystemCryptoHotWalletGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(schwgb.fields) > 1 {
		return nil, errors.New("ent: SystemCryptoHotWalletGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := schwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) IntsX(ctx context.Context) []int {
	v, err := schwgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (schwgb *SystemCryptoHotWalletGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = schwgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = fmt.Errorf("ent: SystemCryptoHotWalletGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) IntX(ctx context.Context) int {
	v, err := schwgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (schwgb *SystemCryptoHotWalletGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(schwgb.fields) > 1 {
		return nil, errors.New("ent: SystemCryptoHotWalletGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := schwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := schwgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (schwgb *SystemCryptoHotWalletGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = schwgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = fmt.Errorf("ent: SystemCryptoHotWalletGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) Float64X(ctx context.Context) float64 {
	v, err := schwgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (schwgb *SystemCryptoHotWalletGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(schwgb.fields) > 1 {
		return nil, errors.New("ent: SystemCryptoHotWalletGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := schwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := schwgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (schwgb *SystemCryptoHotWalletGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = schwgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = fmt.Errorf("ent: SystemCryptoHotWalletGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (schwgb *SystemCryptoHotWalletGroupBy) BoolX(ctx context.Context) bool {
	v, err := schwgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (schwgb *SystemCryptoHotWalletGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range schwgb.fields {
		if !systemcryptohotwallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := schwgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := schwgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (schwgb *SystemCryptoHotWalletGroupBy) sqlQuery() *sql.Selector {
	selector := schwgb.sql.Select()
	aggregation := make([]string, 0, len(schwgb.fns))
	for _, fn := range schwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(schwgb.fields)+len(schwgb.fns))
		for _, f := range schwgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(schwgb.fields...)...)
}

// SystemCryptoHotWalletSelect is the builder for selecting fields of SystemCryptoHotWallet entities.
type SystemCryptoHotWalletSelect struct {
	*SystemCryptoHotWalletQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (schws *SystemCryptoHotWalletSelect) Scan(ctx context.Context, v interface{}) error {
	if err := schws.prepareQuery(ctx); err != nil {
		return err
	}
	schws.sql = schws.SystemCryptoHotWalletQuery.sqlQuery(ctx)
	return schws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) ScanX(ctx context.Context, v interface{}) {
	if err := schws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (schws *SystemCryptoHotWalletSelect) Strings(ctx context.Context) ([]string, error) {
	if len(schws.fields) > 1 {
		return nil, errors.New("ent: SystemCryptoHotWalletSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := schws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) StringsX(ctx context.Context) []string {
	v, err := schws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (schws *SystemCryptoHotWalletSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = schws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = fmt.Errorf("ent: SystemCryptoHotWalletSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) StringX(ctx context.Context) string {
	v, err := schws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (schws *SystemCryptoHotWalletSelect) Ints(ctx context.Context) ([]int, error) {
	if len(schws.fields) > 1 {
		return nil, errors.New("ent: SystemCryptoHotWalletSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := schws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) IntsX(ctx context.Context) []int {
	v, err := schws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (schws *SystemCryptoHotWalletSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = schws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = fmt.Errorf("ent: SystemCryptoHotWalletSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) IntX(ctx context.Context) int {
	v, err := schws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (schws *SystemCryptoHotWalletSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(schws.fields) > 1 {
		return nil, errors.New("ent: SystemCryptoHotWalletSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := schws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) Float64sX(ctx context.Context) []float64 {
	v, err := schws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (schws *SystemCryptoHotWalletSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = schws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = fmt.Errorf("ent: SystemCryptoHotWalletSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) Float64X(ctx context.Context) float64 {
	v, err := schws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (schws *SystemCryptoHotWalletSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(schws.fields) > 1 {
		return nil, errors.New("ent: SystemCryptoHotWalletSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := schws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) BoolsX(ctx context.Context) []bool {
	v, err := schws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (schws *SystemCryptoHotWalletSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = schws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemcryptohotwallet.Label}
	default:
		err = fmt.Errorf("ent: SystemCryptoHotWalletSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (schws *SystemCryptoHotWalletSelect) BoolX(ctx context.Context) bool {
	v, err := schws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (schws *SystemCryptoHotWalletSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := schws.sql.Query()
	if err := schws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (schws *SystemCryptoHotWalletSelect) Modify(modifiers ...func(s *sql.Selector)) *SystemCryptoHotWalletSelect {
	schws.modifiers = append(schws.modifiers, modifiers...)
	return schws
}
