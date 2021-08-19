package go_bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectBuffer(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(16)
	assert.Nil(t, buf.WriteString("abcdefg"))
	assert.EqualValues(t, 7, buf.Len())
	assert.EqualValues(t, 16, buf.Cap())

	s, err := buf.ReadString(2)
	assert.Nil(t, err)
	assert.EqualValues(t, "ab", s)
	assert.EqualValues(t, 5, buf.Len())
	assert.EqualValues(t, 16, buf.Cap())

	stat := allocator.Stat()
	assert.EqualValues(t, 16, stat.Memory())
	assert.EqualValues(t, 1, stat.Count())

	buf.Release()
	stat = allocator.Stat()
	assert.EqualValues(t, 0, stat.Memory())
	assert.EqualValues(t, 0, stat.Count())
}
