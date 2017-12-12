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
	puzzleInput = "day12/input.txt"
)

type (
	Node struct {
		id       int
		children []int
	}
)

func NewNode(id int, children []int) *Node {
	return &Node{
		id:       id,
		children: children,
	}
}

func (n *Node) Check(nodes []*Node, m map[int]bool) {
	if _, ok := m[n.id]; ok {
		return
	}
	m[n.id] = true
	for _, i := range n.children {
		nodes[i].Check(nodes, m)
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

	nodes := []*Node{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " <-> ")
		num, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		snums := strings.Split(s[1], ", ")
		nums := make([]int, 0, len(snums))
		for _, i := range snums {
			n, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, n)
		}
		nodes = append(nodes, NewNode(num, nums))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := map[int]bool{}

	nodes[0].Check(nodes, m)

	fmt.Println(len(m))

	groups := 1

	for len(m) < len(nodes) {
		k := 0
		for ok := true; ok; _, ok = m[k] {
			k++
		}
		nodes[k].Check(nodes, m)
		groups++
	}
	fmt.Println(groups)
}
