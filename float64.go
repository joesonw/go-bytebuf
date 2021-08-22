package bytebuf

import (
	"encoding/binary"
	"math"
)

var _ Float64Buffer = (*wrappedBuffer)(nil)

type Float64Buffer interface {
	ReadFloat64() (float64, error)
	ReadFloat64LE() (float64, error)
	WriteFloat64(v float64) error
	WriteFloat64LE(v float64) error
	GetFloat64(index int) (float64, error)
	GetFloat64LE(index int) (float64, error)
	SetFloat64(index int, v float64) error
	SetFloat64LE(index int, v float64) error
}

func (b *wrappedBuffer) ReadFloat64() (float64, error) {
	p := make([]byte, 8)
	n, err := b.Read(p)
	if n != 8 {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(p)), err
}

func (b *wrappedBuffer) ReadFloat64LE() (float64, error) {
	p := make([]byte, 8)
	n, err := b.Read(p)
	if n != 8 {
		return 0, err
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(p)), err
}

func (b *wrappedBuffer) WriteFloat64(v float64) error {
	p := make([]byte, 8)
	binary.BigEndian.PutUint64(p, math.Float64bits(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) WriteFloat64LE(v float64) error {
	p := make([]byte, 8)
	binary.LittleEndian.PutUint64(p, math.Float64bits(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) GetFloat64(index int) (float64, error) {
	p := make([]byte, 8)
	n, err := b.Get(index, p)
	if n != 8 {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(p)), err
}

func (b *wrappedBuffer) GetFloat64LE(index int) (float64, error) {
	p := make([]byte, 8)
	n, err := b.Get(index, p)
	if n != 8 {
		return 0, err
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(p)), err
}

func (b *wrappedBuffer) SetFloat64(index int, v float64) error {
	p := make([]byte, 8)
	binary.BigEndian.PutUint64(p, math.Float64bits(v))
	_, err := b.Set(index, p)
	return err
}

func (b *wrappedBuffer) SetFloat64LE(index int, v float64) error {
	p := make([]byte, 8)
	binary.LittleEndian.PutUint64(p, math.Float64bits(v))
	_, err := b.Set(index, p)
	return err
}
