package dts

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/test/mock"
	"github.com/mus-format/mus-stream-go/varint"
)

const FooDTM com.DTM = 1

type Foo struct {
	Num int
	Str string
}

var FooSer = fooSer{}

type fooSer struct{}

func (s fooSer) Marshal(foo Foo, w mus.Writer) (n int, err error) {
	n, err = varint.Int.Marshal(foo.Num, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Marshal(foo.Str, w)
	n += n1
	return
}

func (s fooSer) Unmarshal(r mus.Reader) (foo Foo, n int, err error) {
	foo.Num, n, err = varint.Int.Unmarshal(r)
	if err != nil {
		return
	}
	var n1 int
	foo.Str, n1, err = ord.String.Unmarshal(r)
	n += n1
	return
}

func (s fooSer) Size(foo Foo) (size int) {
	size = varint.Int.Size(foo.Num)
	return size + ord.String.Size(foo.Str)
}

func (s fooSer) Skip(r mus.Reader) (n int, err error) {
	n, err = varint.Int.Skip(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Skip(r)
	n += n1
	return
}

func TestDTS(t *testing.T) {
	t.Run("Marshal, Unmarshal, Size, Skip methods should succeed",
		func(t *testing.T) {
			var (
				foo    = Foo{Num: 11, Str: "hello world"}
				fooDTS = New[Foo](FooDTM, FooSer)
				size   = fooDTS.Size(foo)
				buf    = bytes.NewBuffer(make([]byte, 0, size))
			)
			n, err := fooDTS.Marshal(foo, buf)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n != size {
				t.Errorf("unexpected n, want %v, actual %v", size, n)
			}

			afoo, n, err := fooDTS.Unmarshal(buf)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n != size {
				t.Errorf("unexpected n, want %v, actual %v", size, n)
			}
			if !reflect.DeepEqual(afoo, foo) {
				t.Errorf("unexpected afoo, want %v, actual %v", foo, afoo)
			}

			buf.Reset()

			fooDTS.Marshal(foo, buf)
			n, err = fooDTS.Skip(buf)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n != size {
				t.Errorf("unexpected n, want %v, actual %v", size, n)
			}
		})

	t.Run("Marshal, UnmarshalDTM, UnmarshalData, Size, SkipDTM, SkipData methods should succeed",
		func(t *testing.T) {
			var (
				wantDTSize = 1
				foo        = Foo{Num: 11, Str: "hello world"}
				fooDTS     = New[Foo](FooDTM, FooSer)
				size       = fooDTS.Size(foo)
				buf        = bytes.NewBuffer(make([]byte, 0, size))
			)
			n, err := fooDTS.Marshal(foo, buf)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n != size {
				t.Errorf("unexpected n, want %v, actual %v", size, n)
			}

			dtm, n, err := DTMSer.Unmarshal(buf)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n != wantDTSize {
				t.Errorf("unexpected n, want %v, actual %v", wantDTSize, n)
			}
			if dtm != FooDTM {
				t.Errorf("unexpected dtm, want %v, actual %v", FooDTM, dtm)
			}

			afoo, n, err := fooDTS.UnmarshalData(buf)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n != size-wantDTSize {
				t.Errorf("unexpected n, want %v, actual %v", size-wantDTSize, n)
			}
			if !reflect.DeepEqual(afoo, foo) {
				t.Errorf("unexpected afoo, want %v, actual %v", foo, afoo)
			}

			buf.Reset()

			fooDTS.Marshal(foo, buf)
			_, err = DTMSer.Skip(buf)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			n, err = fooDTS.SkipData(buf)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n != size-wantDTSize {
				t.Errorf("unexpected n, want %v, actual %v", size-wantDTSize, n)
			}
		})

	t.Run("DTM method should return correct DTM", func(t *testing.T) {
		var (
			wantDTM = FooDTM

			fooDTS = New[Foo](FooDTM, nil)
		)

		dtm := fooDTS.DTM()
		if dtm != wantDTM {
			t.Errorf("unexpected dtm, want %v, actual %v", wantDTM, dtm)
		}
	})

	t.Run("Unmarshal should fail with ErrWrongDTM, if meets another DTM",
		func(t *testing.T) {
			var (
				actualDTM = FooDTM + 3

				wantDTSize = 1
				wantErr    = com.NewWrongDTMError(FooDTM, actualDTM)
				wantFoo    = Foo{}

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(actualDTM)
						return
					},
				)
				fooDTS = New[Foo](FooDTM, nil)
			)
			foo, n, err := fooDTS.Unmarshal(r)
			if err.Error() != wantErr.Error() {
				t.Errorf("unexpected error, want %v, actual %v", wantErr, err)
			}
			if !reflect.DeepEqual(foo, wantFoo) {
				t.Errorf("unexpected foo, want %v, actual %v", wantFoo, foo)
			}
			if n != wantDTSize {
				t.Errorf("unexpected n, want %v, actual %v", wantDTSize, n)
			}
		})

	t.Run("Skip should fail with ErrWrongDTM, if meets another DTM",
		func(t *testing.T) {
			var (
				actualDTM = FooDTM + 3

				wantDTSize = 1
				wantErr    = com.NewWrongDTMError(FooDTM, actualDTM)

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(actualDTM)
						return
					},
				)
				fooDTS = New[Foo](FooDTM, nil)
			)

			n, err := fooDTS.Skip(r)
			if err.Error() != wantErr.Error() {
				t.Errorf("unexpected error, want %v, actual %v", wantErr, err)
			}
			if n != wantDTSize {
				t.Errorf("unexpected n, want %v, actual %v", wantDTSize, n)
			}
		})

	t.Run("If MarshalDTM fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("write byte error")

				w = mock.NewWriter().RegisterWriteByte(func(c byte) error {
					return wantErr
				})
				fooDTS = New[Foo](FooDTM, nil)
			)
			_, err := fooDTS.Marshal(Foo{}, w)
			if err != wantErr {
				t.Errorf("unexpected error, want %v, actual %v", wantErr, err)
			}
		})

	t.Run("If UnmarshalDTM fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				fooDTS = New[Foo](FooDTM, nil)
			)
			foo, n, err := fooDTS.Unmarshal(r)
			if err != wantErr {
				t.Errorf("unexpected error, want %v, actual %v", wantErr, err)
			}
			if !reflect.DeepEqual(foo, Foo{}) {
				t.Errorf("unexpected foo, want %v, actual %v", Foo{}, foo)
			}
			if n != 0 {
				t.Errorf("unexpected n, want 0, actual %v", n)
			}
		})

	t.Run("If UnmarshalDTM fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				fooDTS = New[Foo](FooDTM, nil)
			)

			n, err := fooDTS.Skip(r)
			if err != wantErr {
				t.Errorf("unexpected error, want %v, actual %v", wantErr, err)
			}
			if n != 0 {
				t.Errorf("unexpected n, want 0, actual %v", n)
			}
		})

	t.Run("If varint.UnmarshalInt fails with an error, UnmarshalDTM should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
			)
			dtm, n, err := DTMSer.Unmarshal(r)
			if err != wantErr {
				t.Errorf("unexpected error, want %v, actual %v", wantErr, err)
			}
			if dtm != com.DTM(0) {
				t.Errorf("unexpected dtm, want 0, actual %v", dtm)
			}
			if n != 0 {
				t.Errorf("unexpected n, want 0, actual %v", n)
			}
		})
}
