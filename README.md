# uuid7

Documentation:
https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-03

## Example

```go
package main

import (
    "log"

    "github.com/GoWebProd/uuid7"
)

func main() {
    u := uuid7.New()

    log.Println(u.Next().String())
}
```

## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/GoWebProd/uuid7
cpu: Intel Xeon Processor (Skylake, IBRS)
BenchmarkNext-8         18314782                65.94 ns/op            0 B/op          0 allocs/op
BenchmarkString-8       17735802                67.66 ns/op           48 B/op          1 allocs/op
BenchmarkParse-8        14948330                78.38 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/GoWebProd/uuid7      3.812s
```