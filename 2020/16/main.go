package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"regexp"
	"strconv"
	"strings"
	// "sort"
	// "math"
	// "strings"
	// "math/big"
)

type Interval struct {
	min int
	max int
}

type IntervalPair struct {
	fst Interval
	snd Interval
}

type Rule struct {
	intervals IntervalPair
	key string
}

type ticket []int

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


func main() {
	lines := LinesInFile("input")
	var myticket ticket
	var tickets []ticket
	rules := []Rule{}

	for _, line := range lines {
		if (len(line) == 0) {; break; }

		reg := "([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)"
		re := regexp.MustCompile(reg)
		match := re.FindStringSubmatch(line)
		minOne,_ := strconv.Atoi(match[2]);
		maxOne,_ := strconv.Atoi(match[3]);
		minTwo,_ := strconv.Atoi(match[4]);
		maxTwo,_ := strconv.Atoi(match[5]);

		rules = append(rules, Rule{
			key: match[1],
			intervals: IntervalPair{
				fst: Interval{ min: minOne, max: maxOne, },
				snd: Interval{ min: minTwo, max: maxTwo, },
			},
		})
	}

	split := strings.Split(lines[len(rules) + 2], ",")
	length := len(split)
	myticket = make(ticket, length)
	for i, val := range split {
		conv,_ := strconv.Atoi(val)
		myticket[i] = conv
	}

	start := len(rules) + 5
	tickets = make([]ticket, len(lines) - start)
	for i, _ := range tickets {; tickets[i] = make(ticket, length); }
	for i := start; i < len(lines); i++ {
		for j, val := range strings.Split(lines[i], ",") {
			conv,_ := strconv.Atoi(val)
			tickets[i - start][j] = conv
		}
	}

	sum := 0
	validTickets := []ticket{}

	validate := func(val int, rule Rule) bool {
		return (val >= rule.intervals.fst.min &&
			val <= rule.intervals.fst.max) || (
			val >= rule.intervals.snd.min &&
				val <= rule.intervals.snd.max)
	}

	for _, ticket := range tickets {
		valid := true

		for _, val := range ticket {
			flag := false

			for _, rule := range rules {
				if validate(val, rule) {
					flag = true
					break
				}
			}

			if !flag {
				sum += val
				valid = false
				break
			}
		}

		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	fmt.Println("Part one: ", sum)

	candidates := map[int]map[Rule]bool{}
	for col := 0; col < length; col++ {
		candidates[col] = map[Rule]bool{}
		for _, rule := range rules {
			candidates[col][rule] = true
		}
	}

	for _, ticket := range validTickets {
		for i, val := range ticket {
			for _, rule := range rules {
				_, ok := candidates[i][rule]
				if ok && !validate(val, rule) {
					delete(candidates[i], rule)
				}
			}
		}
	}

	columns := make([]Rule, length)

	var FetchUnique func() (Rule, int)
	FetchUnique = func() (Rule, int) {
		var chosen Rule
		var column int
		for i, _ := range candidates {
			if len(candidates[i]) == 1 {
				column = i
				for rule := range candidates[i] {; chosen = rule ;}
				delete(candidates, i)
				for k, _ := range candidates {
					delete(candidates[k], chosen)
				}
				break
			}
		}

		return chosen, column
	}

	for i := 0; i < length; i++ {
		rule, col := FetchUnique()
		columns[col] = rule
	}

	for col, rule := range columns {
		fmt.Println(col, rule, myticket[col])
	}

	product := 1
	for col, val := range myticket {
		matched, _ := regexp.Match(`departure.*`, []byte(columns[col].key))
		if matched {; product *= val; }
	}

	fmt.Println("Part two: ", product)
}
