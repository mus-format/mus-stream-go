package raw

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
)

// Byte is a byte serializer.
var Byte = byteSer{}

type byteSer struct{}

// Marshal writes an encoded (Raw) byte value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s byteSer) Marshal(v byte, w mus.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// Unmarshal reads an encoded (Raw) byte value.
//
// In addition to the byte value and the number of bytes read, it may also
// return a Reader error.
func (s byteSer) Unmarshal(r mus.Reader) (v byte, n int, err error) {
	return unmarshalInteger8[byte](r)
}

// Size returns the size of an encoded (Raw) byte value.
func (s byteSer) Size(v byte) (n int) {
	return com.Num8RawSize
}

// Skip skips an encoded (Raw) byte value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s byteSer) Skip(r mus.Reader) (n int, err error) {
	return skipInteger8(r)
}
