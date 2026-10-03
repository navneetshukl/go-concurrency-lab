package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	arr     []int
	mut     *sync.Mutex
	maxSize int
	size    int
	isOpen  bool
	cond    *sync.Cond
}

func NewQueue(size int) *Queue {
	q := &Queue{
		arr:     make([]int, 0, size),
		mut:     &sync.Mutex{},
		maxSize: size,
		size:    0,
		isOpen:  true,
	}
	q.cond = sync.NewCond(q.mut)
	return q
}

func (q *Queue) Enqueue(val int) bool {
	q.mut.Lock()
	defer q.mut.Unlock()
	if !q.isOpen {
		return false
	}

	for q.size >= q.maxSize {
		q.cond.Wait()
	}
	if !q.isOpen {
		return false
	}
	q.arr = append(q.arr, val)
	q.size++
	q.cond.Signal()
	return true
}

func (q *Queue) Dequeue() (int, bool) {
	q.mut.Lock()
	defer q.mut.Unlock()
	if !q.isOpen && q.size == 0 {
		return 0, false
	}
	val := 0

	for q.size <= 0 {
		q.cond.Wait()
		if !q.isOpen && q.size == 0 {
			return 0, false
		}
	}

	val = q.arr[0]
	q.arr = q.arr[1:]
	q.size--
	q.cond.Signal()
	return val, true

}

func (q *Queue) Size() int {
	q.mut.Lock()
	defer q.mut.Unlock()
	return q.size
}

func (q *Queue) Close() {
	q.mut.Lock()
	defer q.mut.Unlock()
	q.isOpen = false
	q.cond.Broadcast()
}
func main() {
	queue := NewQueue(3)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// 3 Producers
	for i := 1; i <= 3; i++ {
		producerWg.Add(1)

		go func(id int) {
			defer producerWg.Done()

			for j := 1; j <= 5; j++ {
				val := id*10 + j

				ok := queue.Enqueue(val)

				if ok {
					fmt.Printf("Producer %d: Enqueued %d | Size: %d\n",
						id, val, queue.Size())
				} else {
					fmt.Printf("Producer %d: Queue closed, couldn't enqueue %d\n",
						id, val)
				}
			}
		}(i)
	}

	// 2 Consumers
	for i := 1; i <= 2; i++ {
		consumerWg.Add(1)

		go func(id int) {
			defer consumerWg.Done()

			for {
				val, ok := queue.Dequeue()

				if !ok {
					fmt.Printf("Consumer %d: Queue closed and empty\n", id)
					return
				}

				fmt.Printf("Consumer %d: Dequeued %d | Size: %d\n",
					id, val, queue.Size())

				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Wait until all producers finish.
	producerWg.Wait()

	fmt.Println("All producers finished")

	// Close the queue.
	queue.Close()

	fmt.Println("Queue closed")

	// Consumers will consume remaining items and then exit.
	consumerWg.Wait()

	fmt.Println("All consumers finished")
	fmt.Println("Final queue size:", queue.Size())
}
