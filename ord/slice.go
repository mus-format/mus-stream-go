package ord

import (
	muscom "github.com/mus-format/mus-common-go"
	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalSlice writes the MUS encoding of a slice. Returns the number of
// used bytes and an error.
//
// The m argument specifies the Marshaler for the slice elements.
func MarshalSlice[T any](v []T, m mustrm.Marshaler[T], w mustrm.Writer) (n int,
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

// UnmarshalSlice reads a MUS-encoded slice. In addition to the slice, it
// returns the number of used bytes and an error.
//
// The u argument specifies the Unmarshaler for the slice elements.
//
// The error returned by UnmarshalSlice can be one of muscom.ErrOverflow,
// muscom.ErrNegativeLength, an Unmarshaler or Reader error.
func UnmarshalSlice[T any](u mustrm.Unmarshaler[T], r mustrm.Reader) (v []T,
	n int, err error) {
	return UnmarshalValidSlice(nil, u, nil, nil, r)
}

// UnmarshalValidSlice reads a MUS-encoded valid slice. In addition to the
// slice, it returns the number of used bytes and an error.
//
// The maxLength argument specifies the slice length Validator. Arguments u,
// vl, sk - the Unmarshaler, Validator and Skipper for slice elements. If one
// of the Validators returns an error, UnmarshalValidSlice uses the Skipper to
// skip the remaining bytes of the slice. If the value of the Skipper is nil, it
// immediately returns the validation error.
//
// The error returned by UnmarshalValidSlice can be one of muscom.ErrOverflow,
// muscom.ErrNegativeLength, an Unmarshaler, Validator, Skipper or Reader error.
func UnmarshalValidSlice[T any](maxLength muscom.Validator[int],
	u mustrm.Unmarshaler[T],
	vl muscom.Validator[T],
	sk mustrm.Skipper,
	r mustrm.Reader,
) (v []T, n int, err error) {
	length, n, err := varint.UnmarshalInt(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = muscom.ErrNegativeLength
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

// SizeSlice returns the size of a MUS-encoded slice.
//
// The s argument specifies the Sizer for the slice elements.
func SizeSlice[T any](v []T, s mustrm.Sizer[T]) (size int) {
	size = varint.SizeInt(len(v))
	for i := 0; i < len(v); i++ {
		size += s.SizeMUS(v[i])
	}
	return
}

// SkipSlice skips a MUS-encoded slice. Returns the number of skiped bytes and
// an error.
//
// The sk argument specifies the Skipper for the slice elements.
//
// The error returned by SkipSlice can be one of muscom.ErrOverflow,
// muscom.ErrNegativeLength, a Skipper or Reader error.
func SkipSlice(sk mustrm.Skipper, r mustrm.Reader) (n int, err error) {
	length, n, err := varint.UnmarshalInt(r)
	if err != nil {
		return
	}
	if length < 0 {
		err = muscom.ErrNegativeLength
		return
	}
	n1, err := skipRemainingSlice(0, length, sk, r)
	n += n1
	return
}

func skipRemainingSlice(from int, length int, sk mustrm.Skipper,
	r mustrm.Reader) (
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
