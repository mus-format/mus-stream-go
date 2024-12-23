package unsafe

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
)

// MarshalByte writes an encoded (Raw) byte value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalByte(v byte, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// UnmarshalByte reads an encoded (Raw) byte value.
//
// In addition to the byte value and the number of used bytes, it may also
// return a Reader error.
func UnmarshalByte(r muss.Reader) (v byte, n int, err error) {
	return unmarshalInteger8[byte](r)
}

// SizeByte returns the size of an encoded (Raw) byte value.
func SizeByte(v byte) (n int) {
	return raw.SizeByte(v)
}

// SkipByte skips an encoded (Raw) byte value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipByte(r muss.Reader) (n int, err error) {
	return raw.SkipByte(r)
}
