### Thread-Safe Blocking Queue — Requirements

1. Implement a **generic thread-safe queue** in Go.

2. The queue should have a **fixed maximum capacity**.

3. `Enqueue(item)`
   - Adds an item to the queue.
   - Must be safe for concurrent producers.
   - If the queue is full, the producer must **block** until space is available.
   - Should return an error if the queue is closed.

4. `Dequeue()`
   - Removes and returns the oldest item.
   - Must be safe for concurrent consumers.
   - If the queue is empty, the consumer must **block** until an item is available.
   - Must maintain **FIFO ordering**.
   - If the queue is closed and empty, it should return an error.

5. `Size()`
   - Returns the current number of items.
   - Must be safe to call concurrently with `Enqueue()` and `Dequeue()`.

6. `Close()`
   - Closes the queue.
   - No new items can be added after closing.
   - Blocked producers must be unblocked.
   - Blocked consumers must be unblocked.
   - Existing items must still be consumable after closing.
   - Once closed and empty, `Dequeue()` should return an error.

7. Multiple producers and consumers must be able to operate on the queue concurrently.

8. The implementation must not have:
   - Data races
   - Lost items
   - Duplicate items
   - Corrupted queue state
   - Goroutine leaks

9. Do **not** use Go channels to implement the queue.

10. Use synchronization primitives such as:
    - `sync.Mutex`
    - `sync.Cond`
    - `sync.WaitGroup` where required.
