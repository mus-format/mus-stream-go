package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type MarshalFn[T any] func(t T, w muss.Writer) (n int, err error)

func NewMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{mok.New("Marshaller")}
}

type Marshaller[T any] struct {
	*mok.Mock
}

func (m Marshaller[T]) RegisterMarshal(fn MarshalFn[T]) Marshaller[T] {
	m.Register("Marshal", fn)
	return m
}

func (m Marshaller[T]) RegisterNMarshal(n int, fn MarshalFn[T]) Marshaller[T] {
	m.RegisterN("Marshal", n, fn)
	return m
}

func (m Marshaller[T]) Marshal(t T, w muss.Writer) (n int, err error) {
	result, err := m.Call("Marshal", mok.SafeVal[T](t),
		mok.SafeVal[muss.Writer](w))
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
