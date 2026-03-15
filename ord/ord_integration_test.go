package ord_test

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/test"
	"github.com/mus-format/mus-stream-go/unsafe"
	"github.com/mus-format/mus-stream-go/varint"
)

func TestIntegrationOrd(t *testing.T) {
	t.Run("pointer", func(t *testing.T) {
		ser := ord.NewPtrSer(ord.String)
		test.Test(ctestutil.PointerTestCases, ser, t)
		test.TestSkip(ctestutil.PointerTestCases, ser, t)
	})

	t.Run("array", func(t *testing.T) {
		ser := unsafe.NewArraySer[[3]int](varint.Int)
		test.Test(ctestutil.ArrayTestCases, ser, t)
		test.TestSkip(ctestutil.ArrayTestCases, ser, t)
	})

	t.Run("valid array", func(t *testing.T) {
		ser := unsafe.NewValidArraySer[[3]int](varint.Int, nil)
		test.Test(ctestutil.ArrayTestCases, ser, t)
		test.TestSkip(ctestutil.ArrayTestCases, ser, t)
	})

	t.Run("slice", func(t *testing.T) {
		ser := ord.NewSliceSer(varint.Int)
		test.Test(ctestutil.SliceTestCases, ser, t)
		test.TestSkip(ctestutil.SliceTestCases, ser, t)
	})

	t.Run("valid slice", func(t *testing.T) {
		ser := ord.NewValidSliceSer(varint.Int, nil, nil)
		test.Test(ctestutil.SliceTestCases, ser, t)
		test.TestSkip(ctestutil.SliceTestCases, ser, t)
	})

	t.Run("map", func(t *testing.T) {
		ser := ord.NewMapSer(varint.Float32, varint.Uint8)
		test.Test(ctestutil.MapTestCases, ser, t)
		test.TestSkip(ctestutil.MapTestCases, ser, t)
	})

	t.Run("valid map", func(t *testing.T) {
		ser := ord.NewValidMapSer(varint.Float32, varint.Uint8, nil, nil, nil)
		test.Test(ctestutil.MapTestCases, ser, t)
		test.TestSkip(ctestutil.MapTestCases, ser, t)
	})
}
