// Package dts provides DTM (Data Type Metadata) support for mus-stream-go
// serializer. It wraps a type serializer together with a DTM value,
// enabling typed data serialization.
package dts

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
)

// New creates a new DTS.
func New[T any](dtm com.DTM, ser mus.Serializer[T],
) DTS[T] {
	return DTS[T]{dtm, ser}
}

// DTS implements the mus.Serializer interface and provides DTM support for the
// mus-stream-go serializer. It helps to serializer DTM + data.
type DTS[T any] struct {
	dtm com.DTM
	ser mus.Serializer[T]
}

// DTM returns the initialization value.
func (d DTS[T]) DTM() com.DTM {
	return d.dtm
}

// Marshal marshals DTM + data.
func (d DTS[T]) Marshal(t T, w mus.Writer) (n int, err error) {
	n, err = DTMSer.Marshal(d.dtm, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = d.ser.Marshal(t, w)
	n += n1
	return
}

// Unmarshal unmarshals DTM + data.
//
// Returns ErrWrongDTM if the unmarshalled DTM differs from the dts.DTM().
func (d DTS[T]) Unmarshal(r mus.Reader) (t T, n int, err error) {
	dtm, n, err := DTMSer.Unmarshal(r)
	if err != nil {
		return
	}
	if dtm != d.dtm {
		err = com.NewWrongDTMError(d.dtm, dtm)
		return
	}
	var n1 int
	t, n1, err = d.UnmarshalData(r)
	n += n1
	return
}

// Size calculates the size of the DTM + data.
func (d DTS[T]) Size(t T) (size int) {
	size = DTMSer.Size(d.dtm)
	return size + d.ser.Size(t)
}

// Skip skips DTM + data.
//
// Returns ErrWrongDTM if the unmarshalled DTM differs from the dts.DTM().
func (d DTS[T]) Skip(r mus.Reader) (n int, err error) {
	dtm, n, err := DTMSer.Unmarshal(r)
	if err != nil {
		return
	}
	if dtm != d.dtm {
		err = com.NewWrongDTMError(d.dtm, dtm)
		return
	}
	n1, err := d.SkipData(r)
	n += n1
	return
}

// UnmarshalData unmarshals only data.
func (d DTS[T]) UnmarshalData(r mus.Reader) (t T, n int, err error) {
	return d.ser.Unmarshal(r)
}

// SkipData skips only data.
func (d DTS[T]) SkipData(r mus.Reader) (n int, err error) {
	return d.ser.Skip(r)
}
