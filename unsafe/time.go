package unsafe

import (
	"time"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-stream-go"
)

var (
	// TimeUnix is a time.Time serializer that encodes a value as a Unix
	// timestamp in seconds.
	TimeUnix = timeUnixSer{}
	// TimeUnixMilli is a time.Time serializer that encodes a value as a Unix
	// timestamp in milliseconds.
	TimeUnixMilli = timeUnixMilliSer{}
	// TimeUnixMicro is a time.Time serializer that encodes a value as a Unix
	// timestamp in microseconds.
	TimeUnixMicro = timeUnixMicroSer{}
	// TimeUnixNano is a time.Time serializer that encodes a value as a Unix
	// timestamp in nanoseconds.
	TimeUnixNano = timeUnixNanoSer{}

	// TimeUnixUTC is a time.Time serializer that encodes a value as a Unix
	// timestamp in seconds. The deserialized value is always in UTC.
	TimeUnixUTC = timeUnixUTCSer{}
	// TimeUnixMilli is a time.Time serializer that encodes a value as a Unix
	// timestamp in milliseconds. The deserialized value is always in UTC.
	TimeUnixMilliUTC = timeUnixMilliUTCSer{}
	// TimeUnixMicroUTC is a time.Time serializer that encodes a value as a Unix
	// timestamp in microseconds. The deserialized value is always in UTC.
	TimeUnixMicroUTC = timeUnixMicroUTCSer{}
	// TimeUnixNanoUTC is a time.Time serializer that encodes a value as a Unix
	// timestamp in nanoseconds. The deserialized value is always in UTC.
	TimeUnixNanoUTC = timeUnixNanoUTCSer{}
)

// -----------------------------------------------------------------------------

type timeUnixSer struct{}

// Marshal writes an encoded time.Time value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s timeUnixSer) Marshal(v time.Time, w mus.Writer) (n int, err error) {
	return Int64.Marshal(v.Unix(), w)
}

// Unmarshal reads an encoded time.Time value.
//
// In addition to the time.Time value and the number of read bytes, it may also
// return a Reader error.
func (s timeUnixSer) Unmarshal(r mus.Reader) (v time.Time, n int, err error) {
	sec, n, err := Int64.Unmarshal(r)
	if err != nil {
		return
	}
	v = time.Unix(sec, 0)
	return
}

// Size returns the size of an encoded time.Time value.
func (s timeUnixSer) Size(v time.Time) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded time.Time value.
//
// In addition to the number of skipped bytes, it may also return a Reader error.
func (s timeUnixSer) Skip(r mus.Reader) (n int, err error) {
	return Int64.Skip(r)
}

// -----------------------------------------------------------------------------

type timeUnixMilliSer struct{}

// Marshal writes an encoded time.Time value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s timeUnixMilliSer) Marshal(v time.Time, w mus.Writer) (n int,
	err error,
) {
	return Int64.Marshal(v.UnixMilli(), w)
}

// Unmarshal reads an encoded time.Time value.
//
// In addition to the time.Time value and the number of read bytes, it may also
// return a Reader error.
func (s timeUnixMilliSer) Unmarshal(r mus.Reader) (v time.Time, n int,
	err error,
) {
	milli, n, err := Int64.Unmarshal(r)
	if err != nil {
		return
	}
	v = time.UnixMilli(milli)
	return
}

// Size returns the size of an encoded time.Time value.
func (s timeUnixMilliSer) Size(v time.Time) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded time.Time value.
//
// In addition to the number of skipped bytes, it may also return a Reader error.
func (s timeUnixMilliSer) Skip(r mus.Reader) (n int, err error) {
	return Int64.Skip(r)
}

// -----------------------------------------------------------------------------

type timeUnixMicroSer struct{}

// Marshal writes an encoded time.Time value.
//
// In addition to the number of bytes written, it may also return a Writer error.
func (s timeUnixMicroSer) Marshal(v time.Time, w mus.Writer) (n int,
	err error,
) {
	return Int64.Marshal(v.UnixMicro(), w)
}

