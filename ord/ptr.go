package ord

import (
	com "github.com/mus-format/common-go"
	mus "github.com/mus-format/mus-stream-go"
)

// MarshalPtr writes an encoded pointer.
//
// The m argument specifies the Marshaller for the pointer base type.
//
// In addition to the number of used bytes, it may also return a Writer or
// Marshaller error.
func MarshalPtr[T any](v *T, m mus.Marshaller[T], w mus.Writer) (n int,
	err error) {
	if v == nil {
		err = w.WriteByte(byte(com.Nil))
		if err != nil {
			return
		}
		n++
		return
	}
	if err = w.WriteByte(byte(com.NotNil)); err != nil {
		return
	}
	n, err = m.Marshal(*v, w)
	n += 1
	return
}

// UnmarshalPtr reads an encoded pointer.
//
// The u argument specifies the Unmarshaller for the pointer base type.
//
// In addition to the pointer and the number of used bytes, it may also return
// com.ErrWrongFormat, Unarshaller or a Reader error.
func UnmarshalPtr[T any](u mus.Unmarshaller[T], r mus.Reader) (v *T, n int,
	err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == byte(com.Nil) {
		return
	}
	if b != byte(com.NotNil) {
		err = com.ErrWrongFormat
		return
	}
	var n1 int
	k, n1, err := u.Unmarshal(r)
	n += n1
	if err != nil {
		return
	}
	v = &k
	return
}

// SizePtr returns the size of an encoded pointer.
//
// The s argument specifies the Sizer for the pointer base type.
func SizePtr[T any](v *T, s mus.Sizer[T]) (size int) {
	if v != nil {
		return 1 + s.Size(*v)
	}
	return 1
}

// SkipPtr skips an encoded pointer.
//
// The sk argument specifies the Skipper for the pointer base type.
//
// In addition to the number of used bytes, it may also return com.ErrWrongFormat,
// a Skipper or Reader error.
func SkipPtr(sk mus.Skipper, r mus.Reader) (n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	n++
	if b == byte(com.Nil) {
		return
	}
	if b != byte(com.NotNil) {
		err = com.ErrWrongFormat
		return
	}
	n, err = sk.Skip(r)
	n += 1
	return
}
