package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)
	go func() {
		for {
			select {
			case m := <-c:
				fmt.Println(m)
			case d := <-time.After(5 * time.Second):
				fmt.Println("time out")
				fmt.Println("current Time :", d)
			}
		}
	}()
	fmt.Print("a")
	<-c
}
