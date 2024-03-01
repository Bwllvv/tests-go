package main

import "fmt"

func main() {
	colors := "BBBAAAABB"
	color := []rune(colors)
	player := true
	GameStatus := true
	for GameStatus {
		if player {
			for i := 0; i < len(color)-2; i++ {
				if string(color[i]) == "A" && string(color[i+1]) == "A" && string(color[i+2]) == "A" {
					color = append(color[:i], color[i+1:]...)
					player = false
					break
				}
			}
			if player == true {
				GameStatus = false
			}
		} else {
			for i := 0; i < len(color)-2; i++ {
				if string(color[i]) == "B" && string(color[i+1]) == "B" && string(color[i+2]) == "B" {
					color = append(color[:i], color[i+1:]...)
					player = true
					break
				}

			}
			if player == false {
				GameStatus = false
			}
		}
	}
	fmt.Println(!player, 234234)
}
