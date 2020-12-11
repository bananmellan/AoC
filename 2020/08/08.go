package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"regexp"
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
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

type Instruction struct {
	code string
	value int
	count int
}

func calcacc(instructions []Instruction) (int, int) {
	// Keep track of last jmp
	var acc = 0
	var linenr = 0
	var max = 0

	for {
		if (linenr > len(instructions) - 1) {
			break
		}

		instruction := &instructions[linenr]
		if instruction.count >= 1 {
			break;
		} else {
			instruction.count++
		}

		if max < linenr {
			max = linenr
		}

		switch instruction.code {
		case "acc":
			acc += instruction.value
			linenr++
		case "nop":
			linenr++
		case "jmp":
			linenr += instruction.value
		}
	}

	return acc, max
}

func main() {
	var instructions = []Instruction{}

	lines := LinesInFile("input")

	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		re := regexp.MustCompile(`([a-p]{3}) ([+|-][0-9]+)`)
		match := re.FindStringSubmatch(line)
		code := match[1]
		value, _ := strconv.Atoi(match[2])

		instructions = append(instructions, Instruction{
			code: code,
			value: value,
			count: 0,
		})
	}

	var clone = make([]Instruction, len(instructions))
	copy(clone, instructions)
	acc, line := calcacc(clone)
	fmt.Println("Part one: ", acc, line)

	line = 0;
	lastline := len(instructions) - 1

	for {
		if line == lastline + 1 {
			fmt.Println(line)
			panic("wtf")
		}

		var clone = make([]Instruction, len(instructions))
		copy(clone, instructions)

		instruction := &clone[line]

		if instruction.code == "nop" {
			instruction.code = "jmp"

		} else if instruction.code == "jmp" {
			instruction.code = "nop"
		}
		if instruction.code != "acc" {
			acc, linenr := calcacc(clone)
			if linenr >= lastline {
				fmt.Println("Part two: ", acc, linenr)
				break;
			}
		}

		line++
	}
}
