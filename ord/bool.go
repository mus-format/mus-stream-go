package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// MarshalBool writes an encoded bool value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalBool(v bool, w muss.Writer) (n int, err error) {
	if v {
		err = w.WriteByte(1)
	} else {
		err = w.WriteByte(0)
	}
	if err != nil {
		return
	}
	n++
	return
}

// UnmarshalBool reads an encoded bool value.
//
// In addition to the bool value and the number of used bytes, it may also
// return com.ErrWrongFormat or a Reader error.
func UnmarshalBool(r muss.Reader) (v bool, n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	if b > 1 {
		return false, 1, com.ErrWrongFormat
	}
	return b == 1, 1, nil
}

// SizeBool returns the size of an encoded bool value.
func SizeBool(v bool) (size int) {
	return 1
}

// SkipBool skips an encoded bool value.
//
// In addition to the number of used bytes, it may also return
// com.ErrWrongFormat or a Reader error.
func SkipBool(r muss.Reader) (n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == 0 || b == 1 {
		return
	}
	err = com.ErrWrongFormat
	return
}
