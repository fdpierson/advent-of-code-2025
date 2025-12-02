package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func intPow10(i int) int {
	return int(math.Pow(10.0, float64(i)))
}

func intConcat(i, n int) int {
	s, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(i), n))
	return s
}

func getInvIdSeq(idStr string, id, n int) int {
	seqLen := len(idStr) / n

	if len(idStr) < n {
		return 1
	} else if len(idStr)%n != 0 {
		return intPow10(seqLen)
	}

	seq, _ := strconv.Atoi(idStr[:seqLen])
	invId := intConcat(seq, n)

	if id <= invId {
		return seq
	} else {
		return seq + 1
	}
}

func main() {
	file, _ := os.Open("2b.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	intervals := strings.Split(line, ",")

	invIds := map[int]bool{}

	for _, interval := range intervals {
		bounds := strings.Split(interval, "-")
		lower, _ := strconv.Atoi(bounds[0])
		upper, _ := strconv.Atoi(bounds[1])

		for n := 2; n <= len(bounds[1]); n++ {
			seq := getInvIdSeq(bounds[0], lower, n)
			id := intConcat(seq, n)

			for id <= upper {
				invIds[id] = true
				seq++
				id = intConcat(seq, n)
			}
		}
	}

	sum := 0

	for id := range invIds {
		sum += id
	}

	fmt.Println(sum)
}
