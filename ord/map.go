package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalMap writes an encoded map value.
//
// The lenM argument specifies the Marshaller for the map length, if nil,
// varint.MarshalPositiveInt() is used.
// Arguments m1, m2 specify Marshallers for the keys and map values,
// respectively.
//
// In addition to the number of used bytes, it may also return a Writer or
// Marshaller error.
func MarshalMap[T comparable, V any](v map[T]V, lenM muss.Marshaller[int],
	m1 muss.Marshaller[T],
	m2 muss.Marshaller[V],
	w muss.Writer,
) (n int, err error) {
	if lenM == nil {
		n, err = varint.MarshalPositiveInt(len(v), w)
	} else {
		n, err = lenM.Marshal(len(v), w)
	}
	if err != nil {
		return
	}
	var n1 int
	for k, v := range v {
		n1, err = m1.Marshal(k, w)
		n += n1
		if err != nil {
			return
		}
		n1, err = m2.Marshal(v, w)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMap reads an encoded map value.
//
// The lenU argument specifies the Unmarshaller for the map length, if nil,
// varint.UnmarshalPositiveInt() is used.
// Arguments u1, u2 specify Unmarshallers for the keys and map values,
// respectively.
//
// In addition to the map value and the number of used bytes, it may also return
// com.ErrOverflow, a Reader or Unmarshaller error.
func UnmarshalMap[T comparable, V any](lenU muss.Unmarshaller[int],
	u1 muss.Unmarshaller[T],
	u2 muss.Unmarshaller[V],
	r muss.Reader,
) (t map[T]V, n int, err error) {
	return UnmarshalValidMap(lenU, nil, u1, u2, nil, nil, nil, nil, r)
}

// UnmarshalValidMap reads an encoded map value.
//
// The lenU argument specifies the Unmarshaller for the map length, if nil,
// varint.UnmarshalPositiveInt() is used.
// The lenVl argument specifies the map length Validator, arguments u1, u2,
// vl1, vl2, sk1, sk2 - Unmarshallers, Validators and Skippers for the keys and
// map values, respectively.
// If one of the Validators returns an error, UnmarshalValidMap uses Skippers to
// skip the remaining bytes of the map. If one of the Skippers is nil, a
// validation error is returned immediately.
//
// In addition to the map value and the number of used bytes, it may also return
// com.ErrOverflow, com.ErrNegativeLength, a Validator or Reader error.
func UnmarshalValidMap[T comparable, V any](lenU muss.Unmarshaller[int],
	lenVl com.Validator[int],
	u1 muss.Unmarshaller[T],
	u2 muss.Unmarshaller[V],
	vl1 com.Validator[T],
	vl2 com.Validator[V],
	sk1, sk2 muss.Skipper,
	r muss.Reader,
) (v map[T]V, n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(r)
	} else {
		length, n, err = lenU.Unmarshal(r)
	}
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
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make(map[T]V)
	for i = 0; i < length; i++ {
		k, n1, err = u1.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		if vl1 != nil {
			if err = vl1.Validate(k); err != nil {
				if sk2 != nil {
					n1, err1 = sk2.Skip(r)
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
		p, n1, err = u2.Unmarshal(r)
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

// SizeMap returns the size of an encoded map value.
//
// The lenS argument specifies the Sizer for the map length, if nil,
// varint.SizePositiveInt() is used.
// Arguments s1, s2 specify Sizers for the keys and map values respectively.
func SizeMap[T comparable, V any](v map[T]V, lenS muss.Sizer[int],
	s1 muss.Sizer[T],
	s2 muss.Sizer[V],
) (size int) {
	if lenS == nil {
		size = varint.SizePositiveInt(len(v))
	} else {
		size = lenS.Size(len(v))
	}
	for k, v := range v {
		size += s1.Size(k)
		size += s2.Size(v)
	}
	return
}

// SkipMap skips an encoded map value.
//
// The lenU argument specifies the Unmarshaller for the map length, if nil,
// varint.UnmarshalPositiveInt() is used.
// Arguments sk1, sk2 specify Skippers for the keys and map values,
// respectively.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow,
// com.ErrNegativeLength, a Skipper or Reader error.
func SkipMap(lenU muss.Unmarshaller[int], sk1, sk2 muss.Skipper, r muss.Reader) (
	n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(r)
	} else {
		length, n, err = lenU.Unmarshal(r)
	}
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	n1, err := skipRemainingMap(0, length, sk1, sk2, r)
	n += n1
	return
}

func skipRemainingMap(from int, length int, sk1, sk2 muss.Skipper,
	r muss.Reader) (n int, err error) {
	var n1 int
	for i := from; i < length; i++ {
		n1, err = sk1.Skip(r)
		n += n1
		if err != nil {
			return
		}
		n1, err = sk2.Skip(r)
		n += n1
		if err != nil {
			return
		}
	}
	return
}
