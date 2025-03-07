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

// Serializer is the interface that groups the Marshal, Unmarshal, Size and
// Skip methods.
//
// Marshal writes an encoded value, returning the number of bytes written and
// any error encountered.
//
// Unmarshal reads an encoded value, returning the value, the number of bytes
// read and any error encountered.
//
// Size returns the number of bytes needed to encode the value.
//
// Skip skips an encoded value, returning the number of skipped bytes and
// any error encountered.
type Serializer[T any] interface {
	Marshal(t T, w Writer) (n int, err error)
	Unmarshal(r Reader) (t T, n int, err error)
	Size(t T) (size int)
	Skip(r Reader) (n int, err error)
}
