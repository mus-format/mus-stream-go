package unsafe

import (
	"io"
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
)

func marshalInteger64[T com.Integer64](t T, w mus.Writer) (n int,
	err error,
) {
	arr := *(*[com.Num64RawSize]byte)(unsafe_mod.Pointer(&t))
	return w.Write(arr[:])
}

func marshalInteger32[T com.Integer32](t T, w mus.Writer) (n int,
	err error,
) {
	arr := *(*[com.Num32RawSize]byte)(unsafe_mod.Pointer(&t))
	return w.Write(arr[:])
}

func marshalInteger16[T com.Integer16](t T, w mus.Writer) (n int,
	err error,
) {
	arr := *(*[com.Num16RawSize]byte)(unsafe_mod.Pointer(&t))
	return w.Write(arr[:])
}

func marshalInteger8[T com.Integer8](t T, w mus.Writer) (n int,
	err error,
) {
	arr := *(*[com.Num8RawSize]byte)(unsafe_mod.Pointer(&t))
	return w.Write(arr[:])
}

func unmarshalInteger64[T com.Integer64](r mus.Reader) (t T, n int,
	err error,
) {
	b := (*byte)(unsafe_mod.Pointer(&t))
	n, err = io.ReadFull(r, unsafe_mod.Slice(b, com.Num64RawSize))
	return
}

func unmarshalInteger32[T com.Integer32](r mus.Reader) (t T, n int,
	err error,
) {
	b := (*byte)(unsafe_mod.Pointer(&t))
	n, err = io.ReadFull(r, unsafe_mod.Slice(b, com.Num32RawSize))
	return
}

func unmarshalInteger16[T com.Integer16](r mus.Reader) (t T, n int,
	err error,
) {
	b := (*byte)(unsafe_mod.Pointer(&t))
	n, err = io.ReadFull(r, unsafe_mod.Slice(b, com.Num16RawSize))
	return
}

func unmarshalInteger8[T com.Integer8](r mus.Reader) (t T, n int,
	err error,
) {
	b := (*byte)(unsafe_mod.Pointer(&t))
	n, err = io.ReadFull(r, unsafe_mod.Slice(b, com.Num8RawSize))
	return
}
