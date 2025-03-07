package unsafe

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
)

// Byte is a byte serializer.
var Byte = byteSer{}

type byteSer struct{}

// Marshal writes an encoded byte value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s byteSer) Marshal(v byte, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// Unmarshal reads an encoded byte value.
//
// In addition to the byte value and the number of bytes read, it may also
// return a Reader error.
func (s byteSer) Unmarshal(r muss.Reader) (v byte, n int, err error) {
	return unmarshalInteger8[byte](r)
}

// Size returns the size of an encoded byte value.
func (s byteSer) Size(v byte) (n int) {
	return raw.Byte.Size(v)
}

// Skip skips an encoded byte value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s byteSer) Skip(r muss.Reader) (n int, err error) {
	return raw.Byte.Skip(r)
}
