// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/k0kishima/golang-realworld-example-app/ent/articletag"
	"github.com/k0kishima/golang-realworld-example-app/ent/predicate"
)

// ArticleTagUpdate is the builder for updating ArticleTag entities.
type ArticleTagUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleTagMutation
}

// Where appends a list predicates to the ArticleTagUpdate builder.
func (atu *ArticleTagUpdate) Where(ps ...predicate.ArticleTag) *ArticleTagUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// Mutation returns the ArticleTagMutation object of the builder.
func (atu *ArticleTagUpdate) Mutation() *ArticleTagMutation {
	return atu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *ArticleTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ArticleTagUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ArticleTagUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ArticleTagUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *ArticleTagUpdate) check() error {
	if _, ok := atu.mutation.ArticleID(); atu.mutation.ArticleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArticleTag.article"`)
	}
	if _, ok := atu.mutation.TagID(); atu.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArticleTag.tag"`)
	}
	return nil
}

func (atu *ArticleTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(articletag.Table, articletag.Columns, sqlgraph.NewFieldSpec(articletag.FieldID, field.TypeUUID))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{articletag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// ArticleTagUpdateOne is the builder for updating a single ArticleTag entity.
type ArticleTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleTagMutation
}

// Mutation returns the ArticleTagMutation object of the builder.
func (atuo *ArticleTagUpdateOne) Mutation() *ArticleTagMutation {
	return atuo.mutation
}

// Where appends a list predicates to the ArticleTagUpdate builder.
func (atuo *ArticleTagUpdateOne) Where(ps ...predicate.ArticleTag) *ArticleTagUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *ArticleTagUpdateOne) Select(field string, fields ...string) *ArticleTagUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated ArticleTag entity.
func (atuo *ArticleTagUpdateOne) Save(ctx context.Context) (*ArticleTag, error) {
	return withHooks(ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ArticleTagUpdateOne) SaveX(ctx context.Context) *ArticleTag {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *ArticleTagUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ArticleTagUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *ArticleTagUpdateOne) check() error {
	if _, ok := atuo.mutation.ArticleID(); atuo.mutation.ArticleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArticleTag.article"`)
	}
	if _, ok := atuo.mutation.TagID(); atuo.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArticleTag.tag"`)
	}
	return nil
}

func (atuo *ArticleTagUpdateOne) sqlSave(ctx context.Context) (_node *ArticleTag, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(articletag.Table, articletag.Columns, sqlgraph.NewFieldSpec(articletag.FieldID, field.TypeUUID))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArticleTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, articletag.FieldID)
		for _, f := range fields {
			if !articletag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != articletag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &ArticleTag{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{articletag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}