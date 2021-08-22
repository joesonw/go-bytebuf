package bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint32(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(0)

	assert.Nil(t, buf.WriteUint32(25))
	assert.Nil(t, buf.WriteUint32LE(25))
	assert.Nil(t, buf.SetUint32(8, 52))
	assert.Nil(t, buf.SetUint32LE(12, 52))

	assertUint32(t, 25)(buf.GetUint32(0))
	assertUint32(t, 25)(buf.GetUint32LE(4))
	assertUint32(t, 25)(buf.ReadUint32())
	assertUint32(t, 25)(buf.ReadUint32LE())
	assertUint32(t, 52)(buf.ReadUint32())
	assertUint32(t, 52)(buf.ReadUint32LE())
}

func assertUint32(t *testing.T, expected uint32) func(uint32, error) {
	return func(v uint32, err error) {
		assert.Nil(t, err)
		assert.EqualValues(t, expected, v)
	}
}
