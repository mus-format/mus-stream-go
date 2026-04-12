package typed

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
)

// Ser implements the mus.Serializer interface and provides DTM support for the
// mus-stream-go serializer. It helps to serializer DTM + data.
type Ser[T any] struct {
	dtm com.DTM
	ser mus.Serializer[T]
}

// NewSer creates a new TypedSer.
func NewSer[T any](dtm com.DTM, ser mus.Serializer[T]) Ser[T] {
	return Ser[T]{dtm, ser}
}

// DTM returns the initialization value.
func (s Ser[T]) DTM() com.DTM {
	return s.dtm
}

// Marshal marshals DTM + data.
func (s Ser[T]) Marshal(v T, w mus.Writer) (n int, err error) {
	n, err = DTMSer.Marshal(s.dtm, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.ser.Marshal(v, w)
	n += n1
	return
}

// Unmarshal unmarshals DTM + data.
//
// Returns com.WrongDTMError if the unmarshalled DTM differs from the expected
// one.
func (s Ser[T]) Unmarshal(r mus.Reader) (v T, n int, err error) {
	dtm, n, err := DTMSer.Unmarshal(r)
	if err != nil {
		return
	}
	if dtm != s.dtm {
		err = com.NewWrongDTMError(s.dtm, dtm)
		return
	}
	var n1 int
	v, n1, err = s.UnmarshalData(r)
	n += n1
	return
}

// Size calculates the size of the DTM + data.
func (s Ser[T]) Size(v T) (size int) {
	size = DTMSer.Size(s.dtm)
	return size + s.ser.Size(v)
}

// Skip skips DTM + data.
//
// Returns com.WrongDTMError if the unmarshalled DTM differs from the expected
// one.
func (s Ser[T]) Skip(r mus.Reader) (n int, err error) {
	dtm, n, err := DTMSer.Unmarshal(r)
	if err != nil {
		return
	}
	if dtm != s.dtm {
		err = com.NewWrongDTMError(s.dtm, dtm)
		return
	}
	n1, err := s.SkipData(r)
	n += n1
	return
}

// UnmarshalData unmarshals only data.
func (s Ser[T]) UnmarshalData(r mus.Reader) (v T, n int, err error) {
	return s.ser.Unmarshal(r)
}

// SkipData skips only data.
func (s Ser[T]) SkipData(r mus.Reader) (n int, err error) {
	return s.ser.Skip(r)
}
