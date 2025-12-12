package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("12a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pixels := map[int][]string{}
	numPixels := map[int]int{}
	regions := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) <= 1 {
			index, _ := strconv.Atoi(strings.TrimSuffix(line, ":"))

			for scanner.Scan() {
				line = scanner.Text()

				if line == "" {
					break
				} else {
					pixels[index] = append(pixels[index], line)
					numPixels[index] += strings.Count(line, "#")
				}
			}
		} else {
			dim := strings.Split(fields[0], "x")
			x, _ := strconv.Atoi(dim[0])
			y, _ := strconv.Atoi(strings.TrimSuffix(dim[1], ":"))

			region := []int{x, y}

			for _, field := range fields[1:] {
				i, _ := strconv.Atoi(field)
				region = append(region, i)
			}

			regions = append(regions, region)
		}
	}

	triviallyPassing := 0
	triviallyNonpassing := 0

	for _, region := range regions {
		numShapesLimit := (region[0] / 3) * (region[1] / 3)
		numShapesTotal := 0

		numPixelsLimit := region[0] * region[1]
		numPixelsTotal := 0

		for index, numShapes := range region[2:] {
			numShapesTotal += numShapes
			numPixelsTotal += numShapes * numPixels[index]
		}

		if (numShapesTotal <= numShapesLimit) {
			triviallyPassing++
		} else if (numPixelsTotal > numPixelsLimit) {
			triviallyNonpassing++
		}
	}

	// Apparently we don't need to take into account combining shapes together to save space
	fmt.Println(triviallyPassing)
	fmt.Println(triviallyNonpassing)
	fmt.Println(len(regions))
}
