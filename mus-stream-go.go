package muss

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

// Marshaller is the interface that wraps the MarshalMUS method.
//
// MarshalMUS marshals data to the MUS format. Returns the number of used
// bytes and an error.
type Marshaller[T any] interface {
	MarshalMUS(t T, w Writer) (n int, err error)
}

// MarshallerFn is a functional implementation of the Marshaller interface.
type MarshallerFn[T any] func(t T, w Writer) (n int, err error)

func (fn MarshallerFn[T]) MarshalMUS(t T, w Writer) (n int, err error) {
	return fn(t, w)
}

// Unmarshaller is the interface that wraps the UnmarshalMUS method.
//
// UnmarshalMUS unmarshals data from the MUS format. Returns data, the number of
// used bytes and an error.
type Unmarshaller[T any] interface {
	UnmarshalMUS(r Reader) (t T, n int, err error)
}

// UnmarshallerFn is a functional implementation of the Unmarshaller interface.
type UnmarshallerFn[T any] func(r Reader) (t T, n int, err error)

func (fn UnmarshallerFn[T]) UnmarshalMUS(r Reader) (t T, n int, err error) {
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
