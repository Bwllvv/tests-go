package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Закрываем контекст после завершения main
	urls := []string{"qwe", "asd", "zxc"}
	for _, url := range urls {
		ch := make(chan int) // Создаем новый канал для каждого вызова
		visitcount(ctx, 3, ch)
		go func(u string) {
			for j := range ch {
				fmt.Println(u, j)
			}
		}(url)
	}
	time.Sleep(time.Second * 3)
}

func visitcount(ctx context.Context, workers int, ch chan int) {
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			case ch <- rand.Intn(10):
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}() // Закрываем канал асинхронно
}
