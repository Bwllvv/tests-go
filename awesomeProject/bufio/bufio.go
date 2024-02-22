package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("file1.txt")
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}
}
