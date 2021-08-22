package bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint64(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(0)

	assert.Nil(t, buf.WriteUint64(25))
	assert.Nil(t, buf.WriteUint64LE(25))
	assert.Nil(t, buf.SetUint64(16, 52))
	assert.Nil(t, buf.SetUint64LE(24, 52))

	assertUint64(t, 25)(buf.GetUint64(0))
	assertUint64(t, 25)(buf.GetUint64LE(8))
	assertUint64(t, 25)(buf.ReadUint64())
	assertUint64(t, 25)(buf.ReadUint64LE())
	assertUint64(t, 52)(buf.ReadUint64())
	assertUint64(t, 52)(buf.ReadUint64LE())
}

func assertUint64(t *testing.T, expected uint64) func(uint64, error) {
	return func(v uint64, err error) {
		assert.Nil(t, err)
		assert.EqualValues(t, expected, v)
	}
}
