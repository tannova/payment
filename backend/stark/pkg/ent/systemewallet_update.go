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
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemewallet"
)

// SystemEWalletUpdate is the builder for updating SystemEWallet entities.
type SystemEWalletUpdate struct {
	config
	hooks    []Hook
	mutation *SystemEWalletMutation
}

// Where appends a list predicates to the SystemEWalletUpdate builder.
func (seu *SystemEWalletUpdate) Where(ps ...predicate.SystemEWallet) *SystemEWalletUpdate {
	seu.mutation.Where(ps...)
	return seu
}

// SetCreatedAt sets the "created_at" field.
func (seu *SystemEWalletUpdate) SetCreatedAt(t time.Time) *SystemEWalletUpdate {
	seu.mutation.SetCreatedAt(t)
	return seu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableCreatedAt(t *time.Time) *SystemEWalletUpdate {
	if t != nil {
		seu.SetCreatedAt(*t)
	}
	return seu
}

// SetUpdatedAt sets the "updated_at" field.
func (seu *SystemEWalletUpdate) SetUpdatedAt(t time.Time) *SystemEWalletUpdate {
	seu.mutation.SetUpdatedAt(t)
	return seu
}

// SetCreatedBy sets the "created_by" field.
func (seu *SystemEWalletUpdate) SetCreatedBy(s string) *SystemEWalletUpdate {
	seu.mutation.SetCreatedBy(s)
	return seu
}

// SetUpdatedBy sets the "updated_by" field.
func (seu *SystemEWalletUpdate) SetUpdatedBy(s string) *SystemEWalletUpdate {
	seu.mutation.SetUpdatedBy(s)
	return seu
}

// SetEWalletName sets the "e_wallet_name" field.
func (seu *SystemEWalletUpdate) SetEWalletName(i int32) *SystemEWalletUpdate {
	seu.mutation.ResetEWalletName()
	seu.mutation.SetEWalletName(i)
	return seu
}

// SetNillableEWalletName sets the "e_wallet_name" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableEWalletName(i *int32) *SystemEWalletUpdate {
	if i != nil {
		seu.SetEWalletName(*i)
	}
	return seu
}

// AddEWalletName adds i to the "e_wallet_name" field.
func (seu *SystemEWalletUpdate) AddEWalletName(i int32) *SystemEWalletUpdate {
	seu.mutation.AddEWalletName(i)
	return seu
}

