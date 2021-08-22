package go_bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffer(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(10)
	assert.EqualValues(t, 10, buf.Cap())
	assert.EqualValues(t, 0, buf.Len())

	buf2 := buf.Clone()
	assert.EqualValues(t, 10, buf2.Cap())
	assert.EqualValues(t, 0, buf2.Len())

	buf.Grow(6)
	assert.EqualValues(t, 16, buf.Cap())
	assert.EqualValues(t, 0, buf.Len())
	assert.EqualValues(t, 10, buf2.Cap())
	assert.EqualValues(t, 0, buf2.Len())

	buf.MarkWriterIndex()
	n, err := buf.WriteString("hello")
	assert.Nil(t, err)
	assert.EqualValues(t, n, 5)
	assert.EqualValues(t, 5, buf.Len())

	buf.ResetWriterIndex()
	assert.EqualValues(t, 0, buf.Len())

	buf.MarkReaderIndex()
	n, err = buf.WriteString("hello")
	assert.Nil(t, err)
	assert.EqualValues(t, n, 5)
	s, err := buf.ReadString(2)
	assert.Nil(t, err)
	assert.EqualValues(t, s, "he")
	assert.EqualValues(t, 3, buf.Len())

	buf.ResetReaderIndex()
	assert.EqualValues(t, 5, buf.Len())

	s, err = buf.ReadString(2)
	assert.Nil(t, err)
	assert.EqualValues(t, s, "he")
	assert.EqualValues(t, 3, buf.Len())

	assert.EqualValues(t, 2, buf.ReaderIndex())
	buf.DiscardReadBytes()
	assert.EqualValues(t, 3, buf.Len())
	assert.EqualValues(t, 0, buf.ReaderIndex())
	s, err = buf.ReadString(3)
	assert.Nil(t, err)
	assert.EqualValues(t, s, "llo")
	assert.EqualValues(t, 0, buf.Len())
}
