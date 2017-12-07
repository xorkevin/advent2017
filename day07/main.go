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
	puzzleInput = "day07/input.txt"
	rootName    = "mwzaxaj"
)

type (
	Node struct {
		name         string
		weight       int
		children     []string
		branchWeight int
	}
)

func (n *Node) CalcWeight(nodes map[string]*Node) int {
	if n.branchWeight > -1 {
		return n.branchWeight
	}
	if n.children == nil {
		n.branchWeight = n.weight
	} else {
		w := 0
		for _, i := range n.children {
			w += nodes[i].CalcWeight(nodes)
		}
		n.branchWeight = w + n.weight
	}
	return n.branchWeight
}

func (n *Node) FindUnbalance(nodes map[string]*Node, expected int) int {
	if n.children == nil {
		return expected
	}

	c := map[int]int{}
	for _, i := range n.children {
		if _, ok := c[nodes[i].branchWeight]; ok {
			c[nodes[i].branchWeight]++
		} else {
			c[nodes[i].branchWeight] = 1
		}
	}

	if len(c) == 1 {
		return expected - n.branchWeight + n.weight
	}

	var differentWeight int
	var sameWeight int
	for k, v := range c {
		if v < 2 {
			differentWeight = k
		} else {
			sameWeight = k
		}
	}
	for _, i := range n.children {
		if nodes[i].branchWeight == differentWeight {
			return nodes[i].FindUnbalance(nodes, sameWeight)
		}
	}
	return -314
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

	nodes := map[string]*Node{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "-")
		s0 := strings.Split(strings.Trim(s[0], " "), "(")
		name := strings.Trim(s0[0], " ")
		weight, err := strconv.Atoi(strings.Trim(strings.Trim(s0[1], " "), ")"))
		if err != nil {
			log.Fatal(err)
		}
		var children []string
		if len(s) > 1 {
			children = strings.Split(strings.Trim(strings.Trim(s[1], ">"), " "), ", ")
		}
		nodes[name] = &Node{
			name:         name,
			weight:       weight,
			children:     children,
			branchWeight: -1,
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//part1()
	part2(nodes)
}

func part1() {
	fmt.Println("Hello World")
}

func part2(nodes map[string]*Node) {
	nodes[rootName].CalcWeight(nodes)
	fmt.Println(nodes[rootName].FindUnbalance(nodes, 0))
}
