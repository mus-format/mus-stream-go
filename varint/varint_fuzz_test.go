package varint

import (
	"bytes"
	"math"
	"testing"

	"github.com/mus-format/mus-stream-go/test"
)

// byte ------------------------------------------------------------------------

func FuzzVarint_Byte(f *testing.F) {
	seeds := []byte{0, 1, 127, 128, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v byte) {
		test.Test([]byte{v}, Byte, t)
		test.TestSkip([]byte{v}, Byte, t)
	})
}

func FuzzVarint_ByteUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Byte.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Byte.Skip(buf2)
	})
}

// uint64 ----------------------------------------------------------------------

func FuzzVarint_Uint64(f *testing.F) {
	seeds := []uint64{0, 1, 127, 128, 255, 256, math.MaxUint64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint64) {
		test.Test([]uint64{v}, Uint64, t)
		test.TestSkip([]uint64{v}, Uint64, t)
	})
}

func FuzzVarint_Uint64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint64.Skip(buf2)
	})
}

// uint32 ----------------------------------------------------------------------

func FuzzVarint_Uint32(f *testing.F) {
	seeds := []uint32{0, 1, 127, 128, 255, 256, math.MaxUint32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint32) {
		test.Test([]uint32{v}, Uint32, t)
		test.TestSkip([]uint32{v}, Uint32, t)
	})
}

func FuzzVarint_Uint32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint32.Skip(buf2)
	})
}

// uint16 ----------------------------------------------------------------------

func FuzzVarint_Uint16(f *testing.F) {
	seeds := []uint16{0, 1, 127, 128, 255, 256, math.MaxUint16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint16) {
		test.Test([]uint16{v}, Uint16, t)
		test.TestSkip([]uint16{v}, Uint16, t)
	})
}

func FuzzVarint_Uint16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint16.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint16.Skip(buf2)
	})
}

// uint8 -----------------------------------------------------------------------

func FuzzVarint_Uint8(f *testing.F) {
	seeds := []uint8{0, 1, 127, 128, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint8) {
		test.Test([]uint8{v}, Uint8, t)
		test.TestSkip([]uint8{v}, Uint8, t)
	})
}

func FuzzVarint_Uint8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint8.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint8.Skip(buf2)
	})
}

// uint ------------------------------------------------------------------------

func FuzzVarint_Uint(f *testing.F) {
	seeds := []uint{0, 1, 127, 128, 255, 256}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint) {
		test.Test([]uint{v}, Uint, t)
		test.TestSkip([]uint{v}, Uint, t)
	})
}

func FuzzVarint_UintUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint.Skip(buf2)
	})
}

// int64 -----------------------------------------------------------------------

func FuzzVarint_Int64(f *testing.F) {
	seeds := []int64{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt64, math.MaxInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		test.Test([]int64{v}, Int64, t)
		test.TestSkip([]int64{v}, Int64, t)
	})
}

func FuzzVarint_Int64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int64.Skip(buf2)
	})
}

// int32 -----------------------------------------------------------------------

func FuzzVarint_Int32(f *testing.F) {
	seeds := []int32{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt32, math.MaxInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		test.Test([]int32{v}, Int32, t)
		test.TestSkip([]int32{v}, Int32, t)
	})
}

func FuzzVarint_Int32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int32.Skip(buf2)
	})
}

// int16 -----------------------------------------------------------------------

func FuzzVarint_Int16(f *testing.F) {
	seeds := []int16{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt16, math.MaxInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		test.Test([]int16{v}, Int16, t)
		test.TestSkip([]int16{v}, Int16, t)
	})
}

func FuzzVarint_Int16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int16.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int16.Skip(buf2)
	})
}

// int8 ------------------------------------------------------------------------

func FuzzVarint_Int8(f *testing.F) {
	seeds := []int8{0, 1, -1, 127, -127, math.MinInt8, math.MaxInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		test.Test([]int8{v}, Int8, t)
		test.TestSkip([]int8{v}, Int8, t)
	})
}

func FuzzVarint_Int8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int8.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int8.Skip(buf2)
	})
}

// int -------------------------------------------------------------------------

func FuzzVarint_Int(f *testing.F) {
	seeds := []int{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		test.Test([]int{v}, Int, t)
		test.TestSkip([]int{v}, Int, t)
	})
}

func FuzzVarint_IntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int.Skip(buf2)
	})
}

// positive_int64 --------------------------------------------------------------

func FuzzVarint_PositiveInt64(f *testing.F) {
	seeds := []int64{0, 1, 127, 128, 255, 256, math.MaxInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		if v < 0 {
			return
		}
		test.Test([]int64{v}, PositiveInt64, t)
		test.TestSkip([]int64{v}, PositiveInt64, t)
	})
}

func FuzzVarint_PositiveInt64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt64.Skip(buf2)
	})
}

// positive_int32 --------------------------------------------------------------

func FuzzVarint_PositiveInt32(f *testing.F) {
	seeds := []int32{0, 1, 127, 128, 255, 256, math.MaxInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		if v < 0 {
			return
		}
		test.Test([]int32{v}, PositiveInt32, t)
		test.TestSkip([]int32{v}, PositiveInt32, t)
	})
}

func FuzzVarint_PositiveInt32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt32.Skip(buf2)
	})
}

// positive_int16 --------------------------------------------------------------

func FuzzVarint_PositiveInt16(f *testing.F) {
	seeds := []int16{0, 1, 127, 128, 255, 256, math.MaxInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		if v < 0 {
			return
		}
		test.Test([]int16{v}, PositiveInt16, t)
		test.TestSkip([]int16{v}, PositiveInt16, t)
	})
}

func FuzzVarint_PositiveInt16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt16.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt16.Skip(buf2)
	})
}

// positive_int8 ---------------------------------------------------------------

func FuzzVarint_PositiveInt8(f *testing.F) {
	seeds := []int8{0, 1, 127, math.MaxInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		if v < 0 {
			return
		}
		test.Test([]int8{v}, PositiveInt8, t)
		test.TestSkip([]int8{v}, PositiveInt8, t)
	})
}

func FuzzVarint_PositiveInt8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt8.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt8.Skip(buf2)
	})
}

// positive_int ----------------------------------------------------------------

func FuzzVarint_PositiveInt(f *testing.F) {
	seeds := []int{0, 1, 127, 128, 255, 256}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		if v < 0 {
			return
		}
		test.Test([]int{v}, PositiveInt, t)
		test.TestSkip([]int{v}, PositiveInt, t)
	})
}

func FuzzVarint_PositiveIntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt.Skip(buf2)
	})
}

// float64 ---------------------------------------------------------------------

func FuzzVarint_Float64(f *testing.F) {
	seeds := []float64{0, 1, -1, 0.1, -0.1, math.Pi, math.E, math.Inf(1), math.Inf(-1), math.NaN()}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float64) {
		test.Test([]float64{v}, Float64, t)
		test.TestSkip([]float64{v}, Float64, t)
	})
}

func FuzzVarint_Float64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Float64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Float64.Skip(buf2)
	})
}

// float32 ---------------------------------------------------------------------

func FuzzVarint_Float32(f *testing.F) {
	seeds := []float32{0, 1, -1, 0.1, -0.1, math.Pi, math.E, float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float32) {
		test.Test([]float32{v}, Float32, t)
		test.TestSkip([]float32{v}, Float32, t)
	})
}

func FuzzVarint_Float32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Float32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Float32.Skip(buf2)
	})
}
