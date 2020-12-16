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
	var myticket []int
	var tickets [][]int
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
	myticket = make([]int, length)
	for i, val := range split {
		conv,_ := strconv.Atoi(val)
		myticket[i] = conv
	}

	start := len(rules) + 5
	tickets = make([][]int, len(lines) - start)
	for i, _ := range tickets {; tickets[i] = make([]int, length); }
	for i := start; i < len(lines); i++ {
		for j, val := range strings.Split(lines[i], ",") {
			conv,_ := strconv.Atoi(val)
			tickets[i - start][j] = conv
		}
	}

	sum := 0
	for _, ticket := range tickets {
		for _, val := range ticket {
			flag := false

			for _, rule := range rules {
				if (val >= rule.intervals.fst.min &&
					val <= rule.intervals.fst.max) || (
					val >= rule.intervals.snd.min &&
					val <= rule.intervals.snd.max) {
					flag = true
					break
				}
			}

			if !flag {; sum += val; }
		}
	}

	fmt.Println("Part one: ", sum)
}
