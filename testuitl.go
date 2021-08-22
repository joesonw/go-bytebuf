package bytebuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func tAssertStat(t *testing.T, stat AllocatorStat, count, memory int64) {
	assert.Equal(t, count, stat.Count())
	assert.EqualValues(t, memory, stat.Memory())
}

func tAssertGetOk(t *testing.T, instrument Instrument, index int, p []byte, n int) {
	tAssertGet(t, instrument, index, p, n, nil)
}

func tAssertGet(t *testing.T, instrument Instrument, index int, p []byte, n int, err error) {
	p2 := make([]byte, len(p))
	m, e := instrument.get(index, p2)
	assert.EqualValues(t, m, n)
	assert.EqualValues(t, p, p2)
	if err != nil {
		assert.ErrorIs(t, e, err)
	} else {
		assert.Nil(t, e)
	}
}

func tAssertSetOk(t *testing.T, instrument Instrument, index int, p []byte, n int) {
	tAssertSet(t, instrument, index, p, n, nil)
}

func tAssertSet(t *testing.T, instrument Instrument, index int, p []byte, n int, err error) {
	m, e := instrument.set(index, p)
	assert.EqualValues(t, m, n)
	if err != nil {
		assert.ErrorIs(t, e, err)
	} else {
		assert.Nil(t, err)
	}
}

func tRangeByteArray(start, end int) []byte {
	p := make([]byte, end-start)
	for i := start; i < end; i++ {
		p[i-start] = byte(i)
	}
	return p
}
