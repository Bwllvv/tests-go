package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	//fmt.Println(nums)
	rangee := 1000
	maxRange := 1000
	Memsum := 1000
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				rangee = int(math.Abs(float64(sum - target)))

				if rangee < maxRange {
					Memsum = sum
					maxRange = rangee
					//fmt.Println(nums[i], nums[j], nums[k])
				}
			}
		}
	}
	return Memsum
}
