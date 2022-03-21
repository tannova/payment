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
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

// PaymentTelcoDetailQuery is the builder for querying PaymentTelcoDetail entities.
type PaymentTelcoDetailQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PaymentTelcoDetail
	// eager-loading edges.
	withPayment *PaymentQuery
	withFKs     bool
	modifiers   []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaymentTelcoDetailQuery builder.
func (ptdq *PaymentTelcoDetailQuery) Where(ps ...predicate.PaymentTelcoDetail) *PaymentTelcoDetailQuery {
	ptdq.predicates = append(ptdq.predicates, ps...)
	return ptdq
}

// Limit adds a limit step to the query.
func (ptdq *PaymentTelcoDetailQuery) Limit(limit int) *PaymentTelcoDetailQuery {
	ptdq.limit = &limit
	return ptdq
}

// Offset adds an offset step to the query.
func (ptdq *PaymentTelcoDetailQuery) Offset(offset int) *PaymentTelcoDetailQuery {
	ptdq.offset = &offset
	return ptdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptdq *PaymentTelcoDetailQuery) Unique(unique bool) *PaymentTelcoDetailQuery {
	ptdq.unique = &unique
	return ptdq
}

// Order adds an order step to the query.
func (ptdq *PaymentTelcoDetailQuery) Order(o ...OrderFunc) *PaymentTelcoDetailQuery {
	ptdq.order = append(ptdq.order, o...)
	return ptdq
}

