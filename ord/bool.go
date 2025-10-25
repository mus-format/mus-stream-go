package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
)

// Bool is the bool serializer.
var Bool = boolSer{}

type boolSer struct{}

// Marshal writes an encoded bool value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (boolSer) Marshal(v bool, w mus.Writer) (n int, err error) {
	if v {
		err = w.WriteByte(1)
	} else {
		err = w.WriteByte(0)
	}
	if err != nil {
		return
	}
	n = 1
	return
}

// Unmarshal reads an encoded bool value.
//
// In addition to the bool value and the number of bytes read, it may also
// return com.ErrWrongFormat, or a Reader error.
func (boolSer) Unmarshal(r mus.Reader) (v bool, n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n = 1
	if b > 1 {
		err = com.ErrWrongFormat
		return
	}
	v = b == 1
	return
}

// Size returns the size of an encoded bool value.
func (boolSer) Size(v bool) (size int) {
	return SizeBool(v)
}

// Skip skips an encoded bool value.
//
// In addition to the number of bytes read, it may also return
// com.ErrWrongFormat, or a Reader error.
func (boolSer) Skip(r mus.Reader) (n int, err error) {
	return SkipBool(r)
}

// -----------------------------------------------------------------------------

func SizeBool(v bool) (size int) {
	return 1
}

func SkipBool(r mus.Reader) (n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n = 1
	if b > 1 {
		err = com.ErrWrongFormat
	}
	return
}
