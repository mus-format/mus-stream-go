package unsafe

import (
	"bytes"
	"errors"
	"io"
	"math"
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	com_mock "github.com/mus-format/common-go/testdata/mock"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/mus-format/mus-stream-go/varint"
	"github.com/ymz-ncnk/mok"
)

func TestUnsafe(t *testing.T) {

	t.Run("setUpUintFuncs", func(t *testing.T) {

		t.Run("If the system int size is not 32 or 64, setUpUintFuncs should panic with ErrUnsupportedIntSize",
			func(t *testing.T) {
				wantErr := com.ErrUnsupportedIntSize
				defer func() {
					if r := recover(); r != nil {
						err := r.(error)
						if err != wantErr {
							t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
						}
					}
				}()
				setUpUintFuncs(16)
			})

		t.Run("If the system int size is equal to 32, setUpUintFuncs should initialize the uint functions with 32-bit versions",
			func(t *testing.T) {
				setUpUintFuncs(32)
				if !com_testdata.ComparePtrs(marshalUint, marshalInteger32[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !com_testdata.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if !com_testdata.ComparePtrs(sizeUint, raw.SizeUint) {
					t.Error("unexpected sizeUint func")
				}
				if !com_testdata.ComparePtrs(skipUint, raw.SkipUint) {
					t.Error("unexpected skipUint func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpUintFuncs(64)
				if !com_testdata.ComparePtrs(marshalUint, marshalInteger64[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !com_testdata.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if !com_testdata.ComparePtrs(sizeUint, raw.SizeUint) {
					t.Error("unexpected sizeUint func")
				}
				if !com_testdata.ComparePtrs(skipUint, raw.SkipUint) {
					t.Error("unexpected skipUint func")
				}
			})

	})

	t.Run("setUpIntFuncs", func(t *testing.T) {

		t.Run("If the system int size is not 32 or 64, setUpIntFuncs should panic with ErrUnsupportedIntSize",
			func(t *testing.T) {
				wantErr := com.ErrUnsupportedIntSize
				defer func() {
					if r := recover(); r != nil {
						err := r.(error)
						if err != wantErr {
							t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
						}
					}
				}()
				setUpIntFuncs(16)
			})

		t.Run("If the system int size is equal to 32, setUpIntFuncs should initialize the uint functions with 32-bit versions",
			func(t *testing.T) {
				setUpIntFuncs(32)
				if !com_testdata.ComparePtrs(marshalInt, marshalInteger32[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !com_testdata.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if !com_testdata.ComparePtrs(sizeInt, raw.SizeInt) {
					t.Error("unexpected sizeInt func")
				}
				if !com_testdata.ComparePtrs(skipInt, raw.SkipInt) {
					t.Error("unexpected skipInt func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpIntFuncs(64)
				if !com_testdata.ComparePtrs(marshalInt, marshalInteger64[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !com_testdata.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if !com_testdata.ComparePtrs(sizeInt, raw.SizeInt) {
					t.Error("unexpected sizeInt func")
				}
				if !com_testdata.ComparePtrs(skipInt, raw.SkipInt) {
					t.Error("unexpected skipInt func")
				}
			})

	})

	t.Run("If Reader fails to read a byte slice, unmarshalInteger64 should return an error",
		func(t *testing.T) {
			var (
				wantV   uint64 = 0
				wantN          = 0
				wantErr        = errors.New("read byte error")
				r              = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger64[uint64](r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger64 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantV   uint64 = 0
				wantN          = 4
				wantErr error  = io.ErrUnexpectedEOF
				r              = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return wantN, nil
					},
				).RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger64[uint64](r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If Reader fails to read a byte slice, unmarshalInteger32 should return an error",
		func(t *testing.T) {
			var (
				wantV   uint32 = 0
				wantN          = 0
				wantErr        = errors.New("read byte error")
				r              = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger32[uint32](r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger32 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantV   uint32 = 0
				wantN          = 2
				wantErr error  = io.ErrUnexpectedEOF
				r              = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return wantN, nil
					},
				).RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger32[uint32](r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If Reader fails to read a byte slice, unmarshalInteger16 should return an error",
		func(t *testing.T) {
			var (
				wantV   uint16 = 0
				wantN          = 0
				wantErr        = errors.New("read byte error")
				r              = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger16[uint16](r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger16 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantV   uint16 = 0
				wantN          = 1
				wantErr error  = io.ErrUnexpectedEOF
				r              = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return wantN, nil
					},
				).RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger16[uint16](r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If Reader fails to read a byte slice, unmarshalInteger8 should return an error",
		func(t *testing.T) {
			var (
				wantV   uint8 = 0
				wantN         = 0
				wantErr       = errors.New("read byte error")
				r             = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger8[uint8](r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger8 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantV   uint8 = 0
				wantN         = 0
				wantErr error = io.EOF
				r             = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger8[uint8](r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("string", func(t *testing.T) {

		t.Run("All MarshalString, UnmarshalString, SizeString, SkipString functions with default lenM, lenU, lenS must work correctly",
			func(t *testing.T) {
				var (
					m muss.MarshallerFn[string] = func(v string, w muss.Writer) (n int,
						err error) {
						return MarshalString(v, nil, w)
					}
					u muss.UnmarshallerFn[string] = func(r muss.Reader) (v string, n int, err error) {
						return UnmarshalString(nil, r)
					}
					s muss.SizerFn[string] = func(v string) (size int) {
						return SizeString(v, nil)
					}
					sk muss.SkipperFn = func(r muss.Reader) (n int, err error) {
						return SkipString(nil, r)
					}
				)
				testdata.Test[string](com_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("All MarshalStringVarint, UnmarshalStringVarint, SizeStringVarint, SkipStringVarint functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[string](MarshalStringVarint)
					u  = muss.UnmarshallerFn[string](UnmarshalStringVarint)
					s  = muss.SizerFn[string](SizeStringVarint)
					sk = muss.SkipperFn(SkipStringVarint)
				)
				testdata.Test[string](com_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("If Writer fails to write a string length, MarshalStringVarint should return an error",
			func(t *testing.T) {
				var (
					s       = "hello world"
					wantN   = 0
					wantErr = errors.New("write error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = MarshalStringVarint(s, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read a string length, UnmarshalStringVarint should return an error",
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
					v, n, err = UnmarshalStringVarint(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails with an io.EOF, UnmarshalStringVarint should return io.ErrUnexpectedEOF",
			func(t *testing.T) {
				var (
					wantV   string = ""
					wantN          = 5
					wantErr error  = io.ErrUnexpectedEOF
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 22, nil },
					).RegisterRead(
						func(p []byte) (n int, err error) {
							return copy(p, "long"), nil
						},
					).RegisterRead(
						func(p []byte) (n int, err error) { return 0, io.EOF },
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalStringVarint(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If encoded string has negative length, UnmarshalStringVarint should return ErrNegativeLength",
			func(t *testing.T) {
				var (
					wantV   string = ""
					wantN          = 1
					wantErr        = com.ErrNegativeLength
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalStringVarint(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("lenVl validator should protect against too much length",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 10
					wantErr = errors.New("lenVl validator error")
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = math.MaxInt64
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					r = func() mock.Reader {
						buf := &bytes.Buffer{}
						varint.MarshalInt64(math.MaxInt64, buf)
						return mock.NewReader().RegisterNReadByte(com.Uint64MaxVarintLen,
							func() (b byte, err error) {
								return buf.ReadByte()
							},
						)
					}()
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidStringVarint(lenVl, false, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If skip == false and lenVl validator fails with an error, UnmarshalValidStringVarint should immediately return it",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 1
					wantErr = errors.New("lenVl validator error")
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidStringVarint(lenVl, false, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If skip == true and lenVl validator fails with an error, UnmarshalValidStringVarint should return it and skip all bytes of the string",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 4
					wantErr = errors.New("lenVl validator error")
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterNReadByte(3,
						func() (b byte, err error) {
							return 0, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidStringVarint(lenVl, true, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If skip == true, lenVl validator != nil and Reader fails with an error, UnmarshalValidStringVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 1
					wantErr = errors.New("Reader error")
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return errors.New("lenVl validator error")
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidStringVarint(lenVl, true, r)
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
				v, n, err = UnmarshalValidStringVarint(lenVl, false, r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("If Reader fails to read string content, UnmarshalStringVarint should return an error",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 2
					wantErr = errors.New("read string content error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterRead(
						func(p []byte) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalStringVarint(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

	})

	t.Run("All MarshalByte, UnmarshalByte, SizeByte, SkipByte functions must work correctly",
		func(t *testing.T) {
			var (
				m  = muss.MarshallerFn[byte](MarshalByte)
				u  = muss.UnmarshallerFn[byte](UnmarshalByte)
				s  = muss.SizerFn[byte](SizeByte)
				sk = muss.SkipperFn(SkipByte)
			)
			testdata.Test[byte](com_testdata.ByteTestCases, m, u, s, t)
			testdata.TestSkip[byte](com_testdata.ByteTestCases, m, sk, s, t)
		})

	t.Run("Unsigned", func(t *testing.T) {

		t.Run("All MarshalUint64, UnmarshalUint64, SizeUint64, SkipUint64 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[uint64](MarshalUint64)
					u  = muss.UnmarshallerFn[uint64](UnmarshalUint64)
					s  = muss.SizerFn[uint64](SizeUint64)
					sk = muss.SkipperFn(SkipUint64)
				)
				testdata.Test[uint64](com_testdata.Uint64TestCases, m, u, s, t)
				testdata.TestSkip[uint64](com_testdata.Uint64TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint32, UnmarshalUint32, SizeUint32, SkipUint32 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[uint32](MarshalUint32)
					u  = muss.UnmarshallerFn[uint32](UnmarshalUint32)
					s  = muss.SizerFn[uint32](SizeUint32)
					sk = muss.SkipperFn(SkipUint32)
				)
				testdata.Test[uint32](com_testdata.Uint32TestCases, m, u, s, t)
				testdata.TestSkip[uint32](com_testdata.Uint32TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint16, UnmarshalUint16, SizeUint16, SkipUint16 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[uint16](MarshalUint16)
					u  = muss.UnmarshallerFn[uint16](UnmarshalUint16)
					s  = muss.SizerFn[uint16](SizeUint16)
					sk = muss.SkipperFn(SkipUint16)
				)
				testdata.Test[uint16](com_testdata.Uint16TestCases, m, u, s, t)
				testdata.TestSkip[uint16](com_testdata.Uint16TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint8, UnmarshalUint8, SizeUint8, SkipUint8 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[uint8](MarshalUint8)
					u  = muss.UnmarshallerFn[uint8](UnmarshalUint8)
					s  = muss.SizerFn[uint8](SizeUint8)
					sk = muss.SkipperFn(SkipUint8)
				)
				testdata.Test[uint8](com_testdata.Uint8TestCases, m, u, s, t)
				testdata.TestSkip[uint8](com_testdata.Uint8TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint, UnmarshalUint, SizeUint, SkipUint functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[uint](MarshalUint)
					u  = muss.UnmarshallerFn[uint](UnmarshalUint)
					s  = muss.SizerFn[uint](SizeUint)
					sk = muss.SkipperFn(SkipUint)
				)
				testdata.Test[uint](com_testdata.UintTestCases, m, u, s, t)
				testdata.TestSkip[uint](com_testdata.UintTestCases, m, sk, s, t)
			})

	})

	t.Run("Signed", func(t *testing.T) {

		t.Run("All MarshalInt64, UnmarshalInt64, SizeInt64, SkipInt64 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[int64](MarshalInt64)
					u  = muss.UnmarshallerFn[int64](UnmarshalInt64)
					s  = muss.SizerFn[int64](SizeInt64)
					sk = muss.SkipperFn(SkipInt64)
				)
				testdata.Test[int64](com_testdata.Int64TestCases, m, u, s, t)
				testdata.TestSkip[int64](com_testdata.Int64TestCases, m, sk, s, t)
			})

		t.Run("All MarshalInt32, UnmarshalInt32, SizeInt32, SkipInt32 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[int32](MarshalInt32)
					u  = muss.UnmarshallerFn[int32](UnmarshalInt32)
					s  = muss.SizerFn[int32](SizeInt32)
					sk = muss.SkipperFn(SkipInt32)
				)
				testdata.Test[int32](com_testdata.Int32TestCases, m, u, s, t)
				testdata.TestSkip[int32](com_testdata.Int32TestCases, m, sk, s, t)
			})

		t.Run("All MarshalInt16, UnmarshalInt16, SizeInt16, SkipInt16 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[int16](MarshalInt16)
					u  = muss.UnmarshallerFn[int16](UnmarshalInt16)
					s  = muss.SizerFn[int16](SizeInt16)
					sk = muss.SkipperFn(SkipInt16)
				)
				testdata.Test[int16](com_testdata.Int16TestCases, m, u, s, t)
				testdata.TestSkip[int16](com_testdata.Int16TestCases, m, sk, s, t)
			})

		t.Run("All MarshalInt8, UnmarshalInt8, SizeInt8, SkipInt8 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[int8](MarshalInt8)
					u  = muss.UnmarshallerFn[int8](UnmarshalInt8)
					s  = muss.SizerFn[int8](SizeInt8)
					sk = muss.SkipperFn(SkipInt8)
				)
				testdata.Test[int8](com_testdata.Int8TestCases, m, u, s, t)
				testdata.TestSkip[int8](com_testdata.Int8TestCases, m, sk, s, t)
			})

		t.Run("All MarshalInt, UnmarshalInt, SizeInt, SkipInt functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[int](MarshalInt)
					u  = muss.UnmarshallerFn[int](UnmarshalInt)
					s  = muss.SizerFn[int](SizeInt)
					sk = muss.SkipperFn(SkipInt)
				)
				testdata.Test[int](com_testdata.IntTestCases, m, u, s, t)
				testdata.TestSkip[int](com_testdata.IntTestCases, m, sk, s, t)
			})

	})

	t.Run("Float", func(t *testing.T) {

		t.Run("float64", func(t *testing.T) {

			t.Run("All MarshalFloat64, UnmarshalFloat64, SizeFloat64, SkipFloat64 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = muss.MarshallerFn[float64](MarshalFloat64)
						u  = muss.UnmarshallerFn[float64](UnmarshalFloat64)
						s  = muss.SizerFn[float64](SizeFloat64)
						sk = muss.SkipperFn(SkipFloat64)
					)
					testdata.Test[float64](com_testdata.Float64TestCases, m, u, s, t)
					testdata.TestSkip[float64](com_testdata.Float64TestCases, m, sk, s,
						t)
				})

			t.Run("If Reader fails to read a byte slice, UnmarshalFloat64 should return an error",
				func(t *testing.T) {
					var (
						wantV   float64 = 0.0
						wantN           = 0
						wantErr         = errors.New("read byte error")
						r               = mock.NewReader().RegisterRead(
							func(p []byte) (n int, err error) {
								return 0, wantErr
							},
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = UnmarshalFloat64(r)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})

		})

		t.Run("float32", func(t *testing.T) {

			t.Run("All MarshalFloat32, UnmarshalFloat32, SizeFloat32, SkipFloat32 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = muss.MarshallerFn[float32](MarshalFloat32)
						u  = muss.UnmarshallerFn[float32](UnmarshalFloat32)
						s  = muss.SizerFn[float32](SizeFloat32)
						sk = muss.SkipperFn(SkipFloat32)
					)
					testdata.Test[float32](com_testdata.Float32TestCases, m, u, s, t)
					testdata.TestSkip[float32](com_testdata.Float32TestCases, m, sk, s,
						t)
				})

			t.Run("If Reader fails to read a byte slice, UnmarshalFloat32 should return an error",
				func(t *testing.T) {
					var (
						wantV   float32 = 0.0
						wantN           = 0
						wantErr         = errors.New("read byte error")
						r               = mock.NewReader().RegisterRead(
							func(p []byte) (n int, err error) {
								return 0, wantErr
							},
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = UnmarshalFloat32(r)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})

		})

	})

	t.Run("bool", func(t *testing.T) {

		t.Run("All MarshalBool, UnmarshalBool, SizeBool, SkipBool functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[bool](MarshalBool)
					u  = muss.UnmarshallerFn[bool](UnmarshalBool)
					s  = muss.SizerFn[bool](SizeBool)
					sk = muss.SkipperFn(SkipBool)
				)
				testdata.Test[bool](com_testdata.BoolTestCases, m, u, s, t)
				testdata.TestSkip[bool](com_testdata.BoolTestCases, m, sk, s, t)
			})

		t.Run("If Writer fails to write a byte, MarshalBool should return an error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("write byte error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = MarshalBool(true, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read a byte, UnmarshalBool should return an error",
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
					v, n, err = UnmarshalBool(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalBool should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV   bool = false
					wantN        = 0
					wantErr      = com.ErrWrongFormat
					r            = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalBool(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

	})

}

// StringVarint

func MarshalStringVarint(v string, w muss.Writer) (n int, err error) {
	return MarshalString(v, muss.MarshallerFn[int](varint.MarshalInt), w)
}

func UnmarshalStringVarint(r muss.Reader) (v string,
	n int, err error) {
	return UnmarshalValidStringVarint(nil, false, r)
}

func UnmarshalValidStringVarint(lenVl com.Validator[int], skip bool, r muss.Reader) (
	v string, n int, err error) {
	return UnmarshalValidString(muss.UnmarshallerFn[int](varint.UnmarshalInt),
		lenVl, skip, r)
}

func SizeStringVarint(v string) (n int) {
	return SizeString(v, muss.SizerFn[int](varint.SizeInt))
}

func SkipStringVarint(r muss.Reader) (n int, err error) {
	return SkipString(muss.UnmarshallerFn[int](varint.UnmarshalInt), r)
}
