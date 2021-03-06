// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymenttelcodetail"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

// PaymentTelcoDetailDelete is the builder for deleting a PaymentTelcoDetail entity.
type PaymentTelcoDetailDelete struct {
	config
	hooks    []Hook
	mutation *PaymentTelcoDetailMutation
}

// Where appends a list predicates to the PaymentTelcoDetailDelete builder.
func (ptdd *PaymentTelcoDetailDelete) Where(ps ...predicate.PaymentTelcoDetail) *PaymentTelcoDetailDelete {
	ptdd.mutation.Where(ps...)
	return ptdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptdd *PaymentTelcoDetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptdd.hooks) == 0 {
		affected, err = ptdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentTelcoDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptdd.mutation = mutation
			affected, err = ptdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptdd.hooks) - 1; i >= 0; i-- {
			if ptdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdd *PaymentTelcoDetailDelete) ExecX(ctx context.Context) int {
	n, err := ptdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptdd *PaymentTelcoDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: paymenttelcodetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: paymenttelcodetail.FieldID,
			},
		},
	}
	if ps := ptdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ptdd.driver, _spec)
}

// PaymentTelcoDetailDeleteOne is the builder for deleting a single PaymentTelcoDetail entity.
type PaymentTelcoDetailDeleteOne struct {
	ptdd *PaymentTelcoDetailDelete
}

// Exec executes the deletion query.
func (ptddo *PaymentTelcoDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := ptddo.ptdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{paymenttelcodetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptddo *PaymentTelcoDetailDeleteOne) ExecX(ctx context.Context) {
	ptddo.ptdd.ExecX(ctx)
}
