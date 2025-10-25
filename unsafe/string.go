package unsafe

import (
	"io"
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
	strops "github.com/mus-format/mus-stream-go/options/string"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/varint"
)

// String is a string serializer.
var String = NewStringSer()

// NewStringSer returns a new string serializer. To specify a length validator,
// use NewValidStringSer instead.
func NewStringSer(ops ...strops.SetOption) stringSer {
	o := strops.Options{}
	strops.Apply(ops, &o)

	return newStringSer(o)
}

// NewValidStringSer returns a new valid string serializer.
func NewValidStringSer(ops ...strops.SetOption) validStringSer {
	o := strops.Options{}
	strops.Apply(ops, &o)
	return validStringSer{newStringSer(o), o.LenVl}
}

func newStringSer(o strops.Options) stringSer {
	var lenSer mus.Serializer[int] = varint.PositiveInt
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return stringSer{lenSer}
}

type stringSer struct {
	lenSer mus.Serializer[int]
}

// Marshal writes an encoded string value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s stringSer) Marshal(v string, w mus.Writer) (n int,
	err error,
) {
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
func (s stringSer) Unmarshal(r mus.Reader) (v string,
	n int, err error,
) {
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
func (s stringSer) Skip(r mus.Reader) (n int, err error) {
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
func (s validStringSer) Unmarshal(r mus.Reader) (v string, n int, err error) {
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
