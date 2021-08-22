package bytebuf

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tAssertPooledStat(t *testing.T, stat AllocatorStat, pages, count, memory int64) {
	assert.EqualValues(t, pages, stat.(*PooledAllocatorStat).Pages())
	tAssertStat(t, stat, count, memory)
}

func TestPooledBuffer(t *testing.T) {
	allocator := PooledAllocator(8)

	buf := allocator.Allocate(16).Instrument().(*pooledBuffer)
	tAssertPooledStat(t, allocator.Stat(), 2, 1, 16)
	assert.EqualValues(t, buf.capacity(), 16)
	tAssertGet(t, buf, 0, make([]byte, 18), 16, io.EOF)
	tAssertGetOk(t, buf, 0, make([]byte, 6), 6)
	tAssertGetOk(t, buf, 0, make([]byte, 8), 8)
	tAssertGetOk(t, buf, 0, make([]byte, 10), 10)

	tAssertSetOk(t, buf, 0, tRangeByteArray(1, 9), 8)
	tAssertPooledStat(t, allocator.Stat(), 2, 1, 16)
	tAssertGetOk(t, buf, 0, tRangeByteArray(1, 9), 8)

	tAssertSetOk(t, buf, 0, tRangeByteArray(1, 17), 16)
	tAssertPooledStat(t, allocator.Stat(), 2, 1, 16)

	tAssertSetOk(t, buf, 2, tRangeByteArray(1, 17), 16)
	tAssertPooledStat(t, allocator.Stat(), 3, 1, 24)
	assert.EqualValues(t, buf.capacity(), 24)
	tAssertGetOk(t, buf, 0, append([]byte{1, 2}, tRangeByteArray(1, 17)...), 18)
	tAssertGetOk(t, buf, 2, tRangeByteArray(1, 17), 16)

	buf2 := buf.clone()
	tAssertPooledStat(t, allocator.Stat(), 6, 2, 48)
	assert.EqualValues(t, buf2.capacity(), 24)

	buf.discard(2)
	tAssertPooledStat(t, allocator.Stat(), 6, 2, 48)
	assert.EqualValues(t, buf.capacity(), 22)
	assert.EqualValues(t, buf2.capacity(), 24)
	tAssertGetOk(t, buf, 2, tRangeByteArray(3, 17), 14)
	tAssertGetOk(t, buf, 0, tRangeByteArray(1, 17), 16)

	buf.discard(7)
	tAssertPooledStat(t, allocator.Stat(), 5, 2, 40)
	assert.EqualValues(t, buf.capacity(), 15)
	assert.EqualValues(t, buf2.capacity(), 24)
	tAssertGetOk(t, buf, 2, tRangeByteArray(10, 17), 7)
	tAssertGetOk(t, buf, 0, tRangeByteArray(8, 17), 9)
	tAssertGetOk(t, buf2, 0, append([]byte{1, 2}, tRangeByteArray(1, 17)...), 18)

	buf.release()
	tAssertPooledStat(t, allocator.Stat(), 3, 1, 24)
	assert.EqualValues(t, buf2.capacity(), 24)
}
