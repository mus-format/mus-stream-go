package pm

import (
	"unsafe"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// MarshalPtr fills bs with the encoding of a pointer.
//
// The m argument specifies the Marshaller for the pointer base type.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalPtr[T any](v *T, m muss.Marshaller[T], mp *com.PtrMap,
	w muss.Writer,
) (n int, err error) {
	if v == nil {
		err = w.WriteByte(byte(com.Nil))
		if err != nil {
			return
		}
		n++
		return
	}
	err = w.WriteByte(byte(com.Mapping))
	if err != nil {
		return
	}
	n++
	id, newOne := maptr(unsafe.Pointer(v), mp)
	var n1 int
	n1, err = varint.MarshalInt(id, w)
	n += n1
	if err != nil {
		return
	}
	if newOne {
		n1, err = m.Marshal(*v, w)
		n += n1
	}
	return
}

// UnmarshalPtr parses a MUS-encoded pointer from bs.
//
// The u argument specifies the Unmarshaller for the base pointer type.
//
// In addition to the pointer, returns the number of used bytes and one of the
// mus.ErrTooSmallByteSlice, com.ErrWrongFormat or Unarshaller errors.
func UnmarshalPtr[T any](u muss.Unmarshaller[T], mp com.ReversePtrMap,
	r muss.Reader,
) (v *T, n int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	switch b {
	case byte(com.Nil):
		n = 1
		return
	case byte(com.Mapping):
		var (
			n1 int
			id int
		)
		n = 1
		id, n1, err = varint.UnmarshalInt(r)
		n += n1
		if err != nil {
			return
		}
		ptr, _ := mp.Get(id)
		if ptr == nil {
			v, n1, err = unmarshalData[T](id, u, mp, r)
			n += n1
		} else {
			v = (*T)(ptr)
		}
	default:
		err = com.ErrWrongFormat
	}
	return
}

// SizePtr returns the size of an encoded pointer.
//
// The s argument specifies the Sizer for the pointer base type.
func SizePtr[T any](v *T, s muss.Sizer[T], mp *com.PtrMap) (size int) {
	size = 1
	if v != nil {
		id, newOne := maptr(unsafe.Pointer(v), mp)
		size += varint.SizeInt(id)
		if newOne {
			return size + s.Size(*v)
		} else {
			return
		}
	}
	return
}

// SkipPtr skips an encoded pointer.
//
// The sk argument specifies the Skipper for the pointer base type.
//
// Returns the number of skiped bytes and one of the mus.ErrTooSmallByteSlice,
// com.ErrWrongFormat or Skipper errors.
func SkipPtr(sk muss.Skipper, mp com.ReversePtrMap, r muss.Reader) (n int,
	err error) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	switch b {
	case byte(com.Nil):
		n = 1
	case byte(com.Mapping):
		n = 1
		var (
			id int
			n1 int
		)
		id, n1, err = varint.UnmarshalInt(r)
		n += n1
		if err != nil {
			return
		}
		_, pst := mp.Get(id)
		if !pst {
			n1, err = sk.Skip(r)
			n += n1
			if err != nil {
				return
			}
			mp.Put(id, nil)
		}
	default:
		err = com.ErrWrongFormat
	}
	return
}

func unmarshalData[T any](id int, u muss.Unmarshaller[T],
	mp com.ReversePtrMap,
	r muss.Reader,
) (v *T, n int, err error) {
	var (
		k  T
		n1 int
	)
	mp.Put(id, unsafe.Pointer(&k))
	k, n1, err = u.Unmarshal(r)
	n += n1
	if err != nil {
		return
	}
	v = &k
	return
}

func maptr(ptr unsafe.Pointer, mp *com.PtrMap) (id int, newOne bool) {
	id, pst := mp.Get(ptr)
	if !pst {
		id = mp.Put(ptr)
		newOne = true
	}
	return
}
