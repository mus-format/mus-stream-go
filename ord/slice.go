package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
	slops "github.com/mus-format/mus-stream-go/options/slice"
	"github.com/mus-format/mus-stream-go/varint"
)

// NewSliceSer returns a new slice serializer with the given element serializer.
// To specify a length or element validator, use NewValidStringSer instead.
func NewSliceSer[T any](elemSer mus.Serializer[T], ops ...slops.SetOption[T]) (
	s sliceSer[T],
) {
	o := slops.Options[T]{}
	slops.Apply(ops, &o)

	return newSliceSer(elemSer, o)
}

// NewValidSliceSer returns a new valid slice serializer.
func NewValidSliceSer[T any](elemSer mus.Serializer[T],
	ops ...slops.SetOption[T],
) validSliceSer[T] {
	o := slops.Options[T]{}
	slops.Apply(ops, &o)

	var (
		lenVl  com.Validator[int]
		elemVl com.Validator[T]
	)
	if o.LenVl != nil {
		lenVl = o.LenVl
	}
	if o.ElemVl != nil {
		elemVl = o.ElemVl
	}
	return validSliceSer[T]{
		sliceSer: newSliceSer(elemSer, o),
		lenVl:    lenVl,
		elemVl:   elemVl,
	}
}

func newSliceSer[T any](elemSer mus.Serializer[T], o slops.Options[T]) (
	s sliceSer[T],
) {
	var lenSer mus.Serializer[int] = varint.PositiveInt
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return sliceSer[T]{
		elemSer: elemSer,
		lenSer:  lenSer,
	}
}

type sliceSer[T any] struct {
	lenSer  mus.Serializer[int]
	elemSer mus.Serializer[T]
}

// Marshal writes an encoded slice value.
//
// In addition to the number of bytes written, it may also return an element
// marshalling error, or a Writer error.
func (s sliceSer[T]) Marshal(v []T, w mus.Writer) (n int, err error) {
	return MarshalSlice(v, s.elemSer, s.lenSer, w)
}

// Unmarshal reads an encoded slice value.
//
// In addition to the slice value and the number of bytes read, it may also return
// com.ErrNegativeLength, a length unmarshalling error, an element unmarshalling
// error, or a Reader error.
func (s sliceSer[T]) Unmarshal(r mus.Reader) (v []T, n int, err error) {
	return UnmarshalSlice(s.elemSer, s.lenSer, r)
}

// Size returns the size of an encoded slice value.
func (s sliceSer[T]) Size(v []T) (size int) {
	return SizeSlice(v, s.elemSer, s.lenSer)
}

// Skip skips an encoded slice value.
//
// In addition to the number of bytes read, it may also return
// com.ErrNegativeLength, a length unmarshalling error, an element skipping
// error, or a Reader error.
func (s sliceSer[T]) Skip(r mus.Reader) (
	n int, err error,
) {
	return SkipSlice(s.elemSer, s.lenSer, r)
}

// -----------------------------------------------------------------------------

type validSliceSer[T any] struct {
	sliceSer[T]
	lenVl  com.Validator[int]
	elemVl com.Validator[T]
}

// Unmarshal reads an encoded valid slice value.
//
// In addition to the slice value and the number of bytes read, it may also
// return com.ErrNegativeLength, a length/element unmarshalling error, a
// length/element validation error, or a Reader error.
func (s validSliceSer[T]) Unmarshal(r mus.Reader) (v []T, n int, err error) {
	return UnmarshalValidSlice(s.elemSer, s.lenSer, s.lenVl, s.elemVl, r)
}

// -----------------------------------------------------------------------------

func MarshalSlice[T any](v []T, elemSer mus.Serializer[T],
	lenSer mus.Serializer[int], w mus.Writer) (n int, err error) {
	n, err = lenSer.Marshal(len(v), w)
	if err != nil {
		return
	}
	var n1 int
	for _, e := range v {
		n1, err = elemSer.Marshal(e, w)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func UnmarshalSlice[T any](elemSer mus.Serializer[T],
	lenSer mus.Serializer[int], r mus.Reader) (v []T, n int, err error) {
	length, n, err := lenSer.Unmarshal(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var (
		n1 int
		e  T
	)
	v = make([]T, length)
	for i := 0; i < length; i++ {
		e, n1, err = elemSer.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		v[i] = e
	}
	return
}

func SizeSlice[T any](v []T, elemSer mus.Serializer[T],
	lenSer mus.Serializer[int]) (size int) {
	size = lenSer.Size(len(v))
	for i := 0; i < len(v); i++ {
		size += elemSer.Size(v[i])
	}
	return
}

func SkipSlice[T any](elemSer mus.Serializer[T],
	lenSer mus.Serializer[int], r mus.Reader) (n int, err error) {
	length, n, err := lenSer.Unmarshal(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var n1 int
	for i := 0; i < length; i++ {
		n1, err = elemSer.Skip(r)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

func UnmarshalValidSlice[T any](elemSer mus.Serializer[T],
	lenSer mus.Serializer[int], lenVl com.Validator[int],
	elemVl com.Validator[T], r mus.Reader) (v []T, n int, err error) {
	length, n, err := lenSer.Unmarshal(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			return
		}
	}
	var (
		n1 int
		e  T
	)
	v = make([]T, length)
	for i := 0; i < length; i++ {
		e, n1, err = elemSer.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		if elemVl != nil {
			if err = elemVl.Validate(e); err != nil {
				return
			}
		}
		v[i] = e
	}
	return
}
