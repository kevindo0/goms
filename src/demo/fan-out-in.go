package main

import (
	"fmt"
	"sync"
	"time"
)

func Produce(nums ...int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()
	return out
}

func Square(in <-chan int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			val := v * v
			out <- val
			time.Sleep(1*time.Second)
		}
	}()
	return out
}

func Merge(ins ...<-chan int) <- chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(i <-chan int) {
			defer wg.Done()
			for v := range i {
				out <- v
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := Produce(3, 4, 5, 6, 2, 1)
	c1 := Square(in)
	c2 := Square(in)
	c3 := Square(in)
	out := Merge(c1, c2, c3)
	for ret := range out {
		fmt.Printf("%3d ", ret)
	}
	  fmt.Println()
}