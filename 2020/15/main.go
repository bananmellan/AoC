package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
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

func NthNumber(n int) int {
	input := LinesInFile("input")[0]
	split := strings.Split(input, ",")
	numbs := make([]int, len(split))

	for i, num := range split {; numbs[i], _ = strconv.Atoi(num); }

	var history = map[int]int{}
	var Speak func(num int, i int) int
	Speak = func(num int, i int) int {
		n, ok := history[num]
		last := 0

		if ok {; last = i - n; }
		history[num] = i

		return last
	}

	var last int

	for i, num := range numbs {
		 last = Speak(num, i)
	}

	for i := len(numbs); i < n - 1; i++ {
		last = Speak(last, i)
	}

	return last
}

func main() {
	fmt.Println("Part one: ", NthNumber(2020))

	fmt.Println("Part two: ", NthNumber(30000000))
}
