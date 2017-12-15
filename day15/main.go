package main

import (
	"fmt"
)

const (
	puzzleInputA = 277
	puzzleInputB = 349
	rounds       = 40000000
	rounds2      = 5000000
)

func main() {
	a := puzzleInputA
	b := puzzleInputB
	k := 0
	for i := 0; i < rounds; i++ {
		match, na, nb := generate(a, b)
		a = na
		b = nb
		if match {
			k++
		}
	}
	fmt.Println(k)

	reviewA := make(chan int)
	reviewB := make(chan int)
	hasher(puzzleInputA, genseedA, modA, reviewA)
	hasher(puzzleInputB, genseedB, modB, reviewB)

	k2 := 0
	for i := 0; i < rounds2; i++ {
		fromA := <-reviewA
		fromB := <-reviewB
		if fromA%65536 == fromB%65536 {
			k2++
		}
	}
	fmt.Println(k2)
}

const (
	genseedA = 16807
	genseedB = 48271
	modA     = 4
	modB     = 8
	genseedD = 2147483647
)

func generate(a, b int) (bool, int, int) {
	newA := a * genseedA % genseedD
	newB := b * genseedB % genseedD
	return newA%65536 == newB%65536, newA, newB
}

func hasher(start, factor, mod int, send chan<- int) {
	go func() {
		for {
			start = start * factor % genseedD
			if start%mod == 0 {
				send <- start
			}
		}
	}()
}
