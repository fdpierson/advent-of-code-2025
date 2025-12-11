package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("10a2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sumPresses := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		sLights := strings.Trim(fields[0], "[]")
		iLights := 0

		for i, c := range sLights {
			if c == '#' {
				iLights ^= 1 << i
			}
		}

		iButtons := make([]int, 0)

		for i := 1; i < len(fields) - 1; i++ {
			sButton := strings.Trim(fields[i], "()")
			sToggles := strings.Split(sButton, ",")
			iButton := 0

			for _, sToggle := range sToggles {
				iToggle, _ := strconv.Atoi(sToggle)
				iButton ^= 1 << iToggle
			}

			iButtons = append(iButtons, iButton)
		}

		minPresses := len(iButtons)

		for presses := range 1 << len(iButtons) {
			result := 0
			numPresses := 0

			for i, iButton := range iButtons {
				if (presses >> i) & 1 == 1 {
					result ^= iButton
					numPresses++
				}
			}

			if result == iLights && numPresses < minPresses {
				minPresses = numPresses
			}
		}

		sumPresses += minPresses
	}

	fmt.Println(sumPresses)
}
