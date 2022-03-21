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
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/systemcryptohotwallet"
)

// SystemCryptoHotWalletUpdate is the builder for updating SystemCryptoHotWallet entities.
type SystemCryptoHotWalletUpdate struct {
	config
	hooks    []Hook
	mutation *SystemCryptoHotWalletMutation
}

// Where appends a list predicates to the SystemCryptoHotWalletUpdate builder.
func (schwu *SystemCryptoHotWalletUpdate) Where(ps ...predicate.SystemCryptoHotWallet) *SystemCryptoHotWalletUpdate {
	schwu.mutation.Where(ps...)
	return schwu
}

// SetCreatedAt sets the "created_at" field.
func (schwu *SystemCryptoHotWalletUpdate) SetCreatedAt(t time.Time) *SystemCryptoHotWalletUpdate {
	schwu.mutation.SetCreatedAt(t)
	return schwu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (schwu *SystemCryptoHotWalletUpdate) SetNillableCreatedAt(t *time.Time) *SystemCryptoHotWalletUpdate {
	if t != nil {
		schwu.SetCreatedAt(*t)
	}
	return schwu
}

// SetUpdatedAt sets the "updated_at" field.
func (schwu *SystemCryptoHotWalletUpdate) SetUpdatedAt(t time.Time) *SystemCryptoHotWalletUpdate {
	schwu.mutation.SetUpdatedAt(t)
	return schwu
}

// SetCreatedBy sets the "created_by" field.
func (schwu *SystemCryptoHotWalletUpdate) SetCreatedBy(s string) *SystemCryptoHotWalletUpdate {
	schwu.mutation.SetCreatedBy(s)
	return schwu
}

// SetUpdatedBy sets the "updated_by" field.
func (schwu *SystemCryptoHotWalletUpdate) SetUpdatedBy(s string) *SystemCryptoHotWalletUpdate {
	schwu.mutation.SetUpdatedBy(s)
	return schwu
}

// SetMerchantID sets the "merchant_id" field.
func (schwu *SystemCryptoHotWalletUpdate) SetMerchantID(i int64) *SystemCryptoHotWalletUpdate {
	schwu.mutation.ResetMerchantID()
	schwu.mutation.SetMerchantID(i)
	return schwu
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (schwu *SystemCryptoHotWalletUpdate) SetNillableMerchantID(i *int64) *SystemCryptoHotWalletUpdate {
	if i != nil {
		schwu.SetMerchantID(*i)
	}
	return schwu
}

// AddMerchantID adds i to the "merchant_id" field.
func (schwu *SystemCryptoHotWalletUpdate) AddMerchantID(i int64) *SystemCryptoHotWalletUpdate {
	schwu.mutation.AddMerchantID(i)
	return schwu
}

// SetCryptoType sets the "crypto_type" field.
func (schwu *SystemCryptoHotWalletUpdate) SetCryptoType(i int32) *SystemCryptoHotWalletUpdate {
	schwu.mutation.ResetCryptoType()
	schwu.mutation.SetCryptoType(i)
	return schwu
}

// SetNillableCryptoType sets the "crypto_type" field if the given value is not nil.
func (schwu *SystemCryptoHotWalletUpdate) SetNillableCryptoType(i *int32) *SystemCryptoHotWalletUpdate {
	if i != nil {
		schwu.SetCryptoType(*i)
	}
	return schwu
}

// AddCryptoType adds i to the "crypto_type" field.
func (schwu *SystemCryptoHotWalletUpdate) AddCryptoType(i int32) *SystemCryptoHotWalletUpdate {
	schwu.mutation.AddCryptoType(i)
	return schwu
}

// SetCryptoNetworkType sets the "crypto_network_type" field.
func (schwu *SystemCryptoHotWalletUpdate) SetCryptoNetworkType(i int32) *SystemCryptoHotWalletUpdate {
	schwu.mutation.ResetCryptoNetworkType()
	schwu.mutation.SetCryptoNetworkType(i)
	return schwu
}

// SetNillableCryptoNetworkType sets the "crypto_network_type" field if the given value is not nil.
func (schwu *SystemCryptoHotWalletUpdate) SetNillableCryptoNetworkType(i *int32) *SystemCryptoHotWalletUpdate {
	if i != nil {
		schwu.SetCryptoNetworkType(*i)
	}
	return schwu
}

// AddCryptoNetworkType adds i to the "crypto_network_type" field.
func (schwu *SystemCryptoHotWalletUpdate) AddCryptoNetworkType(i int32) *SystemCryptoHotWalletUpdate {
	schwu.mutation.AddCryptoNetworkType(i)
	return schwu
}

// SetAddress sets the "address" field.
func (schwu *SystemCryptoHotWalletUpdate) SetAddress(s string) *SystemCryptoHotWalletUpdate {
	schwu.mutation.SetAddress(s)
	return schwu
}

// SetTotalBalance sets the "total_balance" field.
func (schwu *SystemCryptoHotWalletUpdate) SetTotalBalance(f float64) *SystemCryptoHotWalletUpdate {
	schwu.mutation.ResetTotalBalance()
	schwu.mutation.SetTotalBalance(f)
	return schwu
}

// SetNillableTotalBalance sets the "total_balance" field if the given value is not nil.
func (schwu *SystemCryptoHotWalletUpdate) SetNillableTotalBalance(f *float64) *SystemCryptoHotWalletUpdate {
	if f != nil {
		schwu.SetTotalBalance(*f)
	}
	return schwu
}

// AddTotalBalance adds f to the "total_balance" field.
func (schwu *SystemCryptoHotWalletUpdate) AddTotalBalance(f float64) *SystemCryptoHotWalletUpdate {
	schwu.mutation.AddTotalBalance(f)
	return schwu
}

// SetBalance sets the "balance" field.
func (schwu *SystemCryptoHotWalletUpdate) SetBalance(f float64) *SystemCryptoHotWalletUpdate {
	schwu.mutation.ResetBalance()
	schwu.mutation.SetBalance(f)
	return schwu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (schwu *SystemCryptoHotWalletUpdate) SetNillableBalance(f *float64) *SystemCryptoHotWalletUpdate {
	if f != nil {
		schwu.SetBalance(*f)
	}
	return schwu
}

// AddBalance adds f to the "balance" field.
func (schwu *SystemCryptoHotWalletUpdate) AddBalance(f float64) *SystemCryptoHotWalletUpdate {
	schwu.mutation.AddBalance(f)
	return schwu
}

// SetStatus sets the "status" field.
func (schwu *SystemCryptoHotWalletUpdate) SetStatus(i int32) *SystemCryptoHotWalletUpdate {
	schwu.mutation.ResetStatus()
	schwu.mutation.SetStatus(i)
	return schwu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (schwu *SystemCryptoHotWalletUpdate) SetNillableStatus(i *int32) *SystemCryptoHotWalletUpdate {
	if i != nil {
		schwu.SetStatus(*i)
	}
	return schwu
}

// AddStatus adds i to the "status" field.
func (schwu *SystemCryptoHotWalletUpdate) AddStatus(i int32) *SystemCryptoHotWalletUpdate {
	schwu.mutation.AddStatus(i)
	return schwu
}

// Mutation returns the SystemCryptoHotWalletMutation object of the builder.
func (schwu *SystemCryptoHotWalletUpdate) Mutation() *SystemCryptoHotWalletMutation {
	return schwu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (schwu *SystemCryptoHotWalletUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	schwu.defaults()
	if len(schwu.hooks) == 0 {
		if err = schwu.check(); err != nil {
			return 0, err
		}
		affected, err = schwu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemCryptoHotWalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = schwu.check(); err != nil {
				return 0, err
			}
			schwu.mutation = mutation
			affected, err = schwu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(schwu.hooks) - 1; i >= 0; i-- {
			if schwu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = schwu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, schwu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (schwu *SystemCryptoHotWalletUpdate) SaveX(ctx context.Context) int {
	affected, err := schwu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (schwu *SystemCryptoHotWalletUpdate) Exec(ctx context.Context) error {
	_, err := schwu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (schwu *SystemCryptoHotWalletUpdate) ExecX(ctx context.Context) {
	if err := schwu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (schwu *SystemCryptoHotWalletUpdate) defaults() {
	if _, ok := schwu.mutation.UpdatedAt(); !ok {
		v := systemcryptohotwallet.UpdateDefaultUpdatedAt()
		schwu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (schwu *SystemCryptoHotWalletUpdate) check() error {
	if v, ok := schwu.mutation.CreatedBy(); ok {
		if err := systemcryptohotwallet.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.created_by": %w`, err)}
		}
	}
	if v, ok := schwu.mutation.UpdatedBy(); ok {
		if err := systemcryptohotwallet.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.updated_by": %w`, err)}
		}
	}
	if v, ok := schwu.mutation.MerchantID(); ok {
		if err := systemcryptohotwallet.MerchantIDValidator(v); err != nil {
			return &ValidationError{Name: "merchant_id", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.merchant_id": %w`, err)}
		}
	}
	if v, ok := schwu.mutation.CryptoType(); ok {
		if err := systemcryptohotwallet.CryptoTypeValidator(v); err != nil {
			return &ValidationError{Name: "crypto_type", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.crypto_type": %w`, err)}
		}
	}
	if v, ok := schwu.mutation.CryptoNetworkType(); ok {
		if err := systemcryptohotwallet.CryptoNetworkTypeValidator(v); err != nil {
			return &ValidationError{Name: "crypto_network_type", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.crypto_network_type": %w`, err)}
		}
	}
	if v, ok := schwu.mutation.Address(); ok {
		if err := systemcryptohotwallet.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.address": %w`, err)}
		}
	}
	if v, ok := schwu.mutation.Status(); ok {
		if err := systemcryptohotwallet.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.status": %w`, err)}
		}
	}
	return nil
}

func (schwu *SystemCryptoHotWalletUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systemcryptohotwallet.Table,
			Columns: systemcryptohotwallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: systemcryptohotwallet.FieldID,
			},
		},
	}
	if ps := schwu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := schwu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemcryptohotwallet.FieldCreatedAt,
		})
	}
	if value, ok := schwu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemcryptohotwallet.FieldUpdatedAt,
		})
	}
	if value, ok := schwu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemcryptohotwallet.FieldCreatedBy,
		})
	}
	if value, ok := schwu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemcryptohotwallet.FieldUpdatedBy,
		})
	}
	if value, ok := schwu.mutation.MerchantID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemcryptohotwallet.FieldMerchantID,
		})
	}
	if value, ok := schwu.mutation.AddedMerchantID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemcryptohotwallet.FieldMerchantID,
		})
	}
	if value, ok := schwu.mutation.CryptoType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldCryptoType,
		})
	}
	if value, ok := schwu.mutation.AddedCryptoType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldCryptoType,
		})
	}
	if value, ok := schwu.mutation.CryptoNetworkType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldCryptoNetworkType,
		})
	}
	if value, ok := schwu.mutation.AddedCryptoNetworkType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldCryptoNetworkType,
		})
	}
	if value, ok := schwu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemcryptohotwallet.FieldAddress,
		})
	}
	if value, ok := schwu.mutation.TotalBalance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemcryptohotwallet.FieldTotalBalance,
		})
	}
	if value, ok := schwu.mutation.AddedTotalBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemcryptohotwallet.FieldTotalBalance,
		})
	}
	if value, ok := schwu.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemcryptohotwallet.FieldBalance,
		})
	}
	if value, ok := schwu.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemcryptohotwallet.FieldBalance,
		})
	}
	if value, ok := schwu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldStatus,
		})
	}
	if value, ok := schwu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, schwu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemcryptohotwallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SystemCryptoHotWalletUpdateOne is the builder for updating a single SystemCryptoHotWallet entity.
type SystemCryptoHotWalletUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SystemCryptoHotWalletMutation
}

// SetCreatedAt sets the "created_at" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetCreatedAt(t time.Time) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.SetCreatedAt(t)
	return schwuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetNillableCreatedAt(t *time.Time) *SystemCryptoHotWalletUpdateOne {
	if t != nil {
		schwuo.SetCreatedAt(*t)
	}
	return schwuo
}

// SetUpdatedAt sets the "updated_at" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetUpdatedAt(t time.Time) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.SetUpdatedAt(t)
	return schwuo
}

// SetCreatedBy sets the "created_by" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetCreatedBy(s string) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.SetCreatedBy(s)
	return schwuo
}

// SetUpdatedBy sets the "updated_by" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetUpdatedBy(s string) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.SetUpdatedBy(s)
	return schwuo
}

// SetMerchantID sets the "merchant_id" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetMerchantID(i int64) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.ResetMerchantID()
	schwuo.mutation.SetMerchantID(i)
	return schwuo
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetNillableMerchantID(i *int64) *SystemCryptoHotWalletUpdateOne {
	if i != nil {
		schwuo.SetMerchantID(*i)
	}
	return schwuo
}

// AddMerchantID adds i to the "merchant_id" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) AddMerchantID(i int64) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.AddMerchantID(i)
	return schwuo
}

