package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	puzzleInput = "day11/input.txt"
)

func main() {
	file, err := os.Open(puzzleInput)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	line := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	dirs := strings.Split(line, ",")

	i := 0
	j := 0
	k := 0

	max := 0

	for _, d := range dirs {
		switch d {
		case "n":
			i++
			j--
		case "ne":
			i++
			k--
		case "se":
			j++
			k--
		case "s":
			i--
			j++
		case "sw":
			i--
			k++
		case "nw":
			j--
			k++
		}
		if distance(i, j, k) > max {
			max = distance(i, j, k)
		}
	}

	fmt.Println(distance(i, j, k))
	fmt.Println(max)
}

func distance(i, j, k int) int {
	x := abs(i)
	y := abs(j)
	z := abs(k)
	if x > y {
		if x > z {
			return x
		}
		return z
	}
	if y > z {
		return y
	}
	return z
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
