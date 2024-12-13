package varint

import (
	muss "github.com/mus-format/mus-stream-go"
)

// MarshalPositiveInt64 writes an encoded (Varint without ZigZag) int64 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalPositiveInt64(v int64, w muss.Writer) (n int, err error) {
	return marshalUint(uint64(v), w)
}

// MarshalPositiveInt32 writes an encoded (Varint without ZigZag) int32 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalPositiveInt32(v int32, w muss.Writer) (n int, err error) {
	return marshalUint(uint32(v), w)
}

// MarshalPositiveInt16 writes an encoded (Varint without ZigZag) int16 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalPositiveInt16(v int16, w muss.Writer) (n int, err error) {
	return marshalUint(uint16(v), w)
}

// MarshalPositiveInt8 writes an encoded (Varint without ZigZag) int8 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalPositiveInt8(v int8, w muss.Writer) (n int, err error) {
	return marshalUint(uint8(v), w)
}

// MarshalPositiveInt writes an encoded (Varint without ZigZag) int value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalPositiveInt(v int, w muss.Writer) (n int, err error) {
	return marshalUint(uint(v), w)
}

// UnmarshalPositiveInt64 reads an encoded (Varint without ZigZag) int64 value.
// It should only be used with positive values, such as string length.
//
// In addition to the int64 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalPositiveInt64(r muss.Reader) (v int64, n int, err error) {
	uv, n, err := UnmarshalUint64(r)
	if err != nil {
		return
	}
	return int64(uv), n, nil
}

// UnmarshalPositiveInt32 reads an encoded (Varint without ZigZag) int32 value.
// It should only be used with positive values, such as string length.
//
// In addition to the int32 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalPositiveInt32(r muss.Reader) (v int32, n int, err error) {
	uv, n, err := UnmarshalUint32(r)
	if err != nil {
		return
	}
	return int32(uv), n, nil
}

// UnmarshalPositiveInt16 reads an encoded (Varint without ZigZag) int16 value.
// It should only be used with positive values, such as string length.
//
// In addition to the int16 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalPositiveInt16(r muss.Reader) (v int16, n int, err error) {
	uv, n, err := UnmarshalUint16(r)
	if err != nil {
		return
	}
	return int16(uv), n, nil
}

// UnmarshalPositiveInt8 reads an encoded (Varint without ZigZag) int8 value.
// It should only be used with positive values, such as string length.
//
// In addition to the int8 value and the number of used bytes, it may also
// return .ErrOverflow or  aReader error.
func UnmarshalPositiveInt8(r muss.Reader) (v int8, n int, err error) {
	uv, n, err := UnmarshalUint8(r)
	if err != nil {
		return
	}
	return int8(uv), n, nil
}

// UnmarshalPositiveInt reads an encoded (Varint without ZigZag) int value.
// It should only be used with positive values, such as string length.
//
// In addition to the int value and the number of used bytes, it may also
// return .ErrOverflow or  aReader error.
func UnmarshalPositiveInt(r muss.Reader) (v int, n int, err error) {
	uv, n, err := UnmarshalUint(r)
	if err != nil {
		return
	}
	return int(uv), n, nil
}

// SizePositiveInt64 returns the size of an encoded (Varint without ZigZag)
// int64 value.
// It should only be used with positive values, such as string length.
func SizePositiveInt64(v int64) int {
	return sizeUint(uint64(v))
}

// SizePositiveInt32 returns the size of an encoded (Varint without ZigZag)
// int32 value.
// It should only be used with positive values, such as string length.
func SizePositiveInt32(v int32) int {
	return SizeUint32(uint32(v))
}

// SizePositiveInt16 returns the size of an encoded (Varint without ZigZag)
// int16 value.
// It should only be used with positive values, such as string length.
func SizePositiveInt16(v int16) (size int) {
	return SizeUint16(uint16(v))
}

// SizePositiveInt8 returns the size of an encoded (Varint without ZigZag) int8
// value.
// It should only be used with positive values, such as string length.
func SizePositiveInt8(v int8) (size int) {
	return SizeUint8(uint8(v))
}

// SizePositiveInt returns the size of an encoded (Varint without ZigZag) int
// value.
// It should only be used with positive values, such as string length.
func SizePositiveInt(v int) (size int) {
	return SizeUint(uint(v))
}

// SkipPositiveInt64 skips an encoded (Varint without ZigZag) int64 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipPositiveInt64(r muss.Reader) (n int, err error) {
	return SkipUint64(r)
}

// SkipPositiveInt32 skips an encoded (Varint without ZigZag) int32 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipPositiveInt32(r muss.Reader) (n int, err error) {
	return SkipUint32(r)
}

// SkipPositiveInt16 skips an encoded (Varint without ZigZag) int16 value.
// It should only be used with positive values, such as string length.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipPositiveInt16(r muss.Reader) (n int, err error) {
	return SkipUint16(r)
}

// SkipPositiveInt8 skips an encoded (Varint without ZigZag) int8 value.
// It does not use ZigZag, so it should only be used with positive values, such
// as string lengths.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipPositiveInt8(r muss.Reader) (n int, err error) {
	return SkipUint8(r)
}

// SkipPositiveInt skips an encoded (Varint without ZigZag) int value.
// It does not use ZigZag, so it should only be used with positive values, such
// as string lengths.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipPositiveInt(r muss.Reader) (n int, err error) {
	return SkipUint(r)
}
