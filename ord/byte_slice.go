package ord

import (
	"io"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalByteSlice writes an encoded slice value.
//
// The lenM argument specifies the Marshaller for the slice length, if nil,
// varint.MarshalPositiveInt() is used.
//
// In addition to the number of used bytes, it may also return a Writer error.
func MarshalByteSlice(v []byte, lenM muss.Marshaller[int], w muss.Writer) (
	n int, err error) {
	if lenM == nil {
		n, err = varint.MarshalPositiveInt(len(v), w)
	} else {
		n, err = lenM.Marshal(len(v), w)
	}
	if err != nil {
		return
	}
	n1, err := w.Write(v)
	n += n1
	return
}

// UnmarshalSlice reads an encoded slice value.
//
// The lenU argument specifies the Unmarshaller for the slice length, if nil,
// varint.UnmarshalPositiveInt() is used.
//
// In addition to the slice value and the number of used bytes, it may also
// return com.ErrOverflow, com.ErrNegativeLength or a Reader error.
func UnmarshalByteSlice(lenU muss.Unmarshaller[int], r muss.Reader) (v []byte,
	n int, err error) {
	return UnmarshalValidByteSlice(lenU, nil, false, r)
}

// UnmarshalValidByteSlice reads an encoded valid slice value.
//
// The lenU argument specifies the Unmarshaller for the slice length, if nil,
// varint.UnmarshalPositiveInt() is used.
// The lenVl argument specifies the slice length Validator. If it returns an
// error and skip == true, UnmarshalValidByteSlice skips the remaining bytes
// of the slice. If skip == false, a validation error is returned immediately.
//
// In addition to the slice value and the number of used bytes, it may also
// return com.ErrOverflow, com.ErrNegativeLength, a Validator or Reader error.
func UnmarshalValidByteSlice(lenU muss.Unmarshaller[int],
	lenVl com.Validator[int],
	skip bool,
	r muss.Reader,
) (v []byte, n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(r)
	} else {
		length, n, err = lenU.Unmarshal(r)
	}
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var (
		n1   int
		err1 error
	)
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make([]byte, length)
	n1, err = io.ReadFull(r, v)
	n += n1
	if err != nil {
		return
	}
	return
SkipRemainingBytes:
	if !skip {
		return
	}
	n1, err1 = skipRemainingByteSlice(length, r)
	n += n1
	if err1 != nil {
		err = err1
	}
	return
}

// SizeSlice returns the size of an encoded slice value.
//
// The lenS argument specifies the Sizer for the slice length, if nil,
// varint.SizePositiveInt() is used.
func SizeByteSlice(v []byte, lenS muss.Sizer[int]) (size int) {
	length := len(v)
	if lenS == nil {
		size = varint.SizePositiveInt(length)
	} else {
		size = lenS.Size(length)
	}
	return size + length
}

// SkipSlice skips an encoded slice value.
//
// The lenU argument specifies the Unmarshaller for the slice length, if nil,
// varint.UnmarshalPositiveInt() is used.
//
// In addition to the number of used bytes, it may also return com.ErrOverflow,
// com.ErrNegativeLength or a Reader error.
func SkipByteSlice(lenU muss.Unmarshaller[int], r muss.Reader) (
	n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(r)
	} else {
		length, n, err = lenU.Unmarshal(r)
	}
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	n1, err := skipRemainingByteSlice(length, r)
	n += n1
	return
}

func skipRemainingByteSlice(length int, r muss.Reader) (
	n int, err error) {
	for i := 0; i < length; i++ {
		_, err = r.ReadByte()
		if err != nil {
			return
		}
		n += 1
	}
	return
}
