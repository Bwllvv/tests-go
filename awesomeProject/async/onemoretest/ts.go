package main

import (
	"fmt"
	"sync"
)

func main() {
	urls := []string{"nechet", "chet", "nechet1", "chet1"}
	arr := []int{1, 2, 3, 1, 2, 4, 8}
	mapa := make(map[string][]int)
	workers := 1
	for _, value := range urls {
		mapa[value] = appendnumber(arr, workers)
		workers++
	}
	//time.Sleep(time.Second * 5)
	fmt.Println(mapa)
}

func checknumber(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

// попробую найти колличество запросов по конкретному урлу
func appendnumber(arr []int, workers int) []int {
	wg := sync.WaitGroup{}
	var arr2 []int
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < len(arr); j++ {
				if workers%2 == 0 {
					if checknumber(arr[j]) {
						//time.Sleep(time.Second)
						arr2 = append(arr2, arr[j])
					}
				} else {
					if !checknumber(arr[j]) {
						//time.Sleep(time.Second)
						arr2 = append(arr2, arr[j])
					}
				}
			}

		}()
	}
	wg.Wait()
	return arr2
}
