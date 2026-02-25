package mock

import (
	"github.com/ymz-ncnk/mok"
)

func NewWriter() Writer {
	return Writer{mok.New("Writer")}
}

type Writer struct {
	*mok.Mock
}

func (m Writer) RegisterWrite(fn func(p []byte) (n int, err error)) Writer {
	m.Register("Write", fn)
	return m
}

func (m Writer) RegisterWriteByte(fn func(c byte) error) Writer {
	m.Register("WriteByte", fn)
	return m
}

func (m Writer) RegisterNWriteByte(n int, fn func(c byte) error) Writer {
	m.RegisterN("WriteByte", n, fn)
	return m
}

func (m Writer) RegisterWriteString(fn func(s string) (n int, err error)) Writer {
	m.Register("WriteString", fn)
	return m
}

func (m Writer) Write(p []byte) (n int, err error) {
	result, err := m.Call("Write", p)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m Writer) WriteByte(c byte) (err error) {
	result, err := m.Call("WriteByte", c)
	if err != nil {
		panic(err)
	}
	err, _ = result[0].(error)
	return
}

func (m Writer) WriteString(s string) (n int, err error) {
	result, err := m.Call("WriteString", s)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
