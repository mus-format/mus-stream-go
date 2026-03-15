package ord_test

import (
	"testing"

	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/test"
	"github.com/mus-format/mus-stream-go/unsafe"
	"github.com/mus-format/mus-stream-go/varint"
)

func TestOrdIntegration_Pointer(t *testing.T) {
	ser := ord.NewPtrSer(ord.String)
	test.Test(ctest.PointerTestCases, ser, t)
	test.TestSkip(ctest.PointerTestCases, ser, t)
}

func TestOrdIntegration_Array(t *testing.T) {
	ser := unsafe.NewArraySer[[3]int](varint.Int)
	test.Test(ctest.ArrayTestCases, ser, t)
	test.TestSkip(ctest.ArrayTestCases, ser, t)
}

func TestOrdIntegration_ValidArray(t *testing.T) {
	ser := unsafe.NewValidArraySer[[3]int](varint.Int, nil)
	test.Test(ctest.ArrayTestCases, ser, t)
	test.TestSkip(ctest.ArrayTestCases, ser, t)
}

func TestOrdIntegration_Slice(t *testing.T) {
	ser := ord.NewSliceSer(varint.Int)
	test.Test(ctest.SliceTestCases, ser, t)
	test.TestSkip(ctest.SliceTestCases, ser, t)
}

func TestOrdIntegration_ValidSlice(t *testing.T) {
	ser := ord.NewValidSliceSer(varint.Int, nil, nil)
	test.Test(ctest.SliceTestCases, ser, t)
	test.TestSkip(ctest.SliceTestCases, ser, t)
}

func TestOrdIntegration_Map(t *testing.T) {
	ser := ord.NewMapSer(varint.Float32, varint.Uint8)
	test.Test(ctest.MapTestCases, ser, t)
	test.TestSkip(ctest.MapTestCases, ser, t)
}

func TestOrdIntegration_ValidMap(t *testing.T) {
	ser := ord.NewValidMapSer(varint.Float32, varint.Uint8, nil, nil, nil)
	test.Test(ctest.MapTestCases, ser, t)
	test.TestSkip(ctest.MapTestCases, ser, t)
}
