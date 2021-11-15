// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"rpcClient/app/device/service/ent/device"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	mutation *DeviceMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DeviceCreate) SetName(s string) *DeviceCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableName(s *string) *DeviceCreate {
	if s != nil {
		dc.SetName(*s)
	}
	return dc
}

// SetHardwareKey sets the "hardwareKey" field.
func (dc *DeviceCreate) SetHardwareKey(s string) *DeviceCreate {
	dc.mutation.SetHardwareKey(s)
	return dc
}

// SetNillableHardwareKey sets the "hardwareKey" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableHardwareKey(s *string) *DeviceCreate {
	if s != nil {
		dc.SetHardwareKey(*s)
	}
	return dc
}

// SetDefaultLayoutId sets the "defaultLayoutId" field.
func (dc *DeviceCreate) SetDefaultLayoutId(i int32) *DeviceCreate {
	dc.mutation.SetDefaultLayoutId(i)
	return dc
}

// SetNillableDefaultLayoutId sets the "defaultLayoutId" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableDefaultLayoutId(i *int32) *DeviceCreate {
	if i != nil {
		dc.SetDefaultLayoutId(*i)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DeviceCreate) SetStatus(i int32) *DeviceCreate {
	dc.mutation.SetStatus(i)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableStatus(i *int32) *DeviceCreate {
	if i != nil {
		dc.SetStatus(*i)
	}
	return dc
}

// SetStoreNumber sets the "storeNumber" field.
func (dc *DeviceCreate) SetStoreNumber(s string) *DeviceCreate {
	dc.mutation.SetStoreNumber(s)
	return dc
}

// SetNillableStoreNumber sets the "storeNumber" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableStoreNumber(s *string) *DeviceCreate {
	if s != nil {
		dc.SetStoreNumber(*s)
	}
	return dc
}

// SetCreatedAt sets the "createdAt" field.
func (dc *DeviceCreate) SetCreatedAt(t time.Time) *DeviceCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableCreatedAt(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// Mutation returns the DeviceMutation object of the builder.
func (dc *DeviceCreate) Mutation() *DeviceMutation {
	return dc.mutation
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	var (
		err  error
		node *Device
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeviceCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeviceCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeviceCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := device.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeviceCreate) check() error {
	return nil
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DeviceCreate) createSpec() (*Device, *sqlgraph.CreateSpec) {
	var (
		_node = &Device{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: device.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldName,
		})
		_node.Name = &value
	}
	if value, ok := dc.mutation.HardwareKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldHardwareKey,
		})
		_node.HardwareKey = &value
	}
	if value, ok := dc.mutation.DefaultLayoutId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: device.FieldDefaultLayoutId,
		})
		_node.DefaultLayoutId = &value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: device.FieldStatus,
		})
		_node.Status = &value
	}
	if value, ok := dc.mutation.StoreNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldStoreNumber,
		})
		_node.StoreNumber = &value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	return _node, _spec
}

// DeviceCreateBulk is the builder for creating many Device entities in bulk.
type DeviceCreateBulk struct {
	config
	builders []*DeviceCreate
}

// Save creates the Device entities in the database.
func (dcb *DeviceCreateBulk) Save(ctx context.Context) ([]*Device, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Device, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeviceCreateBulk) SaveX(ctx context.Context) []*Device {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeviceCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
