package unsafe

import (
	"bytes"
	"errors"
	"math"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	strops "github.com/mus-format/mus-stream-go/options/string"
	"github.com/mus-format/mus-stream-go/test"
	"github.com/mus-format/mus-stream-go/varint"
)

const maxLen = 1000

// bool ------------------------------------------------------------------------

func FuzzUnsafe_Bool(f *testing.F) {
	seeds := []bool{true, false}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v bool) {
		test.Test([]bool{v}, Bool, t)
		test.TestSkip([]bool{v}, Bool, t)
	})
}

func FuzzUnsafe_BoolUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Bool.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Bool.Skip(buf2)
	})
}

// byte ------------------------------------------------------------------------

func FuzzUnsafe_Byte(f *testing.F) {
	seeds := []byte{0, 1, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v byte) {
		test.Test([]byte{v}, Byte, t)
		test.TestSkip([]byte{v}, Byte, t)
	})
}

func FuzzUnsafe_ByteUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Byte.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Byte.Skip(buf2)
	})
}

// uint64 ----------------------------------------------------------------------

func FuzzUnsafe_Uint64(f *testing.F) {
	seeds := []uint64{0, 1, math.MaxUint64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint64) {
		test.Test([]uint64{v}, Uint64, t)
		test.TestSkip([]uint64{v}, Uint64, t)
	})
}

func FuzzUnsafe_Uint64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint64.Skip(buf2)
	})
}

// uint32 ----------------------------------------------------------------------

func FuzzUnsafe_Uint32(f *testing.F) {
	seeds := []uint32{0, 1, math.MaxUint32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint32) {
		test.Test([]uint32{v}, Uint32, t)
		test.TestSkip([]uint32{v}, Uint32, t)
	})
}

func FuzzUnsafe_Uint32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint32.Skip(buf2)
	})
}

// uint16 ----------------------------------------------------------------------

func FuzzUnsafe_Uint16(f *testing.F) {
	seeds := []uint16{0, 1, math.MaxUint16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint16) {
		test.Test([]uint16{v}, Uint16, t)
		test.TestSkip([]uint16{v}, Uint16, t)
	})
}

func FuzzUnsafe_Uint16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint16.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint16.Skip(buf2)
	})
}

// uint8 -----------------------------------------------------------------------

func FuzzUnsafe_Uint8(f *testing.F) {
	seeds := []uint8{0, 1, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint8) {
		test.Test([]uint8{v}, Uint8, t)
		test.TestSkip([]uint8{v}, Uint8, t)
	})
}

func FuzzUnsafe_Uint8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint8.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint8.Skip(buf2)
	})
}

// uint ------------------------------------------------------------------------

func FuzzUnsafe_Uint(f *testing.F) {
	seeds := []uint{0, 1}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint) {
		test.Test([]uint{v}, Uint, t)
		test.TestSkip([]uint{v}, Uint, t)
	})
}

func FuzzUnsafe_UintUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Uint.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Uint.Skip(buf2)
	})
}

// int64 -----------------------------------------------------------------------

func FuzzUnsafe_Int64(f *testing.F) {
	seeds := []int64{0, 1, -1, math.MaxInt64, math.MinInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		test.Test([]int64{v}, Int64, t)
		test.TestSkip([]int64{v}, Int64, t)
	})
}

func FuzzUnsafe_Int64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int64.Skip(buf2)
	})
}

// int32 -----------------------------------------------------------------------

func FuzzUnsafe_Int32(f *testing.F) {
	seeds := []int32{0, 1, -1, math.MaxInt32, math.MinInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		test.Test([]int32{v}, Int32, t)
		test.TestSkip([]int32{v}, Int32, t)
	})
}

func FuzzUnsafe_Int32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int32.Skip(buf2)
	})
}

// int16 -----------------------------------------------------------------------

func FuzzUnsafe_Int16(f *testing.F) {
	seeds := []int16{0, 1, -1, math.MaxInt16, math.MinInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		test.Test([]int16{v}, Int16, t)
		test.TestSkip([]int16{v}, Int16, t)
	})
}

func FuzzUnsafe_Int16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int16.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int16.Skip(buf2)
	})
}

// int8 ------------------------------------------------------------------------

func FuzzUnsafe_Int8(f *testing.F) {
	seeds := []int8{0, 1, -1, math.MaxInt8, math.MinInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		test.Test([]int8{v}, Int8, t)
		test.TestSkip([]int8{v}, Int8, t)
	})
}

func FuzzUnsafe_Int8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int8.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int8.Skip(buf2)
	})
}

// int -------------------------------------------------------------------------

func FuzzUnsafe_Int(f *testing.F) {
	seeds := []int{0, 1, -1}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		test.Test([]int{v}, Int, t)
		test.TestSkip([]int{v}, Int, t)
	})
}

func FuzzUnsafe_IntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Int.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Int.Skip(buf2)
	})
}

// float64 ---------------------------------------------------------------------

func FuzzUnsafe_Float64(f *testing.F) {
	seeds := []float64{0, 1, -1, 3.14, math.Inf(1), math.Inf(-1), math.NaN()}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float64) {
		test.Test([]float64{v}, Float64, t)
		test.TestSkip([]float64{v}, Float64, t)
	})
}

func FuzzUnsafe_Float64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Float64.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Float64.Skip(buf2)
	})
}

// float32 ---------------------------------------------------------------------

func FuzzUnsafe_Float32(f *testing.F) {
	seeds := []float32{0, 1, -1, 3.14, float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float32) {
		test.Test([]float32{v}, Float32, t)
		test.TestSkip([]float32{v}, Float32, t)
	})
}

func FuzzUnsafe_Float32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Float32.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Float32.Skip(buf2)
	})
}

