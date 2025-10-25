package testdata

import (
	"bytes"
	"testing"

	muss "github.com/mus-format/mus-stream-go"
	mock "github.com/mus-format/mus-stream-go/testdata/mock"
)

func MapSerData(t *testing.T) (mp map[string]int, keySer mock.Serializer[string],
	valueSer mock.Serializer[int],
) {
	var (
		aBs = append([]byte{1}, []byte("a")...)
		bBs = append([]byte{1}, []byte("b")...)
		cBs = append([]byte{1}, []byte("c")...)

		oneBs   = []byte{1}
		twoBs   = []byte{2}
		threeBs = []byte{3}

		keySize = 2
		valSize = 1
	)
	mp = map[string]int{"a": 1, "b": 2, "c": 3}
	keySer = mock.NewSerializer[string]().RegisterMarshalN(6,
		func(v string, w muss.Writer) (n int, err error) {
			switch v {
			case "a":
				return w.Write(aBs)
			case "b":
				return w.Write(bBs)
			case "c":
				return w.Write(cBs)
			default:
				t.Fatalf("keySer.Marshal: unexepcted string %v", v)
				return
			}
		}).RegisterUnmarshalN(3,
		func(r muss.Reader) (v string, n int, err error) {
			bs := make([]byte, keySize)
			n, err = r.Read(bs)
			if err != nil {
				t.Fatalf("keySer.Unmarshal: unexpected error, %v", err)
				return
			}
			switch {
			case bytes.Equal(bs, aBs):
				return "a", n, nil
			case bytes.Equal(bs, bBs):
				return "b", n, nil
			case bytes.Equal(bs, cBs):
				return "c", n, nil
			default:
				t.Fatalf("keySer.Unmarshal: unexepcted bs '%v'", bs)
				return
			}
		}).RegisterSizeN(6,
		func(v string) (size int) {
			switch v {
			case "a":
				return keySize
			case "b":
				return keySize
			case "c":
				return keySize
			default:
				t.Fatalf("keySer.Size: unexepcted string %v", v)
			}
			return
		}).RegisterSkipN(3,
		func(r muss.Reader) (n int, err error) {
			bs := make([]byte, keySize)
			n, err = r.Read(bs)
			if err != nil {
				t.Fatalf("keySer.Skip: unexpected error, %v", err)
				return
			}
			switch {
			case bytes.Equal(bs, aBs):
				return n, nil
			case bytes.Equal(bs, bBs):
				return n, nil
			case bytes.Equal(bs, cBs):
				return n, nil
			default:
				t.Fatalf("keySer.Skip: unexepcted bs '%v'", bs)
			}
			return
		})

	valueSer = mock.NewSerializer[int]().RegisterMarshalN(6,
		func(v int, w muss.Writer) (n int, err error) {
			switch v {
			case 1:
				return w.Write(oneBs)
			case 2:
				return w.Write(twoBs)
			case 3:
				return w.Write(threeBs)
			default:
				t.Fatalf("valueSer.Marshal: unexepcted int %v", v)
			}
			return
		}).RegisterUnmarshalN(3,
		func(r muss.Reader) (v int, n int, err error) {
			bs := make([]byte, valSize)
			n, err = r.Read(bs)
			if err != nil {
				t.Fatalf("valueSer.Unmarshal: unexpected error, %v", err)
				return
			}
			switch {
			case bytes.Equal(bs, oneBs):
				return 1, n, nil
			case bytes.Equal(bs, twoBs):
				return 2, n, nil
			case bytes.Equal(bs, threeBs):
				return 3, n, nil
			default:
				t.Fatalf("valueSer.Unmarshal: unexepcted bs '%v'", bs)
				return
			}
		}).RegisterSizeN(6,
		func(v int) (size int) {
			switch v {
			case 1:
				return valSize
			case 2:
				return valSize
			case 3:
				return valSize
			default:
				t.Fatalf("valueSer.Size: unexepcted int %v", v)
			}
			return
		}).RegisterSkipN(3,
		func(r muss.Reader) (n int, err error) {
			bs := make([]byte, valSize)
			n, err = r.Read(bs)
			if err != nil {
				t.Fatalf("valueSer.Skip: unexpected error, %v", err)
				return
			}
			switch {
			case bytes.Equal(bs, oneBs):
				return n, nil
			case bytes.Equal(bs, twoBs):
				return n, nil
			case bytes.Equal(bs, threeBs):
				return n, nil
			default:
				t.Fatalf("valueSer.Skip: unexepcted bs '%v'", bs)
			}
			return
		})
	return
}

func MapLenSerData(t *testing.T) (mp map[string]int, lenSer mock.Serializer[int],
	keySer mock.Serializer[string], valueSer mock.Serializer[int],
) {
	mp, keySer, valueSer = MapSerData(t)
	var (
		l    = len(mp)
		lBs  = []byte{byte(l * 2)}
		size = 1
	)
	lenSer = mock.NewSerializer[int]().
		// unmarshal
		RegisterMarshal(m(l, lBs, t)).
		RegisterUnmarshal(u(lBs, l, t)).
		RegisterSize(s(l, size, t)).
		// skip
		RegisterMarshal(m(l, lBs, t)).
		RegisterUnmarshal(u(lBs, l, t)).
		RegisterSize(s(l, size, t))
	return
}
