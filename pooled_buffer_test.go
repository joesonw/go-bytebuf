package go_bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPooledBuffer(t *testing.T) {
	allocator := PooledAllocator(8)

	buf := allocator.Allocate(16)
	assert.Nil(t, buf.WriteString("abcdefg"))
	assert.EqualValues(t, 7, buf.Len())
	assert.EqualValues(t, 16, buf.Cap())

	s, err := buf.ReadString(2)
	assert.Nil(t, err)
	assert.EqualValues(t, "ab", s)
	assert.EqualValues(t, 5, buf.Len())
	assert.EqualValues(t, 16, buf.Cap())

	assert.Nil(t, buf.WriteString("hijk"))
	assert.EqualValues(t, 9, buf.Len())
	s, err = buf.ReadString(8)
	assert.Nil(t, err)
	assert.EqualValues(t, "cdefghij", s)
	assert.EqualValues(t, 1, buf.Len())

	stat := allocator.Stat().(*PooledAllocatorStat)
	assert.EqualValues(t, 16, stat.Memory())
	assert.EqualValues(t, 1, stat.Count())
	assert.EqualValues(t, 2, stat.Pages())

	buf.Release()
	stat = allocator.Stat().(*PooledAllocatorStat)
	assert.EqualValues(t, 0, stat.Memory())
	assert.EqualValues(t, 0, stat.Count())
	assert.EqualValues(t, 0, stat.Pages())
}
