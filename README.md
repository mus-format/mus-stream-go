# mus-stream: A Streaming Binary Serialization Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-stream-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-stream-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-stream-go)](https://goreportcard.com/report/github.com/mus-format/mus-stream-go)
[![codecov](https://codecov.io/gh/mus-format/mus-stream-go/graph/badge.svg?token=91OM0S4D9Q)](https://codecov.io/gh/mus-format/mus-stream-go)

**mus-stream** is a streaming version of the [mus](https://github.com/mus-format/mus-go) 
serializer. It keeps the same structure but uses `mus.Writer` and `mus.Reader` 
interfaces instead of byte slices.

## Quick Start

Here is a small example:

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

For network or file-based I/O use `bufio.Writer` and `bufio.Reader`. This 
ensures compatibility with `mus.Writer` and `mus.Reader` interfaces while 
providing the buffering necessary for optimal performance.

Detailed documentation on core serialization concepts and supported types can 
be found in the [mus-go repository](https://github.com/mus-format/mus-go).

## Contributing & Security

We welcome contributions of all kinds! Please see [CONTRIBUTING.md](CONTRIBUTING.md) 
for details on how to get involved.

If you find a security vulnerability, please refer to our 
[Security Policy](SECURITY.md) for instructions on how to report it privately.

## Version Compatibility

For a complete list of compatible module versions, see [VERSIONS.md](VERSIONS.md).
