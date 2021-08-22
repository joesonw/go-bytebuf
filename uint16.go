package go_bytebuf

import (
	"encoding/binary"
)

var _ uint16Buffer = (*wrappedBuffer)(nil)

type uint16Buffer interface {
	ReadUint16() (uint16, error)
	ReadUint16LE() (uint16, error)
	WriteUint16(v uint16) error
	WriteUint16LE(v uint16) error
	GetUint16(index int) (uint16, error)
	GetUint16LE(index int) (uint16, error)
	SetUint16(index int, v uint16) error
	SetUint16LE(index int, v uint16) error
}

func (b *wrappedBuffer) ReadUint16() (uint16, error) {
	p := make([]byte, 2)
	n, err := b.Read(p)
	if n != 2 {
		return 0, err
	}
	return uint16(binary.BigEndian.Uint16(p)), err
}

func (b *wrappedBuffer) ReadUint16LE() (uint16, error) {
	p := make([]byte, 2)
	n, err := b.Read(p)
	if n != 2 {
		return 0, err
	}
	return uint16(binary.LittleEndian.Uint16(p)), err
}

func (b *wrappedBuffer) WriteUint16(v uint16) error {
	p := make([]byte, 2)
	binary.BigEndian.PutUint16(p, uint16(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) WriteUint16LE(v uint16) error {
	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, uint16(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) GetUint16(index int) (uint16, error) {
	p := make([]byte, 2)
	n, err := b.Get(index, p)
	if n != 2 {
		return 0, err
	}
	return uint16(binary.BigEndian.Uint16(p)), err
}

func (b *wrappedBuffer) GetUint16LE(index int) (uint16, error) {
	p := make([]byte, 2)
	n, err := b.Get(index, p)
	if n != 2 {
		return 0, err
	}
	return uint16(binary.LittleEndian.Uint16(p)), err
}

func (b *wrappedBuffer) SetUint16(index int, v uint16) error {
	p := make([]byte, 2)
	binary.BigEndian.PutUint16(p, uint16(v))
	_, err := b.Set(index, p)
	return err
}

func (b *wrappedBuffer) SetUint16LE(index int, v uint16) error {
	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, uint16(v))
	_, err := b.Set(index, p)
	return err
}
