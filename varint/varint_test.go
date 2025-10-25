package varint

import (
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/ymz-ncnk/mok"
)

func TestVarint(t *testing.T) {
	t.Run("marshalUint", func(t *testing.T) {
		t.Run("If Writer fails to write a byte, marshalUint should return an error",
			func(t *testing.T) {
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

		t.Run("If Writer fails to write a last one byte, marshalUint should return an error",
			func(t *testing.T) {
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})
	})

	t.Run("skipUint", func(t *testing.T) {
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
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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

		t.Run("Uint16 serializer should work correctly",
			func(t *testing.T) {
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
		t.Run("int64", func(t *testing.T) {
			t.Run("Int64 serializer should work correctly",
				func(t *testing.T) {
					ser := Int64
					testdata.Test[int64](com_testdata.Int64TestCases, ser, t)
					testdata.TestSkip[int64](com_testdata.Int64TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("int32", func(t *testing.T) {
			t.Run("Int32 serializer should work correctly",
				func(t *testing.T) {
					ser := Int32
					testdata.Test[int32](com_testdata.Int32TestCases, ser, t)
					testdata.TestSkip[int32](com_testdata.Int32TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("int16", func(t *testing.T) {
			t.Run("Int16 serializer should work correctly",
				func(t *testing.T) {
					ser := Int16
					testdata.Test[int16](com_testdata.Int16TestCases, ser, t)
					testdata.TestSkip[int16](com_testdata.Int16TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("int8", func(t *testing.T) {
			t.Run("Int8 serializer should work correctly",
				func(t *testing.T) {
					ser := Int8
					testdata.Test[int8](com_testdata.Int8TestCases, ser, t)
					testdata.TestSkip[int8](com_testdata.Int8TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("int", func(t *testing.T) {
			t.Run("Int serializer should work correctly",
				func(t *testing.T) {
					ser := Int
					testdata.Test[int](com_testdata.IntTestCases, ser, t)
					testdata.TestSkip[int](com_testdata.IntTestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("positive_int64", func(t *testing.T) {
			t.Run("PositiveInt64 serializer should work correctly",
				func(t *testing.T) {
					ser := PositiveInt64
					testdata.Test[int64](com_testdata.Int64TestCases, ser, t)
					testdata.TestSkip[int64](com_testdata.Int64TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("positive_int32", func(t *testing.T) {
			t.Run("PositiveInt32 serializer should work correctly",
				func(t *testing.T) {
					ser := PositiveInt32
					testdata.Test[int32](com_testdata.Int32TestCases, ser, t)
					testdata.TestSkip[int32](com_testdata.Int32TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("positive_int16", func(t *testing.T) {
			t.Run("PositiveInt16 serializer should work correctly",
				func(t *testing.T) {
					ser := PositiveInt16
					testdata.Test[int16](com_testdata.Int16TestCases, ser, t)
					testdata.TestSkip[int16](com_testdata.Int16TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("positive_int8", func(t *testing.T) {
			t.Run("PositiveInt8 serializer should work correctly",
				func(t *testing.T) {
					ser := PositiveInt8
					testdata.Test[int8](com_testdata.Int8TestCases, ser, t)
					testdata.TestSkip[int8](com_testdata.Int8TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("positive_int", func(t *testing.T) {
			t.Run("PositiveInt serializer should work correctly",
				func(t *testing.T) {
					ser := PositiveInt
					testdata.Test[int](com_testdata.IntTestCases, ser, t)
					testdata.TestSkip[int](com_testdata.IntTestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})
	})

	t.Run("byte", func(t *testing.T) {
		t.Run("Byte serializer should work correctly",
			func(t *testing.T) {
				ser := Byte
				testdata.Test[byte](com_testdata.ByteTestCases, ser, t)
				testdata.TestSkip[byte](com_testdata.ByteTestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})

		t.Run("float32", func(t *testing.T) {
			t.Run("Float32 serializer should work correctly",
				func(t *testing.T) {
					ser := Float32
					testdata.Test[float32](com_testdata.Float32TestCases, ser, t)
					testdata.TestSkip[float32](com_testdata.Float32TestCases, ser, t)
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
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						mocks,
						t)
				})
		})
	})
}
