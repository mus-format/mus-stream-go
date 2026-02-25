package raw

import (
	"errors"
	"os"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testutil"
	"github.com/mus-format/mus-stream-go/testutil/mock"
	"github.com/ymz-ncnk/mok"
	"golang.org/x/exp/constraints"
)

func TestRaw(t *testing.T) {
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
				if !ctestutil.ComparePtrs(sizeUint, sizeNum32[uint]) {
					t.Error("unexpected sizeUint func")
				}
				if !ctestutil.ComparePtrs(skipUint, skipInteger32) {
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
				if !ctestutil.ComparePtrs(sizeUint, sizeNum64[uint]) {
					t.Error("unexpected sizeUint func")
				}
				if !ctestutil.ComparePtrs(skipUint, skipInteger64) {
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
				if !ctestutil.ComparePtrs(sizeInt, sizeNum32[int]) {
					t.Error("unexpected sizeInt func")
				}
				if !ctestutil.ComparePtrs(skipInt, skipInteger32) {
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
				if !ctestutil.ComparePtrs(sizeInt, sizeNum64[int]) {
					t.Error("unexpected sizeInt func")
				}
				if !ctestutil.ComparePtrs(skipInt, skipInteger64) {
					t.Error("unexpected skipInt func")
				}
			})
	})

	t.Run("If Writer fails to write a byte, marshalInteger64 should return an error",
		func(t *testing.T) {
			testMarshalIntegerError(0, marshalInteger64[int64], t)
			testMarshalIntegerError(1, marshalInteger64[int64], t)
			testMarshalIntegerError(2, marshalInteger64[int64], t)
			testMarshalIntegerError(3, marshalInteger64[int64], t)
			testMarshalIntegerError(4, marshalInteger64[int64], t)
			testMarshalIntegerError(5, marshalInteger64[int64], t)
			testMarshalIntegerError(6, marshalInteger64[int64], t)
			testMarshalIntegerError(7, marshalInteger64[int64], t)
		})

	t.Run("If Reader fails to read a byte, unmarshalInteger64 should return an error",
		func(t *testing.T) {
			testUnmarshalIntegerError(0, unmarshalInteger64[int64], t)
			testUnmarshalIntegerError(1, unmarshalInteger64[int64], t)
			testUnmarshalIntegerError(2, unmarshalInteger64[int64], t)
			testUnmarshalIntegerError(3, unmarshalInteger64[int64], t)
			testUnmarshalIntegerError(4, unmarshalInteger64[int64], t)
			testUnmarshalIntegerError(5, unmarshalInteger64[int64], t)
			testUnmarshalIntegerError(6, unmarshalInteger64[int64], t)
			testUnmarshalIntegerError(7, unmarshalInteger64[int64], t)
		})

	t.Run("If Reader fails to read a byte, skipInteger64 should return an error",
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
				n, err = skipInteger64(r)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If Writer fails to write a byte, marshalInteger32 should return an error",
		func(t *testing.T) {
			testMarshalIntegerError(0, marshalInteger32[int32], t)
			testMarshalIntegerError(1, marshalInteger32[int32], t)
			testMarshalIntegerError(2, marshalInteger32[int32], t)
			testMarshalIntegerError(3, marshalInteger32[int32], t)
		})

	t.Run("If Reader fails to read a byte, unmarshalInteger32 should return an error",
		func(t *testing.T) {
			testUnmarshalIntegerError(0, unmarshalInteger32[int32], t)
			testUnmarshalIntegerError(1, unmarshalInteger32[int32], t)
			testUnmarshalIntegerError(2, unmarshalInteger32[int32], t)
			testUnmarshalIntegerError(3, unmarshalInteger32[int32], t)
		})

	t.Run("If Reader fails to read a byte, skipInteger32 should return an error",
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
				n, err = skipInteger32(r)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If Writer fails to write a byte, marshalInteger16 should return an error",
		func(t *testing.T) {
			testMarshalIntegerError(0, marshalInteger16[int16], t)
			testMarshalIntegerError(1, marshalInteger16[int16], t)
		})

	t.Run("If Reader fails to read a byte, unmarshalInteger16 should return an error",
		func(t *testing.T) {
			testUnmarshalIntegerError(0, unmarshalInteger16[int16], t)
			testUnmarshalIntegerError(1, unmarshalInteger16[int16], t)
		})

	t.Run("If Reader fails to read a byte, skipInteger16 should return an error",
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
				n, err = skipInteger16(r)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If Writer fails to write a byte, marshalInteger8 should return it",
		func(t *testing.T) {
			testMarshalIntegerError(0, marshalInteger8[int8], t)
		})

	t.Run("If Reader fails to read a byte, unmarshalInteger8 should return it",
		func(t *testing.T) {
			testUnmarshalIntegerError(0, unmarshalInteger8[int8], t)
		})

	t.Run("If Reader fails to read a byte, skipInteger8 should return it",
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
				n, err = skipInteger8(r)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
		t.Run("Uint64 serializer should work correctly",
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
		t.Run("Uint16 serializer should work correctly", func(t *testing.T) {
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

			t.Run("If Reader fails to read a byte, UnmarshalFloat64 should return an error",
				func(t *testing.T) {
					var (
						wantV   float64 = 0
						wantN           = 0
						wantErr         = errors.New("read byte error")
						r               = mock.NewReader().RegisterReadByte(
							func() (b byte, err error) {
								return 0, wantErr
							},
						)
						v, n, err = Float64.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil,
						t)
				})

			t.Run("If Reader fails to read a byte, SkipFloat64 should return an error",
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
						n, err = Float64.Skip(r)
					)
					ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
				})
		})

		t.Run("float32", func(t *testing.T) {
			t.Run("Float32 serializer should work correctly",
				func(t *testing.T) {
					ser := Float32
					testutil.Test[float32](ctestutil.Float32TestCases, ser, t)
					testutil.TestSkip[float32](ctestutil.Float32TestCases, ser, t)
				})

			t.Run("If Reader fails to read a byte, UnmarshalFloat32 should return an error",
				func(t *testing.T) {
					var (
						wantV   float32 = 0
						wantN           = 0
						wantErr         = errors.New("read byte error")
						r               = mock.NewReader().RegisterReadByte(
							func() (b byte, err error) {
								return 0, wantErr
							},
						)
						v, n, err = Float32.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil,
						t)
				})

			t.Run("If Reader fails to read a byte, SkipFloat32 should return an error",
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
						n, err = Float32.Skip(r)
					)
					ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
				})
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
						r       = mock.NewReader().RegisterReadByte(
							func() (b byte, err error) { err = wantErr; return },
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = TimeUnixUTC.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks, t)
				})
		})

		t.Run("time_unix_milli", func(t *testing.T) {
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
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnixMilliUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnixMilliUTC, t)
				})

			t.Run("If Reader fails to read a byte, Unmarshal should return error",
				func(t *testing.T) {
					var (
						wantV   = time.Time{}
						wantN   = 0
						wantErr = errors.New("read byte error")
						r       = mock.NewReader().RegisterReadByte(
							func() (b byte, err error) { err = wantErr; return },
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
						r       = mock.NewReader().RegisterReadByte(
							func() (b byte, err error) { err = wantErr; return },
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
						r       = mock.NewReader().RegisterReadByte(
							func() (b byte, err error) { err = wantErr; return },
						)
						mocks     = []*mok.Mock{r.Mock}
						v, n, err = TimeUnixNanoUTC.Unmarshal(r)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks, t)
				})
		})
	})
}

func testMarshalIntegerError[T constraints.Integer](k int,
	fn func(t T, w mus.Writer) (n int, err error),
	t *testing.T,
) {
	var (
		num     T = 0
		wantN     = k
		wantErr   = errors.New("write byte error")
		w         = mock.NewWriter().RegisterNWriteByte(k,
			func(c byte) error { return nil },
		).RegisterWriteByte(
			func(c byte) error { return wantErr },
		)
		mocks  = []*mok.Mock{w.Mock}
		n, err = fn(num, w)
	)
	testutil.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
}

func testUnmarshalIntegerError[T constraints.Integer](k int,
	fn func(r mus.Reader) (t T, n int, err error),
	t *testing.T,
) {
	var (
		wantV   T = 0
		wantN     = k
		wantErr   = errors.New("read byte error")
		r         = mock.NewReader().RegisterNReadByte(k,
			func() (b byte, err error) { return 0, nil },
		).RegisterReadByte(
			func() (b byte, err error) { return 0, wantErr },
		)
		mocks     = []*mok.Mock{r.Mock}
		v, n, err = fn(r)
	)
	ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
		t)
}
