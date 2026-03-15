package typed

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
)

// TypedSer implements the mus.Serializer interface and provides DTM support for the
// mus-stream-go serializer. It helps to serializer DTM + data.
type TypedSer[T any] struct {
	dtm com.DTM
	ser mus.Serializer[T]
}

// NewTypedSer creates a new TypedSer.
func NewTypedSer[T any](dtm com.DTM, ser mus.Serializer[T]) TypedSer[T] {
	return TypedSer[T]{dtm, ser}
}

// DTM returns the initialization value.
func (t TypedSer[T]) DTM() com.DTM {
	return t.dtm
}

// Marshal marshals DTM + data.
func (t TypedSer[T]) Marshal(v T, w mus.Writer) (n int, err error) {
	n, err = DTMSer.Marshal(t.dtm, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = t.ser.Marshal(v, w)
	n += n1
	return
}

// Unmarshal unmarshals DTM + data.
//
// Returns ErrWrongDTM if the unmarshalled DTM differs from the typed.DTM().
func (t TypedSer[T]) Unmarshal(r mus.Reader) (v T, n int, err error) {
	dtm, n, err := DTMSer.Unmarshal(r)
	if err != nil {
		return
	}
	if dtm != t.dtm {
		err = com.NewWrongDTMError(t.dtm, dtm)
		return
	}
	var n1 int
	v, n1, err = t.UnmarshalData(r)
	n += n1
	return
}

// Size calculates the size of the DTM + data.
func (t TypedSer[T]) Size(v T) (size int) {
	size = DTMSer.Size(t.dtm)
	return size + t.ser.Size(v)
}

// Skip skips DTM + data.
//
// Returns ErrWrongDTM if the unmarshalled DTM differs from the typed.DTM().
func (t TypedSer[T]) Skip(r mus.Reader) (n int, err error) {
	dtm, n, err := DTMSer.Unmarshal(r)
	if err != nil {
		return
	}
	if dtm != t.dtm {
		err = com.NewWrongDTMError(t.dtm, dtm)
		return
	}
	n1, err := t.SkipData(r)
	n += n1
	return
}

// UnmarshalData unmarshals only data.
func (t TypedSer[T]) UnmarshalData(r mus.Reader) (v T, n int, err error) {
	return t.ser.Unmarshal(r)
}

// SkipData skips only data.
func (t TypedSer[T]) SkipData(r mus.Reader) (n int, err error) {
	return t.ser.Skip(r)
}
