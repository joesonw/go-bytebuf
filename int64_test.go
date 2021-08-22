package bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(0)

	assert.Nil(t, buf.WriteInt64(25))
	assert.Nil(t, buf.WriteInt64LE(25))
	assert.Nil(t, buf.SetInt64(16, 52))
	assert.Nil(t, buf.SetInt64LE(24, 52))

	assertInt64(t, 25)(buf.GetInt64(0))
	assertInt64(t, 25)(buf.GetInt64LE(8))
	assertInt64(t, 25)(buf.ReadInt64())
	assertInt64(t, 25)(buf.ReadInt64LE())
	assertInt64(t, 52)(buf.ReadInt64())
	assertInt64(t, 52)(buf.ReadInt64LE())
}

func assertInt64(t *testing.T, expected int64) func(int64, error) {
	return func(v int64, err error) {
		assert.Nil(t, err)
		assert.EqualValues(t, expected, v)
	}
}
