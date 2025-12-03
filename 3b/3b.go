package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, _ := os.Open("3b.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		arr := [12]rune{}

		for li, lc := range line {
			for ai, ac := range arr {
				if lc > ac && li < len(line)+ai-11 {
					arr[ai] = lc
					for i := ai + 1; i < 12; i++ {
						arr[i] = '0'
					}
					break
				}
			}
		}

		for i, c := range arr {
			sum += int(math.Pow10(11-i)) * int(c-'0')
		}
	}

	fmt.Println(sum)
}
