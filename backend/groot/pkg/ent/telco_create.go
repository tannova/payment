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
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent/telco"
)

// TelcoCreate is the builder for creating a Telco entity.
type TelcoCreate struct {
	config
	mutation *TelcoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTelcoName sets the "telco_name" field.
func (tc *TelcoCreate) SetTelcoName(i int64) *TelcoCreate {
	tc.mutation.SetTelcoName(i)
	return tc
}

// SetNillableTelcoName sets the "telco_name" field if the given value is not nil.
func (tc *TelcoCreate) SetNillableTelcoName(i *int64) *TelcoCreate {
	if i != nil {
		tc.SetTelcoName(*i)
	}
	return tc
}

// SetCode sets the "code" field.
func (tc *TelcoCreate) SetCode(s string) *TelcoCreate {
	tc.mutation.SetCode(s)
	return tc
}

// SetSerial sets the "serial" field.
func (tc *TelcoCreate) SetSerial(s string) *TelcoCreate {
	tc.mutation.SetSerial(s)
	return tc
}

// SetAmount sets the "amount" field.
func (tc *TelcoCreate) SetAmount(u uint64) *TelcoCreate {
	tc.mutation.SetAmount(u)
	return tc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tc *TelcoCreate) SetNillableAmount(u *uint64) *TelcoCreate {
	if u != nil {
		tc.SetAmount(*u)
	}
	return tc
}

// SetUsed sets the "used" field.
func (tc *TelcoCreate) SetUsed(b bool) *TelcoCreate {
	tc.mutation.SetUsed(b)
	return tc
}

// SetNillableUsed sets the "used" field if the given value is not nil.
func (tc *TelcoCreate) SetNillableUsed(b *bool) *TelcoCreate {
	if b != nil {
		tc.SetUsed(*b)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TelcoCreate) SetCreatedAt(t time.Time) *TelcoCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TelcoCreate) SetNillableCreatedAt(t *time.Time) *TelcoCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TelcoCreate) SetUpdatedAt(t time.Time) *TelcoCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TelcoCreate) SetNillableUpdatedAt(t *time.Time) *TelcoCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetCreatedBy sets the "created_by" field.
func (tc *TelcoCreate) SetCreatedBy(s string) *TelcoCreate {
	tc.mutation.SetCreatedBy(s)
	return tc
}

// SetUpdatedBy sets the "updated_by" field.
func (tc *TelcoCreate) SetUpdatedBy(s string) *TelcoCreate {
	tc.mutation.SetUpdatedBy(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TelcoCreate) SetID(i int64) *TelcoCreate {
	tc.mutation.SetID(i)
	return tc
}

// Mutation returns the TelcoMutation object of the builder.
func (tc *TelcoCreate) Mutation() *TelcoMutation {
	return tc.mutation
}

// Save creates the Telco in the database.
func (tc *TelcoCreate) Save(ctx context.Context) (*Telco, error) {
	var (
		err  error
		node *Telco
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TelcoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TelcoCreate) SaveX(ctx context.Context) *Telco {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TelcoCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TelcoCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TelcoCreate) defaults() {
	if _, ok := tc.mutation.TelcoName(); !ok {
		v := telco.DefaultTelcoName
		tc.mutation.SetTelcoName(v)
	}
	if _, ok := tc.mutation.Amount(); !ok {
		v := telco.DefaultAmount
		tc.mutation.SetAmount(v)
	}
	if _, ok := tc.mutation.Used(); !ok {
		v := telco.DefaultUsed
		tc.mutation.SetUsed(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := telco.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := telco.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TelcoCreate) check() error {
	if _, ok := tc.mutation.TelcoName(); !ok {
		return &ValidationError{Name: "telco_name", err: errors.New(`ent: missing required field "telco_name"`)}
	}
	if v, ok := tc.mutation.TelcoName(); ok {
		if err := telco.TelcoNameValidator(v); err != nil {
			return &ValidationError{Name: "telco_name", err: fmt.Errorf(`ent: validator failed for field "telco_name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "code"`)}
	}
	if v, ok := tc.mutation.Code(); ok {
		if err := telco.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "code": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Serial(); !ok {
		return &ValidationError{Name: "serial", err: errors.New(`ent: missing required field "serial"`)}
	}
	if v, ok := tc.mutation.Serial(); ok {
		if err := telco.SerialValidator(v); err != nil {
			return &ValidationError{Name: "serial", err: fmt.Errorf(`ent: validator failed for field "serial": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if v, ok := tc.mutation.Amount(); ok {
		if err := telco.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "amount": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Used(); !ok {
		return &ValidationError{Name: "used", err: errors.New(`ent: missing required field "used"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := tc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "created_by"`)}
	}
	if v, ok := tc.mutation.CreatedBy(); ok {
		if err := telco.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "created_by": %w`, err)}
		}
	}
	if _, ok := tc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "updated_by"`)}
	}
	if v, ok := tc.mutation.UpdatedBy(); ok {
		if err := telco.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "updated_by": %w`, err)}
		}
	}
	return nil
}

func (tc *TelcoCreate) sqlSave(ctx context.Context) (*Telco, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (tc *TelcoCreate) createSpec() (*Telco, *sqlgraph.CreateSpec) {
	var (
		_node = &Telco{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: telco.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: telco.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.TelcoName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telco.FieldTelcoName,
		})
		_node.TelcoName = value
	}
	if value, ok := tc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := tc.mutation.Serial(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldSerial,
		})
		_node.Serial = value
	}
	if value, ok := tc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: telco.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := tc.mutation.Used(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: telco.FieldUsed,
		})
		_node.Used = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telco.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telco.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := tc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Telco.Create().
