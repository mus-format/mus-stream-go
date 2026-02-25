package testutil

import (
	"reflect"
	"testing"

	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testutil/mock"
)

func m[T comparable](wantV T, wantBs []byte, t *testing.T) mock.MarshalFn[T] {
	return func(v T, w muss.Writer) (n int, err error) {
		if v == wantV {
			return w.Write(wantBs)
		}
		t.Fatalf("ser.Marshal: unexepcted value, want %v actual %v", wantV, v)
		return
	}
}

func u[T any](wantBs []byte, wantV T, t *testing.T) mock.UnmarshalFn[T] {
	return func(r muss.Reader) (v T, n int, err error) {
		bs := make([]byte, len(wantBs))
		n, err = r.Read(bs)
		if err != nil {
			t.Fatalf("ser.Unmarshal: unexpected error, %v", err)
			return
		}
		if reflect.DeepEqual(bs, wantBs) {
			return wantV, n, nil
		}
		t.Fatalf("ser.Unmarshal: unexepcted bs, want '%v' actual '%v'", wantBs, bs)
		return
	}
}

func s[T comparable](wantV T, wantSize int, t *testing.T) mock.SizeFn[T] {
	return func(v T) (size int) {
		if v == wantV {
			return wantSize
		}
		t.Fatalf("ser.Size: unexepcted value, want %v actual %v", wantV, v)
		return
	}
}

func sk(wantBs []byte, t *testing.T) mock.SkipFn {
	return func(r muss.Reader) (n int, err error) {
		bs := make([]byte, len(wantBs))
		n, err = r.Read(bs)
		if err != nil {
			t.Fatalf("ser.Skip: unexpected error, %v", err)
			return
		}
		if reflect.DeepEqual(bs, wantBs) {
			return n, nil
		}
		t.Fatalf("ser.Skip: unexepcted bs, want '%v' actual '%v'",
			wantBs, bs)
		return
	}
}
