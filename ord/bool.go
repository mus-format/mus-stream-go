package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// MarshalBool writes the MUS encoding of a bool. Returns the number of
// used bytes and an error.
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

// UnmarshalBool reads a MUS-encoded bool. In addition to the bool,
// it returns the number of used bytes and an error.
//
// The error can be one of com.ErrWrongFormat or Reader error.
func UnmarshalBool(r muss.Reader) (v bool, n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == 0 {
		return
	}
	if b == 1 {
		v = true
	} else {
		err = com.ErrWrongFormat
	}
	return
}

// SizeBool returns the size of a MUS-encoded bool.
func SizeBool(v bool) (size int) {
	return 1
}

// SkipBool skips a MUS-encoded bool. Returns the number of skiped bytes
// and an error.
//
// The error can be one of com.ErrWrongFormat or read error.
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
