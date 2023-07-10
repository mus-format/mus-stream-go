package ord

import (
	com "github.com/mus-format/common-go"
	mus "github.com/mus-format/mus-stream-go"
)

// MarshalPtr writes the MUS encoding of a pointer. Returns the number of used
// bytes and error.
//
// The m argument specifies the Marshaler for the pointer base type.
func MarshalPtr[T any](v *T, m mus.Marshaler[T], w mus.Writer) (n int,
	err error) {
	if v == nil {
		err = w.WriteByte(com.NilFlag)
		if err != nil {
			return
		}
		n++
		return
	}
	if err = w.WriteByte(com.NotNilFlag); err != nil {
		return
	}
	n, err = m.MarshalMUS(*v, w)
	n += 1
	return
}

// UnmarshalPtr reads a MUS-encoded pointer. In addition to the pointer, it
// returns the number of used bytes and an error.
//
// The u argument specifies the Unmarshaler for the base pointer type.
//
// The error returned by UnmarshalPtr can be one of com.ErrWrongFormat, an
// Unarshaler or Reader error.
func UnmarshalPtr[T any](u mus.Unmarshaler[T], r mus.Reader) (v *T, n int,
	err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == com.NilFlag {
		return
	}
	if b != com.NotNilFlag {
		err = com.ErrWrongFormat
		return
	}
	var n1 int
	k, n1, err := u.UnmarshalMUS(r)
	n += n1
	if err != nil {
		return
	}
	v = &k
	return
}

// SizePtr returns the size of a MUS-encoded pointer.
//
// The s argument specifies the Sizer for the pointer base type.
func SizePtr[T any](v *T, s mus.Sizer[T]) (size int) {
	if v != nil {
		return 1 + s.SizeMUS(*v)
	}
	return 1
}

// SkipPtr skips a MUS-encoded pointer. Returns the number of skiped bytes and
// an error.
//
// The sk argument specifies the Skipper for the pointer base type.
//
// The error returned by SkipPtr can be one of com.ErrWrongFormat, a
// Skipper or Reader error.
func SkipPtr(sk mus.Skipper, r mus.Reader) (n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == com.NilFlag {
		return
	}
	if b != com.NotNilFlag {
		err = com.ErrWrongFormat
		return
	}
	n, err = sk.SkipMUS(r)
	n += 1
	return
}
