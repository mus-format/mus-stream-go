package varint

import (
	"math"

	muss "github.com/mus-format/mus-stream-go"
)

var (
	// Float64 is a float64 serializer.
	Float64 = float64Ser{}
	// Float32 is a float32 serializer.
	Float32 = float32Ser{}
)

// -----------------------------------------------------------------------------

type float64Ser struct{}

// Marshal writes an encoded (Varint) float64 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (float64Ser) Marshal(v float64, w muss.Writer) (n int, err error) {
	return marshalUint(math.Float64bits(v), w)
}

// Unmarshal reads an encoded (Varint) float64 value.
//
// In addition to the float64 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (float64Ser) Unmarshal(r muss.Reader) (v float64, n int, err error) {
	uv, n, err := Uint64.Unmarshal(r)
	if err != nil {
		return
	}
	return math.Float64frombits(uv), n, nil
}

// Size returns the size of an encoded (Varint) float64 value.
func (float64Ser) Size(v float64) int {
	return sizeUint(math.Float64bits(v))
}

// Skip skips an encoded (Varint) float64 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (float64Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint64.Skip(r)
}

// -----------------------------------------------------------------------------

type float32Ser struct{}

// Marshal writes an encoded (Varint) float32 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (float32Ser) Marshal(v float32, w muss.Writer) (n int, err error) {
	return marshalUint(math.Float32bits(v), w)
}

// Unmarshal reads an encoded (Varint) float32 value.
//
// In addition to the float32 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (float32Ser) Unmarshal(r muss.Reader) (v float32, n int, err error) {
	uv, n, err := Uint32.Unmarshal(r)
	if err != nil {
		return
	}
	return math.Float32frombits(uv), n, nil
}

// Size returns the size of an encoded (Varint) float32 value.
func (float32Ser) Size(v float32) int {
	return sizeUint(math.Float32bits(v))
}

// Skip skips an encoded (Varint) float32 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (float32Ser) Skip(r muss.Reader) (n int, err error) {
	return Uint32.Skip(r)
}
