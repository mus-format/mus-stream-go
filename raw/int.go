package raw

import (
	"strconv"

	muscom "github.com/mus-format/mus-common-go"
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

// MarshalInt64 writes the MUS encoding (Raw) of a int64. Returns the number of
// used bytes.
func MarshalInt64(v int64, w muss.Writer) (n int, err error) {
	return marshalInteger64(v, w)
}

// MarshalInt32 writes the MUS encoding (Raw) of a int32. Returns the number of
// used bytes.
func MarshalInt32(v int32, w muss.Writer) (n int, err error) {
	return marshalInteger32(v, w)
}

// MarshalInt16 writes the MUS encoding (Raw) of a int16. Returns the number of
// used bytes.
func MarshalInt16(v int16, w muss.Writer) (n int, err error) {
	return marshalInteger16(v, w)
}

// MarshalInt8 writes the MUS encoding (Raw) of a int8. Returns the number of
// used bytes.
func MarshalInt8(v int8, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// MarshalInt writes the MUS encoding (Raw) of a int. Returns the number of
// used bytes.
func MarshalInt(v int, w muss.Writer) (n int, err error) {
	return marshalInt(v, w)
}

// -----------------------------------------------------------------------------
// UnmarshalInt64 reads a MUS-encoded (Raw) int64. In addition to the int64, it
// returns the number of used bytes and an error.
func UnmarshalInt64(r muss.Reader) (v int64, n int, err error) {
	return unmarshalInteger64[int64](r)
}

// UnmarshalInt32 reads a MUS-encoded (Raw) int32. In addition to the int32, it
// returns the number of used bytes and an error.
func UnmarshalInt32(r muss.Reader) (v int32, n int, err error) {
	return unmarshalInteger32[int32](r)
}

// UnmarshalInt16 reads a MUS-encoded (Raw) int16. In addition to the int16, it
// returns the number of used bytes and an error.
func UnmarshalInt16(r muss.Reader) (v int16, n int, err error) {
	return unmarshalInteger16[int16](r)
}

// UnmarshalInt8 reads a MUS-encoded (Raw) int8. In addition to the int8, it
// returns the number of used bytes and an error.
func UnmarshalInt8(r muss.Reader) (v int8, n int, err error) {
	return unmarshalInteger8[int8](r)
}

// UnmarshalInt reads a MUS-encoded (Raw) int. In addition to the int, it
// returns the number of used bytes and an error.
func UnmarshalInt(r muss.Reader) (v int, n int, err error) {
	return unmarshalInt(r)
}

// -----------------------------------------------------------------------------
// SizeInt64 returns the size of a MUS-encoded (Raw) int64.
func SizeInt64(v int64) (n int) {
	return sizeNum64(v)
}

// SizeInt32 returns the size of a MUS-encoded (Raw) int32.
func SizeInt32(v int32) (n int) {
	return sizeNum32(v)
}

// SizeInt16 returns the size of a MUS-encoded (Raw) int16.
func SizeInt16(v int16) (n int) {
	return sizeInteger16(v)
}

// SizeInt8 returns the size of a MUS-encoded (Raw) int8.
func SizeInt8(v int8) (n int) {
	return sizeInteger8(v)
}

// SizeInt returns the size of a MUS-encoded (Raw) int.
func SizeInt(v int) (n int) {
	return sizeInt(v)
}

// -----------------------------------------------------------------------------
// SkipInt64 skips a MUS-encoded (Raw) int64. Returns the number of skiped bytes
// and an error.
func SkipInt64(r muss.Reader) (n int, err error) {
	return skipInteger64(r)
}

// SkipInt32 skips a MUS-encoded (Raw) int32. Returns the number of skiped bytes
// and an error.
func SkipInt32(r muss.Reader) (n int, err error) {
	return skipInteger32(r)
}

// SkipInt16 skips a MUS-encoded (Raw) int16. Returns the number of skiped bytes
// and an error.
func SkipInt16(r muss.Reader) (n int, err error) {
	return skipInteger16(r)
}

// SkipInt8 skips a MUS-encoded (Raw) int8. Returns the number of skiped bytes
// and an error.
func SkipInt8(r muss.Reader) (n int, err error) {
	return skipInteger8(r)
}

// SkipInt skips a MUS-encoded (Raw) int. Returns the number of skiped bytes
// and an error.
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
		panic(muscom.ErrUnsupportedIntSize)
	}
}
