package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("6a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Fields(line))
	}

	bottom := len(lines) - 1
	right := len(lines[0]) - 1
	sum := 0

	for j := 0; j <= right; j++ {
		inner, _ := strconv.Atoi(lines[0][j])

		for i := 1; i < bottom; i++ {
			num, _ := strconv.Atoi(lines[i][j])

			if lines[bottom][j] == "+" {
				inner += num
			} else {
				inner *= num
			}
		}

		sum += inner
	}

	fmt.Println(sum)
}
