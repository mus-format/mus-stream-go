package raw

import (
	"errors"
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/ymz-ncnk/mok"
	"golang.org/x/exp/constraints"
)

func TestRaw(t *testing.T) {

	t.Run("setUpUintFuncs", func(t *testing.T) {

		t.Run("ErrUnsupportedIntSize", func(t *testing.T) {
			wantErr := muscom.ErrUnsupportedIntSize
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

		t.Run("intSize == 32", func(t *testing.T) {
			setUpUintFuncs(32)
			if !muscom_testdata.ComparePtrs(marshalUint, marshalInteger32[uint]) {
				t.Error("unexpected marshalUint func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
				t.Error("unexpected unmarshalUint func")
			}
			if !muscom_testdata.ComparePtrs(sizeUint, sizeNum32[uint]) {
				t.Error("unexpected sizeUint func")
			}
			if !muscom_testdata.ComparePtrs(skipUint, skipInteger32) {
				t.Error("unexpected skipUint func")
			}
		})

		t.Run("intSize == 64", func(t *testing.T) {
			setUpUintFuncs(64)
			if !muscom_testdata.ComparePtrs(marshalUint, marshalInteger64[uint]) {
				t.Error("unexpected marshalUint func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
				t.Error("unexpected unmarshalUint func")
			}
			if !muscom_testdata.ComparePtrs(sizeUint, sizeNum64[uint]) {
				t.Error("unexpected sizeUint func")
			}
			if !muscom_testdata.ComparePtrs(skipUint, skipInteger64) {
				t.Error("unexpected skipUint func")
			}
		})

	})

	t.Run("setUpIntFuncs", func(t *testing.T) {

		t.Run("ErrUnsupportedIntSize", func(t *testing.T) {
			wantErr := muscom.ErrUnsupportedIntSize
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

		t.Run("intSize == 32", func(t *testing.T) {
			setUpIntFuncs(32)
			if !muscom_testdata.ComparePtrs(marshalInt, marshalInteger32[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if !muscom_testdata.ComparePtrs(sizeInt, sizeNum32[int]) {
				t.Error("unexpected sizeInt func")
			}
			if !muscom_testdata.ComparePtrs(skipInt, skipInteger32) {
				t.Error("unexpected skipInt func")
			}
		})

		t.Run("intSize == 64", func(t *testing.T) {
			setUpIntFuncs(64)
			if !muscom_testdata.ComparePtrs(marshalInt, marshalInteger64[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if !muscom_testdata.ComparePtrs(sizeInt, sizeNum64[int]) {
				t.Error("unexpected sizeInt func")
			}
			if !muscom_testdata.ComparePtrs(skipInt, skipInteger64) {
				t.Error("unexpected skipInt func")
			}
		})

	})

	t.Run("marshalInteger64 - write byte error", func(t *testing.T) {
		testMarshalIntegerError(0, marshalInteger64[int64], t)
		testMarshalIntegerError(1, marshalInteger64[int64], t)
		testMarshalIntegerError(2, marshalInteger64[int64], t)
		testMarshalIntegerError(3, marshalInteger64[int64], t)
		testMarshalIntegerError(4, marshalInteger64[int64], t)
		testMarshalIntegerError(5, marshalInteger64[int64], t)
		testMarshalIntegerError(6, marshalInteger64[int64], t)
		testMarshalIntegerError(7, marshalInteger64[int64], t)
	})

	t.Run("unmarshalInteger64 - read byte error", func(t *testing.T) {
		testUnmarshalIntegerError(0, unmarshalInteger64[int64], t)
		testUnmarshalIntegerError(1, unmarshalInteger64[int64], t)
		testUnmarshalIntegerError(2, unmarshalInteger64[int64], t)
		testUnmarshalIntegerError(3, unmarshalInteger64[int64], t)
		testUnmarshalIntegerError(4, unmarshalInteger64[int64], t)
		testUnmarshalIntegerError(5, unmarshalInteger64[int64], t)
		testUnmarshalIntegerError(6, unmarshalInteger64[int64], t)
		testUnmarshalIntegerError(7, unmarshalInteger64[int64], t)
	})

	t.Run("skipInteger64 - read byte error", func(t *testing.T) {
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
		muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
	})

	t.Run("marshalInteger32 - write byte error", func(t *testing.T) {
		testMarshalIntegerError(0, marshalInteger32[int32], t)
		testMarshalIntegerError(1, marshalInteger32[int32], t)
		testMarshalIntegerError(2, marshalInteger32[int32], t)
		testMarshalIntegerError(3, marshalInteger32[int32], t)
	})

	t.Run("unmarshalInteger32 - read byte error", func(t *testing.T) {
		testUnmarshalIntegerError(0, unmarshalInteger32[int32], t)
		testUnmarshalIntegerError(1, unmarshalInteger32[int32], t)
		testUnmarshalIntegerError(2, unmarshalInteger32[int32], t)
		testUnmarshalIntegerError(3, unmarshalInteger32[int32], t)
	})

	t.Run("skipInteger32 - read byte error", func(t *testing.T) {
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
		muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
	})

	t.Run("marshalInteger16 - write byte error", func(t *testing.T) {
		testMarshalIntegerError(0, marshalInteger16[int16], t)
		testMarshalIntegerError(1, marshalInteger16[int16], t)
	})

	t.Run("unmarshalInteger16 - read byte error", func(t *testing.T) {
		testUnmarshalIntegerError(0, unmarshalInteger16[int16], t)
		testUnmarshalIntegerError(1, unmarshalInteger16[int16], t)
	})

	t.Run("skipInteger16 - read byte error", func(t *testing.T) {
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
		muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
	})

	t.Run("marshalInteger8 - write byte error", func(t *testing.T) {
		testMarshalIntegerError(0, marshalInteger8[int8], t)
	})

	t.Run("unmarshalInteger8 - read byte error", func(t *testing.T) {
		testUnmarshalIntegerError(0, unmarshalInteger8[int8], t)
	})

	t.Run("skipInteger8 - read byte error", func(t *testing.T) {
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
		muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
	})

	t.Run("byte", func(t *testing.T) {
		var (
			m  = muss.MarshalerFn[byte](MarshalByte)
			u  = muss.UnmarshalerFn[byte](UnmarshalByte)
			s  = muss.SizerFn[byte](SizeByte)
			sk = muss.SkipperFn(SkipByte)
		)
		testdata.Test[byte](muscom_testdata.ByteTestCases, m, u, s, t)
		testdata.TestSkip[byte](muscom_testdata.ByteTestCases, m, sk, s, t)
	})

	t.Run("Unsigned", func(t *testing.T) {

		t.Run("uint64", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[uint64](MarshalUint64)
				u  = muss.UnmarshalerFn[uint64](UnmarshalUint64)
				s  = muss.SizerFn[uint64](SizeUint64)
				sk = muss.SkipperFn(SkipUint64)
			)
			testdata.Test[uint64](muscom_testdata.Uint64TestCases, m, u, s, t)
			testdata.TestSkip[uint64](muscom_testdata.Uint64TestCases, m, sk, s, t)
		})

		t.Run("uint32", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[uint32](MarshalUint32)
				u  = muss.UnmarshalerFn[uint32](UnmarshalUint32)
				s  = muss.SizerFn[uint32](SizeUint32)
				sk = muss.SkipperFn(SkipUint32)
			)
			testdata.Test[uint32](muscom_testdata.Uint32TestCases, m, u, s, t)
			testdata.TestSkip[uint32](muscom_testdata.Uint32TestCases, m, sk, s, t)
		})

		t.Run("uint16", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[uint16](MarshalUint16)
				u  = muss.UnmarshalerFn[uint16](UnmarshalUint16)
				s  = muss.SizerFn[uint16](SizeUint16)
				sk = muss.SkipperFn(SkipUint16)
			)
			testdata.Test[uint16](muscom_testdata.Uint16TestCases, m, u, s, t)
			testdata.TestSkip[uint16](muscom_testdata.Uint16TestCases, m, sk, s, t)
		})

		t.Run("uint8", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[uint8](MarshalUint8)
				u  = muss.UnmarshalerFn[uint8](UnmarshalUint8)
				s  = muss.SizerFn[uint8](SizeUint8)
				sk = muss.SkipperFn(SkipUint8)
			)
			testdata.Test[uint8](muscom_testdata.Uint8TestCases, m, u, s, t)
			testdata.TestSkip[uint8](muscom_testdata.Uint8TestCases, m, sk, s, t)
		})

		t.Run("uint", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[uint](MarshalUint)
				u  = muss.UnmarshalerFn[uint](UnmarshalUint)
				s  = muss.SizerFn[uint](SizeUint)
				sk = muss.SkipperFn(SkipUint)
			)
			testdata.Test[uint](muscom_testdata.UintTestCases, m, u, s, t)
			testdata.TestSkip[uint](muscom_testdata.UintTestCases, m, sk, s, t)
		})

	})

	t.Run("Signed", func(t *testing.T) {

		t.Run("int64", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[int64](MarshalInt64)
				u  = muss.UnmarshalerFn[int64](UnmarshalInt64)
				s  = muss.SizerFn[int64](SizeInt64)
				sk = muss.SkipperFn(SkipInt64)
			)
			testdata.Test[int64](muscom_testdata.Int64TestCases, m, u, s, t)
			testdata.TestSkip[int64](muscom_testdata.Int64TestCases, m, sk, s, t)
		})

		t.Run("int32", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[int32](MarshalInt32)
				u  = muss.UnmarshalerFn[int32](UnmarshalInt32)
				s  = muss.SizerFn[int32](SizeInt32)
				sk = muss.SkipperFn(SkipInt32)
			)
			testdata.Test[int32](muscom_testdata.Int32TestCases, m, u, s, t)
			testdata.TestSkip[int32](muscom_testdata.Int32TestCases, m, sk, s, t)
		})

		t.Run("int16", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[int16](MarshalInt16)
				u  = muss.UnmarshalerFn[int16](UnmarshalInt16)
				s  = muss.SizerFn[int16](SizeInt16)
				sk = muss.SkipperFn(SkipInt16)
			)
			testdata.Test[int16](muscom_testdata.Int16TestCases, m, u, s, t)
			testdata.TestSkip[int16](muscom_testdata.Int16TestCases, m, sk, s, t)
		})

		t.Run("int8", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[int8](MarshalInt8)
				u  = muss.UnmarshalerFn[int8](UnmarshalInt8)
				s  = muss.SizerFn[int8](SizeInt8)
				sk = muss.SkipperFn(SkipInt8)
			)
			testdata.Test[int8](muscom_testdata.Int8TestCases, m, u, s, t)
			testdata.TestSkip[int8](muscom_testdata.Int8TestCases, m, sk, s, t)
		})

		t.Run("int", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[int](MarshalInt)
				u  = muss.UnmarshalerFn[int](UnmarshalInt)
				s  = muss.SizerFn[int](SizeInt)
				sk = muss.SkipperFn(SkipInt)
			)
			testdata.Test[int](muscom_testdata.IntTestCases, m, u, s, t)
			testdata.TestSkip[int](muscom_testdata.IntTestCases, m, sk, s, t)
		})

	})

	t.Run("Float", func(t *testing.T) {

		t.Run("float64", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = muss.MarshalerFn[float64](MarshalFloat64)
					u  = muss.UnmarshalerFn[float64](UnmarshalFloat64)
					s  = muss.SizerFn[float64](SizeFloat64)
					sk = muss.SkipperFn(SkipFloat64)
				)
				testdata.Test[float64](muscom_testdata.Float64TestCases, m, u, s, t)
				testdata.TestSkip[float64](muscom_testdata.Float64TestCases, m, sk, s,
					t)
			})

			t.Run("Unmarshal - read byte error", func(t *testing.T) {
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
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil,
					t)
			})

			t.Run("Skip - read byte error", func(t *testing.T) {
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
				muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
			})

		})

		t.Run("float32", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = muss.MarshalerFn[float32](MarshalFloat32)
					u  = muss.UnmarshalerFn[float32](UnmarshalFloat32)
					s  = muss.SizerFn[float32](SizeFloat32)
					sk = muss.SkipperFn(SkipFloat32)
				)
				testdata.Test[float32](muscom_testdata.Float32TestCases, m, u, s, t)
				testdata.TestSkip[float32](muscom_testdata.Float32TestCases, m, sk, s,
					t)
			})

			t.Run("Unmarshal - read byte error", func(t *testing.T) {
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
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil,
					t)
			})

			t.Run("Skip - read byte error", func(t *testing.T) {
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
				muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
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
	muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
		t)
}
