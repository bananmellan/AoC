package main

import "fmt"
import "bufio"
import "os"
import "log"
import "regexp"
import "strconv"
import "strings"

type Rule struct {
	char rune
	daddy *Rule
	rulesleft []Rule
	rulesright []Rule
	passed bool
	step int
}

func BuildRules(lines []string) (map[int]Rule, int) {
	rules := map[int]Rule{}
	var row int
	var line string
	for row, line = range lines {
		var rule Rule = Rule{}

		if line == "" {; break; }

		rechar := regexp.MustCompile(`(\d+): "([a-z])"`)
		reseq := regexp.MustCompile(`(\d+): ([\d+ ]*)\|?([\d+ ]*)`)

		match := rechar.MatchString(line)
		matches := reseq.FindStringSubmatch(line)

		index := -1
		var char rune
		rulesleft := []Rule{}
		rulesright := []Rule{}

		if (match) {
			m := rechar.FindStringSubmatch(line)
			index, _ = strconv.Atoi(m[1])
			char = rune(m[2][0])

			rules[index] = rule
			rule.char = char
		} else {
			index, _ = strconv.Atoi(matches[1])
			nums1 := strings.Split(matches[2], " ")
			nums2 := strings.Split(matches[3], " ")
			for i := 0; i < len(nums1) - 1; i++ {
				num, _ := strconv.Atoi(nums1[i])
				rulesleft = append(rulesleft, Rule{
					char: rune(num),
					daddy: &rule,
				})
			}

			for i := 1; i < len(nums2); i++ {
				num, _ := strconv.Atoi(nums2[i])
				rulesright = append(rulesright, Rule{
					char: rune(num),
					daddy: &rule,
				})
			}

			rules[index] = rule
			rule.rulesleft = rulesleft
			rule.rulesright = rulesright
		}
	}

	return rules, row
}

func main() {
	lines := LinesInFile("input")
	var row int
	rules, row := BuildRules(lines)

	rows := lines[row + 1:]

	fmt.Println(rules)

	rownr := 0
	for id, _ := range rules {
		cursor := 0
		row := &rows[rownr]
		rule := rules[id]

		for {
			char := rune((*row)[cursor])

			countleft := len(rule.rulesleft)
			countright := len(rule.rulesright)
			count := countleft + countright
			countpassed := 0
			for _, r := range append(rule.rulesleft, rule.rulesright...) {
				if r.passed {; countpassed++; }
			}

			if count == rule.step {
				if count == countpassed {
					rule.passed = true
				}; if rule.daddy != nil {
					rule = *rule.daddy
				}
			}

			if rule.char != 0 {
				if char == rule.char {
					cursor += 1
					rule.passed = true
				} else {
					cursor -= rule.step
					if countright > 0 {
						rule = rule.rulesright[rule.step - countleft]
					} else {
						rule = *rule.daddy
					}
				}
			} else {
				rule.step++
				rule = rule.rulesleft[rule.step - 1]
			}
		}

		rownr++
	}
}

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
