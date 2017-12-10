package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	puzzleInput = "31,2,85,1,80,109,35,63,98,255,0,13,105,254,128,33"
)

var (
	part2Suffix = []int{17, 31, 73, 47, 23}
)

func main() {
	part1()
	part2()
}

func part1() {
	s := strings.Split(puzzleInput, ",")
	nums := make([]int, 0, len(s))
	for _, i := range s {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	list := make([]int, 0, 256)
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	k := 0
	for n, i := range nums {
		flip(list, k, i)
		k = (n + k + i) % len(list)
	}

	fmt.Println(list[0] * list[1])
}

func part2() {
	nums := sequence(puzzleInput)

	list := make([]int, 0, 256)
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	k := 0
	skip := 0

	for i := 0; i < 64; i++ {
		for _, i := range nums {
			flip(list, k, i)
			k = (skip + k + i) % len(list)
			skip++
		}
	}

	list = dense(list)

	for _, i := range list {
		fmt.Printf("%x", i)
	}
	fmt.Println()
}

func flip(nums []int, start, size int) {
	l := len(nums)
	end := (l + start + size - 1) % l
	for i := 0; i < size/2; i++ {
		swap(nums, (start+i)%l, (l+end-i)%l)
	}
}

func swap(nums []int, a, b int) {
	k := nums[a]
	nums[a] = nums[b]
	nums[b] = k
}

func sequence(inp string) []int {
	binp := []byte(inp)
	nums := make([]int, 0, len(binp)+len(part2Suffix))
	for _, i := range binp {
		nums = append(nums, int(i))
	}
	nums = append(nums, part2Suffix...)
	return nums
}

func dense(nums []int) []int {
	list := make([]int, 0, len(nums)/16)
	for i := 0; i < len(nums); i += 16 {
		k := 0
		for j := 0; j < 16; j++ {
			k ^= nums[i+j]
		}
		list = append(list, k)
	}
	return list
}
