package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	// "strconv"
	// "sort"
	// "math"
)

func LinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

const (
	Occupied = '#'
	Empty    = 'L'
	Floor    = '.'
)

type Pos struct {
	x int
	y int
}

func GenLayout(lines []string) map[Pos]rune {
	layout := map[Pos]rune{}

	for i, line := range lines {
		for j, char := range line {
			layout[Pos{ x: i, y: j}] = char
		}
	}

	return layout
}

func IterateLayout(layout map[Pos]rune) (map[Pos]rune, int, int) {
	nayout := map[Pos]rune{}
	diff := 0
	occ  := 0

	for key, seat := range layout {
		occCount := 0

		for _, keys := range [][]int{
			{ 0,  1}, {0, -1}, { 1, 1}, {1,  0},
			{-1, -1}, {-1, 0}, {-1, 1}, {1, -1},
		} {
			s, ok := layout[Pos{x:keys[0] + key.x, y:keys[1] + key.y}]
			if ok && s == Occupied {
				occCount++
			}
		}

		if seat == Empty && occCount == 0 {
			nayout[key] = Occupied; diff++; occ++
		} else if seat == Occupied && occCount >= 4 {
			nayout[key] = Empty; diff++
		} else {
			nayout[key] = seat
		}

		if seat == Floor {
			nayout[key] = Floor
		}
	}

	return nayout, diff, occ
}

func IterateRayLayout(layout map[Pos]rune) (map[Pos]rune, int, int) {
	nayout := map[Pos]rune{}
	diff := 0
	occ  := 0

	for key, seat := range layout {
		occCount := 0

		for _, keys := range [][]int{
			{ 0,  1}, {0, -1}, { 1, 1}, {1,  0},
			{-1, -1}, {-1, 0}, {-1, 1}, {1, -1},
		} {
			pos := Pos{x:key.x, y:key.y}
			for {
				pos.x += keys[0]
				pos.y += keys[1]

				s, ok := layout[pos]
				if !ok || s == Empty {; break; }
				if s == Floor {; continue; }
				if s == Occupied {
					occCount++; break
				}
			}
		}

		if seat == Empty && occCount == 0 {
			nayout[key] = Occupied; diff++; occ++
		} else if seat == Occupied && occCount >= 5 {
			nayout[key] = Empty; diff++
		} else {
			nayout[key] = seat
		}

		if seat == Floor {
			nayout[key] = Floor
		}
	}

	return nayout, diff, occ
}

func OccupiedSeats(layout map[Pos]rune, f func(map[Pos]rune) (map[Pos]rune, int, int)) int {
	occupied := 0

	for {
		diff   := 0
		newOcc := 0
		layout, diff, newOcc = f(layout)

		if diff == 0 {; break; }

		occupied += newOcc - (diff - newOcc)
	}

	return occupied
}

func main() {
	fmt.Println("Day 11")

	fmt.Println("Part one: ", OccupiedSeats(
		GenLayout(LinesInFile("input")), IterateLayout,
	))

	fmt.Println("Part two: ", OccupiedSeats(
		GenLayout(LinesInFile("input")), IterateRayLayout,
	))
}
