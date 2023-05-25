package unsafe

import (
	"errors"
	"io"
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	muscom_mock "github.com/mus-format/mus-common-go/testdata/mock"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/ymz-ncnk/mok"
)

func TestUnsafe(t *testing.T) {

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
			if !muscom_testdata.ComparePtrs(sizeUint, raw.SizeUint) {
				t.Error("unexpected sizeUint func")
			}
			if !muscom_testdata.ComparePtrs(skipUint, raw.SkipUint) {
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
			if !muscom_testdata.ComparePtrs(sizeUint, raw.SizeUint) {
				t.Error("unexpected sizeUint func")
			}
			if !muscom_testdata.ComparePtrs(skipUint, raw.SkipUint) {
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

		t.Run("inSize == 32", func(t *testing.T) {
			setUpIntFuncs(32)
			if !muscom_testdata.ComparePtrs(marshalInt, marshalInteger32[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if !muscom_testdata.ComparePtrs(sizeInt, raw.SizeInt) {
				t.Error("unexpected sizeInt func")
			}
			if !muscom_testdata.ComparePtrs(skipInt, raw.SkipInt) {
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
			if !muscom_testdata.ComparePtrs(sizeInt, raw.SizeInt) {
				t.Error("unexpected sizeInt func")
			}
			if !muscom_testdata.ComparePtrs(skipInt, raw.SkipInt) {
				t.Error("unexpected skipInt func")
			}
		})

	})

	t.Run("unmarshalInteger64 - read error", func(t *testing.T) {
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
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
			mocks,
			t)
	})

	t.Run("unmarshalInteger64 - unexpected EOF", func(t *testing.T) {
		var (
			wantV   uint64 = 0
			wantN          = 4
			wantErr error  = io.ErrUnexpectedEOF
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
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
			mocks,
			t)
	})

	t.Run("unmarshalInteger32 - read error", func(t *testing.T) {
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
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
			mocks,
			t)
	})

	t.Run("unmarshalInteger32 - unexpected EOF", func(t *testing.T) {
		var (
			wantV   uint32 = 0
			wantN          = 2
			wantErr error  = io.ErrUnexpectedEOF
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
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
			mocks,
			t)
	})

	t.Run("unmarshalInteger16 - read error", func(t *testing.T) {
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
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
			mocks,
			t)
	})

	t.Run("unmarshalInteger16 - unexpected EOF", func(t *testing.T) {
		var (
			wantV   uint16 = 0
			wantN          = 1
			wantErr error  = io.ErrUnexpectedEOF
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
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
			mocks,
			t)
	})

	t.Run("unmarshalInteger8 - read error", func(t *testing.T) {
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
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
			mocks,
			t)
	})

	t.Run("unmarshalInteger8 - unexpected EOF", func(t *testing.T) {
		var (
			wantV   uint8 = 0
			wantN         = 0
			wantErr error = io.EOF
			r             = mock.NewReader().RegisterRead(
				func(p []byte) (n int, err error) { return 0, io.EOF },
			)
			mocks     = []*mok.Mock{r.Mock}
			v, n, err = unmarshalInteger8[uint8](r)
		)
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
			mocks,
			t)
	})

	t.Run("string", func(t *testing.T) {

		t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[string](MarshalString)
				u  = muss.UnmarshalerFn[string](UnmarshalString)
				s  = muss.SizerFn[string](SizeString)
				sk = muss.SkipperFn(SkipString)
			)
			testdata.Test[string](muscom_testdata.StringTestCases, m, u, s, t)
			testdata.TestSkip[string](muscom_testdata.StringTestCases, m, sk, s, t)
		})

		t.Run("Marshal - write error", func(t *testing.T) {
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
				n, err = MarshalString(s, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Unmarshal - unmarshal length error", func(t *testing.T) {
			var (
				wantV   string = ""
				wantN          = 0
				wantErr        = errors.New("unmarshal length error")
				r              = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalString(r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("Unmarshal - unexpected EOF", func(t *testing.T) {
			var (
				wantV   string = ""
				wantN          = 5
				wantErr error  = io.ErrUnexpectedEOF
				r              = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 22, nil },
				).RegisterRead(
					func(p []byte) (n int, err error) {
						return copy(p, "long"), nil
					},
				).RegisterRead(
					func(p []byte) (n int, err error) { return 0, io.EOF },
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalString(r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - ErrNegativeLength", func(t *testing.T) {
			var (
				wantV   string = ""
				wantN          = 1
				wantErr        = muscom.ErrNegativeLength
				r              = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 1, nil
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalString(r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - MaxLength validator error, skip == false",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = errors.New("MaxLength validator error")
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidString(maxLength, false, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength validator error, skip == true",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 4
					wantErr   = errors.New("MaxLength validator error")
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterNReadByte(3,
						func() (b byte, err error) {
							return 0, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidString(maxLength, true, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength validator error, skip == true, Reader error",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = errors.New("Reader error")
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return errors.New("MaxLength validator error")
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidString(maxLength, true, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Unmarshal - read string content error", func(t *testing.T) {
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
				v, n, err = UnmarshalString(r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

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

			t.Run("Unmarshal - read error", func(t *testing.T) {
				var (
					wantV   float64 = 0.0
					wantN           = 0
					wantErr         = errors.New("read byte error")
					r               = mock.NewReader().RegisterRead(
						func(p []byte) (n int, err error) {
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
					m  = muss.MarshalerFn[float32](MarshalFloat32)
					u  = muss.UnmarshalerFn[float32](UnmarshalFloat32)
					s  = muss.SizerFn[float32](SizeFloat32)
					sk = muss.SkipperFn(SkipFloat32)
				)
				testdata.Test[float32](muscom_testdata.Float32TestCases, m, u, s, t)
				testdata.TestSkip[float32](muscom_testdata.Float32TestCases, m, sk, s,
					t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
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
					v, n, err = UnmarshalFloat32(r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		})

	})

	t.Run("bool", func(t *testing.T) {

		t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
			var (
				m  = muss.MarshalerFn[bool](MarshalBool)
				u  = muss.UnmarshalerFn[bool](UnmarshalBool)
				s  = muss.SizerFn[bool](SizeBool)
				sk = muss.SkipperFn(SkipBool)
			)
			testdata.Test[bool](muscom_testdata.BoolTestCases, m, u, s, t)
			testdata.TestSkip[bool](muscom_testdata.BoolTestCases, m, sk, s, t)
		})

		t.Run("Marshal - write byte error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("write byte error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				mocks  = []*mok.Mock{w.Mock}
				n, err = MarshalBool(true, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Unmarshal - read byte error", func(t *testing.T) {
			var (
				wantV   bool = false
				wantN        = 0
				wantErr      = errors.New("read byte error")
				r            = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalBool(r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

	})

}
