// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/merchantuserbankaccount"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

// MerchantUserBankAccountUpdate is the builder for updating MerchantUserBankAccount entities.
type MerchantUserBankAccountUpdate struct {
	config
	hooks    []Hook
	mutation *MerchantUserBankAccountMutation
}

// Where appends a list predicates to the MerchantUserBankAccountUpdate builder.
func (mubau *MerchantUserBankAccountUpdate) Where(ps ...predicate.MerchantUserBankAccount) *MerchantUserBankAccountUpdate {
	mubau.mutation.Where(ps...)
	return mubau
}

// SetCreatedAt sets the "created_at" field.
func (mubau *MerchantUserBankAccountUpdate) SetCreatedAt(t time.Time) *MerchantUserBankAccountUpdate {
	mubau.mutation.SetCreatedAt(t)
	return mubau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mubau *MerchantUserBankAccountUpdate) SetNillableCreatedAt(t *time.Time) *MerchantUserBankAccountUpdate {
	if t != nil {
		mubau.SetCreatedAt(*t)
	}
	return mubau
}

// SetUpdatedAt sets the "updated_at" field.
func (mubau *MerchantUserBankAccountUpdate) SetUpdatedAt(t time.Time) *MerchantUserBankAccountUpdate {
	mubau.mutation.SetUpdatedAt(t)
	return mubau
}

// SetCreatedBy sets the "created_by" field.
func (mubau *MerchantUserBankAccountUpdate) SetCreatedBy(s string) *MerchantUserBankAccountUpdate {
	mubau.mutation.SetCreatedBy(s)
	return mubau
}

// SetUpdatedBy sets the "updated_by" field.
func (mubau *MerchantUserBankAccountUpdate) SetUpdatedBy(s string) *MerchantUserBankAccountUpdate {
	mubau.mutation.SetUpdatedBy(s)
	return mubau
}

// SetBankName sets the "bank_name" field.
func (mubau *MerchantUserBankAccountUpdate) SetBankName(i int32) *MerchantUserBankAccountUpdate {
	mubau.mutation.ResetBankName()
	mubau.mutation.SetBankName(i)
	return mubau
}

// SetNillableBankName sets the "bank_name" field if the given value is not nil.
func (mubau *MerchantUserBankAccountUpdate) SetNillableBankName(i *int32) *MerchantUserBankAccountUpdate {
	if i != nil {
		mubau.SetBankName(*i)
	}
	return mubau
}

// AddBankName adds i to the "bank_name" field.
func (mubau *MerchantUserBankAccountUpdate) AddBankName(i int32) *MerchantUserBankAccountUpdate {
	mubau.mutation.AddBankName(i)
	return mubau
}

// SetAccountNumber sets the "account_number" field.
func (mubau *MerchantUserBankAccountUpdate) SetAccountNumber(s string) *MerchantUserBankAccountUpdate {
	mubau.mutation.SetAccountNumber(s)
	return mubau
}

// SetAccountName sets the "account_name" field.
func (mubau *MerchantUserBankAccountUpdate) SetAccountName(s string) *MerchantUserBankAccountUpdate {
	mubau.mutation.SetAccountName(s)
	return mubau
}

// Mutation returns the MerchantUserBankAccountMutation object of the builder.
func (mubau *MerchantUserBankAccountUpdate) Mutation() *MerchantUserBankAccountMutation {
	return mubau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mubau *MerchantUserBankAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mubau.defaults()
	if len(mubau.hooks) == 0 {
		if err = mubau.check(); err != nil {
			return 0, err
		}
		affected, err = mubau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchantUserBankAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mubau.check(); err != nil {
				return 0, err
			}
			mubau.mutation = mutation
			affected, err = mubau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mubau.hooks) - 1; i >= 0; i-- {
			if mubau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mubau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mubau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mubau *MerchantUserBankAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := mubau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mubau *MerchantUserBankAccountUpdate) Exec(ctx context.Context) error {
	_, err := mubau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mubau *MerchantUserBankAccountUpdate) ExecX(ctx context.Context) {
	if err := mubau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mubau *MerchantUserBankAccountUpdate) defaults() {
	if _, ok := mubau.mutation.UpdatedAt(); !ok {
		v := merchantuserbankaccount.UpdateDefaultUpdatedAt()
		mubau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mubau *MerchantUserBankAccountUpdate) check() error {
	if v, ok := mubau.mutation.CreatedBy(); ok {
		if err := merchantuserbankaccount.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.created_by": %w`, err)}
		}
	}
	if v, ok := mubau.mutation.UpdatedBy(); ok {
		if err := merchantuserbankaccount.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.updated_by": %w`, err)}
		}
	}
	if v, ok := mubau.mutation.BankName(); ok {
		if err := merchantuserbankaccount.BankNameValidator(v); err != nil {
			return &ValidationError{Name: "bank_name", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.bank_name": %w`, err)}
		}
	}
	if v, ok := mubau.mutation.AccountNumber(); ok {
		if err := merchantuserbankaccount.AccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_number", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.account_number": %w`, err)}
		}
	}
	if v, ok := mubau.mutation.AccountName(); ok {
		if err := merchantuserbankaccount.AccountNameValidator(v); err != nil {
			return &ValidationError{Name: "account_name", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.account_name": %w`, err)}
		}
	}
	return nil
}

