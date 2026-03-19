package varint

import (
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-stream-go/test"
	"github.com/mus-format/mus-stream-go/test/mock"
	"github.com/ymz-ncnk/mok"
)

func TestVarint_marshalUint(t *testing.T) {
	t.Run("If Writer fails to write a byte, marshalUint should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("write byte error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				want = test.MarshalResults{
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{w.Mock},
				}
			)
			test.TestMarshalOnly(uint(300), w, test.MarshallerFn[uint](marshalUint[uint]), want, t)
		})

	t.Run("If Writer fails to write a last one byte, marshalUint should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("write last byte error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error {
						return nil
					},
				).RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				want = test.MarshalResults{
					N:     1,
					Err:   wantErr,
					Mocks: []*mok.Mock{w.Mock},
				}
			)
			test.TestMarshalOnly(uint(300), w, test.MarshallerFn[uint](marshalUint[uint]), want, t)
		})
}

func TestVarint_unmarshalUint(t *testing.T) {
	t.Run("If Reader fails to read a byte, unmarshalUint should return an error",
		func(t *testing.T) {
			var (
				wantV   uint64 = 0
				wantN          = 0
				wantErr        = errors.New("read byte error")
				r              = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalUint[uint64](com.Uint64MaxVarintLen,
					com.Uint64MaxLastByte, r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	t.Run("unmarshalUint should return ErrOverflow if there is no varint end",
		func(t *testing.T) {
			var (
				wantV   uint16 = 0
				wantN          = 3
				wantErr        = com.ErrOverflow
				r              = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 4, nil },
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalUint[uint16](com.Uint16MaxVarintLen,
					com.Uint16MaxLastByte, r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_skipUint(t *testing.T) {
	t.Run("If Reader fails to read a byte, skipUint should return an error",
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
				n, err = skipUint(com.Uint64MaxVarintLen, com.Uint64MaxLastByte,
					r)
			)
			ctest.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("skipUint should return ErrOverflow if there is no varint end",
		func(t *testing.T) {
			var (
				wantN   = 3
				wantErr = com.ErrOverflow
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				)
				mocks  = []*mok.Mock{r.Mock}
				n, err = skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
					r)
			)
			ctest.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})
}

func TestVarint_Uint64(t *testing.T) {
	ser := Uint64
	test.Test(ctest.Uint64TestCases, ser, t)
	test.TestSkip(ctest.Uint64TestCases, ser, t)
}

func TestVarint_Uint32(t *testing.T) {
	ser := Uint32
	test.Test(ctest.Uint32TestCases, ser, t)
	test.TestSkip(ctest.Uint32TestCases, ser, t)
}

func TestVarint_Uint16(t *testing.T) {
	ser := Uint16
	test.Test(ctest.Uint16TestCases, ser, t)
	test.TestSkip(ctest.Uint16TestCases, ser, t)
}

func TestVarint_Uint8(t *testing.T) {
	ser := Uint8
	test.Test(ctest.Uint8TestCases, ser, t)
	test.TestSkip(ctest.Uint8TestCases, ser, t)
}

func TestVarint_Uint(t *testing.T) {
	ser := Uint
	test.Test(ctest.UintTestCases, ser, t)
	test.TestSkip(ctest.UintTestCases, ser, t)
}

func TestVarint_Int64(t *testing.T) {
	t.Run("Int64 serializer should work correctly",
		func(t *testing.T) {
			ser := Int64
			test.Test(ctest.Int64TestCases, ser, t)
			test.TestSkip(ctest.Int64TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   int64 = 0
				wantN         = 0
				wantErr       = errors.New("read byte error")
				r             = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = Int64.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_Int32(t *testing.T) {
	t.Run("Int32 serializer should work correctly",
		func(t *testing.T) {
			ser := Int32
			test.Test(ctest.Int32TestCases, ser, t)
			test.TestSkip(ctest.Int32TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   int32 = 0
				wantN         = 0
				wantErr       = errors.New("read byte error")
				r             = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = Int32.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_Int16(t *testing.T) {
	t.Run("Int16 serializer should work correctly",
		func(t *testing.T) {
			ser := Int16
			test.Test(ctest.Int16TestCases, ser, t)
			test.TestSkip(ctest.Int16TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   int16 = 0
				wantN         = 0
				wantErr       = errors.New("read byte error")
				r             = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = Int16.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_Int8(t *testing.T) {
	t.Run("Int8 serializer should work correctly",
		func(t *testing.T) {
			ser := Int8
			test.Test(ctest.Int8TestCases, ser, t)
			test.TestSkip(ctest.Int8TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   int8 = 0
				wantN        = 0
				wantErr      = errors.New("read byte error")
				r            = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = Int8.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_Int(t *testing.T) {
	t.Run("Int serializer should work correctly",
		func(t *testing.T) {
			ser := Int
			test.Test(ctest.IntTestCases, ser, t)
			test.TestSkip(ctest.IntTestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   = 0
				wantN   = 0
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = Int.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_PositiveInt64(t *testing.T) {
	t.Run("PositiveInt64 serializer should work correctly",
		func(t *testing.T) {
			ser := PositiveInt64
			test.Test(ctest.Int64TestCases, ser, t)
			test.TestSkip(ctest.Int64TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   int64 = 0
				wantN         = 0
				wantErr       = errors.New("read byte error")
				r             = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = PositiveInt64.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_PositiveInt32(t *testing.T) {
	t.Run("PositiveInt32 serializer should work correctly",
		func(t *testing.T) {
			ser := PositiveInt32
			test.Test(ctest.Int32TestCases, ser, t)
			test.TestSkip(ctest.Int32TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   int32 = 0
				wantN         = 0
				wantErr       = errors.New("read byte error")
				r             = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = PositiveInt32.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_PositiveInt16(t *testing.T) {
	t.Run("PositiveInt16 serializer should work correctly",
		func(t *testing.T) {
			ser := PositiveInt16
			test.Test(ctest.Int16TestCases, ser, t)
			test.TestSkip(ctest.Int16TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   int16 = 0
				wantN         = 0
				wantErr       = errors.New("read byte error")
				r             = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = PositiveInt16.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_PositiveInt8(t *testing.T) {
	t.Run("PositiveInt8 serializer should work correctly",
		func(t *testing.T) {
			ser := PositiveInt8
			test.Test(ctest.Int8TestCases, ser, t)
			test.TestSkip(ctest.Int8TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   int8 = 0
				wantN        = 0
				wantErr      = errors.New("read byte error")
				r            = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = PositiveInt8.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_PositiveInt(t *testing.T) {
	t.Run("PositiveInt serializer should work correctly",
		func(t *testing.T) {
			ser := PositiveInt
			test.Test(ctest.IntTestCases, ser, t)
			test.TestSkip(ctest.IntTestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantV   = 0
				wantN   = 0
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = PositiveInt.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_Byte(t *testing.T) {
	ser := Byte
	test.Test(ctest.ByteTestCases, ser, t)
	test.TestSkip(ctest.ByteTestCases, ser, t)
}

func TestVarint_Float64(t *testing.T) {
	t.Run("Float64 serializer should work correctly",
		func(t *testing.T) {
			ser := Float64
			test.Test(ctest.Float64TestCases, ser, t)
			test.TestSkip(ctest.Float64TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
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
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = Float64.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}

func TestVarint_Float32(t *testing.T) {
	t.Run("Float32 serializer should work correctly",
		func(t *testing.T) {
			ser := Float32
			test.Test(ctest.Float32TestCases, ser, t)
			test.TestSkip(ctest.Float32TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return an error",
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
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = Float32.Unmarshal(r)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})
}
