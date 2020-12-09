package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	// "regexp"
	"strconv"
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

func main() {
	fmt.Println("Day 8")
	lines := LinesInFile("input")

	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	seek := 25
	num := 0
	invalidnum := 0
	for i := seek; i < len(numbers); i++ {
		num = numbers[i]
		valid := false
		for j := i - seek; j < i; j++ {
			for h := j; h < i; h++ {
				sum := numbers[j] + numbers[h]
				if numbers[i] == sum {
					valid = true
				}
			}
		}

		if !valid {
			fmt.Println("Part one: ", num)
			invalidnum = num
			break
		}
	}

	for i := 0; i < len(numbers); i++ {
		min := 0
		max := 0

		num := numbers[i]
		min = num
		max = num

		for j := i + 1; j < len(numbers); j++ {
			jnum := numbers[j]
			num += jnum

			if jnum < min {
				min = jnum
			} else if jnum > max {
				max = jnum
			}

			if num > invalidnum {
				break
			} else if num == invalidnum {
				fmt.Println("Part two: ", min + max)
				panic("Done!")
			}
		}
	}
}
