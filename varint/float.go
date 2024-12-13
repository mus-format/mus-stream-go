package varint

import (
	"math"

	muss "github.com/mus-format/mus-stream-go"
)

// MarshalFloat64 writes an encoded (Varint) float64 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalFloat64(v float64, w muss.Writer) (n int, err error) {
	return MarshalUint64(math.Float64bits(v), w)
}

// MarshalFloat32 writes an encoded (Varint) float32 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalFloat32(v float32, w muss.Writer) (n int, err error) {
	return MarshalUint32(math.Float32bits(v), w)
}

// UnmarshalFloat64 reads an encoded (Varint) float64 value.
//
// In addition to the float64 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalFloat64(r muss.Reader) (v float64, n int, err error) {
	uv, n, err := UnmarshalUint64(r)
	if err != nil {
		return
	}
	v = math.Float64frombits(uv)
	return
}

// UnmarshalFloat32 reads an encoded (Varint) float32 value.
//
// In addition to the float32 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalFloat32(r muss.Reader) (v float32, n int, err error) {
	uv, n, err := UnmarshalUint32(r)
	if err != nil {
		return
	}
	v = math.Float32frombits(uv)
	return
}

// SizeFloat64 returns the size of an encoded (Varint) float64 value.
func SizeFloat64(v float64) (size int) {
	return SizeUint64(math.Float64bits(v))
}

// SizeFloat32 returns the size of an encoded (Varint) float32 value.
func SizeFloat32(v float32) (size int) {
	return SizeUint32(math.Float32bits(v))
}

// SkipFloat64 skips an encoded (Varint) float64 value.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipFloat64(r muss.Reader) (n int, err error) {
	return SkipUint64(r)
}

// SkipFloat32 skips an encoded (Varint) float32 value.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipFloat32(r muss.Reader) (n int, err error) {
	return SkipUint32(r)
}
