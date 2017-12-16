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
	puzzleInput = "day16/input.txt"
	initState   = "abcdefghijklmnop"
	billion     = 1000000000
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

	fmt.Println(run(initState, line))

	seen := map[string]bool{}
	history := []string{}
	state := initState
	for i := 0; i < billion; i++ {
		if _, ok := seen[state]; ok {
			state = history[billion%i]
			break
		} else {
			seen[state] = true
			history = append(history, state)
		}
		state = run(state, line)
	}
	fmt.Println(state)
}

func run(state, line string) string {
	for _, i := range strings.Split(line, ",") {
		s1 := i[0:1]
		s2 := i[1:]
		switch s1 {
		case "s":
			num, err := strconv.Atoi(s2)
			if err != nil {
				log.Fatal(err)
			}
			state = state[16-num:] + state[:16-num]
		case "x":
			s := strings.Split(s2, "/")
			num1, err := strconv.Atoi(s[0])
			if err != nil {
				log.Fatal(err)
			}
			num2, err := strconv.Atoi(s[1])
			if err != nil {
				log.Fatal(err)
			}
			j := []byte(state)
			k := j[num1]
			j[num1] = j[num2]
			j[num2] = k
			state = string(j)
		case "p":
			s := strings.Split(s2, "/")
			char1 := s[0]
			char2 := s[1]
			num1 := strings.Index(state, char1)
			num2 := strings.Index(state, char2)
			j := []byte(state)
			k := j[num1]
			j[num1] = j[num2]
			j[num2] = k
			state = string(j)
		default:
			log.Fatalf("%s does not match any move\n", s1)
		}
	}
	return state
}
