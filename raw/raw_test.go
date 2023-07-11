package raw

import (
	"errors"
	"testing"

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
				n, err = skipInteger64(r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, t)
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
				n, err = skipInteger32(r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, t)
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
				n, err = skipInteger16(r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, t)
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
				n, err = skipInteger8(r)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, t)
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
						v, n, err = UnmarshalFloat64(r)
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
						n, err = SkipFloat64(r)
					)
					com_testdata.TestSkipResults(wantN, n, wantErr, err, t)
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
						v, n, err = UnmarshalFloat32(r)
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
						n, err = SkipFloat32(r)
					)
					com_testdata.TestSkipResults(wantN, n, wantErr, err, t)
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
