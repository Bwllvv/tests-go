package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		defer close(ch2)
		iter := 0
		iter1 := 0
	loop:
		for {
			select {
			case v1 := <-ch:
				fmt.Println("iter № ", iter, "value:", v1)
				iter++
				//time.Sleep(time.Millisecond * 100)
			case v2 := <-ch2:
				fmt.Println("iter № ", iter1, "value:", v2)
				iter1++
				//time.Sleep(time.Millisecond * 100)
			case <-ctx.Done():
				fmt.Println("Work is done")
				break loop
			default:
				//fmt.Println("Loading...")
				//time.Sleep(time.Millisecond * 30)
			}

		}
	}()
	i := 0
	for t := 0; t < 100; t++ {
		go func() {

			for i < 101 {
				mu.Lock()
				if i <= 100 {
					ch <- i * i
					ch2 <- i * 10
					//time.Sleep(time.Millisecond * 50)
					i++
				}
				mu.Unlock()
			}
			time.Sleep(time.Millisecond * 50)
			cancel()
		}()
	}
	wg.Wait()
	fmt.Println("finish")
}
