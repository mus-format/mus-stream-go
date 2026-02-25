package unsafe

import (
	"bytes"
	"errors"
	"math"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	strops "github.com/mus-format/mus-stream-go/options/string"
	"github.com/mus-format/mus-stream-go/testutil"
	"github.com/mus-format/mus-stream-go/varint"
)

const maxLen = 1000

// bool ------------------------------------------------------------------------

func FuzzBool(f *testing.F) {
	seeds := []bool{true, false}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v bool) {
		testutil.Test[bool]([]bool{v}, Bool, t)
		testutil.TestSkip[bool]([]bool{v}, Bool, t)
	})
}

func FuzzBoolUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		Bool.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		Bool.Skip(buf2)
	})
}

// byte ------------------------------------------------------------------------

func FuzzByte(f *testing.F) {
	seeds := []byte{0, 1, 255}
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
	seeds := []uint64{0, 1, math.MaxUint64}
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
	seeds := []uint32{0, 1, math.MaxUint32}
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
	seeds := []uint16{0, 1, math.MaxUint16}
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
	seeds := []uint8{0, 1, 255}
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
	seeds := []uint{0, 1}
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
	seeds := []int64{0, 1, -1, math.MaxInt64, math.MinInt64}
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
	seeds := []int32{0, 1, -1, math.MaxInt32, math.MinInt32}
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
	seeds := []int16{0, 1, -1, math.MaxInt16, math.MinInt16}
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
	seeds := []int8{0, 1, -1, math.MaxInt8, math.MinInt8}
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
	seeds := []int{0, 1, -1}
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

// float64 ---------------------------------------------------------------------

func FuzzFloat64(f *testing.F) {
	seeds := []float64{0, 1, -1, 3.14, math.Inf(1), math.Inf(-1), math.NaN()}
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
	seeds := []float32{0, 1, -1, 3.14, float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())}
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

// string ----------------------------------------------------------------------

func FuzzString(f *testing.F) {
	seeds := []string{"", "hello", "world", "mus-format"}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v string) {
		testutil.Test[string]([]string{v}, String, t)
		testutil.TestSkip[string]([]string{v}, String, t)
	})
}

func FuzzStringUnmarshal(f *testing.F) {
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

func FuzzTimeUnix(f *testing.F) {
	f.Fuzz(func(t *testing.T, sec int64) {
		v := time.Unix(sec, 0)
		testutil.Test[time.Time]([]time.Time{v}, TimeUnix, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnix, t)
	})
}

func FuzzTimeUnixMilli(f *testing.F) {
	f.Fuzz(func(t *testing.T, milli int64) {
		v := time.UnixMilli(milli)
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixMilli, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixMilli, t)
	})
}

func FuzzTimeUnixMicro(f *testing.F) {
	f.Fuzz(func(t *testing.T, micro int64) {
		v := time.UnixMicro(micro)
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixMicro, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixMicro, t)
	})
}

func FuzzTimeUnixNano(f *testing.F) {
	f.Fuzz(func(t *testing.T, nano int64) {
		v := time.Unix(0, nano)
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixNano, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixNano, t)
	})
}

func FuzzTimeUnixUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, sec int64) {
		v := time.Unix(sec, 0).UTC()
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixUTC, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixUTC, t)
	})
}

func FuzzTimeUnixMilliUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, milli int64) {
		v := time.UnixMilli(milli).UTC()
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixMilliUTC, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixMilliUTC, t)
	})
}

func FuzzTimeUnixMicroUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, micro int64) {
		v := time.UnixMicro(micro).UTC()
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixMicroUTC, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixMicroUTC, t)
	})
}

func FuzzTimeUnixNanoUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, nano int64) {
		v := time.Unix(0, nano).UTC()
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixNanoUTC, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixNanoUTC, t)
	})
}

func FuzzTimeUnmarshal(f *testing.F) {
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

func FuzzArray(f *testing.F) {
	f.Fuzz(func(t *testing.T, b1, b2, b3 byte) {
		v := [3]int{int(b1), int(b2), int(b3)}
		ser := NewArraySer[[3]int, int](varint.Int)
		testutil.Test[[3]int]([][3]int{v}, ser, t)
		testutil.TestSkip[[3]int]([][3]int{v}, ser, t)
	})
}

func FuzzArrayUnmarshal(f *testing.F) {
	// Length validator for array is not needed as it has fixed length
	ser := NewArraySer[[3]int, int](varint.Int)
	f.Fuzz(func(t *testing.T, bs []byte) {
		buf := bytes.NewBuffer(bs)
		ser.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		ser.Skip(buf2)
	})
}
