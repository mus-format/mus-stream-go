package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type MarshalMUSFn[T any] func(t T, w muss.Writer) (n int, err error)

func NewMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{mok.New("Marshaller")}
}

type Marshaller[T any] struct {
	*mok.Mock
}

func (m Marshaller[T]) RegisterMarshalMUS(fn MarshalMUSFn[T]) Marshaller[T] {
	m.Register("MarshalMUS", fn)
	return m
}

func (m Marshaller[T]) RegisterNMarshalMUS(n int, fn MarshalMUSFn[T]) Marshaller[T] {
	m.RegisterN("MarshalMUS", n, fn)
	return m
}

func (m Marshaller[T]) MarshalMUS(t T, w muss.Writer) (n int, err error) {
	result, err := m.Call("MarshalMUS", mok.SafeVal[T](t),
		mok.SafeVal[muss.Writer](w))
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
