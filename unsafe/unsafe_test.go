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
	ctest "github.com/mus-format/common-go/test"
	cmock "github.com/mus-format/common-go/test/mock"
	"github.com/mus-format/mus-stream-go"
	arropts "github.com/mus-format/mus-stream-go/options/array"
	stropts "github.com/mus-format/mus-stream-go/options/string"
	"github.com/mus-format/mus-stream-go/raw"
	"github.com/mus-format/mus-stream-go/test"
	"github.com/mus-format/mus-stream-go/test/mock"
	"github.com/mus-format/mus-stream-go/varint"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

func TestUnsafe_setUpUintFuncs(t *testing.T) {
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
			if !ctest.ComparePtrs(sizeUint, raw.Uint.Size) {
				t.Error("unexpected sizeUint func")
			}
			if !ctest.ComparePtrs(skipUint, raw.Uint.Skip) {
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
			if !ctest.ComparePtrs(sizeUint, raw.Uint.Size) {
				t.Error("unexpected sizeUint func")
			}
			if !ctest.ComparePtrs(skipUint, raw.Uint.Skip) {
				t.Error("unexpected skipUint func")
			}
		})
}

func TestUnsafe_setUpIntFuncs(t *testing.T) {
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
			if !ctest.ComparePtrs(sizeInt, raw.Int.Size) {
				t.Error("unexpected sizeInt func")
			}
			if !ctest.ComparePtrs(skipInt, raw.Int.Skip) {
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
			if !ctest.ComparePtrs(sizeInt, raw.Int.Size) {
				t.Error("unexpected sizeInt func")
			}
			if !ctest.ComparePtrs(skipInt, raw.Int.Skip) {
				t.Error("unexpected skipInt func")
			}
		})
}

