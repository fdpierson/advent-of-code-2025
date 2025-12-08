package main

import (
	"bufio"
	"cmp"
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
	z int
}

type Pair struct {
	a Point
	b Point
	d float64
}

func main() {
	start := time.Now()

	file, _ := os.Open("8a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := make([]Point, 0)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		points = append(points, Point{x, y, z})
	}

	fmt.Printf("First part: %v\n", time.Since(start))
	start = time.Now()

	pairs := make([]Pair, 0)

	// Horrible O(n^2 log n) algorithm
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]

			x2 := (b.x - a.x) * (b.x - a.x)
			y2 := (b.y - a.y) * (b.y - a.y)
			z2 := (b.z - a.z) * (b.z - a.z)

			d := math.Sqrt(float64(x2 + y2 + z2))

			pairs = append(pairs, Pair{a, b, d})
		}
	}

	slices.SortFunc(pairs, func(a, b Pair) int {
		return cmp.Compare(a.d, b.d)
	})

	fmt.Printf("Second part: %v\n", time.Since(start))
	start = time.Now()

	// I can't think of anything better than this
	p2c := map[Point]int{}
	c2p := map[int][]Point{}

	for i, point := range points {
		p2c[point] = i + 1
		c2p[i + 1] = append(c2p[i + 1], point)
	}

	for i := range 1000 {
		pair := pairs[i]

		ca := p2c[pair.a]
		cb := p2c[pair.b]

		if ca != cb {
			for _, point := range c2p[cb] {
				p2c[point] = ca
				c2p[ca] = append(c2p[ca], point)
			}

			delete(c2p, cb)
		}
	}

	sizes := make([]int, 0)

	for _, arr := range c2p {
		sizes = append(sizes, len(arr))
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)

	fmt.Printf("Third part: %v\n", time.Since(start))

	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}
