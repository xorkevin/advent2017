package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	puzzleInput = "day21/input.txt"
	startState  = ".#./..#/###"
)

func stringToGrid(board string) [][]int {
	b := [][]int{}
	for _, i := range strings.Split(board, "/") {
		k := []int{}
		for _, j := range strings.Split(i, "") {
			a := 1
			if j == "." {
				a = 0
			}
			k = append(k, a)
		}
		b = append(b, k)
	}
	return b
}

func gridToString(slice [][]int) string {
	s := []string{}
	for _, i := range slice {
		k := ""
		for _, j := range i {
			if j == 0 {
				k += "."
			} else {
				k += "#"
			}
		}
		s = append(s, k)
	}
	return strings.Join(s, "/")
}

type (
	islice [][]int
)

func (s islice) sym() islice {
	k := s[0][1]
	s[0][1] = s[1][0]
	s[1][0] = k
	if len(s) == 3 {
		k = s[0][2]
		s[0][2] = s[2][0]
		s[2][0] = k
		k = s[1][2]
		s[1][2] = s[2][1]
		s[2][1] = k
	}
	return s
}

func (s islice) refl() islice {
	if len(s) == 3 {
		k := s[0][0]
		s[0][0] = s[2][0]
		s[2][0] = k
		k = s[0][1]
		s[0][1] = s[2][1]
		s[2][1] = k
		k = s[0][2]
		s[0][2] = s[2][2]
		s[2][2] = k
	} else {
		k := s[0][0]
		s[0][0] = s[1][0]
		s[1][0] = k
		k = s[0][1]
		s[0][1] = s[1][1]
		s[1][1] = k
	}
	return s
}

func (s islice) rot() islice {
	return s.sym().refl()
}

func (s islice) prettyString() string {
	a := []string{}
	for _, i := range s {
		k := ""
		for _, j := range i {
			if j == 0 {
				k += "."
			} else {
				k += "#"
			}
		}
		a = append(a, k)
	}
	return strings.Join(a, "\n")
}

func (s islice) asString() string {
	return gridToString(s)
}

func addPolicy(state, after string, policy map[string]string) {
	test := islice(stringToGrid(state))
	policy[test.asString()] = after
	policy[test.rot().asString()] = after
	policy[test.rot().asString()] = after
	policy[test.rot().asString()] = after
	policy[test.rot().refl().asString()] = after
	policy[test.rot().asString()] = after
	policy[test.rot().asString()] = after
	policy[test.rot().asString()] = after
}

type (
	grid struct {
		state islice
	}
)

func newGrid(initState islice) *grid {
	return &grid{
		state: initState,
	}
}

func (g *grid) subboard(r, c int) islice {
	l := 3
	if len(g.state)%2 == 0 {
		l = 2
	}

	sub := islice{}
	for i := 0; i < l; i++ {
		k := []int{}
		for j := 0; j < l; j++ {
			k = append(k, g.state[r*l+i][c*l+j])
		}
		sub = append(sub, k)
	}

	return sub
}

func (g *grid) step(policy map[string]string) {
	nextState := islice{}

	l := 3
	if len(g.state)%2 == 0 {
		l = 2
	}

	for i := 0; i < len(g.state)/l; i++ {
		row := []islice{}
		for j := 0; j < len(g.state)/l; j++ {
			row = append(row, stringToGrid(policy[g.subboard(i, j).asString()]))
		}
		nextState = append(nextState, mergeISlice(l+1, row)...)
	}

	g.state = nextState
}

func (g *grid) numLights() int {
	return strings.Count(g.state.asString(), "#")
}

func mergeISlice(l int, arr []islice) islice {
	a := islice{}
	for i := 0; i < l; i++ {
		row := []int{}
		for j := range arr {
			for k := 0; k < l; k++ {
				row = append(row, arr[j][i][k])
			}
		}
		a = append(a, row)
	}
	return a
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

	policy := map[string]string{}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		s := strings.Split(scanner.Text(), " => ")
		addPolicy(s[0], s[1], policy)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g := newGrid(stringToGrid(startState))

	for i := 0; i < 5; i++ {
		g.step(policy)
	}

	fmt.Println(g.numLights())

	for i := 0; i < 13; i++ {
		g.step(policy)
	}

	fmt.Println(g.numLights())
}
