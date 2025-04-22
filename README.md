# mus-stream-go

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-stream-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-stream-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-stream-go)](https://goreportcard.com/report/github.com/mus-format/mus-stream-go)
[![codecov](https://codecov.io/gh/mus-format/mus-stream-go/graph/badge.svg?token=91OM0S4D9Q)](https://codecov.io/gh/mus-format/mus-stream-go)

mus-stream-go offers a streaming version of [mus-go](https://github.com/mus-format/mus-go), 
keeping the same structure but using `Writer` and `Reader` interfaces instead of 
byte slices.

# How To
More information can be found in the mus-go [documentation](https://github.com/mus-format/mus-go#how-to-use). 
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

When working with real connections (e.g., network or file I/O) rather than 
`bytes.Buffer`, you must use `bufio.Writer` and `bufio.Reader`. This is 
required because:
1. They implement the `muss.Writer` and `muss.Reader` interfaces.
2. They provide the necessary buffering for efficient I/O operations.

# DTS (Data Type metadata Support) 
[dts-stream-go](https://github.com/mus-format/dts-stream-go) enables 
typed data serialization using [DTM](https://medium.com/p/21d7be309e8d).

# MarshallerMUS Interface
The `MarshallerMUS` interface is defined in the [ext-mus-stream-go](https://github.com/mus-format/ext-mus-stream-go)
module.
