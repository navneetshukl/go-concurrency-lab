### Worker Pattern — Requirements

1. Implement a **worker pool** in Go that processes jobs concurrently.

2. Create a `Job` structure containing:
   - `ID`
   - `Value`

3. Create a `Result` structure containing:
   - `JobID`
   - `Result`

4. The worker pool should have **5 workers** and process **100 jobs**.

5. Each worker should continuously pick a job, process it by calculating `Value * Value`, and send the result back.

6. Multiple workers must be able to process different jobs concurrently.

7. Every job must be processed **exactly once**, and every job must produce exactly one result.

8. Workers should stop gracefully after all jobs have been processed.

9. The implementation must not have:
   - Lost jobs
   - Duplicate processing
   - Missing results
   - Goroutine leaks
   - Data races

10. Use Go concurrency primitives such as:
   - Goroutines
   - Channels
   - `sync.WaitGroup`

11. Do **not** create one goroutine per job. The system must use exactly **5 worker goroutines**.

12. The result order does not need to match the job order.
