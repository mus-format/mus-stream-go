package unsafe

import (
	"strconv"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
)

func init() {
	setUpUintFuncs(strconv.IntSize)
}

var (
	// Uint64 is a uint64 serializer.
	Uint64 = uint64Ser{}
	// Uint32 is a uint32 serializer.
	Uint32 = uint32Ser{}
	// Uint16 is a uint16 serializer.
	Uint16 = uint16Ser{}
	// Uint8 is a uint8 serializer.
	Uint8 = uint8Ser{}
	// Uint is a uint serializer.
	Uint = uintSer{}
)

var (
	marshalUint   func(v uint, w muss.Writer) (int, error)
	unmarshalUint func(r muss.Reader) (uint, int, error)
	sizeUint      func(v uint) int
	skipUint      func(r muss.Reader) (int, error)
)

type uint64Ser struct{}

// Marshal writes an encoded (Raw) uint64 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s uint64Ser) Marshal(v uint64, w muss.Writer) (n int, err error) {
	return marshalInteger64(v, w)
}

// Unmarshal reads an encoded (Raw) uint64 value.
//
// In addition to the uint64 value and the number of bytes read, it may also
// return a Reader error.
func (s uint64Ser) Unmarshal(r muss.Reader) (v uint64, n int, err error) {
	return unmarshalInteger64[uint64](r)
}

// Size returns the size of an encoded (Raw) uint64 value.
func (s uint64Ser) Size(v uint64) (n int) {
	return raw.Uint64.Size(v)
}

// Skip skips an encoded (Raw) uint64 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s uint64Ser) Skip(r muss.Reader) (n int, err error) {
	return raw.Uint64.Skip(r)
}

// -----------------------------------------------------------------------------

type uint32Ser struct{}

// Marshal writes an encoded (Raw) uint32 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s uint32Ser) Marshal(v uint32, w muss.Writer) (n int, err error) {
	return marshalInteger32(v, w)
}

// Unmarshal reads an encoded (Raw) uint32 value.
//
// In addition to the uint32 value and the number of bytes read, it may also
// return a Reader error.
func (s uint32Ser) Unmarshal(r muss.Reader) (v uint32, n int, err error) {
	return unmarshalInteger32[uint32](r)
}

// Size returns the size of an encoded (Raw) uint32 value.
func (s uint32Ser) Size(v uint32) (n int) {
	return raw.Uint32.Size(v)
}

// Skip skips an encoded (Raw) uint32 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s uint32Ser) Skip(r muss.Reader) (n int, err error) {
	return raw.Uint32.Skip(r)
}

// -----------------------------------------------------------------------------

type uint16Ser struct{}

// Marshal writes an encoded (Raw) uint16 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s uint16Ser) Marshal(v uint16, w muss.Writer) (n int, err error) {
	return marshalInteger16(v, w)
}

// Unmarshal reads an encoded (Raw) uint16 value.
//
// In addition to the uint16 value and the number of bytes read, it may also
// return a Reader error.
func (s uint16Ser) Unmarshal(r muss.Reader) (v uint16, n int, err error) {
	return unmarshalInteger16[uint16](r)
}

// Size returns the size of an encoded (Raw) uint16 value.
func (s uint16Ser) Size(v uint16) (n int) {
	return raw.Uint16.Size(v)
}

// Skip skips an encoded (Raw) uint16 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s uint16Ser) Skip(r muss.Reader) (n int, err error) {
	return raw.Uint16.Skip(r)
}

// -----------------------------------------------------------------------------

type uint8Ser struct{}

// Marshal writes an encoded (Raw) uint8 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s uint8Ser) Marshal(v uint8, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// Unmarshal reads an encoded (Raw) uint8 value.
//
// In addition to the uint8 value and the number of bytes read, it may also
// return a Reader error.
func (s uint8Ser) Unmarshal(r muss.Reader) (v uint8, n int, err error) {
	return unmarshalInteger8[uint8](r)
}

// Size returns the size of an encoded (Raw) uint8 value.
func (s uint8Ser) Size(v uint8) (n int) {
	return raw.Uint8.Size(v)
}

// Skip skips an encoded (Raw) uint8 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s uint8Ser) Skip(r muss.Reader) (n int, err error) {
	return raw.Uint8.Skip(r)
}

// -----------------------------------------------------------------------------

type uintSer struct{}

// Marshal writes an encoded (Raw) uint value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s uintSer) Marshal(v uint, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// Unmarshal reads an encoded (Raw) uint value.
//
// In addition to the uint value and the number of bytes read, it may also
// return a Reader error.
func (s uintSer) Unmarshal(r muss.Reader) (v uint, n int, err error) {
	return unmarshalUint(r)
}

// Size returns the size of an encoded (Raw) uint value.
func (s uintSer) Size(v uint) (n int) {
	return sizeUint(v)
}

// Skip skips an encoded (Raw) uint value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s uintSer) Skip(r muss.Reader) (n int, err error) {
	return skipUint(r)
}

func setUpUintFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalUint = marshalInteger64[uint]
		unmarshalUint = unmarshalInteger64[uint]
	case 32:
		marshalUint = marshalInteger32[uint]
		unmarshalUint = unmarshalInteger32[uint]
	default:
		panic(com.ErrUnsupportedIntSize)
	}
	sizeUint = raw.Uint.Size
	skipUint = raw.Uint.Skip
}
