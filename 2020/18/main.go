package main

import "fmt"
import "bufio"
import "os"
import "log"
// import "regexp"
import "strconv"
import "strings"

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

func simplify(chars []string, operator string) []string {
	for i, char := range chars {
		if char == operator {
			replace := i - 1
			a, _ := strconv.Atoi(chars[replace])
			b, _ := strconv.Atoi(chars[replace + 2])
			new := []string{}
			new = append(chars[:replace], chars[replace + 2:]...)
			sum := ""
			if operator == "+" {
				sum = strconv.Itoa(a + b)
			} else {
				sum = strconv.Itoa(a * b)
			}

			new[replace] = sum

			return new
		}
	}

	return []string{}
}

func Eval2(ex []rune) int {
	result := 0

	// simplifyplus(ex)

	return result
}

func main() {
	sum := 0

	for _, line := range LinesInFile("input") {
		sum += Eval([]rune(line))
	}

	fmt.Println("Part 1:", sum)

	sum = 0
	for _, line := range LinesInFile("input") {
		lastOpening := 0
		i := 0

		for {
			if i >= len(line) {
				if lastOpening == 0 {
					break
				}
				lastOpening = 0
				i = 0
			}
			char := line[i]

			if char == '(' {
				lastOpening = i
			} else if char == ')' {
				tmp := line[lastOpening + 1:]
				tmp = tmp[:i - lastOpening - 1]

				new := strings.Split(tmp, " ")

				fmt.Println(tmp)

				for _, op := range []string{"+", "*"} { for {
					result := simplify(new, op)

					if len(result) == 0 {
						break
					}

					new = result
				}}

				line = line[:lastOpening] + new[0] + line[i + 1:]

				fmt.Println(line)
				lastOpening = 0
				i = 0
			}

			i++
		}

		new := strings.Split(line, " ")
		for _, op := range []string{"+", "*"} { for {
			result := simplify(new, op)

			if len(result) == 0 {
				break
			}

			new = result
		}}


		num, _ := strconv.Atoi(new[0])
		sum += num
	}

	fmt.Println("Part 2: ", sum)
}
