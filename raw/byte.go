package raw

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// MarshalByte writes an encoded (Raw) byte value.
//
// Returns the number of used bytes and a Writer error.
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
	return com.Num8RawSize
}

// SkipByte skips an encoded (Raw) byte value.
//
// In addition to the number of used bytes, it may also return a Reader error.
func SkipByte(r muss.Reader) (n int, err error) {
	return skipInteger8(r)
}
