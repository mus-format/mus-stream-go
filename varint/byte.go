package varint

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// MarshalByte writes an encoded (Varint) byte value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalByte(v byte, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// UnmarshalByte reads an encoded (Varint) byte value.
//
// In addition to the byte value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
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
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipByte(r muss.Reader) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, r)
}
