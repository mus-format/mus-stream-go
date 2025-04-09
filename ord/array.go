package ord

import (
	"reflect"
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	arrops "github.com/mus-format/mus-stream-go/options/array"
	slops "github.com/mus-format/mus-stream-go/options/slice"
)

// NewArraySer returns a new array serializer with the given element serializer.
// To specify a length or element validator, use NewValidArraySer instead.
//
// Panics if T is not an array type.
func NewArraySer[T, V any](elemSer muss.Serializer[V],
	ops ...arrops.SetOption[V]) (s arraySer[T, V]) {
	var (
		o      = arrops.Options[V]{}
		t      = reflect.TypeFor[T]()
		length = t.Len()
	)
	arrops.Apply(ops, &o)

	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSer[V](elemSer, slops.WithLenSer[V](o.LenSer),
			slops.WithLenValidator[V](lenVl))
	)
	return arraySer[T, V]{length, sliceSer}
}

// NewValidArraySer returns a new valid array serializer with the given element
// serializer.
//
// Panics if T is not an array type.
func NewValidArraySer[T, V any](elemSer muss.Serializer[V],
	ops ...arrops.SetOption[V]) arraySer[T, V] {
	var (
		o      = arrops.Options[V]{}
		t      = reflect.TypeFor[T]()
		length = t.Len()
	)
	arrops.Apply(ops, &o)

	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSer[V](elemSer, slops.WithLenSer[V](o.LenSer),
			slops.WithLenValidator[V](lenVl), slops.WithElemValidator(o.ElemVl))
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
