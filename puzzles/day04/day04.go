package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gammaflauge/aoc2017/utils"
)

func parseInput(input string) [][]string {
	var rows [][]string
	for _, line := range strings.Split(input, "\n") {
		rows = append(rows, strings.Split(line, " "))
	}
	return rows
}

func main() {
	inputFile := "input/day04.txt"
	rawInput := utils.ReadInputFile(inputFile)
	input := parseInput(rawInput)
	fmt.Println("******* Part 1 ********")
	validCount := 0
	for _, passphrase := range input {
		m := make(map[string]int)
		valid := 1
		for _, word := range passphrase {
			if m[word] == 1 {
				valid = 0
				break
			}
			m[word] = 1
		}
		validCount += valid
	}
	fmt.Println(validCount)

	fmt.Println("******* Part 2 ********")
	validCount = 0
	for _, passphrase := range input {
		m := make(map[string]int)
		valid := 1
		for _, word := range passphrase {
			chars := strings.Split(word, "")
			sort.Strings(chars)
			sortedWord := strings.Join(chars, "")
			if m[sortedWord] == 1 {
				valid = 0
				break
			}
			m[sortedWord] = 1
		}
		validCount += valid
	}
	fmt.Println(validCount)
}