// SetStatus sets the "status" field.
func (seu *SystemEWalletUpdate) SetStatus(i int32) *SystemEWalletUpdate {
	seu.mutation.ResetStatus()
	seu.mutation.SetStatus(i)
	return seu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableStatus(i *int32) *SystemEWalletUpdate {
	if i != nil {
		seu.SetStatus(*i)
	}
	return seu
}

// AddStatus adds i to the "status" field.
func (seu *SystemEWalletUpdate) AddStatus(i int32) *SystemEWalletUpdate {
	seu.mutation.AddStatus(i)
	return seu
}

// SetMerchantID sets the "merchant_id" field.
func (seu *SystemEWalletUpdate) SetMerchantID(i int64) *SystemEWalletUpdate {
	seu.mutation.ResetMerchantID()
	seu.mutation.SetMerchantID(i)
	return seu
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableMerchantID(i *int64) *SystemEWalletUpdate {
	if i != nil {
		seu.SetMerchantID(*i)
	}
	return seu
}

// AddMerchantID adds i to the "merchant_id" field.
func (seu *SystemEWalletUpdate) AddMerchantID(i int64) *SystemEWalletUpdate {
	seu.mutation.AddMerchantID(i)
	return seu
}

// SetAccountPhoneNumber sets the "account_phone_number" field.
func (seu *SystemEWalletUpdate) SetAccountPhoneNumber(s string) *SystemEWalletUpdate {
	seu.mutation.SetAccountPhoneNumber(s)
	return seu
}

// SetAccountName sets the "account_name" field.
func (seu *SystemEWalletUpdate) SetAccountName(s string) *SystemEWalletUpdate {
	seu.mutation.SetAccountName(s)
	return seu
}

// SetBalance sets the "balance" field.
func (seu *SystemEWalletUpdate) SetBalance(u uint64) *SystemEWalletUpdate {
	seu.mutation.ResetBalance()
	seu.mutation.SetBalance(u)
	return seu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableBalance(u *uint64) *SystemEWalletUpdate {
	if u != nil {
		seu.SetBalance(*u)
	}
	return seu
}

// AddBalance adds u to the "balance" field.
func (seu *SystemEWalletUpdate) AddBalance(u uint64) *SystemEWalletUpdate {
	seu.mutation.AddBalance(u)
	return seu
}

// SetDailyBalance sets the "daily_balance" field.
func (seu *SystemEWalletUpdate) SetDailyBalance(u uint64) *SystemEWalletUpdate {
	seu.mutation.ResetDailyBalance()
	seu.mutation.SetDailyBalance(u)
	return seu
}

// SetNillableDailyBalance sets the "daily_balance" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableDailyBalance(u *uint64) *SystemEWalletUpdate {
	if u != nil {
		seu.SetDailyBalance(*u)
	}
	return seu
}

// AddDailyBalance adds u to the "daily_balance" field.
func (seu *SystemEWalletUpdate) AddDailyBalance(u uint64) *SystemEWalletUpdate {
	seu.mutation.AddDailyBalance(u)
	return seu
}

// SetDailyBalanceLimit sets the "daily_balance_limit" field.
func (seu *SystemEWalletUpdate) SetDailyBalanceLimit(u uint64) *SystemEWalletUpdate {
	seu.mutation.ResetDailyBalanceLimit()
	seu.mutation.SetDailyBalanceLimit(u)
	return seu
}

// SetNillableDailyBalanceLimit sets the "daily_balance_limit" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableDailyBalanceLimit(u *uint64) *SystemEWalletUpdate {
	if u != nil {
		seu.SetDailyBalanceLimit(*u)
	}
	return seu
}

// AddDailyBalanceLimit adds u to the "daily_balance_limit" field.
func (seu *SystemEWalletUpdate) AddDailyBalanceLimit(u uint64) *SystemEWalletUpdate {
	seu.mutation.AddDailyBalanceLimit(u)
	return seu
}

// SetDailyUsedAmount sets the "daily_used_amount" field.
func (seu *SystemEWalletUpdate) SetDailyUsedAmount(i int64) *SystemEWalletUpdate {
	seu.mutation.ResetDailyUsedAmount()
	seu.mutation.SetDailyUsedAmount(i)
	return seu
}

// SetNillableDailyUsedAmount sets the "daily_used_amount" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableDailyUsedAmount(i *int64) *SystemEWalletUpdate {
	if i != nil {
		seu.SetDailyUsedAmount(*i)
	}
	return seu
}

// AddDailyUsedAmount adds i to the "daily_used_amount" field.
func (seu *SystemEWalletUpdate) AddDailyUsedAmount(i int64) *SystemEWalletUpdate {
	seu.mutation.AddDailyUsedAmount(i)
	return seu
}

// SetLastUpdatedBalance sets the "last_updated_balance" field.
func (seu *SystemEWalletUpdate) SetLastUpdatedBalance(t time.Time) *SystemEWalletUpdate {
	seu.mutation.SetLastUpdatedBalance(t)
	return seu
}

// SetNillableLastUpdatedBalance sets the "last_updated_balance" field if the given value is not nil.
func (seu *SystemEWalletUpdate) SetNillableLastUpdatedBalance(t *time.Time) *SystemEWalletUpdate {
	if t != nil {
		seu.SetLastUpdatedBalance(*t)
	}
	return seu
}

// ClearLastUpdatedBalance clears the value of the "last_updated_balance" field.
func (seu *SystemEWalletUpdate) ClearLastUpdatedBalance() *SystemEWalletUpdate {
	seu.mutation.ClearLastUpdatedBalance()
	return seu
}

