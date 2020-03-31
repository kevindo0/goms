package main

import (
	"fmt"
	_"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
            // simulate
            // time.Sleep(time.Second)
		}
	}()

	return out
}

func main() {
	in := producer(10)
	ch := square(in)

	// consumer
	for v := range ch {
		fmt.Println(v)
	}
}