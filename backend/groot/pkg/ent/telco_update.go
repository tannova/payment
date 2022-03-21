// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent/predicate"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent/telco"
)

// TelcoUpdate is the builder for updating Telco entities.
type TelcoUpdate struct {
	config
	hooks    []Hook
	mutation *TelcoMutation
}

// Where appends a list predicates to the TelcoUpdate builder.
func (tu *TelcoUpdate) Where(ps ...predicate.Telco) *TelcoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTelcoName sets the "telco_name" field.
func (tu *TelcoUpdate) SetTelcoName(i int64) *TelcoUpdate {
	tu.mutation.ResetTelcoName()
	tu.mutation.SetTelcoName(i)
	return tu
}

// SetNillableTelcoName sets the "telco_name" field if the given value is not nil.
func (tu *TelcoUpdate) SetNillableTelcoName(i *int64) *TelcoUpdate {
	if i != nil {
		tu.SetTelcoName(*i)
	}
	return tu
}

// AddTelcoName adds i to the "telco_name" field.
func (tu *TelcoUpdate) AddTelcoName(i int64) *TelcoUpdate {
	tu.mutation.AddTelcoName(i)
	return tu
}

// SetCode sets the "code" field.
func (tu *TelcoUpdate) SetCode(s string) *TelcoUpdate {
	tu.mutation.SetCode(s)
	return tu
}

// SetSerial sets the "serial" field.
func (tu *TelcoUpdate) SetSerial(s string) *TelcoUpdate {
	tu.mutation.SetSerial(s)
	return tu
}

// SetAmount sets the "amount" field.
func (tu *TelcoUpdate) SetAmount(u uint64) *TelcoUpdate {
	tu.mutation.ResetAmount()
	tu.mutation.SetAmount(u)
	return tu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tu *TelcoUpdate) SetNillableAmount(u *uint64) *TelcoUpdate {
	if u != nil {
		tu.SetAmount(*u)
	}
	return tu
}

// AddAmount adds u to the "amount" field.
func (tu *TelcoUpdate) AddAmount(u uint64) *TelcoUpdate {
	tu.mutation.AddAmount(u)
	return tu
}

// SetUsed sets the "used" field.
func (tu *TelcoUpdate) SetUsed(b bool) *TelcoUpdate {
	tu.mutation.SetUsed(b)
	return tu
}

// SetNillableUsed sets the "used" field if the given value is not nil.
func (tu *TelcoUpdate) SetNillableUsed(b *bool) *TelcoUpdate {
	if b != nil {
		tu.SetUsed(*b)
	}
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TelcoUpdate) SetCreatedAt(t time.Time) *TelcoUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TelcoUpdate) SetNillableCreatedAt(t *time.Time) *TelcoUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TelcoUpdate) SetUpdatedAt(t time.Time) *TelcoUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetCreatedBy sets the "created_by" field.
func (tu *TelcoUpdate) SetCreatedBy(s string) *TelcoUpdate {
	tu.mutation.SetCreatedBy(s)
	return tu
}

// SetUpdatedBy sets the "updated_by" field.
func (tu *TelcoUpdate) SetUpdatedBy(s string) *TelcoUpdate {
	tu.mutation.SetUpdatedBy(s)
	return tu
}

// Mutation returns the TelcoMutation object of the builder.
func (tu *TelcoUpdate) Mutation() *TelcoMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TelcoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TelcoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TelcoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TelcoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TelcoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TelcoUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := telco.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TelcoUpdate) check() error {
	if v, ok := tu.mutation.TelcoName(); ok {
		if err := telco.TelcoNameValidator(v); err != nil {
			return &ValidationError{Name: "telco_name", err: fmt.Errorf("ent: validator failed for field \"telco_name\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Code(); ok {
		if err := telco.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Serial(); ok {
		if err := telco.SerialValidator(v); err != nil {
			return &ValidationError{Name: "serial", err: fmt.Errorf("ent: validator failed for field \"serial\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Amount(); ok {
		if err := telco.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if v, ok := tu.mutation.CreatedBy(); ok {
		if err := telco.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf("ent: validator failed for field \"created_by\": %w", err)}
		}
	}
	if v, ok := tu.mutation.UpdatedBy(); ok {
		if err := telco.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf("ent: validator failed for field \"updated_by\": %w", err)}
		}
	}
	return nil
}

func (tu *TelcoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   telco.Table,
			Columns: telco.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: telco.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.TelcoName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telco.FieldTelcoName,
		})
	}
	if value, ok := tu.mutation.AddedTelcoName(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telco.FieldTelcoName,
		})
	}
	if value, ok := tu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldCode,
		})
	}
	if value, ok := tu.mutation.Serial(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldSerial,
		})
	}
	if value, ok := tu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: telco.FieldAmount,
		})
	}
	if value, ok := tu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: telco.FieldAmount,
		})
	}
	if value, ok := tu.mutation.Used(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: telco.FieldUsed,
		})
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telco.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telco.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldCreatedBy,
		})
	}
	if value, ok := tu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldUpdatedBy,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telco.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TelcoUpdateOne is the builder for updating a single Telco entity.
type TelcoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TelcoMutation
}

