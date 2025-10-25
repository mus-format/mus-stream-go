package varint

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
	"golang.org/x/exp/constraints"
)

var (
	// Uint64 is a uint64 serializer.
	Uint64 = uint64Ser{}
	// Uint32 is a uint32 serializer.
	Uint32 = uint32Ser{}
	// Uint16 is a uint16 serializer.
	Uint16 = uint16Ser{}
	// Uint8 is a uint8 serializer.
	Uint8 = uint8Ser{}
	// Uint is a uint serializer.
	Uint = uintSer{}
)

type uint64Ser struct{}

// Marshal writes an encoded (Varint) uint64 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (uint64Ser) Marshal(v uint64, w mus.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// Unmarshal reads an encoded (Varint) uint64 value.
//
// In addition to the uint64 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (uint64Ser) Unmarshal(r mus.Reader) (v uint64, n int, err error) {
	return unmarshalUint[uint64](com.Uint64MaxVarintLen,
		com.Uint64MaxLastByte,
		r)
}

// Size returns the size of an encoded (Varint) uint64 value.
func (uint64Ser) Size(v uint64) int {
	return sizeUint(v)
}

// Skip skips an encoded (Varint) uint64 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (uint64Ser) Skip(r mus.Reader) (n int, err error) {
	return skipUint(com.Uint64MaxVarintLen, com.Uint64MaxLastByte, r)
}

// -----------------------------------------------------------------------------

type uint32Ser struct{}

// Marshal writes an encoded (Varint) uint32 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (uint32Ser) Marshal(v uint32, w mus.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// Unmarshal reads an encoded (Varint) uint32 value.
//
// In addition to the uint32 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (uint32Ser) Unmarshal(r mus.Reader) (v uint32, n int, err error) {
	return unmarshalUint[uint32](com.Uint32MaxVarintLen,
		com.Uint32MaxLastByte,
		r)
}

// Size returns the size of an encoded (Varint) uint32 value.
func (uint32Ser) Size(v uint32) int {
	return sizeUint(v)
}

// Skip skips an encoded (Varint) uint32 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (uint32Ser) Skip(r mus.Reader) (n int, err error) {
	return skipUint(com.Uint32MaxVarintLen, com.Uint32MaxLastByte, r)
}

// -----------------------------------------------------------------------------

type uint16Ser struct{}

// Marshal writes an encoded (Varint) uint16 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (uint16Ser) Marshal(v uint16, w mus.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// Unmarshal reads an encoded (Varint) uint16 value.
//
// In addition to the uint16 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (uint16Ser) Unmarshal(r mus.Reader) (v uint16, n int, err error) {
	return unmarshalUint[uint16](com.Uint16MaxVarintLen,
		com.Uint16MaxLastByte,
		r)
}

// Size returns the size of an encoded (Varint) uint16 value.
func (uint16Ser) Size(v uint16) int {
	return sizeUint(v)
}

// Skip skips an encoded (Varint) uint16 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (uint16Ser) Skip(r mus.Reader) (n int, err error) {
	return skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte, r)
}

// -----------------------------------------------------------------------------

type uint8Ser struct{}

// Marshal writes an encoded (Varint) uint8 value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (uint8Ser) Marshal(v uint8, w mus.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// Unmarshal reads an encoded (Varint) uint8 value.
//
// In addition to the uint8 value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (uint8Ser) Unmarshal(r mus.Reader) (v uint8, n int, err error) {
	return unmarshalUint[uint8](com.Uint8MaxVarintLen,
		com.Uint8MaxLastByte,
		r)
}

// Size returns the size of an encoded (Varint) uint8 value.
func (uint8Ser) Size(v uint8) int {
	return sizeUint(v)
}

// Skip skips an encoded (Varint) uint8 value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (uint8Ser) Skip(r mus.Reader) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, r)
}

// -----------------------------------------------------------------------------

type uintSer struct{}

// Marshal writes an encoded (Varint) uint value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (uintSer) Marshal(v uint, w mus.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// Unmarshal reads an encoded (Varint) uint value.
//
// In addition to the uint value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (uintSer) Unmarshal(r mus.Reader) (v uint, n int, err error) {
	return unmarshalUint[uint](com.UintMaxVarintLen(),
		com.UintMaxLastByte(),
		r)
}

// Size returns the size of an encoded (Varint) uint value.
func (uintSer) Size(v uint) int {
	return sizeUint(v)
}

// Skip skips an encoded (Varint) uint value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (uintSer) Skip(r mus.Reader) (n int, err error) {
	return skipUint(com.UintMaxVarintLen(), com.UintMaxLastByte(), r)
}

func marshalUint[T constraints.Unsigned](t T, w mus.Writer) (n int,
	err error,
) {
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
	r mus.Reader,
) (t T, n int, err error) {
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

func skipUint(maxVarintLen int, maxLastByte byte, r mus.Reader) (n int,
	err error,
) {
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
