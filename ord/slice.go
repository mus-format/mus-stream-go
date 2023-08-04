package ord

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalSlice writes the MUS encoding of a slice value.
//
// The m argument specifies the Marshaller for the slice elements.
//
// Returns the number of used bytes and one of the Writer or Marshaller errors.
func MarshalSlice[T any](v []T, m muss.Marshaller[T], w muss.Writer) (n int,
	err error) {
	if n, err = varint.MarshalInt(len(v), w); err != nil {
		return
	}
	var n1 int
	for _, e := range v {
		n1, err = m.MarshalMUS(e, w)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalSlice reads a MUS-encoded slice value.
//
// The u argument specifies the Unmarshaller for the slice elements.
//
// In addition to the slice value, returns the number of used  bytes and one of
// the com.ErrOverflow, com.ErrNegativeLength, Unmarshaller or Reader errors.
func UnmarshalSlice[T any](u muss.Unmarshaller[T], r muss.Reader) (v []T,
	n int, err error) {
	return UnmarshalValidSlice(nil, u, nil, nil, r)
}

// UnmarshalValidSlice reads a MUS-encoded valid slice value.
//
// The maxLength argument specifies the slice length Validator, arguments u,
// vl, sk - Unmarshaller, Validator and Skipper for the slice elements. If one
// of the Validators returns an error, UnmarshalValidSlice uses the Skipper to
// skip the remaining bytes of the slice. If the Skipper is nil, a validation
// error is returned immediately.
//
// In addition to the slice value, returns the number of used bytes and one of
// the com.ErrOverflow, com.ErrNegativeLength, Unmarshaller, Validator, Skipper
// or Reader errors.
func UnmarshalValidSlice[T any](maxLength com.Validator[int],
	u muss.Unmarshaller[T],
	vl com.Validator[T],
	sk muss.Skipper,
	r muss.Reader,
) (v []T, n int, err error) {
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
		e    T
		i    int
		err1 error
	)
	if maxLength != nil {
		if err = maxLength.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make([]T, length)
	for i = 0; i < length; i++ {
		e, n1, err = u.UnmarshalMUS(r)
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

// SizeSlice returns the size of a MUS-encoded slice value.
//
// The s argument specifies the Sizer for the slice elements.
func SizeSlice[T any](v []T, s muss.Sizer[T]) (size int) {
	size = varint.SizeInt(len(v))
	for i := 0; i < len(v); i++ {
		size += s.SizeMUS(v[i])
	}
	return
}

// SkipSlice skips a MUS-encoded slice value.
//
// The sk argument specifies the Skipper for the slice elements.
//
// Returns the number of skiped bytes and one of the com.ErrOverflow,
// com.ErrNegativeLength, Skipper or Reader errors.
func SkipSlice(sk muss.Skipper, r muss.Reader) (n int, err error) {
	length, n, err := varint.UnmarshalInt(r)
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
		n1, err = sk.SkipMUS(r)
		n += n1
		if err != nil {
			return
		}
	}
	return
}
