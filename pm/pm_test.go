package pm

import (
	"errors"
	"fmt"
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/mus-format/mus-stream-go/varint"
	"github.com/ymz-ncnk/mok"
)

func MarshalPointerMappingStruct(v com_testdata.PointerMappingStruct,
	w muss.Writer) (n int, err error) {
	mp := com.NewPtrMap()
	n, err = MarshalPtr[int](v.A1, muss.MarshallerFn[int](varint.MarshalInt),
		mp,
		w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = MarshalPtr[int](v.A2, muss.MarshallerFn[int](varint.MarshalInt),
		mp,
		w)
	n += n1
	if err != nil {
		return
	}
	n1, err = MarshalPtr[int](v.B1, muss.MarshallerFn[int](varint.MarshalInt),
		mp,
		w)
	n += n1
	if err != nil {
		return
	}
	n1, err = MarshalPtr[int](v.B2, muss.MarshallerFn[int](varint.MarshalInt),
		mp,
		w)
	n += n1
	if err != nil {
		return
	}
	n1, err = MarshalPtr[string](v.C1, muss.MarshallerFn[string](MarshalStringVarint),
		mp,
		w)
	n += n1
	if err != nil {
		return
	}
	n1, err = MarshalPtr[string](v.C2, muss.MarshallerFn[string](MarshalStringVarint),
		mp,
		w)
	n += n1
	return
}

func UnmarshalPointerMappingStruct(r muss.Reader) (
	v com_testdata.PointerMappingStruct, n int, err error) {
	mp := com.NewReversePtrMap()
	v.A1, n, err = UnmarshalPtr[int](muss.UnmarshallerFn[int](varint.UnmarshalInt),
		mp,
		r)
	if err != nil {
		return
	}
	var n1 int
	v.A2, n1, err = UnmarshalPtr[int](muss.UnmarshallerFn[int](varint.UnmarshalInt),
		mp,
		r)
	n += n1
	if err != nil {
		return
	}
	v.B1, n1, err = UnmarshalPtr[int](muss.UnmarshallerFn[int](varint.UnmarshalInt),
		mp,
		r)
	n += n1
	if err != nil {
		return
	}
	v.B2, n1, err = UnmarshalPtr[int](muss.UnmarshallerFn[int](varint.UnmarshalInt),
		mp,
		r)
	n += n1
	if err != nil {
		return
	}
	v.C1, n1, err = UnmarshalPtr[string](
		muss.UnmarshallerFn[string](UnmarshalStringVarint),
		mp,
		r)
	n += n1
	if err != nil {
		return
	}
	v.C2, n1, err = UnmarshalPtr[string](
		muss.UnmarshallerFn[string](UnmarshalStringVarint),
		mp,
		r)
	n += n1
	return
}

func SizePointerMappingStruct(v com_testdata.PointerMappingStruct) (size int) {
	mp := com.NewPtrMap()
	size = SizePtr[int](v.A1, muss.SizerFn[int](varint.SizeInt), mp)
	size += SizePtr[int](v.A2, muss.SizerFn[int](varint.SizeInt), mp)
	size += SizePtr[int](v.B1, muss.SizerFn[int](varint.SizeInt), mp)
	size += SizePtr[int](v.B2, muss.SizerFn[int](varint.SizeInt), mp)
	size += SizePtr[string](v.C1, muss.SizerFn[string](SizeStringVarint), mp)
	return size + SizePtr[string](v.C2, muss.SizerFn[string](SizeStringVarint),
		mp)
}

func SkipPointerMappingStruct(r muss.Reader) (n int, err error) {
	mp := com.NewReversePtrMap()
	n, err = SkipPtr(muss.SkipperFn(varint.SkipInt), mp, r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = SkipPtr(muss.SkipperFn(varint.SkipInt), mp, r)
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipPtr(muss.SkipperFn(varint.SkipInt), mp, r)
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipPtr(muss.SkipperFn(varint.SkipInt), mp, r)
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipPtr(muss.SkipperFn(SkipStringVarint), mp, r)
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipPtr(muss.SkipperFn(SkipStringVarint), mp, r)
	n += n1
	return
}

func TestPM(t *testing.T) {

	t.Run("All MarshalPtr, UnmarshalPtr, SizePtr, SkipPtr functions must work correctly",
		func(t *testing.T) {
			var (
				m = muss.MarshallerFn[com_testdata.PointerMappingStruct](
					MarshalPointerMappingStruct)
				u = muss.UnmarshallerFn[com_testdata.PointerMappingStruct](
					UnmarshalPointerMappingStruct)
				s = muss.SizerFn[com_testdata.PointerMappingStruct](
					SizePointerMappingStruct)
				sk = muss.SkipperFn(SkipPointerMappingStruct)
			)
			testdata.Test[com_testdata.PointerMappingStruct](
				com_testdata.MakePointerMappingTestStruct(), m, u, s, t)
			testdata.TestSkip[com_testdata.PointerMappingStruct](
				com_testdata.MakePointerMappingTestStruct(), m, sk, s, t)
		})

	t.Run("MarshalPtr should be able to marshal the nil pointer", func(t *testing.T) {
		var (
			wantN         = 1
			wantErr error = nil
			mp            = com.NewPtrMap()
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
			n, err = MarshalPtr[int](nil, nil, mp, w)
		)
		testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
	})

	t.Run("If marshal of the pointer Nil flag fails with an error, MarshalPtr should return it",
		func(t *testing.T) {
			var (
				wantN         = 0
				wantErr error = errors.New("pointer Nil flag marshal error")
				mp            = com.NewPtrMap()
				w             = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return wantErr
					},
				)
				mocks  = []*mok.Mock{w.Mock}
				n, err = MarshalPtr[int](nil, nil, mp, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If marshal of the pointer Mapping flag fails with an error, MarshalPtr should return it",
		func(t *testing.T) {
			var (
				wantN         = 0
				wantErr error = errors.New("pointer Mapping flag marshal error")
				mp            = com.NewPtrMap()
				w             = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return wantErr
					},
				)
				num    = 2
				mocks  = []*mok.Mock{w.Mock}
				n, err = MarshalPtr[int](&num, nil, mp, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If marshal of the pointer id fails with an error, MarshalPtr should return it",
		func(t *testing.T) {
			var (
				wantN         = 1
				wantErr error = errors.New("pointer id marshal error")
				mp            = com.NewPtrMap()
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
				n, err = MarshalPtr[int](&num, nil, mp, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If Marshaller fails with an error, MarshalPtr should return it",
		func(t *testing.T) {
			var (
				wantN         = 3
				wantErr error = errors.New("marshaller error")
				mp            = com.NewPtrMap()
				w             = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return nil
					},
				).RegisterWriteByte(
					func(c byte) error {
						return nil
					},
				)
				m = mock.NewMarshaller[int]().RegisterMarshal(
					func(t int, w muss.Writer) (n int, err error) {
						return 1, wantErr
					},
				)
				num    = 2
				mocks  = []*mok.Mock{w.Mock}
				n, err = MarshalPtr[int](&num, m, mp, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If unmarshal of the pointer flag fails with an error, UnmarshalPtr should return it",
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
				v, n, err = UnmarshalPtr[int](nil, com.ReversePtrMap{}, r)
			)
			com_testdata.TestUnmarshalResults[*int](wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If unmarshal of the pointer id fails with an error, UnmarshalPtr should return it",
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
				v, n, err = UnmarshalPtr[int](nil, com.ReversePtrMap{}, r)
			)
			com_testdata.TestUnmarshalResults[*int](wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If Unmarshaller fails with an error, UnmarshalPtr should return it",
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
				u = mock.NewUnmarshaller[int]().RegisterUnmarshal(
					func(r muss.Reader) (t int, n int, err error) {
						n = 1
						err = wantErr
						return
					},
				)
				mp        = com.NewReversePtrMap()
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalPtr[int](u, mp, r)
			)
			com_testdata.TestUnmarshalResults[*int](wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("UnmarshalPtr should fail with com.ErrWrongFormat if meets unknown pointer flag",
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
				v, n, err = UnmarshalPtr[int](nil, com.ReversePtrMap{}, r)
			)
			com_testdata.TestUnmarshalResults[*int](wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If unmarshal of the pointer flag fails with an error, SkipPtr should return it",
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
				n, err = SkipPtr(nil, com.ReversePtrMap{}, r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If unmarshal of the pointer id fails with an error, SkipPtr should return it",
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
				n, err = SkipPtr(nil, com.ReversePtrMap{}, r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If Skipper fails with an error, SkipPtr should return it",
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
				sk = mock.NewSkipper().RegisterSkip(
					func(r muss.Reader) (n int, err error) {
						n = 1
						err = wantErr
						return
					},
				)
				mp     = com.NewReversePtrMap()
				mocks  = []*mok.Mock{r.Mock}
				n, err = SkipPtr(sk, mp, r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("SkipPtr should fail with com.ErrWrongFormat if meets unknown pointer flag",
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
				n, err = SkipPtr(nil, com.ReversePtrMap{}, r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

}

// StringVarint

func MarshalStringVarint(v string, w muss.Writer) (n int, err error) {
	return ord.MarshalString(v, muss.MarshallerFn[int](varint.MarshalInt), w)
}

func UnmarshalStringVarint(r muss.Reader) (v string,
	n int, err error) {
	return UnmarshalValidStringVarint(nil, false, r)
}

func UnmarshalValidStringVarint(lenVl com.Validator[int], skip bool, r muss.Reader) (
	v string, n int, err error) {
	return ord.UnmarshalValidString(muss.UnmarshallerFn[int](varint.UnmarshalInt),
		lenVl, skip, r)
}

func SizeStringVarint(v string) (n int) {
	return ord.SizeString(v, muss.SizerFn[int](varint.SizeInt))
}

func SkipStringVarint(r muss.Reader) (n int, err error) {
	return ord.SkipString(muss.UnmarshallerFn[int](varint.UnmarshalInt), r)
}
