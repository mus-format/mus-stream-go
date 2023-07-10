package raw

import (
	"strconv"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

func init() {
	setUpUintFuncs(strconv.IntSize)
}

var (
	marshalUint   func(v uint, w muss.Writer) (int, error)
	unmarshalUint func(r muss.Reader) (uint, int, error)
	sizeUint      func(v uint) int
	skipUint      func(r muss.Reader) (int, error)
)

// MarshalUint64 writes the MUS encoding (Raw) of a uint64. Returns the number
// of used bytes and an error.
func MarshalUint64(v uint64, w muss.Writer) (n int, err error) {
	return marshalInteger64(v, w)
}

// MarshalUint32 writes the MUS encoding (Raw) of a uint32. Returns the number
// of used bytes and an error.
func MarshalUint32(v uint32, w muss.Writer) (n int, err error) {
	return marshalInteger32(v, w)
}

// MarshalUint16 writes the MUS encoding (Raw) of a uint16. Returns the number
// of used bytes and an error.
func MarshalUint16(v uint16, w muss.Writer) (n int, err error) {
	return marshalInteger16(v, w)
}

// MarshalUint8 writes the MUS encoding (Raw) of a uint8. Returns the number of
// used bytes and an error.
func MarshalUint8(v uint8, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// MarshalUint writes the MUS encoding (Raw) of a uint. Returns the number of
// used bytes and an error.
func MarshalUint(v uint, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// -----------------------------------------------------------------------------
// UnmarshalUint64 reads a MUS-encoded (Raw) uint64. In addition to the uint64,
// it returns the number of used bytes and an error.
func UnmarshalUint64(r muss.Reader) (v uint64, n int, err error) {
	return unmarshalInteger64[uint64](r)
}

// UnmarshalUint32 reads a MUS-encoded (Raw) uint32. In addition to the uint32,
// it returns the number of used bytes and an error.
func UnmarshalUint32(r muss.Reader) (v uint32, n int, err error) {
	return unmarshalInteger32[uint32](r)
}

// UnmarshalUint16 reads a MUS-encoded (Raw) uint16. In addition to the uint16,
// it returns the number of used bytes and an error.
func UnmarshalUint16(r muss.Reader) (v uint16, n int, err error) {
	return unmarshalInteger16[uint16](r)
}

// UnmarshalUint8 reads a MUS-encoded (Raw) uint8. In addition to the uint8, it
// returns the number of used bytes and an error.
func UnmarshalUint8(r muss.Reader) (v uint8, n int, err error) {
	return unmarshalInteger8[uint8](r)
}

// UnmarshalUint reads a MUS-encoded (Raw) uint. In addition to the uint, it
// returns the number of used bytes and an error.
func UnmarshalUint(r muss.Reader) (v uint, n int, err error) {
	return unmarshalUint(r)
}

// -----------------------------------------------------------------------------
// SizeUint64 returns the size of a MUS-encoded (Raw) uint64.
func SizeUint64(v uint64) (n int) {
	return sizeNum64(v)
}

// SizeUint32 returns the size of a MUS-encoded (Raw) uint32.
func SizeUint32(v uint32) (n int) {
	return sizeNum32(v)
}

// SizeUint16 returns the size of a MUS-encoded (Raw) uint16.
func SizeUint16(v uint16) (n int) {
	return sizeInteger16(v)
}

// SizeUint8 returns the size of a MUS-encoded (Raw) uint8.
func SizeUint8(v uint8) (n int) {
	return sizeInteger8(v)
}

// SizeUint returns the size of a MUS-encoded (Raw) uint.
func SizeUint(v uint) (n int) {
	return sizeUint(v)
}

// -----------------------------------------------------------------------------
// SkipUint64 skips a MUS-encoded (Raw) uint64. Returns the number of skiped
// bytes and an error.
func SkipUint64(r muss.Reader) (n int, err error) {
	return skipInteger64(r)
}

// SkipUint32 skips a MUS-encoded (Raw) uint32. Returns the number of skiped
// bytes and an error.
func SkipUint32(r muss.Reader) (n int, err error) {
	return skipInteger32(r)
}

// SkipUint16 skips a MUS-encoded (Raw) uint16. Returns the number of skiped
// bytes and an error.
func SkipUint16(r muss.Reader) (n int, err error) {
	return skipInteger16(r)
}

// SkipUint8 skips a MUS-encoded (Raw) uint8. Returns the number of skiped bytes
// and an error.
func SkipUint8(r muss.Reader) (n int, err error) {
	return skipInteger8(r)
}

// SkipUint skips a MUS-encoded (Raw) uint. Returns the number of skiped bytes
// and an error.
func SkipUint(r muss.Reader) (n int, err error) {
	return skipUint(r)
}

func setUpUintFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalUint = marshalInteger64[uint]
		unmarshalUint = unmarshalInteger64[uint]
		sizeUint = sizeNum64[uint]
		skipUint = skipInteger64
	case 32:
		marshalUint = marshalInteger32[uint]
		unmarshalUint = unmarshalInteger32[uint]
		sizeUint = sizeNum32[uint]
		skipUint = skipInteger32
	default:
		panic(com.ErrUnsupportedIntSize)
	}
}
