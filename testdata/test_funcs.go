package testdata

import (
	"bytes"
	"reflect"
	"testing"

	mustrm "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func Test[T any](cases []T, m mustrm.Marshaler[T], u mustrm.Unmarshaler[T],
	s mustrm.Sizer[T],
	t *testing.T,
) {
	var err error
	for i := 0; i < len(cases); i++ {
		var (
			size = s.SizeMUS(cases[i])
			buf  = bytes.NewBuffer(make([]byte, 0, size))
			n    int
			v    T
		)
		n, err = m.MarshalMUS(cases[i], buf)
		if err != nil {
			t.Fatal(err)
		}
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		v, n, err := u.UnmarshalMUS(buf)
		if err != nil {
			t.Fatal(err)
		}
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		if !reflect.DeepEqual(v, cases[i]) {
			t.Errorf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v)
		}
	}
}

func TestSkip[T any](cases []T, m mustrm.Marshaler[T], sk mustrm.Skipper,
	s mustrm.Sizer[T],
	t *testing.T,
) {
	for i := 0; i < len(cases); i++ {
		var (
			size = s.SizeMUS(cases[i])
			buf  = bytes.NewBuffer(make([]byte, 0, size))
		)
		m.MarshalMUS(cases[i], buf)
		n, err := sk.SkipMUS(buf)
		if err != nil {
			t.Fatal(err)
		}
		if n != buf.Cap() {
			t.Fatal("skipped not enough")
		}
	}
}

func TestMarshalResults(wantN, n int, wantErr, err error, mocks []*mok.Mock,
	t *testing.T) {
	// if !reflect.DeepEqual(v, wantV) {
	// 	t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
	// }
	if n != wantN {
		t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
	}
	if err != wantErr {
		t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
	}
	if info := mok.CheckCalls(mocks); len(info) > 0 {
		t.Error(info)
	}
}
