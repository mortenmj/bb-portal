// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/blob"
)

// BlobCreate is the builder for creating a Blob entity.
type BlobCreate struct {
	config
	mutation *BlobMutation
	hooks    []Hook
}

// SetURI sets the "uri" field.
func (bc *BlobCreate) SetURI(s string) *BlobCreate {
	bc.mutation.SetURI(s)
	return bc
}

// SetSizeBytes sets the "size_bytes" field.
func (bc *BlobCreate) SetSizeBytes(i int64) *BlobCreate {
	bc.mutation.SetSizeBytes(i)
	return bc
}

// SetNillableSizeBytes sets the "size_bytes" field if the given value is not nil.
func (bc *BlobCreate) SetNillableSizeBytes(i *int64) *BlobCreate {
	if i != nil {
		bc.SetSizeBytes(*i)
	}
	return bc
}

// SetArchivingStatus sets the "archiving_status" field.
func (bc *BlobCreate) SetArchivingStatus(bs blob.ArchivingStatus) *BlobCreate {
	bc.mutation.SetArchivingStatus(bs)
	return bc
}

// SetNillableArchivingStatus sets the "archiving_status" field if the given value is not nil.
func (bc *BlobCreate) SetNillableArchivingStatus(bs *blob.ArchivingStatus) *BlobCreate {
	if bs != nil {
		bc.SetArchivingStatus(*bs)
	}
	return bc
}

// SetReason sets the "reason" field.
func (bc *BlobCreate) SetReason(s string) *BlobCreate {
	bc.mutation.SetReason(s)
	return bc
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (bc *BlobCreate) SetNillableReason(s *string) *BlobCreate {
	if s != nil {
		bc.SetReason(*s)
	}
	return bc
}

// SetArchiveURL sets the "archive_url" field.
func (bc *BlobCreate) SetArchiveURL(s string) *BlobCreate {
	bc.mutation.SetArchiveURL(s)
	return bc
}

// SetNillableArchiveURL sets the "archive_url" field if the given value is not nil.
func (bc *BlobCreate) SetNillableArchiveURL(s *string) *BlobCreate {
	if s != nil {
		bc.SetArchiveURL(*s)
	}
	return bc
}

// Mutation returns the BlobMutation object of the builder.
func (bc *BlobCreate) Mutation() *BlobMutation {
	return bc.mutation
}

// Save creates the Blob in the database.
func (bc *BlobCreate) Save(ctx context.Context) (*Blob, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlobCreate) SaveX(ctx context.Context) *Blob {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlobCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlobCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlobCreate) defaults() {
	if _, ok := bc.mutation.ArchivingStatus(); !ok {
		v := blob.DefaultArchivingStatus
		bc.mutation.SetArchivingStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlobCreate) check() error {
	if _, ok := bc.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "Blob.uri"`)}
	}
	if _, ok := bc.mutation.ArchivingStatus(); !ok {
		return &ValidationError{Name: "archiving_status", err: errors.New(`ent: missing required field "Blob.archiving_status"`)}
	}
	if v, ok := bc.mutation.ArchivingStatus(); ok {
		if err := blob.ArchivingStatusValidator(v); err != nil {
			return &ValidationError{Name: "archiving_status", err: fmt.Errorf(`ent: validator failed for field "Blob.archiving_status": %w`, err)}
		}
	}
	return nil
}

func (bc *BlobCreate) sqlSave(ctx context.Context) (*Blob, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlobCreate) createSpec() (*Blob, *sqlgraph.CreateSpec) {
	var (
		_node = &Blob{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(blob.Table, sqlgraph.NewFieldSpec(blob.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.URI(); ok {
		_spec.SetField(blob.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if value, ok := bc.mutation.SizeBytes(); ok {
		_spec.SetField(blob.FieldSizeBytes, field.TypeInt64, value)
		_node.SizeBytes = value
	}
	if value, ok := bc.mutation.ArchivingStatus(); ok {
		_spec.SetField(blob.FieldArchivingStatus, field.TypeEnum, value)
		_node.ArchivingStatus = value
	}
	if value, ok := bc.mutation.Reason(); ok {
		_spec.SetField(blob.FieldReason, field.TypeString, value)
		_node.Reason = value
	}
	if value, ok := bc.mutation.ArchiveURL(); ok {
		_spec.SetField(blob.FieldArchiveURL, field.TypeString, value)
		_node.ArchiveURL = value
	}
	return _node, _spec
}

// BlobCreateBulk is the builder for creating many Blob entities in bulk.
type BlobCreateBulk struct {
	config
	err      error
	builders []*BlobCreate
}

// Save creates the Blob entities in the database.
func (bcb *BlobCreateBulk) Save(ctx context.Context) ([]*Blob, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blob, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlobMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlobCreateBulk) SaveX(ctx context.Context) []*Blob {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlobCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlobCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
