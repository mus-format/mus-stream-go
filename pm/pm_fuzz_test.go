package pm

import (
	"bytes"
	"testing"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testutil"
	"github.com/mus-format/mus-stream-go/varint"
)

func FuzzPtr(f *testing.F) {
	f.Add(1, 1, 2)
	f.Fuzz(func(t *testing.T, a1, a2, a3 int) {
		var (
			ptrMap                        = com.NewPtrMap()
			revPtrMap                     = com.NewReversePtrMap()
			baseSer   mus.Serializer[int] = varint.Int
			ptrSer                        = NewPtrSer[int](ptrMap, revPtrMap, baseSer)
			ser                           = Wrap[ctestutil.PtrStruct](ptrMap, revPtrMap, fuzzPtrStructSer{ptrSer})

			v = ctestutil.PtrStruct{A1: &a1, A2: &a1, A3: &a3}
		)
		testutil.Test[ctestutil.PtrStruct]([]ctestutil.PtrStruct{v}, ser, t)
		testutil.TestSkip[ctestutil.PtrStruct]([]ctestutil.PtrStruct{v}, ser, t)

		// Check pointer equality after Test
		buf := bytes.NewBuffer(nil)
		ser.Marshal(v, buf)
		v1, _, err := ser.Unmarshal(buf)
		if err != nil {
			t.Fatal(err)
		}

		// Check pointer equality
		if v1.A1 != v1.A2 {
			t.Errorf("v1.A1 and v1.A2 should be equal, actual %p and %p", v1.A1, v1.A2)
		}
		if v1.A1 == v1.A3 && a1 != a3 {
			t.Errorf("v1.A1 and v1.A3 should not be equal, actual %p and %p", v1.A1, v1.A3)
		}
		if *v1.A1 != a1 || *v1.A2 != a1 || *v1.A3 != a3 {
			t.Errorf("unexpected values, want (%v, %v, %v) actual (%v, %v, %v)", a1, a1, a3, *v1.A1, *v1.A2, *v1.A3)
		}
	})
}

func FuzzPtrUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		var (
			ptrMap                        = com.NewPtrMap()
			revPtrMap                     = com.NewReversePtrMap()
			baseSer   mus.Serializer[int] = varint.Int
			ptrSer                        = NewPtrSer[int](ptrMap, revPtrMap, baseSer)
			ser                           = Wrap[ctestutil.PtrStruct](ptrMap, revPtrMap, fuzzPtrStructSer{ptrSer})
		)
		buf := bytes.NewBuffer(bs)
		ser.Unmarshal(buf)

		buf2 := bytes.NewBuffer(bs)
		ser.Skip(buf2)
	})
}

func FuzzWrapper(f *testing.F) {
	f.Add(1, 2, 3)
	f.Fuzz(func(t *testing.T, a1, a2, a3 int) {
		var (
			ptrMap                        = com.NewPtrMap()
			revPtrMap                     = com.NewReversePtrMap()
			baseSer   mus.Serializer[int] = varint.Int
			ser                           = Wrap[ctestutil.PtrStruct](ptrMap, revPtrMap, fuzzPtrStructSer{NewPtrSer[int](ptrMap, revPtrMap, baseSer)})

			v = ctestutil.PtrStruct{A1: &a1, A2: &a1, A3: &a3}
		)
		testutil.Test[ctestutil.PtrStruct]([]ctestutil.PtrStruct{v}, ser, t)
		testutil.TestSkip[ctestutil.PtrStruct]([]ctestutil.PtrStruct{v}, ser, t)

		// Check pointer equality after Test (Test unmarshals into a new variable)
		// We'll do it manually to be sure.
		size := ser.Size(v)
		buf := bytes.NewBuffer(make([]byte, 0, size))
		ser.Marshal(v, buf)
		v1, _, err := ser.Unmarshal(buf)
		if err != nil {
			t.Fatal(err)
		}
		if v1.A1 != v1.A2 {
			t.Error("v1.A1 and v1.A2 should be equal")
		}
	})
}

// Rename to avoid conflict with wrapper_test.go
type fuzzPtrStructSer struct {
	ptrSer mus.Serializer[*int]
}

func (s fuzzPtrStructSer) Marshal(v ctestutil.PtrStruct, w mus.Writer) (n int,
	err error,
) {
	n, err = s.ptrSer.Marshal(v.A1, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.ptrSer.Marshal(v.A2, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = s.ptrSer.Marshal(v.A3, w)
	n += n1
	return
}

func (s fuzzPtrStructSer) Unmarshal(r mus.Reader) (v ctestutil.PtrStruct, n int,
	err error,
) {
	v.A1, n, err = s.ptrSer.Unmarshal(r)
	if err != nil {
		return
	}
	var n1 int
	v.A2, n1, err = s.ptrSer.Unmarshal(r)
	n += n1
	if err != nil {
		return
	}
	v.A3, n1, err = s.ptrSer.Unmarshal(r)
	n += n1
	return
}

func (s fuzzPtrStructSer) Size(v ctestutil.PtrStruct) (size int) {
	size = s.ptrSer.Size(v.A1)
	size += s.ptrSer.Size(v.A2)
	size += s.ptrSer.Size(v.A3)
	return
}

func (s fuzzPtrStructSer) Skip(r mus.Reader) (n int, err error) {
	n, err = s.ptrSer.Skip(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.ptrSer.Skip(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = s.ptrSer.Skip(r)
	n += n1
	return
}
