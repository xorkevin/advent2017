package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	puzzleInput = "day22/input.txt"
	boardSize   = 1000
	totalSteps  = 10000
	totalSteps2 = 10000000
)

const (
	dirup = iota
	dirdown
	dirleft
	dirright
)

type (
	carrier struct {
		posx, posy int
		dir        int
		count      int
	}
)

func newCarrier(posx, posy int) *carrier {
	return &carrier{
		posx:  posx,
		posy:  posy,
		dir:   dirup,
		count: 0,
	}
}

func (c *carrier) turn(right bool) {
	if right {
		switch c.dir {
		case dirup:
			c.dir = dirright
		case dirright:
			c.dir = dirdown
		case dirdown:
			c.dir = dirleft
		case dirleft:
			c.dir = dirup
		}
	} else {
		switch c.dir {
		case dirup:
			c.dir = dirleft
		case dirleft:
			c.dir = dirdown
		case dirdown:
			c.dir = dirright
		case dirright:
			c.dir = dirup
		}
	}
}

func (c *carrier) forward() {
	switch c.dir {
	case dirup:
		c.posy--
	case dirdown:
		c.posy++
	case dirleft:
		c.posx--
	case dirright:
		c.posx++
	}
}

func (c *carrier) step(grid [][]int) {
	right := false
	k := grid[c.posy][c.posx]
	if k > 0 {
		right = true
	} else {
		c.count++
	}
	grid[c.posy][c.posx] = (k + 1) % 2
	c.turn(right)
	c.forward()
}

func (c *carrier) step2(grid [][]int) {
	k := grid[c.posy][c.posx]
	switch k {
	case 0:
		c.turn(false)
	case 1:
		c.count++
	case 2:
		c.turn(true)
	case 3:
		c.turn(false)
		c.turn(false)
	}
	grid[c.posy][c.posx] = (k + 1) % 4
	c.forward()
}

func printGrid(grid [][]int, x, y, dir int) {
	for a, i := range grid {
		for b, j := range i {
			if x == b && y == a {
				switch dir {
				case dirup:
					fmt.Print("^")
				case dirdown:
					fmt.Print("v")
				case dirleft:
					fmt.Print("<")
				case dirright:
					fmt.Print(">")
				}
			} else if j > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
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

	temp := [][]int{}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		k := []int{}
		for _, i := range strings.Split(scanner.Text(), "") {
			s := 0
			if i == "#" {
				s = 1
			}
			k = append(k, s)
		}
		temp = append(temp, k)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid := make([][]int, boardSize)
	grid2 := make([][]int, boardSize)
	for i := range grid {
		grid[i] = make([]int, boardSize)
		grid2[i] = make([]int, boardSize)
	}

	mid := len(grid) / 2
	start := mid - len(temp)/2

	for i, a := range temp {
		for j, b := range a {
			grid[start+i][start+j] = b
			grid2[start+i][start+j] = b * 2
		}
	}

	c := newCarrier(mid, mid)
	for i := 0; i < totalSteps; i++ {
		c.step(grid)
	}
	fmt.Println(c.count)
	c2 := newCarrier(mid, mid)
	for i := 0; i < totalSteps2; i++ {
		c2.step2(grid2)
	}

	fmt.Println(c2.count)
}
