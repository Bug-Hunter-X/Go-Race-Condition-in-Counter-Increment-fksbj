```go
package main

import (

        "fmt"
        "sync"
        "sync/atomic"
)

func main() {
        var wg sync.WaitGroup
        
        var count int64

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        atomic.AddInt64(&count, int64(i))
                }(i)
        }

        wg.Wait()
        fmt.Println("Count:", count)
}
```