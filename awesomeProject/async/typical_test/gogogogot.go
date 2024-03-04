package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch chan int
	ch = make(chan int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	K := 0
	go func() {
		for {
			for i := range ch {
				fmt.Println(i)
			}

		}
	}()
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for K < 10000 {
				mu.Lock()
				ch <- K
				//time.Sleep(time.Millisecond * 50)
				mu.Unlock()
				K++
			}
		}()
	}
	wg.Wait()
	close(ch)
}
