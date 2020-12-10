package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	// "regexp"
	"strconv"
	"sort"
	"crypto/sha256"
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

func perm(num int) int {
	if num >= 1 {
		return num * perm(num - 1)
	}

	return 1
}

func comb(n int, r int) int {
	return perm(n) / (perm(r) * perm(n - r))
}

func Hash(arr []int) [32]byte {
	b := make([]byte, len(arr));
	for i, num := range arr {
		b[i] = byte(num)
	}

	return sha256.Sum256(b)
}

var isCounted = map[[32]byte]bool{}

func ResetCount(arr []int) int {
	isCounted = map[[32]byte]bool{}
	return CountValidArrangements(arr)
}

func CountValidArrangements(arr []int) int {
	count := 1

	if _, exists := isCounted[Hash(arr)]; exists {
		return 0
	} else {
		isCounted[Hash(arr)] = true
	}

	for i := 0; i < len(arr) - 3; i++ {
		if arr[i + 2] - arr [i] <= 3 {
			narr := make([]int, len(arr) - 1)
			for j := 0; j < i + 1; j++ {
				narr[j] = arr[j]
			}

			for j := i + 1; j < len(narr); j++ {
				narr[j] = arr[j + 1]
			}

			count += CountValidArrangements(narr)
		}
	}

	return count
}

func SliceCount(arr []int, from int, length int) int {
	narr := make([]int, length)

	for j := from; j < from + length; j++ {
		narr[j - from] = arr[j]
	}

	return ResetCount(narr)
}

func EffectiveCount(arr []int) int {
	product := 1
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
			product *= SliceCount(arr, i - consecutive, consecutive + 3)
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
