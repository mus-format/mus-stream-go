package varint

import (
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-stream-go"
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser = test.UnmarshallerFn[uint64](func(r mus.Reader) (uint64, int, error) {
					return unmarshalUint[uint64](com.Uint64MaxVarintLen,
						com.Uint64MaxLastByte, r)
				})
				want = test.UnmarshalResults[uint64]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("unmarshalUint should return ErrOverflow if there is no varint end",
		func(t *testing.T) {
			var (
				wantErr = com.ErrOverflow
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 4, nil },
				)
				ser = test.UnmarshallerFn[uint16](func(r mus.Reader) (uint16, int, error) {
					return unmarshalUint[uint16](com.Uint16MaxVarintLen,
						com.Uint16MaxLastByte, r)
				})
				want = test.UnmarshalResults[uint16]{
					V:     0,
					N:     3,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestVarint_skipUint(t *testing.T) {
	t.Run("If Reader fails to read a byte, skipUint should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser = test.SkipperFn(func(r mus.Reader) (int, error) {
					return skipUint(com.Uint64MaxVarintLen, com.Uint64MaxLastByte,
						r)
				})
				want = test.SkipResults{
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestSkipOnly(r, ser, want, t)
		})

	t.Run("skipUint should return ErrOverflow if there is no varint end",
		func(t *testing.T) {
			var (
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
				ser = test.SkipperFn(func(r mus.Reader) (int, error) {
					return skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
						r)
				})
				want = test.SkipResults{
					N:     3,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestSkipOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = Int64
				want = test.UnmarshalResults[int64]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = Int32
				want = test.UnmarshalResults[int32]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = Int16
				want = test.UnmarshalResults[int16]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = Int8
				want = test.UnmarshalResults[int8]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = Int
				want = test.UnmarshalResults[int]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = PositiveInt64
				want = test.UnmarshalResults[int64]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = PositiveInt32
				want = test.UnmarshalResults[int32]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = PositiveInt16
				want = test.UnmarshalResults[int16]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = PositiveInt8
				want = test.UnmarshalResults[int8]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = PositiveInt
				want = test.UnmarshalResults[int]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = Float64
				want = test.UnmarshalResults[float64]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = Float32
				want = test.UnmarshalResults[float32]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}
