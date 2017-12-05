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
	puzzleInput = "day04/input.txt"
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

	part1(bufio.NewScanner(file))
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	part2(bufio.NewScanner(file))
}

func part1(scanner *bufio.Scanner) {
	numValid := 0
lineloop:
	for scanner.Scan() {
		line := scanner.Text()
		wordMap := map[string]bool{}
		words := strings.Split(line, " ")
		for _, word := range words {
			if _, ok := wordMap[word]; ok {
				continue lineloop
			}
			wordMap[word] = true
		}
		numValid++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(numValid)
}

func part2(scanner *bufio.Scanner) {
	numValid := 0
lineloop:
	for scanner.Scan() {
		line := scanner.Text()
		wordMap := map[string]bool{}
		words := strings.Split(line, " ")
		for _, word := range words {
			chars := make([]int, 26)
			for _, c := range word {
				chars[byte(c)-'a']++
			}
			s := make([]string, 0, 26)
			for _, i := range chars {
				s = append(s, strconv.Itoa(i))
			}
			k := strings.Join(s, ",")
			if _, ok := wordMap[k]; ok {
				continue lineloop
			}
			wordMap[k] = true
		}
		numValid++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(numValid)
}