// QueryPayment chains the current query on the "payment" edge.
func (ptdq *PaymentTelcoDetailQuery) QueryPayment() *PaymentQuery {
	query := &PaymentQuery{config: ptdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(paymenttelcodetail.Table, paymenttelcodetail.FieldID, selector),
			sqlgraph.To(payment.Table, payment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, paymenttelcodetail.PaymentTable, paymenttelcodetail.PaymentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PaymentTelcoDetail entity from the query.
// Returns a *NotFoundError when no PaymentTelcoDetail was found.
func (ptdq *PaymentTelcoDetailQuery) First(ctx context.Context) (*PaymentTelcoDetail, error) {
	nodes, err := ptdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{paymenttelcodetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptdq *PaymentTelcoDetailQuery) FirstX(ctx context.Context) *PaymentTelcoDetail {
	node, err := ptdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PaymentTelcoDetail ID from the query.
// Returns a *NotFoundError when no PaymentTelcoDetail ID was found.
func (ptdq *PaymentTelcoDetailQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ptdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{paymenttelcodetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptdq *PaymentTelcoDetailQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ptdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PaymentTelcoDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PaymentTelcoDetail entity is not found.
// Returns a *NotFoundError when no PaymentTelcoDetail entities are found.
func (ptdq *PaymentTelcoDetailQuery) Only(ctx context.Context) (*PaymentTelcoDetail, error) {
	nodes, err := ptdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{paymenttelcodetail.Label}
	default:
		return nil, &NotSingularError{paymenttelcodetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptdq *PaymentTelcoDetailQuery) OnlyX(ctx context.Context) *PaymentTelcoDetail {
	node, err := ptdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PaymentTelcoDetail ID in the query.
// Returns a *NotSingularError when exactly one PaymentTelcoDetail ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ptdq *PaymentTelcoDetailQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ptdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = &NotSingularError{paymenttelcodetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptdq *PaymentTelcoDetailQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ptdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PaymentTelcoDetails.
func (ptdq *PaymentTelcoDetailQuery) All(ctx context.Context) ([]*PaymentTelcoDetail, error) {
	if err := ptdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ptdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ptdq *PaymentTelcoDetailQuery) AllX(ctx context.Context) []*PaymentTelcoDetail {
	nodes, err := ptdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PaymentTelcoDetail IDs.
func (ptdq *PaymentTelcoDetailQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := ptdq.Select(paymenttelcodetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptdq *PaymentTelcoDetailQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ptdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptdq *PaymentTelcoDetailQuery) Count(ctx context.Context) (int, error) {
	if err := ptdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ptdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ptdq *PaymentTelcoDetailQuery) CountX(ctx context.Context) int {
	count, err := ptdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptdq *PaymentTelcoDetailQuery) Exist(ctx context.Context) (bool, error) {
	if err := ptdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ptdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ptdq *PaymentTelcoDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := ptdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaymentTelcoDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptdq *PaymentTelcoDetailQuery) Clone() *PaymentTelcoDetailQuery {
	if ptdq == nil {
		return nil
	}
	return &PaymentTelcoDetailQuery{
		config:      ptdq.config,
		limit:       ptdq.limit,
		offset:      ptdq.offset,
		order:       append([]OrderFunc{}, ptdq.order...),
		predicates:  append([]predicate.PaymentTelcoDetail{}, ptdq.predicates...),
		withPayment: ptdq.withPayment.Clone(),
		// clone intermediate query.
		sql:  ptdq.sql.Clone(),
		path: ptdq.path,
	}
}

// WithPayment tells the query-builder to eager-load the nodes that are connected to
// the "payment" edge. The optional arguments are used to configure the query builder of the edge.
func (ptdq *PaymentTelcoDetailQuery) WithPayment(opts ...func(*PaymentQuery)) *PaymentTelcoDetailQuery {
	query := &PaymentQuery{config: ptdq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptdq.withPayment = query
	return ptdq
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
//	client.PaymentTelcoDetail.Query().
//		GroupBy(paymenttelcodetail.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ptdq *PaymentTelcoDetailQuery) GroupBy(field string, fields ...string) *PaymentTelcoDetailGroupBy {
	group := &PaymentTelcoDetailGroupBy{config: ptdq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ptdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ptdq.sqlQuery(ctx), nil
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
//	client.PaymentTelcoDetail.Query().
//		Select(paymenttelcodetail.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ptdq *PaymentTelcoDetailQuery) Select(fields ...string) *PaymentTelcoDetailSelect {
	ptdq.fields = append(ptdq.fields, fields...)
	return &PaymentTelcoDetailSelect{PaymentTelcoDetailQuery: ptdq}
}

func (ptdq *PaymentTelcoDetailQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ptdq.fields {
		if !paymenttelcodetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptdq.path != nil {
		prev, err := ptdq.path(ctx)
		if err != nil {
			return err
		}
		ptdq.sql = prev
	}
	return nil
}

func (ptdq *PaymentTelcoDetailQuery) sqlAll(ctx context.Context) ([]*PaymentTelcoDetail, error) {
	var (
		nodes       = []*PaymentTelcoDetail{}
		withFKs     = ptdq.withFKs
		_spec       = ptdq.querySpec()
		loadedTypes = [1]bool{
			ptdq.withPayment != nil,
		}
	)
	if ptdq.withPayment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, paymenttelcodetail.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PaymentTelcoDetail{config: ptdq.config}
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
	if len(ptdq.modifiers) > 0 {
		_spec.Modifiers = ptdq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, ptdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ptdq.withPayment; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*PaymentTelcoDetail)
		for i := range nodes {
			if nodes[i].payment_payment_telco_detail == nil {
				continue
			}
			fk := *nodes[i].payment_payment_telco_detail
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(payment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "payment_payment_telco_detail" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Payment = n
			}
		}
	}

	return nodes, nil
}

func (ptdq *PaymentTelcoDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptdq.querySpec()
	if len(ptdq.modifiers) > 0 {
		_spec.Modifiers = ptdq.modifiers
	}
	return sqlgraph.CountNodes(ctx, ptdq.driver, _spec)
}

func (ptdq *PaymentTelcoDetailQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ptdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ptdq *PaymentTelcoDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymenttelcodetail.Table,
			Columns: paymenttelcodetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: paymenttelcodetail.FieldID,
			},
		},
		From:   ptdq.sql,
		Unique: true,
	}
	if unique := ptdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ptdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymenttelcodetail.FieldID)
		for i := range fields {
			if fields[i] != paymenttelcodetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptdq *PaymentTelcoDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptdq.driver.Dialect())
	t1 := builder.Table(paymenttelcodetail.Table)
	columns := ptdq.fields
	if len(columns) == 0 {
		columns = paymenttelcodetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptdq.sql != nil {
		selector = ptdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptdq.unique != nil && *ptdq.unique {
		selector.Distinct()
	}
	for _, m := range ptdq.modifiers {
		m(selector)
	}
	for _, p := range ptdq.predicates {
		p(selector)
	}
	for _, p := range ptdq.order {
		p(selector)
	}
	if offset := ptdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ptdq *PaymentTelcoDetailQuery) ForUpdate(opts ...sql.LockOption) *PaymentTelcoDetailQuery {
	if ptdq.driver.Dialect() == dialect.Postgres {
		ptdq.Unique(false)
	}
	ptdq.modifiers = append(ptdq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ptdq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ptdq *PaymentTelcoDetailQuery) ForShare(opts ...sql.LockOption) *PaymentTelcoDetailQuery {
	if ptdq.driver.Dialect() == dialect.Postgres {
		ptdq.Unique(false)
	}
	ptdq.modifiers = append(ptdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ptdq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ptdq *PaymentTelcoDetailQuery) Modify(modifiers ...func(s *sql.Selector)) *PaymentTelcoDetailSelect {
	ptdq.modifiers = append(ptdq.modifiers, modifiers...)
	return ptdq.Select()
}

// PaymentTelcoDetailGroupBy is the group-by builder for PaymentTelcoDetail entities.
type PaymentTelcoDetailGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptdgb *PaymentTelcoDetailGroupBy) Aggregate(fns ...AggregateFunc) *PaymentTelcoDetailGroupBy {
	ptdgb.fns = append(ptdgb.fns, fns...)
	return ptdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ptdgb *PaymentTelcoDetailGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ptdgb.path(ctx)
	if err != nil {
		return err
	}
	ptdgb.sql = query
	return ptdgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ptdgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptdgb *PaymentTelcoDetailGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ptdgb.fields) > 1 {
		return nil, errors.New("ent: PaymentTelcoDetailGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ptdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) StringsX(ctx context.Context) []string {
	v, err := ptdgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptdgb *PaymentTelcoDetailGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptdgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = fmt.Errorf("ent: PaymentTelcoDetailGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) StringX(ctx context.Context) string {
	v, err := ptdgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptdgb *PaymentTelcoDetailGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ptdgb.fields) > 1 {
		return nil, errors.New("ent: PaymentTelcoDetailGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ptdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) IntsX(ctx context.Context) []int {
	v, err := ptdgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptdgb *PaymentTelcoDetailGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptdgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = fmt.Errorf("ent: PaymentTelcoDetailGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) IntX(ctx context.Context) int {
	v, err := ptdgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptdgb *PaymentTelcoDetailGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptdgb.fields) > 1 {
		return nil, errors.New("ent: PaymentTelcoDetailGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ptdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ptdgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptdgb *PaymentTelcoDetailGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptdgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = fmt.Errorf("ent: PaymentTelcoDetailGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ptdgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptdgb *PaymentTelcoDetailGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ptdgb.fields) > 1 {
		return nil, errors.New("ent: PaymentTelcoDetailGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ptdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ptdgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptdgb *PaymentTelcoDetailGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptdgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = fmt.Errorf("ent: PaymentTelcoDetailGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptdgb *PaymentTelcoDetailGroupBy) BoolX(ctx context.Context) bool {
	v, err := ptdgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptdgb *PaymentTelcoDetailGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ptdgb.fields {
		if !paymenttelcodetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ptdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ptdgb *PaymentTelcoDetailGroupBy) sqlQuery() *sql.Selector {
	selector := ptdgb.sql.Select()
	aggregation := make([]string, 0, len(ptdgb.fns))
	for _, fn := range ptdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ptdgb.fields)+len(ptdgb.fns))
		for _, f := range ptdgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ptdgb.fields...)...)
}

// PaymentTelcoDetailSelect is the builder for selecting fields of PaymentTelcoDetail entities.
type PaymentTelcoDetailSelect struct {
	*PaymentTelcoDetailQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ptds *PaymentTelcoDetailSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ptds.prepareQuery(ctx); err != nil {
		return err
	}
	ptds.sql = ptds.PaymentTelcoDetailQuery.sqlQuery(ctx)
	return ptds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ptds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ptds *PaymentTelcoDetailSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ptds.fields) > 1 {
		return nil, errors.New("ent: PaymentTelcoDetailSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ptds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) StringsX(ctx context.Context) []string {
	v, err := ptds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ptds *PaymentTelcoDetailSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = fmt.Errorf("ent: PaymentTelcoDetailSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) StringX(ctx context.Context) string {
	v, err := ptds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ptds *PaymentTelcoDetailSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ptds.fields) > 1 {
		return nil, errors.New("ent: PaymentTelcoDetailSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ptds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) IntsX(ctx context.Context) []int {
	v, err := ptds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ptds *PaymentTelcoDetailSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = fmt.Errorf("ent: PaymentTelcoDetailSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) IntX(ctx context.Context) int {
	v, err := ptds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ptds *PaymentTelcoDetailSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptds.fields) > 1 {
		return nil, errors.New("ent: PaymentTelcoDetailSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ptds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ptds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ptds *PaymentTelcoDetailSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = fmt.Errorf("ent: PaymentTelcoDetailSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) Float64X(ctx context.Context) float64 {
	v, err := ptds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ptds *PaymentTelcoDetailSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ptds.fields) > 1 {
		return nil, errors.New("ent: PaymentTelcoDetailSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ptds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) BoolsX(ctx context.Context) []bool {
	v, err := ptds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ptds *PaymentTelcoDetailSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{paymenttelcodetail.Label}
	default:
		err = fmt.Errorf("ent: PaymentTelcoDetailSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptds *PaymentTelcoDetailSelect) BoolX(ctx context.Context) bool {
	v, err := ptds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptds *PaymentTelcoDetailSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ptds.sql.Query()
	if err := ptds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ptds *PaymentTelcoDetailSelect) Modify(modifiers ...func(s *sql.Selector)) *PaymentTelcoDetailSelect {
	ptds.modifiers = append(ptds.modifiers, modifiers...)
	return ptds
}
