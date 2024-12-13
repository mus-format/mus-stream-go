package varint

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"golang.org/x/exp/constraints"
)

// MarshalUint64 writes an encoded (Varint) uint64 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint64(v uint64, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// MarshalUint32 writes an encoded (Varint) uint32 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint32(v uint32, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// MarshalUint16 writes an encoded (Varint) uint16 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint16(v uint16, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// MarshalUint8 writes an encoded (Varint) uint8 value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint8(v uint8, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// MarshalUint writes an encoded (Varint) uint value.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalUint(v uint, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// UnmarshalUint64 reads an encoded (Varint) uint64 value.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalUint64(r muss.Reader) (v uint64, n int, err error) {
	return unmarshalUint[uint64](com.Uint64MaxVarintLen,
		com.Uint64MaxLastByte,
		r)
}

// UnmarshalUint32 reads an encoded (Varint) uint32 value.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalUint32(r muss.Reader) (v uint32, n int, err error) {
	return unmarshalUint[uint32](com.Uint32MaxVarintLen,
		com.Uint32MaxLastByte,
		r)
}

// UnmarshalUint16 reads an encoded (Varint) uint16 value.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalUint16(r muss.Reader) (v uint16, n int, err error) {
	return unmarshalUint[uint16](com.Uint16MaxVarintLen,
		com.Uint16MaxLastByte,
		r)
}

// UnmarshalUint8 reads an encoded (Varint) uint8 value.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalUint8(r muss.Reader) (v uint8, n int, err error) {
	return unmarshalUint[uint8](com.Uint8MaxVarintLen,
		com.Uint8MaxLastByte,
		r)
}

// UnmarshalUint reads an encoded (Varint) uint value.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return com.ErrOverflow or a Reader error.
func UnmarshalUint(r muss.Reader) (v uint, n int, err error) {
	return unmarshalUint[uint](com.UintMaxVarintLen(),
		com.UintMaxLastByte(),
		r)
}

// SizeUint64 returns the size of an encoded (Varint) uint64 value.
func SizeUint64(v uint64) (size int) {
	return sizeUint(v)
}

// SizeUint32 returns the size of an encoded (Varint) uint32 value.
func SizeUint32(v uint32) (size int) {
	return sizeUint(v)
}

// SizeUint16 returns the size of an encoded (Varint) uint16 value.
func SizeUint16(v uint16) (size int) {
	return sizeUint(v)
}

// SizeUint8 returns the size of an encoded (Varint) uint8 value.
func SizeUint8(v uint8) (size int) {
	return sizeUint(v)
}

// SizeUint returns the size of an encoded (Varint) uint value.
func SizeUint(v uint) (size int) {
	return sizeUint(v)
}

// SkipUint64 skips an encoded (Varint) uint64 value.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipUint64(r muss.Reader) (n int, err error) {
	return skipUint(com.Uint64MaxVarintLen, com.Uint64MaxLastByte, r)
}

// SkipUint32 skips an encoded (Varint) uint32 value.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipUint32(r muss.Reader) (n int, err error) {
	return skipUint(com.Uint32MaxVarintLen, com.Uint32MaxLastByte, r)
}

// SkipUint16 skips an encoded (Varint) uint16 value.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipUint16(r muss.Reader) (n int, err error) {
	return skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte, r)
}

// SkipUint8 skips an encoded (Varint) uint8 value.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipUint8(r muss.Reader) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, r)
}

// SkipUint skips an encoded (Varint) uint value.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow
// or a Reader error.
func SkipUint(r muss.Reader) (n int, err error) {
	return skipUint(com.UintMaxVarintLen(), com.UintMaxLastByte(), r)
}

func marshalUint[T constraints.Unsigned](t T, w muss.Writer) (n int,
	err error) {
	for t >= 0x80 {
		err = w.WriteByte(byte(t) | 0x80)
		if err != nil {
			return
		}
		t >>= 7
		n++
	}
	err = w.WriteByte(byte(t))
	if err != nil {
		return
	}
	n++
	return
}

func unmarshalUint[T constraints.Unsigned](maxVarintLen int, maxLastByte byte,
	r muss.Reader) (t T, n int, err error) {
	var (
		b     byte
		shift int
	)
	for {
		b, err = r.ReadByte()
		if err != nil {
			return
		}
		n++
		if n == maxVarintLen && b > maxLastByte {
			return 0, n, com.ErrOverflow
		}
		if b < 0x80 {
			t = t | T(b)<<shift
			return
		}
		t = t | T(b&0x7F)<<shift
		shift += 7
	}
}

func sizeUint[T constraints.Unsigned](t T) (size int) {
	for t >= 0x80 {
		t >>= 7
		size++
	}
	return size + 1
}

func skipUint(maxVarintLen int, maxLastByte byte, r muss.Reader) (n int,
	err error) {
	var b byte
	for {
		b, err = r.ReadByte()
		if err != nil {
			return
		}
		if n == maxVarintLen && b > maxLastByte {
			return n, com.ErrOverflow
		}
		if b < 0x80 {
			n++
			return
		}
		n++
	}
}
