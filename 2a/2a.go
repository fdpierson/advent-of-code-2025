package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func intLog10(a int) int {
	return int(math.Log10(float64(a)))
}

func intPow10(a int) int {
	return int(math.Pow(10.0, float64(a)))
}

func getHalves(id int) (int, int) {
	pow := (intLog10(id) + 1) / 2

	firstHalf := id / intPow10(pow)
	secondHalf := id % intPow10(pow)

	return firstHalf, secondHalf
}

func makeInvalidId(firstHalf int) int {
	nextPow10 := intPow10(intLog10(firstHalf) + 1)
	return firstHalf*nextPow10 + firstHalf
}

func nextInvalidId(id int) int {
	idLog10 := intLog10(id)
	firstHalf := 0

	if idLog10%2 == 0 {
		firstHalf = intPow10(idLog10 / 2)
	} else {
		secondHalf := 0
		firstHalf, secondHalf = getHalves(id)
		if firstHalf < secondHalf {
			firstHalf++
		}
	}

	return makeInvalidId(firstHalf)
}

func prevInvalidId(id int) int {
	idLog10 := intLog10(id)
	firstHalf := 0

	if id <= 11 {
		return 11
	} else if idLog10%2 == 0 {
		firstHalf = intPow10(idLog10/2) - 1
	} else {
		secondHalf := 0
		firstHalf, secondHalf = getHalves(id)
		if firstHalf > secondHalf {
			firstHalf--
		}
	}

	return makeInvalidId(firstHalf)
}

func main() {
	file, _ := os.Open("2a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	intervals := strings.Split(line, ",")

	sum := 0

	for _, interval := range intervals {
		bounds := strings.Split(interval, "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])

		firstInvId := nextInvalidId(lower)
		lastInvId := prevInvalidId(upper)
		id := firstInvId

		for id <= lastInvId {
			sum += id
			id = nextInvalidId(id + 1)
		}
	}

	fmt.Println(sum)
}
