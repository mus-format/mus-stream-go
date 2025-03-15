package bslops

import (
	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

// Options for the byte slice serializer.
type Options struct {
	LenSer muss.Serializer[int]
	LenVl  com.Validator[int]
}

type SetOption func(o *Options)

func WithLenSer(lenSer muss.Serializer[int]) SetOption {
	return func(o *Options) { o.LenSer = lenSer }
}

func WithLenValidator(lenVl com.Validator[int]) SetOption {
	return func(o *Options) { o.LenVl = lenVl }
}

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
