// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/predicate"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where adds a new predicate for the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.predicates = append(pu.mutation.predicates, ps...)
	return pu
}

// SetMerchantID sets the "merchant_id" field.
func (pu *PaymentUpdate) SetMerchantID(i int64) *PaymentUpdate {
	pu.mutation.ResetMerchantID()
	pu.mutation.SetMerchantID(i)
	return pu
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableMerchantID(i *int64) *PaymentUpdate {
	if i != nil {
		pu.SetMerchantID(*i)
	}
	return pu
}

// AddMerchantID adds i to the "merchant_id" field.
func (pu *PaymentUpdate) AddMerchantID(i int64) *PaymentUpdate {
	pu.mutation.AddMerchantID(i)
	return pu
}

// SetType sets the "type" field.
func (pu *PaymentUpdate) SetType(i int32) *PaymentUpdate {
	pu.mutation.ResetType()
	pu.mutation.SetType(i)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableType(i *int32) *PaymentUpdate {
	if i != nil {
		pu.SetType(*i)
	}
	return pu
}

// AddType adds i to the "type" field.
func (pu *PaymentUpdate) AddType(i int32) *PaymentUpdate {
	pu.mutation.AddType(i)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(i int64) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(i)
	return pu
}

// AddAmount adds i to the "amount" field.
func (pu *PaymentUpdate) AddAmount(i int64) *PaymentUpdate {
	pu.mutation.AddAmount(i)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PaymentUpdate) SetCreatedAt(t time.Time) *PaymentUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableCreatedAt(t *time.Time) *PaymentUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PaymentUpdate) SetUpdatedAt(t time.Time) *PaymentUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *PaymentUpdate) SetCreatedBy(s string) *PaymentUpdate {
	pu.mutation.SetCreatedBy(s)
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *PaymentUpdate) SetUpdatedBy(s string) *PaymentUpdate {
	pu.mutation.SetUpdatedBy(s)
	return pu
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PaymentUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := payment.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.MerchantID(); ok {
		if err := payment.MerchantIDValidator(v); err != nil {
			return &ValidationError{Name: "merchant_id", err: fmt.Errorf("ent: validator failed for field \"merchant_id\": %w", err)}
		}
	}
	if v, ok := pu.mutation.GetType(); ok {
		if err := payment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := pu.mutation.CreatedBy(); ok {
		if err := payment.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf("ent: validator failed for field \"created_by\": %w", err)}
		}
	}
	if v, ok := pu.mutation.UpdatedBy(); ok {
		if err := payment.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf("ent: validator failed for field \"updated_by\": %w", err)}
		}
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: payment.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.MerchantID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: payment.FieldMerchantID,
		})
	}
	if value, ok := pu.mutation.AddedMerchantID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: payment.FieldMerchantID,
		})
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: payment.FieldType,
		})
	}
	if value, ok := pu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: payment.FieldType,
		})
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldCreatedBy,
		})
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldUpdatedBy,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// SetMerchantID sets the "merchant_id" field.
func (puo *PaymentUpdateOne) SetMerchantID(i int64) *PaymentUpdateOne {
	puo.mutation.ResetMerchantID()
	puo.mutation.SetMerchantID(i)
	return puo
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableMerchantID(i *int64) *PaymentUpdateOne {
	if i != nil {
		puo.SetMerchantID(*i)
	}
	return puo
}

// AddMerchantID adds i to the "merchant_id" field.
func (puo *PaymentUpdateOne) AddMerchantID(i int64) *PaymentUpdateOne {
	puo.mutation.AddMerchantID(i)
	return puo
}

// SetType sets the "type" field.
func (puo *PaymentUpdateOne) SetType(i int32) *PaymentUpdateOne {
	puo.mutation.ResetType()
	puo.mutation.SetType(i)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableType(i *int32) *PaymentUpdateOne {
	if i != nil {
		puo.SetType(*i)
	}
	return puo
}

// AddType adds i to the "type" field.
func (puo *PaymentUpdateOne) AddType(i int32) *PaymentUpdateOne {
	puo.mutation.AddType(i)
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(i int64) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(i)
	return puo
}

// AddAmount adds i to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(i int64) *PaymentUpdateOne {
	puo.mutation.AddAmount(i)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PaymentUpdateOne) SetCreatedAt(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableCreatedAt(t *time.Time) *PaymentUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PaymentUpdateOne) SetUpdatedAt(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetCreatedBy sets the "created_by" field.
func (puo *PaymentUpdateOne) SetCreatedBy(s string) *PaymentUpdateOne {
	puo.mutation.SetCreatedBy(s)
	return puo
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *PaymentUpdateOne) SetUpdatedBy(s string) *PaymentUpdateOne {
	puo.mutation.SetUpdatedBy(s)
	return puo
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	var (
		err  error
		node *Payment
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PaymentUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := payment.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.MerchantID(); ok {
		if err := payment.MerchantIDValidator(v); err != nil {
			return &ValidationError{Name: "merchant_id", err: fmt.Errorf("ent: validator failed for field \"merchant_id\": %w", err)}
		}
	}
	if v, ok := puo.mutation.GetType(); ok {
		if err := payment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := puo.mutation.CreatedBy(); ok {
		if err := payment.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf("ent: validator failed for field \"created_by\": %w", err)}
		}
	}
	if v, ok := puo.mutation.UpdatedBy(); ok {
		if err := payment.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf("ent: validator failed for field \"updated_by\": %w", err)}
		}
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: payment.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Payment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.MerchantID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: payment.FieldMerchantID,
		})
	}
	if value, ok := puo.mutation.AddedMerchantID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: payment.FieldMerchantID,
		})
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: payment.FieldType,
		})
	}
	if value, ok := puo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: payment.FieldType,
		})
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: payment.FieldAmount,
		})
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldCreatedBy,
		})
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldUpdatedBy,
		})
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
