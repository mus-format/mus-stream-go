package varint

import (
	muss "github.com/mus-format/mus-stream-go"
)

var (
	// PositiveInt64 is a positive int64 serializer.
	PositiveInt64 = positiveInt64Ser{}
	// PositiveInt32 is a positive int32 serializer.
	PositiveInt32 = positiveInt32Ser{}
	// PositiveInt16 is a positive int16 serializer.
	PositiveInt16 = positiveInt16Ser{}
	// PositiveInt8 is a positive int8 serializer.
	PositiveInt8 = positiveInt8Ser{}
	// PositiveInt is a positive int serializer.
	PositiveInt = positiveIntSer{}
)

// -----------------------------------------------------------------------------

type positiveInt64Ser struct{}

// Marshal writes an encoded (Varint without ZigZag) int64 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (positiveInt64Ser) Marshal(v int64, w muss.Writer) (n int, err error) {
	return marshalUint(uint64(v), w)
}

// Unmarshal reads an encoded (Varint without ZigZag) int64 value.
// It should only be used with positive values, such as string length.
//
// In addition to the int64 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (positiveInt64Ser) Unmarshal(r muss.Reader) (v int64, n int, err error) {
	uv, n, err := Uint64.Unmarshal(r)
	if err != nil {
		return
	}
	return int64(uv), n, nil
}

// Size returns the size of an encoded (Varint without ZigZag) int64 value.
// It should only be used with positive values, such as string length.
func (positiveInt64Ser) Size(v int64) int {
	return sizeUint(uint64(v))
}

// Skip skips an encoded (Varint without ZigZag) int64 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (positiveInt64Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint64.Skip(r)
}

// -----------------------------------------------------------------------------

type positiveInt32Ser struct{}

// Marshal writes an encoded (Varint without ZigZag) int32 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (positiveInt32Ser) Marshal(v int32, w muss.Writer) (n int, err error) {
	return marshalUint(uint32(v), w)
}

// Unmarshal reads an encoded (Varint without ZigZag) int32 value.
// It should only be used with positive values, such as string length.
//
// In addition to the int32 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (positiveInt32Ser) Unmarshal(r muss.Reader) (v int32, n int, err error) {
	uv, n, err := Uint32.Unmarshal(r)
	if err != nil {
		return
	}
	return int32(uv), n, nil
}

// Size returns the size of an encoded (Varint without ZigZag) int32 value.
// It should only be used with positive values, such as string length.
func (positiveInt32Ser) Size(v int32) int {
	return Uint32.Size(uint32(v))
}

// Skip skips an encoded (Varint without ZigZag) int32 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (positiveInt32Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint32.Skip(r)
}

// -----------------------------------------------------------------------------

type positiveInt16Ser struct{}

// Marshal writes an encoded (Varint without ZigZag) int16 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (positiveInt16Ser) Marshal(v int16, w muss.Writer) (n int, err error) {
	return marshalUint(uint16(v), w)
}

// Unmarshal reads an encoded (Varint without ZigZag) int16 value.
// It should only be used with positive values, such as string length.
//
// In addition to the int16 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (positiveInt16Ser) Unmarshal(r muss.Reader) (v int16, n int, err error) {
	uv, n, err := Uint16.Unmarshal(r)
	if err != nil {
		return
	}
	return int16(uv), n, nil
}

// Size returns the size of an encoded (Varint without ZigZag) int16 value.
// It should only be used with positive values, such as string length.
func (positiveInt16Ser) Size(v int16) (size int) {
	return Uint16.Size(uint16(v))
}

// Skip skips an encoded (Varint without ZigZag) int16 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (positiveInt16Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint16.Skip(r)
}

// -----------------------------------------------------------------------------

type positiveInt8Ser struct{}

// Marshal writes an encoded (Varint without ZigZag) int8 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (positiveInt8Ser) Marshal(v int8, w muss.Writer) (n int, err error) {
	return marshalUint(uint8(v), w)
}

// Unmarshal reads an encoded (Varint without ZigZag) int8 value.
// It should only be used with positive values, such as string length.
//
// In addition to the int8 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (positiveInt8Ser) Unmarshal(r muss.Reader) (v int8, n int, err error) {
	uv, n, err := Uint8.Unmarshal(r)
	if err != nil {
		return
	}
	return int8(uv), n, nil
}

// Size returns the size of an encoded (Varint without ZigZag) int8 value.
// It should only be used with positive values, such as string length.
func (positiveInt8Ser) Size(v int8) (size int) {
	return Uint8.Size(uint8(v))
}

// Skip skips an encoded (Varint without ZigZag) int8 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (positiveInt8Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint8.Skip(r)
}

// -----------------------------------------------------------------------------

type positiveIntSer struct{}

// Marshal writes an encoded (Varint without ZigZag) int value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (positiveIntSer) Marshal(v int, w muss.Writer) (n int, err error) {
	return marshalUint(uint(v), w)
}

// Unmarshal reads an encoded (Varint without ZigZag) int value.
// It should only be used with positive values, such as string length.
//
// In addition to the int value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (positiveIntSer) Unmarshal(r muss.Reader) (v int, n int, err error) {
	uv, n, err := Uint.Unmarshal(r)
	if err != nil {
		return
	}
	return int(uv), n, nil
}

// Size returns the size of an encoded (Varint without ZigZag) int value.
// It should only be used with positive values, such as string length.
func (positiveIntSer) Size(v int) (size int) {
	return Uint.Size(uint(v))
}

// Skip skips an encoded (Varint without ZigZag) int value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (positiveIntSer) Skip(r muss.Reader) (n int, err error) {
	return Uint.Skip(r)
}
