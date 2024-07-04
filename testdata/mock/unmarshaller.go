package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type UnmarshalFn[T any] func(r muss.Reader) (t T, n int, err error)

func NewUnmarshaller[T any]() Unmarshaller[T] {
	return Unmarshaller[T]{mok.New("Unmarshaller")}
}

type Unmarshaller[T any] struct {
	*mok.Mock
}

func (u Unmarshaller[T]) RegisterUnmarshal(
	fn UnmarshalFn[T]) Unmarshaller[T] {
	u.Register("Unmarshal", fn)
	return u
}

func (u Unmarshaller[T]) RegisterNUnmarshal(n int,
	fn UnmarshalFn[T]) Unmarshaller[T] {
	u.RegisterN("Unmarshal", n, fn)
	return u
}

func (u Unmarshaller[T]) Unmarshal(r muss.Reader) (t T, n int, err error) {
	result, err := u.Call("Unmarshal", mok.SafeVal[muss.Reader](r))
	if err != nil {
		panic(err)
	}
	t, _ = result[0].(T)
	n = result[1].(int)
	err, _ = result[2].(error)
	return
}
