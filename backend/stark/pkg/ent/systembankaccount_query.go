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
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systembankaccount"
)

// SystemBankAccountQuery is the builder for querying SystemBankAccount entities.
type SystemBankAccountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SystemBankAccount
	modifiers  []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SystemBankAccountQuery builder.
func (sbaq *SystemBankAccountQuery) Where(ps ...predicate.SystemBankAccount) *SystemBankAccountQuery {
	sbaq.predicates = append(sbaq.predicates, ps...)
	return sbaq
}

// Limit adds a limit step to the query.
func (sbaq *SystemBankAccountQuery) Limit(limit int) *SystemBankAccountQuery {
	sbaq.limit = &limit
	return sbaq
}

// Offset adds an offset step to the query.
func (sbaq *SystemBankAccountQuery) Offset(offset int) *SystemBankAccountQuery {
	sbaq.offset = &offset
	return sbaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sbaq *SystemBankAccountQuery) Unique(unique bool) *SystemBankAccountQuery {
	sbaq.unique = &unique
	return sbaq
}

// Order adds an order step to the query.
func (sbaq *SystemBankAccountQuery) Order(o ...OrderFunc) *SystemBankAccountQuery {
	sbaq.order = append(sbaq.order, o...)
	return sbaq
}

// First returns the first SystemBankAccount entity from the query.
// Returns a *NotFoundError when no SystemBankAccount was found.
func (sbaq *SystemBankAccountQuery) First(ctx context.Context) (*SystemBankAccount, error) {
	nodes, err := sbaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{systembankaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sbaq *SystemBankAccountQuery) FirstX(ctx context.Context) *SystemBankAccount {
	node, err := sbaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SystemBankAccount ID from the query.
// Returns a *NotFoundError when no SystemBankAccount ID was found.
func (sbaq *SystemBankAccountQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sbaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{systembankaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sbaq *SystemBankAccountQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sbaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SystemBankAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SystemBankAccount entity is not found.
// Returns a *NotFoundError when no SystemBankAccount entities are found.
func (sbaq *SystemBankAccountQuery) Only(ctx context.Context) (*SystemBankAccount, error) {
	nodes, err := sbaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{systembankaccount.Label}
	default:
		return nil, &NotSingularError{systembankaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sbaq *SystemBankAccountQuery) OnlyX(ctx context.Context) *SystemBankAccount {
	node, err := sbaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SystemBankAccount ID in the query.
// Returns a *NotSingularError when exactly one SystemBankAccount ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sbaq *SystemBankAccountQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sbaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = &NotSingularError{systembankaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sbaq *SystemBankAccountQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sbaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SystemBankAccounts.
func (sbaq *SystemBankAccountQuery) All(ctx context.Context) ([]*SystemBankAccount, error) {
	if err := sbaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sbaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sbaq *SystemBankAccountQuery) AllX(ctx context.Context) []*SystemBankAccount {
	nodes, err := sbaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SystemBankAccount IDs.
func (sbaq *SystemBankAccountQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := sbaq.Select(systembankaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sbaq *SystemBankAccountQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sbaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sbaq *SystemBankAccountQuery) Count(ctx context.Context) (int, error) {
	if err := sbaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sbaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sbaq *SystemBankAccountQuery) CountX(ctx context.Context) int {
	count, err := sbaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sbaq *SystemBankAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := sbaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sbaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sbaq *SystemBankAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := sbaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SystemBankAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sbaq *SystemBankAccountQuery) Clone() *SystemBankAccountQuery {
	if sbaq == nil {
		return nil
	}
	return &SystemBankAccountQuery{
		config:     sbaq.config,
		limit:      sbaq.limit,
		offset:     sbaq.offset,
		order:      append([]OrderFunc{}, sbaq.order...),
		predicates: append([]predicate.SystemBankAccount{}, sbaq.predicates...),
		// clone intermediate query.
		sql:  sbaq.sql.Clone(),
		path: sbaq.path,
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
//	client.SystemBankAccount.Query().
//		GroupBy(systembankaccount.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sbaq *SystemBankAccountQuery) GroupBy(field string, fields ...string) *SystemBankAccountGroupBy {
	group := &SystemBankAccountGroupBy{config: sbaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sbaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sbaq.sqlQuery(ctx), nil
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
//	client.SystemBankAccount.Query().
//		Select(systembankaccount.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (sbaq *SystemBankAccountQuery) Select(fields ...string) *SystemBankAccountSelect {
	sbaq.fields = append(sbaq.fields, fields...)
	return &SystemBankAccountSelect{SystemBankAccountQuery: sbaq}
}

func (sbaq *SystemBankAccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sbaq.fields {
		if !systembankaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sbaq.path != nil {
		prev, err := sbaq.path(ctx)
		if err != nil {
			return err
		}
		sbaq.sql = prev
	}
	return nil
}

func (sbaq *SystemBankAccountQuery) sqlAll(ctx context.Context) ([]*SystemBankAccount, error) {
	var (
		nodes = []*SystemBankAccount{}
		_spec = sbaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SystemBankAccount{config: sbaq.config}
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
	if len(sbaq.modifiers) > 0 {
		_spec.Modifiers = sbaq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, sbaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sbaq *SystemBankAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sbaq.querySpec()
	if len(sbaq.modifiers) > 0 {
		_spec.Modifiers = sbaq.modifiers
	}
	return sqlgraph.CountNodes(ctx, sbaq.driver, _spec)
}

func (sbaq *SystemBankAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sbaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sbaq *SystemBankAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systembankaccount.Table,
			Columns: systembankaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: systembankaccount.FieldID,
			},
		},
		From:   sbaq.sql,
		Unique: true,
	}
	if unique := sbaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sbaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systembankaccount.FieldID)
		for i := range fields {
			if fields[i] != systembankaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sbaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sbaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sbaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sbaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sbaq *SystemBankAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sbaq.driver.Dialect())
	t1 := builder.Table(systembankaccount.Table)
	columns := sbaq.fields
	if len(columns) == 0 {
		columns = systembankaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sbaq.sql != nil {
		selector = sbaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sbaq.unique != nil && *sbaq.unique {
		selector.Distinct()
	}
	for _, m := range sbaq.modifiers {
		m(selector)
	}
	for _, p := range sbaq.predicates {
		p(selector)
	}
	for _, p := range sbaq.order {
		p(selector)
	}
	if offset := sbaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sbaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (sbaq *SystemBankAccountQuery) ForUpdate(opts ...sql.LockOption) *SystemBankAccountQuery {
	if sbaq.driver.Dialect() == dialect.Postgres {
		sbaq.Unique(false)
	}
	sbaq.modifiers = append(sbaq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return sbaq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (sbaq *SystemBankAccountQuery) ForShare(opts ...sql.LockOption) *SystemBankAccountQuery {
	if sbaq.driver.Dialect() == dialect.Postgres {
		sbaq.Unique(false)
	}
	sbaq.modifiers = append(sbaq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sbaq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sbaq *SystemBankAccountQuery) Modify(modifiers ...func(s *sql.Selector)) *SystemBankAccountSelect {
	sbaq.modifiers = append(sbaq.modifiers, modifiers...)
	return sbaq.Select()
}

// SystemBankAccountGroupBy is the group-by builder for SystemBankAccount entities.
type SystemBankAccountGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sbagb *SystemBankAccountGroupBy) Aggregate(fns ...AggregateFunc) *SystemBankAccountGroupBy {
	sbagb.fns = append(sbagb.fns, fns...)
	return sbagb
}

// Scan applies the group-by query and scans the result into the given value.
func (sbagb *SystemBankAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sbagb.path(ctx)
	if err != nil {
		return err
	}
	sbagb.sql = query
	return sbagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sbagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sbagb *SystemBankAccountGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sbagb.fields) > 1 {
		return nil, errors.New("ent: SystemBankAccountGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sbagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) StringsX(ctx context.Context) []string {
	v, err := sbagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sbagb *SystemBankAccountGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sbagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = fmt.Errorf("ent: SystemBankAccountGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) StringX(ctx context.Context) string {
	v, err := sbagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sbagb *SystemBankAccountGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sbagb.fields) > 1 {
		return nil, errors.New("ent: SystemBankAccountGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sbagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) IntsX(ctx context.Context) []int {
	v, err := sbagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sbagb *SystemBankAccountGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sbagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = fmt.Errorf("ent: SystemBankAccountGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) IntX(ctx context.Context) int {
	v, err := sbagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sbagb *SystemBankAccountGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sbagb.fields) > 1 {
		return nil, errors.New("ent: SystemBankAccountGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sbagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sbagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sbagb *SystemBankAccountGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sbagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = fmt.Errorf("ent: SystemBankAccountGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sbagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sbagb *SystemBankAccountGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sbagb.fields) > 1 {
		return nil, errors.New("ent: SystemBankAccountGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sbagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sbagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sbagb *SystemBankAccountGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sbagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = fmt.Errorf("ent: SystemBankAccountGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sbagb *SystemBankAccountGroupBy) BoolX(ctx context.Context) bool {
	v, err := sbagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sbagb *SystemBankAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sbagb.fields {
		if !systembankaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sbagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sbagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sbagb *SystemBankAccountGroupBy) sqlQuery() *sql.Selector {
	selector := sbagb.sql.Select()
	aggregation := make([]string, 0, len(sbagb.fns))
	for _, fn := range sbagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sbagb.fields)+len(sbagb.fns))
		for _, f := range sbagb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sbagb.fields...)...)
}

// SystemBankAccountSelect is the builder for selecting fields of SystemBankAccount entities.
type SystemBankAccountSelect struct {
	*SystemBankAccountQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sbas *SystemBankAccountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sbas.prepareQuery(ctx); err != nil {
		return err
	}
	sbas.sql = sbas.SystemBankAccountQuery.sqlQuery(ctx)
	return sbas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sbas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sbas *SystemBankAccountSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sbas.fields) > 1 {
		return nil, errors.New("ent: SystemBankAccountSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sbas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) StringsX(ctx context.Context) []string {
	v, err := sbas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sbas *SystemBankAccountSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sbas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = fmt.Errorf("ent: SystemBankAccountSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) StringX(ctx context.Context) string {
	v, err := sbas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sbas *SystemBankAccountSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sbas.fields) > 1 {
		return nil, errors.New("ent: SystemBankAccountSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sbas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) IntsX(ctx context.Context) []int {
	v, err := sbas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sbas *SystemBankAccountSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sbas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = fmt.Errorf("ent: SystemBankAccountSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) IntX(ctx context.Context) int {
	v, err := sbas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sbas *SystemBankAccountSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sbas.fields) > 1 {
		return nil, errors.New("ent: SystemBankAccountSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sbas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sbas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sbas *SystemBankAccountSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sbas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = fmt.Errorf("ent: SystemBankAccountSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) Float64X(ctx context.Context) float64 {
	v, err := sbas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sbas *SystemBankAccountSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sbas.fields) > 1 {
		return nil, errors.New("ent: SystemBankAccountSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sbas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) BoolsX(ctx context.Context) []bool {
	v, err := sbas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sbas *SystemBankAccountSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sbas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systembankaccount.Label}
	default:
		err = fmt.Errorf("ent: SystemBankAccountSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sbas *SystemBankAccountSelect) BoolX(ctx context.Context) bool {
	v, err := sbas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sbas *SystemBankAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sbas.sql.Query()
	if err := sbas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sbas *SystemBankAccountSelect) Modify(modifiers ...func(s *sql.Selector)) *SystemBankAccountSelect {
	sbas.modifiers = append(sbas.modifiers, modifiers...)
	return sbas
}