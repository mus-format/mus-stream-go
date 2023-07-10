package varint

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// MarshalByte writes the MUS encoding (Varint) of a byte. Returns the
// number of used bytes and an error.
func MarshalByte(v byte, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// UnmarshalByte reads a MUS-encoded (Varint) byte. In addition to the byte, it
// returns the number of used bytes and an error.
//
// The error can be one of com.ErrOverflow or a Reader error.
func UnmarshalByte(r muss.Reader) (v byte, n int, err error) {
	return unmarshalUint[byte](com.Uint8MaxVarintLen, com.Uint8MaxLastByte,
		r)
}

// SizeByte returns the size of a MUS-encoded (Varint) byte.
func SizeByte(v byte) (size int) {
	return sizeUint(v)
}

// SkipByte skips a MUS-encoded (Varint) byte in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of com.ErrOverflow or a Reader error.
func SkipByte(r muss.Reader) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, r)
}
