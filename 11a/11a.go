package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dfs(graph map[string][]string, deviceKey string, oldVisitedKeys map[string]bool, numPaths int) int {
	visitedKeys := map[string]bool{}

	for key, val := range oldVisitedKeys {
		visitedKeys[key] = val
	}

	visitedKeys[deviceKey] = true

	for _, outputKey := range graph[deviceKey] {
		if visitedKeys[outputKey] {
			continue
		}

		if outputKey == "out" {
			return numPaths + 1
		} else {
			numPaths = dfs(graph, outputKey, visitedKeys, numPaths)
		}
	}

	return numPaths
}

func main() {
	file, _ := os.Open("11a.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		key := strings.TrimSuffix(fields[0], ":")
		val := make([]string, 0)

		for _, field := range fields[1:] {
			val = append(val, field)
		}

		graph[key] = val
	}

	fmt.Println(dfs(graph, "you", map[string]bool{}, 0))
}
