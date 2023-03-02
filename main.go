package main

import (
	"bufio"
	"flag"
	"os"
	"runtime/pprof"
	"strconv"
)

const max = 32

func combine(nums [][]byte, list [][]byte, start, index int) {
	if index == size {
		for j := 0; j < size; j++ {
			_, _ = writer.Write(list[j])
		}
		_ = writer.WriteByte('\n')
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
	saveProfile := flag.Bool("profile", false, "save profile")
	flag.Parse()

	if *saveProfile {
		f, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	nums := [][]byte{}
	leftPad := "0"
	for i := 1; i <= max; i++ {
		if i > 9 {
			leftPad = ""
		}
		nums = append(nums, []byte(leftPad+strconv.Itoa(i)))
	}
	size = 8
	end = len(nums) - 1
	list := make([][]byte, size)
	writer = bufio.NewWriter(os.Stdout)
	combine(nums, list, 0, 0)
	_ = writer.Flush()
}
