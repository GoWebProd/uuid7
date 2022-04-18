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

    log.Println(u.Next())
}
```

## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/GoWebProd/uuid7
cpu: Intel Xeon Processor (Skylake, IBRS)
BenchmarkNext-8          9736629               125.1 ns/op            48 B/op          1 allocs/op
PASS
ok      github.com/GoWebProd/uuid7      1.353s
```