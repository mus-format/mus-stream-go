package unsafe

import (
	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
)

// MarshalByte writes the MUS encoding (Raw) of a byte. Returns the number of
// used bytes.
func MarshalByte(v byte, w mustrm.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// UnmarshalByte reads a MUS-encoded (Raw) byte. In addition to the byte, it
// returns the number of used bytes and an error.
func UnmarshalByte(r mustrm.Reader) (v byte, n int, err error) {
	return unmarshalInteger8[byte](r)
}

// SizeByte returns the size of a MUS-encoded (Raw) byte.
func SizeByte(v byte) (n int) {
	return raw.SizeByte(v)
}

// SkipByte skips a MUS-encoded (Raw) byte. Returns the number of skiped bytes
// and an error.
func SkipByte(r mustrm.Reader) (n int, err error) {
	return raw.SkipByte(r)
}
