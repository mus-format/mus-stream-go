package varint

import (
	muscom "github.com/mus-format/mus-common-go"
	mustrm "github.com/mus-format/mus-stream-go"
	"golang.org/x/exp/constraints"
)

// MarshalUint64 fills bs with the MUS encoding (Varint) of a uint64. Returns
// the number of used bytes.
//
// It will panic if receives too small bs.
func MarshalUint64(v uint64, w mustrm.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// MarshalUint32 fills bs with the MUS encoding (Varint) of a uint32. Returns
// the number of used bytes.
//
// It will panic if receives too small bs.
func MarshalUint32(v uint32, w mustrm.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// MarshalUint16 fills bs with the MUS encoding (Varint) of a uint16. Returns
// the number of used bytes.
//
// It will panic if receives too small bs.
func MarshalUint16(v uint16, w mustrm.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// MarshalUint8 fills bs with the MUS encoding (Varint) of a uint8. Returns
// the number of used bytes.
//
// It will panic if receives too small bs.
func MarshalUint8(v uint8, w mustrm.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// MarshalUint fills bs with the MUS encoding (Varint) of a uint. Returns
// the number of used bytes.
//
// It will panic if receives too small bs.
func MarshalUint(v uint, w mustrm.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// -----------------------------------------------------------------------------
// UnmarshalUint64 parses a MUS-encoded (Varint) uint64 from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalUint64(r mustrm.Reader) (v uint64, n int, err error) {
	return unmarshalUint[uint64](muscom.Uint64MaxVarintLen,
		muscom.Uint64MaxLastByte,
		r)
}

// UnmarshalUint32 parses a MUS-encoded (Varint) uint32 from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalUint32(r mustrm.Reader) (v uint32, n int, err error) {
	return unmarshalUint[uint32](muscom.Uint32MaxVarintLen,
		muscom.Uint32MaxLastByte,
		r)
}

// UnmarshalUint16 parses a MUS-encoded (Varint) uint16 from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalUint16(r mustrm.Reader) (v uint16, n int, err error) {
	return unmarshalUint[uint16](muscom.Uint16MaxVarintLen,
		muscom.Uint16MaxLastByte,
		r)
}

// UnmarshalUint8 parses a MUS-encoded (Varint) uint8 from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalUint8(r mustrm.Reader) (v uint8, n int, err error) {
	return unmarshalUint[uint8](muscom.Uint8MaxVarintLen,
		muscom.Uint8MaxLastByte,
		r)
}

// UnmarshalUint parses a MUS-encoded (Varint) uint from bs. In addition to
// the byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalUint(r mustrm.Reader) (v uint, n int, err error) {
	return unmarshalUint[uint](muscom.UintMaxVarintLen(),
		muscom.UintMaxLastByte(),
		r)
}

// -----------------------------------------------------------------------------
// SizeUint64 returns the size of a MUS-encoded uint64.
func SizeUint64(v uint64) (size int) {
	return sizeUint(v)
}

// SizeUint32 returns the size of a MUS-encoded uint32.
func SizeUint32(v uint32) (size int) {
	return sizeUint(v)
}

// SizeUint16 returns the size of a MUS-encoded uint16.
func SizeUint16(v uint16) (size int) {
	return sizeUint(v)
}

// SizeUint8 returns the size of a MUS-encoded uint8.
func SizeUint8(v uint8) (size int) {
	return sizeUint(v)
}

// SizeUint returns the size of a MUS-encoded uint.
func SizeUint(v uint) (size int) {
	return sizeUint(v)
}

// -----------------------------------------------------------------------------
// SkipUint64 skips a MUS-encoded uint64 in bs. Returns the number of skiped
// bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipUint64(r mustrm.Reader) (n int, err error) {
	return skipUint(muscom.Uint64MaxVarintLen, muscom.Uint64MaxLastByte, r)
}

// SkipUint32 skips a MUS-encoded uint32 in bs. Returns the number of skiped
// bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipUint32(r mustrm.Reader) (n int, err error) {
	return skipUint(muscom.Uint32MaxVarintLen, muscom.Uint32MaxLastByte, r)
}

// SkipUint16 skips a MUS-encoded uint16 in bs. Returns the number of skiped
// bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipUint16(r mustrm.Reader) (n int, err error) {
	return skipUint(muscom.Uint16MaxVarintLen, muscom.Uint16MaxLastByte, r)
}

// SkipUint8 skips a MUS-encoded uint8 in bs. Returns the number of skiped
// bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipUint8(r mustrm.Reader) (n int, err error) {
	return skipUint(muscom.Uint8MaxVarintLen, muscom.Uint8MaxLastByte, r)
}

// SkipUint skips a MUS-encoded uint in bs. Returns the number of skiped
// bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipUint(r mustrm.Reader) (n int, err error) {
	return skipUint(muscom.UintMaxVarintLen(), muscom.UintMaxLastByte(), r)
}

func marshalUint[T constraints.Unsigned](t T, w mustrm.Writer) (n int,
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
	r mustrm.Reader) (t T, n int, err error) {
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
			return 0, n, muscom.ErrOverflow
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

func skipUint(maxVarintLen int, maxLastByte byte, r mustrm.Reader) (n int,
	err error) {
	var b byte
	for {
		b, err = r.ReadByte()
		if err != nil {
			return
		}
		if n == maxVarintLen && b > maxLastByte {
			return n, muscom.ErrOverflow
		}
		if b < 0x80 {
			n++
			return
		}
		n++
	}
}
