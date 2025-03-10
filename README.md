# mus-stream-go

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-stream-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-stream-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-stream-go)](https://goreportcard.com/report/github.com/mus-format/mus-stream-go)
[![codecov](https://codecov.io/gh/mus-format/mus-stream-go/graph/badge.svg?token=91OM0S4D9Q)](https://codecov.io/gh/mus-format/mus-stream-go)

mus-stream-go is a streaming version of [mus-go](https://github.com/mus-format/mus-go). 
It maintains the same structure but replaces byte slices with the `Writer` and 
`Reader` interfaces for streaming data.

# How To
You can learn more about this in the mus-go [documentation](https://github.com/mus-format/mus-go#how-to-use). 
Here is just a small example:
```go
package main

import "github.com/mus-format/mus-go/varint"

func main() {
		var (
			num  = 100
			size = varint.Int.Size(num)
			bs   = make([]byte, size)
			buf  = bytes.NewBuffer(bs) // Create a Writer/Reader.
		)
		n, err := varint.Int.Marshal(num, buf)
    // ...
    num, n, err = varint.Int.Unmarshal(buf)
    // ...
}
```

Another thing to note is that with a real connection (instead of `bytes.Buffer`), 
you need to use the `bufio` package. This is because `bufio.Writer` and 
`bufio.Reader` implement the `muss.Writer` and `muss.Reader` interfaces.

# DTM (Data Type Metadata) Support
[mus-stream-dts-go](https://github.com/mus-format/mus-stream-dts-go) provides [DTM](https://medium.com/p/21d7be309e8d) 
support.
