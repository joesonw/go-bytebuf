package go_bytebuf

import (
	"io"
)

const MinRead = 512

type Buffer interface {
	// Instrument get the underlying instrument
	Instrument() Instrument

	Clear()
	Release()
	Reset()
	Clone() Buffer
	Cap() int
	Len() int
	Grow(size int)
	DiscardReadBytes()
	Bytes() []byte
	UnsafeBytes() []byte
	String() string

	ReaderIndex() int
	MarkReaderIndex()
	ResetReaderIndex()
	WriterIndex() int
	MarkWriterIndex()
	ResetWriterIndex()

	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Get(index int, p []byte) (int, error)
	Set(index int, p []byte) (int, error)
	ReadFrom(r io.Reader) (int64, error)
	WriteTo(w io.Writer) (n int64, err error)

	bufferHelper
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
	capacity := b.instrument.capacity()
	if (b.markedWriterIndex + size) > capacity {
		b.instrument.growTo(b.markedWriterIndex + size)
	}
}

func (b *wrappedBuffer) DiscardReadBytes() {
	b.instrument.discard(b.readerIndex)
	b.writerIndex -= b.readerIndex
	b.readerIndex = 0
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
