package bytebuf

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectBuffer(t *testing.T) {
	allocator := DirectAllocator(WithDirectSizer(func(capacity, newCapacity int) int {
		return newCapacity
	}))

	buf := allocator.Allocate(8).Instrument().(*directBuffer)
	tAssertStat(t, allocator.Stat(), 1, 8)
	assert.EqualValues(t, buf.capacity(), 8)
	tAssertGet(t, buf, 0, make([]byte, 10), 8, io.EOF)

	tAssertSetOk(t, buf, 0, tRangeByteArray(1, 9), 8)
	tAssertStat(t, allocator.Stat(), 1, 8)
	assert.EqualValues(t, buf.capacity(), 8)
	tAssertGet(t, buf, 0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 0, 0}, 8, io.EOF)

	tAssertSetOk(t, buf, 8, []byte{9, 10}, 2)
	tAssertStat(t, allocator.Stat(), 1, 10)
	assert.EqualValues(t, buf.capacity(), 10)
	tAssertGetOk(t, buf, 0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)

	buf2 := buf.clone()
	tAssertStat(t, allocator.Stat(), 2, 20)
	assert.EqualValues(t, buf2.capacity(), 10)

	buf.discard(2)
	tAssertStat(t, allocator.Stat(), 2, 20)
	assert.EqualValues(t, buf.capacity(), 10)
	tAssertGetOk(t, buf, 0, []byte{3, 4, 5, 6, 7, 8, 9, 10, 9, 10}, 10)

	buf.release()
	assert.True(t, buf.released)
	tAssertStat(t, allocator.Stat(), 1, 10)
	tAssertGetOk(t, buf2, 0, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)

	buf2.release()
	tAssertStat(t, allocator.Stat(), 0, 0)
}
