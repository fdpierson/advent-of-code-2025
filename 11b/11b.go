package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	a string
	b string
}

func dfs(graph map[string][]string, visitedPairs map[Pair]bool, deviceKey string, dir bool) {
	for _, outputKey := range graph[deviceKey] {
		ab := Pair{deviceKey, outputKey}
		ba := Pair{outputKey, deviceKey}

		if visitedPairs[ab] || visitedPairs[ba] {
			continue
		} else {
			if dir {
				visitedPairs[ab] = true
			} else {
				visitedPairs[ba] = true
			}

			dfs(graph, visitedPairs, outputKey, dir)
		}
	}
}

func dfs2(graph map[string][]string, oldVisitedKeys map[string]bool, deviceKey string, endKey string, numPaths int) int {
	visitedKeys := map[string]bool{}

	for key, val := range oldVisitedKeys {
		visitedKeys[key] = val
	}

	visitedKeys[deviceKey] = true

	for _, outputKey := range graph[deviceKey] {
		if outputKey == endKey {
			return numPaths + 1
		} else if visitedKeys[outputKey] {
			continue
		} else {
			numPaths = dfs2(graph, visitedKeys, outputKey, endKey, numPaths)
		}
	}

	return numPaths
}

func numPaths(forwardGraph map[string][]string, backwardGraph map[string][]string, startKey string, endKey string) int {
	forwardVisitedPairs := map[Pair]bool{}
	dfs(forwardGraph, forwardVisitedPairs, startKey, true)

	backwardVisitedPairs := map[Pair]bool{}
	dfs(backwardGraph, backwardVisitedPairs, endKey, false)

	sharedGraph := map[string][]string{}

	for pair, _ := range forwardVisitedPairs {
		if backwardVisitedPairs[pair] {
			sharedGraph[pair.a] = append(sharedGraph[pair.a], pair.b)
		}
	}

	return dfs2(sharedGraph, map[string]bool{}, startKey, endKey, 0)
}

func main() {
	file, _ := os.Open("11b.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	forwardGraph := map[string][]string{}
	backwardGraph := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		key := strings.TrimSuffix(fields[0], ":")

		for _, field := range fields[1:] {
			forwardGraph[key] = append(forwardGraph[key], field)
			backwardGraph[field] = append(backwardGraph[field], key)
		}
	}

	// No paths go from dac -> fft so don't waste time calculating it.
	// Also, calculating dac -> fft takes forever and that's a sign I probably need to optimize further.
	svrFftNumPaths := numPaths(forwardGraph, backwardGraph, "svr", "fft")
	fmt.Println("svrFftNumPaths", svrFftNumPaths)
	fftDacNumPaths := numPaths(forwardGraph, backwardGraph, "fft", "dac")
	fmt.Println("fftDacNumPaths", fftDacNumPaths)
	dacOutNumPaths := numPaths(forwardGraph, backwardGraph, "dac", "out")
	fmt.Println("dacOutNumPaths", dacOutNumPaths)

	fmt.Println("totalPaths", svrFftNumPaths * fftDacNumPaths * dacOutNumPaths)
}
