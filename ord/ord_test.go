package ord

import (
	"errors"
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	muscom_mock "github.com/mus-format/mus-common-go/testdata/mock"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/ymz-ncnk/mok"
)

func TestOrd(t *testing.T) {

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

		t.Run("Marshal - write error", func(t *testing.T) {
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

		t.Run("Unmarshal - read error", func(t *testing.T) {
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

		t.Run("Unmarshal - ErrWrongFormat", func(t *testing.T) {
			var (
				wantV   bool = false
				wantN        = 1
				wantErr      = muscom.ErrWrongFormat
				r            = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 3, nil
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalBool(r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
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
				n, err = SkipBool(r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - ErrWrongFormat", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = muscom.ErrWrongFormat
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 3, nil
					},
				)
				n, err = SkipBool(r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

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

		t.Run("Marshal - marshal length error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("marshal length error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error {
						return wantErr
					},
				)
				mocks  = []*mok.Mock{w.Mock}
				n, err = MarshalString("hello world", w)
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

		t.Run("Unmarshal - ErrNegativeLength", func(t *testing.T) {
			var (
				wantV   = ""
				wantN   = 1
				wantErr = muscom.ErrNegativeLength
				r       = mock.NewReader().RegisterReadByte(
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

		t.Run("Unmarshal - unmarshal string content error", func(t *testing.T) {
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
				v, n, err = UnmarshalString(r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - MaxLength validator error, skip = false",
			func(t *testing.T) {
				var (
					wantV      = ""
					wantN      = 1
					wantErr    = errors.New("MaxLength validator error")
					wantLength = 3
					maxLength  = muscom_mock.NewValidator[int]().RegisterValidate(
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
					v, n, err = UnmarshalValidString(maxLength, false, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength validator error, skip = true",
			func(t *testing.T) {
				var (
					wantV      = ""
					wantN      = 4
					wantErr    = errors.New("MaxLength validator error")
					wantLength = 3
					maxLength  = muscom_mock.NewValidator[int]().RegisterValidate(
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
					v, n, err = UnmarshalValidString(maxLength, true, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength validator error, skip error",
			func(t *testing.T) {
				var (
					wantV      = ""
					wantN      = 2
					wantErr    = errors.New("skip error")
					wantLength = 3
					maxLength  = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != wantLength {
								t.Errorf("unexpected length, want '%v' actual '%v'", wantLength,
									v)
							}
							return errors.New("MaxLength validator error")
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
					v, n, err = UnmarshalValidString(maxLength, true, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Skip - unmarshal length error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("unmarshal length error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				n, err = SkipString(r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - ErrNegativeLength", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = muscom.ErrNegativeLength
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 1, nil
					},
				)
				n, err = SkipString(r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - skip string content error", func(t *testing.T) {
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
				n, err = SkipString(r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

	t.Run("pointer", func(t *testing.T) {

		t.Run("Marshal, Unmarshal, Size, Skip nil pointer", func(t *testing.T) {
			var (
				m = func() muss.MarshalerFn[*string] {
					return func(v *string, w muss.Writer) (n int, err error) {
						return MarshalPtr(v, nil, w)
					}
				}()
				u = func() muss.UnmarshalerFn[*string] {
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

		t.Run("Marshal, Unmarshal, Size, Skip not nil pointer", func(t *testing.T) {
			var (
				str1    = "one"
				str1Raw = append([]byte{6}, []byte(str1)...)
				ptr     = &str1
				m1      = func() mock.Marshaler[string] {
					return mock.NewMarshaler[string]().RegisterNMarshalMUS(2,
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
				u1 = func() mock.Unmarshaler[string] {
					return mock.NewUnmarshaler[string]().RegisterNUnmarshalMUS(1,
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
				m = func() muss.MarshalerFn[*string] {
					return func(v *string, w muss.Writer) (n int, err error) {
						return MarshalPtr(v, muss.Marshaler[string](m1), w)
					}
				}()
				u = func() muss.UnmarshalerFn[*string] {
					return func(r muss.Reader) (t *string, n int, err error) {
						return UnmarshalPtr(muss.Unmarshaler[string](u1), r)
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

		t.Run("Marshal nil - write nil flag error", func(t *testing.T) {
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

		t.Run("Marshal not nil - write nil flag error", func(t *testing.T) {
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

		t.Run("Marshal not nil - Marshaler error", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = errors.New("Marshaler error")
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
					muss.MarshalerFn[string](MarshalString),
					w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Unmarshal - read nil flag error", func(t *testing.T) {
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
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr,
				err,
				mocks,
				t)
		})

		t.Run("Unmarshal - ErrWrongFormat", func(t *testing.T) {
			var (
				wantV   *string = nil
				wantN           = 1
				wantErr         = muscom.ErrWrongFormat
				r               = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 2, nil
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalPtr[string](nil, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("Unmarshal - Unmarshaler error", func(t *testing.T) {
			var (
				wantV   *string = nil
				wantN           = 1
				wantErr         = errors.New("Unmarshaler error")
				r               = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, nil
					},
				)
				u = mock.NewUnmarshaler[string]().RegisterUnmarshalMUS(
					func(r muss.Reader) (t string, n int, err error) {
						return "", 0, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock, u.Mock}
				v, n, err = UnmarshalPtr[string](u, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("Skip - read nil flag error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("read nil flag error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 0, wantErr
					},
				)
				n, err = SkipPtr(nil, r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - ErrWrongFormat", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = muscom.ErrWrongFormat
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 3, nil
					},
				)
				n, err = SkipPtr(nil, r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - Skipper error", func(t *testing.T) {
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
				n, err = SkipPtr(sk, r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

	t.Run("slice", func(t *testing.T) {

		t.Run("Marshal, Unmarshal, Size Skip empty slice", func(t *testing.T) {
			var (
				sl = []string{}
				m  = func() muss.MarshalerFn[[]string] {
					return func(v []string, w muss.Writer) (n int, err error) {
						return MarshalSlice(v, nil, w)
					}
				}()
				u = func() muss.UnmarshalerFn[[]string] {
					return func(r muss.Reader) (v []string, n int, err error) {
						return UnmarshalSlice[string](nil, r)
					}
				}()
				s = func() muss.SizerFn[[]string] {
					return func(v []string) (size int) {
						return SizeSlice(v, nil)
					}
				}()
				sk = func() muss.SkipperFn {
					return func(r muss.Reader) (n int, err error) {
						return SkipSlice(nil, r)
					}
				}()
			)
			testdata.Test[[]string]([][]string{sl}, m, u, s, t)
			testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
		})

		t.Run("Marshal, Unmarshal, Size, Skip not empty slice", func(t *testing.T) {
			var (
				str1    = "one"
				str1Raw = append([]byte{6}, []byte(str1)...)
				str2    = "two"
				str2Raw = append([]byte{6}, []byte(str2)...)
				sl      = []string{str1, str2}

				m1 = func() mock.Marshaler[string] {
					return mock.NewMarshaler[string]().RegisterNMarshalMUS(4,
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
				u1 = func() mock.Unmarshaler[string] {
					return mock.NewUnmarshaler[string]().RegisterUnmarshalMUS(
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
				m = func() muss.MarshalerFn[[]string] {
					return func(v []string, w muss.Writer) (n int, err error) {
						return MarshalSlice(v, muss.Marshaler[string](m1), w)
					}
				}()
				u = func() muss.UnmarshalerFn[[]string] {
					return func(r muss.Reader) (t []string, n int, err error) {
						return UnmarshalSlice(muss.Unmarshaler[string](u1), r)
					}
				}()
				s = func() muss.SizerFn[[]string] {
					return func(t []string) (size int) {
						return SizeSlice(t, muss.Sizer[string](s1))
					}
				}()
				sk = func() muss.SkipperFn {
					return func(r muss.Reader) (n int, err error) {
						return SkipSlice(muss.Skipper(sk1), r)
					}
				}()
			)
			testdata.Test[[]string]([][]string{sl}, m, u, s, t)
			testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
		})

		t.Run("Marshal - marshal length error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("marshal length error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error { return wantErr },
				)
				mocks  = []*mok.Mock{w.Mock}
				n, err = MarshalSlice([]uint{1}, nil, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Marshal - Marshaler error", func(t *testing.T) {
			var (
				wantN   = 2
				wantErr = errors.New("Marshaler error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error { return nil },
				)
				m = mock.NewMarshaler[uint]().RegisterMarshalMUS(
					func(t uint, w muss.Writer) (n int, err error) {
						return 1, wantErr
					},
				)
				mocks  = []*mok.Mock{w.Mock}
				n, err = MarshalSlice[uint]([]uint{1}, m, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Unmarshal - unmarshal length error", func(t *testing.T) {
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
				v, n, err = UnmarshalSlice[string](nil, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("Unmarhal - ErrNegativeLength", func(t *testing.T) {
			var (
				wantV   []string = nil
				wantN            = 1
				wantErr          = muscom.ErrNegativeLength
				r                = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 1, nil
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalSlice[string](nil, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("Unmarshal - Unmarshaler error", func(t *testing.T) {
			var (
				wantV   []string = make([]string, 2)
				wantN            = 3
				wantErr          = errors.New("Unmarshaler error")
				r                = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 4, nil
					},
				)
				u = mock.NewUnmarshaler[string]().RegisterUnmarshalMUS(
					func(r muss.Reader) (t string, n int, err error) {
						return "", 2, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalSlice[string](u, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - MaxLength validator error", func(t *testing.T) {
			var (
				wantV     []uint = nil
				wantN            = 5
				wantErr          = errors.New("MaxLength validator error")
				maxLength        = muscom_mock.NewValidator[int]().RegisterValidate(
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
				v, n, err = UnmarshalValidSlice[uint](maxLength, nil, nil, sk, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - MaxLength validator error, Skipper error",
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
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 5, v)
							}
							return errors.New("MaxLength validator error")
						},
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) { return 3, wantErr },
					)
					mocks     = []*mok.Mock{sk.Mock}
					v, n, err = UnmarshalValidSlice[uint](maxLength, nil, nil, sk, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength Validator error, Skipper == nil",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 1
					wantErr          = errors.New("MaxLength Validator error")
					maxLength        = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					r = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) {
							return 4, nil
						},
					)
					mocks     = []*mok.Mock{maxLength.Mock}
					v, n, err = UnmarshalValidSlice[uint](maxLength, nil, nil, nil, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - Validator error", func(t *testing.T) {
			var (
				wantV   []uint = []uint{10, 0, 0}
				wantN          = 4
				wantErr        = errors.New("Validator error")
				r              = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 6, nil },
				).RegisterReadByte(
					func() (b byte, err error) { return 10, nil },
				)
				vl = muscom_mock.NewValidator[uint]().RegisterValidate(
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
				u = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
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
				v, n, err = UnmarshalValidSlice[uint](nil, u, vl, sk, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - Validator error, Skipper error",
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
					vl = muscom_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return errors.New("validator error")
						},
					)
					u = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
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
					v, n, err = UnmarshalValidSlice[uint](nil, u, vl, sk, r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Skip - unmarshal length error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("unmarshal length error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 0, wantErr },
				)
				n, err = SkipSlice(nil, r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - ErrNegativeLength", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = muscom.ErrNegativeLength
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 1, nil },
				)
				n, err = SkipSlice(nil, r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

	t.Run("map", func(t *testing.T) {

		t.Run("Marshal, Unmarshal, Size, Skip", func(t *testing.T) {
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
				m1           = func() mock.Marshaler[string] {
					return mock.NewMarshaler[string]().RegisterNMarshalMUS(4,
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
				m2 = func() mock.Marshaler[uint] {
					return mock.NewMarshaler[uint]().RegisterNMarshalMUS(4,
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
				u1 = func() mock.Unmarshaler[string] {
					return mock.NewUnmarshaler[string]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v string, n int, err error) {
							return str1, len(str1Raw), nil
						},
					).RegisterUnmarshalMUS(
						func(r muss.Reader) (t string, n int, err error) {
							return str2, len(str2Raw), nil
						},
					)
				}()
				u2 = func() mock.Unmarshaler[uint] {
					return mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
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
				m = func() muss.MarshalerFn[map[string]uint] {
					return func(v map[string]uint, w muss.Writer) (n int, err error) {
						return MarshalMap(v, muss.Marshaler[string](m1),
							muss.Marshaler[uint](m2),
							w)
					}
				}()
				u = func() muss.UnmarshalerFn[map[string]uint] {
					return func(r muss.Reader) (t map[string]uint, n int, err error) {
						return UnmarshalMap(
							muss.Unmarshaler[string](u1),
							muss.Unmarshaler[uint](u2),
							r)
					}
				}()
				s = func() muss.SizerFn[map[string]uint] {
					return func(v map[string]uint) (size int) {
						return SizeMap(v,
							muss.Sizer[string](s1),
							muss.Sizer[uint](s2))
					}
				}()
				sk = func() muss.SkipperFn {
					return func(r muss.Reader) (n int, err error) {
						return SkipMap(muss.Skipper(sk1), muss.Skipper(sk2), r)
					}
				}()
				mocks = []*mok.Mock{m1.Mock, m2.Mock, u1.Mock, u2.Mock, s1.Mock,
					s2.Mock}
			)
			testdata.Test[map[string]uint]([]map[string]uint{mp}, m, u, s, t)
			testdata.TestSkip[map[string]uint]([]map[string]uint{mp}, m, sk, s, t)
			if info := mok.CheckCalls(mocks); len(info) > 0 {
				t.Error(info)
			}
		})

		t.Run("Marshal - marshal length error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("marshal length error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error { return wantErr },
				)
				mocks  = []*mok.Mock{w.Mock}
				n, err = MarshalMap[uint, uint](nil, nil, nil, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Marshal - key Marshaler error", func(t *testing.T) {
			var (
				wantN   = 2
				wantErr = errors.New("key Marshaler error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error { return nil },
				)
				m1 = mock.NewMarshaler[uint]().RegisterMarshalMUS(
					func(t uint, w muss.Writer) (n int, err error) {
						return 1, wantErr
					},
				)
				mocks  = []*mok.Mock{w.Mock, m1.Mock}
				n, err = MarshalMap[uint](map[uint]uint{1: 1}, m1, nil, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Marshal - value Marshaler error", func(t *testing.T) {
			var (
				wantN   = 3
				wantErr = errors.New("value Marshaler error")
				w       = mock.NewWriter().RegisterWriteByte(
					func(c byte) error { return nil },
				)
				m1 = mock.NewMarshaler[uint]().RegisterMarshalMUS(
					func(t uint, w muss.Writer) (n int, err error) {
						return 1, nil
					},
				)
				m2 = mock.NewMarshaler[uint]().RegisterMarshalMUS(
					func(t uint, w muss.Writer) (n int, err error) {
						return 1, wantErr
					},
				)
				mocks  = []*mok.Mock{w.Mock, m1.Mock, m2.Mock}
				n, err = MarshalMap[uint, uint](map[uint]uint{1: 1}, m1, m2, w)
			)
			testdata.TestMarshalResults(wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Unmarshal - unmarshal length error", func(t *testing.T) {
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
				v, n, err = UnmarshalMap[uint, uint](nil, nil, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("Unmarshal - ErrNegativeLength", func(t *testing.T) {
			var (
				wantV   map[uint]uint = nil
				wantN                 = 1
				wantErr               = muscom.ErrNegativeLength
				r                     = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 1, nil
					},
				)
				mocks     = []*mok.Mock{r.Mock}
				v, n, err = UnmarshalMap[uint, uint](nil, nil, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("Unmarshal - key Unmarshaler error", func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 1)
				wantN   = 3
				wantErr = errors.New("key Unmarshaler error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 2, nil
					},
				)
				u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(r muss.Reader) (v uint, n int, err error) {
						return 0, 2, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock, u1.Mock}
				v, n, err = UnmarshalMap[uint, uint](u1, nil, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("Unmarshal - value Unmarshaler error", func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 1)
				wantN   = 4
				wantErr = errors.New("value Unmarshaler error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						return 2, nil
					},
				)
				u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(r muss.Reader) (v uint, n int, err error) {
						return 1, 1, nil
					},
				)
				u2 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(r muss.Reader) (v uint, n int, err error) {

						return 0, 2, wantErr
					},
				)
				mocks     = []*mok.Mock{r.Mock, u1.Mock, u2.Mock}
				v, n, err = UnmarshalMap[uint, uint](u1, u2, r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - MaxLength Validator error", func(t *testing.T) {
			var (
				wantV   map[uint]uint = nil
				wantN                 = 5
				wantErr               = errors.New("MaxLength validator error")
				r                     = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 4, nil },
				)
				maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
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
				mocks     = []*mok.Mock{r.Mock, maxLength.Mock, sk1.Mock, sk2.Mock}
				v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil, nil,
					sk1,
					sk2,
					r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - MaxLength Validator error, key Skipper error",
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
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return errors.New("MaxLength validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 1, wantErr
						},
					)
					sk2       = mock.NewSkipper()
					mocks     = []*mok.Mock{r.Mock, maxLength.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						sk2,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength Validator error, value Skipper error",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 3
					wantErr               = errors.New("value Skipper error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return errors.New("MaxLength Validator error")
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
					mocks     = []*mok.Mock{r.Mock, maxLength.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						sk2,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength Validator erorr, key Skipper == nil",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					wantErr               = errors.New("MaxLength Validator error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					sk2       = mock.NewSkipper()
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					mocks     = []*mok.Mock{r.Mock, maxLength.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil,
						nil,
						nil,
						sk2,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength Validator erorr, value Skipper == nil",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					wantErr               = errors.New("MaxLength Validator error")
					r                     = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					sk1       = mock.NewSkipper()
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					mocks     = []*mok.Mock{r.Mock, maxLength.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						nil,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - key Validator error", func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 2)
				wantN   = 5
				wantErr = errors.New("key Validator error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 4, nil },
				)
				u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(r muss.Reader) (v uint, n int, err error) {
						return 10, 1, nil
					},
				)
				v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
				v, n, err = UnmarshalValidMap[uint, uint](nil, u1, nil, v1, nil, sk1,
					sk2,
					r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - key Validator error, value Skipper error",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("value Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) { return errors.New("key Validator error") },
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(r muss.Reader) (n int, err error) {
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, u1.Mock, v1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, nil, v1, nil, nil,
						sk2,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - key Validator error, value Skipper == nil",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 2
					wantErr = errors.New("key Validator error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{r.Mock, u1.Mock, v1.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, nil, v1, nil, nil,
						nil,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - key Validator error, key Skipper error",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("key Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, nil, v1, nil, sk1,
						sk2,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - value Validator error", func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 2)
				wantN   = 5
				wantErr = errors.New("value Validator error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 4, nil },
				)
				u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(r muss.Reader) (v uint, n int, err error) {
						return 10, 1, nil
					},
				)
				u2 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(r muss.Reader) (v uint, n int, err error) {
						return 11, 1, nil
					},
				)
				v2 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
				v, n, err = UnmarshalValidMap[uint, uint](nil, u1, u2, nil, v2, sk1,
					sk2,
					r)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks,
				t)
		})

		t.Run("UnmarshalValid - value Validator error, key Skipper error",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("key Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v2 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, u2, nil, v2,
						sk1,
						sk2,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("UnmarshalValid - value Validator error, value Skipper error",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("value Skipper error")
					r       = mock.NewReader().RegisterReadByte(
						func() (b byte, err error) { return 4, nil },
					)
					u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(r muss.Reader) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return nil
						},
					)
					v2 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, u2, v1, v2, sk1,
						sk2,
						r)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks,
					t)
			})

		t.Run("Skip - ErrNegativeLength", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = muscom.ErrNegativeLength
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 1, nil },
				)
				n, err = SkipMap(nil, nil, r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - unmarshal length error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("unmarshal length error")
				r       = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) { return 0, wantErr },
				)
				n, err = SkipMap(nil, nil, r)
			)
			muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

}
