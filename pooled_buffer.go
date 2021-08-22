package bytebuf

import (
	"io"
)

type pooledBuffer struct {
	released       bool
	allocator      *pooledAllocator
	pages          []pooledBufferPage
	discardedIndex int
}

func (b *pooledBuffer) get(index int, p []byte) (n int, err error) {
	if b.released {
		panic(ErrBufferReleased)
	}
	realIndex := index + b.discardedIndex
	n = len(p)
	if n == 0 {
		return
	}
	if (n + realIndex) > b.allocator.pageSize*len(b.pages) {
		n = b.allocator.pageSize*len(b.pages) - realIndex
		err = io.EOF
	}

	startPage := realIndex / b.allocator.pageSize
	endPage := (realIndex + n - 1) / b.allocator.pageSize
	page := startPage
	totalPages := endPage - startPage
	bottomIndex := realIndex - page*b.allocator.pageSize // start index of current page
	if totalPages == 0 {
		copy(p, b.pages[page][bottomIndex:bottomIndex+n])
		return
	}

	targetIndex := b.allocator.pageSize - bottomIndex
	// copy first page
	copy(p, b.pages[page][bottomIndex:])
	// copy pages in-between
	for page = startPage + 1; page < endPage; page++ {
		copy(p[targetIndex:], b.pages[page])
		targetIndex += b.allocator.pageSize
	}
	// copy last page
	copy(p[targetIndex:], b.pages[page][:n-targetIndex])
	return
}

func (b *pooledBuffer) set(index int, p []byte) (n int, err error) {
	if b.released {
		panic(ErrBufferReleased)
	}
	realIndex := index + b.discardedIndex
	n = len(p)
	if n == 0 {
		return
	}
	if (n + realIndex) > b.allocator.pageSize*len(b.pages) {
		b.growTo(n + realIndex)
	}

	startPage := realIndex / b.allocator.pageSize
	endPage := (realIndex + n - 1) / b.allocator.pageSize
	page := startPage
	totalPages := endPage - startPage
	bottomIndex := realIndex - page*b.allocator.pageSize // start index of current page
	if totalPages == 0 {
		copy(b.pages[page][bottomIndex:bottomIndex+n], p)
		return
	}

	targetIndex := b.allocator.pageSize - bottomIndex
	// copy first page
	copy(b.pages[page][bottomIndex:], p[:b.allocator.pageSize-bottomIndex])
	// copy pages in-between
	for page = startPage + 1; page < endPage; page++ {
		copy(b.pages[page], p[targetIndex:targetIndex+b.allocator.pageSize])
		targetIndex += b.allocator.pageSize
	}
	// copy last page
	copy(b.pages[page][:n-targetIndex], p[targetIndex:])
	return
}

func (b *pooledBuffer) discard(size int) {
	if b.released {
		panic(ErrBufferReleased)
	}
	realSize := size + b.discardedIndex
	discardPageSize := realSize / b.allocator.pageSize
	for i := 0; i < discardPageSize; i++ {
		b.allocator.release(b.pages[i])
	}
	b.pages = b.pages[discardPageSize:]
	b.discardedIndex = realSize % b.allocator.pageSize
}
func (b *pooledBuffer) capacity() int {
	if b.released {
		panic(ErrBufferReleased)
	}
	return b.allocator.pageSize*len(b.pages) - b.discardedIndex
}

func (b *pooledBuffer) growTo(newCapacity int) {
	if b.released {
		panic(ErrBufferReleased)
	}
	capacity := b.capacity()
	if newCapacity <= capacity {
		return
	}

	pages := b.allocator.allocate(newCapacity - capacity)
	b.pages = append(b.pages, pages...)
}

func (b *pooledBuffer) bytes() []byte {
	if len(b.pages) == 0 {
		return nil
	}
	p := make([]byte, b.capacity())
	off := b.allocator.pageSize - b.discardedIndex
	copy(p, b.pages[0][:b.discardedIndex])
	n := len(b.pages)
	for i := 1; i < n; i++ {
		copy(p[off:], b.pages[i][:])
		off += b.allocator.pageSize
	}
	return p
}

func (b *pooledBuffer) release() {
	if b.released {
		panic(ErrBufferReleased)
	}
	_ = b.allocator.Release(b)
}

func (b *pooledBuffer) clone() Instrument {
	if b.released {
		panic(ErrBufferReleased)
	}
	return b.allocator.clone(b)
}
