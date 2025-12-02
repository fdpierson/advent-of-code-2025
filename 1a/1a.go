package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("1a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dial := 50
	zero := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line[1:])

		if line[0] == 'L' {
			dial -= num
		} else {
			dial += num
		}

		if dial%100 == 0 {
			zero++
		}
	}

	fmt.Println(zero)
}
