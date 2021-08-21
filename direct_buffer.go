package go_bytebuf

import (
	"io"
)

type directBuffer struct {
	allocator *directAllocator
	released  bool
	memory    []byte
}

func (b *directBuffer) get(index int, p []byte) (n int, err error) {
	if b.released {
		panic(ErrBufferReleased)
	}
	n = len(p)
	capacity := b.capacity()
	if n+index > capacity {
		n = capacity - index
		err = io.EOF
	}
	copy(p, b.memory[index:index+n])
	return
}

func (b *directBuffer) set(index int, p []byte) (n int, err error) {
	if b.released {
		panic(ErrBufferReleased)
	}
	n = len(p)
	if n+index > len(b.memory) {
		b.growTo(n + index)
	}
	copy(b.memory[index:index+n], p)
	return
}

func (b *directBuffer) discard(size int) {
	if b.released {
		panic(ErrBufferReleased)
	}
	copy(b.memory, b.memory[size:])
}

func (b *directBuffer) capacity() int {
	if b.released {
		panic(ErrBufferReleased)
	}
	return len(b.memory)
}

func (b *directBuffer) growTo(newCapacity int) {
	if b.released {
		panic(ErrBufferReleased)
	}
	capacity := b.capacity()
	if newCapacity <= capacity {
		return
	}
	newMemory := b.allocator.allocate(b.allocator.sizer(capacity, newCapacity))
	copy(newMemory, b.memory)
	b.allocator.release(b.memory)
	b.memory = newMemory
}

func (b *directBuffer) bytes() []byte {
	return b.memory[:]
}

func (b *directBuffer) release() {
	if b.released {
		panic(ErrBufferReleased)
	}
	_ = b.allocator.Release(b)
}

func (b *directBuffer) clone() Instrument {
	if b.released {
		panic(ErrBufferReleased)
	}
	return b.allocator.clone(b)
}
