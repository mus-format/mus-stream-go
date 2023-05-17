package varint

import (
	"errors"
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/ymz-ncnk/mok"
)

func TestVarint(t *testing.T) {

	t.Run("marshalUint", func(t *testing.T) {

		t.Run("write byte error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("write byte error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				mocks  = []*mok.Mock{w.Mock}
				n, err = marshalUint[uint](300, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("write last one byte error", func(t *testing.T) {
			var (
				wantN   = 1
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
				mocks  = []*mok.Mock{w.Mock}
				n, err = marshalUint[uint](300, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

	})

	t.Run("unmarshalUint", func(t *testing.T) {

		t.Run("Read byte error", func(t *testing.T) {
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
				v, n, err = unmarshalUint[uint64](muscom.Uint64MaxVarintLen,
					muscom.Uint64MaxLastByte, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("ErrOverflow", func(t *testing.T) {
			var (
				wantV   uint16 = 0
				wantN          = 3
				wantErr        = muscom.ErrOverflow
				r              = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 4, nil },
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = unmarshalUint[uint16](muscom.Uint16MaxVarintLen,
					muscom.Uint16MaxLastByte, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	})

	t.Run("skipUint", func(t *testing.T) {

		t.Run("Read byte error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("read byte error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				n, err = skipUint(muscom.Uint64MaxVarintLen, muscom.Uint64MaxLastByte,
					r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("ErrOverflow", func(t *testing.T) {
			var (
				wantN   = 3
				wantErr = muscom.ErrOverflow
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 200, nil },
				)
				n, err = skipUint(muscom.Uint16MaxVarintLen, muscom.Uint16MaxLastByte,
					r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

	t.Run("Unsigned", func(t *testing.T) {

		t.Run("uint64", func(t *testing.T) {
			var (
				m  = mustrm.MarshalerFn[uint64](MarshalUint64)
				u  = mustrm.UnmarshalerFn[uint64](UnmarshalUint64)
				s  = mustrm.SizerFn[uint64](SizeUint64)
				sk = mustrm.SkipperFn(SkipUint64)
			)
			testdata.Test[uint64](muscom_testdata.Uint64TestCases, m, u, s, t)
			testdata.TestSkip[uint64](muscom_testdata.Uint64TestCases, m, sk, s, t)
		})

		t.Run("uint32", func(t *testing.T) {
			var (
				m  = mustrm.MarshalerFn[uint32](MarshalUint32)
				u  = mustrm.UnmarshalerFn[uint32](UnmarshalUint32)
				s  = mustrm.SizerFn[uint32](SizeUint32)
				sk = mustrm.SkipperFn(SkipUint32)
			)
			testdata.Test[uint32](muscom_testdata.Uint32TestCases, m, u, s, t)
			testdata.TestSkip[uint32](muscom_testdata.Uint32TestCases, m, sk, s, t)
		})

		t.Run("uint16", func(t *testing.T) {
			var (
				m  = mustrm.MarshalerFn[uint16](MarshalUint16)
				u  = mustrm.UnmarshalerFn[uint16](UnmarshalUint16)
				s  = mustrm.SizerFn[uint16](SizeUint16)
				sk = mustrm.SkipperFn(SkipUint16)
			)
			testdata.Test[uint16](muscom_testdata.Uint16TestCases, m, u, s, t)
			testdata.TestSkip[uint16](muscom_testdata.Uint16TestCases, m, sk, s, t)
		})

		t.Run("uint8", func(t *testing.T) {
			var (
				m  = mustrm.MarshalerFn[uint8](MarshalUint8)
				u  = mustrm.UnmarshalerFn[uint8](UnmarshalUint8)
				s  = mustrm.SizerFn[uint8](SizeUint8)
				sk = mustrm.SkipperFn(SkipUint8)
			)
			testdata.Test[uint8](muscom_testdata.Uint8TestCases, m, u, s, t)
			testdata.TestSkip[uint8](muscom_testdata.Uint8TestCases, m, sk, s, t)
		})

		t.Run("uint", func(t *testing.T) {
			var (
				m  = mustrm.MarshalerFn[uint](MarshalUint)
				u  = mustrm.UnmarshalerFn[uint](UnmarshalUint)
				s  = mustrm.SizerFn[uint](SizeUint)
				sk = mustrm.SkipperFn(SkipUint)
			)
			testdata.Test[uint](muscom_testdata.UintTestCases, m, u, s, t)
			testdata.TestSkip[uint](muscom_testdata.UintTestCases, m, sk, s, t)
		})

	})

	t.Run("Signed", func(t *testing.T) {

		t.Run("int64", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = mustrm.MarshalerFn[int64](MarshalInt64)
					u  = mustrm.UnmarshalerFn[int64](UnmarshalInt64)
					s  = mustrm.SizerFn[int64](SizeInt64)
					sk = mustrm.SkipperFn(SkipInt64)
				)
				testdata.Test[int64](muscom_testdata.Int64TestCases, m, u, s, t)
				testdata.TestSkip[int64](muscom_testdata.Int64TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - read byte error", func(t *testing.T) {
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
					v, n, err = UnmarshalInt64(r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		})

		t.Run("int32", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = mustrm.MarshalerFn[int32](MarshalInt32)
					u  = mustrm.UnmarshalerFn[int32](UnmarshalInt32)
					s  = mustrm.SizerFn[int32](SizeInt32)
					sk = mustrm.SkipperFn(SkipInt32)
				)
				testdata.Test[int32](muscom_testdata.Int32TestCases, m, u, s, t)
				testdata.TestSkip[int32](muscom_testdata.Int32TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - read byte error", func(t *testing.T) {
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
					v, n, err = UnmarshalInt32(r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		})

		t.Run("int16", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = mustrm.MarshalerFn[int16](MarshalInt16)
					u  = mustrm.UnmarshalerFn[int16](UnmarshalInt16)
					s  = mustrm.SizerFn[int16](SizeInt16)
					sk = mustrm.SkipperFn(SkipInt16)
				)
				testdata.Test[int16](muscom_testdata.Int16TestCases, m, u, s, t)
				testdata.TestSkip[int16](muscom_testdata.Int16TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - read byte error", func(t *testing.T) {
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
					v, n, err = UnmarshalInt16(r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		})

		t.Run("int8", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = mustrm.MarshalerFn[int8](MarshalInt8)
					u  = mustrm.UnmarshalerFn[int8](UnmarshalInt8)
					s  = mustrm.SizerFn[int8](SizeInt8)
					sk = mustrm.SkipperFn(SkipInt8)
				)
				testdata.Test[int8](muscom_testdata.Int8TestCases, m, u, s, t)
				testdata.TestSkip[int8](muscom_testdata.Int8TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
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
					v, n, err = UnmarshalInt8(r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		})

		t.Run("int", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = mustrm.MarshalerFn[int](MarshalInt)
					u  = mustrm.UnmarshalerFn[int](UnmarshalInt)
					s  = mustrm.SizerFn[int](SizeInt)
					sk = mustrm.SkipperFn(SkipInt)
				)
				testdata.Test[int](muscom_testdata.IntTestCases, m, u, s, t)
				testdata.TestSkip[int](muscom_testdata.IntTestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
				var (
					wantV   int = 0
					wantN       = 0
					wantErr     = errors.New("read byte error")
					r           = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalInt(r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		})

	})

	t.Run("byte", func(t *testing.T) {
		var (
			m  = mustrm.MarshalerFn[byte](MarshalByte)
			u  = mustrm.UnmarshalerFn[byte](UnmarshalByte)
			s  = mustrm.SizerFn[byte](SizeByte)
			sk = mustrm.SkipperFn(SkipByte)
		)
		testdata.Test[byte](muscom_testdata.ByteTestCases, m, u, s, t)
		testdata.TestSkip[byte](muscom_testdata.ByteTestCases, m, sk, s, t)
	})

	t.Run("Float", func(t *testing.T) {

		t.Run("float64", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = mustrm.MarshalerFn[float64](MarshalFloat64)
					u  = mustrm.UnmarshalerFn[float64](UnmarshalFloat64)
					s  = mustrm.SizerFn[float64](SizeFloat64)
					sk = mustrm.SkipperFn(SkipFloat64)
				)
				testdata.Test[float64](muscom_testdata.Float64TestCases, m, u, s, t)
				testdata.TestSkip[float64](muscom_testdata.Float64TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
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
					v, n, err = UnmarshalFloat64(r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		})

		t.Run("float32", func(t *testing.T) {

			t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
				var (
					m  = mustrm.MarshalerFn[float32](MarshalFloat32)
					u  = mustrm.UnmarshalerFn[float32](UnmarshalFloat32)
					s  = mustrm.SizerFn[float32](SizeFloat32)
					sk = mustrm.SkipperFn(SkipFloat32)
				)
				testdata.Test[float32](muscom_testdata.Float32TestCases, m, u, s, t)
				testdata.TestSkip[float32](muscom_testdata.Float32TestCases, m, sk, s,
					t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
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
					v, n, err = UnmarshalFloat32(r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		})

	})

}
