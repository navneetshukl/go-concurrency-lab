package main

import "sync"

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
		arr:     make([]int, size), // here is small issue
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

	// what if queue is closed here after sync.cond.wait

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

		// what is queue is closed and q.size=0
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
