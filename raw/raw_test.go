package raw

import (
	"errors"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/test"
	"github.com/mus-format/mus-stream-go/test/mock"
	"github.com/ymz-ncnk/mok"
	"golang.org/x/exp/constraints"
)

func TestRaw_setUpUintFuncs(t *testing.T) {
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
			if !ctest.ComparePtrs(marshalUint, marshalInteger32[uint]) {
				t.Error("unexpected marshalUint func")
			}
			if !ctest.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
				t.Error("unexpected unmarshalUint func")
			}
			if !ctest.ComparePtrs(sizeUint, sizeNum32[uint]) {
				t.Error("unexpected sizeUint func")
			}
			if !ctest.ComparePtrs(skipUint, skipInteger32) {
				t.Error("unexpected skipUint func")
			}
		})

	t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
		func(t *testing.T) {
			setUpUintFuncs(64)
			if !ctest.ComparePtrs(marshalUint, marshalInteger64[uint]) {
				t.Error("unexpected marshalUint func")
			}
			if !ctest.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
				t.Error("unexpected unmarshalUint func")
			}
			if !ctest.ComparePtrs(sizeUint, sizeNum64[uint]) {
				t.Error("unexpected sizeUint func")
			}
			if !ctest.ComparePtrs(skipUint, skipInteger64) {
				t.Error("unexpected skipUint func")
			}
		})
}

func TestRaw_setUpIntFuncs(t *testing.T) {
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
			if !ctest.ComparePtrs(marshalInt, marshalInteger32[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !ctest.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if !ctest.ComparePtrs(sizeInt, sizeNum32[int]) {
				t.Error("unexpected sizeInt func")
			}
			if !ctest.ComparePtrs(skipInt, skipInteger32) {
				t.Error("unexpected skipInt func")
			}
		})

	t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
		func(t *testing.T) {
			setUpIntFuncs(64)
			if !ctest.ComparePtrs(marshalInt, marshalInteger64[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !ctest.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if !ctest.ComparePtrs(sizeInt, sizeNum64[int]) {
				t.Error("unexpected sizeInt func")
			}
			if !ctest.ComparePtrs(skipInt, skipInteger64) {
				t.Error("unexpected skipInt func")
			}
		})
}

func TestRaw_IntegerErrorHandling(t *testing.T) {
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
			ctest.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
			ctest.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
			ctest.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
			ctest.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})
}

func TestRaw_Byte(t *testing.T) {
	t.Run("Byte serializer should work correctly",
		func(t *testing.T) {
			ser := Byte
			test.Test(ctest.ByteTestCases, ser, t)
			test.TestSkip(ctest.ByteTestCases, ser, t)
		})
}

func TestRaw_Uint64(t *testing.T) {
	t.Run("Uint64 serializer should work correctly",
		func(t *testing.T) {
			ser := Uint64
			test.Test(ctest.Uint64TestCases, ser, t)
			test.TestSkip(ctest.Uint64TestCases, ser, t)
		})
}

func TestRaw_Uint32(t *testing.T) {

	t.Run("Uint32 serializer should work correctly",
		func(t *testing.T) {
			ser := Uint32
			test.Test(ctest.Uint32TestCases, ser, t)
			test.TestSkip(ctest.Uint32TestCases, ser, t)
		})
}

func TestRaw_Uint16(t *testing.T) {
	t.Run("Uint16 serializer should work correctly", func(t *testing.T) {
		ser := Uint16
		test.Test(ctest.Uint16TestCases, ser, t)
		test.TestSkip(ctest.Uint16TestCases, ser, t)
	})
}

func TestRaw_Uint8(t *testing.T) {

	t.Run("Uint8 serializer should work correctly",
		func(t *testing.T) {
			ser := Uint8
			test.Test(ctest.Uint8TestCases, ser, t)
			test.TestSkip(ctest.Uint8TestCases, ser, t)
		})
}

func TestRaw_Uint(t *testing.T) {

	t.Run("Uint serializer should work correctly",
		func(t *testing.T) {
			ser := Uint
			test.Test(ctest.UintTestCases, ser, t)
			test.TestSkip(ctest.UintTestCases, ser, t)
		})
}

func TestRaw_Int64(t *testing.T) {
	t.Run("Int64 serializer should work correctly",
		func(t *testing.T) {
			ser := Int64
			test.Test(ctest.Int64TestCases, ser, t)
			test.TestSkip(ctest.Int64TestCases, ser, t)
		})
}

func TestRaw_Int32(t *testing.T) {

	t.Run("Int32 serializer should work correctly",
		func(t *testing.T) {
			ser := Int32
			test.Test(ctest.Int32TestCases, ser, t)
			test.TestSkip(ctest.Int32TestCases, ser, t)
		})
}

