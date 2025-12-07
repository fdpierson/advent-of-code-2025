package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("6b.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	width := 0
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if len(line) > width {
			width = len(line)
		}
	}

	height := len(lines)
	trans := make([][]byte, width)

	for i := 0; i < height; i++ {
		for j := 0; j < len(lines[i]); j++ {
			trans[j] = append(trans[j], lines[i][j])
		}
	}

	oper := byte(' ')
	inner := 0
	sum := 0

	for _, arr := range trans {
		line := strings.TrimSpace(string(arr))

		if line == "" {
			sum += inner
			inner = 0
			continue
		}

		last := line[len(line) - 1]

		if last == byte('+') || last == byte('*') {
			line = line[:len(line) - 1]
			inner, _ = strconv.Atoi(strings.TrimSpace(line))
			oper = last
			continue
		}

		num, _ := strconv.Atoi(line)

		if oper == byte('+') {
			inner += num
		} else {
			inner *= num
		}
	}

	fmt.Println(sum + inner)
}
