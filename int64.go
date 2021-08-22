package go_bytebuf

import (
	"encoding/binary"
)

var _ int64Buffer = (*wrappedBuffer)(nil)

type int64Buffer interface {
	ReadInt64() (int64, error)
	ReadInt64LE() (int64, error)
	WriteInt64(v int64) error
	WriteInt64LE(v int64) error
	GetInt64(index int) (int64, error)
	GetInt64LE(index int) (int64, error)
	SetInt64(index int, v int64) error
	SetInt64LE(index int, v int64) error
}

func (b *wrappedBuffer) ReadInt64() (int64, error) {
	p := make([]byte, 8)
	n, err := b.Read(p)
	if n != 8 {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(p)), err
}

func (b *wrappedBuffer) ReadInt64LE() (int64, error) {
	p := make([]byte, 8)
	n, err := b.Read(p)
	if n != 8 {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(p)), err
}

func (b *wrappedBuffer) WriteInt64(v int64) error {
	p := make([]byte, 8)
	binary.BigEndian.PutUint64(p, uint64(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) WriteInt64LE(v int64) error {
	p := make([]byte, 8)
	binary.LittleEndian.PutUint64(p, uint64(v))
	_, err := b.Write(p)
	return err
}

func (b *wrappedBuffer) GetInt64(index int) (int64, error) {
	p := make([]byte, 8)
	n, err := b.Get(index, p)
	if n != 8 {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(p)), err
}

func (b *wrappedBuffer) GetInt64LE(index int) (int64, error) {
	p := make([]byte, 8)
	n, err := b.Get(index, p)
	if n != 8 {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(p)), err
}

func (b *wrappedBuffer) SetInt64(index int, v int64) error {
	p := make([]byte, 8)
	binary.BigEndian.PutUint64(p, uint64(v))
	_, err := b.Set(index, p)
	return err
}

func (b *wrappedBuffer) SetInt64LE(index int, v int64) error {
	p := make([]byte, 8)
	binary.LittleEndian.PutUint64(p, uint64(v))
	_, err := b.Set(index, p)
	return err
}
