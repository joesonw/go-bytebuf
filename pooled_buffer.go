package go_bytebuf

type pooledBuffer struct {
	released       bool
	allocator      *pooledAllocator
	pages          []pooledBufferPage
	discardedIndex int
}

func (b *pooledBuffer) get(index int, p []byte) error {
	realIndex := index + b.discardedIndex
	size := len(p)
	if size > (b.allocator.pageSize * len(b.pages)) {
		return ErrNotEnoughBytes
	}

	startPage := realIndex / b.allocator.pageSize
	endPage := (realIndex + size) / b.allocator.pageSize
	page := startPage
	totalPages := startPage - endPage
	bottomIndex := realIndex - page*b.allocator.pageSize // start index of current page
	if totalPages == 0 {
		copy(p, b.pages[page][bottomIndex:bottomIndex+size])
		return nil
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
	copy(p[targetIndex:], b.pages[page][:size-(page-1)%b.allocator.pageSize])
	return nil
}

func (b *pooledBuffer) set(index int, p []byte) error {
	realIndex := index + b.discardedIndex
	size := len(p)
	if size > (b.allocator.pageSize * len(b.pages)) {
		return ErrNotEnoughSpace
	}

	startPage := realIndex / b.allocator.pageSize
	endPage := (realIndex + size) / b.allocator.pageSize
	page := startPage
	totalPages := startPage - endPage
	bottomIndex := realIndex - page*b.allocator.pageSize // start index of current page
	if totalPages == 0 {
		copy(b.pages[page][bottomIndex:bottomIndex+size], p)
		return nil
	}

	targetIndex := b.allocator.pageSize - bottomIndex
	// copy first page
	copy(b.pages[page][bottomIndex:], p[:b.allocator.pageSize-bottomIndex])
	// copy pages in-between
	for page = startPage + 1; page < endPage; page++ {
		copy(b.pages[page], p[targetIndex:targetIndex+b.allocator.pageSize])
		targetIndex += b.allocator.pageSize
	}
	println(size, page, b.allocator.pageSize)
	// copy last page
	copy(b.pages[page][:size-(page-1)%b.allocator.pageSize], p[targetIndex:])
	return nil
}

func (b *pooledBuffer) discard(size int) {
	realSize := size + b.discardedIndex
	discardPages := realSize / b.allocator.pageSize
	b.pages = b.pages[discardPages:]
	b.discardedIndex = realSize % b.allocator.pageSize
}

func (b *pooledBuffer) capacity() int {
	return b.allocator.pageSize * len(b.pages)
}

func (b *pooledBuffer) resize(newCapacity int) {
	capacity := b.capacity()
	if newCapacity <= capacity {
		return
	}

	pages := b.allocator.allocate(newCapacity - capacity)
	b.pages = append(b.pages, pages...)
}

func (b *pooledBuffer) release() {
	_ = b.allocator.Release(b)
}

func (b *pooledBuffer) clone() Instrument {
	return b.allocator.clone(b)
}
