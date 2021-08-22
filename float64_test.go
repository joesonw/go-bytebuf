package go_bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(0)

	assert.Nil(t, buf.WriteFloat64(25.25))
	assert.Nil(t, buf.WriteFloat64LE(25.25))
	assert.Nil(t, buf.SetFloat64(16, 52.52))
	assert.Nil(t, buf.SetFloat64LE(24, 52.52))

	assertFloat64(t, 25.25)(buf.GetFloat64(0))
	assertFloat64(t, 25.25)(buf.GetFloat64LE(8))
	assertFloat64(t, 25.25)(buf.ReadFloat64())
	assertFloat64(t, 25.25)(buf.ReadFloat64LE())
	assertFloat64(t, 52.52)(buf.ReadFloat64())
	assertFloat64(t, 52.52)(buf.ReadFloat64LE())
}

func assertFloat64(t *testing.T, expected float64) func(float64, error) {
	return func(v float64, err error) {
		assert.Nil(t, err)
		assert.EqualValues(t, expected, v)
	}
}
