package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	ch := make(chan int)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			ch <- i
			mu.Unlock()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	count := 0
	for i := range ch {
		fmt.Println(i, count)
		count++
	}
}
