package bytebuf

import (
	"io"
)

const MinRead = 512

type Buffer interface {
	// Instrument get the underlying instrument
	Instrument() Instrument

	// Clear clear marked indexes
	Clear()
	// Release release underlying instrument, this buffer can no longer be used
	Release()
	// Reset reset the buffer to its full capacity, means length = 0
	Reset()
	// Clone clone an entire new buffer, with the same instrument implementation (different instance)
	Clone() Buffer
	// Cap returns current instrument allocated capacity
	Cap() int
	// Len current readable number of bytes
	Len() int
	// Grow grow capacity with at least this much more to write
	Grow(size int)
	// Ensure at least this much capacity
	Ensure(capacity int)
	// DiscardReadBytes discard already read bytes, this will reset marks and indexes
	DiscardReadBytes()
	// Bytes return all readable bytes
	Bytes() []byte
	// UnsafeBytes return all underlying bytes, this may contain discarded bytes
	UnsafeBytes() []byte
	// String return all readable as string
	String() string

	// ReaderIndex current reader index in memory
	ReaderIndex() int
	MarkReaderIndex()
	ResetReaderIndex()
	// WriterIndex current writer index in memory
	WriterIndex() int
	MarkWriterIndex()
	ResetWriterIndex()

	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Get(index int, p []byte) (int, error)
	Set(index int, p []byte) (int, error)
	ReadFrom(r io.Reader) (int64, error)
	WriteTo(w io.Writer) (n int64, err error)

	ReadString(n int) (string, error)
	WriteString(s string) (int, error)
	ReadByte() (byte, error)
	WriteByte(b byte) error
	GetByte(index int) (byte, error)
	SetByte(index int, b byte) error

	Int16Buffer
	Uint16Buffer
	Int32Buffer
	Uint32Buffer
	Int64Buffer
	Uint64Buffer
	Float32Buffer
	Float64Buffer
}

type Instrument interface {
	get(index int, p []byte) (int, error)
	set(index int, p []byte) (int, error)
	discard(size int)
	capacity() int
	growTo(newCapacity int)
	clone() Instrument
	bytes() []byte
	release()
}

type basicBuffer interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Get(index int, p []byte) (int, error)
	Set(index int, p []byte) (int, error)
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

func (b *wrappedBuffer) Clear() {
	b.markedReaderIndex = 0
	b.markedWriterIndex = 0
}

func (b *wrappedBuffer) Release() {
	b.instrument.release()
}

func (b *wrappedBuffer) Reset() {
	b.readerIndex = 0
	b.writerIndex = 0
}

func (b *wrappedBuffer) Clone() Buffer {
	return &wrappedBuffer{
		instrument: b.instrument.clone(),
	}
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
	b.instrument.growTo(b.instrument.capacity() + size)
}

func (b *wrappedBuffer) Ensure(capacity int) {
	if capacity > b.instrument.capacity() {
		b.instrument.growTo(capacity)
	}
}

func (b *wrappedBuffer) DiscardReadBytes() {
	b.instrument.discard(b.readerIndex)
	b.writerIndex -= b.readerIndex
	b.readerIndex = 0
	b.markedWriterIndex = b.writerIndex
	b.markedReaderIndex = 0
}

func (b *wrappedBuffer) Bytes() []byte {
	p := make([]byte, b.Len())
	_, _ = b.Read(p)
	return p
}

func (b *wrappedBuffer) UnsafeBytes() []byte {
	return b.instrument.bytes()
}

func (b *wrappedBuffer) String() string {
	return string(b.Bytes())
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

func (b *wrappedBuffer) Resize(newCapacity int) {
	b.instrument.growTo(newCapacity)
}

func (b *wrappedBuffer) Read(p []byte) (n int, err error) {
	n, err = b.instrument.get(b.readerIndex, p)
	b.readerIndex += n
	return
}

func (b *wrappedBuffer) Write(p []byte) (n int, err error) {
	n, err = b.instrument.set(b.writerIndex, p)
	b.writerIndex += n
	return
}

func (b *wrappedBuffer) Get(index int, p []byte) (int, error) {
	return b.instrument.get(index, p)
}

func (b *wrappedBuffer) Set(index int, p []byte) (int, error) {
	return b.instrument.set(index, p)
}

func (b *wrappedBuffer) ReadFrom(r io.Reader) (n int64, err error) {
	for {
		b.Grow(MinRead)
		p := make([]byte, MinRead)
		read, e := r.Read(p)
		n += int64(read)
		wrote, e := b.Write(p)
		if e != nil {
			return n, e
		}
		if wrote != read {
			return n, io.ErrShortWrite
		}
		if e == io.EOF {
			return n, nil // e is EOF, so return nil explicitly
		}
		if e != nil {
			return n, err
		}
	}
}
func (b *wrappedBuffer) WriteTo(w io.Writer) (n int64, err error) {
	if nBytes := b.Len(); nBytes > 0 {
		wrote, e := w.Write(b.Bytes())
		if wrote > nBytes {
			panic(ErrInvalidLengthWrote)
		}
		n = int64(nBytes)
		if e != nil {
			return n, e
		}
		if wrote != nBytes {
			return n, io.ErrShortWrite
		}
	}
	b.Reset()
	return n, nil
}

func (b *wrappedBuffer) ReadString(n int) (string, error) {
	p := make([]byte, n)
	m, err := b.Read(p)
	return string(p[:m]), err
}

func (b *wrappedBuffer) WriteString(s string) (int, error) {
	return b.Write([]byte(s))
}

func (b *wrappedBuffer) ReadByte() (byte, error) {
	p := make([]byte, 1)
	_, err := b.Read(p)
	return p[0], err
}

func (b *wrappedBuffer) WriteByte(b2 byte) error {
	_, err := b.Write([]byte{b2})
	return err
}

func (b *wrappedBuffer) GetByte(index int) (byte, error) {
	p := make([]byte, 1)
	_, err := b.Get(index, p)
	return p[0], err
}

func (b *wrappedBuffer) SetByte(index int, b2 byte) error {
	_, err := b.Set(index, []byte{b2})
	return err
}
