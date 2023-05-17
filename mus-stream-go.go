package mustrm

import "io"

// Writer is the interface that groups the WriteByte, Write and WriteString
// methods.
type Writer interface {
	io.ByteWriter
	io.Writer
	io.StringWriter
}

// Reader is the interface that groups the basic ReadByte and Read methods.
type Reader interface {
	io.ByteReader
	io.Reader
}

// Marshaler is the interface that wraps the MarshalMUS method.
//
// MarshalMUS marshals data to the MUS format. Returns the number of used
// bytes and an error.
type Marshaler[T any] interface {
	MarshalMUS(t T, w Writer) (n int, err error)
}

// MarshalerFn is a functional implementation of the Marshaler interface.
type MarshalerFn[T any] func(t T, w Writer) (n int, err error)

func (fn MarshalerFn[T]) MarshalMUS(t T, w Writer) (n int, err error) {
	return fn(t, w)
}

// Unmarshaler is the interface that wraps the UnmarshalMUS method.
//
// UnmarshalMUS unmarshals data from the MUS format. Returns data, the number of
// used bytes and an error.
type Unmarshaler[T any] interface {
	UnmarshalMUS(r Reader) (t T, n int, err error)
}

// UnmarshalerFn is a functional implementation of the Unmarshaler interface.
type UnmarshalerFn[T any] func(r Reader) (t T, n int, err error)

func (fn UnmarshalerFn[T]) UnmarshalMUS(r Reader) (t T, n int, err error) {
	return fn(r)
}

// Sizer is the interface that wraps the SizeMUS method.
//
// SizeMUS calculates the size of data in the MUS format.
type Sizer[T any] interface {
	SizeMUS(t T) (size int)
}

// SizerFn is a functional implementation of the Sizer interface.
type SizerFn[T any] func(t T) (size int)

func (fn SizerFn[T]) SizeMUS(t T) (size int) {
	return fn(t)
}

// Skipper is the interface that wraps the SkipMUS method.
//
// SkipMUS skips data in the MUS format. Returns the number of skipped bytes and
// an error.
type Skipper interface {
	SkipMUS(r Reader) (n int, err error)
}

// SkipperFn is a functional implementation of the Skipper interface.
type SkipperFn func(r Reader) (n int, err error)

func (fn SkipperFn) SkipMUS(r Reader) (n int, err error) {
	return fn(r)
}
