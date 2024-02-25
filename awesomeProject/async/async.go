package main

import (
	"context"
	"fmt"
	"math/rand"
)

func main() {
	countAlive := 0
	countDead := 0
	soldersCount := 50
	ctx, cancel := context.WithCancel(context.Background())
	sostoyaniye := make(chan string)
	var status string
	for i := 1; i <= soldersCount; i++ {
		fmt.Printf("Solder №%d disanted\n", i)
		go solder(ctx, i, sostoyaniye)
		status = <-sostoyaniye
		if status == "Dead" {
			countDead++
		} else if status == "bless" {
			cancel()
		} else {
			countAlive++
		}
		fmt.Println("Solder №", i, status)
	}
	fmt.Printf("\nalive solders:%d dead solders:%d", countAlive, countDead)
	cancel()
}

func solder(ctx context.Context, num int, condition chan<- string) {
	Living_randomizer := rand.Intn(100)
	var LiveOrDie string
	if Living_randomizer > 50 {
		LiveOrDie = "Alive"
	} else if Living_randomizer < 5 {
		LiveOrDie = "bless"
	} else {
		LiveOrDie = "Dead"
	}
	select {
	case <-ctx.Done():
		return
	default:
		condition <- LiveOrDie
	}
}