// SetTelcoName sets the "telco_name" field.
func (tuo *TelcoUpdateOne) SetTelcoName(i int64) *TelcoUpdateOne {
	tuo.mutation.ResetTelcoName()
	tuo.mutation.SetTelcoName(i)
	return tuo
}

// SetNillableTelcoName sets the "telco_name" field if the given value is not nil.
func (tuo *TelcoUpdateOne) SetNillableTelcoName(i *int64) *TelcoUpdateOne {
	if i != nil {
		tuo.SetTelcoName(*i)
	}
	return tuo
}

// AddTelcoName adds i to the "telco_name" field.
func (tuo *TelcoUpdateOne) AddTelcoName(i int64) *TelcoUpdateOne {
	tuo.mutation.AddTelcoName(i)
	return tuo
}

// SetCode sets the "code" field.
func (tuo *TelcoUpdateOne) SetCode(s string) *TelcoUpdateOne {
	tuo.mutation.SetCode(s)
	return tuo
}

// SetSerial sets the "serial" field.
func (tuo *TelcoUpdateOne) SetSerial(s string) *TelcoUpdateOne {
	tuo.mutation.SetSerial(s)
	return tuo
}

// SetAmount sets the "amount" field.
func (tuo *TelcoUpdateOne) SetAmount(u uint64) *TelcoUpdateOne {
	tuo.mutation.ResetAmount()
	tuo.mutation.SetAmount(u)
	return tuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tuo *TelcoUpdateOne) SetNillableAmount(u *uint64) *TelcoUpdateOne {
	if u != nil {
		tuo.SetAmount(*u)
	}
	return tuo
}

// AddAmount adds u to the "amount" field.
func (tuo *TelcoUpdateOne) AddAmount(u uint64) *TelcoUpdateOne {
	tuo.mutation.AddAmount(u)
	return tuo
}

// SetUsed sets the "used" field.
func (tuo *TelcoUpdateOne) SetUsed(b bool) *TelcoUpdateOne {
	tuo.mutation.SetUsed(b)
	return tuo
}

// SetNillableUsed sets the "used" field if the given value is not nil.
func (tuo *TelcoUpdateOne) SetNillableUsed(b *bool) *TelcoUpdateOne {
	if b != nil {
		tuo.SetUsed(*b)
	}
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TelcoUpdateOne) SetCreatedAt(t time.Time) *TelcoUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TelcoUpdateOne) SetNillableCreatedAt(t *time.Time) *TelcoUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TelcoUpdateOne) SetUpdatedAt(t time.Time) *TelcoUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetCreatedBy sets the "created_by" field.
func (tuo *TelcoUpdateOne) SetCreatedBy(s string) *TelcoUpdateOne {
	tuo.mutation.SetCreatedBy(s)
	return tuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tuo *TelcoUpdateOne) SetUpdatedBy(s string) *TelcoUpdateOne {
	tuo.mutation.SetUpdatedBy(s)
	return tuo
}

// Mutation returns the TelcoMutation object of the builder.
func (tuo *TelcoUpdateOne) Mutation() *TelcoMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TelcoUpdateOne) Select(field string, fields ...string) *TelcoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Telco entity.
func (tuo *TelcoUpdateOne) Save(ctx context.Context) (*Telco, error) {
	var (
		err  error
		node *Telco
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TelcoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TelcoUpdateOne) SaveX(ctx context.Context) *Telco {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TelcoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TelcoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TelcoUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := telco.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TelcoUpdateOne) check() error {
	if v, ok := tuo.mutation.TelcoName(); ok {
		if err := telco.TelcoNameValidator(v); err != nil {
			return &ValidationError{Name: "telco_name", err: fmt.Errorf("ent: validator failed for field \"telco_name\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Code(); ok {
		if err := telco.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Serial(); ok {
		if err := telco.SerialValidator(v); err != nil {
			return &ValidationError{Name: "serial", err: fmt.Errorf("ent: validator failed for field \"serial\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Amount(); ok {
		if err := telco.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.CreatedBy(); ok {
		if err := telco.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf("ent: validator failed for field \"created_by\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.UpdatedBy(); ok {
		if err := telco.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf("ent: validator failed for field \"updated_by\": %w", err)}
		}
	}
	return nil
}

func (tuo *TelcoUpdateOne) sqlSave(ctx context.Context) (_node *Telco, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   telco.Table,
			Columns: telco.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: telco.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Telco.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, telco.FieldID)
		for _, f := range fields {
			if !telco.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != telco.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.TelcoName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telco.FieldTelcoName,
		})
	}
	if value, ok := tuo.mutation.AddedTelcoName(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telco.FieldTelcoName,
		})
	}
	if value, ok := tuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldCode,
		})
	}
	if value, ok := tuo.mutation.Serial(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldSerial,
		})
	}
	if value, ok := tuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: telco.FieldAmount,
		})
	}
	if value, ok := tuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: telco.FieldAmount,
		})
	}
	if value, ok := tuo.mutation.Used(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: telco.FieldUsed,
		})
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telco.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telco.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldCreatedBy,
		})
	}
	if value, ok := tuo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telco.FieldUpdatedBy,
		})
	}
	_node = &Telco{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telco.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}