//		SetTelcoName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TelcoUpsert) {
//			SetTelcoName(v+v).
//		}).
//		Exec(ctx)
//
func (tc *TelcoCreate) OnConflict(opts ...sql.ConflictOption) *TelcoUpsertOne {
	tc.conflict = opts
	return &TelcoUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Telco.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tc *TelcoCreate) OnConflictColumns(columns ...string) *TelcoUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TelcoUpsertOne{
		create: tc,
	}
}

type (
	// TelcoUpsertOne is the builder for "upsert"-ing
	//  one Telco node.
	TelcoUpsertOne struct {
		create *TelcoCreate
	}

	// TelcoUpsert is the "OnConflict" setter.
	TelcoUpsert struct {
		*sql.UpdateSet
	}
)

// SetTelcoName sets the "telco_name" field.
func (u *TelcoUpsert) SetTelcoName(v int64) *TelcoUpsert {
	u.Set(telco.FieldTelcoName, v)
	return u
}

// UpdateTelcoName sets the "telco_name" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateTelcoName() *TelcoUpsert {
	u.SetExcluded(telco.FieldTelcoName)
	return u
}

// SetCode sets the "code" field.
func (u *TelcoUpsert) SetCode(v string) *TelcoUpsert {
	u.Set(telco.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateCode() *TelcoUpsert {
	u.SetExcluded(telco.FieldCode)
	return u
}

// SetSerial sets the "serial" field.
func (u *TelcoUpsert) SetSerial(v string) *TelcoUpsert {
	u.Set(telco.FieldSerial, v)
	return u
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateSerial() *TelcoUpsert {
	u.SetExcluded(telco.FieldSerial)
	return u
}

// SetAmount sets the "amount" field.
func (u *TelcoUpsert) SetAmount(v uint64) *TelcoUpsert {
	u.Set(telco.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateAmount() *TelcoUpsert {
	u.SetExcluded(telco.FieldAmount)
	return u
}

// SetUsed sets the "used" field.
func (u *TelcoUpsert) SetUsed(v bool) *TelcoUpsert {
	u.Set(telco.FieldUsed, v)
	return u
}

// UpdateUsed sets the "used" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateUsed() *TelcoUpsert {
	u.SetExcluded(telco.FieldUsed)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TelcoUpsert) SetCreatedAt(v time.Time) *TelcoUpsert {
	u.Set(telco.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateCreatedAt() *TelcoUpsert {
	u.SetExcluded(telco.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TelcoUpsert) SetUpdatedAt(v time.Time) *TelcoUpsert {
	u.Set(telco.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateUpdatedAt() *TelcoUpsert {
	u.SetExcluded(telco.FieldUpdatedAt)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *TelcoUpsert) SetCreatedBy(v string) *TelcoUpsert {
	u.Set(telco.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateCreatedBy() *TelcoUpsert {
	u.SetExcluded(telco.FieldCreatedBy)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *TelcoUpsert) SetUpdatedBy(v string) *TelcoUpsert {
	u.Set(telco.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *TelcoUpsert) UpdateUpdatedBy() *TelcoUpsert {
	u.SetExcluded(telco.FieldUpdatedBy)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Telco.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(telco.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TelcoUpsertOne) UpdateNewValues() *TelcoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(telco.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Telco.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TelcoUpsertOne) Ignore() *TelcoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TelcoUpsertOne) DoNothing() *TelcoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TelcoCreate.OnConflict
// documentation for more info.
func (u *TelcoUpsertOne) Update(set func(*TelcoUpsert)) *TelcoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TelcoUpsert{UpdateSet: update})
	}))
	return u
}

// SetTelcoName sets the "telco_name" field.
func (u *TelcoUpsertOne) SetTelcoName(v int64) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetTelcoName(v)
	})
}

// UpdateTelcoName sets the "telco_name" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateTelcoName() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateTelcoName()
	})
}

// SetCode sets the "code" field.
func (u *TelcoUpsertOne) SetCode(v string) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateCode() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateCode()
	})
}