func (mubau *MerchantUserBankAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   merchantuserbankaccount.Table,
			Columns: merchantuserbankaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: merchantuserbankaccount.FieldID,
			},
		},
	}
	if ps := mubau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mubau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchantuserbankaccount.FieldCreatedAt,
		})
	}
	if value, ok := mubau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchantuserbankaccount.FieldUpdatedAt,
		})
	}
	if value, ok := mubau.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantuserbankaccount.FieldCreatedBy,
		})
	}
	if value, ok := mubau.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantuserbankaccount.FieldUpdatedBy,
		})
	}
	if value, ok := mubau.mutation.BankName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: merchantuserbankaccount.FieldBankName,
		})
	}
	if value, ok := mubau.mutation.AddedBankName(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: merchantuserbankaccount.FieldBankName,
		})
	}
	if value, ok := mubau.mutation.AccountNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantuserbankaccount.FieldAccountNumber,
		})
	}
	if value, ok := mubau.mutation.AccountName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantuserbankaccount.FieldAccountName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mubau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchantuserbankaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MerchantUserBankAccountUpdateOne is the builder for updating a single MerchantUserBankAccount entity.
type MerchantUserBankAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MerchantUserBankAccountMutation
}

// SetCreatedAt sets the "created_at" field.
func (mubauo *MerchantUserBankAccountUpdateOne) SetCreatedAt(t time.Time) *MerchantUserBankAccountUpdateOne {
	mubauo.mutation.SetCreatedAt(t)
	return mubauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mubauo *MerchantUserBankAccountUpdateOne) SetNillableCreatedAt(t *time.Time) *MerchantUserBankAccountUpdateOne {
	if t != nil {
		mubauo.SetCreatedAt(*t)
	}
	return mubauo
}

// SetUpdatedAt sets the "updated_at" field.
func (mubauo *MerchantUserBankAccountUpdateOne) SetUpdatedAt(t time.Time) *MerchantUserBankAccountUpdateOne {
	mubauo.mutation.SetUpdatedAt(t)
	return mubauo
}

// SetCreatedBy sets the "created_by" field.
func (mubauo *MerchantUserBankAccountUpdateOne) SetCreatedBy(s string) *MerchantUserBankAccountUpdateOne {
	mubauo.mutation.SetCreatedBy(s)
	return mubauo
}

// SetUpdatedBy sets the "updated_by" field.
func (mubauo *MerchantUserBankAccountUpdateOne) SetUpdatedBy(s string) *MerchantUserBankAccountUpdateOne {
	mubauo.mutation.SetUpdatedBy(s)
	return mubauo
}

// SetBankName sets the "bank_name" field.
func (mubauo *MerchantUserBankAccountUpdateOne) SetBankName(i int32) *MerchantUserBankAccountUpdateOne {
	mubauo.mutation.ResetBankName()
	mubauo.mutation.SetBankName(i)
	return mubauo
}

// SetNillableBankName sets the "bank_name" field if the given value is not nil.
func (mubauo *MerchantUserBankAccountUpdateOne) SetNillableBankName(i *int32) *MerchantUserBankAccountUpdateOne {
	if i != nil {
		mubauo.SetBankName(*i)
	}
	return mubauo
}

// AddBankName adds i to the "bank_name" field.
func (mubauo *MerchantUserBankAccountUpdateOne) AddBankName(i int32) *MerchantUserBankAccountUpdateOne {
	mubauo.mutation.AddBankName(i)
	return mubauo
}

// SetAccountNumber sets the "account_number" field.
func (mubauo *MerchantUserBankAccountUpdateOne) SetAccountNumber(s string) *MerchantUserBankAccountUpdateOne {
	mubauo.mutation.SetAccountNumber(s)
	return mubauo
}

