package main

import (
	"fmt"
)

const (
	puzzleInput = 335
)

func main() {
	k := []int{0}
	j := 0
	for i := 0; i < 2017; i++ {
		j = (j+puzzleInput)%len(k) + 1
		k = append(k, 0)
		copy(k[j+1:], k[j:])
		k[j] = i + 1
	}
	fmt.Println(k[j+1])
	l := len(k)
	afterZero := k[1]
	for i := 2017; i < 50000000; i++ {
		j = (j+puzzleInput)%l + 1
		l++
		if j == 1 {
			afterZero = i + 1
		}
	}
	fmt.Println(afterZero)
}
