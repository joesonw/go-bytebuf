package bytebuf

import (
	"encoding/binary"
	"math"
)

var _ Float32Buffer = (*wrappedBuffer)(nil)

type Float32Buffer interface {
	ReadFloat32() (float32, error)
	ReadFloat32LE() (float32, error)
	WriteFloat32(v float32) error
	WriteFloat32LE(v float32) error
	GetFloat32(index int) (float32, error)
	GetFloat32LE(index int) (float32, error)
	SetFloat32(index int, v float32) error
	SetFloat32LE(index int, v float32) error
}

func (b *wrappedBuffer) ReadFloat32() (float32, error) {
	p := make([]byte, 4)
	n, err := b.Read(p)
	if n != 4 {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(p)), err
}

func (b *wrappedBuffer) ReadFloat32LE() (float32, error) {
	p := make([]byte, 4)
	n, err := b.Read(p)
	if n != 4 {
		return 0, err
	}
	return math.Float32frombits(binary.LittleEndian.Uint32(p)), err
}

func (b *wrappedBuffer) WriteFloat32(v float32) error {
	p := make([]byte, 4)
	binary.BigEndian.PutUint32(p, math.Float32bits(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) WriteFloat32LE(v float32) error {
	p := make([]byte, 4)
	binary.LittleEndian.PutUint32(p, math.Float32bits(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) GetFloat32(index int) (float32, error) {
	p := make([]byte, 4)
	n, err := b.Get(index, p)
	if n != 4 {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(p)), err
}

func (b *wrappedBuffer) GetFloat32LE(index int) (float32, error) {
	p := make([]byte, 4)
	n, err := b.Get(index, p)
	if n != 4 {
		return 0, err
	}
	return math.Float32frombits(binary.LittleEndian.Uint32(p)), err
}

func (b *wrappedBuffer) SetFloat32(index int, v float32) error {
	p := make([]byte, 4)
	binary.BigEndian.PutUint32(p, math.Float32bits(v))
	_, err := b.Set(index, p)
	return err
}

func (b *wrappedBuffer) SetFloat32LE(index int, v float32) error {
	p := make([]byte, 4)
	binary.LittleEndian.PutUint32(p, math.Float32bits(v))
	_, err := b.Set(index, p)
	return err
}
