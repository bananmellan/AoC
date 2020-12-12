package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"regexp"
	"strconv"
	// "sort"
	"math"
)

func LinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {; log.Fatal(err); }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {; continue; }
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

const NORTH   = 'N'
const WEST    = 'W'
const SOUTH   = 'S'
const EAST    = 'E'
const LEFT    = 'L'
const RIGHT   = 'R'
const FORWARD = 'F'

func Advance(x int, y int, command rune, param int) (int, int) {
	switch command {
		case NORTH: y += param
		case SOUTH: y -= param
		case EAST:  x += param
		case WEST:  x -= param
	}

	return x, y
}

func main() {
	x := 0
	y := 0
	dirs := []rune{EAST, SOUTH, WEST, NORTH}
	d := 0

	lines := LinesInFile("input")

	for _, line := range lines {
		re := regexp.MustCompile(`([A-Z])([0-9]+)`)
		match := re.FindStringSubmatch(line)
		command := rune(match[1][0])
		param, _ := strconv.Atoi(match[2])

		if command == LEFT {
			d = (d - param / 90) % 4
			if d < 0 {; d += 4; }
			command = dirs[d]
		} else if command == RIGHT {
			d = (d + param / 90) % 4
			command = dirs[d]
		} else if command == FORWARD {
			command = dirs[d]
			x, y = Advance(x, y, command, param)
		} else {
			x, y = Advance(x, y, command, param)
		}
	}

	fmt.Println("Part one: ", math.Abs(float64(x)) + math.Abs(float64(y)))

	x = 0
	y = 0
	wx := 10
	wy := 1

	for _, line := range lines {
		re := regexp.MustCompile(`([A-Z])([0-9]+)`)
		match := re.FindStringSubmatch(line)
		command := rune(match[1][0])
		param, _ := strconv.Atoi(match[2])
		if command == FORWARD {
			for i := 0; i < param; i++ {
				x, y = Advance(x, y, 'E', wx)
				x, y = Advance(x, y, 'N', wy)
			}
		} else if command == LEFT {
			for i := 0; i < param / 90; i++ {
				tempx := wx
				wx = -wy; wy = tempx
			}
		} else if command == RIGHT {
			for i := 0; i < param / 90; i++ {
				tempx := wx
				wx =  wy; wy = -tempx
			}
		} else {
			wx, wy = Advance(wx, wy, command, param)
		}
	}

	fmt.Println("Part two: ", math.Abs(float64(x)) + math.Abs(float64(y)))
}
