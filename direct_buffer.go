package go_bytebuf

type directBuffer struct {
	allocator *directAllocator
	released  bool
	memory    []byte
}

func (b *directBuffer) get(index int, p []byte) error {
	n := len(p)
	if n+index > len(b.memory) {
		return ErrNotEnoughBytes
	}
	copy(p, b.memory[index:index+n])
	return nil
}

func (b *directBuffer) set(index int, p []byte) error {
	size := len(p)
	if size+index > len(b.memory) {
		return ErrNotEnoughSpace
	}
	copy(b.memory[index:index+size], p)
	return nil
}

func (b *directBuffer) discard(size int) {
	copy(b.memory, b.memory[size:])
}

func (b *directBuffer) capacity() int {
	return len(b.memory)
}

func (b *directBuffer) resize(newCapacity int) {
	capacity := b.capacity()
	if newCapacity <= capacity {
		return
	}
	newMemory := b.allocator.allocate(b.allocator.sizer(capacity, newCapacity))
	copy(newMemory, b.memory)
	b.allocator.release(b.memory)
	b.memory = newMemory
}

func (b *directBuffer) release() {
	_ = b.allocator.Release(b)
}

func (b *directBuffer) clone() Instrument {
	return b.allocator.clone(b)
}
