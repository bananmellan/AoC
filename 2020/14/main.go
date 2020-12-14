package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"regexp"
	"strconv"
	// "sort"
	// "math"
	// "strings"
	// "math/big"
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

func main() {
	lines := LinesInFile("input")
	mem := map[int]int{}
	var zermask int
	var onemask int

	for _, line := range lines {
		re := regexp.MustCompile("([a-z]+).*")
		switch re.FindStringSubmatch(line)[1] {
		case "mask":
			re = regexp.MustCompile("mask = ([1|0|X]+)")
			mask := re.FindStringSubmatch(line)[1]
			zermask = 0
			onemask = 0

			for n, x := range mask {
				if x == '1' || x == 'X' {; zermask += 1 << (len(mask) - n - 1); }
			}; for n, x := range mask {
				if x == '1' {; onemask += 1 << (len(mask) - n - 1); }
			}
		case "mem":
			re = regexp.MustCompile("mem\\[([0-9]+)\\] = ([0-9]+)")
			match := re.FindStringSubmatch(line)
			addr, _ := strconv.Atoi(match[1])
			val, _ := strconv.Atoi(match[2])
			val |= onemask
			val &= zermask
			mem[addr] = val
		}
	}

	sum := 0
	for _, val := range mem {
		sum += val
	}

	fmt.Println("Part one: ", sum)

	mem = map[int]int{}
	var mask string

	for _, line := range lines {
		re := regexp.MustCompile("([a-z]+).*")
		switch re.FindStringSubmatch(line)[1] {
		case "mask":
			re = regexp.MustCompile("mask = ([1|0|X]+)")
			mask = re.FindStringSubmatch(line)[1]
		case "mem":
			re = regexp.MustCompile("mem\\[([0-9]+)\\] = ([0-9]+)")
			match := re.FindStringSubmatch(line)
			val, _ := strconv.Atoi(match[2])
			addr, _ := strconv.Atoi(match[1])
			baddrconv := []rune(strconv.FormatInt(int64(addr), 2))
			raddr := make([]rune, len(mask))
			baddr := make([]rune, len(mask))

			diff := len(mask) - len(baddrconv)
			for i, _ := range baddr {
				baddr[i] = '0'
			}; for i, bit := range baddrconv {
				baddr[i + diff] = bit
			}

			for i, bit := range mask {
				switch bit {
				case '0': raddr[i] = baddr[i]
				case '1': raddr[i] = '1'
				case 'X': raddr[i] = 'X'
				}
			}

			var comb func (bits []rune)
			comb = func (bits []rune) {
				for i, bit := range bits {
					if bit == 'X' {
						clone := make([]rune, len(bits))
						copy(clone, bits)
						clone[i] = '0'; comb(clone)
						clone = make([]rune, len(bits))
						copy(clone, bits)
						clone[i] = '1'; comb(clone)
						return
					}
				}

				addr, _ := strconv.ParseInt(string(bits), 2, 64)
				mem[int(addr)] = val
			}

			comb(raddr)
		}
	}

	sum = 0
	for _, val := range mem {
		sum += val
	}

	fmt.Println("Part two: ", sum)
}
