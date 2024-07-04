package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type SkipFn func(r muss.Reader) (n int, err error)

func NewSkipper() Skipper {
	return Skipper{mok.New("Skipper")}
}

type Skipper struct {
	*mok.Mock
}

func (u Skipper) RegisterSkip(fn SkipFn) Skipper {
	u.Register("Skip", fn)
	return u
}

func (u Skipper) RegisterNSkip(n int, fn SkipFn) Skipper {
	u.RegisterN("Skip", n, fn)
	return u
}

func (u Skipper) Skip(r muss.Reader) (n int, err error) {
	result, err := u.Call("Skip", mok.SafeVal[muss.Reader](r))
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
