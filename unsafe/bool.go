package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
)

// MarshalBool writes an encoded bool value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalBool(v bool, w muss.Writer) (n int, err error) {
	err = w.WriteByte(*(*byte)(unsafe_mod.Pointer(&v)))
	if err != nil {
		return
	}
	n++
	return
}

// UnmarshalBool reads an encoded bool value.
//
// In addition to the bool value and the number of used bytes, it may also
// return the com.ErrWrongFormat or a Reader error.
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

// SizeBool returns the size of an encoded bool value.
func SizeBool(v bool) (n int) {
	return ord.SizeBool(v)
}

// SkipBool skips an encoded bool value.
//
// In addition to the number of used bytes, it may also return
// com.ErrWrongFormat or a Reader error.
func SkipBool(r muss.Reader) (n int, err error) {
	return ord.SkipBool(r)
}
