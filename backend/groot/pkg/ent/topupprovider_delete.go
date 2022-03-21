// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent/predicate"
	"gitlab.com/mcuc/monorepo/backend/groot/pkg/ent/topupprovider"
)

// TopUpProviderDelete is the builder for deleting a TopUpProvider entity.
type TopUpProviderDelete struct {
	config
	hooks    []Hook
	mutation *TopUpProviderMutation
}

// Where appends a list predicates to the TopUpProviderDelete builder.
func (tupd *TopUpProviderDelete) Where(ps ...predicate.TopUpProvider) *TopUpProviderDelete {
	tupd.mutation.Where(ps...)
	return tupd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tupd *TopUpProviderDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tupd.hooks) == 0 {
		affected, err = tupd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopUpProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tupd.mutation = mutation
			affected, err = tupd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tupd.hooks) - 1; i >= 0; i-- {
			if tupd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tupd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tupd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tupd *TopUpProviderDelete) ExecX(ctx context.Context) int {
	n, err := tupd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tupd *TopUpProviderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: topupprovider.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: topupprovider.FieldID,
			},
		},
	}
	if ps := tupd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tupd.driver, _spec)
}

// TopUpProviderDeleteOne is the builder for deleting a single TopUpProvider entity.
type TopUpProviderDeleteOne struct {
	tupd *TopUpProviderDelete
}

// Exec executes the deletion query.
func (tupdo *TopUpProviderDeleteOne) Exec(ctx context.Context) error {
	n, err := tupdo.tupd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{topupprovider.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tupdo *TopUpProviderDeleteOne) ExecX(ctx context.Context) {
	tupdo.tupd.ExecX(ctx)
}
