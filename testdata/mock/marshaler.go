package mock

import (
	"reflect"

	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func NewMarshaler[T any]() Marshaler[T] {
	return Marshaler[T]{mok.New("Marshaler")}
}

type Marshaler[T any] struct {
	*mok.Mock
}

func (m Marshaler[T]) RegisterMarshalMUS(
	fn func(t T, w mustrm.Writer) (n int, err error)) Marshaler[T] {
	m.Register("MarshalMUS", fn)
	return m
}

func (m Marshaler[T]) RegisterNMarshalMUS(n int,
	fn func(t T, w mustrm.Writer) (n int, err error)) Marshaler[T] {
	m.RegisterN("MarshalMUS", n, fn)
	return m
}

func (m Marshaler[T]) MarshalMUS(t T, w mustrm.Writer) (n int, err error) {
	var tVal reflect.Value
	if v := reflect.ValueOf(t); (v.Kind() == reflect.Ptr) && v.IsNil() {
		tVal = reflect.Zero(reflect.TypeOf((*T)(nil)).Elem())
	} else {
		tVal = reflect.ValueOf(t)
	}
	var wVal reflect.Value
	if w == nil {
		wVal = reflect.Zero(reflect.TypeOf((*mustrm.Writer)(nil)).Elem())
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
