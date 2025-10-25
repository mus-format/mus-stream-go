package pm

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
)

// Wrap function wraps the serializer that uses one or more pm pointer
// serializers (all created with the same pointer and reverse pointer maps), so
// it can be used like a regular serializer.
func Wrap[T any](ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap,
	ser mus.Serializer[T],
) wrapper[T] {
	return wrapper[T]{ptrMap, revPtrMap, ser}
}

type wrapper[T any] struct {
	ptrMap    *com.PtrMap
	revPtrMap *com.ReversePtrMap
	ser       mus.Serializer[T]
}

// Marshal writes an encoded pointer.
//
// In addition to the number of bytes written, it may also return an inner
// serializer marshalling error.
func (p wrapper[T]) Marshal(v T, w mus.Writer) (n int, err error) {
	defer func() {
		*p.ptrMap = *com.NewPtrMap()
	}()
	return p.ser.Marshal(v, w)
}

// Unmarshal reads an encoded pointer.
//
// In addition to the pointer and the number of bytes read, it may also return
// an inner serializer unmarshalling error.
func (p wrapper[T]) Unmarshal(r mus.Reader) (t T, n int, err error) {
	defer func() {
		*p.revPtrMap = *com.NewReversePtrMap()
	}()
	return p.ser.Unmarshal(r)
}

// Size returns the size of an encoded pointer.
func (p wrapper[T]) Size(v T) int {
	defer func() {
		*p.ptrMap = *com.NewPtrMap()
	}()
	return p.ser.Size(v)
}

// Skip skips an encoded pointer.
//
// In addition to the number of bytes read, it may also return an inner
// serializer error.
func (p wrapper[T]) Skip(r mus.Reader) (n int, err error) {
	defer func() {
		*p.revPtrMap = *com.NewReversePtrMap()
	}()
	return p.ser.Skip(r)
}
