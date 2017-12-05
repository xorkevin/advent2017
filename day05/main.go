package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	puzzleInput = "day05/input.txt"
)

func main() {
	file, err := os.Open(puzzleInput)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := []int{}
	nums2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
		nums2 = append(nums2, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(nums)
	part2(nums2)
}

func part1(nums []int) {
	k := 0
	l := len(nums)
	a := 0
	for k < l {
		i := nums[k]
		nums[k]++
		k += i
		a++
	}
	fmt.Println(a)
}

func part2(nums []int) {
	k := 0
	l := len(nums)
	a := 0
	for k < l {
		i := nums[k]
		if i > 2 {
			nums[k]--
		} else {
			nums[k]++
		}
		k += i
		a++
	}
	fmt.Println(a)
}
