// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/paymentbankingdetail"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

// PaymentBankingDetailDelete is the builder for deleting a PaymentBankingDetail entity.
type PaymentBankingDetailDelete struct {
	config
	hooks    []Hook
	mutation *PaymentBankingDetailMutation
}

// Where appends a list predicates to the PaymentBankingDetailDelete builder.
func (pbdd *PaymentBankingDetailDelete) Where(ps ...predicate.PaymentBankingDetail) *PaymentBankingDetailDelete {
	pbdd.mutation.Where(ps...)
	return pbdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pbdd *PaymentBankingDetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pbdd.hooks) == 0 {
		affected, err = pbdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentBankingDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pbdd.mutation = mutation
			affected, err = pbdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pbdd.hooks) - 1; i >= 0; i-- {
			if pbdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pbdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pbdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbdd *PaymentBankingDetailDelete) ExecX(ctx context.Context) int {
	n, err := pbdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pbdd *PaymentBankingDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: paymentbankingdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: paymentbankingdetail.FieldID,
			},
		},
	}
	if ps := pbdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pbdd.driver, _spec)
}

// PaymentBankingDetailDeleteOne is the builder for deleting a single PaymentBankingDetail entity.
type PaymentBankingDetailDeleteOne struct {
	pbdd *PaymentBankingDetailDelete
}

// Exec executes the deletion query.
func (pbddo *PaymentBankingDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := pbddo.pbdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{paymentbankingdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pbddo *PaymentBankingDetailDeleteOne) ExecX(ctx context.Context) {
	pbddo.pbdd.ExecX(ctx)
}
