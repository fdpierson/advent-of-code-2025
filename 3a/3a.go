package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("3a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		c1 := '0'
		c2 := '0'

		for i, c := range line {
			if c > c1 && i < len(line)-1 {
				c1 = c
				c2 = '0'
			} else if c > c2 {
				c2 = c
			}
		}

		sum += int(c1-'0')*10 + int(c2-'0')
	}

	fmt.Println(sum)
}
