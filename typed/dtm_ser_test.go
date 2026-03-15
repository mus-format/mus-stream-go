package typed

import (
	"bytes"
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go/test/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestDTMSer(t *testing.T) {
	t.Run("Marshal, Unmarshal, Size, Skip should succeed", func(t *testing.T) {
		var (
			dtm  = com.DTM(10)
			size = DTMSer.Size(dtm)
			buf  = bytes.NewBuffer(make([]byte, 0, size))
		)
		n, err := DTMSer.Marshal(dtm, buf)
		assertfatal.EqualError(t, err, nil)
		asserterror.Equal(t, n, size)

		adtm, n, err := DTMSer.Unmarshal(buf)
		assertfatal.EqualError(t, err, nil)
		asserterror.Equal(t, n, size)
		asserterror.Equal(t, adtm, dtm)

		buf.Reset()

		DTMSer.Marshal(dtm, buf)
		n, err = DTMSer.Skip(buf)
		assertfatal.EqualError(t, err, nil)
		asserterror.Equal(t, n, size)
	})

	t.Run("If varint.UnmarshalInt fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
			)
			dtm, n, err := DTMSer.Unmarshal(r)
			asserterror.EqualError(t, err, wantErr)
			asserterror.Equal(t, dtm, com.DTM(0))
			asserterror.Equal(t, n, 0)
		})

	t.Run("If varint.SkipInt fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("read byte error")

				r = mock.NewReader().RegisterReadByte(
					func() (b byte, err error) {
						err = wantErr
						return
					},
				)
			)
			n, err := DTMSer.Skip(r)
			asserterror.EqualError(t, err, wantErr)
			asserterror.Equal(t, n, 0)
		})
}
