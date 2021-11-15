package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
// Name            string
// HardwareKey     string
// Defaultlayoutid int32
// Status          string
// Storenumber     string
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional().
			Nillable(),
		field.String("hardwareKey").
			Optional().
			Unique().
			Nillable(),
		field.Int32("defaultLayoutId").
			Optional().
			Nillable(),
		field.Int32("status").
			Optional().
			Nillable(),
		field.String("storeNumber").
			Optional().
			Nillable(),
		field.Time("createdAt").
			Optional().
			Default(time.Now).
			Immutable(),
	}
}

func (Device) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		// index.Fields("field1", "field2"),
		// 唯一约束索引
		// index.Fields("hardwareKey").
		// 	Unique(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return nil
}
