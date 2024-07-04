package varint

import (
	muss "github.com/mus-format/mus-stream-go"
	"golang.org/x/exp/constraints"
)

// MarshalInt64 writes the encoding (Varint) of a int64 value.
//
// Returns the number of used bytes and a Writer error.
func MarshalInt64(v int64, w muss.Writer) (n int, err error) {
	return marshalUint(uint64(EncodeZigZag(v)), w)
}

// MarshalInt32 writes the encoding (Varint) of a int32 value.
//
// Returns the number of used bytes and a Writer error.
func MarshalInt32(v int32, w muss.Writer) (n int, err error) {
	return marshalUint(uint32(EncodeZigZag(v)), w)
}

// MarshalInt16 writes the encoding (Varint) of a int16 value.
//
// Returns the number of used bytes and a Writer error.
func MarshalInt16(v int16, w muss.Writer) (n int, err error) {
	return marshalUint(uint16(EncodeZigZag(v)), w)
}

// MarshalInt8 writes the encoding (Varint) of a int8 value.
//
// Returns the number of used bytes and a Writer error.
func MarshalInt8(v int8, w muss.Writer) (n int, err error) {
	return marshalUint(uint8(EncodeZigZag(v)), w)
}

// MarshalInt writes the encoding (Varint) of a int value.
//
// Returns the number of used bytes and a Writer error.
func MarshalInt(v int, w muss.Writer) (n int, err error) {
	return marshalUint(uint(EncodeZigZag(v)), w)
}

// UnmarshalInt64 reads an encoded (Varint) int64 value.
//
// In addition to the int64 value, returns the number of used bytes and one
// of the com.ErrOverflow or Reader errors.
func UnmarshalInt64(r muss.Reader) (v int64, n int, err error) {
	uv, n, err := UnmarshalUint64(r)
	if err != nil {
		return
	}
	return int64(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt32 reads an encoded (Varint) int32 value.
//
// In addition to the int32 value, returns the number of used bytes and one
// of the com.ErrOverflow or Reader errors.
func UnmarshalInt32(r muss.Reader) (v int32, n int, err error) {
	uv, n, err := UnmarshalUint32(r)
	if err != nil {
		return
	}
	return int32(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt16 reads an encoded (Varint) int16 value.
//
// In addition to the int16 value, returns the number of used bytes and one
// of the com.ErrOverflow or Reader errors.
func UnmarshalInt16(r muss.Reader) (v int16, n int, err error) {
	uv, n, err := UnmarshalUint16(r)
	if err != nil {
		return
	}
	return int16(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt8 reads an encoded (Varint) int8 value.
//
// In addition to the int8 value, returns the number of used bytes and one of
// the com.ErrOverflow or Reader errors.
func UnmarshalInt8(r muss.Reader) (v int8, n int, err error) {
	uv, n, err := UnmarshalUint8(r)
	if err != nil {
		return
	}
	return int8(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt reads an encoded (Varint) int value.
//
// In addition to the int value, returns the number of used bytes and one of
// the com.ErrOverflow or Reader errors.
func UnmarshalInt(r muss.Reader) (v int, n int, err error) {
	uv, n, err := UnmarshalUint(r)
	if err != nil {
		return
	}
	return int(DecodeZigZag(uv)), n, nil
}

// SizeInt64 returns the size of an encoded (Varint) int64 value.
func SizeInt64(v int64) int {
	return sizeUint(uint64(EncodeZigZag(v)))
}

// SizeInt32 returns the size of an encoded (Varint) int32 value.
func SizeInt32(v int32) int {
	return SizeUint32(uint32(EncodeZigZag(v)))
}

// SizeInt16 returns the size of an encoded (Varint) int16 value.
func SizeInt16(v int16) (size int) {
	return SizeUint16(uint16(EncodeZigZag(v)))
}

// SizeInt8 returns the size of an encoded (Varint) int8 value.
func SizeInt8(v int8) (size int) {
	return SizeUint8(uint8(EncodeZigZag(v)))
}

// SizeInt returns the size of an encoded (Varint) int value.
func SizeInt(v int) (size int) {
	return SizeUint(uint(EncodeZigZag(v)))
}

// SkipInt64 skips an encoded (Varint) int64 value.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipInt64(r muss.Reader) (n int, err error) {
	return SkipUint64(r)
}

// SkipInt32 skips an encoded (Varint) int32 value.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipInt32(r muss.Reader) (n int, err error) {
	return SkipUint32(r)
}

// SkipInt16 skips an encoded (Varint) int16 value.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipInt16(r muss.Reader) (n int, err error) {
	return SkipUint16(r)
}

// SkipInt8 skips an encoded (Varint) int8 value.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipInt8(r muss.Reader) (n int, err error) {
	return SkipUint8(r)
}

// SkipInt skips an encoded (Varint) int value.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipInt(r muss.Reader) (n int, err error) {
	return SkipUint(r)
}

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
