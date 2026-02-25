package varint

import (
	"bytes"
	"math"
	"testing"

	"github.com/mus-format/mus-stream-go/testutil"
)

// byte ------------------------------------------------------------------------

func FuzzByte(f *testing.F) {
	seeds := []byte{0, 1, 127, 128, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v byte) {
		testutil.Test[byte]([]byte{v}, Byte, t)
		testutil.TestSkip[byte]([]byte{v}, Byte, t)
	})
}

func FuzzByteUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Byte.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Byte.Skip(buf2)
	})
}

// uint64 ----------------------------------------------------------------------

func FuzzUint64(f *testing.F) {
	seeds := []uint64{0, 1, 127, 128, 255, 256, math.MaxUint64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint64) {
		testutil.Test[uint64]([]uint64{v}, Uint64, t)
		testutil.TestSkip[uint64]([]uint64{v}, Uint64, t)
	})
}

func FuzzUint64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint64.Skip(buf2)
	})
}

// uint32 ----------------------------------------------------------------------

func FuzzUint32(f *testing.F) {
	seeds := []uint32{0, 1, 127, 128, 255, 256, math.MaxUint32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint32) {
		testutil.Test[uint32]([]uint32{v}, Uint32, t)
		testutil.TestSkip[uint32]([]uint32{v}, Uint32, t)
	})
}

func FuzzUint32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint32.Skip(buf2)
	})
}

// uint16 ----------------------------------------------------------------------

func FuzzUint16(f *testing.F) {
	seeds := []uint16{0, 1, 127, 128, 255, 256, math.MaxUint16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint16) {
		testutil.Test[uint16]([]uint16{v}, Uint16, t)
		testutil.TestSkip[uint16]([]uint16{v}, Uint16, t)
	})
}

func FuzzUint16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint16.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint16.Skip(buf2)
	})
}

// uint8 -----------------------------------------------------------------------

func FuzzUint8(f *testing.F) {
	seeds := []uint8{0, 1, 127, 128, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint8) {
		testutil.Test[uint8]([]uint8{v}, Uint8, t)
		testutil.TestSkip[uint8]([]uint8{v}, Uint8, t)
	})
}

func FuzzUint8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint8.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint8.Skip(buf2)
	})
}

// uint ------------------------------------------------------------------------

func FuzzUint(f *testing.F) {
	seeds := []uint{0, 1, 127, 128, 255, 256}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint) {
		testutil.Test[uint]([]uint{v}, Uint, t)
		testutil.TestSkip[uint]([]uint{v}, Uint, t)
	})
}

func FuzzUintUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint.Skip(buf2)
	})
}

// int64 -----------------------------------------------------------------------

func FuzzInt64(f *testing.F) {
	seeds := []int64{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt64, math.MaxInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		testutil.Test[int64]([]int64{v}, Int64, t)
		testutil.TestSkip[int64]([]int64{v}, Int64, t)
	})
}

func FuzzInt64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int64.Skip(buf2)
	})
}

// int32 -----------------------------------------------------------------------

func FuzzInt32(f *testing.F) {
	seeds := []int32{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt32, math.MaxInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		testutil.Test[int32]([]int32{v}, Int32, t)
		testutil.TestSkip[int32]([]int32{v}, Int32, t)
	})
}

func FuzzInt32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int32.Skip(buf2)
	})
}

// int16 -----------------------------------------------------------------------

func FuzzInt16(f *testing.F) {
	seeds := []int16{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt16, math.MaxInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		testutil.Test[int16]([]int16{v}, Int16, t)
		testutil.TestSkip[int16]([]int16{v}, Int16, t)
	})
}

func FuzzInt16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int16.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int16.Skip(buf2)
	})
}

// int8 ------------------------------------------------------------------------

func FuzzInt8(f *testing.F) {
	seeds := []int8{0, 1, -1, 127, -127, math.MinInt8, math.MaxInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		testutil.Test[int8]([]int8{v}, Int8, t)
		testutil.TestSkip[int8]([]int8{v}, Int8, t)
	})
}

func FuzzInt8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int8.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int8.Skip(buf2)
	})
}

// int -------------------------------------------------------------------------

func FuzzInt(f *testing.F) {
	seeds := []int{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		testutil.Test[int]([]int{v}, Int, t)
		testutil.TestSkip[int]([]int{v}, Int, t)
	})
}

func FuzzIntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int.Skip(buf2)
	})
}

// positive_int64 --------------------------------------------------------------

func FuzzPositiveInt64(f *testing.F) {
	seeds := []int64{0, 1, 127, 128, 255, 256, math.MaxInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		if v < 0 {
			return
		}
		testutil.Test[int64]([]int64{v}, PositiveInt64, t)
		testutil.TestSkip[int64]([]int64{v}, PositiveInt64, t)
	})
}

func FuzzPositiveInt64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt64.Skip(buf2)
	})
}

// positive_int32 --------------------------------------------------------------

func FuzzPositiveInt32(f *testing.F) {
	seeds := []int32{0, 1, 127, 128, 255, 256, math.MaxInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		if v < 0 {
			return
		}
		testutil.Test[int32]([]int32{v}, PositiveInt32, t)
		testutil.TestSkip[int32]([]int32{v}, PositiveInt32, t)
	})
}

func FuzzPositiveInt32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt32.Skip(buf2)
	})
}

// positive_int16 --------------------------------------------------------------

func FuzzPositiveInt16(f *testing.F) {
	seeds := []int16{0, 1, 127, 128, 255, 256, math.MaxInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		if v < 0 {
			return
		}
		testutil.Test[int16]([]int16{v}, PositiveInt16, t)
		testutil.TestSkip[int16]([]int16{v}, PositiveInt16, t)
	})
}

func FuzzPositiveInt16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt16.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt16.Skip(buf2)
	})
}

// positive_int8 ---------------------------------------------------------------

func FuzzPositiveInt8(f *testing.F) {
	seeds := []int8{0, 1, 127, math.MaxInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		if v < 0 {
			return
		}
		testutil.Test[int8]([]int8{v}, PositiveInt8, t)
		testutil.TestSkip[int8]([]int8{v}, PositiveInt8, t)
	})
}

func FuzzPositiveInt8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt8.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt8.Skip(buf2)
	})
}

// positive_int ----------------------------------------------------------------

func FuzzPositiveInt(f *testing.F) {
	seeds := []int{0, 1, 127, 128, 255, 256}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		if v < 0 {
			return
		}
		testutil.Test[int]([]int{v}, PositiveInt, t)
		testutil.TestSkip[int]([]int{v}, PositiveInt, t)
	})
}

func FuzzPositiveIntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		PositiveInt.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		PositiveInt.Skip(buf2)
	})
}

// float64 ---------------------------------------------------------------------

func FuzzFloat64(f *testing.F) {
	seeds := []float64{0, 1, -1, 0.1, -0.1, math.Pi, math.E, math.Inf(1), math.Inf(-1), math.NaN()}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float64) {
		testutil.Test[float64]([]float64{v}, Float64, t)
		testutil.TestSkip[float64]([]float64{v}, Float64, t)
	})
}

func FuzzFloat64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Float64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Float64.Skip(buf2)
	})
}

// float32 ---------------------------------------------------------------------

func FuzzFloat32(f *testing.F) {
	seeds := []float32{0, 1, -1, 0.1, -0.1, math.Pi, math.E, float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float32) {
		testutil.Test[float32]([]float32{v}, Float32, t)
		testutil.TestSkip[float32]([]float32{v}, Float32, t)
	})
}

func FuzzFloat32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Float32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Float32.Skip(buf2)
	})
}
