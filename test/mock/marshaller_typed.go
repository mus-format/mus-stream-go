package mock

import (
	"github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type MarshalTypedMUSFn func(w mus.Writer) (n int, err error)
type SizeTypedMUSFn func() (size int)

type MarshallerTyped struct {
	*mok.Mock
}

func NewMarshallerTyped() MarshallerTyped {
	return MarshallerTyped{mok.New("MarshallerTyped")}
}

func (m MarshallerTyped) RegisterMarshalTypedMUS(fn MarshalTypedMUSFn) MarshallerTyped {
	m.Register("MarshalTypedMUS", fn)
	return m
}

func (m MarshallerTyped) RegisterMarshalTypedMUSN(n int, fn MarshalTypedMUSFn) MarshallerTyped {
	m.RegisterN("MarshalTypedMUS", n, fn)
	return m
}

func (m MarshallerTyped) RegisterSizeTypedMUS(fn SizeTypedMUSFn) MarshallerTyped {
	m.Register("SizeTypedMUS", fn)
	return m
}

func (m MarshallerTyped) RegisterSizeTypedMUSN(n int, fn SizeTypedMUSFn) MarshallerTyped {
	m.RegisterN("SizeTypedMUS", n, fn)
	return m
}

func (m MarshallerTyped) MarshalTypedMUS(w mus.Writer) (n int, err error) {
	result, err := m.Call("MarshalTypedMUS", w)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m MarshallerTyped) SizeTypedMUS() (size int) {
	result, err := m.Call("SizeTypedMUS")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
