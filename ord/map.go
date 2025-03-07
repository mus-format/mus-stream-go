package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// NewMapSer returns a new map serializer with the given key and value serializers.
func NewMapSer[T comparable, V any](
	keySer muss.Serializer[T],
	valSer muss.Serializer[V],
) mapSer[T, V] {
	return NewMapSerWith(varint.PositiveInt, keySer, valSer)
}

// NewMapSerWith returns a new map serializer with the given length serializer,
// key serializer and value serializer.
func NewMapSerWith[T comparable, V any](
	lenSer muss.Serializer[int],
	keySer muss.Serializer[T],
	valSer muss.Serializer[V],
) mapSer[T, V] {
	return mapSer[T, V]{lenSer: lenSer, keySer: keySer, valSer: valSer}
}

// NewValidMapSer returns a new valid map serializer with the given key and value
// serializers and length validator.
func NewValidMapSer[T comparable, V any](
	keySer muss.Serializer[T],
	valSer muss.Serializer[V],
	lenVl com.Validator[int],
	keyVl com.Validator[T],
	valVl com.Validator[V],
) validMapSer[T, V] {
	return NewValidMapSerWith(varint.PositiveInt, keySer, valSer, lenVl, keyVl, valVl)
}

// NewValidMapSerWith returns a new valid map serializer with the given length
// serializer, key serializer, value serializer, length validator, key validator
// and value validator.
func NewValidMapSerWith[T comparable, V any](
	lenSer muss.Serializer[int],
	keySer muss.Serializer[T],
	valSer muss.Serializer[V],
	lenVl com.Validator[int],
	keyVl com.Validator[T],
	valVl com.Validator[V],
) validMapSer[T, V] {
	return validMapSer[T, V]{NewMapSerWith(lenSer, keySer, valSer), lenVl, keyVl,
		valVl}
}

// -----------------------------------------------------------------------------

type mapSer[T comparable, V any] struct {
	lenSer muss.Serializer[int]
	keySer muss.Serializer[T]
	valSer muss.Serializer[V]
}

// Marshal writes an encoded map value.
//
// In addition to the number of bytes written, it may also return a length/key/value
// marshalling error, or a Writer error.
func (s mapSer[T, V]) Marshal(v map[T]V, w muss.Writer) (n int, err error) {
	n, err = s.lenSer.Marshal(len(v), w)
	if err != nil {
		return
	}
	var n1 int
	for k, v := range v {
		n1, err = s.keySer.Marshal(k, w)
		n += n1
		if err != nil {
			return
		}
		n1, err = s.valSer.Marshal(v, w)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// Unmarshal reads an encoded map value.
//
// In addition to the map value and the number of bytes read, it may also return
// a length/key/value unmarshalling error, or a Reader error.
func (s mapSer[T, V]) Unmarshal(r muss.Reader) (v map[T]V, n int, err error) {
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
		i  int
		k  T
		p  V
	)
	v = make(map[T]V)
	for i = 0; i < length; i++ {
		k, n1, err = s.keySer.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		p, n1, err = s.valSer.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		v[k] = p
	}
	return
}

// Size returns the size of an encoded map value.
func (s mapSer[T, V]) Size(v map[T]V) (size int) {
	size = s.lenSer.Size(len(v))
	for k, v := range v {
		size += s.keySer.Size(k)
		size += s.valSer.Size(v)
	}
	return
}

// Skip skips an encoded map value.
//
// In addition to the number of bytes read, it may also return
// com.ErrNegativeLength, a length unmarshalling error, a key/value skipping
// error, or a Reader error.
func (s mapSer[T, V]) Skip(r muss.Reader) (
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
		n1, err = s.keySer.Skip(r)
		n += n1
		if err != nil {
			return
		}
		n1, err = s.valSer.Skip(r)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// -----------------------------------------------------------------------------

type validMapSer[T comparable, V any] struct {
	mapSer[T, V]
	lenVl com.Validator[int]
	keyVl com.Validator[T]
	valVl com.Validator[V]
}

// Unmarshal reads an encoded map value.
//
// In addition to the map value and the number of bytes read, it may also return
// com.ErrNegativeLength, a length/key/value unmarshalling error, a
// length/key/value validation error, or a Reader error.
func (s validMapSer[T, V]) Unmarshal(r muss.Reader) (v map[T]V, n int,
	err error) {
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
		i  int
		k  T
		p  V
	)
	if s.lenVl != nil {
		if err = s.lenVl.Validate(length); err != nil {
			return
		}
	}
	v = make(map[T]V)
	for i = 0; i < length; i++ {
		k, n1, err = s.keySer.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		if s.keyVl != nil {
			if err = s.keyVl.Validate(k); err != nil {
				return
			}
		}
		p, n1, err = s.valSer.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		if s.valVl != nil {
			if err = s.valVl.Validate(p); err != nil {
				return
			}
		}
		v[k] = p
	}
	return
}
