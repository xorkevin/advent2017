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
	puzzleInput = "day20/input.txt"
)

type (
	point struct {
		px, py, pz int
		vx, vy, vz int
		ax, ay, az int
	}
)

func (p *point) step() {
	p.vx += p.ax
	p.vy += p.ay
	p.vz += p.az
	p.px += p.vx
	p.py += p.vy
	p.pz += p.vz
}

func (p *point) dist() int {
	return abs(p.px) + abs(p.py) + abs(p.pz)
}

func (p *point) key() string {
	return fmt.Sprintf("%d,%d,%d", p.px, p.py, p.pz)
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

	nodes := []*point{}
	nodes2 := map[int]*point{}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		s := strings.Split(scanner.Text(), ", ")
		p := strings.Split(strings.Trim(s[0], "p=<>"), ",")
		v := strings.Split(strings.Trim(s[1], "v=<>"), ",")
		a := strings.Split(strings.Trim(s[2], "a=<>"), ",")

		n := point{}
		if num, err := strconv.Atoi(a[0]); err != nil {
			log.Fatal(err)
		} else {
			n.ax = num
		}
		if num, err := strconv.Atoi(v[0]); err != nil {
			log.Fatal(err)
		} else {
			n.vx = num
		}
		if num, err := strconv.Atoi(p[0]); err != nil {
			log.Fatal(err)
		} else {
			n.px = num
		}

		if num, err := strconv.Atoi(a[1]); err != nil {
			log.Fatal(err)
		} else {
			n.ay = num
		}
		if num, err := strconv.Atoi(v[1]); err != nil {
			log.Fatal(err)
		} else {
			n.vy = num
		}
		if num, err := strconv.Atoi(p[1]); err != nil {
			log.Fatal(err)
		} else {
			n.py = num
		}

		if num, err := strconv.Atoi(a[2]); err != nil {
			log.Fatal(err)
		} else {
			n.az = num
		}
		if num, err := strconv.Atoi(v[2]); err != nil {
			log.Fatal(err)
		} else {
			n.vz = num
		}
		if num, err := strconv.Atoi(p[2]); err != nil {
			log.Fatal(err)
		} else {
			n.pz = num
		}

		n2 := n

		nodes = append(nodes, &n)
		nodes2[i] = &n2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1024; i++ {
		for _, i := range nodes {
			i.step()
		}
	}

	min := 999999999999
	minIndex := 0
	for n, i := range nodes {
		if i.dist() < min {
			minIndex = n
			min = i.dist()
		}
	}

	fmt.Println(minIndex)

	for i := 0; i < 1024; i++ {
		checkCollide(nodes2)
	}

	fmt.Println(len(nodes2))
}

func checkCollide(nodes map[int]*point) {
	grid := map[string][]int{}
	for k, v := range nodes {
		v.step()
		key := v.key()
		if _, ok := grid[key]; ok {
			grid[key] = append(grid[key], k)
		} else {
			grid[key] = []int{k}
		}
	}

	for _, v := range grid {
		if len(v) > 1 {
			for _, i := range v {
				delete(nodes, i)
			}
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
