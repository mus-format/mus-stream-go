package mock

import (
	"reflect"

	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func NewUnmarshaler[T any]() Unmarshaler[T] {
	return Unmarshaler[T]{mok.New("Unmarshaler")}
}

type Unmarshaler[T any] struct {
	*mok.Mock
}

func (u Unmarshaler[T]) RegisterUnmarshalMUS(
	fn func(r mustrm.Reader) (t T, n int, err error)) Unmarshaler[T] {
	u.Register("UnmarshalMUS", fn)
	return u
}

func (u Unmarshaler[T]) RegisterNUnmarshalMUS(n int,
	fn func(r mustrm.Reader) (t T, n int, err error)) Unmarshaler[T] {
	u.RegisterN("UnmarshalMUS", n, fn)
	return u
}

func (u Unmarshaler[T]) UnmarshalMUS(r mustrm.Reader) (t T, n int, err error) {
	var rVal reflect.Value
	if r == nil {
		rVal = reflect.Zero(reflect.TypeOf((*mustrm.Writer)(nil)).Elem())
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
