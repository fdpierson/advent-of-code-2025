package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2b.txt")
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

		// TODO: Implement the rest
	}

	fmt.Println(sum)
}
