package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
)

// MarshalBool writes the MUS encoding of a bool value.
//
// Returns the number of used bytes and a Writer error.
func MarshalBool(v bool, w muss.Writer) (n int, err error) {
	err = w.WriteByte(*(*byte)(unsafe_mod.Pointer(&v)))
	if err != nil {
		return
	}
	n++
	return
}

// UnmarshalBool reads a MUS-encoded bool value.
//
// In addition to the bool value, returns the number of used bytes and one of
// the com.ErrWrongFormat or Reader errors.
func UnmarshalBool(r muss.Reader) (v bool, n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	if b > 1 {
		err = com.ErrWrongFormat
		return
	}
	return *(*bool)(unsafe_mod.Pointer(&b)), 1, nil
}

// SizeBool returns the size of a MUS-encoded bool value.
func SizeBool(v bool) (n int) {
	return ord.SizeBool(v)
}

// SkipBool skips a MUS-encoded bool value.
//
// Returns the number of skiped bytes and one of the com.ErrWrongFormat or
// Reader errors.
func SkipBool(r muss.Reader) (n int, err error) {
	return ord.SkipBool(r)
}
