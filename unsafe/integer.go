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
	bs := make([]byte, muscom.Num64RawSize)
	n, err = io.ReadFull(r, bs)
	if err != nil {
		return
	}
	t = *(*T)(unsafe_mod.Pointer(&bs[0]))
	return
}

func unmarshalInteger32[T muscom.Integer32](r muss.Reader) (t T, n int,
	err error) {
	bs := make([]byte, muscom.Num32RawSize)
	n, err = io.ReadFull(r, bs)
	if err != nil {
		return
	}
	t = *(*T)(unsafe_mod.Pointer(&bs[0]))
	return
}

func unmarshalInteger16[T muscom.Integer16](r muss.Reader) (t T, n int,
	err error) {
	bs := make([]byte, muscom.Num16RawSize)
	n, err = io.ReadFull(r, bs)
	if err != nil {
		return
	}
	t = *(*T)(unsafe_mod.Pointer(&bs[0]))
	return
}

func unmarshalInteger8[T muscom.Integer8](r muss.Reader) (t T, n int,
	err error) {
	bs := make([]byte, muscom.Num8RawSize)
	n, err = io.ReadFull(r, bs)
	if err != nil {
		return
	}
	t = *(*T)(unsafe_mod.Pointer(&bs[0]))
	return
}
