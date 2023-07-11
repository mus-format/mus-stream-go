package mock

import (
	"reflect"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func NewUnmarshaller[T any]() Unmarshaller[T] {
	return Unmarshaller[T]{mok.New("Unmarshaller")}
}

type Unmarshaller[T any] struct {
	*mok.Mock
}

func (u Unmarshaller[T]) RegisterUnmarshalMUS(
	fn func(r muss.Reader) (t T, n int, err error)) Unmarshaller[T] {
	u.Register("UnmarshalMUS", fn)
	return u
}

func (u Unmarshaller[T]) RegisterNUnmarshalMUS(n int,
	fn func(r muss.Reader) (t T, n int, err error)) Unmarshaller[T] {
	u.RegisterN("UnmarshalMUS", n, fn)
	return u
}

func (u Unmarshaller[T]) UnmarshalMUS(r muss.Reader) (t T, n int, err error) {
	var rVal reflect.Value
	if r == nil {
		rVal = reflect.Zero(reflect.TypeOf((*muss.Writer)(nil)).Elem())
	} else {
		rVal = reflect.ValueOf(r)
	}
	result, err := u.Call("UnmarshalMUS", rVal)
	if err != nil {
		panic(err)
	}
	t, _ = result[0].(T)
	n = result[1].(int)
	err, _ = result[2].(error)
	return
}
