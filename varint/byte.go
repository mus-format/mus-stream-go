package varint

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// MarshalByte writes the encoding (Varint) of a byte value.
//
// Returns the number of used bytes and a Writer error.
func MarshalByte(v byte, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// UnmarshalByte reads an encoded (Varint) byte value.
//
// In addition to the byte value, returns the number of used bytes and one of
// the com.ErrOverflow or Reader errors.
func UnmarshalByte(r muss.Reader) (v byte, n int, err error) {
	return unmarshalUint[byte](com.Uint8MaxVarintLen, com.Uint8MaxLastByte,
		r)
}

// SizeByte returns the size of an encoded (Varint) byte value.
func SizeByte(v byte) (size int) {
	return sizeUint(v)
}

// SkipByte skips an encoded (Varint) byte value.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow or Reader
// errors.
func SkipByte(r muss.Reader) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, r)
}
