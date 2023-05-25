package varint

import (
	muss "github.com/mus-format/mus-stream-go"
	"golang.org/x/exp/constraints"
)

// MarshalInt64 fills bs with the MUS encoding (Varint) of a int64. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt64(v int64, w muss.Writer) (n int, err error) {
	return marshalUint(uint64(EncodeZigZag(v)), w)
}

// MarshalInt32 fills bs with the MUS encoding (Varint) of a int32. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt32(v int32, w muss.Writer) (n int, err error) {
	return marshalUint(uint32(EncodeZigZag(v)), w)
}

// MarshalInt16 fills bs with the MUS encoding (Varint) of a int16. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt16(v int16, w muss.Writer) (n int, err error) {
	return marshalUint(uint16(EncodeZigZag(v)), w)
}

// MarshalInt8 fills bs with the MUS encoding (Varint) of a int8. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt8(v int8, w muss.Writer) (n int, err error) {
	return marshalUint(uint8(EncodeZigZag(v)), w)
}

// MarshalInt fills bs with the MUS encoding (Varint) of a int. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt(v int, w muss.Writer) (n int, err error) {
	return marshalUint(uint(EncodeZigZag(v)), w)
}

// -----------------------------------------------------------------------------
// UnmarshalInt64 parses a MUS-encoded (Varint) int64 from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalInt64(r muss.Reader) (v int64, n int, err error) {
	uv, n, err := UnmarshalUint64(r)
	if err != nil {
		return
	}
	return int64(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt32 parses a MUS-encoded (Varint) int32 from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalInt32(r muss.Reader) (v int32, n int, err error) {
	uv, n, err := UnmarshalUint32(r)
	if err != nil {
		return
	}
	return int32(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt16 parses a MUS-encoded (Varint) int16 from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalInt16(r muss.Reader) (v int16, n int, err error) {
	uv, n, err := UnmarshalUint16(r)
	if err != nil {
		return
	}
	return int16(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt8 parses a MUS-encoded (Varint) int8 from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalInt8(r muss.Reader) (v int8, n int, err error) {
	uv, n, err := UnmarshalUint8(r)
	if err != nil {
		return
	}
	return int8(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt parses a MUS-encoded (Varint) int from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalInt(r muss.Reader) (v int, n int, err error) {
	uv, n, err := UnmarshalUint(r)
	if err != nil {
		return
	}
	return int(DecodeZigZag(uv)), n, nil
}

// -----------------------------------------------------------------------------
// SizeInt64 returns the size of a MUS-encoded (Varint) int64.
func SizeInt64(v int64) int {
	return sizeUint(uint64(EncodeZigZag(v)))
}

// SizeInt32 returns the size of a MUS-encoded (Varint) int32.
func SizeInt32(v int32) int {
	return SizeUint32(uint32(EncodeZigZag(v)))
}

// SizeInt16 returns the size of a MUS-encoded (Varint) int16.
func SizeInt16(v int16) (size int) {
	return SizeUint16(uint16(EncodeZigZag(v)))
}

// SizeInt8 returns the size of a MUS-encoded (Varint) int8.
func SizeInt8(v int8) (size int) {
	return SizeUint8(uint8(EncodeZigZag(v)))
}

// SizeInt returns the size of a MUS-encoded (Varint) int.
func SizeInt(v int) (size int) {
	return SizeUint(uint(EncodeZigZag(v)))
}

// -----------------------------------------------------------------------------
// SkipInt64 skips a MUS-encoded (Varint) int64 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipInt64(r muss.Reader) (n int, err error) {
	return SkipUint64(r)
}

// SkipInt32 skips a MUS-encoded (Varint) int32 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipInt32(r muss.Reader) (n int, err error) {
	return SkipUint32(r)
}

// SkipInt16 skips a MUS-encoded (Varint) int16 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipInt16(r muss.Reader) (n int, err error) {
	return SkipUint16(r)
}

// SkipInt8 skips a MUS-encoded (Varint) int8 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipInt8(r muss.Reader) (n int, err error) {
	return SkipUint8(r)
}

// SkipInt skips a MUS-encoded (Varint) int in bs. Returns the number of skiped
// bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipInt(r muss.Reader) (n int, err error) {
	return SkipUint(r)
}

// -----------------------------------------------------------------------------
func EncodeZigZag[T constraints.Signed](t T) T {
	if t < 0 {
		return ^(t << 1)
	} else {
		return t << 1
	}
}

func DecodeZigZag[T constraints.Unsigned](t T) T {
	if t&1 == 1 {
		return ^(t >> 1)
	} else {
		return t >> 1
	}
}
