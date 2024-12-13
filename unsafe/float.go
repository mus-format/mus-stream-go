package unsafe

import (
	"math"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
)

// MarshalFloat64 writes an encoded (Raw) float64 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalFloat64(v float64, w muss.Writer) (n int, err error) {
	return marshalInteger64(math.Float64bits(v), w)
}

// MarshalFloat32 writes an encoded (Raw) float32 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalFloat32(v float32, w muss.Writer) (n int, err error) {
	return marshalInteger32(math.Float32bits(v), w)
}

// UnmarshalFloat64 reads an encoded (Raw) float64 value.
//
// In addition to the fkoat64 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalFloat64(r muss.Reader) (v float64, n int, err error) {
	uv, n, err := unmarshalInteger64[uint64](r)
	if err != nil {
		return
	}
	return math.Float64frombits(uv), n, nil
}

// UnmarshalFloat32 reads an encoded (Raw) float32 value.
//
// In addition to the float32 value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalFloat32(r muss.Reader) (v float32, n int, err error) {
	uv, n, err := unmarshalInteger32[uint32](r)
	if err != nil {
		return
	}
	return math.Float32frombits(uv), n, nil
}

// SizeFloat64 returns the size of an encoded (Raw) float64 value.
func SizeFloat64(v float64) (n int) {
	return raw.SizeFloat64(v)
}

// SizeFloat32 returns the size of an encoded (Raw) float32 value.
func SizeFloat32(v float32) (n int) {
	return raw.SizeFloat32(v)
}

// SkipFloat64 skips an encoded (Raw) float64 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipFloat64(r muss.Reader) (n int, err error) {
	return raw.SkipFloat64(r)
}

// SkipFloat32 skips an encoded (Raw) float32 value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipFloat32(r muss.Reader) (n int, err error) {
	return raw.SkipFloat32(r)
}
