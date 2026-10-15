package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	id       int
	duration int
}

type result struct {
	id     int
	status string
}

const numWorkers int = 3

func execute(in chan job, out chan result, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for v := range in {
		res := result{
			id: v.id,
		}
		select {
		case <-ctx.Done():
			res.status = "Context Cancelled At Start"
			out <- res
			return

		default:
			timer := time.NewTimer(time.Duration(v.duration) * time.Second)
			defer timer.Stop()
			select {
			case <-ctx.Done():
				res.status = "Context Cancelled Just Before Start"
				out <- res
				return

			case <-timer.C:
				select {
				case <-ctx.Done():
					res.status = "Context Cancelled Before Finish"
					out <- res
					return
				default:
					res.status = "SUCCESS"
					out <- res

				}
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	inp, out := make(chan job), make(chan result)
	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go execute(inp, out, &wg, ctx)
	}

	go func() {
		defer func() {
			close(inp)
			wg.Wait()
			close(out)
		}()
		for i := 1; i <= 200; i++ {
			select {
			case <-ctx.Done():
				return

			case inp <- job{
				id:       i,
				duration: rand.Intn(3),
			}:
			}

		}

	}()

	for v := range out {
		log.Printf("%v\n", v)
	}

}
