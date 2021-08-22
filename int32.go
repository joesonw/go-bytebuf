package go_bytebuf

import (
	"encoding/binary"
)

var _ int32Buffer = (*wrappedBuffer)(nil)

type int32Buffer interface {
	ReadInt32() (int32, error)
	ReadInt32LE() (int32, error)
	WriteInt32(v int32) error
	WriteInt32LE(v int32) error
	GetInt32(index int) (int32, error)
	GetInt32LE(index int) (int32, error)
	SetInt32(index int, v int32) error
	SetInt32LE(index int, v int32) error
}

func (b *wrappedBuffer) ReadInt32() (int32, error) {
	p := make([]byte, 4)
	n, err := b.Read(p)
	if n != 4 {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(p)), err
}

func (b *wrappedBuffer) ReadInt32LE() (int32, error) {
	p := make([]byte, 4)
	n, err := b.Read(p)
	if n != 4 {
		return 0, err
	}
	return int32(binary.LittleEndian.Uint32(p)), err
}

func (b *wrappedBuffer) WriteInt32(v int32) error {
	p := make([]byte, 4)
	binary.BigEndian.PutUint32(p, uint32(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) WriteInt32LE(v int32) error {
	p := make([]byte, 4)
	binary.LittleEndian.PutUint32(p, uint32(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) GetInt32(index int) (int32, error) {
	p := make([]byte, 4)
	n, err := b.Get(index, p)
	if n != 4 {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(p)), err
}

func (b *wrappedBuffer) GetInt32LE(index int) (int32, error) {
	p := make([]byte, 4)
	n, err := b.Get(index, p)
	if n != 4 {
		return 0, err
	}
	return int32(binary.LittleEndian.Uint32(p)), err
}

func (b *wrappedBuffer) SetInt32(index int, v int32) error {
	p := make([]byte, 4)
	binary.BigEndian.PutUint32(p, uint32(v))
	_, err := b.Set(index, p)
	return err
}

func (b *wrappedBuffer) SetInt32LE(index int, v int32) error {
	p := make([]byte, 4)
	binary.LittleEndian.PutUint32(p, uint32(v))
	_, err := b.Set(index, p)
	return err
}
