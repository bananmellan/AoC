package main

import "fmt"
import "bufio"
import "os"
import "log"
// import "strconv"
// import "strings"

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

func Match(ex []rune) int {
	depth := 0

	for i, char := range ex {
		if char == '(' {
			depth++
		} else if char == ')' {
			depth--
		}

		if depth == 0 {
			return i
		}
	}

	return -1
}

func Eval(ex []rune) int {
	result := 0
	var operator rune

	for i := 0; i < len(ex); i++ {
		char := []rune(ex)[i]
		eval := int(char) - '0'
		if char == '(' {
			match := i + Match(ex[i:])
			eval = Eval(ex[i + 1:match])
			i = match + 1
		} else if char == ' ' {
			continue
		} else if eval < 0 {
			operator = char
		}

		if eval >= 0 {
			if operator == '*' {
				result *= eval
			} else {
				result += eval
			}
		}
	}

	return result
}

func main() {
	lines := LinesInFile("input")

	sum := 0

	for _, line := range lines {
		sum += Eval([]rune(line))
	}

	fmt.Println("Part 1:", sum)
}
