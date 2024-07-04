package varint

import (
	muss "github.com/mus-format/mus-stream-go"
)

// MarshalPositiveInt64 writes the encoding (Varint) of an int64 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes and a Writer error.
func MarshalPositiveInt64(v int64, w muss.Writer) (n int, err error) {
	return marshalUint(uint64(v), w)
}

// MarshalPositiveInt32 writes the encoding (Varint) of an int32 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes and a Writer error.
func MarshalPositiveInt32(v int32, w muss.Writer) (n int, err error) {
	return marshalUint(uint32(v), w)
}

// MarshalPositiveInt16 writes the encoding (Varint) of an int16 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes and a Writer error.
func MarshalPositiveInt16(v int16, w muss.Writer) (n int, err error) {
	return marshalUint(uint16(v), w)
}

// MarshalPositiveInt8 writes the encoding (Varint) of an int8 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes and a Writer error.
func MarshalPositiveInt8(v int8, w muss.Writer) (n int, err error) {
	return marshalUint(uint8(v), w)
}

// MarshalPositiveInt writes the encoding (Varint) of an int value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes and a Writer error.
func MarshalPositiveInt(v int, w muss.Writer) (n int, err error) {
	return marshalUint(uint(v), w)
}

// UnmarshalPositiveInt64 reads an encoded (Varint) int64 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int64 value, returns the number of used bytes and one
// of the com.ErrOverflow or Reader errors.
func UnmarshalPositiveInt64(r muss.Reader) (v int64, n int, err error) {
	uv, n, err := UnmarshalUint64(r)
	if err != nil {
		return
	}
	return int64(uv), n, nil
}

// UnmarshalPositiveInt32 reads an encoded (Varint) int32 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int32 value, returns the number of used bytes and one
// of the com.ErrOverflow or Reader errors.
func UnmarshalPositiveInt32(r muss.Reader) (v int32, n int, err error) {
	uv, n, err := UnmarshalUint32(r)
	if err != nil {
		return
	}
	return int32(uv), n, nil
}

// UnmarshalPositiveInt16 reads an encoded (Varint) int16 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int16 value, returns the number of used bytes and one
// of the com.ErrOverflow or Reader errors.
func UnmarshalPositiveInt16(r muss.Reader) (v int16, n int, err error) {
	uv, n, err := UnmarshalUint16(r)
	if err != nil {
		return
	}
	return int16(uv), n, nil
}

// UnmarshalPositiveInt8 reads an encoded (Varint) int8 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int8 value, returns the number of used bytes and one of
// the com.ErrOverflow or Reader errors.
func UnmarshalPositiveInt8(r muss.Reader) (v int8, n int, err error) {
	uv, n, err := UnmarshalUint8(r)
	if err != nil {
		return
	}
	return int8(uv), n, nil
}

// UnmarshalPositiveInt reads an encoded (Varint) int value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int value, returns the number of used bytes and one of
// the com.ErrOverflow or Reader errors.
func UnmarshalPositiveInt(r muss.Reader) (v int, n int, err error) {
	uv, n, err := UnmarshalUint(r)
	if err != nil {
		return
	}
	return int(uv), n, nil
}

// SizePositiveInt64 returns the size of an encoded (Varint) int64 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt64(v int64) int {
	return sizeUint(uint64(v))
}

// SizePositiveInt32 returns the size of an encoded (Varint) int32 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt32(v int32) int {
	return SizeUint32(uint32(v))
}

// SizePositiveInt16 returns the size of an encoded (Varint) int16 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt16(v int16) (size int) {
	return SizeUint16(uint16(v))
}

// SizePositiveInt8 returns the size of an encoded (Varint) int8 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt8(v int8) (size int) {
	return SizeUint8(uint8(v))
}

// SizePositiveInt returns the size of an encoded (Varint) int value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt(v int) (size int) {
	return SizeUint(uint(v))
}

// SkipPositiveInt64 skips an encoded (Varint) int64 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipPositiveInt64(r muss.Reader) (n int, err error) {
	return SkipUint64(r)
}

// SkipPositiveInt32 skips an encoded (Varint) int32 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipPositiveInt32(r muss.Reader) (n int, err error) {
	return SkipUint32(r)
}

// SkipPositiveInt16 skips an encoded (Varint) int16 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipPositiveInt16(r muss.Reader) (n int, err error) {
	return SkipUint16(r)
}

// SkipPositiveInt8 skips an encoded (Varint) int8 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipPositiveInt8(r muss.Reader) (n int, err error) {
	return SkipUint8(r)
}

// SkipPositiveInt skips an encoded (Varint) int value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipPositiveInt(r muss.Reader) (n int, err error) {
	return SkipUint(r)
}
