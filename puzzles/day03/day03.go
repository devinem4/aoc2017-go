package main

import (
	"fmt"

	"github.com/gammaflauge/aoc2017/utils"
)

var puzzleInput = 361527

type Point struct {
	x, y int
}

type Square struct {
	x, y, value int
}

type Dir struct {
	dx, dy int
}

// change direction of next point when necessary
// top right or bottom left of spiral
func updateDir(d Dir, curr Point) Dir {
	if curr.x == curr.y ||
		// top left of spiral
		(curr.x < 0 && -curr.x == curr.y) ||
		// expanding spiral outward
		(curr.x > 0 && curr.x == -curr.y+1) {
		if d.dx != 0 {
			d.dy = d.dx
			d.dx = 0
			return d
		}
		d.dx = -d.dy
		d.dy = 0
	}
	return d
}

func getNeighborSum(m map[Point]Square, curr Point) int {
	value := 0
	for neighborX := -1; neighborX <= 1; neighborX++ {
		for neighborY := -1; neighborY <= 1; neighborY++ {
			neighbor := Point{curr.x + neighborX, curr.y + neighborY}
			if sq, ok := m[neighbor]; ok {
				value += sq.value
			}
		}
	}
	return value
}

func main() {
	m := make(map[int]Point)
	curr := Point{0, 0}
	dir := Dir{1, 0}
	m[1] = curr

	for i := 2; i <= puzzleInput; i++ {
		curr.x += dir.dx
		curr.y += dir.dy
		m[i] = Point{curr.x, curr.y}

		dir = updateDir(dir, curr)
	}

	fmt.Println("******* Part 1 ********")
	finalPoint := m[puzzleInput]
	fmt.Println(utils.CalcManhattanDist(finalPoint.x, finalPoint.y))

	fmt.Println("******* Part 2 ********")
	m2 := make(map[Point]Square)
	curr = Point{0, 0}
	dir = Dir{1, 0}
	m2[curr] = Square{0, 0, 1}

	for i := 2; i <= 250; i++ {
		curr.x += dir.dx
		curr.y += dir.dy
		value := getNeighborSum(m2, curr)

		m2[curr] = Square{curr.x, curr.y, value}
		if value > puzzleInput {
			break
		}
		dir = updateDir(dir, curr)
	}
	fmt.Println(m2[curr].value)
}
