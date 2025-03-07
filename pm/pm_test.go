package pm

import (
	"errors"
	"fmt"
	"testing"

	com "github.com/mus-format/common-go"
	mock "github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/ymz-ncnk/mok"

	com_testdata "github.com/mus-format/common-go/testdata"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
)

func TestPM(t *testing.T) {

	t.Run("Marshal should be able to marshal the nil pointer", func(t *testing.T) {
		var (
			wantN         = 1
			wantErr error = nil
			ptrMap        = com.NewPtrMap()
			w             = mock.NewWriter().RegisterWriteByte(
				func(c byte) (err error) {
					if c != byte(com.Nil) {
						err = fmt.Errorf("unexpected byte, want '%v' actual '%v'", com.Nil,
							c)
					}
					return
				},
			)
			mocks  = []*mok.Mock{w.Mock}
			n, err = NewPtrSer[int](ptrMap, nil, nil).Marshal(nil, w)
		)
		testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
	})

	t.Run("If marshal of the pointer Nil flag fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantN         = 0
				wantErr error = errors.New("pointer Nil flag marshal error")
				ptrMap        = com.NewPtrMap()
				w             = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return wantErr
					},
				)
				mocks  = []*mok.Mock{w.Mock}
				n, err = NewPtrSer[int](ptrMap, nil, nil).Marshal(nil, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If marshal of the pointer mapping flag fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantN         = 0
				wantErr error = errors.New("pointer mapping flag marshal error")
				ptrMap        = com.NewPtrMap()
				w             = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return wantErr
					},
				)
				num    = 2
				mocks  = []*mok.Mock{w.Mock}
				n, err = NewPtrSer[int](ptrMap, nil, nil).Marshal(&num, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If marshal of the pointer id fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantN         = 1
				wantErr error = errors.New("pointer id marshal error")
				ptrMap        = com.NewPtrMap()
				w             = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return nil
					},
				).RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				num    = 2
				mocks  = []*mok.Mock{w.Mock}
				n, err = NewPtrSer[int](ptrMap, nil, nil).Marshal(&num, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If baseSer.Marshal fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantN         = 3
				wantErr error = errors.New("marshaller error")
				ptrMap        = com.NewPtrMap()
				w             = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return nil
					},
				).RegisterWriteByte(
					func(c byte) error {
						return nil
					},
				)
				baseSer = mock.NewSerializer[int]().RegisterMarshal(
					func(t int, w muss.Writer) (n int, err error) {
						return 1, wantErr
					},
				)
				num    = 2
				mocks  = []*mok.Mock{w.Mock, baseSer.Mock}
				n, err = NewPtrSer[int](ptrMap, nil, baseSer).Marshal(&num, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If unmarshal of the pointer flag fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   *int = nil
				wantN        = 0
				wantErr      = errors.New("pointer flag unmarshal error")
				r            = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = NewPtrSer[int](nil, nil, nil).Unmarshal(r)
			)
			com_testdata.TestUnmarshalResults[*int](wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("If unmarshal of the pointer id fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   *int = nil
				wantN        = 1
				wantErr      = errors.New("pointer id unmarshal error")
				r            = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(com.Mapping)
						return
					},
				).RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = NewPtrSer[int](nil, com.NewReversePtrMap(), nil).Unmarshal(r)
			)
			com_testdata.TestUnmarshalResults[*int](wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If baseSer.Unmarshal fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   *int = nil
				wantN        = 3
				wantErr      = errors.New("Unmarshaller error")
				r            = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(com.Mapping)
						return
					},
				).RegisterReadByte(
					func() (b byte, err error) {
						b = 1
						return
					},
				)
				baseSer = mock.NewSerializer[int]().RegisterUnmarshal(
					func(r muss.Reader) (t int, n int, err error) {
						n = 1
						err = wantErr
						return
					},
				)
				revPtrMap = com.NewReversePtrMap()
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = NewPtrSer[int](nil, revPtrMap, baseSer).Unmarshal(r)
			)
			com_testdata.TestUnmarshalResults[*int](wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("Unmarshal should fail with com.ErrWrongFormat if meets unknown pointer flag",
		func(t *testing.T) {
			var (
				wantV   *int = nil
				wantN        = 0
				wantErr      = com.ErrWrongFormat
				r            = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(com.Mapping) + 1
						return
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = NewPtrSer[int](nil, com.NewReversePtrMap(), nil).Unmarshal(r)
			)
			com_testdata.TestUnmarshalResults[*int](wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("If unmarshal of the pointer flag fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("pointer flag unmarshal error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				mocks  = []*mok.Mock{r.Mock}
				n, err = NewPtrSer[int](nil, com.NewReversePtrMap(), nil).Skip(r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If unmarshal of the pointer id fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = errors.New("pointer id unmarshal error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(com.Mapping)
						return
					},
				).RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				mocks  = []*mok.Mock{r.Mock}
				n, err = NewPtrSer[int](nil, com.NewReversePtrMap(), nil).Skip(r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If baseSer.Skip fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 3
				wantErr = errors.New("Skipper error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(com.Mapping)
						return
					},
				).RegisterReadByte(
					func() (b byte, err error) {
						b = 1
						return
					},
				)
				baseSer = mock.NewSerializer[int]().RegisterSkip(
					func(r muss.Reader) (n int, err error) {
						n = 1
						err = wantErr
						return
					},
				)
				revPtrMap = com.NewReversePtrMap()
				mocks     = []*mok.Mock{r.Mock}
				n, err    = NewPtrSer[int](nil, revPtrMap, baseSer).Skip(r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Skip should fail with com.ErrWrongFormat if meets unknown pointer flag",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = com.ErrWrongFormat
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(com.Mapping) + 1
						return
					},
				)
				mocks  = []*mok.Mock{r.Mock}
				n, err = NewPtrSer[int](nil, com.NewReversePtrMap(), nil).Skip(r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

}
