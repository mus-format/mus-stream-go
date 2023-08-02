package mock

import (
	"reflect"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func NewMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{mok.New("Marshaller")}
}

type Marshaller[T any] struct {
	*mok.Mock
}

func (m Marshaller[T]) RegisterMarshalMUS(
	fn func(t T, w muss.Writer) (n int, err error)) Marshaller[T] {
	m.Register("MarshalMUS", fn)
	return m
}

func (m Marshaller[T]) RegisterNMarshalMUS(n int,
	fn func(t T, w muss.Writer) (n int, err error)) Marshaller[T] {
	m.RegisterN("MarshalMUS", n, fn)
	return m
}

func (m Marshaller[T]) MarshalMUS(t T, w muss.Writer) (n int, err error) {
	var tVal reflect.Value
	if v := reflect.ValueOf(t); (v.Kind() == reflect.Ptr) && v.IsNil() {
		tVal = reflect.Zero(reflect.TypeOf((*T)(nil)).Elem())
	} else {
		tVal = reflect.ValueOf(t)
	}
	var wVal reflect.Value
	if w == nil {
		wVal = reflect.Zero(reflect.TypeOf((*muss.Writer)(nil)).Elem())
	} else {
		wVal = reflect.ValueOf(w)
	}
	result, err := m.Call("MarshalMUS", tVal, wVal)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
