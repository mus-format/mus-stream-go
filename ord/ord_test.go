package ord

import (
	"bytes"
	"errors"
	"math"
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	com_mock "github.com/mus-format/common-go/testdata/mock"
	mus "github.com/mus-format/mus-stream-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/mus-format/mus-stream-go/varint"
	"github.com/ymz-ncnk/mok"
)

func TestOrd(t *testing.T) {

	t.Run("bool", func(t *testing.T) {

		t.Run("All MarshalBool, UnmarshalBool, SizeBool, SkipBool functions must work correctly",
			func(t *testing.T) {
				var (
					m  = muss.MarshallerFn[bool](MarshalBool)
					u  = muss.UnmarshallerFn[bool](UnmarshalBool)
					s  = muss.SizerFn[bool](SizeBool)
					sk = muss.SkipperFn(SkipBool)
				)
				testdata.Test[bool](com_testdata.BoolTestCases, m, u, s, t)
				testdata.TestSkip[bool](com_testdata.BoolTestCases, m, sk, s, t)
			})

		t.Run("If Writer fails to write a byte, MarshalBool should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("write error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return wantErr },
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = MarshalBool(true, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read a byte, UnmarshalBool should return error",
			func(t *testing.T) {
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalBool should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV   bool = false
					wantN        = 1
					wantErr      = com.ErrWrongFormat
					r            = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalBool(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read a byte, SkipBool should return error",
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
					n, err = SkipBool(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("SkipBool should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrWrongFormat
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipBool(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

	})

	t.Run("string", func(t *testing.T) {

		t.Run("All MarshalString, UnmarshalString, SizeString, SkipString functions with default lenM, lenU, lenS must work correctly",
			func(t *testing.T) {
				var (
					m muss.MarshallerFn[string] = func(v string, w muss.Writer) (n int,
						err error) {
						return MarshalString(v, nil, w)
					}
					u muss.UnmarshallerFn[string] = func(r muss.Reader) (v string, n int,
						err error) {
						return UnmarshalString(nil, r)
					}
					s muss.SizerFn[string] = func(v string) (size int) {
						return SizeString(v, nil)
					}
					sk muss.SkipperFn = func(r muss.Reader) (n int, err error) {
						return SkipString(nil, r)
					}
				)
				testdata.Test[string](com_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("If Writer fails to write string length, MarshalStringVarint should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("marshal length error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = MarshalStringVarint("hello world", w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read string length, UnmarshalStringVarint should return error",
			func(t *testing.T) {
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
					v, n, err = UnmarshalStringVarint(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalVarintString should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 1
					wantErr = com.ErrNegativeLength
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalStringVarint(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read string content, UnmarshalStringVarint should return error",
			func(t *testing.T) {
				var (
					wantV   string = ""
					wantN          = 1 + 2
					wantErr        = errors.New("unmarshal string content error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterRead(
						func(p []byte) (n int, err error) {
							p[0] = 110
							p[1] = 111
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalStringVarint(r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("lenVl validator should protect against too much length",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 10
					wantErr = errors.New("lenVl validator error")
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = math.MaxInt64
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					r = func() mock.Reader {
						buf := &bytes.Buffer{}
						varint.MarshalInt64(math.MaxInt64, buf)
						return mock.NewReader().RegisterNReadByte(com.Uint64MaxVarintLen,
							func() (b byte, err error) {
								return buf.ReadByte()
							},
						)
					}()
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidStringVarint(lenVl, false, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If lenVl validator fails with an error, UnmarshalValidStringVarint should immediately return it if skip == false",
			func(t *testing.T) {
				var (
					wantV      = ""
					wantN      = 1
					wantErr    = errors.New("lenVl validator error")
					wantLength = 3
					lenVl      = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != wantLength {
								t.Errorf("unexpected length, want '%v' actual '%v'", wantLength,
									v)
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
					v, n, err = UnmarshalValidStringVarint(lenVl, false, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If lenVl validator fails with an error, UnmarshalValidStringVarint should return it and skip all bytes of the string if skip == true",
			func(t *testing.T) {
				var (
					wantV      = ""
					wantN      = 4
					wantErr    = errors.New("lenVl validator error")
					wantLength = 3
					lenVl      = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != wantLength {
								t.Errorf("unexpected length, want '%v' actual '%v'", wantLength,
									v)
							}
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterNReadByte(3, func() (b byte, err error) { return })
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidStringVarint(lenVl, true, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If string length == 0 lenVl should work", func(t *testing.T) {
			var (
				wantV                        = ""
				wantN                        = 1
				wantErr                      = errors.New("empty string")
				lenVl   com.ValidatorFn[int] = func(t int) (err error) {
					return wantErr
				}
				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, nil
					},
				)
				v, n, err = UnmarshalValidStringVarint(lenVl, false, r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("If Reader fails to read string content, UnmarshalValidStringVarint, with lenVl validator and skip == true, should return error",
			func(t *testing.T) {
				var (
					wantV      = ""
					wantN      = 2
					wantErr    = errors.New("skip error")
					wantLength = 3
					lenVl      = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != wantLength {
								t.Errorf("unexpected length, want '%v' actual '%v'", wantLength,
									v)
							}
							return errors.New("lenVl validator error")
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterReadByte(
						func() (b byte, err error) { return 0, nil },
					).RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalValidStringVarint(lenVl, true, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read string length, SkipString should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipStringVarint(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("SkipStringVarint should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrNegativeLength
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 1, nil
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipStringVarint(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read string content, SkipStringVarint should return error",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = errors.New("skip string content error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 6, nil
						},
					).RegisterReadByte(
						func() (b byte, err error) { return },
					).RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipStringVarint(r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

	})

	t.Run("pointer", func(t *testing.T) {

		t.Run("All MarshalPtr, UnmarshalPtr, SizePtr, SkipPtr functions must work correctly with nil pointer",
			func(t *testing.T) {
				var (
					m = func() muss.MarshallerFn[*string] {
						return func(v *string, w muss.Writer) (n int, err error) {
							return MarshalPtr(v, nil, w)
						}
					}()
					u = func() muss.UnmarshallerFn[*string] {
						return func(r muss.Reader) (t *string, n int, err error) {
							return UnmarshalPtr[string](nil, r)
						}
					}()
					s = func() muss.SizerFn[*string] {
						return func(v *string) (size int) {
							return SizePtr(v, nil)
						}
					}()
					sk = func() muss.SkipperFn {
						return func(r muss.Reader) (n int, err error) {
							return SkipPtr(nil, r)
						}
					}()
				)
				testdata.Test[*string]([]*string{nil}, m, u, s, t)
				testdata.TestSkip[*string]([]*string{nil}, m, sk, s, t)
			})

		t.Run("All MarshalPtr, UnmarshalPtr, SizePtr, SkipPtr functions must work correctly with not nil pointer",
			func(t *testing.T) {
				var (
					str1    = "one"
					str1Raw = append([]byte{6}, []byte(str1)...)
					ptr     = &str1
					m1      = func() mock.Marshaller[string] {
						return mock.NewMarshaller[string]().RegisterNMarshalMUS(2,
							func(v string, w muss.Writer) (n int, err error) {
								switch v {
								case str1:
									return 4, nil
								default:
									t.Fatalf("unexepcted string, want '%v' actual '%v'", str1, v)
									return
								}
							},
						)
					}()
					u1 = func() mock.Unmarshaller[string] {
						return mock.NewUnmarshaller[string]().RegisterNUnmarshalMUS(1,
							func(r muss.Reader) (v string, n int, err error) {
								return str1, len(str1Raw), nil
							},
						)
					}()
					s1 = func() mock.Sizer[string] {
						return mock.NewSizer[string]().RegisterNSizeMUS(2,
							func(v string) (size int) {
								switch v {
								case str1:
									return len(str1Raw)
								default:
									t.Fatalf("unexepcted string, want '%v' actual '%v'", str1, v)
									return
								}
							},
						)
					}()
					sk1 = func() mock.Skipper {
						return mock.NewSkipper().RegisterNSkipMUS(1,
							func(r muss.Reader) (n int, err error) {
								return len(str1Raw), nil
							},
						)
					}()
					m = func() muss.MarshallerFn[*string] {
						return func(v *string, w muss.Writer) (n int, err error) {
							return MarshalPtr(v, muss.Marshaller[string](m1), w)
						}
					}()
					u = func() muss.UnmarshallerFn[*string] {
						return func(r muss.Reader) (t *string, n int, err error) {
							return UnmarshalPtr(muss.Unmarshaller[string](u1), r)
						}
					}()
					s = func() muss.SizerFn[*string] {
						return func(v *string) (size int) {
							return SizePtr(v, muss.Sizer[string](s1))
						}
					}()
					sk = func() muss.SkipperFn {
						return func(r muss.Reader) (n int, err error) {
							return SkipPtr(muss.Skipper(sk1), r)
						}
					}()
				)
				testdata.Test[*string]([]*string{ptr}, m, u, s, t)
				testdata.TestSkip[*string]([]*string{ptr}, m, sk, s, t)
			})

		t.Run("If Writer fails to write nil flag == 0, MarshalPtr should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("write nil flag error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = MarshalPtr[string](nil, nil, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Writer fails to write nil flag == 1, MarshalPtr should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("write nil flag error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error {
							return wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					str    = "str"
					strPtr = &str
					n, err = MarshalPtr(strPtr, nil, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Writer fails to write pointer content, MarshalPtr should return error",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = errors.New("Marshaller error")
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
					str    = "str"
					strPtr = &str
					n, err = MarshalPtr[string](strPtr,
						muss.MarshallerFn[string](MarshalStringVarint),
						w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read nil flag, UnmarshalPtr should return error",
			func(t *testing.T) {
				var (
					wantV   *string = nil
					wantN           = 0
					wantErr         = errors.New("read nil flag error")
					r               = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalPtr[string](nil, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr,
					err,
					mocks,
					t)
			})

		t.Run("UnmarshalPtr should return ErrWrongFormat if meets wrong format", func(t *testing.T) {
			var (
				wantV   *string = nil
				wantN           = 1
				wantErr         = com.ErrWrongFormat
				r               = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 2, nil
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalPtr[string](nil, r)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("If Unmarshaller fails with an error, UnmarshalPtr should return it",
			func(t *testing.T) {
				var (
					wantV   *string = nil
					wantN           = 1
					wantErr         = errors.New("Unmarshaller error")
					r               = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, nil
						},
					)
					u = mock.NewUnmarshaller[string]().RegisterUnmarshalMUS(
						func(r muss.Reader) (t string, n int, err error) {
							return "", 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, u.Mock}
					v, n, err = UnmarshalPtr[string](u, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read nil flag, Skip should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("read nil flag error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipPtr(nil, r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrWrongFormat
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 3, nil
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipPtr(nil, r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Skipper fails with an error, Skip should return it",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = errors.New("Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, nil
						},
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipPtr(sk, r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

	})

	t.Run("slice", func(t *testing.T) {

		t.Run("All MarshalSlice, UnmarshalSlice, SizeSlice, SkipSlice functions with default lenM, lenU, lenS must work correctly for empty slice",
			func(t *testing.T) {
				var (
					sl = []string{}
					m  = func() muss.MarshallerFn[[]string] {
						return func(v []string, w muss.Writer) (n int, err error) {
							return MarshalSlice(v, nil, nil, w)
						}
					}()
					u = func() muss.UnmarshallerFn[[]string] {
						return func(r muss.Reader) (v []string, n int, err error) {
							return UnmarshalSlice[string](nil, nil, r)
						}
					}()
					s = func() muss.SizerFn[[]string] {
						return func(v []string) (size int) {
							return SizeSlice(v, nil, nil)
						}
					}()
					sk = func() muss.SkipperFn {
						return func(r muss.Reader) (n int, err error) {
							return SkipSlice(nil, nil, r)
						}
					}()
				)
				testdata.Test[[]string]([][]string{sl}, m, u, s, t)
				testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
			})

		t.Run("All MarshalSlice, UnmarshalSlice, SizeSlice, SkipSlice functions with default lenM, lenU, lenS must work correctly for not empty slice",
			func(t *testing.T) {
				testAllSliceFunctions(
					func(m1 muss.Marshaller[string]) muss.MarshallerFn[[]string] {
						return func(v []string, w muss.Writer) (n int, err error) {
							return MarshalSlice(v, nil, m1, w)
						}
					},
					func(u1 muss.Unmarshaller[string]) muss.UnmarshallerFn[[]string] {
						return func(r muss.Reader) (v []string, n int, err error) {
							return UnmarshalSlice[string](nil, u1, r)
						}
					},
					func(s1 muss.Sizer[string]) muss.SizerFn[[]string] {
						return func(v []string) (size int) {
							return SizeSlice(v, nil, s1)
						}
					},
					func(sk1 muss.Skipper) muss.SkipperFn {
						return func(r muss.Reader) (n int, err error) {
							return SkipSlice(nil, sk1, r)
						}
					},
					t,
				)
			})

		t.Run("All MarshalSliceVarint, UnmarshalSliceVarint, SizeSliceVarint, SkipSliceVarint functions must work correctly for empty slice",
			func(t *testing.T) {
				var (
					sl = []string{}
					m  = func() muss.MarshallerFn[[]string] {
						return func(v []string, w muss.Writer) (n int, err error) {
							return MarshalSliceVarint(v, nil, w)
						}
					}()
					u = func() muss.UnmarshallerFn[[]string] {
						return func(r muss.Reader) (v []string, n int, err error) {
							return UnmarshalSliceVarint[string](nil, r)
						}
					}()
					s = func() muss.SizerFn[[]string] {
						return func(v []string) (size int) {
							return SizeSliceVarint(v, nil)
						}
					}()
					sk = func() muss.SkipperFn {
						return func(r muss.Reader) (n int, err error) {
							return SkipSliceVarint(nil, r)
						}
					}()
				)
				testdata.Test[[]string]([][]string{sl}, m, u, s, t)
				testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
			})

		t.Run("All MarshalSliceVarint, UnmarshalSliceVarint, SizeSliceVarint, SkipSliceVarint functions wmust work correctly for not empty slice",
			func(t *testing.T) {
				testAllSliceFunctions(
					func(m1 muss.Marshaller[string]) muss.MarshallerFn[[]string] {
						return func(v []string, w muss.Writer) (n int, err error) {
							return MarshalSliceVarint(v, m1, w)
						}
					},
					func(u1 muss.Unmarshaller[string]) muss.UnmarshallerFn[[]string] {
						return func(r muss.Reader) (v []string, n int, err error) {
							return UnmarshalSliceVarint[string](u1, r)
						}
					},
					func(s1 muss.Sizer[string]) muss.SizerFn[[]string] {
						return func(v []string) (size int) {
							return SizeSliceVarint(v, s1)
						}
					},
					func(sk1 muss.Skipper) muss.SkipperFn {
						return func(r muss.Reader) (n int, err error) {
							return SkipSliceVarint(sk1, r)
						}
					},
					t,
				)
			})

		t.Run("If Writer fails to write slice length, MarshalSliceVarint should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("marshal length error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return wantErr },
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = MarshalSliceVarint([]uint{1}, nil, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Marshaller fails with an error, MarshalSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = errors.New("Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return nil },
					)
					m = mock.NewMarshaller[uint]().RegisterMarshalMUS(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = MarshalSliceVarint[uint]([]uint{1}, m, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read slice length, UnmarshalSliceVarint should return error",
			func(t *testing.T) {
				var (
					wantV   []string = nil
					wantN            = 0
					wantErr          = errors.New("unmarshal length error")
					r                = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalSliceVarint[string](nil, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalSliceVarint should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV   []string = nil
					wantN            = 1
					wantErr          = com.ErrNegativeLength
					r                = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalSliceVarint[string](nil, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Unmarshaller fails with an error, UnmarshalSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV   []string = make([]string, 2)
					wantN            = 3
					wantErr          = errors.New("Unmarshaller error")
					r                = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 4, nil
						},
					)
					u = mock.NewUnmarshaller[string]().RegisterUnmarshalMUS(
						func(r muss.Reader) (t string, n int, err error) {
							return "", 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalSliceVarint[string](u, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Skipper != nil and lenVl validator returns an error, UnmarshalValidSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = nil
					wantN          = 5
					wantErr        = errors.New("lenVl validator error")
					lenVl          = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) { return 4, nil },
					)
					mocks     = []*mok.Mock{r.Mock, sk.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](lenVl, nil, nil, sk, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Skipper returns an error, UnmarshalValidSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = nil
					wantN          = 4
					wantErr        = errors.New("Skipper error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 4, nil
						},
					)
					lenVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 5, v)
							}
							return errors.New("lenVl validator error")
						},
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) { return 3, wantErr },
					)
					mocks     = []*mok.Mock{sk.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](lenVl, nil, nil, sk, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Skipper == nil and lenVl validator returns an error, UnmarshalValidSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = nil
					wantN          = 1
					wantErr        = errors.New("lenVl Validator error")
					lenVl          = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 4, nil
						},
					)
					mocks     = []*mok.Mock{lenVl.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](lenVl, nil, nil, nil, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Validator returns an error, UnmarshalValidSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = []uint{10, 0, 0}
					wantN          = 4
					wantErr        = errors.New("Validator error")
					r              = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 6, nil },
					).RegisterReadByte(
						func() (b byte, err error) { return 10, nil },
					)
					vl = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return nil
						},
					).RegisterValidate(
						func(v uint) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return wantErr
						},
					)
					u = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					).RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 2, 1, nil
						},
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{vl.Mock, u.Mock, sk.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](nil, u, vl, sk, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Skipper fails with an error, UnmarshalValidSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = []uint{0, 0, 0}
					wantN   = 4
					wantErr = errors.New("Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 6, nil },
					).RegisterReadByte(
						func() (b byte, err error) { return 10, nil },
					)
					vl = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return errors.New("validator error")
						},
					)
					u = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{vl.Mock, u.Mock, sk.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](nil, u, vl, sk, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Reader fails to read slice length, SkipSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipSliceVarint(nil, r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("SkipSliceVarint should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrNegativeLength
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 1, nil },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipSliceVarint(nil, r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

	})

	t.Run("map", func(t *testing.T) {

		t.Run("All MarshalMap, UnmarshalMap, SizeMap, SkipMap functions with default lenM, lenU, lenS, lenSk must work correctly",
			func(t *testing.T) {
				testAllMapFunctions(
					func(m1 muss.Marshaller[string], m2 muss.Marshaller[uint]) muss.MarshallerFn[map[string]uint] {
						return func(v map[string]uint, w muss.Writer) (n int, err error) {
							return MarshalMap(v, nil, m1, m2, w)
						}
					},
					func(u1 muss.Unmarshaller[string], u2 muss.Unmarshaller[uint]) muss.UnmarshallerFn[map[string]uint] {
						return func(r muss.Reader) (v map[string]uint, n int, err error) {
							return UnmarshalMap[string, uint](nil, u1, u2, r)
						}
					},
					func(s1 muss.Sizer[string], s2 muss.Sizer[uint]) muss.SizerFn[map[string]uint] {
						return func(v map[string]uint) (size int) {
							return SizeMap(v, nil, s1, s2)
						}
					},
					func(sk1, sk2 muss.Skipper) muss.SkipperFn {
						return func(r muss.Reader) (n int, err error) {
							return SkipMap(nil, sk1, sk2, r)
						}
					},
					t,
				)
			})

		t.Run("All MarshalMap, UnmarshalMap, SizeMap, SkipMap functions must work correctly",
			func(t *testing.T) {
				testAllMapFunctions(
					func(m1 muss.Marshaller[string], m2 muss.Marshaller[uint]) muss.MarshallerFn[map[string]uint] {
						return func(v map[string]uint, w muss.Writer) (n int, err error) {
							return MarshalMapVarint(v, m1, m2, w)
						}
					},
					func(u1 muss.Unmarshaller[string], u2 muss.Unmarshaller[uint]) muss.UnmarshallerFn[map[string]uint] {
						return func(r muss.Reader) (v map[string]uint, n int, err error) {
							return UnmarshalMapVarint[string, uint](u1, u2, r)
						}
					},
					func(s1 muss.Sizer[string], s2 muss.Sizer[uint]) muss.SizerFn[map[string]uint] {
						return func(v map[string]uint) (size int) {
							return SizeMapVarint(v, s1, s2)
						}
					},
					func(sk1, sk2 muss.Skipper) muss.SkipperFn {
						return func(r muss.Reader) (n int, err error) {
							return SkipMapVarint(sk1, sk2, r)
						}
					},
					t,
				)
			})

		t.Run("If Writer fails to write map length, MarshalMapVarint should return it",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("marshal length error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return wantErr },
					)
					mocks  = []*mok.Mock{w.Mock}
					n, err = MarshalMapVarint[uint, uint](nil, nil, nil, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Key Marshaller fails with an error, MarshalMapVarint should return it",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = errors.New("key Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return nil },
					)
					m1 = mock.NewMarshaller[uint]().RegisterMarshalMUS(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock, m1.Mock}
					n, err = MarshalMapVarint[uint](map[uint]uint{1: 1}, m1, nil, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Value Marshaller fails with an error, MarshalMapVarint should return it",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = errors.New("value Marshaller error")
					w       = mock.NewWriter().RegisterWriteByte(
						func(c byte) error { return nil },
					)
					m1 = mock.NewMarshaller[uint]().RegisterMarshalMUS(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, nil
						},
					)
					m2 = mock.NewMarshaller[uint]().RegisterMarshalMUS(
						func(t uint, w muss.Writer) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks  = []*mok.Mock{w.Mock, m1.Mock, m2.Mock}
					n, err = MarshalMapVarint[uint, uint](map[uint]uint{1: 1}, m1, m2, w)
				)
				testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read map length, UnmarshalMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 0
					wantErr               = errors.New("unmarshal length error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 0, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalMapVarint[uint, uint](nil, nil, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalMapVarint should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					wantErr               = com.ErrNegativeLength
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock}
					v, n, err = UnmarshalMapVarint[uint, uint](nil, nil, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Key Unmarshaller fails with an error, UnmarshalMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 3
					wantErr = errors.New("key Unmarshaller error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, u1.Mock}
					v, n, err = UnmarshalMapVarint[uint, uint](u1, nil, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Unmarshaller fails with an error, UnmarshalMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 4
					wantErr = errors.New("value Unmarshaller error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 2, nil
						},
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 1, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {

							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, u1.Mock, u2.Mock}
					v, n, err = UnmarshalMapVarint[uint, uint](u1, u2, r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If lenVl validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 5
					wantErr               = errors.New("lenVl validator error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					lenVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return wantErr
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					).RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					).RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock, lenVl.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](lenVl, nil, nil, nil, nil,
						sk1,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Key Skipper fails with an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 2
					wantErr               = errors.New("key Skipper error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 4, nil
						},
					)
					lenVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return errors.New("lenVl validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, wantErr
						},
					)
					sk2       = mock.NewSkipper()
					mocks     = []*mok.Mock{r.Mock, lenVl.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](lenVl, nil, nil, nil,
						nil,
						sk1,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Value Skipper fails with an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 3
					wantErr               = errors.New("value Skipper error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					lenVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return errors.New("lenVl Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, lenVl.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](lenVl, nil, nil, nil,
						nil,
						sk1,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Key Skipper == nil and lenVl validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					wantErr               = errors.New("lenVl Validator error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					sk2   = mock.NewSkipper()
					lenVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					mocks     = []*mok.Mock{r.Mock, lenVl.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](lenVl, nil, nil, nil,
						nil,
						nil,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Value Skipper == nil and lenVl validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					wantErr               = errors.New("lenVl Validator error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					sk1   = mock.NewSkipper()
					lenVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					mocks     = []*mok.Mock{r.Mock, lenVl.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](lenVl, nil, nil, nil,
						nil,
						sk1,
						nil,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If KeyValidator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("key Validator error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return wantErr
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					).RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock, u1.Mock, v1.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, nil, v1, nil, sk1,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Key Validator != nil and Value Skipper fails with an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("value Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) { return errors.New("key Validator error") },
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, u1.Mock, v1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, nil, v1, nil, nil,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Value Skipper == nil and Key Validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 2
					wantErr = errors.New("key Validator error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, u1.Mock, v1.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, nil, v1, nil, nil,
						nil,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Key Validator != nil and Key Skipper fails with an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("key Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("key Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{r.Mock, u1.Mock, v1.Mock, sk1.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, nil, v1, nil, sk1,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Value Validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("value Validator error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v2 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 11 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 11, v)
							}
							return wantErr
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, nil
						},
					)
					mocks = []*mok.Mock{r.Mock, u1.Mock, u2.Mock, v2.Mock, sk1.Mock,
						sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, u2, nil, v2, sk1,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Value Validator != nil and Key Skipper fails with an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("key Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v2 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("value Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) { return 1, wantErr },
					)
					sk2   = mock.NewSkipper()
					mocks = []*mok.Mock{r.Mock, u1.Mock, u2.Mock, v2.Mock, sk1.Mock,
						sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, u2, nil, v2,
						sk1,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("If Value Validator != nil and Value Skipper fails with an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("value Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return nil
						},
					)
					v2 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("value Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) { return 1, nil },
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) { return 1, wantErr },
					)
					mocks = []*mok.Mock{r.Mock, u1.Mock, u2.Mock, v1.Mock, v2.Mock,
						sk1.Mock,
						sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, u2, v1, v2, sk1,
						sk2,
						r)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Skip should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrNegativeLength
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 1, nil },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipMapVarint(nil, nil, r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Reader fails to read map length, Skip should return error",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = errors.New("unmarshal length error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 0, wantErr },
					)
					mocks  = []*mok.Mock{r.Mock}
					n, err = SkipMapVarint(nil, nil, r)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

	})

}

func testAllMapFunctions(
	m func(m1 muss.Marshaller[string], m2 muss.Marshaller[uint]) muss.MarshallerFn[map[string]uint],
	u func(u1 muss.Unmarshaller[string], u2 muss.Unmarshaller[uint]) muss.UnmarshallerFn[map[string]uint],
	s func(s1 muss.Sizer[string], s2 muss.Sizer[uint]) muss.SizerFn[map[string]uint],
	sk func(sk1 muss.Skipper, sk2 muss.Skipper) muss.SkipperFn,
	t *testing.T,

) {
	var (
		str1         = "one"
		str1Raw      = append([]byte{6}, []byte(str1)...)
		str2         = "two"
		str2Raw      = append([]byte{6}, []byte(str2)...)
		int1    uint = 5
		int1Raw      = []byte{5}
		int2    uint = 8
		int2Raw      = []byte{8}
		mp           = map[string]uint{str1: int1, str2: int2}
		m1           = func() mock.Marshaller[string] {
			return mock.NewMarshaller[string]().RegisterNMarshalMUS(4,
				func(v string, w muss.Writer) (n int, err error) {
					switch v {
					case str1:
						return len(str1Raw), nil
					case str2:
						return len(str2Raw), nil
					default:
						t.Fatalf("unexepcted string, want '%v' or '%v' actual '%v'",
							str1, str2, v)
						return
					}
				},
			)
		}()
		m2 = func() mock.Marshaller[uint] {
			return mock.NewMarshaller[uint]().RegisterNMarshalMUS(4,
				func(v uint, w muss.Writer) (n int, err error) {
					switch v {
					case int1:
						return len(int1Raw), nil
					case int2:
						return len(int2Raw), nil
					default:
						t.Fatalf("unexepcted uint, want '%v' or '%v' actual '%v'",
							int1, int2, v)
						return
					}
				},
			)
		}()
		u1 = func() mock.Unmarshaller[string] {
			return mock.NewUnmarshaller[string]().RegisterUnmarshalMUS(
				func(r muss.Reader) (v string, n int, err error) {
					return str1, len(str1Raw), nil
				},
			).RegisterUnmarshalMUS(
				func(r muss.Reader) (t string, n int, err error) {
					return str2, len(str2Raw), nil
				},
			)
		}()
		u2 = func() mock.Unmarshaller[uint] {
			return mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
				func(r muss.Reader) (v uint, n int, err error) {
					return int1, len(int1Raw), nil
				},
			).RegisterUnmarshalMUS(
				func(r muss.Reader) (t uint, n int, err error) {
					return int2, len(int2Raw), nil
				},
			)
		}()
		s1 = func() mock.Sizer[string] {
			return mock.NewSizer[string]().RegisterNSizeMUS(4,
				func(v string) (size int) {
					switch v {
					case str1:
						return len(str1Raw)
					case str2:
						return len(str2Raw)
					default:
						t.Fatalf("unexepcted string, want '%v' or '%v' actual '%v'",
							str1, str2, v)
						return
					}
				},
			)
		}()
		s2 = func() mock.Sizer[uint] {
			return mock.NewSizer[uint]().RegisterNSizeMUS(4,
				func(v uint) (size int) {
					switch v {
					case int1:
						return len(int1Raw)
					case int2:
						return len(int2Raw)
					default:
						t.Fatalf("unexepcted uint, want '%v' or '%v' actual '%v'", int1,
							int2, v)
						return
					}
				},
			)
		}()
		sk1 = func() mock.Skipper {
			return mock.NewSkipper().RegisterSkipMUS(
				func(r muss.Reader) (n int, err error) {
					return len(str1Raw), nil
				},
			).RegisterSkipMUS(
				func(r muss.Reader) (n int, err error) {
					return len(str2Raw), nil
				},
			)
		}()
		sk2 = func() mock.Skipper {
			return mock.NewSkipper().RegisterSkipMUS(
				func(r muss.Reader) (n int, err error) {
					return len(int1Raw), nil
				},
			).RegisterSkipMUS(
				func(r muss.Reader) (n int, err error) {
					return len(int2Raw), nil
				},
			)
		}()
		mocks = []*mok.Mock{m1.Mock, m2.Mock, u1.Mock, u2.Mock, s1.Mock,
			s2.Mock}
	)
	testdata.Test[map[string]uint]([]map[string]uint{mp},
		m(m1, m2),
		u(u1, u2),
		s(s1, s2),
		t)
	testdata.TestSkip[map[string]uint]([]map[string]uint{mp},
		m(m1, m2),
		sk(sk1, sk2),
		s(s1, s2),
		t)
	if info := mok.CheckCalls(mocks); len(info) > 0 {
		t.Error(info)
	}
}

func testAllSliceFunctions(
	m func(m1 muss.Marshaller[string]) muss.MarshallerFn[[]string],
	u func(u1 muss.Unmarshaller[string]) muss.UnmarshallerFn[[]string],
	s func(s1 muss.Sizer[string]) muss.SizerFn[[]string],
	sk func(sk1 muss.Skipper) muss.SkipperFn,
	t *testing.T,
) {
	var (
		str1    = "one"
		str1Raw = append([]byte{6}, []byte(str1)...)
		str2    = "two"
		str2Raw = append([]byte{6}, []byte(str2)...)
		sl      = []string{str1, str2}

		m1 = func() mock.Marshaller[string] {
			return mock.NewMarshaller[string]().RegisterNMarshalMUS(4,
				func(v string, w muss.Writer) (n int, err error) {
					switch v {
					case str1:
						return 4, nil
					case str2:
						return 4, nil
					default:
						t.Fatalf("unexepcted string, want '%v' or '%v' actual '%v'",
							str1, str2, v)
						return
					}
				},
			)
		}()
		u1 = func() mock.Unmarshaller[string] {
			return mock.NewUnmarshaller[string]().RegisterUnmarshalMUS(
				func(r muss.Reader) (v string, n int, err error) {
					return str1, len(str1Raw), nil
				},
			).RegisterUnmarshalMUS(
				func(r muss.Reader) (t string, n int, err error) {
					return str2, len(str2Raw), nil
				},
			)
		}()
		s1 = func() mock.Sizer[string] {
			return mock.NewSizer[string]().RegisterNSizeMUS(4,
				func(v string) (size int) {
					switch v {
					case str1:
						return len(str1Raw)
					case str2:
						return len(str2Raw)
					default:
						t.Fatalf("unexepcted string, want '%v' or '%v' actual '%v'",
							str1, str2, v)
						return
					}
				},
			)
		}()
		sk1 = func() mock.Skipper {
			return mock.NewSkipper().RegisterSkipMUS(
				func(r muss.Reader) (n int, err error) {
					return len(str1Raw), nil
				},
			).RegisterSkipMUS(
				func(r muss.Reader) (n int, err error) {
					return len(str2Raw), nil
				},
			)
		}()
	)
	testdata.Test[[]string]([][]string{sl}, m(m1), u(u1), s(s1), t)
	testdata.TestSkip[[]string]([][]string{sl}, m(m1), sk(sk1), s(s1), t)
}

// StringVarint

func MarshalStringVarint(v string, w muss.Writer) (n int, err error) {
	return MarshalString(v, muss.MarshallerFn[int](varint.MarshalInt), w)
}

func UnmarshalStringVarint(r muss.Reader) (v string,
	n int, err error) {
	return UnmarshalValidStringVarint(nil, false, r)
}

func UnmarshalValidStringVarint(lenVl com.Validator[int], skip bool, r muss.Reader) (
	v string, n int, err error) {
	return UnmarshalValidString(mus.UnmarshallerFn[int](varint.UnmarshalInt),
		lenVl, skip, r)
}

func SizeStringVarint(v string) (n int) {
	return SizeString(v, mus.SizerFn[int](varint.SizeInt))
}

func SkipStringVarint(r muss.Reader) (n int, err error) {
	return SkipString(mus.UnmarshallerFn[int](varint.UnmarshalInt), r)
}

// SliceVarint

func MarshalSliceVarint[T any](v []T, m mus.Marshaller[T], w muss.Writer) (
	n int, err error) {
	return MarshalSlice(v, muss.MarshallerFn[int](varint.MarshalInt), m, w)
}

func UnmarshalSliceVarint[T any](u mus.Unmarshaller[T], r muss.Reader) (v []T,
	n int, err error) {
	return UnmarshalValidSliceVarint[T](nil, u, nil, nil, r)
}

func UnmarshalValidSliceVarint[T any](lenVl com.Validator[int],
	u mus.Unmarshaller[T],
	vl com.Validator[T],
	sk mus.Skipper,
	r muss.Reader,
) (v []T, n int, err error) {
	return UnmarshalValidSlice[T](muss.UnmarshallerFn[int](varint.UnmarshalInt),
		lenVl, u, vl, sk, r)
}

func SizeSliceVarint[T any](v []T, s mus.Sizer[T]) (size int) {
	return SizeSlice(v, muss.SizerFn[int](varint.SizeInt), s)
}

func SkipSliceVarint(sk mus.Skipper, r muss.Reader) (n int,
	err error) {
	return SkipSlice(muss.UnmarshallerFn[int](varint.UnmarshalInt), sk, r)
}

// MapVarint

func MarshalMapVarint[T comparable, V any](v map[T]V, m1 mus.Marshaller[T],
	m2 mus.Marshaller[V],
	w muss.Writer,
) (n int, err error) {
	return MarshalMap(v, muss.MarshallerFn[int](varint.MarshalInt), m1, m2, w)
}

func UnmarshalMapVarint[T comparable, V any](u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	r muss.Reader,
) (v map[T]V, n int, err error) {
	return UnmarshalValidMapVarint[T, V](nil, u1, u2, nil, nil, nil, nil, r)
}

func UnmarshalValidMapVarint[T comparable, V any](lenVl com.Validator[int],
	u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	vl1 com.Validator[T],
	vl2 com.Validator[V],
	sk1, sk2 mus.Skipper,
	r muss.Reader,
) (v map[T]V, n int, err error) {
	return UnmarshalValidMap[T, V](muss.UnmarshallerFn[int](varint.UnmarshalInt),
		lenVl, u1, u2, vl1, vl2, sk1, sk2, r)
}

func SizeMapVarint[T comparable, V any](v map[T]V, s1 mus.Sizer[T],
	s2 mus.Sizer[V]) (size int) {
	return SizeMap(v, muss.SizerFn[int](varint.SizeInt), s1, s2)
}

func SkipMapVarint(sk1, sk2 mus.Skipper, r muss.Reader) (n int, err error) {
	return SkipMap(muss.UnmarshallerFn[int](varint.UnmarshalInt), sk1, sk2, r)
}