// Mutation returns the SystemEWalletMutation object of the builder.
func (seu *SystemEWalletUpdate) Mutation() *SystemEWalletMutation {
	return seu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (seu *SystemEWalletUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	seu.defaults()
	if len(seu.hooks) == 0 {
		if err = seu.check(); err != nil {
			return 0, err
		}
		affected, err = seu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemEWalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = seu.check(); err != nil {
				return 0, err
			}
			seu.mutation = mutation
			affected, err = seu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(seu.hooks) - 1; i >= 0; i-- {
			if seu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = seu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, seu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (seu *SystemEWalletUpdate) SaveX(ctx context.Context) int {
	affected, err := seu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seu *SystemEWalletUpdate) Exec(ctx context.Context) error {
	_, err := seu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seu *SystemEWalletUpdate) ExecX(ctx context.Context) {
	if err := seu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (seu *SystemEWalletUpdate) defaults() {
	if _, ok := seu.mutation.UpdatedAt(); !ok {
		v := systemewallet.UpdateDefaultUpdatedAt()
		seu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (seu *SystemEWalletUpdate) check() error {
	if v, ok := seu.mutation.CreatedBy(); ok {
		if err := systemewallet.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.created_by": %w`, err)}
		}
	}
	if v, ok := seu.mutation.UpdatedBy(); ok {
		if err := systemewallet.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.updated_by": %w`, err)}
		}
	}
	if v, ok := seu.mutation.EWalletName(); ok {
		if err := systemewallet.EWalletNameValidator(v); err != nil {
			return &ValidationError{Name: "e_wallet_name", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.e_wallet_name": %w`, err)}
		}
	}
	if v, ok := seu.mutation.Status(); ok {
		if err := systemewallet.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.status": %w`, err)}
		}
	}
	if v, ok := seu.mutation.MerchantID(); ok {
		if err := systemewallet.MerchantIDValidator(v); err != nil {
			return &ValidationError{Name: "merchant_id", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.merchant_id": %w`, err)}
		}
	}
	if v, ok := seu.mutation.AccountPhoneNumber(); ok {
		if err := systemewallet.AccountPhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_phone_number", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.account_phone_number": %w`, err)}
		}
	}
	if v, ok := seu.mutation.AccountName(); ok {
		if err := systemewallet.AccountNameValidator(v); err != nil {
			return &ValidationError{Name: "account_name", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.account_name": %w`, err)}
		}
	}
	return nil
}

func (seu *SystemEWalletUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systemewallet.Table,
			Columns: systemewallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: systemewallet.FieldID,
			},
		},
	}
	if ps := seu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemewallet.FieldCreatedAt,
		})
	}
	if value, ok := seu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemewallet.FieldUpdatedAt,
		})
	}
	if value, ok := seu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemewallet.FieldCreatedBy,
		})
	}
	if value, ok := seu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemewallet.FieldUpdatedBy,
		})
	}
	if value, ok := seu.mutation.EWalletName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemewallet.FieldEWalletName,
		})
	}
	if value, ok := seu.mutation.AddedEWalletName(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemewallet.FieldEWalletName,
		})
	}
	if value, ok := seu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemewallet.FieldStatus,
		})
	}
	if value, ok := seu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemewallet.FieldStatus,
		})
	}
	if value, ok := seu.mutation.MerchantID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemewallet.FieldMerchantID,
		})
	}
	if value, ok := seu.mutation.AddedMerchantID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemewallet.FieldMerchantID,
		})
	}
	if value, ok := seu.mutation.AccountPhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemewallet.FieldAccountPhoneNumber,
		})
	}
	if value, ok := seu.mutation.AccountName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemewallet.FieldAccountName,
		})
	}
	if value, ok := seu.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldBalance,
		})
	}
	if value, ok := seu.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldBalance,
		})
	}
	if value, ok := seu.mutation.DailyBalance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldDailyBalance,
		})
	}
	if value, ok := seu.mutation.AddedDailyBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldDailyBalance,
		})
	}
	if value, ok := seu.mutation.DailyBalanceLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldDailyBalanceLimit,
		})
	}
	if value, ok := seu.mutation.AddedDailyBalanceLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldDailyBalanceLimit,
		})
	}
	if value, ok := seu.mutation.DailyUsedAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemewallet.FieldDailyUsedAmount,
		})
	}
	if value, ok := seu.mutation.AddedDailyUsedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemewallet.FieldDailyUsedAmount,
		})
	}
	if value, ok := seu.mutation.LastUpdatedBalance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemewallet.FieldLastUpdatedBalance,
		})
	}
	if seu.mutation.LastUpdatedBalanceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: systemewallet.FieldLastUpdatedBalance,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemewallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SystemEWalletUpdateOne is the builder for updating a single SystemEWallet entity.
type SystemEWalletUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SystemEWalletMutation
}

