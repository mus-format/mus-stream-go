package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type UnmarshalMUSFn[T any] func(r muss.Reader) (t T, n int, err error)

func NewUnmarshaller[T any]() Unmarshaller[T] {
	return Unmarshaller[T]{mok.New("Unmarshaller")}
}

type Unmarshaller[T any] struct {
	*mok.Mock
}

func (u Unmarshaller[T]) RegisterUnmarshalMUS(
	fn UnmarshalMUSFn[T]) Unmarshaller[T] {
	u.Register("UnmarshalMUS", fn)
	return u
}

func (u Unmarshaller[T]) RegisterNUnmarshalMUS(n int,
	fn UnmarshalMUSFn[T]) Unmarshaller[T] {
	u.RegisterN("UnmarshalMUS", n, fn)
	return u
}

func (u Unmarshaller[T]) UnmarshalMUS(r muss.Reader) (t T, n int, err error) {
	result, err := u.Call("UnmarshalMUS", mok.SafeVal[muss.Reader](r))
	if err != nil {
		panic(err)
	}
	t, _ = result[0].(T)
	n = result[1].(int)
	err, _ = result[2].(error)
	return
}
