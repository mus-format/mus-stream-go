package ord

import (
	"testing"

	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	"github.com/mus-format/mus-stream-go/varint"
)

func TestIntegrationOrd(t *testing.T) {

	t.Run("pointer", func(t *testing.T) {
		var (
			m = func() muss.MarshalerFn[*string] {
				return func(t *string, w muss.Writer) (n int, err error) {
					return MarshalPtr[string](t, muss.MarshalerFn[string](MarshalString),
						w)
				}
			}()
			u = func() muss.UnmarshalerFn[*string] {
				return func(r muss.Reader) (t *string, n int, err error) {
					return UnmarshalPtr[string](muss.UnmarshalerFn[string](UnmarshalString), r)
				}
			}()
			s = func() muss.SizerFn[*string] {
				return func(t *string) (size int) {
					return SizePtr[string](t, muss.SizerFn[string](SizeString))
				}
			}()
			sk = func() muss.SkipperFn {
				return func(r muss.Reader) (n int, err error) {
					return SkipPtr(muss.SkipperFn(SkipString), r)
				}
			}()
		)
		testdata.Test[*string](muscom_testdata.PointerTestCases, m, u, s, t)
		testdata.TestSkip[*string](muscom_testdata.PointerTestCases, m, sk, s, t)
	})

	t.Run("slice", func(t *testing.T) {
		var (
			m = func() muss.MarshalerFn[[]int] {
				return func(t []int, w muss.Writer) (n int, err error) {
					return MarshalSlice[int](t, muss.MarshalerFn[int](varint.MarshalInt),
						w)
				}
			}()
			u = func() muss.UnmarshalerFn[[]int] {
				return func(r muss.Reader) (t []int, n int, err error) {
					return UnmarshalSlice[int](muss.UnmarshalerFn[int](varint.UnmarshalInt),
						r)
				}
			}()
			s = func() muss.SizerFn[[]int] {
				return func(t []int) (size int) {
					return SizeSlice[int](t, muss.SizerFn[int](varint.SizeInt))
				}
			}()
			sk = func() muss.SkipperFn {
				return func(r muss.Reader) (n int, err error) {
					return SkipSlice(muss.SkipperFn(varint.SkipInt), r)
				}
			}()
		)
		testdata.Test[[]int](muscom_testdata.SliceTestCases, m, u, s, t)
		testdata.TestSkip[[]int](muscom_testdata.SliceTestCases, m, sk, s, t)
	})

	t.Run("map", func(t *testing.T) {
		var (
			m = func() muss.MarshalerFn[map[float32]uint8] {
				return func(t map[float32]uint8, w muss.Writer) (n int, err error) {
					return MarshalMap[float32, uint8](t,
						muss.MarshalerFn[float32](varint.MarshalFloat32),
						muss.MarshalerFn[uint8](varint.MarshalUint8),
						w)
				}
			}()
			u = func() muss.UnmarshalerFn[map[float32]uint8] {
				return func(r muss.Reader) (t map[float32]uint8, n int, err error) {
					return UnmarshalMap[float32, uint8](
						muss.UnmarshalerFn[float32](varint.UnmarshalFloat32),
						muss.UnmarshalerFn[uint8](varint.UnmarshalUint8),
						r)
				}
			}()
			s = func() muss.SizerFn[map[float32]uint8] {
				return func(t map[float32]uint8) (size int) {
					return SizeMap[float32, uint8](t,
						muss.SizerFn[float32](varint.SizeFloat32),
						muss.SizerFn[uint8](varint.SizeUint8))
				}
			}()
			sk = func() muss.SkipperFn {
				return func(r muss.Reader) (n int, err error) {
					return SkipMap(muss.SkipperFn(varint.SkipFloat32),
						muss.SkipperFn(varint.SkipUint8),
						r)
				}
			}()
		)
		testdata.Test[map[float32]uint8](muscom_testdata.MapTestCases, m, u, s, t)
		testdata.TestSkip[map[float32]uint8](muscom_testdata.MapTestCases, m, sk, s, t)
	})

}
