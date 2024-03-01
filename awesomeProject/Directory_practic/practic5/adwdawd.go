package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height1 := []int{2, 3, 4, 5, 18, 17, 6}
	height3 := []int{1, 2, 1}
	fmt.Println(maxArea(height))
	fmt.Println(maxArea(height1))
	fmt.Println(maxArea(height3))
}
func maxArea(height []int) int {
	//biggest := height[len(height)-1]
	//iNeed := 0
	maxMulti := 0
	multi := 0
	i := 0
	j := len(height) - 1
	point1 := 0
	point2 := 0
	for z := 0; z < len(height); z++ {
		point1 = height[i]
		point2 = height[j]
		if point2 > point1 {
			multi = (j - i) * point1
			i++
		} else {
			multi = (j - i) * point2
			j--
		}
		if maxMulti < multi {
			maxMulti = multi
		}
	}
	return maxMulti
}
