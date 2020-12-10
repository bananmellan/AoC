package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"sort"
	"math"
)

func Perm(r int) int {; if r >= 1 {; return r * Perm(r - 1) ;}; return 1;}
func Comb(n int, r int) int {; return Perm(n)/(Perm(r)*Perm(n - r)); }

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

func EffectiveCount(arr []int) int {
	product := 1

	sums := make([]int, 12)
	sums[0] = 1
	for i := 0; i < 2; {
		sum := int(math.Pow(2, float64(i))) + 1
		for j := 0; j <= i; j++ {
			if j % 2 == 1 {
				sum -= i - 2
			}
		}

		i++; sums[i] = sum
	}

	for i := 3; i < len(sums); i++ {
		sums[i] = sums[i - 3] + sums[i - 2] + sums[i - 1]
	}

	consecutive := 0
	for i := 0; i < len(arr) - 2; i++ {
		diff := arr[i + 2] - arr[i]

		if diff <= 3 {
			if diff == 2 {
				consecutive++
			} else {
				product *= 2
			}
		} else if consecutive > 0 {
			product *= sums[consecutive]
			consecutive = 0
		}
	}

	return product
}

func main() {
	fmt.Println("Day 10")
	lines := LinesInFile("input")

	max := 0
	jolts := make([]int, len(lines) + 2)
	for i, line := range lines {
		jolts[i], _ = strconv.Atoi(line)

		if max < jolts[i] {
			max = jolts[i]
		}
	}

	jolts[len(jolts) - 2] = 0
	jolts[len(jolts) - 1] = max + 3
	sort.Ints(jolts)

	diffs := map[int]int{}
	for i := 0; i < len(jolts) - 1; i++ {
		diffs[jolts[i+1] - jolts[i]]++
	}

	fmt.Println("Part one: ", diffs[1], "*", diffs[3], "=", diffs[1] * diffs[3])
	fmt.Println("Part two: ", EffectiveCount(jolts))
}
