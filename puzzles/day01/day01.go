package main

import (
	"fmt"

	"github.com/gammaflauge/aoc2017/utils"
)

func captcha(input string) (sum int) {
	sum = 0
	for i, c := range input {
		nextIndex := i + 1
		if nextIndex == len(input) {
			nextIndex = 0
		}
		next := input[nextIndex]
		if c == rune(next) {
			sum += int(rune(next) - '0')
		}
	}
	return sum
}

func captchaV2(input string) (sum int) {
	sum = 0
	for i, c := range input {
		nextIndex := i + len(input)/2
		if nextIndex >= len(input) {
			nextIndex -= len(input)
		}
		next := input[nextIndex]
		if c == rune(next) {
			sum += int(rune(next) - '0')
		}
	}

	return sum
}

func main() {
	inputFile := "./input/day01.txt"
	input := utils.ReadInputFile(inputFile)

	fmt.Println("**** part 1 solution ****")
	fmt.Println("sum =", captcha(input))
	fmt.Println()

	fmt.Println("**** part 2 solution ****")
	fmt.Println("sum =", captchaV2(input))
}
