package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// NewSliceSer returns a new slice serializer with the given element serializer.
func NewSliceSer[T any](elemSer muss.Serializer[T]) sliceSer[T] {
	return NewSliceSerWith(varint.PositiveInt, elemSer)
}

// NewSliceSerWith returns a new slice serializer with the given length and
// element serializers.
func NewSliceSerWith[T any](lenSer muss.Serializer[int], elemSer muss.Serializer[T]) sliceSer[T] {
	return sliceSer[T]{varint.PositiveInt, elemSer}
}

// NewValidSliceSer returns a new valid slice serializer with the given element
// serializer and length validator.
func NewValidSliceSer[T any](elemSer muss.Serializer[T],
	lenVl com.Validator[int], elemVl com.Validator[T]) validSliceSer[T] {
	return NewValidSliceSerWith(varint.PositiveInt, elemSer, lenVl, elemVl)
}

// NewValidSliceSerWith returns a new valid slice serializer with the given
// length serializer, element serializer, length, and element validators.
func NewValidSliceSerWith[T any](lenSer muss.Serializer[int],
	elemSer muss.Serializer[T],
	lenVl com.Validator[int],
	elemVl com.Validator[T],
) validSliceSer[T] {
	return validSliceSer[T]{NewSliceSerWith(lenSer, elemSer), lenVl, elemVl}
}

type sliceSer[T any] struct {
	lenSer  muss.Serializer[int]
	elemSer muss.Serializer[T]
}

// Marshal writes an encoded slice value.
//
// In addition to the number of bytes written, it may also return an element
// marshalling error, or a Writer error.
func (s sliceSer[T]) Marshal(v []T, w muss.Writer) (n int, err error) {
	n, err = s.lenSer.Marshal(len(v), w)
	if err != nil {
		return
	}
	var n1 int
	for _, e := range v {
		n1, err = s.elemSer.Marshal(e, w)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// Unmarshal reads an encoded slice value.
//
// In addition to the slice value and the number of bytes read, it may also return
// com.ErrNegativeLength, a length unmarshalling error, an element unmarshalling
// error, or a Reader error.
func (s sliceSer[T]) Unmarshal(r muss.Reader) (v []T, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(r)
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
		i  int
	)
	v = make([]T, length)
	for i = 0; i < length; i++ {
		e, n1, err = s.elemSer.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		v[i] = e
	}
	return
}

// Size returns the size of an encoded slice value.
func (s sliceSer[T]) Size(v []T) (size int) {
	size = s.lenSer.Size(len(v))
	for i := 0; i < len(v); i++ {
		size += s.elemSer.Size(v[i])
	}
	return
}

// Skip skips an encoded slice value.
//
// In addition to the number of bytes read, it may also return
// com.ErrNegativeLength, a length unmarshalling error, an element skipping
// error, or a Reader error.
func (s sliceSer[T]) Skip(r muss.Reader) (
	n int, err error) {
	length, n, err := s.lenSer.Unmarshal(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var n1 int
	for i := 0; i < length; i++ {
		n1, err = s.elemSer.Skip(r)
		n += n1
		if err != nil {
			return
		}
	}
	return
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
func (s validSliceSer[T]) Unmarshal(r muss.Reader) (v []T, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(r)
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
		i  int
	)
	if s.lenVl != nil {
		if err = s.lenVl.Validate(length); err != nil {
			return
		}
	}
	v = make([]T, length)
	for i = 0; i < length; i++ {
		e, n1, err = s.elemSer.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		if s.elemVl != nil {
			if err = s.elemVl.Validate(e); err != nil {
				return
			}
		}
		v[i] = e
	}
	return
}