// SetCryptoType sets the "crypto_type" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetCryptoType(i int32) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.ResetCryptoType()
	schwuo.mutation.SetCryptoType(i)
	return schwuo
}

// SetNillableCryptoType sets the "crypto_type" field if the given value is not nil.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetNillableCryptoType(i *int32) *SystemCryptoHotWalletUpdateOne {
	if i != nil {
		schwuo.SetCryptoType(*i)
	}
	return schwuo
}

// AddCryptoType adds i to the "crypto_type" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) AddCryptoType(i int32) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.AddCryptoType(i)
	return schwuo
}

// SetCryptoNetworkType sets the "crypto_network_type" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetCryptoNetworkType(i int32) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.ResetCryptoNetworkType()
	schwuo.mutation.SetCryptoNetworkType(i)
	return schwuo
}

// SetNillableCryptoNetworkType sets the "crypto_network_type" field if the given value is not nil.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetNillableCryptoNetworkType(i *int32) *SystemCryptoHotWalletUpdateOne {
	if i != nil {
		schwuo.SetCryptoNetworkType(*i)
	}
	return schwuo
}

// AddCryptoNetworkType adds i to the "crypto_network_type" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) AddCryptoNetworkType(i int32) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.AddCryptoNetworkType(i)
	return schwuo
}