func TestRaw_Int16(t *testing.T) {

	t.Run("Int16 serializer should work correctly",
		func(t *testing.T) {
			ser := Int16
			test.Test(ctest.Int16TestCases, ser, t)
			test.TestSkip(ctest.Int16TestCases, ser, t)
		})
}

func TestRaw_Int8(t *testing.T) {

	t.Run("Int8 serializer should work correctly",
		func(t *testing.T) {
			ser := Int8
			test.Test(ctest.Int8TestCases, ser, t)
			test.TestSkip(ctest.Int8TestCases, ser, t)
		})
}

func TestRaw_Int(t *testing.T) {

	t.Run("Int serializer should work correctly",
		func(t *testing.T) {
			ser := Int
			test.Test(ctest.IntTestCases, ser, t)
			test.TestSkip(ctest.IntTestCases, ser, t)
		})
}

func TestRaw_Float64(t *testing.T) {
	t.Run("Float64 serializer should work correctly",
		func(t *testing.T) {
			ser := Float64
			test.Test(ctest.Float64TestCases, ser, t)
			test.TestSkip(ctest.Float64TestCases, ser, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			ctest.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})
}

func TestRaw_Float32(t *testing.T) {
	t.Run("Float32 serializer should work correctly",
		func(t *testing.T) {
			ser := Float32
			test.Test(ctest.Float32TestCases, ser, t)
			test.TestSkip(ctest.Float32TestCases, ser, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			ctest.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})
}

func TestRaw_TimeUnixUTC(t *testing.T) {
	t.Run("TimeUnixUTC serializer should work correctly",
		func(t *testing.T) {
			var (
				sec = time.Now().Unix()
				tm  = time.Unix(sec, 0)
			)
			test.Test([]time.Time{tm}, TimeUnixUTC, t)
			test.TestSkip([]time.Time{tm}, TimeUnixUTC, t)
		})

	t.Run("We should be able to serializer the zero Time",
		func(t *testing.T) {
			test.Test([]time.Time{{}}, TimeUnixUTC, t)
			test.TestSkip([]time.Time{{}}, TimeUnixUTC, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})
}

func TestRaw_TimeUnixMilliUTC(t *testing.T) {
	t.Run("TimeUnixMilliUTC serializer should work correctly",
		func(t *testing.T) {
			var (
				milli = time.Now().UnixMilli()
				tm    = time.UnixMilli(milli)
			)
			test.Test([]time.Time{tm}, TimeUnixMilliUTC, t)
			test.TestSkip([]time.Time{tm}, TimeUnixMilliUTC, t)
		})

	t.Run("We should be able to serializer the zero Time",
		func(t *testing.T) {
			test.Test([]time.Time{{}}, TimeUnixMilliUTC, t)
			test.TestSkip([]time.Time{{}}, TimeUnixMilliUTC, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})
}

func TestRaw_TimeUnixMicroUTC(t *testing.T) {
	t.Run("TimeUnixMicroUTC serializer should work correctly",
		func(t *testing.T) {
			var (
				milli = time.Now().UnixMicro()
				tm    = time.UnixMicro(milli)
			)
			test.Test([]time.Time{tm}, TimeUnixMicroUTC, t)
			test.TestSkip([]time.Time{tm}, TimeUnixMicroUTC, t)
		})

	t.Run("We should be able to serializer the zero Time",
		func(t *testing.T) {
			test.Test([]time.Time{{}}, TimeUnix, t)
			test.TestSkip([]time.Time{{}}, TimeUnix, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})
}

func TestRaw_TimeUnixNanoUTC(t *testing.T) {
	t.Run("TimeUnixNanoUTC serializer should work correctly",
		func(t *testing.T) {
			var (
				nano = time.Now().UnixNano()
				tm   = time.Unix(0, nano)
			)
			test.Test([]time.Time{tm}, TimeUnixNanoUTC, t)
			test.TestSkip([]time.Time{tm}, TimeUnixNanoUTC, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})
}

func testMarshalIntegerError[T constraints.Integer](k int,
	fn func(t T, w mus.Writer) (n int, err error),
	t *testing.T,
) {
	var (
		num     T = 0
		wantErr   = errors.New("write byte error")
		w         = mock.NewWriter().RegisterNWriteByte(k,
			func(c byte) error { return nil },
		).RegisterWriteByte(
			func(c byte) error { return wantErr },
		)
		want = test.MarshalResults{
			N:     k,
			Err:   wantErr,
			Mocks: []*mok.Mock{w.Mock},
		}
	)
	test.TestMarshalOnly(num, w, test.MarshallerFn[T](fn), want, t)
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
	ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
}
