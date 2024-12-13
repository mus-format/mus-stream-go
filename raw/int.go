package raw

import (
	"strconv"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

func init() {
	setUpIntFuncs(strconv.IntSize)
}

var (
	marshalInt   func(v int, w muss.Writer) (int, error)
	unmarshalInt func(r muss.Reader) (int, int, error)
	sizeInt      func(v int) int
	skipInt      func(r muss.Reader) (int, error)
)

// MarshalInt64 writes an encoded (Raw) int64 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalInt64(v int64, w muss.Writer) (n int, err error) {
	return marshalInteger64(v, w)
}

// MarshalInt32 writes an encoded (Raw) int32 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalInt32(v int32, w muss.Writer) (n int, err error) {
	return marshalInteger32(v, w)
}

// MarshalInt16 writes an encoded (Raw) int16 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalInt16(v int16, w muss.Writer) (n int, err error) {
	return marshalInteger16(v, w)
}

// MarshalInt8 writes an encoded (Raw) int8 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalInt8(v int8, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// MarshalInt writes an encoded (Raw) int value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalInt(v int, w muss.Writer) (n int, err error) {
	return marshalInt(v, w)
}

// UnmarshalInt64 reads an encoded (Raw) int64 value.
//
// In addition to the int64 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalInt64(r muss.Reader) (v int64, n int, err error) {
	return unmarshalInteger64[int64](r)
}

// UnmarshalInt32 reads an encoded (Raw) int32 value.
//
// In addition to the int32 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalInt32(r muss.Reader) (v int32, n int, err error) {
	return unmarshalInteger32[int32](r)
}

// UnmarshalInt16 reads an encoded (Raw) int16 value.
//
// In addition to the int16 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalInt16(r muss.Reader) (v int16, n int, err error) {
	return unmarshalInteger16[int16](r)
}

// UnmarshalInt8 reads an encoded (Raw) int8 value.
//
// In addition to the int8 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalInt8(r muss.Reader) (v int8, n int, err error) {
	return unmarshalInteger8[int8](r)
}

// UnmarshalInt reads an encoded (Raw) int value.
//
// In addition to the int value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalInt(r muss.Reader) (v int, n int, err error) {
	return unmarshalInt(r)
}

// SizeInt64 returns the size of an encoded (Raw) int64 value.
func SizeInt64(v int64) (n int) {
	return sizeNum64(v)
}

// SizeInt32 returns the size of an encoded (Raw) int32 value.
func SizeInt32(v int32) (n int) {
	return sizeNum32(v)
}

// SizeInt16 returns the size of an encoded (Raw) int16 value.
func SizeInt16(v int16) (n int) {
	return com.Num16RawSize
}

// SizeInt8 returns the size of an encoded (Raw) int8 value.
func SizeInt8(v int8) (n int) {
	return com.Num8RawSize
}

// SizeInt returns the size of an encoded (Raw) int value.
func SizeInt(v int) (n int) {
	return sizeInt(v)
}

// SkipInt64 skips an encoded (Raw) int64 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipInt64(r muss.Reader) (n int, err error) {
	return skipInteger64(r)
}

// SkipInt32 skips an encoded (Varint) int32 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipInt32(r muss.Reader) (n int, err error) {
	return skipInteger32(r)
}

// SkipInt16 skips an encoded (Varint) int16 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipInt16(r muss.Reader) (n int, err error) {
	return skipInteger16(r)
}

// SkipInt8 skips an encoded (Varint) int8 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipInt8(r muss.Reader) (n int, err error) {
	return skipInteger8(r)
}

// SkipInt skips an encoded (Varint) int value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipInt(r muss.Reader) (n int, err error) {
	return skipInt(r)
}

func setUpIntFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalInt = marshalInteger64[int]
		unmarshalInt = unmarshalInteger64[int]
		sizeInt = sizeNum64[int]
		skipInt = skipInteger64
	case 32:
		marshalInt = marshalInteger32[int]
		unmarshalInt = unmarshalInteger32[int]
		sizeInt = sizeNum32[int]
		skipInt = skipInteger32
	default:
		panic(com.ErrUnsupportedIntSize)
	}
}
