package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gammaflauge/aoc2017/utils"
)

func parseInput(input string) [][]int {
	var rows [][]int
	for _, line := range strings.Split(input, "\n") {
		var nums []int
		for _, num := range strings.Split(line, "\t") {
			num, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		rows = append(rows, nums)
	}
	return rows
}

// getChecksum finds the min and max value from each line in the input,
// then sums the difference (max - min) of all lines
func getChecksum(input [][]int) (checksum int) {
	checksum = 0
	for _, line := range input {
		min := line[0]
		max := line[0]
		for _, num := range line[1:] {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		checksum += max - min
	}
	return
}

// getChecksumV2 finds the two numbers that evenly divide,
// then sums those divisions (max - min) of all lines
func getChecksumV2(input [][]int) (checksum int) {
	checksum = 0
	for _, line := range input {
		for i, num := range line {
			for _, newnum := range line[i+1:] {
				if newnum%num == 0 || num%newnum == 0 {
					if newnum > num {
						checksum += newnum / num
					} else {
						checksum += num / newnum
					}
				}
			}
		}
	}
	return
}

func main() {
	inputFile := "./input/day02.txt"
	rawInput := utils.ReadInputFile(inputFile)
	input := parseInput(rawInput)

	fmt.Println("******* Part 1 ********")
	fmt.Println(getChecksum(input))

	fmt.Println("******* Part 2 ********")
	fmt.Println(getChecksumV2(input))
}
