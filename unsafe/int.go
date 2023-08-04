package unsafe

import (
	"strconv"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
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

// MarshalInt64 writes the MUS encoding (Raw) of a int64 value.
//
// Returns the number of used bytes and a Writer error.
func MarshalInt64(v int64, w muss.Writer) (n int, err error) {
	return marshalInteger64(v, w)
}

// MarshalInt32 writes the MUS encoding (Raw) of a int32 value.
//
// Returns the number of used bytes and a Writer value.
func MarshalInt32(v int32, w muss.Writer) (n int, err error) {
	return marshalInteger32(v, w)
}

// MarshalInt16 writes the MUS encoding (Raw) of a int16 value.
//
// Returns the number of used bytes and a Writer value.
func MarshalInt16(v int16, w muss.Writer) (n int, err error) {
	return marshalInteger16(v, w)
}

// MarshalInt8 writes the MUS encoding (Raw) of a int8 value.
//
// Returns the number of used bytes and a Writer value.
func MarshalInt8(v int8, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// MarshalInt writes the MUS encoding (Raw) of a int value.
//
// Returns the number of used bytes and a Writer value.
func MarshalInt(v int, w muss.Writer) (n int, err error) {
	return marshalInt(v, w)
}

// -----------------------------------------------------------------------------
// UnmarshalInt64 reads a MUS-encoded (Raw) int64 value.
//
// In addition to the int64 value, returns the number of used bytes and a Reader
// error.
func UnmarshalInt64(r muss.Reader) (v int64, n int, err error) {
	return unmarshalInteger64[int64](r)
}

// UnmarshalInt32 reads a MUS-encoded (Raw) int32 value.
//
// In addition to the int32 value, returns the number of used bytes and a Reader
// error.
func UnmarshalInt32(r muss.Reader) (v int32, n int, err error) {
	return unmarshalInteger32[int32](r)
}

// UnmarshalInt16 reads a MUS-encoded (Raw) int16 value.
//
// In addition to the int16 value, returns the number of used bytes and a Reader
// error.
func UnmarshalInt16(r muss.Reader) (v int16, n int, err error) {
	return unmarshalInteger16[int16](r)
}

// UnmarshalInt8 reads a MUS-encoded (Raw) int8 value.
//
// In addition to the int8 value, returns the number of used bytes and a Reader
// error.
func UnmarshalInt8(r muss.Reader) (v int8, n int, err error) {
	return unmarshalInteger8[int8](r)
}

// UnmarshalInt reads a MUS-encoded (Raw) int value.
//
// In addition to the int value, returns the number of used bytes and a Reader
// error.
func UnmarshalInt(r muss.Reader) (v int, n int, err error) {
	return unmarshalInt(r)
}

// -----------------------------------------------------------------------------
// SizeInt64 returns the size of a MUS-encoded (Raw) int64 value.
func SizeInt64(v int64) (n int) {
	return raw.SizeInt64(v)
}

// SizeInt32 returns the size of a MUS-encoded (Raw) int32 value.
func SizeInt32(v int32) (n int) {
	return raw.SizeInt32(v)
}

// SizeInt16 returns the size of a MUS-encoded (Raw) int16 value.
func SizeInt16(v int16) (n int) {
	return raw.SizeInt16(v)
}

// SizeInt8 returns the size of a MUS-encoded (Raw) int8 value.
func SizeInt8(v int8) (n int) {
	return raw.SizeInt8(v)
}

// SizeInt returns the size of a MUS-encoded (Raw) int value.
func SizeInt(v int) (n int) {
	return sizeInt(v)
}

// -----------------------------------------------------------------------------
// SkipInt64 skips a MUS-encoded (Raw) int64 value.
//
// Returns the number of skiped bytes and a Reader error.
func SkipInt64(r muss.Reader) (n int, err error) {
	return raw.SkipInt64(r)
}

// SkipInt32 skips a MUS-encoded (Raw) int32 value.
//
// Returns the number of skiped bytes and a Reader error.
func SkipInt32(r muss.Reader) (n int, err error) {
	return raw.SkipInt32(r)
}

// SkipInt16 skips a MUS-encoded (Raw) int16 value.
//
// Returns the number of skiped bytes and a Reader error.
func SkipInt16(r muss.Reader) (n int, err error) {
	return raw.SkipInt16(r)
}

// SkipInt8 skips a MUS-encoded (Raw) int8 value.
//
// Returns the number of skiped bytes and a Reader error.
func SkipInt8(r muss.Reader) (n int, err error) {
	return raw.SkipInt8(r)
}

// SkipInt skips a MUS-encoded (Raw) int value.
//
// / Returns the number of skiped bytes and a Reader error.
func SkipInt(r muss.Reader) (n int, err error) {
	return skipInt(r)
}

func setUpIntFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalInt = marshalInteger64[int]
		unmarshalInt = unmarshalInteger64[int]
	case 32:
		marshalInt = marshalInteger32[int]
		unmarshalInt = unmarshalInteger32[int]
	default:
		panic(com.ErrUnsupportedIntSize)
	}
	sizeInt = raw.SizeInt
	skipInt = raw.SkipInt
}
