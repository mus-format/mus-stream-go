package raw

import (
	muscom "github.com/mus-format/mus-common-go"
	muss "github.com/mus-format/mus-stream-go"
)

func marshalInteger64[T muscom.Integer64](t T, w muss.Writer) (n int,
	err error) {
	err = w.WriteByte(byte(t))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 8))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 16))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 24))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 32))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 40))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 48))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 56))
	if err != nil {
		return
	}
	n++
	return
}

func marshalInteger32[T muscom.Integer32](t T, w muss.Writer) (n int,
	err error) {
	err = w.WriteByte(byte(t))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 8))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 16))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 24))
	if err != nil {
		return
	}
	n++
	return
}

func marshalInteger16[T muscom.Integer16](t T, w muss.Writer) (n int,
	err error) {
	err = w.WriteByte(byte(t))
	if err != nil {
		return
	}
	n++
	err = w.WriteByte(byte(t >> 8))
	if err != nil {
		return
	}
	n++
	return
}

func marshalInteger8[T muscom.Integer8](t T, w muss.Writer) (n int,
	err error) {
	err = w.WriteByte(byte(t))
	if err != nil {
		return
	}
	n++
	return
}

// -----------------------------------------------------------------------------
func unmarshalInteger64[T muscom.Integer64](r muss.Reader) (t T, n int,
	err error) {
	var b byte
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t = T(b)
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 8
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 16
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 24
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 32
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 40
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 48
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 56
	return
}

func unmarshalInteger32[T muscom.Integer32](r muss.Reader) (t T, n int,
	err error) {
	var b byte
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t = T(b)
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 8
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 16
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 24
	return
}

func unmarshalInteger16[T muscom.Integer16](r muss.Reader) (t T, n int,
	err error) {
	var b byte
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t = T(b)
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t |= T(b) << 8
	return
}

func unmarshalInteger8[T muscom.Integer8](r muss.Reader) (t T, n int,
	err error) {
	var b byte
	b, err = r.ReadByte()
	if err != nil {
		return
	}
	n++
	t = T(b)
	return
}

// -----------------------------------------------------------------------------
func sizeNum64[T muscom.Num64](t T) int { // Remove this
	return muscom.Num64RawSize
}

func sizeNum32[T muscom.Num32](t T) int {
	return muscom.Num32RawSize
}

func sizeInteger16[T muscom.Integer16](t T) int {
	return muscom.Num16RawSize
}

func sizeInteger8[T muscom.Integer8](t T) int {
	return muscom.Num8RawSize
}

// -----------------------------------------------------------------------------
func skipInteger64(r muss.Reader) (int, error) {
	return skipInteger(muscom.Num64RawSize, r)
}

func skipInteger32(r muss.Reader) (int, error) {
	return skipInteger(muscom.Num32RawSize, r)
}

func skipInteger16(r muss.Reader) (int, error) {
	return skipInteger(muscom.Num16RawSize, r)
}

func skipInteger8(r muss.Reader) (int, error) {
	return skipInteger(muscom.Num8RawSize, r)
}

func skipInteger(integerSize int, r muss.Reader) (n int, err error) {
	for i := 0; i < integerSize; i++ {
		_, err = r.ReadByte()
		if err != nil {
			return
		}
		n++
	}
	return
}
