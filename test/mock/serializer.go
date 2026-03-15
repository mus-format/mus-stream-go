package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type MarshalFn[T any] func(t T, w muss.Writer) (n int, err error)

type UnmarshalFn[T any] func(r muss.Reader) (t T, n int, err error)

type SizeFn[T any] func(t T) (size int)

type SkipFn func(r muss.Reader) (n int, err error)

func NewSerializer[T any]() Serializer[T] {
	return Serializer[T]{mok.New("Serializer")}
}

type Serializer[T any] struct {
	*mok.Mock
}

func (s Serializer[T]) RegisterMarshal(fn MarshalFn[T]) Serializer[T] {
	s.Register("Marshal", fn)
	return s
}

func (m Serializer[T]) RegisterMarshalN(n int, fn MarshalFn[T]) Serializer[T] {
	m.RegisterN("Marshal", n, fn)
	return m
}

func (s Serializer[T]) RegisterUnmarshal(fn UnmarshalFn[T]) Serializer[T] {
	s.Register("Unmarshal", fn)
	return s
}

func (u Serializer[T]) RegisterUnmarshalN(n int, fn UnmarshalFn[T]) Serializer[T] {
	u.RegisterN("Unmarshal", n, fn)
	return u
}

func (s Serializer[T]) RegisterSize(fn SizeFn[T]) Serializer[T] {
	s.Register("Size", fn)
	return s
}

func (m Serializer[T]) RegisterSizeN(n int, fn SizeFn[T]) Serializer[T] {
	m.RegisterN("Size", n, fn)
	return m
}

func (s Serializer[T]) RegisterSkip(fn SkipFn) Serializer[T] {
	s.Register("Skip", fn)
	return s
}

func (s Serializer[T]) RegisterSkipN(n int, fn SkipFn) Serializer[T] {
	s.RegisterN("Skip", n, fn)
	return s
}

func (m Serializer[T]) Marshal(t T, w muss.Writer) (n int, err error) {
	result, err := m.Call("Marshal", mok.SafeVal[T](t),
		mok.SafeVal[muss.Writer](w))
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (s Serializer[T]) Unmarshal(r muss.Reader) (t T, n int, err error) {
	result, err := s.Call("Unmarshal", mok.SafeVal[muss.Reader](r))
	if err != nil {
		panic(err)
	}
	t, _ = result[0].(T)
	n = result[1].(int)
	err, _ = result[2].(error)
	return
}

func (s Serializer[T]) Size(t T) (size int) {
	result, err := s.Call("Size", mok.SafeVal[T](t))
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}

func (s Serializer[T]) Skip(r muss.Reader) (n int, err error) {
	result, err := s.Call("Skip", mok.SafeVal[muss.Reader](r))
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
