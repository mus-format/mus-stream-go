package unsafe

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	com_mock "github.com/mus-format/common-go/testutil/mock"
	muss "github.com/mus-format/mus-stream-go"
	arrops "github.com/mus-format/mus-stream-go/options/array"
	strops "github.com/mus-format/mus-stream-go/options/string"
	"github.com/mus-format/mus-stream-go/raw"
	"github.com/mus-format/mus-stream-go/testutil"
	"github.com/mus-format/mus-stream-go/testutil/mock"
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
				if !ctestutil.ComparePtrs(marshalUint, marshalInteger32[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !ctestutil.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if !ctestutil.ComparePtrs(sizeUint, raw.Uint.Size) {
					t.Error("unexpected sizeUint func")
				}
				if !ctestutil.ComparePtrs(skipUint, raw.Uint.Skip) {
					t.Error("unexpected skipUint func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpUintFuncs(64)
				if !ctestutil.ComparePtrs(marshalUint, marshalInteger64[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !ctestutil.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if !ctestutil.ComparePtrs(sizeUint, raw.Uint.Size) {
					t.Error("unexpected sizeUint func")
				}
				if !ctestutil.ComparePtrs(skipUint, raw.Uint.Skip) {
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
				if !ctestutil.ComparePtrs(marshalInt, marshalInteger32[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !ctestutil.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if !ctestutil.ComparePtrs(sizeInt, raw.Int.Size) {
					t.Error("unexpected sizeInt func")
				}
				if !ctestutil.ComparePtrs(skipInt, raw.Int.Skip) {
					t.Error("unexpected skipInt func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpIntFuncs(64)
				if !ctestutil.ComparePtrs(marshalInt, marshalInteger64[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !ctestutil.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if !ctestutil.ComparePtrs(sizeInt, raw.Int.Size) {
					t.Error("unexpected sizeInt func")
				}
				if !ctestutil.ComparePtrs(skipInt, raw.Int.Skip) {
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger64 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantV   uint64 = 0
				wantN          = 4
				wantErr        = io.ErrUnexpectedEOF
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger32 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantV   uint32 = 0
				wantN          = 2
				wantErr        = io.ErrUnexpectedEOF
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger16 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantV   uint16 = 0
				wantN          = 1
				wantErr        = io.ErrUnexpectedEOF
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger8 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantV   uint8 = 0
				wantN         = 0
				wantErr       = io.EOF
				r             = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalInteger8[uint8](r)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("string", func(t *testing.T) {
		t.Run("String serializer should work correctly",
			func(t *testing.T) {
				ser := String
				testutil.Test[string](ctestutil.StringTestCases, ser, t)
				testutil.TestSkip[string](ctestutil.StringTestCases, ser, t)
			})

		t.Run("We should be able to set a length serializer",
			func(t *testing.T) {
				var (
					str, lenSer = testutil.StringSerData(t)
					ser         = NewStringSer(strops.WithLenSer(lenSer))
					mocks       = []*mok.Mock{lenSer.Mock}
				)
				testutil.Test[string]([]string{str}, ser, t)
				testutil.TestSkip[string]([]string{str}, ser, t)

				if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
					t.Error(infomap)
				}
			})

		t.Run("If Writer fails to write a string length, Marshal should return an error",
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
					n, err = String.Marshal(s, w)
				)
				testutil.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails with an io.EOF, Unmarshal should return io.ErrUnexpectedEOF",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 5
					wantErr = io.ErrUnexpectedEOF
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 22, nil },
					).RegisterRead(
						func(p []byte) (n int, err error) {
							return copy(p, "long"), nil
						},
					).RegisterRead(
						func(p []byte) (n int, err error) { return 0, io.EOF },
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = String.Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If Reader fails to read a string length, Unmarshal should return an error",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = String.Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
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
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If Reader fails to read string content, Unmarshal should return an error",
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
					v, n, err = String.Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("ValidString should work correctly",
			func(t *testing.T) {
				ser := NewValidStringSer(nil)
				testutil.Test[string](ctestutil.StringTestCases, ser, t)
				testutil.TestSkip[string](ctestutil.StringTestCases, ser, t)
			})

		t.Run("If lenVl validator returns an error, ValidString.Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV      = ""
					wantLength = math.MaxInt64 - 1
					wantN      = 9
					wantErr    = errors.New("lenVl error")
					lenVl      = com_mock.NewValidator[int]().RegisterValidate(
						func(length int) (err error) {
							if length != wantLength {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantLength, length)
							}
							return wantErr
						},
					)
					r         = LengthReader(wantLength)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidStringSer(strops.WithLenValidator(lenVl)).Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
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
				v, n, err = NewValidStringSer(strops.WithLenValidator(lenVl)).Unmarshal(r)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("If Reader fails to read a string length, ValidString.Unmarshal should return an error",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidStringSer(nil).Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("ValidString.Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 10
					wantErr   = com.ErrNegativeLength
					r         = LengthReader(-1)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = NewValidStringSer(nil).Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("If Reader fails to read string content, ValidString.Unmarshal should return an error",
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
					v, n, err = NewValidStringSer(nil).Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})
	})

	t.Run("byte", func(t *testing.T) {
		t.Run("Byte serializer should work correctly",
			func(t *testing.T) {
				ser := Byte
				testutil.Test[byte](ctestutil.ByteTestCases, ser, t)
				testutil.TestSkip[byte](ctestutil.ByteTestCases, ser, t)
			})
	})

	t.Run("unsigned", func(t *testing.T) {
		t.Run("Uint64 should work correctly",
			func(t *testing.T) {
				ser := Uint64
				testutil.Test[uint64](ctestutil.Uint64TestCases, ser, t)
				testutil.TestSkip[uint64](ctestutil.Uint64TestCases, ser, t)
			})

		t.Run("Uint32 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint32
				testutil.Test[uint32](ctestutil.Uint32TestCases, ser, t)
				testutil.TestSkip[uint32](ctestutil.Uint32TestCases, ser, t)
			})

		t.Run("Uint16 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint16
				testutil.Test[uint16](ctestutil.Uint16TestCases, ser, t)
				testutil.TestSkip[uint16](ctestutil.Uint16TestCases, ser, t)
			})

		t.Run("Uint8 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint8
				testutil.Test[uint8](ctestutil.Uint8TestCases, ser, t)
				testutil.TestSkip[uint8](ctestutil.Uint8TestCases, ser, t)
			})

		t.Run("Uint serializer should work correctly",
			func(t *testing.T) {
				ser := Uint
				testutil.Test[uint](ctestutil.UintTestCases, ser, t)
				testutil.TestSkip[uint](ctestutil.UintTestCases, ser, t)
			})
	})

	t.Run("signed", func(t *testing.T) {
		t.Run("Int64 serializer should work correctly",
			func(t *testing.T) {
				ser := Int64
				testutil.Test[int64](ctestutil.Int64TestCases, ser, t)
				testutil.TestSkip[int64](ctestutil.Int64TestCases, ser, t)
			})

		t.Run("Int32 serializer should work correctly",
			func(t *testing.T) {
				ser := Int32
				testutil.Test[int32](ctestutil.Int32TestCases, ser, t)
				testutil.TestSkip[int32](ctestutil.Int32TestCases, ser, t)
			})

		t.Run("Int16 serializer should work correctly",
			func(t *testing.T) {
				ser := Int16
				testutil.Test[int16](ctestutil.Int16TestCases, ser, t)
				testutil.TestSkip[int16](ctestutil.Int16TestCases, ser, t)
			})

		t.Run("Int8 serializer should work correctly",
			func(t *testing.T) {
				ser := Int8
				testutil.Test[int8](ctestutil.Int8TestCases, ser, t)
				testutil.TestSkip[int8](ctestutil.Int8TestCases, ser, t)
			})

		t.Run("Int serializer should work correctly",
			func(t *testing.T) {
				ser := Int
				testutil.Test[int](ctestutil.IntTestCases, ser, t)
				testutil.TestSkip[int](ctestutil.IntTestCases, ser, t)
			})
	})

	t.Run("float", func(t *testing.T) {
		t.Run("float64", func(t *testing.T) {
			t.Run("Float64 serializer should work correctly",
				func(t *testing.T) {
					ser := Float64
					testutil.Test[float64](ctestutil.Float64TestCases, ser, t)
					testutil.TestSkip[float64](ctestutil.Float64TestCases, ser, t)
				})

			t.Run("If Reader fails to read a byte slice, Unmarshal should return an error",
				func(t *testing.T) {
					var (
						wantV   = 0.0
						wantN   = 0
						wantErr = errors.New("read byte error")
						r       = mock.NewReader().RegisterRead(
							func(p []byte) (n int, err error) {
								return 0, wantErr
							},
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = Float64.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("float32", func(t *testing.T) {
			t.Run("Float32 serializer should work correctly",
				func(t *testing.T) {
					ser := Float32
					testutil.Test[float32](ctestutil.Float32TestCases, ser, t)
					testutil.TestSkip[float32](ctestutil.Float32TestCases, ser, t)
				})

			t.Run("If Reader fails to read a byte slice, Unmarshal should return an error",
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
						v, n, err = Float32.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})
	})

	t.Run("bool", func(t *testing.T) {
		t.Run("Bool serializer should work correctly",
			func(t *testing.T) {
				ser := Bool
				testutil.Test[bool](ctestutil.BoolTestCases, ser, t)
				testutil.TestSkip[bool](ctestutil.BoolTestCases, ser, t)
			})

		t.Run("If Writer fails to write a byte, Marshal should return an error",
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
					n, err = Bool.Marshal(true, w)
				)
				testutil.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read a byte, Unmarshal should return an error",
			func(t *testing.T) {
				var (
					wantV   = false
					wantN   = 0
					wantErr = errors.New("read byte error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = Bool.Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV   = false
					wantN   = 1
					wantErr = com.ErrWrongFormat
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = Bool.Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})
	})

	t.Run("time", func(t *testing.T) {
		os.Setenv("TZ", "")

		t.Run("time_unix_utc", func(t *testing.T) {
			t.Run("TimeUnixUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						sec = time.Now().Unix()
						tm  = time.Unix(sec, 0)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnixUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnixUTC, t)
				})

			t.Run("If Reader fails to read a byte, Unmarshal should return error",
				func(t *testing.T) {
					var (
						wantV   = time.Time{}
						wantN   = 0
						wantErr = errors.New("read byte error")
						r       = mock.NewReader().RegisterRead(
							func(p []byte) (n int, err error) { err = wantErr; return },
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = TimeUnixUTC.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks, t)
				})
		})

		t.Run("time_unix_milli_utc", func(t *testing.T) {
			t.Run("TimeUnixMilliUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						milli = time.Now().UnixMilli()
						tm    = time.UnixMilli(milli)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixMilliUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixMilliUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnix, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnix, t)
				})

			t.Run("If Reader fails to read a byte, Unmarshal should return error",
				func(t *testing.T) {
					var (
						wantV   = time.Time{}
						wantN   = 0
						wantErr = errors.New("read byte error")
						r       = mock.NewReader().RegisterRead(
							func(p []byte) (n int, err error) { err = wantErr; return },
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = TimeUnixMilliUTC.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks, t)
				})
		})

		t.Run("time_unix_micro_utc", func(t *testing.T) {
			t.Run("TimeUnixMicroUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						milli = time.Now().UnixMicro()
						tm    = time.UnixMicro(milli)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixMicroUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixMicroUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnix, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnix, t)
				})

			t.Run("If Reader fails to read a byte, Unmarshal should return error",
				func(t *testing.T) {
					var (
						wantV   = time.Time{}
						wantN   = 0
						wantErr = errors.New("read byte error")
						r       = mock.NewReader().RegisterRead(
							func(p []byte) (n int, err error) { err = wantErr; return },
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = TimeUnixMicroUTC.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks, t)
				})
		})

		t.Run("time_unix_nano_utc", func(t *testing.T) {
			t.Run("TimeUnixNanoUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						nano = time.Now().UnixNano()
						tm   = time.Unix(0, nano)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixNanoUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixNanoUTC, t)
				})

			t.Run("If Reader fails to read a byte, Unmarshal should return error",
				func(t *testing.T) {
					var (
						wantV   = time.Time{}
						wantN   = 0
						wantErr = errors.New("read byte error")
						r       = mock.NewReader().RegisterRead(
							func(p []byte) (n int, err error) { err = wantErr; return },
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = TimeUnixNanoUTC.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks, t)
				})
		})
	})

	t.Run("array", func(t *testing.T) {
		t.Run("Array serializer should work correctly", func(t *testing.T) {
			var (
				arr, elemSer = testutil.ArraySerData(t)
				mocks        = []*mok.Mock{elemSer.Mock}
				ser          = NewArraySer[[3]int, int](elemSer)
			)
			testutil.Test[[3]int]([][3]int{arr}, ser, t)
			testutil.TestSkip[[3]int]([][3]int{arr}, ser, t)

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
					v, n, err = NewArraySer[[3]int, int](nil).Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Valid array serializer should work correctly", func(t *testing.T) {
			var (
				arr, elemSer = testutil.ArraySerData(t)
				mocks        = []*mok.Mock{elemSer.Mock}
				ser          = NewValidArraySer[[3]int, int](elemSer, nil)
			)
			testutil.Test[[3]int]([][3]int{arr}, ser, t)
			testutil.TestSkip[[3]int]([][3]int{arr}, ser, t)

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
					v, n, err = NewValidArraySer[[3]int, int](nil, nil).Unmarshal(r)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
					ser   = NewValidArraySer[[3]int, int](elemSer, arrops.WithElemValidator[int](elemVl))
					mocks = []*mok.Mock{r.Mock, elemSer.Mock, elemVl.Mock}
				)
				v, n, err := ser.Unmarshal(r)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("We should be able to set a length serializer for NewArraySer",
			func(t *testing.T) {
				var (
					arr = [3]int{1, 2, 3}
					ser = NewArraySer[[3]int, int](varint.Int, arrops.WithLenSer[int](varint.Int))
				)
				testutil.Test[[3]int]([][3]int{arr}, ser, t)
				testutil.TestSkip[[3]int]([][3]int{arr}, ser, t)
			})

		t.Run("We should be able to set a length serializer for NewValidArraySer",
			func(t *testing.T) {
				var (
					arr = [3]int{1, 2, 3}
					ser = NewValidArraySer[[3]int, int](varint.Int, arrops.WithLenSer[int](varint.Int))
				)
				testutil.Test[[3]int]([][3]int{arr}, ser, t)
				testutil.TestSkip[[3]int]([][3]int{arr}, ser, t)
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
