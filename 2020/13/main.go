package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	// "regexp"
	"strconv"
	// "sort"
	// "math"
	"strings"
	"math/big"
)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(big.NewInt(1)) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

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

func IsGoldStamp(busids []string, timestamp int) int {
	streak := 0

	for j := len(busids) - 1; j >= 0; j-- {
		id, _ := strconv.Atoi(busids[j])
		if busids[j] == "x" {; continue; }

		if (timestamp + j) % id != 0 {
			break
		} else {
			streak++
		}
	}

	return streak
}

func main() {
	lines := LinesInFile("input")
	timestamp, _ := strconv.Atoi(lines[0])
	min := -1
	busid := 0
	busids := strings.Split(lines[1], ",")
	exes := 0

	for _, mod := range busids {
		if mod == "x" {
			exes++
			continue
		}
		conv, _ := strconv.Atoi(mod)
		val := conv - (timestamp % conv)
		if val < min || min == -1 {
			min = val
			busid = conv
		}
	}

	fmt.Println("Part one: ", min * busid)

	busid, _ = strconv.Atoi(busids[0])
	n := make([]*big.Int, len(busids) - exes)
	a := make([]*big.Int, len(busids) - exes)
	shift := 0
	for i := 0; i < len(busids); i++ {
		if busids[i] == "x" {; shift++; continue; }

		id,_ := strconv.Atoi(busids[i])
		n[i - shift] = big.NewInt(int64(id))
		a[i - shift] = big.NewInt(int64(id - i))
	}

	ans,_ := crt(a, n)
	fmt.Println("Part two: ", ans)
}
