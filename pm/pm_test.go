package pm

import (
	"errors"
	"fmt"
	"testing"
	"unsafe"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	mock "github.com/mus-format/mus-stream-go/test/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
	"github.com/ymz-ncnk/mok"

	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/test"
)

func TestPM_Pointer(t *testing.T) {
	t.Run("Marshal should be able to marshal the nil pointer", func(t *testing.T) {
		var (
			ptrMap = com.NewPtrMap()
			w      = mock.NewWriter().RegisterWriteByte(
				func(c byte) (err error) {
					if c != byte(com.Nil) {
						err = fmt.Errorf("unexpected byte, want '%v' actual '%v'", com.Nil,
							c)
					}
					return
				},
			)
			ser  = NewPtrSer[int](ptrMap, nil, nil)
			want = test.MarshalResults{
				N:     1,
				Err:   nil,
				Mocks: []*mok.Mock{w.Mock},
			}
		)
		test.TestMarshalOnly(nil, w, ser, want, t)
	})

	t.Run("If marshal of the pointer Nil flag fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("pointer Nil flag marshal error")
				ptrMap  = com.NewPtrMap()
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return wantErr
					},
				)
				ser  = NewPtrSer[int](ptrMap, nil, nil)
				want = test.MarshalResults{
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{w.Mock},
				}
			)
			test.TestMarshalOnly(nil, w, ser, want, t)
		})

	t.Run("If marshal of the pointer mapping flag fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("pointer mapping flag marshal error")
				ptrMap  = com.NewPtrMap()
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return wantErr
					},
				)
				num  = 2
				ser  = NewPtrSer[int](ptrMap, nil, nil)
				want = test.MarshalResults{
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{w.Mock},
				}
			)
			test.TestMarshalOnly(&num, w, ser, want, t)
		})

	t.Run("If marshal of the pointer id fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("pointer id marshal error")
				ptrMap  = com.NewPtrMap()
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return nil
					},
				).RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				num  = 2
				ser  = NewPtrSer[int](ptrMap, nil, nil)
				want = test.MarshalResults{
					N:     1,
					Err:   wantErr,
					Mocks: []*mok.Mock{w.Mock},
				}
			)
			test.TestMarshalOnly(&num, w, ser, want, t)
		})

	t.Run("If baseSer.Marshal fails with an error, Marshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("marshaller error")
				ptrMap  = com.NewPtrMap()
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) (err error) {
						return nil
					},
				).RegisterWriteByte(
					func(c byte) error {
						return nil
					},
				)
				baseSer = mock.NewSerializer[int]().RegisterMarshal(
					func(t int, w mus.Writer) (n int, err error) {
						return 1, wantErr
					},
				)
				num  = 2
				ser  = NewPtrSer(ptrMap, nil, baseSer)
				want = test.MarshalResults{
					N:     3,
					Err:   wantErr,
					Mocks: []*mok.Mock{w.Mock, baseSer.Mock},
				}
			)
			test.TestMarshalOnly(&num, w, ser, want, t)
		})

	t.Run("If unmarshal of the pointer flag fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("pointer flag unmarshal error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				ser  = NewPtrSer[int](nil, nil, nil)
				want = test.UnmarshalResults[*int]{
					V:     nil,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If unmarshal of the pointer id fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
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
				ser  = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
				want = test.UnmarshalResults[*int]{
					V:     nil,
					N:     1,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If baseSer.Unmarshal fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("Unmarshaller error")
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
				baseSer = mock.NewSerializer[int]().RegisterUnmarshal(
					func(r mus.Reader) (t int, n int, err error) {
						n = 1
						err = wantErr
						return
					},
				)
				revPtrMap = com.NewReversePtrMap()
				ser       = NewPtrSer(nil, revPtrMap, baseSer)
				want      = test.UnmarshalResults[*int]{
					V:     nil,
					N:     3,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock, baseSer.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("Unmarshal should fail with com.ErrWrongFormat if meets unknown pointer flag",
		func(t *testing.T) {
			var (
				wantErr = com.ErrWrongFormat
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(com.Mapping) + 1
						return
					},
				)
				ser  = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
				want = test.UnmarshalResults[*int]{
					V:     nil,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If unmarshal of the pointer flag fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("pointer flag unmarshal error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
				ser  = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
				want = test.SkipResults{
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestSkipOnly(r, ser, want, t)
		})

	t.Run("If unmarshal of the pointer id fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
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
				ser  = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
				want = test.SkipResults{
					N:     1,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestSkipOnly(r, ser, want, t)
		})

	t.Run("If baseSer.Skip fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
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
					func(r mus.Reader) (n int, err error) {
						n = 1
						err = wantErr
						return
					},
				)
				revPtrMap = com.NewReversePtrMap()
				ser       = NewPtrSer(nil, revPtrMap, baseSer)
				want      = test.SkipResults{
					N:     3,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock, baseSer.Mock},
				}
			)
			test.TestSkipOnly(r, ser, want, t)
		})

	t.Run("Skip should fail with com.ErrWrongFormat if meets unknown pointer flag",
		func(t *testing.T) {
			var (
				wantErr = com.ErrWrongFormat
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = byte(com.Mapping) + 1
						return
					},
				)
				ser  = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
				want = test.SkipResults{
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestSkipOnly(r, ser, want, t)
		})
}

func TestPM_Wrapper(t *testing.T) {
	t.Run("wrapper serializer should work correctly",
		func(t *testing.T) {
			var (
				st, baseSer = test.PtrStructSerData(t)
				ptrMap      = com.NewPtrMap()
				revPtrMap   = com.NewReversePtrMap()
				ser         = Wrap(ptrMap, revPtrMap, newPtrStructSer(ptrMap, revPtrMap,
					baseSer))
			)
			test.Test([]ctest.PtrStruct{st}, ser, t)
			test.TestSkip([]ctest.PtrStruct{st}, ser, t)
		})

	t.Run("Marshal should call ser.Marshal and empty the ptrMap",
		func(t *testing.T) {
			var (
				wantV  byte = 1
				wantN       = 1
				ptrMap      = com.NewPtrMap()
				ptrSer      = mock.NewSerializer[byte]().RegisterMarshal(
					func(v byte, w mus.Writer) (n int, err error) {
						ptrMap.Put(unsafe.Pointer(&v))
						n = wantN
						err = w.WriteByte(v)
						return
					},
				)
				ser = Wrap(ptrMap, nil, ptrSer)
				w   = mock.NewWriter().RegisterWriteByte(
					func(b byte) error {
						if b != wantV {
							return fmt.Errorf("unexpected v, want %v, actual %v", wantV, b)
						}
						return nil
					},
				)
				mocks = []*mok.Mock{ptrSer.Mock, w.Mock}
			)
			n, err := ser.Marshal(wantV, w)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, wantN)
			assertfatal.Equal(t, ptrMap.Len(), 0)
			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap)
		})

	t.Run("Unmarshal should call ser.Unmarshal and empty the revPtrMap",
		func(t *testing.T) {
			var (
				wantV     byte = 1
				wantN          = 1
				revPtrMap      = com.NewReversePtrMap()
				ptrSer         = mock.NewSerializer[byte]().RegisterUnmarshal(
					func(r mus.Reader) (v byte, n int, err error) {
						v, err = r.ReadByte()
						if err != nil {
							return
						}
						if v != wantV {
							err = fmt.Errorf("unexpected v, want %v, actual %v", wantV, v)
							return
						}
						n = wantN
						revPtrMap.Put(1, unsafe.Pointer(&v))
						return
					},
				)
				ser = Wrap(nil, revPtrMap, ptrSer)
				r   = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = wantV
						return
					},
				)
				mocks = []*mok.Mock{ptrSer.Mock, r.Mock}
			)
			v, n, err := ser.Unmarshal(r)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, v, wantV)
			asserterror.Equal(t, n, wantN)
			assertfatal.Equal(t, revPtrMap.Len(), 0)
			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap)
		})

	t.Run("Size should call ser.Size and empty the ptrMap",
		func(t *testing.T) {
			var (
				wantV    byte = 1
				wantSize      = 1
				ptrMap        = com.NewPtrMap()
				ptrSer        = mock.NewSerializer[byte]().RegisterSize(
					func(v byte) (size int) {
						ptrMap.Put(unsafe.Pointer(&v))
						return wantSize
					},
				)
				ser   = Wrap(ptrMap, nil, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			size := ser.Size(wantV)
			asserterror.Equal(t, size, wantSize)
			assertfatal.Equal(t, ptrMap.Len(), 0)
			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap)
		})

	t.Run("Skip should call ser.Skip and empty the revPtrMap",
		func(t *testing.T) {
			var (
				wantV     byte = 1
				wantN          = 1
				revPtrMap      = com.NewReversePtrMap()
				ptrSer         = mock.NewSerializer[byte]().RegisterSkip(
					func(r mus.Reader) (n int, err error) {
						v, err := r.ReadByte()
						if err != nil {
							return
						}
						if v != wantV {
							err = fmt.Errorf("unexpected v, want %v, actual %v", wantV, v)
							return
						}
						revPtrMap.Put(1, unsafe.Pointer(&v))
						n = wantN
						return
					},
				)
				ser = Wrap(nil, revPtrMap, ptrSer)
				r   = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = wantV
						return
					},
				)
				mocks = []*mok.Mock{ptrSer.Mock, r.Mock}
			)
			n, err := ser.Skip(r)
			assertfatal.EqualError(t, err, nil)
			asserterror.Equal(t, n, wantN)
			assertfatal.Equal(t, revPtrMap.Len(), 0)
			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap)
		})
}

