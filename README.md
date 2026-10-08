# Go Concurrency Patterns Practice

This repository contains implementations of various concurrency patterns and problems in Go (Golang). Each pattern is implemented in its own directory with a clear, runnable example.

## 📁 Project Structure

```
machine-coding/
├── Fan-In-Fan-out/          # Fan-In / Fan-Out pattern
└── Worker-Pattern/          # Worker pool pattern
```

---

## 🚀 Patterns Implemented

### 1. Fan-In / Fan-Out (`Fan-In-Fan-out/`)

**Concept**: Distribute work across multiple workers (Fan-Out), then combine results back into a single channel (Fan-In).

**Key Features**:
- Multiple worker goroutines processing from a shared input channel
- Fan-in goroutine that merges multiple output channels into one global channel
- Uses `sync.WaitGroup` for synchronization
- Demonstrates pipeline pattern with channels

**Run**:
```bash
cd Fan-In-Fan-out && go run main.go
```

---

### 2. Thread-Safe Queue (`Thread-Safe-Queue/`)

**Concept**: A bounded, blocking queue implementation using condition variables (`sync.Cond`).

#### Correct Solution (`correct-solution/`)
- Properly handles queue closure with `isOpen` flag
- Uses `sync.Cond.Wait()` in loops to handle spurious wakeups
- Checks `isOpen` after wakeup before proceeding
- Uses `Broadcast()` on close to wake all waiters
- Includes producer/consumer demo with multiple goroutines

#### My Solution (`my-solution/`)
- Initial implementation with known issues (commented in code):
  - Array initialized with `make([]int, size)` instead of `make([]int, 0, size)` — pre-allocates zero values
  - Missing `isOpen` check after `cond.Wait()` returns (race condition on close)
  - No handling for spurious wakeups properly

**Run Correct Solution**:
```bash
cd Thread-Safe-Queue/correct-solution && go run main.go
```

---

### 3. Worker Pattern / Worker Pool (`Worker-Pattern/`)

**Concept**: Fixed number of workers processing jobs from a shared queue, with results collected via a separate channel.

**Key Features**:
- Fixed worker pool (5 workers)
- Jobs and results are structured types (`job`, `response`)
- Simulated variable work duration with `time.Sleep`
- Proper channel closing sequence: close input → wait workers → close output
- Uses two `WaitGroup`s: one for workers, one for result collector

**Run**:
```bash
cd Worker-Pattern && go run main.go
```

---

## 🎯 Learning Objectives

This repository demonstrates:

| Pattern | Go Concepts Used |
|---------|------------------|
| Fan-In/Fan-Out | Channels, goroutines, `sync.WaitGroup`, pipeline pattern |
| Thread-Safe Queue | `sync.Mutex`, `sync.Cond`, condition variables, blocking operations |
| Worker Pool | Channels, goroutines, `sync.WaitGroup`, job/result structs |

---

## 🛠 Prerequisites

- Go 1.18+ (uses generics in some patterns)

## 📦 Running All Examples

```bash
# Run each pattern
for dir in Fan-In-Fan-out Thread-Safe-Queue/correct-solution Worker-Pattern; do
  echo "=== Running $dir ==="
  (cd "$dir" && go run main.go)
  echo
done
```

---

## 📚 Resources & References

- [Go Concurrency Patterns](https://go.dev/blog/pipelines) - Official Go blog on pipelines
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
- [sync package documentation](https://pkg.go.dev/sync)
- [Go Memory Model](https://go.dev/ref/mem)

---

## 🤝 Contributing

This is a personal learning repository, but feel free to:
- Open issues for bugs or improvements
- Submit PRs with additional concurrency patterns
- Suggest better implementations

---

## 📄 License

MIT License - Feel free to use for learning purposes.