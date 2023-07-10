package ord

import (
	"io"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalString writes the MUS encoding of a string. Returns the number of
// used bytes.
func MarshalString(v string, w muss.Writer) (n int, err error) {
	n, err = varint.MarshalInt(len(v), w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = w.WriteString(v)
	n += n1
	return
}

// UnmarshalString reads a MUS-encoded string. In addition to the string, it
// returns the number of used bytes and an error.
//
// The error can be one of com.ErrOverflow or com.ErrNegativeLength.
//
// It will panic if the length of the resulting string is too long.
func UnmarshalString(r muss.Reader) (v string, n int, err error) {
	return UnmarshalValidString(nil, false, r)
}

// UnmarshalValidString reads a MUS-encoded valid string. In addition to the
// string, it returns the number of used bytes and an error.
//
// The maxLength argument specifies the string length Validator. If it returns
// an error and skip == true UnmarshalValidString skips the remaining bytes of
// the string.
//
// The error returned by UnmarshalValidString can be one of com.ErrOverflow,
// com.ErrNegativeLength, a Validator or Reader error.
//
// It will panic if there is no maxLength validator and the length of the
// resulting string is too long.
func UnmarshalValidString(maxLength com.Validator[int], skip bool,
	r muss.Reader) (v string, n int, err error) {
	length, n, err := varint.UnmarshalInt(r)
	if err != nil || length == 0 {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var (
		c  []byte
		n1 int
	)
	if maxLength != nil {
		if err = maxLength.Validate(length); err != nil {
			if skip {
				goto SkipRemainingBytes
			}
			return
		}
	}
	c = make([]byte, length)
	n1, err = io.ReadFull(r, c)
	n += n1
	if err != nil {
		return
	}
	v = string(c)
	return
SkipRemainingBytes:
	n1, err1 := skipRemainingString(length, r)
	n += n1
	if err1 != nil {
		err = err1
	}
	return
}

// SizeString returns the size of a MUS-encoded string.
func SizeString(v string) (n int) {
	return varint.SizeInt(len(v)) + len(v)
}

// SkipString skips a MUS-encoded string. Returns the number of skiped bytes
// and an error.
//
// The error can be one of com.ErrOverflow, mus.ErrNegativeLength or a
// Reader error.
func SkipString(r muss.Reader) (n int, err error) {
	length, n, err := varint.UnmarshalInt(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	for i := 0; i < length; i++ {
		_, err = r.ReadByte()
		if err != nil {
			return
		}
		n++
	}
	return
}

func skipRemainingString(lenght int, r muss.Reader) (n int, err error) {
	for i := 0; i < lenght; i++ {
		_, err = r.ReadByte()
		if err != nil {
			return
		}
		n++
	}
	return
}
