package varint

import (
	muss "github.com/mus-format/mus-stream-go"
	"golang.org/x/exp/constraints"
)

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

type int64Ser struct{}

// Marshal writes an encoded (Varint) int64 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (int64Ser) Marshal(v int64, w muss.Writer) (n int, err error) {
	return marshalUint(uint64(EncodeZigZag(v)), w)
}

// Unmarshal reads an encoded (Varint) int64 value.
//
// In addition to the int64 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (int64Ser) Unmarshal(r muss.Reader) (v int64, n int, err error) {
	uv, n, err := Uint64.Unmarshal(r)
	if err != nil {
		return
	}
	return int64(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int64 value.
func (int64Ser) Size(v int64) int {
	return sizeUint(uint64(EncodeZigZag(v)))
}

// Skip skips an encoded (Varint) int64 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (int64Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint64.Skip(r)
}

// -----------------------------------------------------------------------------

type int32Ser struct{}

// Marshal writes an encoded (Varint) int32 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (int32Ser) Marshal(v int32, w muss.Writer) (n int, err error) {
	return marshalUint(uint32(EncodeZigZag(v)), w)
}

// Unmarshal reads an encoded (Varint) int32 value.
//
// In addition to the int32 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (int32Ser) Unmarshal(r muss.Reader) (v int32, n int, err error) {
	uv, n, err := Uint32.Unmarshal(r)
	if err != nil {
		return
	}
	return int32(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int32 value.
func (int32Ser) Size(v int32) int {
	return Uint32.Size(uint32(EncodeZigZag(v)))
}

// Skip skips an encoded (Varint) int32 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (int32Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint32.Skip(r)
}

// -----------------------------------------------------------------------------

type int16Ser struct{}

// Marshal writes an encoded (Varint) int16 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (int16Ser) Marshal(v int16, w muss.Writer) (n int, err error) {
	return marshalUint(uint16(EncodeZigZag(v)), w)
}

// Unmarshal reads an encoded (Varint) int16 value.
//
// In addition to the int16 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (int16Ser) Unmarshal(r muss.Reader) (v int16, n int, err error) {
	uv, n, err := Uint16.Unmarshal(r)
	if err != nil {
		return
	}
	return int16(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int16 value.
func (int16Ser) Size(v int16) (size int) {
	return Uint16.Size(uint16(EncodeZigZag(v)))
}

// Skip skips an encoded (Varint) int16 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (int16Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint16.Skip(r)
}

// -----------------------------------------------------------------------------

type int8Ser struct{}

// Marshal writes an encoded (Varint) int8 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (int8Ser) Marshal(v int8, w muss.Writer) (n int, err error) {
	return marshalUint(uint8(EncodeZigZag(v)), w)
}

// Unmarshal reads an encoded (Varint) int8 value.
//
// In addition to the int8 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (int8Ser) Unmarshal(r muss.Reader) (v int8, n int, err error) {
	uv, n, err := Uint8.Unmarshal(r)
	if err != nil {
		return
	}
	return int8(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int8 value.
func (int8Ser) Size(v int8) (size int) {
	return Uint8.Size(uint8(EncodeZigZag(v)))
}

// Skip skips an encoded (Varint) int8 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (int8Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint8.Skip(r)
}

// -----------------------------------------------------------------------------

type intSer struct{}

// Marshal writes an encoded (Varint) int value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (intSer) Marshal(v int, w muss.Writer) (n int, err error) {
	return marshalUint(uint(EncodeZigZag(v)), w)
}

// Unmarshal reads an encoded (Varint) int value.
//
// In addition to the int value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (intSer) Unmarshal(r muss.Reader) (v int, n int, err error) {
	uv, n, err := Uint.Unmarshal(r)
	if err != nil {
		return
	}
	return int(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int value.
func (intSer) Size(v int) (size int) {
	return Uint.Size(uint(EncodeZigZag(v)))
}

// Skip skips an encoded (Varint) int value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (intSer) Skip(r muss.Reader) (n int, err error) {
	return Uint.Skip(r)
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
