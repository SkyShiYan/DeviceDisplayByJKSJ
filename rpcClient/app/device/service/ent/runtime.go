// Code generated by entc, DO NOT EDIT.

package ent

import (
	"rpcClient/app/device/service/ent/device"
	"rpcClient/app/device/service/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	deviceFields := schema.Device{}.Fields()
	_ = deviceFields
	// deviceDescCreatedAt is the schema descriptor for createdAt field.
	deviceDescCreatedAt := deviceFields[5].Descriptor()
	// device.DefaultCreatedAt holds the default value on creation for the createdAt field.
	device.DefaultCreatedAt = deviceDescCreatedAt.Default.(func() time.Time)
}
