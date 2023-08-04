package raw

import muss "github.com/mus-format/mus-stream-go"

// MarshalByte writes the MUS encoding (Raw) of a byte value.
//
// Returns the number of used bytes and a Writer error.
func MarshalByte(v byte, w muss.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// UnmarshalByte reads a MUS-encoded (Raw) byte value.
//
// In addition to the byte value, returns the number of used bytes and a
// Reader error.
func UnmarshalByte(r muss.Reader) (v byte, n int, err error) {
	return unmarshalInteger8[byte](r)
}

// SizeByte returns the size of a MUS-encoded (Raw) byte value.
func SizeByte(v byte) (n int) {
	return sizeInteger8(v)
}

// SkipByte skips a MUS-encoded (Raw) byte value.
//
// Returns the number of skiped bytes and a Reader error.
func SkipByte(r muss.Reader) (n int, err error) {
	return skipInteger8(r)
}