// string ----------------------------------------------------------------------

func FuzzUnsafe_String(f *testing.F) {
	seeds := []string{"", "hello", "world", "mus-format"}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v string) {
		test.Test([]string{v}, String, t)
		test.TestSkip([]string{v}, String, t)
	})
}

func FuzzUnsafe_StringUnmarshal(f *testing.F) {
	// We use Valid serializer to avoid OOM during fuzzing.
	ser := NewValidStringSer(strops.WithLenValidator(
		com.ValidatorFn[int](func(v int) error {
			if v > maxLen {
				return errors.New("too large length")
			}
			return nil
		}),
	))
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		ser.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		ser.Skip(buf2)
	})
}

// time ------------------------------------------------------------------------

func FuzzUnsafe_TimeUnix(f *testing.F) {
	f.Fuzz(func(t *testing.T, sec int64) {
		v := time.Unix(sec, 0)
		test.Test([]time.Time{v}, TimeUnix, t)
		test.TestSkip([]time.Time{v}, TimeUnix, t)
	})
}

func FuzzUnsafe_TimeUnixMilli(f *testing.F) {
	f.Fuzz(func(t *testing.T, milli int64) {
		v := time.UnixMilli(milli)
		test.Test([]time.Time{v}, TimeUnixMilli, t)
		test.TestSkip([]time.Time{v}, TimeUnixMilli, t)
	})
}

func FuzzUnsafe_TimeUnixMicro(f *testing.F) {
	f.Fuzz(func(t *testing.T, micro int64) {
		v := time.UnixMicro(micro)
		test.Test([]time.Time{v}, TimeUnixMicro, t)
		test.TestSkip([]time.Time{v}, TimeUnixMicro, t)
	})
}

func FuzzUnsafe_TimeUnixNano(f *testing.F) {
	f.Fuzz(func(t *testing.T, nano int64) {
		v := time.Unix(0, nano)
		test.Test([]time.Time{v}, TimeUnixNano, t)
		test.TestSkip([]time.Time{v}, TimeUnixNano, t)
	})
}

func FuzzUnsafe_TimeUnixUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, sec int64) {
		v := time.Unix(sec, 0).UTC()
		test.Test([]time.Time{v}, TimeUnixUTC, t)
		test.TestSkip([]time.Time{v}, TimeUnixUTC, t)
	})
}

func FuzzUnsafe_TimeUnixMilliUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, milli int64) {
		v := time.UnixMilli(milli).UTC()
		test.Test([]time.Time{v}, TimeUnixMilliUTC, t)
		test.TestSkip([]time.Time{v}, TimeUnixMilliUTC, t)
	})
}

func FuzzUnsafe_TimeUnixMicroUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, micro int64) {
		v := time.UnixMicro(micro).UTC()
		test.Test([]time.Time{v}, TimeUnixMicroUTC, t)
		test.TestSkip([]time.Time{v}, TimeUnixMicroUTC, t)
	})
}

func FuzzUnsafe_TimeUnixNanoUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, nano int64) {
		v := time.Unix(0, nano).UTC()
		test.Test([]time.Time{v}, TimeUnixNanoUTC, t)
		test.TestSkip([]time.Time{v}, TimeUnixNanoUTC, t)
	})
}

func FuzzUnsafe_TimeUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		TimeUnixUTC.Unmarshal(buf)

		bufSkip := bytes.NewBuffer(bs)
		TimeUnixUTC.Skip(bufSkip)

		buf2 := bytes.NewBuffer(bs)
		TimeUnixMilliUTC.Unmarshal(buf2)

		bufSkip2 := bytes.NewBuffer(bs)
		TimeUnixMilliUTC.Skip(bufSkip2)

		buf3 := bytes.NewBuffer(bs)
		TimeUnixMicroUTC.Unmarshal(buf3)

		bufSkip3 := bytes.NewBuffer(bs)
		TimeUnixMicroUTC.Skip(bufSkip3)

		buf4 := bytes.NewBuffer(bs)
		TimeUnixNanoUTC.Unmarshal(buf4)

		bufSkip4 := bytes.NewBuffer(bs)
		TimeUnixNanoUTC.Skip(bufSkip4)

		buf5 := bytes.NewBuffer(bs)
		TimeUnix.Unmarshal(buf5)

		bufSkip5 := bytes.NewBuffer(bs)
		TimeUnix.Skip(bufSkip5)

		buf6 := bytes.NewBuffer(bs)
		TimeUnixMilli.Unmarshal(buf6)

		bufSkip6 := bytes.NewBuffer(bs)
		TimeUnixMilli.Skip(bufSkip6)

		buf7 := bytes.NewBuffer(bs)
		TimeUnixMicro.Unmarshal(buf7)

		bufSkip7 := bytes.NewBuffer(bs)
		TimeUnixMicro.Skip(bufSkip7)

		buf8 := bytes.NewBuffer(bs)
		TimeUnixNano.Unmarshal(buf8)

		bufSkip8 := bytes.NewBuffer(bs)
		TimeUnixNano.Skip(bufSkip8)
	})
}

// array -----------------------------------------------------------------------

func FuzzUnsafe_Array(f *testing.F) {
	f.Fuzz(func(t *testing.T, b1, b2, b3 byte) {
		v := [3]int{int(b1), int(b2), int(b3)}
		ser := NewArraySer[[3]int](varint.Int)
		test.Test([][3]int{v}, ser, t)
		test.TestSkip([][3]int{v}, ser, t)
	})
}

func FuzzUnsafe_ArrayUnmarshal(f *testing.F) {
	// Length validator for array is not needed as it has fixed length
	ser := NewArraySer[[3]int](varint.Int)
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		ser.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		ser.Skip(buf2)
	})
}
