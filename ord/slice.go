package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalSlice writes an encoded slice value.
//
// The lenM argument specifies the Marshaller for the slice length, if nil,
// varint.MarshalPositiveInt() is used.
// The m argument specifies the Marshaller for the slice elements.
//
// In addition to the number of used bytes, it may also return a Writer or
// Marshaller error.
func MarshalSlice[T any](v []T, lenM muss.Marshaller[int], m muss.Marshaller[T],
	w muss.Writer) (n int, err error) {
	if lenM == nil {
		n, err = varint.MarshalPositiveInt(len(v), w)
	} else {
		n, err = lenM.Marshal(len(v), w)
	}
	if err != nil {
		return
	}
	var n1 int
	for _, e := range v {
		n1, err = m.Marshal(e, w)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalSlice reads an encoded slice value.
//
// The lenU argument specifies the Unmarshaller for the slice length, if nil,
// varint.UnmarshalPositiveInt() is used.
// The u argument specifies the Unmarshaller for the slice elements.
//
// In addition to the slice value and the number of used bytes, it may also
// return com.ErrOverflow, com.ErrNegativeLength, a Unmarshaller or Reader error.
func UnmarshalSlice[T any](lenU muss.Unmarshaller[int], u muss.Unmarshaller[T],
	r muss.Reader) (v []T, n int, err error) {
	return UnmarshalValidSlice(lenU, nil, u, nil, nil, r)
}

// UnmarshalValidSlice reads an encoded valid slice value.
//
// The lenU argument specifies the Unmarshaller for the slice length, if nil,
// varint.UnmarshalPositiveInt() is used.
// The lenVl argument specifies the slice length Validator, arguments u,
// vl, sk - Unmarshaller, Validator and Skipper for the slice elements. If one
// of the Validators returns an error, UnmarshalValidSlice uses the Skipper to
// skip the remaining bytes of the slice. If the Skipper is nil, a validation
// error is returned immediately.
//
// In addition to the slice value and the number of used bytes, it may also
// return com.ErrOverflow, com.ErrNegativeLength, a Unmarshaller, Validator,
// Skipper or Reader error.
func UnmarshalValidSlice[T any](lenU muss.Unmarshaller[int],
	lenVl com.Validator[int],
	u muss.Unmarshaller[T],
	vl com.Validator[T],
	sk muss.Skipper,
	r muss.Reader,
) (v []T, n int, err error) {
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
		e    T
		i    int
		err1 error
	)
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make([]T, length)
	for i = 0; i < length; i++ {
		e, n1, err = u.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		if vl != nil {
			if err = vl.Validate(e); err != nil {
				goto SkipRemainingBytes
			}
		}
		v[i] = e
	}
	return
SkipRemainingBytes:
	if sk == nil {
		return
	}
	n1, err1 = skipRemainingSlice(i+1, length, sk, r)
	n += n1
	if err1 != nil {
		err = err1
	}
	return
}

// SizeSlice returns the size of an encoded slice value.
//
// The lenS argument specifies the Sizer for the slice length, if nil,
// varint.SizePositiveInt() is used.
// The s argument specifies the Sizer for the slice elements.
func SizeSlice[T any](v []T, lenS muss.Sizer[int], s muss.Sizer[T]) (size int) {
	if lenS == nil {
		size = varint.SizePositiveInt(len(v))
	} else {
		size = lenS.Size(len(v))
	}
	for i := 0; i < len(v); i++ {
		size += s.Size(v[i])
	}
	return
}

// SkipSlice skips an encoded slice value.
//
// The lenU argument specifies the Unmarshaller for the slice length, if nil,
// varint.UnmarshalPositiveInt() is used.
// The sk argument specifies the Skipper for the slice elements.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow,
// com.ErrNegativeLength, a Skipper or Reader error.
func SkipSlice(lenU muss.Unmarshaller[int], sk muss.Skipper, r muss.Reader) (
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
	n1, err := skipRemainingSlice(0, length, sk, r)
	n += n1
	return
}

func skipRemainingSlice(from int, length int, sk muss.Skipper,
	r muss.Reader) (
	n int, err error) {
	var n1 int
	for i := from; i < length; i++ {
		n1, err = sk.Skip(r)
		n += n1
		if err != nil {
			return
		}
	}
	return
}
