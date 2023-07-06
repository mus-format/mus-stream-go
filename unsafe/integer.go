package unsafe

import (
	"io"
	unsafe_mod "unsafe"

	muscom "github.com/mus-format/mus-common-go"
	muss "github.com/mus-format/mus-stream-go"
)

func marshalInteger64[T muscom.Integer64](t T, w muss.Writer) (n int,
	err error) {
	arr := *(*[muscom.Num64RawSize]byte)(unsafe_mod.Pointer(&t))
	return w.Write(arr[:])
}

func marshalInteger32[T muscom.Integer32](t T, w muss.Writer) (n int,
	err error) {
	arr := *(*[muscom.Num32RawSize]byte)(unsafe_mod.Pointer(&t))
	return w.Write(arr[:])
}

func marshalInteger16[T muscom.Integer16](t T, w muss.Writer) (n int,
	err error) {
	arr := *(*[muscom.Num16RawSize]byte)(unsafe_mod.Pointer(&t))
	return w.Write(arr[:])
}

func marshalInteger8[T muscom.Integer8](t T, w muss.Writer) (n int,
	err error) {
	arr := *(*[muscom.Num8RawSize]byte)(unsafe_mod.Pointer(&t))
	return w.Write(arr[:])
}

// -----------------------------------------------------------------------------
func unmarshalInteger64[T muscom.Integer64](r muss.Reader) (t T, n int,
	err error) {
	b := (*byte)(unsafe_mod.Pointer(&t))
	n, err = io.ReadFull(r, unsafe_mod.Slice(b, muscom.Num64RawSize))
	return
}

func unmarshalInteger32[T muscom.Integer32](r muss.Reader) (t T, n int,
	err error) {
	b := (*byte)(unsafe_mod.Pointer(&t))
	n, err = io.ReadFull(r, unsafe_mod.Slice(b, muscom.Num32RawSize))
	return
}

func unmarshalInteger16[T muscom.Integer16](r muss.Reader) (t T, n int,
	err error) {
	b := (*byte)(unsafe_mod.Pointer(&t))
	n, err = io.ReadFull(r, unsafe_mod.Slice(b, muscom.Num16RawSize))
	return
}

func unmarshalInteger8[T muscom.Integer8](r muss.Reader) (t T, n int,
	err error) {
	b := (*byte)(unsafe_mod.Pointer(&t))
	n, err = io.ReadFull(r, unsafe_mod.Slice(b, muscom.Num8RawSize))
	return
}
