package go_bytebuf

import (
	"encoding/binary"
)

var _ int16Buffer = (*wrappedBuffer)(nil)

type int16Buffer interface {
	ReadInt16() (int16, error)
	ReadInt16LE() (int16, error)
	WriteInt16(v int16) error
	WriteInt16LE(v int16) error
	GetInt16(index int) (int16, error)
	GetInt16LE(index int) (int16, error)
	SetInt16(index int, v int16) error
	SetInt16LE(index int, v int16) error
}

func (b *wrappedBuffer) ReadInt16() (int16, error) {
	p := make([]byte, 2)
	n, err := b.Read(p)
	if n != 2 {
		return 0, err
	}
	return int16(binary.BigEndian.Uint16(p)), err
}

func (b *wrappedBuffer) ReadInt16LE() (int16, error) {
	p := make([]byte, 2)
	n, err := b.Read(p)
	if n != 2 {
		return 0, err
	}
	return int16(binary.LittleEndian.Uint16(p)), err
}

func (b *wrappedBuffer) WriteInt16(v int16) error {
	p := make([]byte, 2)
	binary.BigEndian.PutUint16(p, uint16(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) WriteInt16LE(v int16) error {
	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, uint16(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) GetInt16(index int) (int16, error) {
	p := make([]byte, 2)
	n, err := b.Get(index, p)
	if n != 2 {
		return 0, err
	}
	return int16(binary.BigEndian.Uint16(p)), err
}

func (b *wrappedBuffer) GetInt16LE(index int) (int16, error) {
	p := make([]byte, 2)
	n, err := b.Get(index, p)
	if n != 2 {
		return 0, err
	}
	return int16(binary.LittleEndian.Uint16(p)), err
}

func (b *wrappedBuffer) SetInt16(index int, v int16) error {
	p := make([]byte, 2)
	binary.BigEndian.PutUint16(p, uint16(v))
	_, err := b.Set(index, p)
	return err
}

func (b *wrappedBuffer) SetInt16LE(index int, v int16) error {
	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, uint16(v))
	_, err := b.Set(index, p)
	return err
}
