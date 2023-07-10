package mock

import (
	"reflect"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func NewUnMarshaller[T any]() UnMarshaller[T] {
	return UnMarshaller[T]{mok.New("UnMarshaller")}
}

type UnMarshaller[T any] struct {
	*mok.Mock
}

func (u UnMarshaller[T]) RegisterUnmarshalMUS(
	fn func(r muss.Reader) (t T, n int, err error)) UnMarshaller[T] {
	u.Register("UnmarshalMUS", fn)
	return u
}

func (u UnMarshaller[T]) RegisterNUnmarshalMUS(n int,
	fn func(r muss.Reader) (t T, n int, err error)) UnMarshaller[T] {
	u.RegisterN("UnmarshalMUS", n, fn)
	return u
}

func (u UnMarshaller[T]) UnmarshalMUS(r muss.Reader) (t T, n int, err error) {
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