// Unmarshal reads an encoded time.Time value.
//
// In addition to the time.Time value and the number of read bytes, it may also
// return a Reader error.
func (s timeUnixMicroSer) Unmarshal(r mus.Reader) (v time.Time, n int,
	err error,
) {
	micro, n, err := Int64.Unmarshal(r)
	if err != nil {
		return
	}
	v = time.UnixMicro(micro)
	return
}

// Size returns the size of an encoded time.Time value.
func (s timeUnixMicroSer) Size(v time.Time) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded time.Time value.
//
// In addition to the number of skipped bytes, it may also return a Reader error.
func (s timeUnixMicroSer) Skip(r mus.Reader) (n int, err error) {
	return Int64.Skip(r)
}

// -----------------------------------------------------------------------------

type timeUnixNanoSer struct{}

// Marshal writes an encoded time.Time value.
//
// In addition to the number of bytes written, it may also return a Writer error.
// The result will be unpredictable if v is the zero Time.
func (s timeUnixNanoSer) Marshal(v time.Time, w mus.Writer) (n int, err error) {
	return Int64.Marshal(v.UnixNano(), w)
}

// Unmarshal reads an encoded time.Time value.
//
// In addition to the time.Time value and the number of read bytes, it may also
// return a Reader error.
func (s timeUnixNanoSer) Unmarshal(r mus.Reader) (v time.Time, n int, err error) {
	nano, n, err := Int64.Unmarshal(r)
	if err != nil {
		return
	}
	v = time.Unix(0, nano)
	return
}

// Size returns the size of an encoded time.Time value. The result will be
// unpredictable if v is the zero Time.
func (s timeUnixNanoSer) Size(v time.Time) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded time.Time value.
//
// In addition to the number of skipped bytes, it may also return a Reader error.
func (s timeUnixNanoSer) Skip(r mus.Reader) (n int, err error) {
	return Int64.Skip(r)
}

// -----------------------------------------------------------------------------

type timeUnixUTCSer struct {
	timeUnixSer
}

// Unmarshal reads an encoded time.Time value.
//
// In addition to the time.Time value and the number of read bytes, it may also
// return a Reader error.
func (s timeUnixUTCSer) Unmarshal(r mus.Reader) (v time.Time, n int, err error) {
	v, n, err = s.timeUnixSer.Unmarshal(r)
	if err == nil {
		v = v.UTC()
	}
	return
}

// -----------------------------------------------------------------------------

type timeUnixMilliUTCSer struct {
	timeUnixMilliSer
}

// Unmarshal reads an encoded time.Time value.
//
// In addition to the time.Time value and the number of read bytes, it may also
// return a Reader error.
func (s timeUnixMilliUTCSer) Unmarshal(r mus.Reader) (v time.Time, n int, err error) {
	v, n, err = s.timeUnixMilliSer.Unmarshal(r)
	if err == nil {
		v = v.UTC()
	}
	return
}

// -----------------------------------------------------------------------------

type timeUnixMicroUTCSer struct {
	timeUnixMicroSer
}

// Unmarshal reads an encoded time.Time value.
//
// In addition to the time.Time value and the number of read bytes, it may also
// return a Reader error.
func (s timeUnixMicroUTCSer) Unmarshal(r mus.Reader) (v time.Time, n int, err error) {
	v, n, err = s.timeUnixMicroSer.Unmarshal(r)
	if err == nil {
		v = v.UTC()
	}
	return
}

// -----------------------------------------------------------------------------

type timeUnixNanoUTCSer struct {
	timeUnixNanoSer
}

// Unmarshal reads an encoded time.Time value.
//
// In addition to the time.Time value and the number of read bytes, it may also
// return a Reader error.
func (s timeUnixNanoUTCSer) Unmarshal(r mus.Reader) (v time.Time, n int, err error) {
	v, n, err = s.timeUnixNanoSer.Unmarshal(r)
	if err == nil {
		v = v.UTC()
	}
	return
}
