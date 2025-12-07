package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("7a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	prev := map[int]bool{}
	splits := 0

	for scanner.Scan() {
		line := scanner.Text()
		curr := map[int]bool{}

		for i, c := range line {
			if prev[i] && c == '^' {
				curr[i - 1] = true
				curr[i + 1] = true
				splits++
			} else if prev[i] || c == 'S' {
				curr[i] = true
			}
		}

		prev = curr
		scanner.Scan()
	}

	fmt.Println(splits)
}