// SetAddress sets the "address" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetAddress(s string) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.SetAddress(s)
	return schwuo
}

// SetTotalBalance sets the "total_balance" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetTotalBalance(f float64) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.ResetTotalBalance()
	schwuo.mutation.SetTotalBalance(f)
	return schwuo
}

// SetNillableTotalBalance sets the "total_balance" field if the given value is not nil.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetNillableTotalBalance(f *float64) *SystemCryptoHotWalletUpdateOne {
	if f != nil {
		schwuo.SetTotalBalance(*f)
	}
	return schwuo
}

// AddTotalBalance adds f to the "total_balance" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) AddTotalBalance(f float64) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.AddTotalBalance(f)
	return schwuo
}

// SetBalance sets the "balance" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetBalance(f float64) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.ResetBalance()
	schwuo.mutation.SetBalance(f)
	return schwuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetNillableBalance(f *float64) *SystemCryptoHotWalletUpdateOne {
	if f != nil {
		schwuo.SetBalance(*f)
	}
	return schwuo
}

// AddBalance adds f to the "balance" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) AddBalance(f float64) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.AddBalance(f)
	return schwuo
}

// SetStatus sets the "status" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetStatus(i int32) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.ResetStatus()
	schwuo.mutation.SetStatus(i)
	return schwuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (schwuo *SystemCryptoHotWalletUpdateOne) SetNillableStatus(i *int32) *SystemCryptoHotWalletUpdateOne {
	if i != nil {
		schwuo.SetStatus(*i)
	}
	return schwuo
}

// AddStatus adds i to the "status" field.
func (schwuo *SystemCryptoHotWalletUpdateOne) AddStatus(i int32) *SystemCryptoHotWalletUpdateOne {
	schwuo.mutation.AddStatus(i)
	return schwuo
}

