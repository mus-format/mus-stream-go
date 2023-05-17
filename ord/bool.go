package ord

import (
	muscom "github.com/mus-format/mus-common-go"
	mustrm "github.com/mus-format/mus-stream-go"
)

// MarshalBool writes the MUS encoding of a bool. Returns the number of
// used bytes and an error.
func MarshalBool(v bool, w mustrm.Writer) (n int, err error) {
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
// The error can be one of muscom.ErrWrongFormat or Reader error.
func UnmarshalBool(r mustrm.Reader) (v bool, n int, err error) {
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
		err = muscom.ErrWrongFormat
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
// The error can be one of muscom.ErrWrongFormat or read error.
func SkipBool(r mustrm.Reader) (n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == 0 || b == 1 {
		return
	}
	err = muscom.ErrWrongFormat
	return
}
