package go_bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt32(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(0)

	assert.Nil(t, buf.WriteInt32(25))
	assert.Nil(t, buf.WriteInt32LE(25))
	assert.Nil(t, buf.SetInt32(8, 52))
	assert.Nil(t, buf.SetInt32LE(12, 52))

	assertInt32(t, 25)(buf.GetInt32(0))
	assertInt32(t, 25)(buf.GetInt32LE(4))
	assertInt32(t, 25)(buf.ReadInt32())
	assertInt32(t, 25)(buf.ReadInt32LE())
	assertInt32(t, 52)(buf.ReadInt32())
	assertInt32(t, 52)(buf.ReadInt32LE())
}

func assertInt32(t *testing.T, expected int32) func(int32, error) {
	return func(v int32, err error) {
		assert.Nil(t, err)
		assert.EqualValues(t, expected, v)
	}
}
