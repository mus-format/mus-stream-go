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

// MarshalUint64 writes the encoding (Raw) of a uint64 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint64(v uint64, w muss.Writer) (n int, err error) {
	return marshalInteger64(v, w)
}

// MarshalUint32 writes the encoding (Raw) of a uint32 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint32(v uint32, w muss.Writer) (n int, err error) {
	return marshalInteger32(v, w)
}

// MarshalUint16 writes the encoding (Raw) of a uint16 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint16(v uint16, w muss.Writer) (n int, err error) {
	return marshalInteger16(v, w)
}

// MarshalUint8 writes the encoding (Raw) of a uint8 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint8(v uint8, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// MarshalUint writes the encoding (Raw) of a uint value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint(v uint, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// UnmarshalUint64 reads an encoded (Raw) uint64 value.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalUint64(r muss.Reader) (v uint64, n int, err error) {
	return unmarshalInteger64[uint64](r)
}

// UnmarshalUint32 reads an encoded (Raw) uint32 value.
//
// In addition to the uint32 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalUint32(r muss.Reader) (v uint32, n int, err error) {
	return unmarshalInteger32[uint32](r)
}

// UnmarshalUint16 reads an encoded (Raw) uint16 value.
//
// In addition to the uint16 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalUint16(r muss.Reader) (v uint16, n int, err error) {
	return unmarshalInteger16[uint16](r)
}

// UnmarshalUint8 reads an encoded (Raw) uint8 value.
//
// In addition to the uint8 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalUint8(r muss.Reader) (v uint8, n int, err error) {
	return unmarshalInteger8[uint8](r)
}

// UnmarshalUint reads an encoded (Raw) uint value.
//
// In addition to the uint value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalUint(r muss.Reader) (v uint, n int, err error) {
	return unmarshalUint(r)
}

// SizeUint64 returns the size of an encoded (Raw) uint64 value.
func SizeUint64(v uint64) (n int) {
	return sizeNum64(v)
}

// SizeUint32 returns the size of an encoded (Raw) uint32 value.
func SizeUint32(v uint32) (n int) {
	return sizeNum32(v)
}

// SizeUint16 returns the size of an encoded (Raw) uint16 value.
func SizeUint16(v uint16) (n int) {
	return com.Num16RawSize
}

// SizeUint8 returns the size of an encoded (Raw) uint8 value.
func SizeUint8(v uint8) (n int) {
	return com.Num8RawSize
}

// SizeUint returns the size of an encoded (Raw) uint value.
func SizeUint(v uint) (n int) {
	return sizeUint(v)
}

// SkipUint64 skips an encoded (Raw) uint64 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipUint64(r muss.Reader) (n int, err error) {
	return skipInteger64(r)
}

// SkipUint32 skips an encoded (Raw) uint32 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipUint32(r muss.Reader) (n int, err error) {
	return skipInteger32(r)
}

// SkipUint16 skips an encoded (Raw) uint16 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipUint16(r muss.Reader) (n int, err error) {
	return skipInteger16(r)
}

// SkipUint8 skips an encoded (Raw) uint8 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipUint8(r muss.Reader) (n int, err error) {
	return skipInteger8(r)
}

// SkipUint skips an encoded (Raw) uint value.
//
// In addition to the number of used bytes, it may also return a Reader error.
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
