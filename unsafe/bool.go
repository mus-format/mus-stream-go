package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
)

// Bool is a bool serializer.
var Bool = boolSer{}

type boolSer struct{}

// Marshal writes an encoded bool value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s boolSer) Marshal(v bool, w muss.Writer) (n int, err error) {
	err = w.WriteByte(*(*byte)(unsafe_mod.Pointer(&v)))
	if err != nil {
		return
	}
	n = 1
	return
}

// Unmarshal reads an encoded bool value.
//
// In addition to the bool value and the number of bytes read, it may also
// return a Reader error.
func (s boolSer) Unmarshal(r muss.Reader) (v bool, n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n = 1
	if b > 1 {
		err = com.ErrWrongFormat
		return
	}
	v = *(*bool)(unsafe_mod.Pointer(&b))
	return
}

// Size returns the size of an encoded bool value.
func (s boolSer) Size(v bool) (n int) {
	return ord.SizeBool(v)
}

// Skip skips an encoded bool value.
//
// In addition to the number of bytes read, it may also return a Reader error.
func (s boolSer) Skip(r muss.Reader) (n int, err error) {
	return ord.SkipBool(r)
}
