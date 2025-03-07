package ord

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

func NewArraySer[T, V any](length int, elemSer muss.Serializer[V]) arraySer[T, V] {
	return NewArraySerWith[T, V](length, varint.PositiveInt, elemSer)
}

func NewArraySerWith[T, V any](length int, lenSer muss.Serializer[int],
	elemSer muss.Serializer[V]) arraySer[T, V] {
	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSerWith[V](lenSer, elemSer, lenVl, nil)
	)
	return arraySer[T, V]{length, sliceSer}
}

func NewValidArraySer[T, V any](length int, elemSer muss.Serializer[V],
	elemVl com.Validator[V]) arraySer[T, V] {
	return NewValidArraySerWith[T, V](length, varint.PositiveInt, elemSer, elemVl)
}

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

func (s arraySer[T, V]) Marshal(v T, w muss.Writer) (n int, err error) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return s.sliceSer.Marshal(sl, w)
}

func (s arraySer[T, V]) Unmarshal(r muss.Reader) (v T, n int, err error) {
	sl, n, err := s.sliceSer.Unmarshal(r)
	if err != nil {
		return
	}
	v = *(*T)(unsafe_mod.Pointer(unsafe_mod.SliceData(sl)))
	return
}

func (s arraySer[T, V]) Size(v T) (size int) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return s.sliceSer.Size(sl)
}

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
