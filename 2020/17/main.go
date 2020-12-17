package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func LinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {; log.Fatal(err); }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

type Cube struct {
	x int; y int; z int; w int
}

func main() {
	lines := LinesInFile("input")
	var cubes map[Cube]bool
	comb := []int{-1, 0, 1}
	adj  := []Cube{}
	adj4 := []Cube{}
	for _, x := range comb {
		for _, y := range comb {
			for _, z := range comb {
				if !(x == 0 && x == y && y == z) {
					adj = append(adj, Cube{x: x, y: y, z: z})
				}


				for _, w := range comb {
					if !(x == 0 && x == y && y == z && z == w) {
						adj4 = append(adj4, Cube{x: x, y: y, z: z, w: w})
					}
				}
			}
		}
	}

	var buildInitial func()
	buildInitial = func() {
		cubes = map[Cube]bool{}

		// Build initial cubes
		for row, line := range lines {
			for col, rune := range line {
				if rune == '#' {
					cubes[Cube{ x: col, y: row }] = true
				}
			}
		}
	}

	var neighbors func(cube Cube, adja []Cube) []Cube
	neighbors = func(cube Cube, adja []Cube) []Cube {
		result := []Cube{}

		for _, ad := range adja {
			result = append(result, Cube{
				x: cube.x + ad.x,
				y: cube.y + ad.y,
				z: cube.z + ad.z,
				w: cube.w + ad.w,
			})
		}

		return result
	}

	var countAdjacents func(cube Cube, adja []Cube) int
	countAdjacents = func(cube Cube, adja []Cube) int {
		count := 0

		for _, neighbor := range neighbors(cube, adja) {
			_, ok := cubes[neighbor]
			if ok {; count++ ;}
		}

		return count
	}

	var iterate func(adj []Cube) int
	iterate = func(adja []Cube) int {
		buildInitial()

		for it := 0; it < 6; it++ {
			clone := map[Cube]bool{}

			for cube, _ := range cubes {
				count := countAdjacents(cube, adja)
				if count == 2 || count == 3 {
					clone[cube] = true
				}

				for _, neighbor := range neighbors(cube, adja) {
					_, ok := clone[neighbor]
					if !ok && countAdjacents(neighbor, adja) == 3 {
						clone[neighbor] = true
					}
				}
			}

			cubes = clone
		}

		return len(cubes)
	}

	fmt.Println("Part one:", iterate(adj))
	fmt.Println("Part two:", iterate(adj4))
}
