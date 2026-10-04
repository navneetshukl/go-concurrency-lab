# Requirements

1. Create an `input` channel to send integers.
2. Create **3 worker goroutines**.
3. Each worker should:
   - Read integers from the `input` channel.
   - Calculate the square of each integer.
   - Send the result to its own output channel.
4. Use **Fan-Out** to distribute work among the workers.
5. Use **Fan-In** to merge all worker output channels into a single `global` channel.
6. Read and print the results from the `global` channel.
7. Properly close all channels.
8. Use `sync.WaitGroup` to synchronize the goroutines.
9. Ensure the program does not have any **deadlocks or goroutine leaks**.
10. The order of the output does not need to be deterministic.
