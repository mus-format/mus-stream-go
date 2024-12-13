# mus-stream-go

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-stream-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-stream-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-stream-go)](https://goreportcard.com/report/github.com/mus-format/mus-stream-go)
[![codecov](https://codecov.io/gh/mus-format/mus-stream-go/graph/badge.svg?token=91OM0S4D9Q)](https://codecov.io/gh/mus-format/mus-stream-go)

mus-stream-go is a streaming version of [mus-go](https://github.com/mus-format/mus-go). 
It has the same structure, but uses the `Writer` and `Reader` interfaces instead
of a byte slices.

# How To Use
You can learn more about this in the mus-go [documentation](https://github.com/mus-format/mus-go#how-to-use). 
Here is just a small example.

mus-stream-go is able to skip invalid data from the byte stream:
```go
package main

import (
  "bytes"
  "errors"
  "fmt"

  com "github.com/mus-format/common-go"
  "github.com/mus-format/mus-stream-go/ord"
)

func main() {
  var (
    str1                              = "hello"
    str2                              = "very long string" // Invalid string.
    str3                              = "world"
    // String length validator.
    maxLength com.ValidatorFn[int] = func(length int) error {
      if length > 5 {
        return errors.New("too long")
      }
      return nil
    }
    size = ord.SizeString(str1) + ord.SizeString(str2) + ord.SizeString(str3)
    bs   = make([]byte, 0, size)
    buf  = bytes.NewBuffer(bs) // Create a Writer/Reader.
  )

  // Fill the buffer.
  ord.MarshalString(str1, buf)
  ord.MarshalString(str2, buf)
  ord.MarshalString(str3, buf)

  var (
    skip = true // The invalid string will be skipped. If false, the Unmarshal 
    // function will immediately return a validation error.
    str  string
    err  error
  )
  for i := 0; i < 3; i++ {
    str, _, err = ord.UnmarshalValidString(maxLength, skip, buf)
    if err == nil {
      fmt.Println(str)
    } else {
      // In this case the string is skipped.
      fmt.Printf("validation error - \"%v\"\n", err)
    }
  }
  // The output will be:
  // hello
  // validation error - "too long"
  // world
}
```
Another thing to note is that with a real connection (instead of `bytes.Buffer`), 
we need to use the `bufio` package. This is because `bufio.Writer` and 
`bufio.Reader` implement the `muss.Writer` and `muss.Reader` interfaces.

# Data Type Metadata (DTM) Support
[mus-stream-dts-go](https://github.com/mus-format/mus-stream-dts-go) provides 
DTM support.
