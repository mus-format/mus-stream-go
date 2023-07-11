package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalMap writes the MUS encoding of a map. Returns the number of
// used bytes and an error.
//
// Arguments m1, m2 specify Marshallers for the keys and map values,
// respectively.
func MarshalMap[T comparable, V any](v map[T]V, m1 muss.Marshaller[T],
	m2 muss.Marshaller[V],
	w muss.Writer,
) (n int, err error) {
	if n, err = varint.MarshalInt(len(v), w); err != nil {
		return
	}
	var n1 int
	for k, v := range v {
		n1, err = m1.MarshalMUS(k, w)
		n += n1
		if err != nil {
			return
		}
		n1, err = m2.MarshalMUS(v, w)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMap reads a MUS-encoded map. In addition to the map,
// it returns the number of used bytes and an error.
//
// Arguments u1, u2 specify Unmarshallers for the keys and map values,
// respectively.
//
// The error can be one of com.ErrNegativeLength or Reader error.
func UnmarshalMap[T comparable, V any](u1 muss.Unmarshaller[T],
	u2 muss.Unmarshaller[V],
	r muss.Reader,
) (t map[T]V, n int, err error) {
	return UnmarshalValidMap(nil, u1, u2, nil, nil, nil, nil, r)
}

// UnmarshalValidMap reads a MUS-encoded map. In addition to the map,
// it returns the number of used bytes and an error.
//
// The maxLength argument specifies the map length Validator. Arguments
// u1, u2, vl1, vl2, sk1, sk2 - Unmarshallers, Validators and Skippers for the
// map keys and values, respectively.
// If one of the Validators returns an error, UnmarshalValidMap uses the
// Skippers to skip the remaining bytes of the map. If one of the Skippers is
// nil, it immediately returns a validation error.
//
// The error can be one of com.ErrOverflow, com.ErrNegativeLength,
// validation or Reader error.
func UnmarshalValidMap[T comparable, V any](maxLength com.Validator[int],
	u1 muss.Unmarshaller[T],
	u2 muss.Unmarshaller[V],
	vl1 com.Validator[T],
	vl2 com.Validator[V],
	sk1, sk2 muss.Skipper,
	r muss.Reader,
) (v map[T]V, n int, err error) {
	length, n, err := varint.UnmarshalInt(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var (
		n1   int
		i    int
		err1 error
		k    T
		p    V
	)
	if maxLength != nil {
		if err = maxLength.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make(map[T]V)
	for i = 0; i < length; i++ {
		k, n1, err = u1.UnmarshalMUS(r)
		n += n1
		if err != nil {
			return
		}
		if vl1 != nil {
			if err = vl1.Validate(k); err != nil {
				if sk2 != nil {
					n1, err1 = sk2.SkipMUS(r)
					n += n1
					if err1 != nil {
						err = err1
						return
					}
					i++
				}
				goto SkipRemainingBytes
			}
		}
		p, n1, err = u2.UnmarshalMUS(r)
		n += n1
		if err != nil {
			return
		}
		if vl2 != nil {
			if err = vl2.Validate(p); err != nil {
				i++
				goto SkipRemainingBytes
			}
		}
		v[k] = p
	}
	return
SkipRemainingBytes:
	if sk1 == nil || sk2 == nil {
		return
	}
	n1, err1 = skipRemainingMap(i, length, sk1, sk2, r)
	n += n1
	if err1 != nil {
		err = err1
	}
	return
}

// SizeMap returns the size of a MUS-encoded map.
//
// Arguments s1, s2 specify Sizers for the keys and map values respectively.
func SizeMap[T comparable, V any](v map[T]V, s1 muss.Sizer[T],
	s2 muss.Sizer[V]) (size int) {
	size += varint.SizeInt(len(v))
	for k, v := range v {
		size += s1.SizeMUS(k)
		size += s2.SizeMUS(v)
	}
	return
}

// SkipMap skips a MUS-encoded map. Returns the number of skiped bytes
// and an error.
//
// Arguments sk1, sk2 specify Skippers for the keys and map values,
// respectively.
//
// The error returned by SkipMap can be one of com.ErrOverflow,
// com.ErrNegativeLength, a Skipper or Reader error.
func SkipMap(s1, s2 muss.Skipper, r muss.Reader) (n int, err error) {
	length, n, err := varint.UnmarshalInt(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	n1, err := skipRemainingMap(0, length, s1, s2, r)
	n += n1
	return
}

func skipRemainingMap(from int, length int, sk1, sk2 muss.Skipper,
	r muss.Reader) (n int, err error) {
	var n1 int
	for i := from; i < length; i++ {
		n1, err = sk1.SkipMUS(r)
		n += n1
		if err != nil {
			return
		}
		n1, err = sk2.SkipMUS(r)
		n += n1
		if err != nil {
			return
		}
	}
	return
}
