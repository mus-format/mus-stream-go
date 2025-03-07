package ord

import (
	com "github.com/mus-format/common-go"
	mus "github.com/mus-format/mus-stream-go"
)

// NewPtrSer returns a new pointer serializer with the given base serializer.
func NewPtrSer[T any](baseSer mus.Serializer[T]) ptrSer[T] {
	return ptrSer[T]{baseSer}
}

type ptrSer[T any] struct {
	baseSer mus.Serializer[T]
}

// Marshal writes an encoded pointer.
//
// In addition to the number of bytes written, it may also return a base type
// marshalling error, or a Writer error.
func (s ptrSer[T]) Marshal(v *T, w mus.Writer) (n int, err error) {
	if v == nil {
		err = w.WriteByte(byte(com.Nil))
		if err != nil {
			return
		}
		n++
		return
	}
	if err = w.WriteByte(byte(com.NotNil)); err != nil {
		return
	}
	n, err = s.baseSer.Marshal(*v, w)
	n += 1
	return
}

// Unmarshal reads an encoded pointer.
//
// In addition to the pointer and the number of bytes read, it may also return
// com.ErrWrongFormat, a base type unmarshalling error, or a Reader error.
func (s ptrSer[T]) Unmarshal(r mus.Reader) (v *T, n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == byte(com.Nil) {
		return
	}
	if b != byte(com.NotNil) {
		err = com.ErrWrongFormat
		return
	}
	var n1 int
	k, n1, err := s.baseSer.Unmarshal(r)
	n += n1
	if err != nil {
		return
	}
	v = &k
	return
}

// Size returns the size of an encoded pointer.
func (s ptrSer[T]) Size(v *T) (size int) {
	if v != nil {
		return 1 + s.baseSer.Size(*v)
	}
	return 1
}

// Skip skips an encoded pointer.
//
// In addition to the number of bytes skipped, it may also return com.ErrWrongFormat,
// a base type skipping error, or a Reader error.
func (s ptrSer[T]) Skip(r mus.Reader) (n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == byte(com.Nil) {
		return
	}
	if b != byte(com.NotNil) {
		err = com.ErrWrongFormat
		return
	}
	n, err = s.baseSer.Skip(r)
	n += 1
	return
}
