package varint

import (
	muscom "github.com/mus-format/mus-common-go"
	mustrm "github.com/mus-format/mus-stream-go"
)

// MarshalByte writes the MUS encoding (Varint) of a byte. Returns the
// number of used bytes and an error.
func MarshalByte(v byte, w mustrm.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// UnmarshalByte reads a MUS-encoded (Varint) byte. In addition to the byte, it
// returns the number of used bytes and an error.
//
// The error can be one of muscom.ErrOverflow or a Reader error.
func UnmarshalByte(r mustrm.Reader) (v byte, n int, err error) {
	return unmarshalUint[byte](muscom.Uint8MaxVarintLen, muscom.Uint8MaxLastByte,
		r)
}

// SizeByte returns the size of a MUS-encoded (Varint) byte.
func SizeByte(v byte) (size int) {
	return sizeUint(v)
}

// SkipByte skips a MUS-encoded (Varint) byte in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of muscom.ErrOverflow or a Reader error.
func SkipByte(r mustrm.Reader) (n int, err error) {
	return skipUint(muscom.Uint8MaxVarintLen, muscom.Uint8MaxLastByte, r)
}
