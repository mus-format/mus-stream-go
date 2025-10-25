package unsafe

import (
	"math"

	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
)

var (
	// Float64 is a float64 serializer.
	Float64 = float64Ser{}
	// Float32 is a float32 serializer.
	Float32 = float32Ser{}
)

type float64Ser struct{}

// Marshal writes an encoded float64 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s float64Ser) Marshal(v float64, w mus.Writer) (n int, err error) {
	return marshalInteger64(math.Float64bits(v), w)
}

// Unmarshal reads an encoded float64 value.
//
// In addition to the float64 value and the number of bytes read, it may also
// return a Reader error.
func (s float64Ser) Unmarshal(r mus.Reader) (v float64, n int, err error) {
	uv, n, err := unmarshalInteger64[uint64](r)
	if err != nil {
		return
	}
	return math.Float64frombits(uv), n, nil
}

// Size returns the size of an encoded float64 value.
func (s float64Ser) Size(v float64) (n int) {
	return raw.Float64.Size(v)
}

// Skip skips an encoded float64 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s float64Ser) Skip(r mus.Reader) (n int, err error) {
	return raw.Float64.Skip(r)
}

// -----------------------------------------------------------------------------

type float32Ser struct{}

// Marshal writes an encoded float32 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s float32Ser) Marshal(v float32, w mus.Writer) (n int, err error) {
	return marshalInteger32(math.Float32bits(v), w)
}

// Unmarshal reads an encoded float32 value.
//
// In addition to the float32 value and the number of bytes read, it may also
// return a Reader error.
func (s float32Ser) Unmarshal(r mus.Reader) (v float32, n int, err error) {
	uv, n, err := unmarshalInteger32[uint32](r)
	if err != nil {
		return
	}
	return math.Float32frombits(uv), n, nil
}

// Size returns the size of an encoded float32 value.
func (s float32Ser) Size(v float32) (n int) {
	return raw.Float32.Size(v)
}

// Skip skips an encoded float32 value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s float32Ser) Skip(r mus.Reader) (n int, err error) {
	return raw.Float32.Skip(r)
}
