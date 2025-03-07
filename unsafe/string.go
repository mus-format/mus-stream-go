package unsafe

import (
	"io"
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/varint"
)

// String is a string serializer.
var String = NewStringSerWith(varint.PositiveInt)

// NewStringSerWith returns a new string serializer with the given length
// serializer.
func NewStringSerWith(lenSer muss.Serializer[int]) stringSer {
	return stringSer{lenSer}
}

// NewValidStringSer returns a new valid string serializer with the given length
// validator.
func NewValidStringSer(lenVl com.Validator[int]) validStringSer {
	return NewValidStringSerWith(varint.PositiveInt, lenVl)
}

// NewValidStringSerWith returns a new valid string serializer with the given
// length serializer and length validator.
func NewValidStringSerWith(lenSer muss.Serializer[int],
	lenVl com.Validator[int]) validStringSer {
	return validStringSer{NewStringSerWith(lenSer), lenVl}
}

// -----------------------------------------------------------------------------
type stringSer struct {
	lenSer muss.Serializer[int]
}

// Marshal writes an encoded string value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s stringSer) Marshal(v string, w muss.Writer) (n int,
	err error) {
	return ord.MarshalString(v, s.lenSer, w)
}

// Unmarshal reads an encoded string value.
//
// In addition to the string value and the number of bytes read, it may also
// return com.ErrNegativeLength, a length unmarshalling error, or a Reader
// error.
//
// Unmarshal will panic if the length of the resulting string is too big
// for the string type.
func (s stringSer) Unmarshal(r muss.Reader) (v string,
	n int, err error) {
	length, n, err := s.lenSer.Unmarshal(r)
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
}

// Size returns the size of an encoded string value.
func (s stringSer) Size(v string) (n int) {
	return ord.SizeString(v, s.lenSer)
}

// Skip skips an encoded string value.
//
// In addition to the number of bytes read, it may also return
// mus.ErrNegativeLength, a length unmarshalling error, or a Reader error.
func (s stringSer) Skip(r muss.Reader) (n int, err error) {
	return ord.SkipString(r, s.lenSer)
}

// -----------------------------------------------------------------------------

type validStringSer struct {
	stringSer
	lenVl com.Validator[int]
}

// Unmarshal reads an encoded valid string value.
//
// In addition to the string value and the number of bytes read, it may also
// return com.ErrNegativeLength, a length unmarshalling error, a length
// validation error, or a Reader error.
//
// Unmarshal will panic if the length of the resulting string is too big
// for the string type.
func (s validStringSer) Unmarshal(r muss.Reader) (v string, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(r)
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
	if s.lenVl != nil {
		if err = s.lenVl.Validate(length); err != nil {
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
}
