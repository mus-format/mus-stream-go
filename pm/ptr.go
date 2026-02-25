package pm

import (
	"unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// NewPtrSer returns a new pointer serializer with the given pointer map,
// reverse pointer map and base serializer.
func NewPtrSer[T any](ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap,
	baseSer mus.Serializer[T],
) ptrSer[T] {
	return ptrSer[T]{ptrMap, revPtrMap, baseSer}
}

type ptrSer[T any] struct {
	ptrMap    *com.PtrMap
	revPtrMap *com.ReversePtrMap
	baseSer   mus.Serializer[T]
}

// Marshal writes an encoded pointer.
//
// In addition to the number of bytes written, it may also return a base type
// marshalling error, or a Writer error.
func (s ptrSer[T]) Marshal(v *T, w mus.Writer) (n int, err error) {
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
	id, newOne := maptr(unsafe.Pointer(v), s.ptrMap)
	var n1 int
	n1, err = varint.PositiveInt.Marshal(id, w)
	n += n1
	if err != nil {
		return
	}
	if newOne {
		n1, err = s.baseSer.Marshal(*v, w)
		n += n1
	}
	return
}

// Unmarshal parses a MUS-encoded pointer from bs.
//
// In addition to the pointer and the number of bytes read, it may also return
// mus.ErrTooSmallByteSlice, com.ErrWrongFormat, a base type unmarshalling error,
// or a Reader error.
func (s ptrSer[T]) Unmarshal(r mus.Reader) (v *T, n int, err error) {
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
		id, n1, err = varint.PositiveInt.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		ptr, _ := s.revPtrMap.Get(id)
		if ptr == nil {
			v, n1, err = s.unmarshalData(id, r)
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
func (s ptrSer[T]) Size(v *T) (size int) {
	size = 1
	if v != nil {
		id, newOne := maptr(unsafe.Pointer(v), s.ptrMap)
		size += varint.PositiveInt.Size(id)
		if newOne {
			return size + s.baseSer.Size(*v)
		} else {
			return
		}
	}
	return
}

// SkipPtr skips an encoded pointer.
//
// In addition to the number of bytes read, it may also return
// mus.ErrTooSmallByteSlice, com.ErrWrongFormat, a base type skipping error,
// or a Reader error.
func (s ptrSer[T]) Skip(r mus.Reader) (n int,
	err error,
) {
	b, err := r.ReadByte()
	if err != nil {
		return
	}
	switch b {
	case byte(com.Nil):
		n = 1
		return
	case byte(com.Mapping):
		n = 1
		var (
			id int
			n1 int
		)
		id, n1, err = varint.PositiveInt.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		_, pst := s.revPtrMap.Get(id)
		if !pst {
			s.revPtrMap.Put(id, nil)
			n1, err = s.baseSer.Skip(r)
			n += n1
			if err != nil {
				return
			}
		}
	default:
		err = com.ErrWrongFormat
	}
	return
}

func (s ptrSer[T]) unmarshalData(id int, r mus.Reader) (v *T, n int, err error) {
	var (
		k  T
		n1 int
	)
	s.revPtrMap.Put(id, unsafe.Pointer(&k))
	k, n1, err = s.baseSer.Unmarshal(r)
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
