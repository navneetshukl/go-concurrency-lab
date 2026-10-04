package main

import (
	"log"
	"sync"
)

var numWorkers int = 3

func execute(input, output chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range input {
		output <- v * v
	}
	close(output)
}

func main() {
	input := make(chan int)
	wg := &sync.WaitGroup{}
	output := make([]chan int, numWorkers)
	global := make(chan int)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		output[i] = make(chan int)
		go execute(input, output[i], wg)
	}

	wg1 := &sync.WaitGroup{}

	go func() {

		for i := 0; i < numWorkers; i++ {
			wg1.Add(1)
			go func(out chan int) {
				defer wg1.Done()
				for v := range out {
					global <- v
				}

			}(output[i])
		}
		wg1.Wait()
		close(global)
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			input <- i
		}
		close(input)
	}()
	for val := range global {
		log.Println(val)
	}
	wg.Wait()

}