func newPtrStructSer(ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap,
	baseSer mus.Serializer[int],
) mus.Serializer[ctest.PtrStruct] {
	return ptrStructSer{NewPtrSer(ptrMap, revPtrMap, baseSer)}
}

type ptrStructSer struct {
	intPtrSer mus.Serializer[*int]
}

func (s ptrStructSer) Marshal(v ctest.PtrStruct, w mus.Writer) (n int,
	err error,
) {
	n, err = s.intPtrSer.Marshal(v.A1, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.intPtrSer.Marshal(v.A2, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = s.intPtrSer.Marshal(v.A3, w)
	n += n1
	return
}

func (s ptrStructSer) Unmarshal(r mus.Reader) (v ctest.PtrStruct, n int,
	err error,
) {
	v.A1, n, err = s.intPtrSer.Unmarshal(r)
	if err != nil {
		return
	}
	var n1 int
	v.A2, n1, err = s.intPtrSer.Unmarshal(r)
	n += n1
	if err != nil {
		return
	}
	v.A3, n1, err = s.intPtrSer.Unmarshal(r)
	n += n1
	return
}

func (s ptrStructSer) Size(v ctest.PtrStruct) (size int) {
	size = s.intPtrSer.Size(v.A1)
	size += s.intPtrSer.Size(v.A2)
	size += s.intPtrSer.Size(v.A3)

	return
}

func (s ptrStructSer) Skip(r mus.Reader) (n int, err error) {
	n, err = s.intPtrSer.Skip(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.intPtrSer.Skip(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = s.intPtrSer.Skip(r)
	n += n1
	return
}
