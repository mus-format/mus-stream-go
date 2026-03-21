package test

import (
	"bytes"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/mus-format/mus-stream-go"
	asserterror "github.com/ymz-ncnk/assert/error"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
	"github.com/ymz-ncnk/mok"
)

func Test[T any](cases []T, ser mus.Serializer[T], t *testing.T) {
	var err error
	for i := range cases {
		var (
			size = ser.Size(cases[i])
			buf  = bytes.NewBuffer(make([]byte, 0, size))
			n    int
			v    T
		)
		n, err = ser.Marshal(cases[i], buf)
		assertfatal.EqualError(t, err, nil)
		asserterror.Equal(t, n, size,
			fmt.Sprintf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n),
		)
		v, n, err := ser.Unmarshal(buf)
		assertfatal.EqualError(t, err, nil)
		asserterror.Equal(t, n, size,
			fmt.Sprintf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n),
		)
		if tm, ok := any(v).(time.Time); ok {
			tm1 := any(cases[i]).(time.Time)
			if !tm.Equal(tm1) {
				t.Errorf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v)
			}
			continue
		}
		if f64, ok := any(v).(float64); ok {
			if math.Float64bits(f64) == math.Float64bits(any(cases[i]).(float64)) {
				continue
			}
		}
		if f32, ok := any(v).(float32); ok {
			if math.Float32bits(f32) == math.Float32bits(any(cases[i]).(float32)) {
				continue
			}
		}
		asserterror.EqualDeep(t, v, cases[i],
			fmt.Sprintf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v),
		)
	}
}

func TestSkip[T any](cases []T, ser mus.Serializer[T], t *testing.T) {
	for i := range cases {
		var (
			size = ser.Size(cases[i])
			buf  = bytes.NewBuffer(make([]byte, 0, size))
		)
		ser.Marshal(cases[i], buf)
		n, err := ser.Skip(buf)
		assertfatal.EqualError(t, err, nil)
		asserterror.Equal(t, n, buf.Cap(), "skipped not enough")
	}
}

func TestValidation[T any](testCase T, ser mus.Serializer[T], wantErr error,
	t *testing.T,
) {
	var (
		size = ser.Size(testCase)
		buf  = bytes.NewBuffer(make([]byte, 0, size))
	)
	ser.Marshal(testCase, buf)
	_, _, err := ser.Unmarshal(buf)
	asserterror.EqualError(t, err, wantErr, "unexpected error")
}

func TestMarshalOnly[T any](v T, w mus.Writer, ser interface {
	Marshal(T, mus.Writer) (int, error)
}, want MarshalResults, t *testing.T,
) {
	n, err := ser.Marshal(v, w)
	assertfatal.EqualError(t, err, want.Err)
	asserterror.Equal(t, n, want.N)
	asserterror.EqualDeep(t, mok.CheckCalls(want.Mocks), mok.EmptyInfomap)
}

func TestUnmarshalOnly[T any](r mus.Reader, ser interface {
	Unmarshal(mus.Reader) (T, int, error)
}, want UnmarshalResults[T], t *testing.T,
) {
	v, n, err := ser.Unmarshal(r)
	assertfatal.EqualError(t, err, want.Err)
	asserterror.Equal(t, n, want.N)
	if tm, ok := any(v).(time.Time); ok {
		tm1 := any(want.V).(time.Time)
		if !tm.Equal(tm1) {
			t.Errorf("unexpected v, want '%v' actual '%v'", want.V, v)
		}
	} else {
		if f64, ok := any(v).(float64); ok {
			if math.Float64bits(f64) == math.Float64bits(any(want.V).(float64)) {
				goto Mocks
			}
		}
		if f32, ok := any(v).(float32); ok {
			if math.Float32bits(f32) == math.Float32bits(any(want.V).(float32)) {
				goto Mocks
			}
		}
		asserterror.EqualDeep(t, v, want.V)
	}
Mocks:
	asserterror.EqualDeep(t, mok.CheckCalls(want.Mocks), mok.EmptyInfomap)
}

func TestSkipOnly(r mus.Reader, ser interface {
	Skip(mus.Reader) (int, error)
}, want SkipResults, t *testing.T,
) {
	n, err := ser.Skip(r)
	assertfatal.EqualError(t, err, want.Err)
	asserterror.Equal(t, n, want.N)
	asserterror.EqualDeep(t, mok.CheckCalls(want.Mocks), mok.EmptyInfomap)
}

// -----------------------------------------------------------------------------

type MarshallerFn[T any] func(v T, w mus.Writer) (int, error)

func (f MarshallerFn[T]) Marshal(v T, w mus.Writer) (int, error) {
	return f(v, w)
}

type UnmarshallerFn[T any] func(r mus.Reader) (T, int, error)

func (f UnmarshallerFn[T]) Unmarshal(r mus.Reader) (T, int, error) {
	return f(r)
}

type SkipperFn func(r mus.Reader) (int, error)

func (f SkipperFn) Skip(r mus.Reader) (int, error) {
	return f(r)
}

type MarshallerResults = MarshalResults

type MarshalResults struct {
	N     int
	Err   error
	Mocks []*mok.Mock
}

type UnmarshalResults[T any] struct {
	V     T
	N     int
	Err   error
	Mocks []*mok.Mock
}

type SkipResults struct {
	N     int
	Err   error
	Mocks []*mok.Mock
}
