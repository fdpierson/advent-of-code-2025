package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("4b.txt")
	defer file.Close()

	var grid [138][138]int

	scanner := bufio.NewScanner(file)
	i := 1

	for scanner.Scan() {
		for j, c := range scanner.Text() {
			if c == '@' {
				grid[i][j+1] = 1
			} else {
				grid[i][j+1] = 0
			}
		}
		i++
	}

	rolls := 0

	for true {
		oldRolls := rolls

		for i = 1; i <= 136; i++ {
			for j := 1; j <= 136; j++ {
				if grid[i][j] != 1 {
					continue
				}

				adj := grid[i-1][j-1] +
					grid[i-1][j] +
					grid[i-1][j+1] +
					grid[i][j-1] +
					grid[i][j+1] +
					grid[i+1][j-1] +
					grid[i+1][j] +
					grid[i+1][j+1]

				if adj < 4 {
					grid[i][j] = 0
					rolls++
				}
			}
		}

		if oldRolls == rolls {
			break
		}
	}

	fmt.Println(rolls)
}
