package ord

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	com_mock "github.com/mus-format/common-go/testdata/mock"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/mus-format/mus-stream-go/varint"
	"github.com/ymz-ncnk/mok"
)

func TestOrd(t *testing.T) {

	t.Run("bool", func(t *testing.T) {

		t.Run("Bool serializer should work correctly",
			func(t *testing.T) {
				ser := Bool
				testdata.Test[bool](com_testdata.BoolTestCases, ser, t)
				testdata.TestSkip[bool](com_testdata.BoolTestCases, ser, t)
			})

		t.Run("If Writer fails to write a byte, Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("write error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return wantErr },
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = Bool.Marshal(true, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read a byte, Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   bool = false
					wantN        = 0
					wantErr      = errors.New("read byte error")
					r            = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = Bool.Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV   bool = false
					wantN        = 1
					wantErr      = com.ErrWrongFormat
					r            = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = Bool.Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read a byte, Skip should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("read byte error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = Bool.Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrWrongFormat
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = Bool.Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

	})

	t.Run("string", func(t *testing.T) {

		t.Run("String seralizer should work correctly",
			func(t *testing.T) {
				ser := String
				testdata.Test[string](com_testdata.StringTestCases, ser, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, ser, t)
			})

		t.Run("If Writer fails to write string length, Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("marshal length error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = String.Marshal("hello world", w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read string length, Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   string = ""
					wantN          = 0
					wantErr        = errors.New("unmarshal length error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = String.Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 10
					wantErr   = com.ErrNegativeLength
					r         = LengthReader(-1)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = String.Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read string content, Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   string = ""
					wantN          = 1 + 2
					wantErr        = errors.New("unmarshal string content error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterRead(
						func(p []byte) (n int, err error) {
							p[0] = 110
							p[1] = 111
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = String.Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read string length, Skip should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = String.Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN   = 10
					wantErr = com.ErrNegativeLength
					r       = LengthReader(-1)
					mocks   = []*mok.Mock{r.Mock}
					n, err  = String.Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read string content, Skip should return error",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = errors.New("skip string content error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterReadByte(
						func() (b byte, err error) { return },
					).RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = String.Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("ValidString serializer should work", func(t *testing.T) {
			ser := NewValidStringSer(nil)
			testdata.Test[string](com_testdata.StringTestCases, ser, t)
			testdata.TestSkip[string](com_testdata.StringTestCases, ser, t)
		})

		t.Run("If lenVl returns an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV      = ""
					wantLength = math.MaxInt64 - 1
					wantN      = 9
					wantErr    = errors.New("lenVl validator error")
					lenVl      = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != wantLength {
								t.Errorf("unexpected length, want '%v' actual '%v'", wantLength, v)
							}
							return wantErr
						},
					)
					r         = LengthReader(wantLength)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidStringSer(lenVl).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If string length == 0 lenVl should work", func(t *testing.T) {
			var (
				wantV                        = ""
				wantN                        = 1
				wantErr                      = errors.New("empty string")
				lenVl   com.ValidatorFn[int] = func(t int) (err error) {
					return wantErr
				}
				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, nil
					},
				)
				v, n, err = NewValidStringSer(lenVl).Unmarshal(r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("If Reader fails to read string length, valid Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   string = ""
					wantN          = 0
					wantErr        = errors.New("unmarshal length error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidStringSer(nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Valid Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 10
					wantErr   = com.ErrNegativeLength
					r         = LengthReader(-1)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidStringSer(nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read string content, valid Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   string = ""
					wantN          = 1 + 2
					wantErr        = errors.New("unmarshal string content error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterRead(
						func(p []byte) (n int, err error) {
							p[0] = 110
							p[1] = 111
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidStringSer(nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

	})

	t.Run("pointer", func(t *testing.T) {

		t.Run("Pointer seralizer should work correctly",
			func(t *testing.T) {
				var (
					ptr, baseSer = testdata.PtrSerData(t)
					mocks        = []*mok.Mock{baseSer.Mock}
					ser          = NewPtrSer[int](baseSer)
				)
				testdata.Test[*int]([]*int{ptr}, ser, t)
				testdata.TestSkip[*int]([]*int{ptr}, ser, t)

				if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
					t.Error(infomap)
				}
			})

		t.Run("Ptr serializer should work correctly with nil pointer",
			func(t *testing.T) {
				ser := NewPtrSer[string](String)
				testdata.Test[*string]([]*string{nil}, ser, t)
				testdata.TestSkip[*string]([]*string{nil}, ser, t)
			})

		t.Run("If Writer fails to write nil flag == 0, Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("write nil flag error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = NewPtrSer[string](nil).Marshal(nil, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Writer fails to write nil flag == 1, Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("write nil flag error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					str    = "str"
					strPtr = &str
					n, err = NewPtrSer[string](String).Marshal(strPtr, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Writer fails to write pointer content, Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = errors.New("Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return nil
						},
					).RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					str    = "str"
					strPtr = &str
					n, err = NewPtrSer[string](String).Marshal(strPtr, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read nil flag, Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   *string = nil
					wantN           = 0
					wantErr         = errors.New("read nil flag error")
					r               = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewPtrSer[string](nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("UnmarshalPtr should return ErrWrongFormat if meets wrong format", func(t *testing.T) {
			var (
				wantV   *string = nil
				wantN           = 1
				wantErr         = com.ErrWrongFormat
				r               = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 2, nil
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = NewPtrSer[string](nil).Unmarshal(r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
				t)
		})

		t.Run("If base serializer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   *string = nil
					wantN           = 1
					wantErr         = errors.New("base serializer error")
					r               = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, nil
						},
					)
					baseSer = mock.NewSerializer[string]().RegisterUnmarshal(
						func(r muss.Reader) (t string, n int, err error) {
							return "", 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, baseSer.Mock}
					v, n, err = NewPtrSer[string](baseSer).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If Reader fails to read nil flag, Skip should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("read nil flag error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = NewPtrSer[string](nil).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrWrongFormat
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = NewPtrSer[string](nil).Skip(r)
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
							return 0, nil
						},
					)
					baseSer = mock.NewSerializer[string]().RegisterSkip(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = NewPtrSer[string](baseSer).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

	})

	t.Run("byte_slice", func(t *testing.T) {

		t.Run("ByteSlice seializer should work correctly for empty slice",
			func(t *testing.T) {
				var (
					sl  = []byte{}
					ser = ByteSlice
				)
				testdata.Test[[]byte]([][]byte{sl}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
			})

		t.Run("ByteSlice serializer should work correctly for not empty slice",
			func(t *testing.T) {
				var (
					sl  = []byte{1, 2, 45, 255, 123, 70, 0, 0}
					ser = ByteSlice
				)
				testdata.Test[[]byte]([][]byte{sl}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
			})

		t.Run("If Writer fails to write slice length, Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("marshal length error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return wantErr },
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = ByteSlice.Marshal([]byte{1}, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read slice length, Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 0
					wantErr        = errors.New("unmarshal length error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = ByteSlice.Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 10
					wantErr          = com.ErrNegativeLength
					r                = LengthReader(-1)
					mocks            = []*mok.Mock{r.Mock}
					v, n, err        = ByteSlice.Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If Reader fails to read slice content, Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   []byte = make([]byte, 2)
					wantN          = 1
					wantErr        = errors.New("read slice content error")

					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return byte(len(wantV)), nil
						},
					).RegisterRead(
						func(p []byte) (n int, err error) { return 0, wantErr },
					)
					v, n, err = ByteSlice.Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil, t)
			})

		t.Run("If Reader fails to read slice length, Skip should return it",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = ByteSlice.Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN   = 10
					wantErr = com.ErrNegativeLength
					r       = LengthReader(-1)
					mocks   = []*mok.Mock{r.Mock}
					n, err  = ByteSlice.Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read slice content, Skip should return error",
			func(t *testing.T) {
				var (
					wantV   []byte = make([]byte, 2)
					wantN          = 1
					wantErr        = errors.New("read slice content error")

					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return byte(len(wantV)), nil
						},
					).RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
						// func(p []byte) (n int, err error) { return 0, wantErr },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = ByteSlice.Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("ValidByteSlice seializer should work correctly for empty slice",
			func(t *testing.T) {
				var (
					sl  = []byte{}
					ser = NewValidByteSliceSer(nil)
				)
				testdata.Test[[]byte]([][]byte{sl}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
			})

		t.Run("ValidByteSlice seializer should work correctly for not empty slice",
			func(t *testing.T) {
				var (
					sl  = []byte{1, 2, 3}
					ser = NewValidByteSliceSer(nil)
				)
				testdata.Test[[]byte]([][]byte{sl}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
			})

		t.Run("If lenVl returns an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 1
					wantErr        = errors.New("lenVl error")
					lenVl          = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 4, nil
						},
					)
					mocks     = []*mok.Mock{lenVl.Mock}
					v, n, err = NewValidByteSliceSer(lenVl).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read slice length, ValidByteSlice.Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 0
					wantErr        = errors.New("unmarshal length error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidByteSliceSer(nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("ValidByteSlice.Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 10
					wantErr          = com.ErrNegativeLength
					r                = LengthReader(-1)
					mocks            = []*mok.Mock{r.Mock}
					v, n, err        = NewValidByteSliceSer(nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If Reader fails to read slice content, ValidByteSlice.Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   []byte = make([]byte, 2)
					wantN          = 1
					wantErr        = errors.New("read slice content error")

					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return byte(len(wantV)), nil
						},
					).RegisterRead(
						func(p []byte) (n int, err error) { return 0, wantErr },
					)
					v, n, err = NewValidByteSliceSer(nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil, t)
			})

	})

	t.Run("array", func(t *testing.T) {

		t.Run("Array serializer should work correctly", func(t *testing.T) {
			var (
				arr, elemSer = testdata.ArraySerData(t)
				mocks        = []*mok.Mock{elemSer.Mock}
				ser          = NewArraySer[[3]int, int](3,
					elemSer)
			)
			testdata.Test[[3]int]([][3]int{arr}, ser, t)
			testdata.TestSkip[[3]int]([][3]int{arr}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("Unmarshal of the too large array should return ErrTooLargeLength",
			func(t *testing.T) {
				var (
					wantV   [3]int = [3]int{0, 0, 0}
					wantN          = 1
					wantErr        = com.ErrTooLargeLength
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 4, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewArraySer[[3]int, int](3, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Valid array serializer should work correctly", func(t *testing.T) {
			var (
				arr, elemSer = testdata.ArraySerData(t)
				mocks        = []*mok.Mock{elemSer.Mock}
				ser          = NewValidArraySer[[3]int, int](3,
					elemSer, nil)
			)
			testdata.Test[[3]int]([][3]int{arr}, ser, t)
			testdata.TestSkip[[3]int]([][3]int{arr}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("Valid Unmarshal of the too large array should return ErrTooLargeLength",
			func(t *testing.T) {
				var (
					wantV   [3]int = [3]int{0, 0, 0}
					wantN          = 1
					wantErr        = com.ErrTooLargeLength
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 4, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidArraySer[[3]int, int](3, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If elemVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV    [3]int = [3]int{0, 0, 0}
					wantElem        = 11
					wantN           = 2
					wantErr         = errors.New("elemVl error")
					r               = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					elemSer = mock.NewSerializer[int]().RegisterUnmarshal(
						func(r muss.Reader) (t int, n int, err error) {
							return 11, 1, nil
						},
					)
					elemVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != wantElem {
								return fmt.Errorf("unexpected v, want %v actual %v", wantElem, v)
							}
							return wantErr
						},
					)
					ser   = NewValidArraySer[[3]int, int](3, elemSer, elemVl)
					mocks = []*mok.Mock{r.Mock, elemSer.Mock, elemVl.Mock}
				)
				v, n, err := ser.Unmarshal(r)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

	})

	t.Run("slice", func(t *testing.T) {

		t.Run("Slice serializer should work correctly for empty slice",
			func(t *testing.T) {
				var (
					sl  = []string{}
					ser = NewSliceSer[string](String)
				)
				testdata.Test[[]string]([][]string{sl}, ser, t)
				testdata.TestSkip[[]string]([][]string{sl}, ser, t)
			})

		t.Run("Slice serializer should work correctly for not empty slice",
			func(t *testing.T) {
				var (
					sl, elemSer = testdata.SliceSerData(t)
					mocks       = []*mok.Mock{elemSer.Mock}
					ser         = NewSliceSer[string](elemSer)
				)
				testdata.Test[[]string]([][]string{sl}, ser, t)
				testdata.TestSkip[[]string]([][]string{sl}, ser, t)

				if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
					t.Error(infomap)
				}
			})

		t.Run("If Writer fails to write slice length, Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("marshal length error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return wantErr },
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = NewSliceSer[uint](nil).Marshal([]uint{1}, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Marshaller fails with an error, Marshal should return it",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = errors.New("Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return nil },
					)
					elemSer = mock.NewSerializer[uint]().RegisterMarshal(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock, elemSer.Mock}
					n, err = NewSliceSer[uint](elemSer).Marshal([]uint{1}, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read slice length, Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   []string = nil
					wantN            = 0
					wantErr          = errors.New("unmarshal length error")
					r                = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewSliceSer[string](nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     []string = nil
					wantN              = 10
					wantErr            = com.ErrNegativeLength
					r                  = LengthReader(-1)
					mocks              = []*mok.Mock{r.Mock}
					v, n, err          = NewSliceSer[string](nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If elemSer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []string = make([]string, 2)
					wantN            = 3
					wantErr          = errors.New("Unmarshaller error")
					r                = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					elemSer = mock.NewSerializer[string]().RegisterUnmarshal(
						func(r muss.Reader) (t string, n int, err error) {
							return "", 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewSliceSer[string](elemSer).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read slice length, Skip should return it",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = NewSliceSer[string](nil).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN   = 10
					wantErr = com.ErrNegativeLength
					r       = LengthReader(-1)
					mocks   = []*mok.Mock{r.Mock}
					n, err  = NewSliceSer[string](nil).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If elemSer fails with an error, Skip should return it",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = errors.New("Unmarshaller error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					elemSer = mock.NewSerializer[string]().RegisterSkip(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock, elemSer.Mock}
					n, err = NewSliceSer[string](elemSer).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("ValidSlice serializer should work correctly for empty slice",
			func(t *testing.T) {
				var (
					sl  = []string{}
					ser = NewValidSliceSer[string](nil, nil, nil)
				)
				testdata.Test[[]string]([][]string{sl}, ser, t)
				testdata.TestSkip[[]string]([][]string{sl}, ser, t)
			})

		t.Run("ValidSlice serializer should work correctly for not empty slice",
			func(t *testing.T) {
				var (
					sl, elemSer = testdata.SliceSerData(t)
					mocks       = []*mok.Mock{elemSer.Mock}
					ser         = NewValidSliceSer[string](elemSer, nil, nil)
				)
				testdata.Test[[]string]([][]string{sl}, ser, t)
				testdata.TestSkip[[]string]([][]string{sl}, ser, t)

				if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
					t.Error(infomap)
				}
			})

		t.Run("If Reader fails to read slice length, ValidSlice.Unmarshal should return error",
			func(t *testing.T) {
				var (
					wantV   []string = nil
					wantN            = 0
					wantErr          = errors.New("unmarshal length error")
					r                = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidSliceSer[string](nil, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("ValidSlice.Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     []string = nil
					wantN              = 10
					wantErr            = com.ErrNegativeLength
					r                  = LengthReader(-1)
					mocks              = []*mok.Mock{r.Mock}
					v, n, err          = NewValidSliceSer[string](nil, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If elemSer fails with an error, ValidSlice.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []string = make([]string, 2)
					wantN            = 3
					wantErr          = errors.New("Unmarshaller error")
					r                = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					elemSer = mock.NewSerializer[string]().RegisterUnmarshal(
						func(r muss.Reader) (t string, n int, err error) {
							return "", 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, elemSer.Mock}
					v, n, err = NewValidSliceSer[string](elemSer, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If lenVl validator returns an error, ValidSlice.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = nil
					wantN          = 1
					wantErr        = errors.New("lenVl validator error")
					lenVl          = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 4 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 4, v)
							}
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidSliceSer[uint](nil, lenVl, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If elemVl returns an error, ValidSlice.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = []uint{10, 0, 0}
					wantN          = 3
					wantErr        = errors.New("Validator error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 3, nil },
					).RegisterReadByte(
						func() (b byte, err error) { return 10, nil },
					)
					elemVl = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return nil
						},
					).RegisterValidate(
						func(v uint) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return wantErr
						},
					)
					elemSer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					).RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 2, 1, nil
						},
					)
					mocks     = []*mok.Mock{elemVl.Mock, elemSer.Mock}
					v, n, err = NewValidSliceSer[uint](elemSer, nil, elemVl).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

	})

	t.Run("map", func(t *testing.T) {

		t.Run("Map serializer should work correctly for empty map",
			func(t *testing.T) {
				var (
					mp  = map[string]int{}
					ser = NewMapSer[string, int](nil, nil)
				)
				testdata.Test[map[string]int]([]map[string]int{mp}, ser, t)
				testdata.TestSkip[map[string]int]([]map[string]int{mp}, ser, t)
			})

		t.Run("Map serializer should work correctly",
			func(t *testing.T) {
				var (
					mp, keySer, valueSer = testdata.MapSerData(t)
					ser                  = NewMapSer[string, int](keySer, valueSer)
					mocks                = []*mok.Mock{keySer.Mock, valueSer.Mock}
				)
				testdata.Test[map[string]int]([]map[string]int{mp}, ser, t)
				testdata.TestSkip[map[string]int]([]map[string]int{mp}, ser, t)

				if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
					t.Error(infomap)
				}
			})

		t.Run("If Writer fails to write map length, Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("marshal length error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return wantErr },
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = NewMapSer[uint, uint](nil, nil).Marshal(nil, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If keySer fails with an error, Marshal should return it",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = errors.New("key Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return nil },
					)
					keySer = mock.NewSerializer[uint]().RegisterMarshal(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock, keySer.Mock}
					n, err = NewMapSer[uint, uint](keySer, nil).Marshal(map[uint]uint{1: 1}, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If valueSer fails with an error, Marshal should return it",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = errors.New("value Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return nil },
					)
					keySer = mock.NewSerializer[uint]().RegisterMarshal(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, nil
						},
					)
					valueSer = mock.NewSerializer[uint]().RegisterMarshal(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock, keySer.Mock, valueSer.Mock}
					n, err = NewMapSer[uint, uint](keySer, valueSer).Marshal(map[uint]uint{1: 1}, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read map length, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 0
					wantErr               = errors.New("unmarshal length error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewMapSer[uint, uint](nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 10
					wantErr                 = com.ErrNegativeLength
					r                       = LengthReader(-1)
					mocks                   = []*mok.Mock{r.Mock}
					v, n, err               = NewMapSer[uint, uint](nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If keySer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 3
					wantErr = errors.New("keySer error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					keySer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, keySer.Mock}
					v, n, err = NewMapSer[uint, uint](keySer, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If valueSer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 4
					wantErr = errors.New("valueSer error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					keySer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 1, 1, nil
						},
					)
					valueSer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {

							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, keySer.Mock, valueSer.Mock}
					v, n, err = NewMapSer[uint, uint](keySer, valueSer).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If Reader fails to read map length, Skip should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = NewMapSer[uint, uint](nil, nil).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN   = 10
					wantErr = com.ErrNegativeLength
					r       = LengthReader(-1)
					mocks   = []*mok.Mock{r.Mock}
					n, err  = NewMapSer[uint, uint](nil, nil).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If keySer fails with an error, Skip should return it",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = errors.New("keySer error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					keySer = mock.NewSerializer[uint]().RegisterSkip(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock, keySer.Mock}
					n, err = NewMapSer[uint, uint](keySer, nil).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If valueSer fails with an error, Skip should return it",
			func(t *testing.T) {
				var (
					wantN   = 4
					wantErr = errors.New("valueSer error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					keySer = mock.NewSerializer[uint]().RegisterSkip(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					valueSer = mock.NewSerializer[uint]().RegisterSkip(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock, keySer.Mock, valueSer.Mock}
					n, err = NewMapSer[uint, uint](keySer, valueSer).Skip(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("ValidMap serializer should work correctly",
			func(t *testing.T) {
				var (
					mp, keySer, valueSer = testdata.MapSerData(t)
					mocks                = []*mok.Mock{keySer.Mock, valueSer.Mock}
					ser                  = NewValidMapSer[string, int](keySer, valueSer, nil, nil, nil)
				)
				testdata.Test[map[string]int]([]map[string]int{mp}, ser, t)
				testdata.TestSkip[map[string]int]([]map[string]int{mp}, ser, t)

				if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
					t.Error(infomap)
				}
			})

		t.Run("If Writer fails to write map length, ValidMap.Marshal should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("marshal length error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return wantErr },
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = NewValidMapSer[uint, uint](nil, nil, nil, nil, nil).Marshal(nil, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If keySer fails with an error, ValidMap.Marshal should return it",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = errors.New("key Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return nil },
					)
					keySer = mock.NewSerializer[uint]().RegisterMarshal(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock, keySer.Mock}
					n, err = NewValidMapSer[uint, uint](keySer, nil, nil, nil, nil).Marshal(map[uint]uint{1: 1}, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If valueSer fails with an error, ValidMap.Marshal should return it",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = errors.New("value Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return nil },
					)
					keySer = mock.NewSerializer[uint]().RegisterMarshal(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, nil
						},
					)
					valueSer = mock.NewSerializer[uint]().RegisterMarshal(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock, keySer.Mock, valueSer.Mock}
					n, err = NewValidMapSer[uint, uint](keySer, valueSer, nil, nil, nil).Marshal(map[uint]uint{1: 1}, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read map length, ValidMap.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 0
					wantErr               = errors.New("unmarshal length error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidMapSer[uint, uint](nil, nil, nil, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("ValidMap.Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 10
					wantErr                 = com.ErrNegativeLength
					r                       = LengthReader(-1)
					mocks                   = []*mok.Mock{r.Mock}
					v, n, err               = NewValidMapSer[uint, uint](nil, nil, nil, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If keySer fails with an error, ValidMap.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 3
					wantErr = errors.New("key Unmarshaller error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					keySer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, keySer.Mock}
					v, n, err = NewValidMapSer[uint, uint](keySer, nil, nil, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If valueSer fails with an error, ValidMap.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 4
					wantErr = errors.New("valueSer error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					keySer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 1, 1, nil
						},
					)
					valueSer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {

							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, keySer.Mock, valueSer.Mock}
					v, n, err = NewValidMapSer[uint, uint](keySer, valueSer, nil, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If lenVl validator returns an error, ValidMap.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					mapLen                = 4
					wantErr               = errors.New("lenVl validator error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return byte(mapLen), nil },
					)
					lenVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != mapLen {
								t.Errorf("unexpected v, want '%v' actual '%v'", mapLen, v)
							}
							return wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, lenVl.Mock}
					v, n, err = NewValidMapSer[uint, uint](nil, nil, lenVl, nil, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If keyVl returns an error, ValidMap.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 2
					wantErr = errors.New("key Validator error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					keySer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					keyVl = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, keySer.Mock, keyVl.Mock}
					v, n, err = NewValidMapSer[uint, uint](keySer, nil, nil, keyVl, nil).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If valueVl returns an error, ValidMap.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 3
					wantErr = errors.New("value Validator error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					keySer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					valueSer = mock.NewSerializer[uint]().RegisterUnmarshal(
						func(r muss.Reader) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					valueVl = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 11 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 11, v)
							}
							return wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, keySer.Mock, valueSer.Mock, valueVl.Mock}
					v, n, err = NewValidMapSer[uint, uint](keySer, valueSer, nil, nil, valueVl).Unmarshal(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

	})

}

func LengthReader(length int) mock.Reader {
	r := mock.NewReader()
	buf := &bytes.Buffer{}
	varint.PositiveInt.Marshal(length, buf)
	for _, b := range buf.Bytes() {
		func(b byte) {
			r.RegisterReadByte(func() (byte, error) {
				return b, nil
			})
		}(b)
	}
	return r
}
