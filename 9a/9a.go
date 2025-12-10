package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("9a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := make([]Point, 0)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		points = append(points, Point{x, y})
	}

	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pi := points[i]
			pj := points[j]

			dx := math.Abs(float64(pj.x - pi.x)) + 1.0
			dy := math.Abs(float64(pj.y - pi.y)) + 1.0

			area := int(dx * dy)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Println(maxArea)
}
