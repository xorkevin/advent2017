package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	puzzleInput = "day19/input.txt"
)

const (
	vertical = iota
	horizontal
	cross
	space
	letter
)

const (
	up = iota
	down
	left
	right
)

type (
	square struct {
		class int
		val   string
	}

	// Network represents a network
	Network struct {
		pr        int
		pc        int
		direction int
		data      [][]*square
		letters   string
		steps     int
	}
)

func newNetwork(startX int, data [][]*square) *Network {
	return &Network{
		pr:        0,
		pc:        startX,
		direction: down,
		data:      data,
		letters:   "",
		steps:     0,
	}
}

func (n *Network) step() bool {
	sq := n.data[n.pr][n.pc]
	switch sq.class {
	case cross:
		switch n.direction {
		case up, down:
			if n.data[n.pr][n.pc-1].class != space {
				n.direction = left
			} else if n.data[n.pr][n.pc+1].class != space {
				n.direction = right
			}
		case left, right:
			if n.data[n.pr-1][n.pc].class != space {
				n.direction = up
			} else if n.data[n.pr+1][n.pc].class != space {
				n.direction = down
			}
		}
	case letter:
		if len(sq.val) > 0 {
			n.letters += sq.val
		}
	case space:
		return false
	}

	switch n.direction {
	case up:
		n.pr--
	case down:
		n.pr++
	case left:
		n.pc--
	case right:
		n.pc++
	}

	n.steps++

	return n.pr >= 0 && n.pr < len(n.data) && n.pc >= 0 && n.pc < len(n.data[0])
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

	data := [][]*square{}
	scanner := bufio.NewScanner(file)
	firstRow := true
	startX := 0
	for scanner.Scan() {
		row := []*square{}
		for n, i := range scanner.Text() {
			var s *square
			switch i {
			case '|':
				s = &square{
					class: vertical,
					val:   "",
				}
				if firstRow {
					firstRow = false
					startX = n
				}
			case '-':
				s = &square{
					class: horizontal,
					val:   "",
				}
			case '+':
				s = &square{
					class: cross,
					val:   "",
				}
			case ' ':
				s = &square{
					class: space,
					val:   "",
				}
			default:
				if i-'A' > 26 {
					log.Fatalf("%s is not a valid char\n", string(i))
				}
				s = &square{
					class: letter,
					val:   string(i),
				}
			}
			row = append(row, s)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	network := newNetwork(startX, data)

	for network.step() {
	}

	fmt.Println(network.letters)
	fmt.Println(network.steps)
}
