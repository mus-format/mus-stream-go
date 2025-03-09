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
	// Int64 is an int64 serializer.
	Int64 = int64Ser{}
	// Int32 is an int32 serializer.
	Int32 = int32Ser{}
	// Int16 is an int16 serializer.
	Int16 = int16Ser{}
	// Int8 is an int8 serializer.
	Int8 = int8Ser{}
	// Int is an int serializer.
	Int = intSer{}
)

var (
	marshalInt   func(v int, w muss.Writer) (int, error)
	unmarshalInt func(r muss.Reader) (int, int, error)
	sizeInt      func(v int) int
	skipInt      func(r muss.Reader) (int, error)
)

type int64Ser struct{}

// Marshal writes an encoded (Raw) int64 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s int64Ser) Marshal(v int64, w muss.Writer) (n int, err error) {
	return marshalInteger64(v, w)
}

// Unmarshal reads an encoded (Raw) int64 value.
//
// In addition to the int64 value and the number of bytes read, it may also
// return a Reader error.
func (s int64Ser) Unmarshal(r muss.Reader) (v int64, n int, err error) {
	return unmarshalInteger64[int64](r)
}

// Size returns the size of an encoded (Raw) int64 value.
func (s int64Ser) Size(v int64) (n int) {
	return sizeNum64(v)
}

// Skip skips an encoded (Raw) int64 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s int64Ser) Skip(r muss.Reader) (n int, err error) {
	return skipInteger64(r)
}

// -----------------------------------------------------------------------------

type int32Ser struct{}

// Marshal writes an encoded (Raw) int32 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s int32Ser) Marshal(v int32, w muss.Writer) (n int, err error) {
	return marshalInteger32(v, w)
}

// Unmarshal reads an encoded (Raw) int32 value.
//
// In addition to the int32 value and the number of bytes read, it may also
// return a Reader error.
func (s int32Ser) Unmarshal(r muss.Reader) (v int32, n int, err error) {
	return unmarshalInteger32[int32](r)
}

// Size returns the size of an encoded (Raw) int32 value.
func (s int32Ser) Size(v int32) (n int) {
	return sizeNum32(v)
}

// Skip skips an encoded (Raw) int32 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func (s int32Ser) Skip(r muss.Reader) (n int, err error) {
	return skipInteger32(r)
}

// -----------------------------------------------------------------------------

type int16Ser struct{}

// Marshal writes an encoded (Raw) int16 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s int16Ser) Marshal(v int16, w muss.Writer) (n int, err error) {
	return marshalInteger16(v, w)
}

// Unmarshal reads an encoded (Raw) int16 value.
//
// In addition to the int16 value and the number of bytes read, it may also
// return a Reader error.
func (s int16Ser) Unmarshal(r muss.Reader) (v int16, n int, err error) {
	return unmarshalInteger16[int16](r)
}

// Size returns the size of an encoded (Raw) int16 value.
func (s int16Ser) Size(v int16) (n int) {
	return com.Num16RawSize
}

// Skip skips an encoded (Raw) int16 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s int16Ser) Skip(r muss.Reader) (n int, err error) {
	return skipInteger16(r)
}

// -----------------------------------------------------------------------------

type int8Ser struct{}

// Marshal writes an encoded (Raw) int8 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s int8Ser) Marshal(v int8, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// Unmarshal reads an encoded (Raw) int8 value.
//
// In addition to the int8 value and the number of bytes read, it may also
// return a Reader error.
func (s int8Ser) Unmarshal(r muss.Reader) (v int8, n int, err error) {
	return unmarshalInteger8[int8](r)
}

// Size returns the size of an encoded (Raw) int8 value.
func (s int8Ser) Size(v int8) (n int) {
	return com.Num8RawSize
}

// Skip skips an encoded (Raw) int8 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s int8Ser) Skip(r muss.Reader) (n int, err error) {
	return skipInteger8(r)
}

// -----------------------------------------------------------------------------

type intSer struct{}

// Marshal writes an encoded (Raw) int value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s intSer) Marshal(v int, w muss.Writer) (n int, err error) {
	return marshalInt(v, w)
}

// Unmarshal reads an encoded (Raw) int value.
//
// In addition to the int value and the number of bytes read, it may also
// return a Reader error.
func (s intSer) Unmarshal(r muss.Reader) (v int, n int, err error) {
	return unmarshalInt(r)
}

// Size returns the size of an encoded (Raw) int value.
func (s intSer) Size(v int) (n int) {
	return sizeInt(v)
}

// Skip skips an encoded (Raw) int value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s intSer) Skip(r muss.Reader) (n int, err error) {
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
