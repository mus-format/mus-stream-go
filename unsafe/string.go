package unsafe

import (
	"io"
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalString writes the MUS encoding of a string value.
//
// The lenM argument specifies the Marshaller for the string length.
//
// Returns the number of used bytes and a Writer error.
func MarshalString(v string, lenM muss.Marshaller[int], w muss.Writer) (n int,
	err error) {
	length := len(v)
	if lenM == nil {
		n, err = varint.MarshalPositiveInt(length, w)
	} else {
		n, err = lenM.MarshalMUS(length, w)
	}
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
// The lenU argument specifies the Unmarshaller for the string length.
//
// In addition to the string, returns the number of used bytes and one of the
// com.ErrOverflow, com.ErrNegativeLength or Reader errors.
//
// UnmarshalString will panic if the length of the resulting string is too big
// for the string type.
func UnmarshalString(lenU muss.Unmarshaller[int], r muss.Reader) (v string,
	n int, err error) {
	return UnmarshalValidString(lenU, nil, false, r)
}

// UnmarshalValidString reads a MUS-encoded valid string value.
//
// The lenU argument specifies the Unmarshaller for the string length.
// The lenVl argument specifies the string length Validator. If it returns
// an error and skip == true UnmarshalValidString skips the remaining string
// bytes.
//
// In addition to the string value, returns the number of used bytes and one
// of the com.ErrOverflow, com.ErrNegativeLength, Validator or Reader errors.
//
// UnmarshalValidString will panic if the length of the resulting string is too
// big for the string type.
func UnmarshalValidString(lenU muss.Unmarshaller[int],
	lenVl com.Validator[int],
	skip bool,
	r muss.Reader,
) (v string, n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(r)
	} else {
		length, n, err = lenU.UnmarshalMUS(r)
	}
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
	v = unsafe_mod.String(&c[0], len(c))
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
//
// The lenS argument specifies the Sizer for the string length.
func SizeString(v string, lenS muss.Sizer[int]) (n int) {
	return ord.SizeString(v, lenS)
}

// SkipString skips a MUS-encoded string value.
//
// The lenU argument specifies the Unmarshaller for the string length.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow,
// mus.ErrNegativeLength or Reader errors.
func SkipString(lenU muss.Unmarshaller[int], r muss.Reader) (n int, err error) {
	return ord.SkipString(lenU, r)
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
