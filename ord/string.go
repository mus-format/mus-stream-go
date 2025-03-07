package ord

import (
	"io"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// String is the string serializer.
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
func NewValidStringSerWith(lenSer muss.Serializer[int], lenVl com.Validator[int]) validStringSer {
	return validStringSer{NewStringSerWith(lenSer), lenVl}
}

// -----------------------------------------------------------------------------

type stringSer struct {
	lenSer muss.Serializer[int]
}

// Marshal writes an encoded string value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s stringSer) Marshal(v string, w muss.Writer) (n int, err error) {
	return MarshalString(v, s.lenSer, w)
}

// Unmarshal reads an encoded string value.
//
// In addition to the string value and the number of bytes read, it may also
// return com.ErrNegativeLength, a length unmarshalling error, or a Reader error.
//
// Unmarshal will panic if the length of the resulting string is too big
// for the string type.
func (s stringSer) Unmarshal(r muss.Reader) (v string, n int, err error) {
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
	v = string(c)
	return
}

// Size returns the size of an encoded string value.
func (s stringSer) Size(v string) (n int) {
	return SizeString(v, s.lenSer)
}

// Skip skips an encoded string value.
//
// In addition to the number of bytes read, it may also return
// mus.ErrNegativeLength, or a Reader error.
func (s stringSer) Skip(r muss.Reader) (n int, err error) {
	return SkipString(r, s.lenSer)
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
// UnmarshalValidString will panic if the length of the resulting string is too
// big for the string type.
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
	v = string(c)
	return
}

// -----------------------------------------------------------------------------

func MarshalString(v string, lenSer muss.Serializer[int], w muss.Writer) (n int, err error) {
	n, err = lenSer.Marshal(len(v), w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = w.WriteString(v)
	n += n1
	return
}

func SizeString(v string, lenSer muss.Serializer[int]) (n int) {
	length := len(v)
	return lenSer.Size(length) + length
}

func SkipString(r muss.Reader, lenSer muss.Serializer[int]) (n int, err error) {
	length, n, err := lenSer.Unmarshal(r)
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
