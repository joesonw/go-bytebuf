package go_bytebuf

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	benchmarkBatchSize = 128
)

var (
	benchmarkSizes         = []int{512, 1024, 2048, 4096}
	directBufferGrowRatios = []float64{2, 4}
	pooledBufferPageSizes  = []int{128, 512, 1024}
)

func BenchmarkPooledBuffer(b *testing.B) {
	for _, size := range benchmarkSizes {
		for _, pageSize := range pooledBufferPageSizes {
			b.Run(fmt.Sprintf("BenchmarkPooledBuffer_%d_%d", size, pageSize), func(b *testing.B) {
				runPooledBuffer(b, pageSize, size)
			})
		}
	}
}

func BenchmarkDirectBuffer(b *testing.B) {
	for _, size := range benchmarkSizes {
		for _, ratio := range directBufferGrowRatios {
			b.Run(fmt.Sprintf("BenchmarkDirectBuffer_%d_%.0f", size, ratio), func(b *testing.B) {
				runDirectBuffer(b, ratio, size)
			})
		}
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("BenchmarkBytesBuffer_%d", size), func(b *testing.B) {
			runBytesBuffer(b, size)
		})
	}
}

func runPooledBuffer(b *testing.B, pageSize int, size int) {
	allocator := PooledAllocator(pageSize)
	doBenchmark(b, size, func() io.ReadWriter {
		return allocator.Allocate(benchmarkBatchSize)
	})
}

func runDirectBuffer(b *testing.B, growRatio float64, size int) {
	allocator := DirectAllocator(WithDirectMultiplier(growRatio))
	doBenchmark(b, size, func() io.ReadWriter {
		return allocator.Allocate(benchmarkBatchSize)
	})
}

func runBytesBuffer(b *testing.B, size int) {
	doBenchmark(b, size, func() io.ReadWriter {
		return bytes.NewBuffer(nil)
	})
}

func doBenchmark(b *testing.B, size int, newFunc func() io.ReadWriter) {
	arr := make([]io.ReadWriter, b.N)
	batches := size / benchmarkBatchSize
	n := b.N

	b.Cleanup(func() {
		for i := 0; i < n; i++ {
			if r, ok := arr[i].(interface {
				Release()
			}); ok {
				r.Release()
			}
		}
	})

	for i := 0; i < n; i++ {
		buf := newFunc()

		for j := 0; j < batches; j++ {
			p := make([]byte, benchmarkBatchSize)
			_, err := rand.Read(p)
			assert.Nil(b, err)
			_, err = buf.Write(p)
			assert.Nil(b, err)
		}

		arr[i] = buf
	}

	for i := 0; i < n; i++ {
		p := make([]byte, size)
		_, err := arr[i].Read(p)
		assert.Nil(b, err)
		assert.Equal(b, size, len(p))
	}
}
