package go_bytebuf

type Buffer interface {
	Instrument() Instrument
	Release()
	Reset()

	Bytes() []byte
	Cap() int
	Len() int
	Grow(size int)
	Truncate()

	// ReaderIndex Returns the ReaderIndex of this buffer.
	ReaderIndex() int
	// MarkReaderIndex Marks the current ReaderIndex in this buffer
	MarkReaderIndex()
	// ResetReaderIndex Repositions the current ReaderIndex to the marked ReaderIndex in this buffer.
	ResetReaderIndex()
	// WriterIndex Returns the WriterIndex of this buffer.
	WriterIndex() int
	// MarkWriterIndex Marks the current WriterIndex in this buffer
	MarkWriterIndex()
	// ResetWriterIndex Repositions the current WriterIndex to the marked WriterIndex in this buffer.
	ResetWriterIndex()
	// Clone Returns a cloned buffer of this buffer.
	Clone() Buffer
	// Clear Sets the ReaderIndex and WriterIndex of this buffer to 0.
	Clear()
	// Resize Resize the capacity of this buffer.
	Resize(newCapacity int)
	// IsWritable Returns true if and only if this buffer has enough room to allow writing the specified number of elements.
	IsWritable(size int) bool
	// ReadableBytes Returns the number of readable bytes which is equal to (WriterIndex() - ReaderIndex()).

	bufferHelper
}

type Instrument interface {
	get(index int, p []byte) error
	set(index int, p []byte) error
	discard(size int)
	capacity() int
	resize(newCapacity int)
	clone() Instrument
	release()
}

type wrappedBuffer struct {
	instrument        Instrument
	readerIndex       int
	writerIndex       int
	markedReaderIndex int
	markedWriterIndex int
}

func (b *wrappedBuffer) Instrument() Instrument {
	return b.instrument
}

func (b *wrappedBuffer) Reset() {
	b.readerIndex = 0
	b.writerIndex = 0
}

func (b *wrappedBuffer) ReaderIndex() int {
	return b.readerIndex
}

func (b *wrappedBuffer) MarkReaderIndex() {
	b.markedReaderIndex = b.readerIndex
}

func (b *wrappedBuffer) ResetReaderIndex() {
	b.readerIndex = b.markedReaderIndex
	b.markedReaderIndex = 0
}

func (b *wrappedBuffer) WriterIndex() int {
	return b.writerIndex
}

func (b *wrappedBuffer) MarkWriterIndex() {
	b.markedWriterIndex = b.writerIndex
}

func (b *wrappedBuffer) ResetWriterIndex() {
	b.writerIndex = b.markedWriterIndex
	b.markedWriterIndex = 0
}

func (b *wrappedBuffer) Clone() Buffer {
	return &wrappedBuffer{
		instrument: b.instrument.clone(),
	}
}

func (b *wrappedBuffer) Release() {
	b.instrument.release()
}

func (b *wrappedBuffer) Bytes() []byte {
	p := make([]byte, b.Len())
	_ = b.ReadBytes(p)
	return p
}

func (b *wrappedBuffer) Clear() {
	b.markedReaderIndex = 0
	b.markedWriterIndex = 0
}

func (b *wrappedBuffer) Cap() int {
	return b.instrument.capacity()
}

func (b *wrappedBuffer) Len() int {
	n := b.writerIndex - b.readerIndex
	if n < 0 {
		return 0
	}
	return n
}

func (b *wrappedBuffer) Grow(size int) {
	capacity := b.instrument.capacity()
	if (b.markedWriterIndex + size) > capacity {
		b.instrument.resize(b.markedWriterIndex + size)
	}
}

func (b *wrappedBuffer) Resize(newCapacity int) {
	b.instrument.resize(newCapacity)
}

func (b *wrappedBuffer) Truncate() {
	b.instrument.discard(b.readerIndex)
	b.writerIndex -= b.readerIndex
	b.readerIndex = 0
}

func (b *wrappedBuffer) IsWritable(size int) bool {
	capacity := b.instrument.capacity()
	return b.markedWriterIndex+size <= capacity
}

func (b *wrappedBuffer) ReadBytes(p []byte) error {
	n := len(p)
	if n > b.Len() {
		return ErrNotEnoughBytes
	}
	err := b.instrument.get(b.readerIndex, p)
	if err != nil {
		return err
	}

	b.readerIndex += n
	return nil
}

func (b *wrappedBuffer) WriteBytes(p []byte) error {
	n := len(p)
	if !b.IsWritable(n) {
		return ErrNotEnoughSpace
	}
	err := b.instrument.set(b.writerIndex, p)
	if err != nil {
		return err
	}

	b.writerIndex += n
	return nil
}

func (b *wrappedBuffer) GetBytes(index int, p []byte) error {
	return b.instrument.get(index, p)
}

func (b *wrappedBuffer) SetBytes(index int, p []byte) error {
	return b.instrument.set(index, p)
}
