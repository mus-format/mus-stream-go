package dts

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/varint"
)

// DTMSer serializes DTM values. It implements the mus.Serializer[com.DTM]
// interface.
var DTMSer = dtmSer{}

type dtmSer struct{}

func (s dtmSer) Marshal(dtm com.DTM, w mus.Writer) (n int, err error) {
	return varint.PositiveInt.Marshal(int(dtm), w)
}

func (s dtmSer) Unmarshal(r mus.Reader) (dtm com.DTM, n int, err error) {
	num, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	dtm = com.DTM(num)
	return
}

func (s dtmSer) Size(dtm com.DTM) (size int) {
	return varint.PositiveInt.Size(int(dtm))
}

func (s dtmSer) Skip(r mus.Reader) (n int, err error) {
	return varint.PositiveInt.Skip(r)
}
