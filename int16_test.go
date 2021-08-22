package bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt16(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(0)

	assert.Nil(t, buf.WriteInt16(25))
	assert.Nil(t, buf.WriteInt16LE(25))
	assert.Nil(t, buf.SetInt16(4, 52))
	assert.Nil(t, buf.SetInt16LE(6, 52))

	assertInt16(t, 25)(buf.GetInt16(0))
	assertInt16(t, 25)(buf.GetInt16LE(2))
	assertInt16(t, 25)(buf.ReadInt16())
	assertInt16(t, 25)(buf.ReadInt16LE())
	assertInt16(t, 52)(buf.ReadInt16())
	assertInt16(t, 52)(buf.ReadInt16LE())
}

func assertInt16(t *testing.T, expected int16) func(int16, error) {
	return func(v int16, err error) {
		assert.Nil(t, err)
		assert.EqualValues(t, expected, v)
	}
}
