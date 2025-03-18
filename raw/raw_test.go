package raw

import (
	"errors"
	"os"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
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
				if !com_testdata.ComparePtrs(marshalUint, marshalInteger32[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !com_testdata.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if !com_testdata.ComparePtrs(sizeUint, sizeNum32[uint]) {
					t.Error("unexpected sizeUint func")
				}
				if !com_testdata.ComparePtrs(skipUint, skipInteger32) {
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
				if !com_testdata.ComparePtrs(sizeUint, sizeNum64[uint]) {
					t.Error("unexpected sizeUint func")
				}
				if !com_testdata.ComparePtrs(skipUint, skipInteger64) {
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
				if !com_testdata.ComparePtrs(sizeInt, sizeNum32[int]) {
					t.Error("unexpected sizeInt func")
				}
				if !com_testdata.ComparePtrs(skipInt, skipInteger32) {
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
				if !com_testdata.ComparePtrs(sizeInt, sizeNum64[int]) {
					t.Error("unexpected sizeInt func")
				}
				if !com_testdata.ComparePtrs(skipInt, skipInteger64) {
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
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
			com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("byte", func(t *testing.T) {

		t.Run("Byte serializer should work correctly",
			func(t *testing.T) {
				ser := Byte
				testdata.Test[byte](com_testdata.ByteTestCases, ser, t)
				testdata.TestSkip[byte](com_testdata.ByteTestCases, ser, t)
			})

	})

	t.Run("unsigned", func(t *testing.T) {

		t.Run("Uint64 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint64
				testdata.Test[uint64](com_testdata.Uint64TestCases, ser, t)
				testdata.TestSkip[uint64](com_testdata.Uint64TestCases, ser, t)
			})

		t.Run("Uint32 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint32
				testdata.Test[uint32](com_testdata.Uint32TestCases, ser, t)
				testdata.TestSkip[uint32](com_testdata.Uint32TestCases, ser, t)
			})
		t.Run("Uint16 serializer should work correctly", func(t *testing.T) {
			ser := Uint16
			testdata.Test[uint16](com_testdata.Uint16TestCases, ser, t)
			testdata.TestSkip[uint16](com_testdata.Uint16TestCases, ser, t)
		})

		t.Run("Uint8 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint8
				testdata.Test[uint8](com_testdata.Uint8TestCases, ser, t)
				testdata.TestSkip[uint8](com_testdata.Uint8TestCases, ser, t)
			})

		t.Run("Uint serializer should work correctly",
			func(t *testing.T) {
				ser := Uint
				testdata.Test[uint](com_testdata.UintTestCases, ser, t)
				testdata.TestSkip[uint](com_testdata.UintTestCases, ser, t)
			})

	})

	t.Run("signed", func(t *testing.T) {

		t.Run("Int64 serializer should work correctly",
			func(t *testing.T) {
				ser := Int64
				testdata.Test[int64](com_testdata.Int64TestCases, ser, t)
				testdata.TestSkip[int64](com_testdata.Int64TestCases, ser, t)
			})

		t.Run("Int32 serializer should work correctly",
			func(t *testing.T) {
				ser := Int32
				testdata.Test[int32](com_testdata.Int32TestCases, ser, t)
				testdata.TestSkip[int32](com_testdata.Int32TestCases, ser, t)
			})

		t.Run("Int16 serializer should work correctly",
			func(t *testing.T) {
				ser := Int16
				testdata.Test[int16](com_testdata.Int16TestCases, ser, t)
				testdata.TestSkip[int16](com_testdata.Int16TestCases, ser, t)
			})

		t.Run("Int8 serializer should work correctly",
			func(t *testing.T) {
				ser := Int8
				testdata.Test[int8](com_testdata.Int8TestCases, ser, t)
				testdata.TestSkip[int8](com_testdata.Int8TestCases, ser, t)
			})

		t.Run("Int serializer should work correctly",
			func(t *testing.T) {
				ser := Int
				testdata.Test[int](com_testdata.IntTestCases, ser, t)
				testdata.TestSkip[int](com_testdata.IntTestCases, ser, t)
			})

	})

	t.Run("float", func(t *testing.T) {

		t.Run("float64", func(t *testing.T) {

			t.Run("Float64 serializer should work correctly",
				func(t *testing.T) {
					ser := Float64
					testdata.Test[float64](com_testdata.Float64TestCases, ser, t)
					testdata.TestSkip[float64](com_testdata.Float64TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
					com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
				})

		})

		t.Run("float32", func(t *testing.T) {

			t.Run("Float32 serializer should work correctly",
				func(t *testing.T) {
					ser := Float32
					testdata.Test[float32](com_testdata.Float32TestCases, ser, t)
					testdata.TestSkip[float32](com_testdata.Float32TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
					com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
					testdata.Test[time.Time]([]time.Time{tm}, TimeUnixUTC, t)
					testdata.TestSkip[time.Time]([]time.Time{tm}, TimeUnixUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testdata.Test[time.Time]([]time.Time{time.Time{}}, TimeUnixUTC, t)
					testdata.TestSkip[time.Time]([]time.Time{time.Time{}}, TimeUnixUTC, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
					testdata.Test[time.Time]([]time.Time{tm}, TimeUnixMilliUTC, t)
					testdata.TestSkip[time.Time]([]time.Time{tm}, TimeUnixMilliUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testdata.Test[time.Time]([]time.Time{time.Time{}}, TimeUnixMilliUTC, t)
					testdata.TestSkip[time.Time]([]time.Time{time.Time{}}, TimeUnixMilliUTC, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
					testdata.Test[time.Time]([]time.Time{tm}, TimeUnixMicroUTC, t)
					testdata.TestSkip[time.Time]([]time.Time{tm}, TimeUnixMicroUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testdata.Test[time.Time]([]time.Time{time.Time{}}, TimeUnix, t)
					testdata.TestSkip[time.Time]([]time.Time{time.Time{}}, TimeUnix, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
					testdata.Test[time.Time]([]time.Time{tm}, TimeUnixNanoUTC, t)
					testdata.TestSkip[time.Time]([]time.Time{tm}, TimeUnixNanoUTC, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks, t)
				})

		})

	})

}

func testMarshalIntegerError[T constraints.Integer](k int,
	fn func(t T, w muss.Writer) (n int, err error),
	t *testing.T) {
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
	testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
}

func testUnmarshalIntegerError[T constraints.Integer](k int,
	fn func(r muss.Reader) (t T, n int, err error),
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
	com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
		t)
}
