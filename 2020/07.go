package main

import "fmt"
import "bufio"
import "os"
import "log"
import "strings"
import "regexp"
import "strconv"

type Bag struct {
	color string
	bags []Bag
	quantity int
}

var rules []Bag = []Bag{}

func fetch(key string) Bag {
	for _, rulebag := range rules {
		if rulebag.color == key {
			return rulebag
		}
	}

	return Bag{}
}

func checkbagforbag(keyColor string, bag Bag) bool {

	// If currentColor is key, return true.
	if keyColor == bag.color {
		return true
	} else {
		if (len(bag.bags) > 0) {
			for _, sbag := range bag.bags {
				if checkbagforbag(keyColor, sbag) {
					return true
				}
			}
		} else {
			for _, rulebag := range rules {
				if rulebag.color == bag.color {
					if (len(rulebag.bags) > 0 ) {
						if checkbagforbag(keyColor, rulebag) {
							return true
						}
					}
				}
			}
		}

		return false
	}
}

func countbagsinbag(bag Bag) int {
	var count = 0

	if (len(bag.bags) > 0) {
		fmt.Println(bag)

		count += bag.quantity

		for _, sbag := range bag.bags {
			fmt.Println(sbag)
			count += bag.quantity * countbagsinbag(sbag)
		}
	} else {
		rulebag := fetch(bag.color)

		if (len(rulebag.bags) == 0) {
			count = bag.quantity
		} else {
			rulebag.quantity = bag.quantity
			count = countbagsinbag(rulebag)
		}
	}

	return count;
}

func main() {
	file, err := os.Open("input")

	// If file wasn't opened correctly, print error.
	if err != nil {
		log.Fatal(err)
	}

	// Close file at end of file execution.
	defer file.Close()

	// Iterate lines.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		split := strings.Split(scanner.Text(), " bags contain ");
		color := split[0]
		bags := split[1]

		bag := Bag{
			quantity: 1,
			color: color,
			bags: []Bag{},
		}

		for _, sbag := range strings.Split(bags, ", ") {
			re := regexp.MustCompile(`([0-9]+) ([A-Za-z]+ [A-Za-z]+) bags?.?`)

			if (sbag == "no other bags.") {
				continue;
			}

			match := re.FindStringSubmatch(sbag)
			quantity, _ := strconv.Atoi(match[1])

			color = match[2]
			bag.bags = append(bag.bags, Bag{
				quantity: quantity,
				color: color,
			})
		}

		rules = append(rules, bag)
	}

	// var count = -1;

	// for _, bag := range rules {
	//	if (checkbagforbag("shiny gold", bag)) {
	//		count++;
	//	}
	// }

	var shinygold Bag = Bag{}

	for _, sbag := range rules {
		if sbag.color == "shiny gold" {
			shinygold = sbag
		}
	}

	fmt.Println("Required number of bags in shiny gold bag: ",
		countbagsinbag(shinygold) - 1)

	// If scanner catches error, print it.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
