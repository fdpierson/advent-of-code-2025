package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("7b.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	prev := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()
		curr := map[int]int{}

		for i, c := range line {
			if prev[i] > 0 {
				if c == '^' {
					curr[i - 1] += prev[i]
					curr[i + 1] += prev[i]
				} else {
					curr[i] += prev[i]
				}
			} else if c == 'S' {
				curr[i] = 1
			}
		}

		prev = curr
		scanner.Scan()
	}

	sum := 0

	for _, i := range prev {
		sum += i
	}

	fmt.Println(sum)
}
