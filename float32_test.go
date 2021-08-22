package go_bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat32(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(0)

	assert.Nil(t, buf.WriteFloat32(25.25))
	assert.Nil(t, buf.WriteFloat32LE(25.25))
	assert.Nil(t, buf.SetFloat32(8, 52.52))
	assert.Nil(t, buf.SetFloat32LE(12, 52.52))

	assertFloat32(t, 25.25)(buf.GetFloat32(0))
	assertFloat32(t, 25.25)(buf.GetFloat32LE(4))
	assertFloat32(t, 25.25)(buf.ReadFloat32())
	assertFloat32(t, 25.25)(buf.ReadFloat32LE())
	assertFloat32(t, 52.52)(buf.ReadFloat32())
	assertFloat32(t, 52.52)(buf.ReadFloat32LE())
}

func assertFloat32(t *testing.T, expected float32) func(float32, error) {
	return func(v float32, err error) {
		assert.Nil(t, err)
		assert.EqualValues(t, expected, v)
	}
}
