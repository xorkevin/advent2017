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
	puzzleInput = "day24/input.txt"
)

type (
	component struct {
		left, right, id int
	}
)

func newComponent(left, right, id int) *component {
	return &component{
		left:  left,
		right: right,
		id:    id,
	}
}

func parseComponentString(componentString string, id int) *component {
	s := strings.Split(componentString, "/")
	num1, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatal(err)
	}
	num2, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	return newComponent(num1, num2, id)
}

type (
	bridge struct {
		parent   *bridge
		c        *component
		unused   int
		strength int
		length   int
	}
)

func newBridge(c *component) *bridge {
	return &bridge{
		parent:   nil,
		c:        c,
		unused:   c.right,
		strength: c.left + c.right,
		length:   1,
	}
}

func (b *bridge) add(other *bridge) *bridge {
	other.parent = b
	other.strength = b.strength + other.strength
	if other.c.left == b.unused {
		other.unused = other.c.right
	} else if other.c.right == b.unused {
		other.unused = other.c.left
	}
	other.length = b.length + 1
	return other
}

type (
	componentBank struct {
		bank map[int][]*component
	}
)

func newComponentBank() *componentBank {
	return &componentBank{
		bank: map[int][]*component{},
	}
}

func (b *componentBank) add(c *component) {
	if _, ok := b.bank[c.left]; !ok {
		b.bank[c.left] = []*component{}
	}
	if _, ok := b.bank[c.right]; !ok {
		b.bank[c.right] = []*component{}
	}
	b.bank[c.left] = append(b.bank[c.left], c)
	b.bank[c.right] = append(b.bank[c.right], c)
}

func (b *componentBank) findStrongestBridge(bg *bridge, part2 bool) *bridge {
	used := map[int]bool{}
	for i := bg; i != nil; i = i.parent {
		used[i.c.id] = true
	}
	max := bg
	for _, i := range b.bank[bg.unused] {
		if _, ok := used[i.id]; !ok {
			used[i.id] = true
			k := b.findStrongestBridge(bg.add(newBridge(i)), part2)
			if part2 {
				if k.length >= max.length {
					if k.strength > max.strength {
						max = k
					}
				}
			} else {
				if k.strength > max.strength {
					max = k
				}
			}
		}
	}
	return max
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

	bank := newComponentBank()
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		bank.add(parseComponentString(scanner.Text(), i))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	max := bank.findStrongestBridge(newBridge(newComponent(0, 0, -1)), false)
	fmt.Println(max.strength)
	max2 := bank.findStrongestBridge(newBridge(newComponent(0, 0, -1)), true)
	fmt.Println(max2.strength)
}
