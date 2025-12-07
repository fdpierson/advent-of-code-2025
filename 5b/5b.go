package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	Lower int
	Upper int
}

func getSortedIntervals(scanner *bufio.Scanner) []Interval {
	ivs := make([]Interval, 0)

	for scanner.Scan() {
		line := scanner.Text()

		bounds := strings.Split(line, "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])

		ivs = append(ivs, Interval{lower, upper})
	}

	slices.SortFunc(ivs, func(a, b Interval) int {
		return a.Lower - b.Lower
	})

	return ivs
}

func main() {
	file, _ := os.Open("5b.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ivs := getSortedIntervals(scanner)
	all := 0

	for i := 0; i < len(ivs); {
		minLower := ivs[i].Lower
		maxUpper := ivs[i].Upper

		for i += 1; i < len(ivs); i++ {
			if maxUpper < ivs[i].Lower {
				break
			} else if maxUpper < ivs[i].Upper {
				maxUpper = ivs[i].Upper
			}
		}

		all += maxUpper - minLower + 1
	}

	fmt.Println(all)
}
