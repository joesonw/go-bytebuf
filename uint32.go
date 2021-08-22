package go_bytebuf

import (
	"encoding/binary"
)

var _ uint32Buffer = (*wrappedBuffer)(nil)

type uint32Buffer interface {
	ReadUint32() (uint32, error)
	ReadUint32LE() (uint32, error)
	WriteUint32(v uint32) error
	WriteUint32LE(v uint32) error
	GetUint32(index int) (uint32, error)
	GetUint32LE(index int) (uint32, error)
	SetUint32(index int, v uint32) error
	SetUint32LE(index int, v uint32) error
}

func (b *wrappedBuffer) ReadUint32() (uint32, error) {
	p := make([]byte, 4)
	n, err := b.Read(p)
	if n != 4 {
		return 0, err
	}
	return uint32(binary.BigEndian.Uint32(p)), err
}

func (b *wrappedBuffer) ReadUint32LE() (uint32, error) {
	p := make([]byte, 4)
	n, err := b.Read(p)
	if n != 4 {
		return 0, err
	}
	return uint32(binary.LittleEndian.Uint32(p)), err
}

func (b *wrappedBuffer) WriteUint32(v uint32) error {
	p := make([]byte, 4)
	binary.BigEndian.PutUint32(p, uint32(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) WriteUint32LE(v uint32) error {
	p := make([]byte, 4)
	binary.LittleEndian.PutUint32(p, uint32(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) GetUint32(index int) (uint32, error) {
	p := make([]byte, 4)
	n, err := b.Get(index, p)
	if n != 4 {
		return 0, err
	}
	return uint32(binary.BigEndian.Uint32(p)), err
}

func (b *wrappedBuffer) GetUint32LE(index int) (uint32, error) {
	p := make([]byte, 4)
	n, err := b.Get(index, p)
	if n != 4 {
		return 0, err
	}
	return uint32(binary.LittleEndian.Uint32(p)), err
}

func (b *wrappedBuffer) SetUint32(index int, v uint32) error {
	p := make([]byte, 4)
	binary.BigEndian.PutUint32(p, uint32(v))
	_, err := b.Set(index, p)
	return err
}

func (b *wrappedBuffer) SetUint32LE(index int, v uint32) error {
	p := make([]byte, 4)
	binary.LittleEndian.PutUint32(p, uint32(v))
	_, err := b.Set(index, p)
	return err
}
