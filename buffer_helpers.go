package go_bytebuf

import (
	"encoding/binary"
	"math"
)

type bufferHelper interface {
	ReadString(n int) (string, error)
	WriteString(s string) (int, error)

	ReadByte() (byte, error)
	WriteByte(b byte) error
	GetByte(index int) (byte, error)
	SetByte(index int, b byte) error

	ReadUint16() (uint16, error)
	ReadUint16LE() (uint16, error)
	WriteUint16(v uint16) error
	WriteUint16LE(v uint16) error
	GetUint16(index int) (uint16, error)
	GetUint16LE(index int) (uint16, error)
	SetUint16(index int, v uint16) error
	SetUint16LE(index int, v uint16) error

	ReadInt16() (int16, error)
	ReadInt16LE() (int16, error)
	WriteInt16(v int16) error
	WriteInt16LE(v int16) error
	GetInt16(index int) (int16, error)
	GetInt16LE(index int) (int16, error)
	SetInt16(index int, v int16) error
	SetInt16LE(index int, v int16) error

	ReadUint32() (uint32, error)
	ReadUint32LE() (uint32, error)
	WriteUint32(v uint32) error
	WriteUint32LE(v uint32) error
	GetUint32(index int) (uint32, error)
	GetUint32LE(index int) (uint32, error)
	SetUint32(index int, v uint32) error
	SetUint32LE(index int, v uint32) error

	ReadInt32() (int32, error)
	ReadInt32LE() (int32, error)
	WriteInt32(v int32) error
	WriteInt32LE(v int32) error
	GetInt32(index int) (int32, error)
	GetInt32LE(index int) (int32, error)
	SetInt32(index int, v int32) error
	SetInt32LE(index int, v int32) error

	ReadUint64() (uint64, error)
	ReadUint64LE() (uint64, error)
	WriteUint64(v uint64) error
	WriteUint64LE(v uint64) error
	GetUint64(index int) (uint64, error)
	GetUint64LE(index int) (uint64, error)
	SetUint64(index int, v uint64) error
	SetUint64LE(index int, v uint64) error

	ReadInt64() (int64, error)
	ReadInt64LE() (int64, error)
	WriteInt64(v int64) error
	WriteInt64LE(v int64) error
	GetInt64(index int) (int64, error)
	GetInt64LE(index int) (int64, error)
	SetInt64(index int, v int64) error
	SetInt64LE(index int, v int64) error

	ReadFloat32() (float32, error)
	ReadFloat32LE() (float32, error)
	WriteFloat32(v float32) error
	WriteFloat32LE(v float32) error
	GetFloat32(index int) (float32, error)
	GetFloat32LE(index int) (float32, error)
	SetFloat32(index int, v float32) error
	SetFloat32LE(index int, v float32) error

	ReadFloat64() (float64, error)
	ReadFloat64LE() (float64, error)
	WriteFloat64(v float64) error
	WriteFloat64LE(v float64) error
	GetFloat64(index int) (float64, error)
	GetFloat64LE(index int) (float64, error)
	SetFloat64(index int, v float64) error
	SetFloat64LE(index int, v float64) error
}

func (b *wrappedBuffer) ReadString(n int) (string, error) {
	p := make([]byte, n)
	n, err := b.Read(p)
	return string(p[:n]), err
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
