package main

import "fmt"

func combine(nums []int, list map[int]int, start, end, index, size int) {
	if index == size {
		for j := 0; j < size; j++ {
			fmt.Printf(" %2d", list[j])
		}
		fmt.Println("")
		return
	}

	for i := start; i <= end && end >= size-index; i++ {
		list[index] = nums[i]
		combine(nums, list, i+1, end, index+1, size)
	}
}

func main() {
	size := 8
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	list := make(map[int]int)
	combine(nums, list, 0, len(nums)-1, 0, size)
}
