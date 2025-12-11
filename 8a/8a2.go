package main

import (
	"container/heap"
	"bufio"
	"fmt"
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
	d int
}

// TODO: implement this ourselves
// --- BEGIN AI CODE ---

// Bounded max-heap

const K = 1000

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }

// MAX-heap: largest d at root
func (h PairHeap) Less(i, j int) bool {
	return h[i].d > h[j].d
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *PairHeap) Insert(p Pair) {
	if h.Len() < K {
		heap.Push(h, p)
		return
	}

	// Root holds the worst of the kept pairs
	if p.d < (*h)[0].d {
		(*h)[0] = p
		heap.Fix(h, 0)
	}
	// else: discard
}

// --- END AI CODE ---

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

	h := &PairHeap{}
	heap.Init(h)

	// Horrible O(n^2 log n) algorithm
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]

			x2 := (b.x - a.x) * (b.x - a.x)
			y2 := (b.y - a.y) * (b.y - a.y)
			z2 := (b.z - a.z) * (b.z - a.z)

			d := x2 + y2 + z2

			h.Insert(Pair{a, b, d})
		}
	}

	fmt.Printf("Second part: %v\n", time.Since(start))
	start = time.Now()

	pairs := make([]Pair, 0)

	for range 1000 {
		pairs = append(pairs, heap.Pop(h).(Pair))
	}

	fmt.Printf("Third part: %v\n", time.Since(start))
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

	fmt.Printf("Fourth part: %v\n", time.Since(start))

	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}
