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

		if len(line) == 0 {
			break
		}

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

func getSortedIds(scanner *bufio.Scanner) []int {
	ids := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		id, _ := strconv.Atoi(line)

		ids = append(ids, id)
	}

	slices.Sort(ids)

	return ids
}

func main() {
	file, _ := os.Open("5a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ivs := getSortedIntervals(scanner)
	ids := getSortedIds(scanner)

	avail := 0
	ivIndex := 0
	idIndex := 0

	for idIndex < len(ids) && ivIndex < len(ivs) {
		id := ids[idIndex]
		iv := ivs[ivIndex]

		if id >= iv.Lower && id <= iv.Upper {
			avail++
			idIndex++
		} else if id < iv.Lower {
			idIndex++
		} else if id > iv.Upper {
			ivIndex++
		}
	}

	fmt.Println(avail)
}
