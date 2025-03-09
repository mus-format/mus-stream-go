package ord

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// NewArraySer returns a new array serializer with the given array length and
// element serializer.
func NewArraySer[T, V any](length int, elemSer muss.Serializer[V]) arraySer[T, V] {
	return NewArraySerWith[T, V](length, varint.PositiveInt, elemSer)
}

// NewArraySerWith returns a new array serializer with the given array length,
// length and element serializers.
func NewArraySerWith[T, V any](length int, lenSer muss.Serializer[int],
	elemSer muss.Serializer[V]) arraySer[T, V] {
	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSerWith[V](lenSer, elemSer, lenVl, nil)
	)
	return arraySer[T, V]{length, sliceSer}
}

// NewValidArraySer returns a new valid array serializer with the given array
// length, element serializer and length validator.
func NewValidArraySer[T, V any](length int, elemSer muss.Serializer[V],
	elemVl com.Validator[V]) arraySer[T, V] {
	return NewValidArraySerWith[T, V](length, varint.PositiveInt, elemSer, elemVl)
}

// NewValidArraySerWith returns a new valid array serializer with the given
// array length, length serializer, element serializer, length, and element
// validators.
func NewValidArraySerWith[T, V any](length int, lenSer muss.Serializer[int],
	elemSer muss.Serializer[V], elemVl com.Validator[V]) arraySer[T, V] {
	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSerWith[V](lenSer, elemSer, lenVl, elemVl)
	)
	return arraySer[T, V]{length, sliceSer}
}

type arraySer[T, V any] struct {
	length   int
	sliceSer validSliceSer[V]
}

// Marshal writes an encoded array value.
//
// In addition to the number of bytes written, it may also return an element
// marshalling error, or a Writer error.
func (s arraySer[T, V]) Marshal(v T, w muss.Writer) (n int, err error) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return s.sliceSer.Marshal(sl, w)
}

// Unmarshal reads an encoded array value.
//
// In addition to the array value and the number of bytes read, it may also
// return com.ErrNegativeLength, a length unmarshalling error, an element
// unmarshalling error, or a Reader error.
func (s arraySer[T, V]) Unmarshal(r muss.Reader) (v T, n int, err error) {
	sl, n, err := s.sliceSer.Unmarshal(r)
	if err != nil {
		return
	}
	v = *(*T)(unsafe_mod.Pointer(unsafe_mod.SliceData(sl)))
	return
}

// Size returns the size of an encoded array value.
func (s arraySer[T, V]) Size(v T) (size int) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return s.sliceSer.Size(sl)
}

// Skip skips an encoded array value.
//
// In addition to the number of bytes read, it may also return
// com.ErrNegativeLength, a length unmarshalling error, an element skipping
// error, or a Reader error.
func (s arraySer[T, V]) Skip(r muss.Reader) (n int, err error) {
	return s.sliceSer.Skip(r)
}

func newLenVl(length int) com.ValidatorFn[int] {
	return func(t int) (err error) {
		if t > length {
			err = com.ErrTooLargeLength
		}
		return
	}
}
