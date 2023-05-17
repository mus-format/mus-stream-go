package unsafe

import (
	unsafe_mod "unsafe"

	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
)

// MarshalBool writes the MUS encoding of a bool. Returns the number of used
// bytes.
func MarshalBool(v bool, w mustrm.Writer) (n int, err error) {
	err = w.WriteByte(*(*byte)(unsafe_mod.Pointer(&v)))
	if err != nil {
		return
	}
	n++
	return
}

// UnmarshalBool reads a MUS-encoded bool. In addition to the bool, it returns
// the number of used bytes and an error.
//
// The error can be one of muscom.ErrWrongFormat or a Reader error.
func UnmarshalBool(r mustrm.Reader) (v bool, n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	return *(*bool)(unsafe_mod.Pointer(&b)), 1, nil
}

// SizeBool returns the size of a MUS-encoded bool.
func SizeBool(v bool) (n int) {
	return ord.SizeBool(v)
}

// SkipBool skips a MUS-encoded bool. Returns the number of skiped bytes and an
// error.
//
// The error can be one of muscom.ErrWrongFormat or Reader error.
func SkipBool(r mustrm.Reader) (n int, err error) {
	return ord.SkipBool(r)
}
