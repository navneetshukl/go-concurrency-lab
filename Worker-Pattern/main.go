package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	id    int
	value int
}

type response struct {
	id     int
	result int
}

const numWorkers int = 5

func execute(inp chan job, out chan response, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range inp {
		out <- response{
			id:     val.id,
			result: val.value * val.value,
		}
		random := rand.Intn(3)
		time.Sleep(time.Duration(random) * time.Second)
	}
}

func main() {
	input := make(chan job)
	output := make(chan response)
	wg := &sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go execute(input, output, wg)
	}
	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for val := range output {
			log.Printf("Id -> %d | Value -> %d \n", val.id, val.result)
		}
	}()
	for i := 1; i <= 100; i++ {
		input <- job{
			id:    i,
			value: i,
		}
	}
	close(input)
	wg.Wait()
	close(output)
	wg1.Wait()

}
