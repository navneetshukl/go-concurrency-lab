### Context Cancellation — Requirements

1. Implement a **context-aware task processor** in Go that supports cancellation.

2. Create a `Job` structure containing:
   - `ID`
   - `Duration`

3. Create a `Result` structure containing:
   - `JobID`
   - `Status`

4. The task processor should start **3 worker goroutines** and process **20 jobs**.

5. Each worker should continuously receive a job, simulate processing using a timer, and report whether the job completed or was cancelled.

6. Use a `context.Context` to signal cancellation to all workers.

7. If cancellation occurs while a worker is processing a job, that worker must stop processing the job promptly.

8. Workers must not start processing new jobs after they observe cancellation.

9. The main goroutine should cancel the context after a configurable duration.

10. Workers should stop gracefully, and the main goroutine must wait for all workers to exit.

11. The implementation must not have:
    - Goroutine leaks
    - Deadlocks
    - Workers stuck waiting for jobs after cancellation
    - Workers stuck processing after cancellation
    - Unsafe channel operations

12. Use Go concurrency primitives such as:
    - `context.WithCancel`
    - Goroutines
    - Channels
    - `select`
    - `sync.WaitGroup`
    - Timers

13. Do **not** use `time.Sleep` as a synchronization mechanism.

14. Handle both cases:
    - All jobs finish before cancellation.
    - Cancellation occurs while jobs are still being processed.

15. The output should identify which jobs completed and which jobs were interrupted by cancellation. The result order does not need to match the job order.