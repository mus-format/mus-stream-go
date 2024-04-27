package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type SkipMUSFn func(r muss.Reader) (n int, err error)

func NewSkipper() Skipper {
	return Skipper{mok.New("Skipper")}
}

type Skipper struct {
	*mok.Mock
}

func (u Skipper) RegisterSkipMUS(fn SkipMUSFn) Skipper {
	u.Register("SkipMUS", fn)
	return u
}

func (u Skipper) RegisterNSkipMUS(n int, fn SkipMUSFn) Skipper {
	u.RegisterN("SkipMUS", n, fn)
	return u
}

func (u Skipper) SkipMUS(r muss.Reader) (n int, err error) {
	result, err := u.Call("SkipMUS", mok.SafeVal[muss.Reader](r))
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
