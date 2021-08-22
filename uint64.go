package bytebuf

import (
	"encoding/binary"
)

var _ uint64Buffer = (*wrappedBuffer)(nil)

type uint64Buffer interface {
	ReadUint64() (uint64, error)
	ReadUint64LE() (uint64, error)
	WriteUint64(v uint64) error
	WriteUint64LE(v uint64) error
	GetUint64(index int) (uint64, error)
	GetUint64LE(index int) (uint64, error)
	SetUint64(index int, v uint64) error
	SetUint64LE(index int, v uint64) error
}

func (b *wrappedBuffer) ReadUint64() (uint64, error) {
	p := make([]byte, 8)
	n, err := b.Read(p)
	if n != 8 {
		return 0, err
	}
	return uint64(binary.BigEndian.Uint64(p)), err
}

func (b *wrappedBuffer) ReadUint64LE() (uint64, error) {
	p := make([]byte, 8)
	n, err := b.Read(p)
	if n != 8 {
		return 0, err
	}
	return uint64(binary.LittleEndian.Uint64(p)), err
}

func (b *wrappedBuffer) WriteUint64(v uint64) error {
	p := make([]byte, 8)
	binary.BigEndian.PutUint64(p, uint64(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) WriteUint64LE(v uint64) error {
	p := make([]byte, 8)
	binary.LittleEndian.PutUint64(p, uint64(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) GetUint64(index int) (uint64, error) {
	p := make([]byte, 8)
	n, err := b.Get(index, p)
	if n != 8 {
		return 0, err
	}
	return uint64(binary.BigEndian.Uint64(p)), err
}

func (b *wrappedBuffer) GetUint64LE(index int) (uint64, error) {
	p := make([]byte, 8)
	n, err := b.Get(index, p)
	if n != 8 {
		return 0, err
	}
	return uint64(binary.LittleEndian.Uint64(p)), err
}

func (b *wrappedBuffer) SetUint64(index int, v uint64) error {
	p := make([]byte, 8)
	binary.BigEndian.PutUint64(p, uint64(v))
	_, err := b.Set(index, p)
	return err
}

func (b *wrappedBuffer) SetUint64LE(index int, v uint64) error {
	p := make([]byte, 8)
	binary.LittleEndian.PutUint64(p, uint64(v))
	_, err := b.Set(index, p)
	return err
}
