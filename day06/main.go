package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	puzzleInput = "10	3	15	10	5	15	5	15	9	2	5	8	5	2	3	6"
)

func main() {
	part1()
	part2()
}

func part1() {
	snums := strings.Split(puzzleInput, "\t")
	nums := make([]int, 0, len(snums))
	for _, i := range snums {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	prevStates := map[string]bool{}
	numkey := stringifyIntArray(nums)
	found := false
	count := 0
	for !found {
		prevStates[numkey] = true

		redistribute(nums)

		numkey = stringifyIntArray(nums)
		_, ok := prevStates[numkey]
		found = ok
		count++
	}
	fmt.Println(count)
}

func redistribute(nums []int) {
	l := len(nums)
	index := 0
	max := nums[index]
	for n, i := range nums {
		if i > max {
			index = n
			max = i
		}
	}
	nums[index] = 0
	index = (index + 1) % l
	for i := 0; i < max; i++ {
		nums[index]++
		index = (index + 1) % l
	}
}

func stringifyIntArray(nums []int) string {
	k := make([]string, 0, len(nums))
	for _, i := range nums {
		k = append(k, strconv.Itoa(i))
	}
	return strings.Join(k, ",")
}

func part2() {
	snums := strings.Split(puzzleInput, "\t")
	nums := make([]int, 0, len(snums))
	for _, i := range snums {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	prevStates := map[string]int{}
	numkey := stringifyIntArray(nums)
	found := -1
	count := 0
	for found < 0 {
		prevStates[numkey] = count

		redistribute(nums)

		numkey = stringifyIntArray(nums)
		if i, ok := prevStates[numkey]; ok {
			found = i
		}
		count++
	}
	fmt.Println(count - found)
}
