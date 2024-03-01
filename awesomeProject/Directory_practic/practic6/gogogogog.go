package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums) // Сортируем массив
	var res [][]int
	fmt.Println(nums)
	for i := 0; i < len(nums)-2; i++ {
		// Пропускаем повторяющиеся значения
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l, r = next(nums, l, r)
				//break
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

// Функция для пропуска повторяющихся значений
func next(nums []int, l int, r int) (int, int) {
	l++
	for l < r && nums[l] == nums[l-1] {
		l++
	}
	r--
	for l < r && nums[r] == nums[r+1] {
		r--
	}
	return l, r
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
