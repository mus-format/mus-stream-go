package typed

import (
	"bytes"
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/test/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

const FooDTM com.DTM = 1

type Foo struct {
	Num int
	Str string
}

func TestTypedSer(t *testing.T) {
	t.Run("Marshal, Unmarshal, Size, Skip methods should succeed",
		func(t *testing.T) {
			var (
				foo     = Foo{Num: 11, Str: "hello world"}
				wantN   = 5
				serMock = mock.NewSerializer[Foo]().
					RegisterSize(func(v Foo) (size int) {
						return wantN
					}).
					RegisterMarshal(func(v Foo, w mus.Writer) (n int, err error) {
						return wantN, nil
					}).
					RegisterUnmarshal(func(r mus.Reader) (v Foo, n int, err error) {
						return foo, wantN, nil
					}).
					RegisterMarshal(func(v Foo, w mus.Writer) (n int, err error) {
						return wantN, nil
					}).
					RegisterSkip(func(r mus.Reader) (n int, err error) {
						return wantN, nil
					})
				fooTypedSer = NewTypedSer[Foo](FooDTM, serMock)
				dtmSize     = DTMSer.Size(FooDTM)
				size        = dtmSize + wantN
				buf         = bytes.NewBuffer(make([]byte, 0, size))
			)
			n, err := fooTypedSer.Marshal(foo, buf)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, size)

			afoo, n, err := fooTypedSer.Unmarshal(buf)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, size)
			asserterror.EqualDeep(t, afoo, foo)

			buf.Reset()

			fooTypedSer.Marshal(foo, buf)
			n, err = fooTypedSer.Skip(buf)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, size)
		})

	t.Run("Marshal, UnmarshalDTM, UnmarshalData, Size, SkipDTM, SkipData methods should succeed",
		func(t *testing.T) {
			var (
				foo     = Foo{Num: 11, Str: "hello world"}
				wantN   = 5
				serMock = mock.NewSerializer[Foo]().
					RegisterSize(func(v Foo) (size int) {
						return wantN
					}).
					RegisterMarshal(func(v Foo, w mus.Writer) (n int, err error) {
						return wantN, nil
					}).
					RegisterUnmarshal(func(r mus.Reader) (v Foo, n int, err error) {
						return foo, wantN, nil
					}).
					RegisterMarshal(func(v Foo, w mus.Writer) (n int, err error) {
						return wantN, nil
					}).
					RegisterSkip(func(r mus.Reader) (n int, err error) {
						return wantN, nil
					})
				fooTypedSer = NewTypedSer[Foo](FooDTM, serMock)
				dtmSize     = DTMSer.Size(FooDTM)
				size        = dtmSize + wantN
				buf         = bytes.NewBuffer(make([]byte, 0, size))
			)
			n, err := fooTypedSer.Marshal(foo, buf)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, size)

			dtm, n, err := DTMSer.Unmarshal(buf)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, dtmSize)
			asserterror.Equal(t, dtm, FooDTM)

			afoo, n, err := fooTypedSer.UnmarshalData(buf)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, wantN)
			asserterror.EqualDeep(t, afoo, foo)

			buf.Reset()

			fooTypedSer.Marshal(foo, buf)
			_, err = DTMSer.Skip(buf)
			assertfatal.EqualError(t, err, nil)

			n, err = fooTypedSer.SkipData(buf)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, wantN)
		})

	t.Run("DTM method should return correct DTM", func(t *testing.T) {
		var (
			wantDTM = FooDTM

			fooTypedSer = NewTypedSer[Foo](FooDTM, nil)
		)

		dtm := fooTypedSer.DTM()
		asserterror.Equal(t, dtm, wantDTM)
	})

	t.Run("Unmarshal should fail with ErrWrongDTM, if meets another DTM",
		func(t *testing.T) {
			var (
				actualDTM = FooDTM + 3

				wantErr = com.NewWrongDTMError(FooDTM, actualDTM)
				wantFoo = Foo{}

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(actualDTM)
						return
					},
				)
				fooTypedSer = NewTypedSer[Foo](FooDTM, nil)
				dtmSize     = DTMSer.Size(actualDTM)
			)
			foo, n, err := fooTypedSer.Unmarshal(r)
			asserterror.EqualError(t, err, wantErr)
			asserterror.EqualDeep(t, foo, wantFoo)
			asserterror.Equal(t, n, dtmSize)
		})

	t.Run("Skip should fail with ErrWrongDTM, if meets another DTM",
		func(t *testing.T) {
			var (
				actualDTM = FooDTM + 3

				wantErr = com.NewWrongDTMError(FooDTM, actualDTM)

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(actualDTM)
						return
					},
				)
				fooTypedSer = NewTypedSer[Foo](FooDTM, nil)
				dtmSize     = DTMSer.Size(actualDTM)
			)

			n, err := fooTypedSer.Skip(r)
			asserterror.EqualError(t, err, wantErr)
			asserterror.Equal(t, n, dtmSize)
		})

	t.Run("If MarshalDTM fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("write byte error")

				w = mock.NewWriter().RegisterWriteByte(func(c byte) error {
					return wantErr
				})
				fooTypedSer = NewTypedSer[Foo](FooDTM, nil)
			)
			_, err := fooTypedSer.Marshal(Foo{}, w)
			asserterror.EqualError(t, err, wantErr)
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
				fooTypedSer = NewTypedSer[Foo](FooDTM, nil)
			)
			foo, n, err := fooTypedSer.Unmarshal(r)
			asserterror.EqualError(t, err, wantErr)
			asserterror.EqualDeep(t, foo, Foo{})
			asserterror.Equal(t, n, 0)
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
				fooTypedSer = NewTypedSer[Foo](FooDTM, nil)
			)

			n, err := fooTypedSer.Skip(r)
			asserterror.EqualError(t, err, wantErr)
			asserterror.Equal(t, n, 0)
		})
}
