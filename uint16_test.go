package go_bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint16(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(0)

	assert.Nil(t, buf.WriteUint16(25))
	assert.Nil(t, buf.WriteUint16LE(25))
	assert.Nil(t, buf.SetUint16(4, 52))
	assert.Nil(t, buf.SetUint16LE(6, 52))

	assertUint16(t, 25)(buf.GetUint16(0))
	assertUint16(t, 25)(buf.GetUint16LE(2))
	assertUint16(t, 25)(buf.ReadUint16())
	assertUint16(t, 25)(buf.ReadUint16LE())
	assertUint16(t, 52)(buf.ReadUint16())
	assertUint16(t, 52)(buf.ReadUint16LE())
}

func assertUint16(t *testing.T, expected uint16) func(uint16, error) {
	return func(v uint16, err error) {
		assert.Nil(t, err)
		assert.EqualValues(t, expected, v)
	}
}
