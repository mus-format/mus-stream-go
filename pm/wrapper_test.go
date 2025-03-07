package pm

import (
	"fmt"
	"testing"
	"unsafe"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata"
	mock "github.com/mus-format/mus-stream-go/testdata/mock"
	"github.com/ymz-ncnk/mok"
)

func TestWrapper(t *testing.T) {

	t.Run("wrapper serializer should work correctly",
		func(t *testing.T) {
			var (
				st, baseSer = testdata.PtrStructSerData(t)
				ptrMap      = com.NewPtrMap()
				revPtrMap   = com.NewReversePtrMap()
				ser         = Wrap(ptrMap, revPtrMap, newPtrStructSer(ptrMap, revPtrMap, baseSer))
			)
			testdata.Test[com_testdata.PtrStruct]([]com_testdata.PtrStruct{st}, ser, t)
			testdata.TestSkip[com_testdata.PtrStruct]([]com_testdata.PtrStruct{st}, ser, t)
		})

	t.Run("Marshal should call ser.Marshal and empty the ptrMap",
		func(t *testing.T) {
			var (
				wantV  byte = 1
				wantN  int  = 1
				ptrMap      = com.NewPtrMap()
				ptrSer      = mock.NewSerializer[byte]().RegisterMarshal(
					func(v byte, w muss.Writer) (n int, err error) {
						ptrMap.Put(unsafe.Pointer(&v))
						n = wantN
						err = w.WriteByte(v)
						return
					},
				)
				ser = Wrap[byte](ptrMap, nil, ptrSer)
				w   = mock.NewWriter().RegisterWriteByte(
					func(b byte) error {
						if b != wantV {
							return fmt.Errorf("unexpected v, want %v, actual %v", wantV, b)
						}
						return nil
					},
				)
				mocks = []*mok.Mock{ptrSer.Mock, w.Mock}
			)
			n, err := ser.Marshal(wantV, w)
			if n != wantN {
				t.Fatalf("unexpected n, want %v, actual %v", wantN, n)
			}
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			if ptrMap.Len() != 0 {
				t.Fatal("ptrMap should be empty")
			}
			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}

		})

	t.Run("Unmarshal should call ser.Unmarshal and empty the revPtrMap",
		func(t *testing.T) {
			var (
				wantV     byte = 1
				wantN     int  = 1
				revPtrMap      = com.NewReversePtrMap()
				ptrSer         = mock.NewSerializer[byte]().RegisterUnmarshal(
					func(r muss.Reader) (v byte, n int, err error) {
						v, err = r.ReadByte()
						if err != nil {
							return
						}
						if v != wantV {
							err = fmt.Errorf("unexpected v, want %v, actual %v", wantV, v)
							return
						}
						n = wantN
						revPtrMap.Put(1, unsafe.Pointer(&v))
						return
					},
				)
				ser = Wrap[byte](nil, revPtrMap, ptrSer)
				r   = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = wantV
						return
					},
				)
				mocks = []*mok.Mock{ptrSer.Mock, r.Mock}
			)
			v, n, err := ser.Unmarshal(r)
			if v != wantV {
				t.Fatalf("unexpected v, want %v, actual %v", wantV, v)
			}
			if n != wantN {
				t.Fatalf("unexpected n, want %v, actual %v", wantN, n)
			}
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			if revPtrMap.Len() != 0 {
				t.Fatal("revPtrMap should be empty")
			}
			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("Size should call ser.Size and empty the ptrMap",
		func(t *testing.T) {
			var (
				wantV    byte = 1
				wantSize int  = 1
				ptrMap        = com.NewPtrMap()
				ptrSer        = mock.NewSerializer[byte]().RegisterSize(
					func(v byte) (size int) {
						ptrMap.Put(unsafe.Pointer(&v))
						return wantSize
					},
				)
				ser   = Wrap[byte](ptrMap, nil, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			size := ser.Size(wantV)
			if size != wantSize {
				t.Fatalf("unexpected size, want %v, actual %v", wantSize, size)
			}
			if ptrMap.Len() != 0 {
				t.Fatal("ptrMap should be empty")
			}
			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("Skip should call ser.Skip and empty the revPtrMap",
		func(t *testing.T) {
			var (
				wantV     byte = 1
				wantN     int  = 1
				revPtrMap      = com.NewReversePtrMap()
				ptrSer         = mock.NewSerializer[byte]().RegisterSkip(
					func(r muss.Reader) (n int, err error) {
						v, err := r.ReadByte()
						if err != nil {
							return
						}
						if v != wantV {
							err = fmt.Errorf("unexpected v, want %v, actual %v", wantV, v)
							return
						}
						revPtrMap.Put(1, unsafe.Pointer(&v))
						n = wantN
						return
					},
				)
				ser = Wrap[byte](nil, revPtrMap, ptrSer)
				r   = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						b = wantV
						return
					},
				)
				mocks = []*mok.Mock{ptrSer.Mock, r.Mock}
			)
			n, err := ser.Skip(r)
			if n != wantN {
				t.Fatalf("unexpected n, want %v, actual %v", wantN, n)
			}
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			if revPtrMap.Len() != 0 {
				t.Fatal("revPtrMap should be empty")
			}
			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

}

func newPtrStructSer(ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap,
	baseSer muss.Serializer[int]) muss.Serializer[com_testdata.PtrStruct] {
	return ptrStructSer{NewPtrSer[int](ptrMap, revPtrMap, baseSer)}
}

type ptrStructSer struct {
	intPtrSer muss.Serializer[*int]
}

func (s ptrStructSer) Marshal(v com_testdata.PtrStruct, w muss.Writer) (n int, err error) {
	n, err = s.intPtrSer.Marshal(v.A1, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.intPtrSer.Marshal(v.A2, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = s.intPtrSer.Marshal(v.A3, w)
	n += n1
	return
}

func (s ptrStructSer) Unmarshal(r muss.Reader) (v com_testdata.PtrStruct, n int, err error) {
	v.A1, n, err = s.intPtrSer.Unmarshal(r)
	if err != nil {
		return
	}
	var n1 int
	v.A2, n1, err = s.intPtrSer.Unmarshal(r)
	n += n1
	if err != nil {
		return
	}
	v.A3, n1, err = s.intPtrSer.Unmarshal(r)
	n += n1
	return
}

func (s ptrStructSer) Size(v com_testdata.PtrStruct) (size int) {
	size = s.intPtrSer.Size(v.A1)
	size += s.intPtrSer.Size(v.A2)
	size += s.intPtrSer.Size(v.A3)

	return
}

func (s ptrStructSer) Skip(r muss.Reader) (n int, err error) {
	n, err = s.intPtrSer.Skip(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.intPtrSer.Skip(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = s.intPtrSer.Skip(r)
	n += n1
	return
}
