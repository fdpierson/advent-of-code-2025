package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input/1b.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dial := 50
	zero := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line[1:])

		if line[0] == 'L' {
			for range num {
				if dial%100 == 0 {
					zero++
				}
				dial--
			}
		} else if line[0] == 'R' {
			for range num {
				if dial%100 == 0 {
					zero++
				}
				dial++
			}
		}
	}

	fmt.Println(zero)
}
