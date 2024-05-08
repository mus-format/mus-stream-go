package ord

import (
	"io"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalString writes the MUS encoding of a string value.
//
// Returns the number of used bytes and a Writer error.
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

// UnmarshalString reads a MUS-encoded string value.
//
// In addition to the string value, returns the number of used bytes and one of
// the com.ErrOverflow, com.ErrNegativeLength or Reader errors.
//
// UnmarshalString will panic if the length of the resulting string is too big
// for the string type.
func UnmarshalString(r muss.Reader) (v string, n int, err error) {
	return UnmarshalValidString(nil, false, r)
}

// UnmarshalValidString reads a MUS-encoded valid string value.
//
// The lenVl argument specifies the string length Validator. If it returns
// an error and skip == true UnmarshalValidString skips the remaining bytes of
// the string.
//
// In addition to the string value, returns the number of used bytes and one of
// the com.ErrOverflow, com.ErrNegativeLength, Validator or Reader errors.
//
// UnmarshalValidString will panic if the length of the resulting string is too
// big for the string type.
func UnmarshalValidString(lenVl com.Validator[int], skip bool,
	r muss.Reader) (v string, n int, err error) {
	length, n, err := varint.UnmarshalInt(r)
	if err != nil {
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
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			if skip {
				goto SkipRemainingBytes
			}
			return
		}
	}
	if length == 0 {
		return
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

// SizeString returns the size of a MUS-encoded string value.
func SizeString(v string) (n int) {
	return varint.SizeInt(len(v)) + len(v)
}

// SkipString skips a MUS-encoded string value.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow,
// mus.ErrNegativeLength or Reader errors.
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
