package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type MarshalMUSFn func(w muss.Writer) (n int, err error)
type SizeMUSFn func() (size int)

type Marshaller struct {
	*mok.Mock
}

func NewMarshaller() Marshaller {
	return Marshaller{mok.New("Marshaller")}
}

func (m Marshaller) RegisterMarshalMUS(fn MarshalMUSFn) Marshaller {
	m.Register("MarshalMUS", fn)
	return m
}

func (m Marshaller) RegisterMarshalMUSN(n int, fn MarshalMUSFn) Marshaller {
	m.RegisterN("MarshalMUS", n, fn)
	return m
}

func (m Marshaller) RegisterSizeMUS(fn SizeMUSFn) Marshaller {
	m.Register("SizeMUS", fn)
	return m
}

func (m Marshaller) RegisterSizeMUSN(n int, fn SizeMUSFn) Marshaller {
	m.RegisterN("SizeMUS", n, fn)
	return m
}

func (m Marshaller) MarshalMUS(w muss.Writer) (n int, err error) {
	result, err := m.Call("MarshalMUS", w)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m Marshaller) SizeMUS() (size int) {
	result, err := m.Call("SizeMUS")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
