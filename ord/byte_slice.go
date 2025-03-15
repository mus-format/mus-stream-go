package ord

import (
	"io"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	bslops "github.com/mus-format/mus-stream-go/options/byte_slice"
	"github.com/mus-format/mus-stream-go/varint"
)

// ByteSlice is the byte slice serializer.
var ByteSlice = NewByteSliceSer()

// NewByteSliceSer returns a new byte slice serializer. To specify a length
// validator, use NewValidByteSliceSer instead.
func NewByteSliceSer(ops ...bslops.SetOption) byteSliceSer {
	o := bslops.Options{}
	bslops.Apply(ops, &o)

	return newByteSliceSer(o)
}

// NewValidByteSliceSer returns a new valid byte slice serializer.
func NewValidByteSliceSer(ops ...bslops.SetOption) validByteSliceSer {
	o := bslops.Options{}
	bslops.Apply(ops, &o)

	var lenVl com.Validator[int]
	if o.LenVl != nil {
		lenVl = o.LenVl
	}
	return validByteSliceSer{newByteSliceSer(o), lenVl}
}

func newByteSliceSer(o bslops.Options) byteSliceSer {
	var lenSer muss.Serializer[int] = varint.PositiveInt
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return byteSliceSer{lenSer}
}

type byteSliceSer struct {
	lenSer muss.Serializer[int]
}

// Marshal writes an encoded slice value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s byteSliceSer) Marshal(v []byte, w muss.Writer) (
	n int, err error) {
	n, err = s.lenSer.Marshal(len(v), w)
	if err != nil {
		return
	}
	n1, err := w.Write(v)
	n += n1
	return
}

// Unmarshal reads an encoded slice value.

// In addition to the slice value and the number of bytes read, it may also
// return com.ErrOverflow, com.ErrNegativeLength, or a Reader error.
func (s byteSliceSer) Unmarshal(r muss.Reader) (v []byte, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var n1 int
	v = make([]byte, length)
	n1, err = io.ReadFull(r, v)
	n += n1
	return
}

// Size returns the size of an encoded slice value.
func (s byteSliceSer) Size(v []byte) (size int) {
	length := len(v)
	return s.lenSer.Size(length) + length
}

// Skip skips an encoded slice value.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow,
// com.ErrNegativeLength, or a Reader error.
func (s byteSliceSer) Skip(r muss.Reader) (n int, err error) {
	length, n, err := s.lenSer.Unmarshal(r)
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
		n += 1
	}
	return
}

// -----------------------------------------------------------------------------

type validByteSliceSer struct {
	byteSliceSer
	lenVl com.Validator[int]
}

// Unmarshal reads an encoded valid slice value.
//
// In addition to the slice value and the number of bytes read, it may also
// return com.ErrOverflow, com.ErrNegativeLength, a length validation error,
// or a Reader error.
func (s validByteSliceSer) Unmarshal(r muss.Reader) (v []byte, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	if s.lenVl != nil {
		if err = s.lenVl.Validate(length); err != nil {
			return
		}
	}
	var n1 int
	v = make([]byte, length)
	n1, err = io.ReadFull(r, v)
	n += n1
	return
}
