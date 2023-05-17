package unsafe

import (
	"strconv"

	muscom "github.com/mus-format/mus-common-go"
	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/raw"
)

func init() {
	setUpUintFuncs(strconv.IntSize)
}

var (
	marshalUint   func(v uint, w mustrm.Writer) (int, error)
	unmarshalUint func(r mustrm.Reader) (uint, int, error)
	sizeUint      func(v uint) int
	skipUint      func(r mustrm.Reader) (int, error)
)

// MarshalUint64 writes the MUS encoding (Raw) of a uint64. Returns the number
// of used bytes and an error.
func MarshalUint64(v uint64, w mustrm.Writer) (n int, err error) {
	return marshalInteger64(v, w)
}

// MarshalUint32 writes the MUS encoding (Raw) of a uint32. Returns the number
// of used bytes and error.
func MarshalUint32(v uint32, w mustrm.Writer) (n int, err error) {
	return marshalInteger32(v, w)
}

// MarshalUint16 writes the MUS encoding (Raw) of a uint16. Returns the number
// of used bytes and error.
func MarshalUint16(v uint16, w mustrm.Writer) (n int, err error) {
	return marshalInteger16(v, w)
}

// MarshalUint8 writes the MUS encoding (Raw) of a uint8. Returns the number
// of used bytes and error.
func MarshalUint8(v uint8, w mustrm.Writer) (n int, err error) {
	return marshalInteger8(v, w)
}

// MarshalUint writes the MUS encoding (Raw) of a uint. Returns the number
// of used bytes and error.
func MarshalUint(v uint, w mustrm.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// -----------------------------------------------------------------------------
// UnmarshalUint64 reads a MUS-encoded (Raw) uint64. In addition to  uint64, it
// returns the number of used bytes and an error.
func UnmarshalUint64(r mustrm.Reader) (v uint64, n int, err error) {
	return unmarshalInteger64[uint64](r)
}

// UnmarshalUint32 reads a MUS-encoded (Raw) uint32. In addition to  uint32, it
// returns the number of used bytes and an error.
func UnmarshalUint32(r mustrm.Reader) (v uint32, n int, err error) {
	return unmarshalInteger32[uint32](r)
}

// UnmarshalUint16 reads a MUS-encoded (Raw) uint16. In addition to  uint16, it
// returns the number of used bytes and an error.
func UnmarshalUint16(r mustrm.Reader) (v uint16, n int, err error) {
	return unmarshalInteger16[uint16](r)
}

// UnmarshalUint8 reads a MUS-encoded (Raw) uint8. In addition to  uint8, it
// returns the number of used bytes and an error.
func UnmarshalUint8(r mustrm.Reader) (v uint8, n int, err error) {
	return unmarshalInteger8[uint8](r)
}

// UnmarshalUint reads a MUS-encoded (Raw) uint. In addition to  uint, it
// returns the number of used bytes and an error.
func UnmarshalUint(r mustrm.Reader) (v uint, n int, err error) {
	return unmarshalUint(r)
}

// -----------------------------------------------------------------------------
// SizeUint64 returns the size of a MUS-encoded (Raw) uint64.
func SizeUint64(v uint64) (n int) {
	return raw.SizeUint64(v)
}

// SizeUint32 returns the size of a MUS-encoded (Raw) uint32.
func SizeUint32(v uint32) (n int) {
	return raw.SizeUint32(v)
}

// SizeUint16 returns the size of a MUS-encoded (Raw) uint16.
func SizeUint16(v uint16) (n int) {
	return raw.SizeUint16(v)
}

// SizeUint8 returns the size of a MUS-encoded (Raw) uint8.
func SizeUint8(v uint8) (n int) {
	return raw.SizeUint8(v)
}

// SizeUint returns the size of a MUS-encoded (Raw) uint.
func SizeUint(v uint) (n int) {
	return sizeUint((v))
}

// -----------------------------------------------------------------------------
// SkipUint64 skips a MUS-encoded (Raw) uint64. Returns the number of skiped
// bytes and an error.
func SkipUint64(r mustrm.Reader) (n int, err error) {
	return raw.SkipUint64(r)
}

// SkipUint32 skips a MUS-encoded (Raw) uint32. Returns the number of skiped
// bytes and an error.
func SkipUint32(r mustrm.Reader) (n int, err error) {
	return raw.SkipUint32(r)
}

// SkipUint16 skips a MUS-encoded (Raw) uint16. Returns the number of skiped
// bytes and an error.
func SkipUint16(r mustrm.Reader) (n int, err error) {
	return raw.SkipUint16(r)
}

// SkipUint8 skips a MUS-encoded (Raw) uint8. Returns the number of skiped
// bytes and an error.
func SkipUint8(r mustrm.Reader) (n int, err error) {
	return raw.SkipUint8(r)
}

// SkipUint skips a MUS-encoded (Raw) uint. Returns the number of skiped
// bytes and an error.
func SkipUint(r mustrm.Reader) (n int, err error) {
	return skipUint(r)
}

func setUpUintFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalUint = marshalInteger64[uint]
		unmarshalUint = unmarshalInteger64[uint]
	case 32:
		marshalUint = marshalInteger32[uint]
		unmarshalUint = unmarshalInteger32[uint]
	default:
		panic(muscom.ErrUnsupportedIntSize)
	}
	sizeUint = raw.SizeUint
	skipUint = raw.SkipUint
}
