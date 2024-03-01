package main

import (
	"fmt"
	"sync"
	"time"
)

var Mutex sync.Mutex

func main() {
	mapa := make(map[int]int)
	for i := 1; i <= 10; i += 10 {
		go setvalue(mapa, i)
	}
	//go setvalue(mapa)
	time.Sleep(time.Second)
	a := getvalue(mapa)
	if a != nil {
		fmt.Println(a)
		return
	}

}

func setvalue(mapa map[int]int, m int) {
	for i := 10 * m; i < 100*m; i++ {
		Mutex.Lock()
		mapa[i] = i + 1
		Mutex.Unlock()
	}
}

func getvalue(mapa map[int]int) error {
	for j := 1; j <= 10; j += 10 {

		for i := 10 * j; i < 100*j; i++ {
			//fmt.Println(mapa[i], i)
			Mutex.Lock()
			if mapa[i]-i != 1 {
				return fmt.Errorf("произошел пиздец")
			}
			Mutex.Unlock()
			fmt.Printf("key:%d value:%d\n", i, mapa[i])
		}
	}
	return nil
}