// SetAccountName sets the "account_name" field.
func (mubauo *MerchantUserBankAccountUpdateOne) SetAccountName(s string) *MerchantUserBankAccountUpdateOne {
	mubauo.mutation.SetAccountName(s)
	return mubauo
}

// Mutation returns the MerchantUserBankAccountMutation object of the builder.
func (mubauo *MerchantUserBankAccountUpdateOne) Mutation() *MerchantUserBankAccountMutation {
	return mubauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mubauo *MerchantUserBankAccountUpdateOne) Select(field string, fields ...string) *MerchantUserBankAccountUpdateOne {
	mubauo.fields = append([]string{field}, fields...)
	return mubauo
}

// Save executes the query and returns the updated MerchantUserBankAccount entity.
func (mubauo *MerchantUserBankAccountUpdateOne) Save(ctx context.Context) (*MerchantUserBankAccount, error) {
	var (
		err  error
		node *MerchantUserBankAccount
	)
	mubauo.defaults()
	if len(mubauo.hooks) == 0 {
		if err = mubauo.check(); err != nil {
			return nil, err
		}
		node, err = mubauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchantUserBankAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mubauo.check(); err != nil {
				return nil, err
			}
			mubauo.mutation = mutation
			node, err = mubauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mubauo.hooks) - 1; i >= 0; i-- {
			if mubauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mubauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mubauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mubauo *MerchantUserBankAccountUpdateOne) SaveX(ctx context.Context) *MerchantUserBankAccount {
	node, err := mubauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mubauo *MerchantUserBankAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := mubauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mubauo *MerchantUserBankAccountUpdateOne) ExecX(ctx context.Context) {
	if err := mubauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mubauo *MerchantUserBankAccountUpdateOne) defaults() {
	if _, ok := mubauo.mutation.UpdatedAt(); !ok {
		v := merchantuserbankaccount.UpdateDefaultUpdatedAt()
		mubauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mubauo *MerchantUserBankAccountUpdateOne) check() error {
	if v, ok := mubauo.mutation.CreatedBy(); ok {
		if err := merchantuserbankaccount.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.created_by": %w`, err)}
		}
	}
	if v, ok := mubauo.mutation.UpdatedBy(); ok {
		if err := merchantuserbankaccount.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.updated_by": %w`, err)}
		}
	}
	if v, ok := mubauo.mutation.BankName(); ok {
		if err := merchantuserbankaccount.BankNameValidator(v); err != nil {
			return &ValidationError{Name: "bank_name", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.bank_name": %w`, err)}
		}
	}
	if v, ok := mubauo.mutation.AccountNumber(); ok {
		if err := merchantuserbankaccount.AccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_number", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.account_number": %w`, err)}
		}
	}
	if v, ok := mubauo.mutation.AccountName(); ok {
		if err := merchantuserbankaccount.AccountNameValidator(v); err != nil {
			return &ValidationError{Name: "account_name", err: fmt.Errorf(`ent: validator failed for field "MerchantUserBankAccount.account_name": %w`, err)}
		}
	}
	return nil
}

func (mubauo *MerchantUserBankAccountUpdateOne) sqlSave(ctx context.Context) (_node *MerchantUserBankAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   merchantuserbankaccount.Table,
			Columns: merchantuserbankaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: merchantuserbankaccount.FieldID,
			},
		},
	}
	id, ok := mubauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MerchantUserBankAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mubauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchantuserbankaccount.FieldID)
		for _, f := range fields {
			if !merchantuserbankaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != merchantuserbankaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mubauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mubauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchantuserbankaccount.FieldCreatedAt,
		})
	}
	if value, ok := mubauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchantuserbankaccount.FieldUpdatedAt,
		})
	}
	if value, ok := mubauo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantuserbankaccount.FieldCreatedBy,
		})
	}
	if value, ok := mubauo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantuserbankaccount.FieldUpdatedBy,
		})
	}
	if value, ok := mubauo.mutation.BankName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: merchantuserbankaccount.FieldBankName,
		})
	}
	if value, ok := mubauo.mutation.AddedBankName(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: merchantuserbankaccount.FieldBankName,
		})
	}
	if value, ok := mubauo.mutation.AccountNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantuserbankaccount.FieldAccountNumber,
		})
	}
	if value, ok := mubauo.mutation.AccountName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchantuserbankaccount.FieldAccountName,
		})
	}
	_node = &MerchantUserBankAccount{config: mubauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mubauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchantuserbankaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}