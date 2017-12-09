package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	puzzleInput = "day09/input.txt"
)

type (
	// Node is a tree structure
	Node struct {
		parent   *Node
		children []*Node
	}
)

// NewNode creates a new node
func NewNode() *Node {
	return &Node{
		children: []*Node{},
	}
}

// AddChild adds a node as the current nodes child
func (n *Node) AddChild(other *Node) {
	n.children = append(n.children, other)
	other.parent = n
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

	line := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	chars := strings.Split(line, "")

	buffer := &bytes.Buffer{}

	garbageCount := 0

	garbage := false
	for i := 0; i < len(chars); i++ {
		if garbage {
			switch chars[i] {
			case "!":
				i++
			case ">":
				garbage = false
			default:
				garbageCount++
			}
		} else {
			switch chars[i] {
			case "<":
				garbage = true
			default:
				buffer.WriteString(chars[i])
			}
		}
	}

	chars = strings.Split(buffer.String(), "")

	n := NewNode()

	for i := 0; i < len(chars); i++ {
		switch chars[i] {
		case "{":
			k := NewNode()
			n.AddChild(k)
			n = k
		case "}":
			n = n.parent
		}
	}

	fmt.Println(score(n, 0))
	fmt.Println(garbageCount)
}

func score(n *Node, s int) int {
	k := s
	for _, i := range n.children {
		k += score(i, s+1)
	}
	return k
}
