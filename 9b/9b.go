package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
}

type Pair struct {
	a Point
	b Point
}

func getPoints() []Point {
	file, _ := os.Open("9b.txt")
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

	return append(points, points[0])
}

func getPairs(points []Point) []Pair {
	pairs := make([]Pair, 0)

	for i := 0; i < len(points) - 1; i++ {
		a := points[i]
		b := points[i + 1]

		pairs = append(pairs, Pair{a, b})
	}

	return pairs
}

func getHorizPairsSorted(pairs []Pair) []Pair {
	horiz := make([]Pair, 0)

	for _, pair := range pairs {
		if pair.a.y == pair.b.y {
			horiz = append(horiz, pair)
		}
	}

	slices.SortFunc(horiz, func(p1, p2 Pair) int {
		return p1.a.y - p2.a.y
	})

	return horiz
}

func getVertPairsSorted(pairs []Pair) []Pair {
	vert := make([]Pair, 0)

	for _, pair := range pairs {
		if pair.a.x == pair.b.x {
			vert = append(vert, pair)
		}
	}

	slices.SortFunc(vert, func(p1, p2 Pair) int {
		return p1.a.x - p2.a.x
	})

	return vert
}

func isBetween(a, i, b int) bool {
	if a <= b && a <= i && i <= b {
		return true
	} else if b <= a && b <= i && i <= a {
		return true
	} else {
		return false
	}
}

func isStrictlyBetween(a, i, b int) bool {
	if a < b && a < i && i < b {
		return true
	} else if b < a && b < i && i < a {
		return true
	} else {
		return false
	}
}

func getMaxArea(a, b Point, maxArea int) int {
	dx := math.Abs(float64(a.x - b.x)) + 1.0
	dy := math.Abs(float64(a.y - b.y)) + 1.0

	area := int(dx * dy)

	if area > maxArea {
		return area
	} else {
		return maxArea
	}
}

func main() {
	start := time.Now()

	points := getPoints()
	pairs := getPairs(points)

	hPairs := getHorizPairsSorted(pairs)
	vPairs := getVertPairsSorted(pairs)

	hInt := make(map[Point][]Pair)
	vInt := make(map[Point][]Pair)

	for i := 0; i < len(points) - 1; i++ {
		point := points[i]

		for _, hPair := range hPairs {
			if isBetween(hPair.a.x, point.x, hPair.b.x) {
				hInt[point] = append(hInt[point], hPair)
			}
		}

		for _, vPair := range vPairs {
			if isBetween(vPair.a.y, point.y, vPair.b.y) {
				vInt[point] = append(vInt[point], vPair)
			}
		}
	}

	maxArea := 0

	for i := 0; i < len(points) - 1; i++ {
		for j := i + 1; j < len(points) - 1; j++ {
			a := points[i]
			b := points[j]

			// There can't be any hPair intersections between (a.x, a.y) and (a.x, b.y)
			for _, hPair := range hInt[a] {
				if isStrictlyBetween(a.y, hPair.a.y, b.y) {
					if a.x == hPair.a.x && !isStrictlyBetween(a.x, hPair.b.x, b.x) {
						// do nothing
					} else if a.x == hPair.b.x && !isStrictlyBetween(a.x, hPair.a.x, b.x) {
						// do nothing
					} else {
						goto inter
					}
				}
			}

			// There can't be any hPair intersections between (b.x, a.y) and (b.x, b.y)
			for _, hPair := range hInt[b] {
				if isStrictlyBetween(a.y, hPair.a.y, b.y) {
					if a.x == hPair.a.x && !isStrictlyBetween(a.x, hPair.b.x, b.x) {
						// do nothing
					} else if a.x == hPair.b.x && !isStrictlyBetween(a.x, hPair.a.x, b.x) {
						// do nothing
					} else {
						goto inter
					}
				}
			}

			// There can't be any vPair intersections between (a.x, a.y) and (b.x, a.y)
			for _, vPair := range vInt[a] {
				if isStrictlyBetween(a.x, vPair.a.x, b.x) {
					if a.y == vPair.a.y && !isStrictlyBetween(a.y, vPair.b.y, b.y) {
						// do nothing
					} else if a.y == vPair.b.y && !isStrictlyBetween(a.y, vPair.a.y, b.y) {
						// do nothing
					} else {
						goto inter
					}
				}
			}

			// There can't be any vPair intersections between (a.x, b.y) and (b.x, b.y)
			for _, vPair := range vInt[b] {
				if isStrictlyBetween(a.x, vPair.a.x, b.x) {
					if a.y == vPair.a.y && !isStrictlyBetween(a.y, vPair.b.y, b.y) {
						// do nothing
					} else if a.y == vPair.b.y && !isStrictlyBetween(a.y, vPair.a.y, b.y) {
						// do nothing
					} else {
						goto inter
					}
				}
			}

			maxArea = getMaxArea(a, b, maxArea)

			inter: continue
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Time:", elapsed)
	fmt.Println("Max area:", maxArea)
}
