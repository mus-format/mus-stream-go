package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	mapops "github.com/mus-format/mus-stream-go/options/map"
	"github.com/mus-format/mus-stream-go/varint"
)

// NewMapSer returns a new map serializer with the given key and value
// serializers. To specify a length, key or value validator, use NewValidMapSer
// instead.
func NewMapSer[T comparable, V any](keySer muss.Serializer[T],
	valueSer muss.Serializer[V], ops ...mapops.SetOption[T, V]) mapSer[T, V] {
	o := mapops.Options[T, V]{}
	mapops.Apply(ops, &o)

	return newMapSer(keySer, valueSer, o)
}

// NewValidMapSer returns a new valid map serializer.
func NewValidMapSer[T comparable, V any](keySer muss.Serializer[T],
	valueSer muss.Serializer[V], ops ...mapops.SetOption[T, V]) validMapSer[T, V] {
	o := mapops.Options[T, V]{}
	mapops.Apply(ops, &o)

	var (
		lenVl   com.Validator[int]
		keyVl   com.Validator[T]
		valueVl com.Validator[V]
	)
	if o.LenVl != nil {
		lenVl = o.LenVl
	}
	if o.KeyVl != nil {
		keyVl = o.KeyVl
	}
	if o.ValueVl != nil {
		valueVl = o.ValueVl
	}

	return validMapSer[T, V]{
		mapSer:  newMapSer(keySer, valueSer, o),
		lenVl:   lenVl,
		keyVl:   keyVl,
		valueVl: valueVl,
	}
}

func newMapSer[T comparable, V any](keySer muss.Serializer[T],
	valueSer muss.Serializer[V], o mapops.Options[T, V]) mapSer[T, V] {
	var lenSer muss.Serializer[int] = varint.PositiveInt
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return mapSer[T, V]{lenSer, keySer, valueSer}
}

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
	lenVl   com.Validator[int]
	keyVl   com.Validator[T]
	valueVl com.Validator[V]
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
		if s.valueVl != nil {
			if err = s.valueVl.Validate(p); err != nil {
				return
			}
		}
		v[k] = p
	}
	return
}
