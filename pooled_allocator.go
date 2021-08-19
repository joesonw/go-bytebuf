package go_bytebuf

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type pooledBufferPage []byte

type pooledAllocator struct {
	statMemory int64
	statCount  int64
	statPages  int64
	pageSize   int
	pagePool   *sync.Pool
}

func PooledAllocator(pageSize int) Allocator {
	return &pooledAllocator{
		pageSize: pageSize,
		pagePool: &sync.Pool{
			New: func() interface{} {
				p := make([]byte, pageSize)
				return pooledBufferPage(p)
			},
		},
	}
}

func (a *pooledAllocator) Allocate(capacity int) Buffer {
	instrument := &pooledBuffer{
		allocator: a,
		pages:     a.allocate(capacity),
	}
	runtime.SetFinalizer(instrument, a.Release)
	atomic.AddInt64(&a.statCount, 1)
	return &wrappedBuffer{
		instrument: instrument,
	}
}

func (a *pooledAllocator) newPage() pooledBufferPage {
	page := a.pagePool.Get().(pooledBufferPage)
	atomic.AddInt64(&a.statPages, 1)
	atomic.AddInt64(&a.statMemory, int64(len(page)))
	return page
}

func (a *pooledAllocator) allocate(capacity int) []pooledBufferPage {
	n := capacity / a.pageSize
	if capacity%a.pageSize > 0 {
		n += 1
	}
	pages := make([]pooledBufferPage, n)
	for i := range pages {
		pages[i] = a.newPage()
	}
	return pages
}

func (a *pooledAllocator) release(page pooledBufferPage) {
	a.pagePool.Put(page)
	atomic.AddInt64(&a.statPages, -1)
	atomic.AddInt64(&a.statMemory, int64(-len(page)))
}

func (a *pooledAllocator) clone(b *pooledBuffer) *pooledBuffer {
	pages := make([]pooledBufferPage, len(b.pages))
	for i := range pages {
		page := b.allocator.newPage()
		copy(page, b.pages[i])
	}

	instrument := &pooledBuffer{
		allocator: a,
		pages:     pages,
	}
	runtime.SetFinalizer(instrument, a.Release)
	atomic.AddInt64(&a.statCount, 1)
	return instrument
}

func (a *pooledAllocator) Release(instrument Instrument) error {
	buf, ok := instrument.(*pooledBuffer)
	if !ok {
		return ErrBufferTypeNotMatch
	}
	if !buf.released {
		buf.released = true
		atomic.AddInt64(&a.statCount, -1)
		for _, page := range buf.pages {
			a.release(page)
		}
	}
	return nil
}

type PooledAllocatorStat struct {
	memory int64
	count  int64
	pages  int64
}

func (s PooledAllocatorStat) Memory() int64 {
	return s.memory
}

func (s PooledAllocatorStat) Count() int64 {
	return s.count
}

func (s PooledAllocatorStat) Pages() int64 {
	return s.pages
}

func (a *pooledAllocator) Stat() AllocatorStat {
	return &PooledAllocatorStat{
		memory: a.statMemory,
		count:  a.statCount,
		pages:  a.statPages,
	}
}
