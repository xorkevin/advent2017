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
	puzzleInput = "day08/input.txt"
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

	registers := map[string]int{}

	procMax := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "if")
		s0 := strings.Split(strings.Trim(s[0], " "), " ")
		word := s0[1]
		register := s0[0]
		amount, err := strconv.Atoi(s0[2])
		if err != nil {
			log.Fatal(err)
		}

		s1 := strings.Split(strings.Trim(s[1], " "), " ")
		condition := s1[1]
		condRegister := s1[0]
		condAmount, err := strconv.Atoi(s1[2])
		if err != nil {
			log.Fatal(err)
		}

		if _, ok := registers[register]; !ok {
			registers[register] = 0
		}
		if _, ok := registers[condRegister]; !ok {
			registers[condRegister] = 0
		}

		if checkCondition(registers[condRegister], condition, condAmount) {
			switch word {
			case "inc":
				registers[register] += amount
			case "dec":
				registers[register] -= amount
			default:
				log.Fatal("word not matched: ", word)
			}
		}

		if registers[register] > procMax {
			procMax = registers[register]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	max := -99999
	for _, v := range registers {
		if v > max {
			max = v
		}
	}

	fmt.Println(max)
	fmt.Println(procMax)

	//part1()
	//part2()
}

func checkCondition(r int, condition string, v int) bool {
	switch condition {
	case "==":
		return r == v
	case "!=":
		return r != v
	case ">":
		return r > v
	case "<":
		return r < v
	case ">=":
		return r >= v
	case "<=":
		return r <= v
	}
	log.Fatal("condition not matched: ", condition)
	return false
}

func part1() {
	fmt.Println("Hello World")
}

func part2() {
}
