package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	puzzleInput = "day13/input.txt"
)

type (
	firewall struct {
		layers []int
	}
)

func newFirewall(layers []int) *firewall {
	return &firewall{
		layers: layers,
	}
}

func (f *firewall) checkSeverity() int {
	severity := 0
	for n, i := range f.layers {
		if i > 0 && n%(i*2-2) == 0 {
			severity += n * i
		}
	}
	return severity
}

func (f *firewall) checkCaught(delay int) bool {
	for n, i := range f.layers {
		if i > 0 && (delay+n)%(i*2-2) == 0 {
			return true
		}
	}
	return false
}

func (f *firewall) findDelay() int {
	i := -1
	caught := true
	for caught {
		i++
		caught = f.checkCaught(i)
	}
	return i
}

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

	f := make([]int, 100)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ": ")
		num, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		f[num] = num2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fire := newFirewall(f)

	fmt.Println(fire.checkSeverity())

	fmt.Println(fire.findDelay())
}
