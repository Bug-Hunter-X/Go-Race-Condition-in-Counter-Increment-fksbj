# Go Race Condition in Counter Increment

This repository demonstrates a common race condition bug in Go programs that use goroutines and mutexes to increment a shared counter.  The bug occurs because multiple goroutines attempt to update the counter concurrently without proper synchronization, leading to incorrect results.

The `bug.go` file contains the buggy code, while `bugSolution.go` demonstrates a corrected version using atomic operations for thread-safe counter updates.

## How to run

1. Clone the repository:
   ```bash
git clone https://github.com/<username>/go-race-condition.git
```
2. Navigate to the repository:
   ```bash
cd go-race-condition
```
3. Run the buggy code:
   ```bash
go run bug.go
```
4. Run the corrected code:
   ```bash
go run bugSolution.go
```

Observe the difference in the output. The buggy version may produce an incorrect count, whereas the corrected version consistently produces the correct count.