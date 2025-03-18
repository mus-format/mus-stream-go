package testdata

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

func Test[T any](cases []T, ser muss.Serializer[T], t *testing.T) {
	var err error
	for i := 0; i < len(cases); i++ {
		var (
			size = ser.Size(cases[i])
			buf  = bytes.NewBuffer(make([]byte, 0, size))
			n    int
			v    T
		)
		n, err = ser.Marshal(cases[i], buf)
		if err != nil {
			t.Fatal(err)
		}
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		v, n, err := ser.Unmarshal(buf)
		if err != nil {
			t.Fatal(err)
		}
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		if tm, ok := any(v).(time.Time); ok {
			tm1 := any(cases[i]).(time.Time)
			if !tm.Equal(tm1) {
				t.Errorf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v)
			}
		} else if !reflect.DeepEqual(v, cases[i]) {
			t.Errorf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v)
		}
	}
}

func TestSkip[T any](cases []T, ser muss.Serializer[T], t *testing.T) {
	for i := 0; i < len(cases); i++ {
		var (
			size = ser.Size(cases[i])
			buf  = bytes.NewBuffer(make([]byte, 0, size))
		)
		ser.Marshal(cases[i], buf)
		n, err := ser.Skip(buf)
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
	if n != wantN {
		t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
	}
	if err != wantErr {
		t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
	}
	if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
		t.Error(infomap)
	}
}
