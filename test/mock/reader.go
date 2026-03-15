package mock

import "github.com/ymz-ncnk/mok"

func NewReader() Reader {
	return Reader{mok.New("Reader")}
}

type Reader struct {
	*mok.Mock
}

func (m Reader) RegisterRead(fn func(p []byte) (n int, err error)) Reader {
	m.Register("Read", fn)
	return m
}

func (m Reader) RegisterNReadByte(n int,
	fn func() (b byte, err error),
) Reader {
	m.RegisterN("ReadByte", n, fn)
	return m
}

func (m Reader) RegisterReadByte(fn func() (b byte, err error)) Reader {
	m.Register("ReadByte", fn)
	return m
}

func (m Reader) Read(p []byte) (n int, err error) {
	result, err := m.Call("Read", p)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m Reader) ReadByte() (b byte, err error) {
	result, err := m.Call("ReadByte")
	if err != nil {
		panic(err)
	}
	b = result[0].(byte)
	err, _ = result[1].(error)
	return
}
