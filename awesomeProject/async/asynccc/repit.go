package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	//mu := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//mu.Lock()
			fmt.Println(i)
			//mu.Unlock()
		}(i)
	}
	wg.Wait()
}