// SetCreatedAt sets the "created_at" field.
func (seuo *SystemEWalletUpdateOne) SetCreatedAt(t time.Time) *SystemEWalletUpdateOne {
	seuo.mutation.SetCreatedAt(t)
	return seuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableCreatedAt(t *time.Time) *SystemEWalletUpdateOne {
	if t != nil {
		seuo.SetCreatedAt(*t)
	}
	return seuo
}

// SetUpdatedAt sets the "updated_at" field.
func (seuo *SystemEWalletUpdateOne) SetUpdatedAt(t time.Time) *SystemEWalletUpdateOne {
	seuo.mutation.SetUpdatedAt(t)
	return seuo
}

// SetCreatedBy sets the "created_by" field.
func (seuo *SystemEWalletUpdateOne) SetCreatedBy(s string) *SystemEWalletUpdateOne {
	seuo.mutation.SetCreatedBy(s)
	return seuo
}

// SetUpdatedBy sets the "updated_by" field.
func (seuo *SystemEWalletUpdateOne) SetUpdatedBy(s string) *SystemEWalletUpdateOne {
	seuo.mutation.SetUpdatedBy(s)
	return seuo
}

// SetEWalletName sets the "e_wallet_name" field.
func (seuo *SystemEWalletUpdateOne) SetEWalletName(i int32) *SystemEWalletUpdateOne {
	seuo.mutation.ResetEWalletName()
	seuo.mutation.SetEWalletName(i)
	return seuo
}

// SetNillableEWalletName sets the "e_wallet_name" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableEWalletName(i *int32) *SystemEWalletUpdateOne {
	if i != nil {
		seuo.SetEWalletName(*i)
	}
	return seuo
}

// AddEWalletName adds i to the "e_wallet_name" field.
func (seuo *SystemEWalletUpdateOne) AddEWalletName(i int32) *SystemEWalletUpdateOne {
	seuo.mutation.AddEWalletName(i)
	return seuo
}

// SetStatus sets the "status" field.
func (seuo *SystemEWalletUpdateOne) SetStatus(i int32) *SystemEWalletUpdateOne {
	seuo.mutation.ResetStatus()
	seuo.mutation.SetStatus(i)
	return seuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableStatus(i *int32) *SystemEWalletUpdateOne {
	if i != nil {
		seuo.SetStatus(*i)
	}
	return seuo
}

// AddStatus adds i to the "status" field.
func (seuo *SystemEWalletUpdateOne) AddStatus(i int32) *SystemEWalletUpdateOne {
	seuo.mutation.AddStatus(i)
	return seuo
}

// SetMerchantID sets the "merchant_id" field.
func (seuo *SystemEWalletUpdateOne) SetMerchantID(i int64) *SystemEWalletUpdateOne {
	seuo.mutation.ResetMerchantID()
	seuo.mutation.SetMerchantID(i)
	return seuo
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableMerchantID(i *int64) *SystemEWalletUpdateOne {
	if i != nil {
		seuo.SetMerchantID(*i)
	}
	return seuo
}

// AddMerchantID adds i to the "merchant_id" field.
func (seuo *SystemEWalletUpdateOne) AddMerchantID(i int64) *SystemEWalletUpdateOne {
	seuo.mutation.AddMerchantID(i)
	return seuo
}

// SetAccountPhoneNumber sets the "account_phone_number" field.
func (seuo *SystemEWalletUpdateOne) SetAccountPhoneNumber(s string) *SystemEWalletUpdateOne {
	seuo.mutation.SetAccountPhoneNumber(s)
	return seuo
}

// SetAccountName sets the "account_name" field.
func (seuo *SystemEWalletUpdateOne) SetAccountName(s string) *SystemEWalletUpdateOne {
	seuo.mutation.SetAccountName(s)
	return seuo
}

// SetBalance sets the "balance" field.
func (seuo *SystemEWalletUpdateOne) SetBalance(u uint64) *SystemEWalletUpdateOne {
	seuo.mutation.ResetBalance()
	seuo.mutation.SetBalance(u)
	return seuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableBalance(u *uint64) *SystemEWalletUpdateOne {
	if u != nil {
		seuo.SetBalance(*u)
	}
	return seuo
}

// AddBalance adds u to the "balance" field.
func (seuo *SystemEWalletUpdateOne) AddBalance(u uint64) *SystemEWalletUpdateOne {
	seuo.mutation.AddBalance(u)
	return seuo
}

// SetDailyBalance sets the "daily_balance" field.
func (seuo *SystemEWalletUpdateOne) SetDailyBalance(u uint64) *SystemEWalletUpdateOne {
	seuo.mutation.ResetDailyBalance()
	seuo.mutation.SetDailyBalance(u)
	return seuo
}

// SetNillableDailyBalance sets the "daily_balance" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableDailyBalance(u *uint64) *SystemEWalletUpdateOne {
	if u != nil {
		seuo.SetDailyBalance(*u)
	}
	return seuo
}

// AddDailyBalance adds u to the "daily_balance" field.
func (seuo *SystemEWalletUpdateOne) AddDailyBalance(u uint64) *SystemEWalletUpdateOne {
	seuo.mutation.AddDailyBalance(u)
	return seuo
}

// SetDailyBalanceLimit sets the "daily_balance_limit" field.
func (seuo *SystemEWalletUpdateOne) SetDailyBalanceLimit(u uint64) *SystemEWalletUpdateOne {
	seuo.mutation.ResetDailyBalanceLimit()
	seuo.mutation.SetDailyBalanceLimit(u)
	return seuo
}

// SetNillableDailyBalanceLimit sets the "daily_balance_limit" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableDailyBalanceLimit(u *uint64) *SystemEWalletUpdateOne {
	if u != nil {
		seuo.SetDailyBalanceLimit(*u)
	}
	return seuo
}

// AddDailyBalanceLimit adds u to the "daily_balance_limit" field.
func (seuo *SystemEWalletUpdateOne) AddDailyBalanceLimit(u uint64) *SystemEWalletUpdateOne {
	seuo.mutation.AddDailyBalanceLimit(u)
	return seuo
}

// SetDailyUsedAmount sets the "daily_used_amount" field.
func (seuo *SystemEWalletUpdateOne) SetDailyUsedAmount(i int64) *SystemEWalletUpdateOne {
	seuo.mutation.ResetDailyUsedAmount()
	seuo.mutation.SetDailyUsedAmount(i)
	return seuo
}

// SetNillableDailyUsedAmount sets the "daily_used_amount" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableDailyUsedAmount(i *int64) *SystemEWalletUpdateOne {
	if i != nil {
		seuo.SetDailyUsedAmount(*i)
	}
	return seuo
}