func TestUnsafe_IntegerErrorHandling(t *testing.T) {
	t.Run("If Reader fails to read a byte slice, unmarshalInteger64 should return an error", func(t *testing.T) {
		var (
			wantErr = errors.New("read byte error")
			r       = mock.NewReader().RegisterRead(
				func(p []byte) (n int, err error) {
					return 0, wantErr
				},
			)
			ser  = test.UnmarshallerFn[uint64](unmarshalInteger64[uint64])
			want = test.UnmarshalResults[uint64]{
				V:     0,
				N:     0,
				Err:   wantErr,
				Mocks: []*mok.Mock{r.Mock},
			}
		)
		test.TestUnmarshalOnly(r, ser, want, t)
	})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger64 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantErr = io.ErrUnexpectedEOF
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 4, nil
					},
				).RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				ser  = test.UnmarshallerFn[uint64](unmarshalInteger64[uint64])
				want = test.UnmarshalResults[uint64]{
					V:     0,
					N:     4,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails to read a byte slice, unmarshalInteger32 should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				ser  = test.UnmarshallerFn[uint32](unmarshalInteger32[uint32])
				want = test.UnmarshalResults[uint32]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger32 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantErr = io.ErrUnexpectedEOF
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 2, nil
					},
				).RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				ser  = test.UnmarshallerFn[uint32](unmarshalInteger32[uint32])
				want = test.UnmarshalResults[uint32]{
					V:     0,
					N:     2,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails to read a byte slice, unmarshalInteger16 should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				ser  = test.UnmarshallerFn[uint16](unmarshalInteger16[uint16])
				want = test.UnmarshalResults[uint16]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger16 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantErr = io.ErrUnexpectedEOF
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 1, nil
					},
				).RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				ser  = test.UnmarshallerFn[uint16](unmarshalInteger16[uint16])
				want = test.UnmarshalResults[uint16]{
					V:     0,
					N:     1,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails to read a byte slice, unmarshalInteger8 should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				ser  = test.UnmarshallerFn[uint8](unmarshalInteger8[uint8])
				want = test.UnmarshalResults[uint8]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails with an io.EOF, unmarshalInteger8 should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
				wantErr = io.EOF
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				ser  = test.UnmarshallerFn[uint8](unmarshalInteger8[uint8])
				want = test.UnmarshalResults[uint8]{
					V:     0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_String(t *testing.T) {
	t.Run("String serializer should work correctly",
		func(t *testing.T) {
			ser := String
			test.Test(ctest.StringTestCases, ser, t)
			test.TestSkip(ctest.StringTestCases, ser, t)
		})

	t.Run("We should be able to set a length serializer",
		func(t *testing.T) {
			var (
				str, lenSer = test.StringSerData(t)
				ser         = NewStringSer(stropts.WithLenSer(lenSer))
				mocks       = []*mok.Mock{lenSer.Mock}
			)
			test.Test([]string{str}, ser, t)
			test.TestSkip([]string{str}, ser, t)

			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap)
		})

	t.Run("If Writer fails to write a string length, Marshal should return an error",
		func(t *testing.T) {
			var (
				s       = "hello world"
				wantErr = errors.New("write error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				ser  = String
				want = test.MarshalResults{
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{w.Mock},
				}
			)
			test.TestMarshalOnly(s, w, ser, want, t)
		})

	t.Run("If Reader fails with an io.EOF, Unmarshal should return io.ErrUnexpectedEOF",
		func(t *testing.T) {
			var (
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
				ser  = String
				want = test.UnmarshalResults[string]{
					V:     "",
					N:     5,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails to read a string length, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("unmarshal length error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = String
				want = test.UnmarshalResults[string]{
					V:     "",
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
		func(t *testing.T) {
			var (
				wantErr = com.ErrNegativeLength
				r       = LengthReader(-1)
				ser     = String
				want    = test.UnmarshalResults[string]{
					V:     "",
					N:     10,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails to read string content, Unmarshal should return an error",
		func(t *testing.T) {
			var (
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
				ser  = String
				want = test.UnmarshalResults[string]{
					V:     "",
					N:     2,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("ValidString should work correctly",
		func(t *testing.T) {
			ser := NewValidStringSer(nil)
			test.Test(ctest.StringTestCases, ser, t)
			test.TestSkip(ctest.StringTestCases, ser, t)
		})

	t.Run("If lenVl validator returns an error, ValidString.Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantLength = math.MaxInt64 - 1
				wantErr    = errors.New("lenVl error")
				lenVl      = cmock.NewValidator[int]().RegisterValidate(
					func(length int) (err error) {
						if length != wantLength {
							t.Errorf("unexpected v, want '%v' actual '%v'", wantLength, length)
						}
						return wantErr
					},
				)
				r    = LengthReader(wantLength)
				ser  = NewValidStringSer(stropts.WithLenValidator(lenVl))
				want = test.UnmarshalResults[string]{
					V:     "",
					N:     9,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock, lenVl.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If string length == 0 lenVl should work", func(t *testing.T) {
		var (
			wantErr                      = errors.New("empty string")
			lenVl   com.ValidatorFn[int] = func(t int) (err error) {
				return wantErr
			}
			r = mock.NewReader().RegisterReadByte(
				func() (b byte, err error) {
					return 0, nil
				},
			)
			ser  = NewValidStringSer(stropts.WithLenValidator(lenVl))
			want = test.UnmarshalResults[string]{
				V:   "",
				N:   1,
				Err: wantErr,
			}
		)
		test.TestUnmarshalOnly(r, ser, want, t)
	})

	t.Run("If Reader fails to read a string length, ValidString.Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("unmarshal length error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				ser  = NewValidStringSer(nil)
				want = test.UnmarshalResults[string]{
					V:     "",
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("ValidString.Unmarshal should return ErrNegativeLength if meets negative length",
		func(t *testing.T) {
			var (
				wantErr = com.ErrNegativeLength
				r       = LengthReader(-1)
				ser     = NewValidStringSer(nil)
				want    = test.UnmarshalResults[string]{
					V:     "",
					N:     10,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If Reader fails to read string content, ValidString.Unmarshal should return an error",
		func(t *testing.T) {
			var (
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
				ser  = NewValidStringSer(nil)
				want = test.UnmarshalResults[string]{
					V:     "",
					N:     2,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_Byte(t *testing.T) {
	t.Run("Byte serializer should work correctly",
		func(t *testing.T) {
			ser := Byte
			test.Test(ctest.ByteTestCases, ser, t)
			test.TestSkip(ctest.ByteTestCases, ser, t)
		})
}

func TestUnsafe_Uint64(t *testing.T) {
	t.Run("Uint64 should work correctly",
		func(t *testing.T) {
			ser := Uint64
			test.Test(ctest.Uint64TestCases, ser, t)
			test.TestSkip(ctest.Uint64TestCases, ser, t)
		})
}

func TestUnsafe_Uint32(t *testing.T) {
	t.Run("Uint32 serializer should work correctly",
		func(t *testing.T) {
			ser := Uint32
			test.Test(ctest.Uint32TestCases, ser, t)
			test.TestSkip(ctest.Uint32TestCases, ser, t)
		})
}

func TestUnsafe_Uint16(t *testing.T) {
	t.Run("Uint16 serializer should work correctly",
		func(t *testing.T) {
			ser := Uint16
			test.Test(ctest.Uint16TestCases, ser, t)
			test.TestSkip(ctest.Uint16TestCases, ser, t)
		})
}

func TestUnsafe_Uint8(t *testing.T) {
	t.Run("Uint8 serializer should work correctly",
		func(t *testing.T) {
			ser := Uint8
			test.Test(ctest.Uint8TestCases, ser, t)
			test.TestSkip(ctest.Uint8TestCases, ser, t)
		})
}

func TestUnsafe_Uint(t *testing.T) {
	t.Run("Uint serializer should work correctly",
		func(t *testing.T) {
			ser := Uint
			test.Test(ctest.UintTestCases, ser, t)
			test.TestSkip(ctest.UintTestCases, ser, t)
		})
}

func TestUnsafe_Int64(t *testing.T) {
	t.Run("Int64 serializer should work correctly", func(t *testing.T) {
		ser := Int64
		test.Test(ctest.Int64TestCases, ser, t)
		test.TestSkip(ctest.Int64TestCases, ser, t)
	})
}

func TestUnsafe_Int32(t *testing.T) {
	t.Run("Int32 serializer should work correctly",
		func(t *testing.T) {
			ser := Int32
			test.Test(ctest.Int32TestCases, ser, t)
			test.TestSkip(ctest.Int32TestCases, ser, t)
		})
}

func TestUnsafe_Int16(t *testing.T) {
	t.Run("Int16 serializer should work correctly",
		func(t *testing.T) {
			ser := Int16
			test.Test(ctest.Int16TestCases, ser, t)
			test.TestSkip(ctest.Int16TestCases, ser, t)
		})
}

func TestUnsafe_Int8(t *testing.T) {
	t.Run("Int8 serializer should work correctly",
		func(t *testing.T) {
			ser := Int8
			test.Test(ctest.Int8TestCases, ser, t)
			test.TestSkip(ctest.Int8TestCases, ser, t)
		})
}

func TestUnsafe_Int(t *testing.T) {
	t.Run("Int serializer should work correctly",
		func(t *testing.T) {
			ser := Int
			test.Test(ctest.IntTestCases, ser, t)
			test.TestSkip(ctest.IntTestCases, ser, t)
		})
}

func TestUnsafe_Float64(t *testing.T) {
	t.Run("Float64 serializer should work correctly",
		func(t *testing.T) {
			ser := Float64
			test.Test(ctest.Float64TestCases, ser, t)
			test.TestSkip(ctest.Float64TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte slice, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				ser  = Float64
				want = test.UnmarshalResults[float64]{
					V:     0.0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_Float32(t *testing.T) {
	t.Run("Float32 serializer should work correctly",
		func(t *testing.T) {
			ser := Float32
			test.Test(ctest.Float32TestCases, ser, t)
			test.TestSkip(ctest.Float32TestCases, ser, t)
		})

	t.Run("If Reader fails to read a byte slice, Unmarshal should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				ser  = Float32
				want = test.UnmarshalResults[float32]{
					V:     0.0,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_Bool(t *testing.T) {
	t.Run("Bool serializer should work correctly",
		func(t *testing.T) {
			ser := Bool
			test.Test(ctest.BoolTestCases, ser, t)
			test.TestSkip(ctest.BoolTestCases, ser, t)
		})

	t.Run("If Writer fails to write a byte, Marshal should return an error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("write byte error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				ser  = Bool
				want = test.MarshalResults{
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{w.Mock},
				}
			)
			test.TestMarshalOnly(true, w, ser, want, t)
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
				ser  = Bool
				want = test.UnmarshalResults[bool]{
					V:     false,
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
		func(t *testing.T) {
			var (
				wantErr = com.ErrWrongFormat
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 3, nil
					},
				)
				ser  = Bool
				want = test.UnmarshalResults[bool]{
					V:     false,
					N:     1,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_TimeUnixUTC(t *testing.T) {
	_ = os.Setenv("TZ", "")
	t.Run("TimeUnixUTC serializer should work correctly", func(t *testing.T) {
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) { err = wantErr; return },
				)
				ser  = TimeUnixUTC
				want = test.UnmarshalResults[time.Time]{
					V:     time.Time{},
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_TimeUnixMilliUTC(t *testing.T) {
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
			test.Test([]time.Time{{}}, TimeUnix, t)
			test.TestSkip([]time.Time{{}}, TimeUnix, t)
		})

	t.Run("If Reader fails to read a byte, Unmarshal should return error",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) { err = wantErr; return },
				)
				ser  = TimeUnixMilliUTC
				want = test.UnmarshalResults[time.Time]{
					V:     time.Time{},
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_TimeUnixMicroUTC(t *testing.T) {
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) { err = wantErr; return },
				)
				ser  = TimeUnixMicroUTC
				want = test.UnmarshalResults[time.Time]{
					V:     time.Time{},
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_TimeUnixNanoUTC(t *testing.T) {
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
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterRead(
					func(p []byte) (n int, err error) { err = wantErr; return },
				)
				ser  = TimeUnixNanoUTC
				want = test.UnmarshalResults[time.Time]{
					V:     time.Time{},
					N:     0,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})
}

func TestUnsafe_Array(t *testing.T) {
	t.Run("Array serializer should work correctly", func(t *testing.T) {
		var (
			arr, elemSer = test.ArraySerData(t)
			mocks        = []*mok.Mock{elemSer.Mock}
			ser          = NewArraySer[[3]int](elemSer)
		)
		test.Test([][3]int{arr}, ser, t)
		test.TestSkip([][3]int{arr}, ser, t)

		asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap)
	})

	t.Run("Unmarshal of the too large array should return ErrTooLargeLength",
		func(t *testing.T) {
			var (
				wantErr = com.ErrTooLargeLength
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 4, nil
					},
				)
				ser  = NewArraySer[[3]int, int](nil)
				want = test.UnmarshalResults[[3]int]{
					V:     [3]int{0, 0, 0},
					N:     1,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("Valid array serializer should work correctly", func(t *testing.T) {
		var (
			arr, elemSer = test.ArraySerData(t)
			mocks        = []*mok.Mock{elemSer.Mock}
			ser          = NewValidArraySer[[3]int](elemSer, nil)
		)
		test.Test([][3]int{arr}, ser, t)
		test.TestSkip([][3]int{arr}, ser, t)

		asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap)
	})

	t.Run("Valid Unmarshal of the too large array should return ErrTooLargeLength",
		func(t *testing.T) {
			var (
				wantErr = com.ErrTooLargeLength
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 4, nil
					},
				)
				ser  = NewValidArraySer[[3]int, int](nil, nil)
				want = test.UnmarshalResults[[3]int]{
					V:     [3]int{0, 0, 0},
					N:     1,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("If elemVl returns an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantElem = 11
				wantErr  = errors.New("elemVl error")
				r        = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 3, nil
					},
				)
				elemSer = mock.NewSerializer[int]().RegisterUnmarshal(
					func(r mus.Reader) (t int, n int, err error) {
						return 11, 1, nil
					},
				)
				elemVl = cmock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) {
						if v != wantElem {
							return fmt.Errorf("unexpected v, want %v actual %v", wantElem, v)
						}
						return wantErr
					},
				)
				ser  = NewValidArraySer[[3]int](elemSer, arropts.WithElemValidator(elemVl))
				want = test.UnmarshalResults[[3]int]{
					V:     [3]int{0, 0, 0},
					N:     2,
					Err:   wantErr,
					Mocks: []*mok.Mock{r.Mock, elemSer.Mock, elemVl.Mock},
				}
			)
			test.TestUnmarshalOnly(r, ser, want, t)
		})

	t.Run("We should be able to set a length serializer for NewArraySer",
		func(t *testing.T) {
			var (
				arr = [3]int{1, 2, 3}
				ser = NewArraySer[[3]int](varint.Int, arropts.WithLenSer[int](varint.Int))
			)
			test.Test([][3]int{arr}, ser, t)
			test.TestSkip([][3]int{arr}, ser, t)
		})

	t.Run("We should be able to set a length serializer for NewValidArraySer",
		func(t *testing.T) {
			var (
				arr = [3]int{1, 2, 3}
				ser = NewValidArraySer[[3]int](varint.Int, arropts.WithLenSer[int](varint.Int))
			)
			test.Test([][3]int{arr}, ser, t)
			test.TestSkip([][3]int{arr}, ser, t)
		})
}

func LengthReader(length int) mock.Reader {
	r := mock.NewReader()
	buf := &bytes.Buffer{}
	_, _ = varint.PositiveInt.Marshal(length, buf)
	for _, b := range buf.Bytes() {
		func(b byte) {
			r.RegisterReadByte(func() (byte, error) {
				return b, nil
			})
		}(b)
	}
	return r
}
