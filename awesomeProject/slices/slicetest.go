package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 2, 4) // len 1, cap 2
	fmt.Println(nums)
	appendSlice(nums, 1024)
	fmt.Println(nums) // t here
	mutateSlice(nums, 1, 512)
	fmt.Println(nums) // t and
	slicee := []int{1, 2, 3, 4, 5, 6, 7}
	cut := slicee[2:6]
	fmt.Println(slicee, len(slicee), cap(slicee))
	fmt.Println(cut, len(cut), cap(cut))
	cut = cut[:cap(cut)]
	fmt.Println(cut, len(cut), cap(cut))
}
func appendSlice(sl []int, val int) {
	sl = append(sl, val)
	fmt.Println(sl, len(sl), cap(sl))
}
func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
