package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/snewell/advent2021/internal/aoc"
)

type cave struct {
	connections []*cave
	name        string
	big         bool
	id          int
}

func readMap(input io.Reader) map[string]*cave {
	nextId := 0
	caves := map[string]*cave{}

	getCave := func(name string) *cave {
		value, found := caves[name]
		if found {
			return value
		}
		value = &cave{
			name: name,
			big:  name[0] >= 'A' && name[0] <= 'Z',
			id:   nextId,
		}
		caves[name] = value
		nextId++
		return value
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		endpoints := strings.Split(scanner.Text(), "-")
		first := getCave(endpoints[0])
		second := getCave(endpoints[1])
		first.connections = append(first.connections, second)
		second.connections = append(second.connections, first)
	}

	return caves
}

func countSubPath(end *cave, caves *map[string]*cave, visited *[]bool, last *cave, bigLoop *map[string]struct{}) int {
	count := 0
	for _, candidate := range last.connections {
		if candidate == end {
			count++
		} else {
			if candidate.big {
				_, found := (*bigLoop)[candidate.name]
				if !found {
					(*bigLoop)[candidate.name] = struct{}{}
					count += countSubPath(end, caves, visited, candidate, bigLoop)
					delete(*bigLoop, candidate.name)
				}
			} else {
				if !(*visited)[candidate.id] {
					(*visited)[candidate.id] = true
					count += countSubPath(end, caves, visited, candidate, &map[string]struct{}{})
					(*visited)[candidate.id] = false
				}
			}
		}
	}
	return count
}

func countPaths(caves *map[string]*cave) int {
	visited := make([]bool, len(*caves))
	start := (*caves)["start"]
	end := (*caves)["end"]
	visited[start.id] = true
	return countSubPath(end, caves, &visited, start, &map[string]struct{}{})
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		caves := readMap(input)
		count := countPaths(&caves)
		return count
	})
}
