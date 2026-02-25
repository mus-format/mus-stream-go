package ord_test

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/testutil"
	"github.com/mus-format/mus-stream-go/unsafe"
	"github.com/mus-format/mus-stream-go/varint"
)

func TestIntegrationOrd(t *testing.T) {
	t.Run("pointer", func(t *testing.T) {
		ser := ord.NewPtrSer[string](ord.String)
		testutil.Test[*string](ctestutil.PointerTestCases, ser, t)
		testutil.TestSkip[*string](ctestutil.PointerTestCases, ser, t)
	})

	t.Run("array", func(t *testing.T) {
		ser := unsafe.NewArraySer[[3]int, int](varint.Int)
		testutil.Test[[3]int](ctestutil.ArrayTestCases, ser, t)
		testutil.TestSkip[[3]int](ctestutil.ArrayTestCases, ser, t)
	})

	t.Run("valid array", func(t *testing.T) {
		ser := unsafe.NewValidArraySer[[3]int, int](varint.Int, nil)
		testutil.Test[[3]int](ctestutil.ArrayTestCases, ser, t)
		testutil.TestSkip[[3]int](ctestutil.ArrayTestCases, ser, t)
	})

	t.Run("slice", func(t *testing.T) {
		ser := ord.NewSliceSer[int](varint.Int)
		testutil.Test[[]int](ctestutil.SliceTestCases, ser, t)
		testutil.TestSkip[[]int](ctestutil.SliceTestCases, ser, t)
	})

	t.Run("valid slice", func(t *testing.T) {
		ser := ord.NewValidSliceSer[int](varint.Int, nil, nil)
		testutil.Test[[]int](ctestutil.SliceTestCases, ser, t)
		testutil.TestSkip[[]int](ctestutil.SliceTestCases, ser, t)
	})

	t.Run("map", func(t *testing.T) {
		ser := ord.NewMapSer[float32, uint8](varint.Float32, varint.Uint8)
		testutil.Test[map[float32]uint8](ctestutil.MapTestCases, ser, t)
		testutil.TestSkip[map[float32]uint8](ctestutil.MapTestCases, ser, t)
	})

	t.Run("valid map", func(t *testing.T) {
		ser := ord.NewValidMapSer[float32, uint8](varint.Float32, varint.Uint8, nil, nil, nil)
		testutil.Test[map[float32]uint8](ctestutil.MapTestCases, ser, t)
		testutil.TestSkip[map[float32]uint8](ctestutil.MapTestCases, ser, t)
	})
}