// AddDailyUsedAmount adds i to the "daily_used_amount" field.
func (seuo *SystemEWalletUpdateOne) AddDailyUsedAmount(i int64) *SystemEWalletUpdateOne {
	seuo.mutation.AddDailyUsedAmount(i)
	return seuo
}

// SetLastUpdatedBalance sets the "last_updated_balance" field.
func (seuo *SystemEWalletUpdateOne) SetLastUpdatedBalance(t time.Time) *SystemEWalletUpdateOne {
	seuo.mutation.SetLastUpdatedBalance(t)
	return seuo
}

// SetNillableLastUpdatedBalance sets the "last_updated_balance" field if the given value is not nil.
func (seuo *SystemEWalletUpdateOne) SetNillableLastUpdatedBalance(t *time.Time) *SystemEWalletUpdateOne {
	if t != nil {
		seuo.SetLastUpdatedBalance(*t)
	}
	return seuo
}

// ClearLastUpdatedBalance clears the value of the "last_updated_balance" field.
func (seuo *SystemEWalletUpdateOne) ClearLastUpdatedBalance() *SystemEWalletUpdateOne {
	seuo.mutation.ClearLastUpdatedBalance()
	return seuo
}

// Mutation returns the SystemEWalletMutation object of the builder.
func (seuo *SystemEWalletUpdateOne) Mutation() *SystemEWalletMutation {
	return seuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (seuo *SystemEWalletUpdateOne) Select(field string, fields ...string) *SystemEWalletUpdateOne {
	seuo.fields = append([]string{field}, fields...)
	return seuo
}

// Save executes the query and returns the updated SystemEWallet entity.
func (seuo *SystemEWalletUpdateOne) Save(ctx context.Context) (*SystemEWallet, error) {
	var (
		err  error
		node *SystemEWallet
	)
	seuo.defaults()
	if len(seuo.hooks) == 0 {
		if err = seuo.check(); err != nil {
			return nil, err
		}
		node, err = seuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemEWalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = seuo.check(); err != nil {
				return nil, err
			}
			seuo.mutation = mutation
			node, err = seuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(seuo.hooks) - 1; i >= 0; i-- {
			if seuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = seuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, seuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (seuo *SystemEWalletUpdateOne) SaveX(ctx context.Context) *SystemEWallet {
	node, err := seuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (seuo *SystemEWalletUpdateOne) Exec(ctx context.Context) error {
	_, err := seuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seuo *SystemEWalletUpdateOne) ExecX(ctx context.Context) {
	if err := seuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (seuo *SystemEWalletUpdateOne) defaults() {
	if _, ok := seuo.mutation.UpdatedAt(); !ok {
		v := systemewallet.UpdateDefaultUpdatedAt()
		seuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (seuo *SystemEWalletUpdateOne) check() error {
	if v, ok := seuo.mutation.CreatedBy(); ok {
		if err := systemewallet.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.created_by": %w`, err)}
		}
	}
	if v, ok := seuo.mutation.UpdatedBy(); ok {
		if err := systemewallet.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.updated_by": %w`, err)}
		}
	}
	if v, ok := seuo.mutation.EWalletName(); ok {
		if err := systemewallet.EWalletNameValidator(v); err != nil {
			return &ValidationError{Name: "e_wallet_name", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.e_wallet_name": %w`, err)}
		}
	}
	if v, ok := seuo.mutation.Status(); ok {
		if err := systemewallet.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.status": %w`, err)}
		}
	}
	if v, ok := seuo.mutation.MerchantID(); ok {
		if err := systemewallet.MerchantIDValidator(v); err != nil {
			return &ValidationError{Name: "merchant_id", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.merchant_id": %w`, err)}
		}
	}
	if v, ok := seuo.mutation.AccountPhoneNumber(); ok {
		if err := systemewallet.AccountPhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_phone_number", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.account_phone_number": %w`, err)}
		}
	}
	if v, ok := seuo.mutation.AccountName(); ok {
		if err := systemewallet.AccountNameValidator(v); err != nil {
			return &ValidationError{Name: "account_name", err: fmt.Errorf(`ent: validator failed for field "SystemEWallet.account_name": %w`, err)}
		}
	}
	return nil
}

func (seuo *SystemEWalletUpdateOne) sqlSave(ctx context.Context) (_node *SystemEWallet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systemewallet.Table,
			Columns: systemewallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: systemewallet.FieldID,
			},
		},
	}
	id, ok := seuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SystemEWallet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := seuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemewallet.FieldID)
		for _, f := range fields {
			if !systemewallet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != systemewallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := seuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemewallet.FieldCreatedAt,
		})
	}
	if value, ok := seuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemewallet.FieldUpdatedAt,
		})
	}
	if value, ok := seuo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemewallet.FieldCreatedBy,
		})
	}
	if value, ok := seuo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemewallet.FieldUpdatedBy,
		})
	}
	if value, ok := seuo.mutation.EWalletName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemewallet.FieldEWalletName,
		})
	}
	if value, ok := seuo.mutation.AddedEWalletName(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemewallet.FieldEWalletName,
		})
	}
	if value, ok := seuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemewallet.FieldStatus,
		})
	}
	if value, ok := seuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemewallet.FieldStatus,
		})
	}
	if value, ok := seuo.mutation.MerchantID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemewallet.FieldMerchantID,
		})
	}
	if value, ok := seuo.mutation.AddedMerchantID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemewallet.FieldMerchantID,
		})
	}
	if value, ok := seuo.mutation.AccountPhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemewallet.FieldAccountPhoneNumber,
		})
	}
	if value, ok := seuo.mutation.AccountName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemewallet.FieldAccountName,
		})
	}
	if value, ok := seuo.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldBalance,
		})
	}
	if value, ok := seuo.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldBalance,
		})
	}
	if value, ok := seuo.mutation.DailyBalance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldDailyBalance,
		})
	}
	if value, ok := seuo.mutation.AddedDailyBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldDailyBalance,
		})
	}
	if value, ok := seuo.mutation.DailyBalanceLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldDailyBalanceLimit,
		})
	}
	if value, ok := seuo.mutation.AddedDailyBalanceLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: systemewallet.FieldDailyBalanceLimit,
		})
	}
	if value, ok := seuo.mutation.DailyUsedAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemewallet.FieldDailyUsedAmount,
		})
	}
	if value, ok := seuo.mutation.AddedDailyUsedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemewallet.FieldDailyUsedAmount,
		})
	}
	if value, ok := seuo.mutation.LastUpdatedBalance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemewallet.FieldLastUpdatedBalance,
		})
	}
	if seuo.mutation.LastUpdatedBalanceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: systemewallet.FieldLastUpdatedBalance,
		})
	}
	_node = &SystemEWallet{config: seuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, seuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemewallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
