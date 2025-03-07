package varint

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// Byte is a byte serializer.
var Byte = byteSer{}

type byteSer struct{}

// Marshal writes an encoded (Varint) byte value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (byteSer) Marshal(v byte, w muss.Writer) (n int, err error) {
	return marshalUint(v, w)
}

// Unmarshal reads an encoded (Varint) byte value.
//
// In addition to the byte value and the number of bytes read, it may also
// return com.ErrOverflow or a Reader error.
func (byteSer) Unmarshal(r muss.Reader) (v byte, n int, err error) {
	return unmarshalUint[byte](com.Uint8MaxVarintLen, com.Uint8MaxLastByte,
		r)
}

// Size returns the size of an encoded (Varint) byte value.
func (byteSer) Size(v byte) (size int) {
	return sizeUint(v)
}

// Skip skips an encoded (Varint) byte value.
//
// In addition to the number of bytes read, it may also return com.ErrOverflow
// or a Reader error.
func (byteSer) Skip(r muss.Reader) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, r)
}
