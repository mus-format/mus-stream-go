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
		ser := ord.NewPtrSer(ord.String)
		testutil.Test(ctestutil.PointerTestCases, ser, t)
		testutil.TestSkip(ctestutil.PointerTestCases, ser, t)
	})

	t.Run("array", func(t *testing.T) {
		ser := unsafe.NewArraySer[[3]int](varint.Int)
		testutil.Test(ctestutil.ArrayTestCases, ser, t)
		testutil.TestSkip(ctestutil.ArrayTestCases, ser, t)
	})

	t.Run("valid array", func(t *testing.T) {
		ser := unsafe.NewValidArraySer[[3]int](varint.Int, nil)
		testutil.Test(ctestutil.ArrayTestCases, ser, t)
		testutil.TestSkip(ctestutil.ArrayTestCases, ser, t)
	})

	t.Run("slice", func(t *testing.T) {
		ser := ord.NewSliceSer(varint.Int)
		testutil.Test(ctestutil.SliceTestCases, ser, t)
		testutil.TestSkip(ctestutil.SliceTestCases, ser, t)
	})

	t.Run("valid slice", func(t *testing.T) {
		ser := ord.NewValidSliceSer(varint.Int, nil, nil)
		testutil.Test(ctestutil.SliceTestCases, ser, t)
		testutil.TestSkip(ctestutil.SliceTestCases, ser, t)
	})

	t.Run("map", func(t *testing.T) {
		ser := ord.NewMapSer(varint.Float32, varint.Uint8)
		testutil.Test(ctestutil.MapTestCases, ser, t)
		testutil.TestSkip(ctestutil.MapTestCases, ser, t)
	})

	t.Run("valid map", func(t *testing.T) {
		ser := ord.NewValidMapSer(varint.Float32, varint.Uint8, nil, nil, nil)
		testutil.Test(ctestutil.MapTestCases, ser, t)
		testutil.TestSkip(ctestutil.MapTestCases, ser, t)
	})
}
