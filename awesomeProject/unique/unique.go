package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	uni := make(map[string]bool)
	for reader.Scan() {
		txt := reader.Text()
		if _, found := uni[txt]; found {
			continue
		}
		uni[txt] = true
		fmt.Println(txt)
	}

}
