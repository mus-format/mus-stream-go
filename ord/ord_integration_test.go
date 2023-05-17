package ord

import (
	"testing"

	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/varint"
)

func TestIntegrationOrd(t *testing.T) {

	t.Run("Marshal, Unmarshal, Size Skip pointer", func(t *testing.T) {
		var (
			m = func() mustrm.MarshalerFn[*string] {
				return func(t *string, w mustrm.Writer) (n int, err error) {
					return MarshalPtr[string](t, mustrm.MarshalerFn[string](MarshalString),
						w)
				}
			}()
			u = func() mustrm.UnmarshalerFn[*string] {
				return func(r mustrm.Reader) (t *string, n int, err error) {
					return UnmarshalPtr[string](mustrm.UnmarshalerFn[string](UnmarshalString), r)
				}
			}()
			s = func() mustrm.SizerFn[*string] {
				return func(t *string) (size int) {
					return SizePtr[string](t, mustrm.SizerFn[string](SizeString))
				}
			}()
			sk = func() mustrm.SkipperFn {
				return func(r mustrm.Reader) (n int, err error) {
					return SkipPtr(mustrm.SkipperFn(SkipString), r)
				}
			}()
		)
		testdata.Test[*string](muscom_testdata.PointerTestCases, m, u, s, t)
		testdata.TestSkip[*string](muscom_testdata.PointerTestCases, m, sk, s, t)
	})

	t.Run("Marshal, Unmarshal, Size, Skip slice", func(t *testing.T) {
		var (
			m = func() mustrm.MarshalerFn[[]int] {
				return func(t []int, w mustrm.Writer) (n int, err error) {
					return MarshalSlice[int](t, mustrm.MarshalerFn[int](varint.MarshalInt),
						w)
				}
			}()
			u = func() mustrm.UnmarshalerFn[[]int] {
				return func(r mustrm.Reader) (t []int, n int, err error) {
					return UnmarshalSlice[int](mustrm.UnmarshalerFn[int](varint.UnmarshalInt),
						r)
				}
			}()
			s = func() mustrm.SizerFn[[]int] {
				return func(t []int) (size int) {
					return SizeSlice[int](t, mustrm.SizerFn[int](varint.SizeInt))
				}
			}()
			sk = func() mustrm.SkipperFn {
				return func(r mustrm.Reader) (n int, err error) {
					return SkipSlice(mustrm.SkipperFn(varint.SkipInt), r)
				}
			}()
		)
		testdata.Test[[]int](muscom_testdata.SliceTestCases, m, u, s, t)
		testdata.TestSkip[[]int](muscom_testdata.SliceTestCases, m, sk, s, t)
	})

	t.Run("Marshal, Unmarshal, Size, Skip map", func(t *testing.T) {
		var (
			m = func() mustrm.MarshalerFn[map[float32]uint8] {
				return func(t map[float32]uint8, w mustrm.Writer) (n int, err error) {
					return MarshalMap[float32, uint8](t,
						mustrm.MarshalerFn[float32](varint.MarshalFloat32),
						mustrm.MarshalerFn[uint8](varint.MarshalUint8),
						w)
				}
			}()
			u = func() mustrm.UnmarshalerFn[map[float32]uint8] {
				return func(r mustrm.Reader) (t map[float32]uint8, n int, err error) {
					return UnmarshalMap[float32, uint8](
						mustrm.UnmarshalerFn[float32](varint.UnmarshalFloat32),
						mustrm.UnmarshalerFn[uint8](varint.UnmarshalUint8),
						r)
				}
			}()
			s = func() mustrm.SizerFn[map[float32]uint8] {
				return func(t map[float32]uint8) (size int) {
					return SizeMap[float32, uint8](t,
						mustrm.SizerFn[float32](varint.SizeFloat32),
						mustrm.SizerFn[uint8](varint.SizeUint8))
				}
			}()
			sk = func() mustrm.SkipperFn {
				return func(r mustrm.Reader) (n int, err error) {
					return SkipMap(mustrm.SkipperFn(varint.SkipFloat32),
						mustrm.SkipperFn(varint.SkipUint8),
						r)
				}
			}()
		)
		testdata.Test[map[float32]uint8](muscom_testdata.MapTestCases, m, u, s, t)
		testdata.TestSkip[map[float32]uint8](muscom_testdata.MapTestCases, m, sk, s, t)
	})

}
