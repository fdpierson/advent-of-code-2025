// AI-generated version that uses Kahn's algorithm and dynamic programming
// Algorithm assumes that the graph is a directed acyclic graph, which the AoC graph is,
// even if it didn't actually imply that.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Topological sort (Kahn's Algorithm)
func topoSort(graph map[string][]string) []string {
	indegree := map[string]int{}
	for u := range graph {
		// ensure every key exists even if no outgoing edges
		if _, ok := indegree[u]; !ok {
			indegree[u] = 0
		}
		for _, v := range graph[u] {
			indegree[v]++
		}
	}

	queue := []string{}
	for node, deg := range indegree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}

	order := []string{}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)

		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return order
}

// Count number of unique paths from start → end in a DAG
func countPaths(graph map[string][]string, start, end string) int {
	topo := topoSort(graph)

	// dp[node] = number of ways to reach node from start
	dp := map[string]int{}
	dp[start] = 1

	for _, u := range topo {
		for _, v := range graph[u] {
			dp[v] += dp[u]
		}
	}

	return dp[end]
}

func main() {
	file, _ := os.Open("11b.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	forwardGraph := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		key := strings.TrimSuffix(fields[0], ":")

		for _, field := range fields[1:] {
			forwardGraph[key] = append(forwardGraph[key], field)
		}
	}

	// DAG path counts
	svrFftNumPaths := countPaths(forwardGraph, "svr", "fft")
	fmt.Println("svrFftNumPaths", svrFftNumPaths)

	fftDacNumPaths := countPaths(forwardGraph, "fft", "dac")
	fmt.Println("fftDacNumPaths", fftDacNumPaths)

	dacOutNumPaths := countPaths(forwardGraph, "dac", "out")
	fmt.Println("dacOutNumPaths", dacOutNumPaths)

	fmt.Println("totalPaths", svrFftNumPaths*fftDacNumPaths*dacOutNumPaths)
}