// SetSerial sets the "serial" field.
func (u *TelcoUpsertOne) SetSerial(v string) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetSerial(v)
	})
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateSerial() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateSerial()
	})
}

// SetAmount sets the "amount" field.
func (u *TelcoUpsertOne) SetAmount(v uint64) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateAmount() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateAmount()
	})
}

// SetUsed sets the "used" field.
func (u *TelcoUpsertOne) SetUsed(v bool) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetUsed(v)
	})
}

// UpdateUsed sets the "used" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateUsed() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateUsed()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TelcoUpsertOne) SetCreatedAt(v time.Time) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateCreatedAt() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TelcoUpsertOne) SetUpdatedAt(v time.Time) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateUpdatedAt() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *TelcoUpsertOne) SetCreatedBy(v string) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateCreatedBy() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *TelcoUpsertOne) SetUpdatedBy(v string) *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *TelcoUpsertOne) UpdateUpdatedBy() *TelcoUpsertOne {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateUpdatedBy()
	})
}

// Exec executes the query.
func (u *TelcoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TelcoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TelcoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TelcoUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TelcoUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TelcoCreateBulk is the builder for creating many Telco entities in bulk.
type TelcoCreateBulk struct {
	config
	builders []*TelcoCreate
	conflict []sql.ConflictOption
}

// Save creates the Telco entities in the database.
func (tcb *TelcoCreateBulk) Save(ctx context.Context) ([]*Telco, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Telco, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TelcoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TelcoCreateBulk) SaveX(ctx context.Context) []*Telco {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TelcoCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TelcoCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Telco.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TelcoUpsert) {
//			SetTelcoName(v+v).
//		}).
//		Exec(ctx)
//
func (tcb *TelcoCreateBulk) OnConflict(opts ...sql.ConflictOption) *TelcoUpsertBulk {
	tcb.conflict = opts
	return &TelcoUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Telco.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tcb *TelcoCreateBulk) OnConflictColumns(columns ...string) *TelcoUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TelcoUpsertBulk{
		create: tcb,
	}
}

// TelcoUpsertBulk is the builder for "upsert"-ing
// a bulk of Telco nodes.
type TelcoUpsertBulk struct {
	create *TelcoCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Telco.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(telco.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TelcoUpsertBulk) UpdateNewValues() *TelcoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(telco.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Telco.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TelcoUpsertBulk) Ignore() *TelcoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TelcoUpsertBulk) DoNothing() *TelcoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TelcoCreateBulk.OnConflict
// documentation for more info.
func (u *TelcoUpsertBulk) Update(set func(*TelcoUpsert)) *TelcoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TelcoUpsert{UpdateSet: update})
	}))
	return u
}

// SetTelcoName sets the "telco_name" field.
func (u *TelcoUpsertBulk) SetTelcoName(v int64) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetTelcoName(v)
	})
}

// UpdateTelcoName sets the "telco_name" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateTelcoName() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateTelcoName()
	})
}

// SetCode sets the "code" field.
func (u *TelcoUpsertBulk) SetCode(v string) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateCode() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateCode()
	})
}

// SetSerial sets the "serial" field.
func (u *TelcoUpsertBulk) SetSerial(v string) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetSerial(v)
	})
}

// UpdateSerial sets the "serial" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateSerial() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateSerial()
	})
}

// SetAmount sets the "amount" field.
func (u *TelcoUpsertBulk) SetAmount(v uint64) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateAmount() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateAmount()
	})
}

// SetUsed sets the "used" field.
func (u *TelcoUpsertBulk) SetUsed(v bool) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetUsed(v)
	})
}

// UpdateUsed sets the "used" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateUsed() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateUsed()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TelcoUpsertBulk) SetCreatedAt(v time.Time) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateCreatedAt() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TelcoUpsertBulk) SetUpdatedAt(v time.Time) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateUpdatedAt() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *TelcoUpsertBulk) SetCreatedBy(v string) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateCreatedBy() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *TelcoUpsertBulk) SetUpdatedBy(v string) *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *TelcoUpsertBulk) UpdateUpdatedBy() *TelcoUpsertBulk {
	return u.Update(func(s *TelcoUpsert) {
		s.UpdateUpdatedBy()
	})
}

// Exec executes the query.
func (u *TelcoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TelcoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TelcoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TelcoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}