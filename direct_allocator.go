package bytebuf

import (
	"runtime"
	"sync/atomic"
)

type DirectAllocateSizer func(capacity, newCapacity int) int

type DirectAllocatorOptions func(*directAllocator)

type directAllocator struct {
	statMemory int64
	statCount  int64

	sizer DirectAllocateSizer
}

func DirectAllocator(options ...DirectAllocatorOptions) Allocator {
	d := &directAllocator{}
	WithDirectMultiplier(1.2)(d)
	for _, o := range options {
		o(d)
	}
	return d
}

func (a *directAllocator) Allocate(capacity int) Buffer {
	instrument := &directBuffer{
		allocator: a,
		memory:    a.allocate(capacity),
	}
	runtime.SetFinalizer(instrument, a.Release)
	atomic.AddInt64(&a.statCount, 1)
	return &wrappedBuffer{instrument: instrument}
}

func (a *directAllocator) allocate(capacity int) []byte {
	p := make([]byte, capacity)
	atomic.AddInt64(&a.statMemory, int64(capacity))
	return p
}

func (a *directAllocator) clone(b *directBuffer) *directBuffer {
	clonedMemory := a.allocate(len(b.memory))
	copy(clonedMemory, b.memory)
	instrument := &directBuffer{
		memory:    clonedMemory,
		allocator: a,
	}
	runtime.SetFinalizer(instrument, a.Release)
	atomic.AddInt64(&a.statCount, 1)
	return instrument
}

func (a *directAllocator) release(p []byte) {
	atomic.AddInt64(&a.statMemory, int64(-len(p)))
}

func (a *directAllocator) Release(instrument Instrument) error {
	b, ok := instrument.(*directBuffer)
	if !ok {
		return ErrBufferTypeNotMatch
	}
	if !b.released {
		b.released = true
		a.release(b.memory)
		atomic.AddInt64(&a.statCount, -1)
	}

	return nil
}

type DirectAllocatorStat struct {
	memory int64
	count  int64
}

func (s DirectAllocatorStat) Memory() int64 {
	return s.memory
}

func (s DirectAllocatorStat) Count() int64 {
	return s.count
}

func (a *directAllocator) Stat() AllocatorStat {
	return &DirectAllocatorStat{
		memory: a.statMemory,
		count:  a.statCount,
	}
}

func WithDirectSizer(sizer DirectAllocateSizer) DirectAllocatorOptions {
	return func(allocator *directAllocator) {
		allocator.sizer = sizer
	}
}

func WithDirectMultiplier(mul float64) DirectAllocatorOptions {
	return WithDirectSizer(func(capacity, newCapacity int) int {
		c := capacity
		for c < newCapacity {
			c = int(float64(c) * 1.2)
		}
		return c
	})
}

func WithDirectStepper(step int) DirectAllocatorOptions {
	return WithDirectSizer(func(capacity, newCapacity int) int {
		c := capacity
		for c < newCapacity {
			c += step
		}
		return c
	})
}
