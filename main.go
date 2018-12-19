package main

import (
	"bufio"
	"os"
)

func combine(nums [][]byte, list [][]byte, start, index int) {
	if index == size {
		for j := 0; j < size; j++ {
			if j > 0 {
				writer.WriteByte(' ')
			}
			writer.Write(list[j])
		}
		writer.WriteByte('\n')
		return
	}

	for i := start; i <= end && end >= size-index; i++ {
		list[index] = nums[i]
		combine(nums, list, i+1, index+1)
	}
}

var writer *bufio.Writer
var size, end int

func main() {
	args := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32"}
	nums := [][]byte{}
	for _, num := range args {
		nums = append(nums, []byte(num))
	}
	size = 8
	end = len(nums) - 1
	list := make([][]byte, size)
	writer = bufio.NewWriter(os.Stdout)
	combine(nums, list, 0, 0)
	writer.Flush()
}