// Mutation returns the SystemCryptoHotWalletMutation object of the builder.
func (schwuo *SystemCryptoHotWalletUpdateOne) Mutation() *SystemCryptoHotWalletMutation {
	return schwuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (schwuo *SystemCryptoHotWalletUpdateOne) Select(field string, fields ...string) *SystemCryptoHotWalletUpdateOne {
	schwuo.fields = append([]string{field}, fields...)
	return schwuo
}

// Save executes the query and returns the updated SystemCryptoHotWallet entity.
func (schwuo *SystemCryptoHotWalletUpdateOne) Save(ctx context.Context) (*SystemCryptoHotWallet, error) {
	var (
		err  error
		node *SystemCryptoHotWallet
	)
	schwuo.defaults()
	if len(schwuo.hooks) == 0 {
		if err = schwuo.check(); err != nil {
			return nil, err
		}
		node, err = schwuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemCryptoHotWalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = schwuo.check(); err != nil {
				return nil, err
			}
			schwuo.mutation = mutation
			node, err = schwuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(schwuo.hooks) - 1; i >= 0; i-- {
			if schwuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = schwuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, schwuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (schwuo *SystemCryptoHotWalletUpdateOne) SaveX(ctx context.Context) *SystemCryptoHotWallet {
	node, err := schwuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (schwuo *SystemCryptoHotWalletUpdateOne) Exec(ctx context.Context) error {
	_, err := schwuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (schwuo *SystemCryptoHotWalletUpdateOne) ExecX(ctx context.Context) {
	if err := schwuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (schwuo *SystemCryptoHotWalletUpdateOne) defaults() {
	if _, ok := schwuo.mutation.UpdatedAt(); !ok {
		v := systemcryptohotwallet.UpdateDefaultUpdatedAt()
		schwuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (schwuo *SystemCryptoHotWalletUpdateOne) check() error {
	if v, ok := schwuo.mutation.CreatedBy(); ok {
		if err := systemcryptohotwallet.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.created_by": %w`, err)}
		}
	}
	if v, ok := schwuo.mutation.UpdatedBy(); ok {
		if err := systemcryptohotwallet.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.updated_by": %w`, err)}
		}
	}
	if v, ok := schwuo.mutation.MerchantID(); ok {
		if err := systemcryptohotwallet.MerchantIDValidator(v); err != nil {
			return &ValidationError{Name: "merchant_id", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.merchant_id": %w`, err)}
		}
	}
	if v, ok := schwuo.mutation.CryptoType(); ok {
		if err := systemcryptohotwallet.CryptoTypeValidator(v); err != nil {
			return &ValidationError{Name: "crypto_type", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.crypto_type": %w`, err)}
		}
	}
	if v, ok := schwuo.mutation.CryptoNetworkType(); ok {
		if err := systemcryptohotwallet.CryptoNetworkTypeValidator(v); err != nil {
			return &ValidationError{Name: "crypto_network_type", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.crypto_network_type": %w`, err)}
		}
	}
	if v, ok := schwuo.mutation.Address(); ok {
		if err := systemcryptohotwallet.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.address": %w`, err)}
		}
	}
	if v, ok := schwuo.mutation.Status(); ok {
		if err := systemcryptohotwallet.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SystemCryptoHotWallet.status": %w`, err)}
		}
	}
	return nil
}

func (schwuo *SystemCryptoHotWalletUpdateOne) sqlSave(ctx context.Context) (_node *SystemCryptoHotWallet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systemcryptohotwallet.Table,
			Columns: systemcryptohotwallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: systemcryptohotwallet.FieldID,
			},
		},
	}
	id, ok := schwuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SystemCryptoHotWallet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := schwuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemcryptohotwallet.FieldID)
		for _, f := range fields {
			if !systemcryptohotwallet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != systemcryptohotwallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := schwuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := schwuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemcryptohotwallet.FieldCreatedAt,
		})
	}
	if value, ok := schwuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemcryptohotwallet.FieldUpdatedAt,
		})
	}
	if value, ok := schwuo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemcryptohotwallet.FieldCreatedBy,
		})
	}
	if value, ok := schwuo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemcryptohotwallet.FieldUpdatedBy,
		})
	}
	if value, ok := schwuo.mutation.MerchantID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemcryptohotwallet.FieldMerchantID,
		})
	}
	if value, ok := schwuo.mutation.AddedMerchantID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systemcryptohotwallet.FieldMerchantID,
		})
	}
	if value, ok := schwuo.mutation.CryptoType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldCryptoType,
		})
	}
	if value, ok := schwuo.mutation.AddedCryptoType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldCryptoType,
		})
	}
	if value, ok := schwuo.mutation.CryptoNetworkType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldCryptoNetworkType,
		})
	}
	if value, ok := schwuo.mutation.AddedCryptoNetworkType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldCryptoNetworkType,
		})
	}
	if value, ok := schwuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemcryptohotwallet.FieldAddress,
		})
	}
	if value, ok := schwuo.mutation.TotalBalance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemcryptohotwallet.FieldTotalBalance,
		})
	}
	if value, ok := schwuo.mutation.AddedTotalBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemcryptohotwallet.FieldTotalBalance,
		})
	}
	if value, ok := schwuo.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemcryptohotwallet.FieldBalance,
		})
	}
	if value, ok := schwuo.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemcryptohotwallet.FieldBalance,
		})
	}
	if value, ok := schwuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldStatus,
		})
	}
	if value, ok := schwuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systemcryptohotwallet.FieldStatus,
		})
	}
	_node = &SystemCryptoHotWallet{config: schwuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, schwuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemcryptohotwallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
