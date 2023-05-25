package mock

import (
	"reflect"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func NewSkipper() Skipper {
	return Skipper{mok.New("Skipper")}
}

type Skipper struct {
	*mok.Mock
}

func (u Skipper) RegisterSkipMUS(
	fn func(r muss.Reader) (n int, err error)) Skipper {
	u.Register("SkipMUS", fn)
	return u
}

func (u Skipper) RegisterNSkipMUS(n int,
	fn func(r muss.Reader) (n int, err error)) Skipper {
	u.RegisterN("SkipMUS", n, fn)
	return u
}

func (u Skipper) SkipMUS(r muss.Reader) (n int, err error) {
	var rVal reflect.Value
	if r == nil {
		rVal = reflect.Zero(reflect.TypeOf((*muss.Writer)(nil)).Elem())
	} else {
		rVal = reflect.ValueOf(r)
	}
	result, err := u.Call("SkipMUS", rVal)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
