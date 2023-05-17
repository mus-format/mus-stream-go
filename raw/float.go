package raw

import (
	"math"

	mustrm "github.com/mus-format/mus-stream-go"
)

// MarshalFloat64 writes the MUS encoding (Raw) of a float64. Returns the number
// of used bytes.
func MarshalFloat64(v float64, w mustrm.Writer) (n int, err error) {
	return marshalInteger64(math.Float64bits(v), w)
}

// MarshalFloat32 writes the MUS encoding (Raw) of a float32. Returns the number
// of used bytes.
func MarshalFloat32(v float32, w mustrm.Writer) (n int, err error) {
	return marshalInteger32(math.Float32bits(v), w)
}

// -----------------------------------------------------------------------------
// UnmarshalFloat64 reads a MUS-encoded (Raw) float64. In addition to the
// float64, it returns the number of used bytes and an error.
func UnmarshalFloat64(r mustrm.Reader) (v float64, n int, err error) {
	uv, n, err := unmarshalInteger64[uint64](r)
	if err != nil {
		return
	}
	return math.Float64frombits(uv), n, nil
}

// UnmarshalFloat32 reads a MUS-encoded (Raw) float32 from bs. In addition
// to the float32, it returns the number of used bytes and an error.
func UnmarshalFloat32(r mustrm.Reader) (v float32, n int, err error) {
	uv, n, err := unmarshalInteger32[uint32](r)
	if err != nil {
		return
	}
	return math.Float32frombits(uv), n, nil
}

// -----------------------------------------------------------------------------
// SizeFloat64 returns the size of a MUS-encoded (Raw) float64.
func SizeFloat64(v float64) (n int) {
	return sizeNum64(v)
}

// SizeFloat32 returns the size of a MUS-encoded (Raw) float32.
func SizeFloat32(v float32) (n int) {
	return sizeNum32(v)
}

// -----------------------------------------------------------------------------
// SkipFloat64 skips a MUS-encoded (Raw) float64. Returns the number of skiped
// bytes and an error.
func SkipFloat64(r mustrm.Reader) (n int, err error) {
	return skipInteger64(r)
}

// SkipFloat32 skips a MUS-encoded (Raw) float32. Returns the number of skiped
// bytes and an error.
func SkipFloat32(r mustrm.Reader) (n int, err error) {
	return skipInteger32(r)
}
