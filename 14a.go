package main

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strings"

	"github.com/snewell/advent2021/internal/aoc"
)

func readData(input io.Reader) (string, map[string]rune) {
	insertionPattern, _ := regexp.Compile(`^([A-Z]{2}) -> ([A-Z])$`)

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	first := scanner.Text()
	scanner.Scan() // burn the blank
	rules := map[string]rune{}
	for scanner.Scan() {
		subs := insertionPattern.FindStringSubmatch(scanner.Text())
		rules[subs[1]] = rune(subs[2][0])
	}

	return first, rules
}

func apply(start string, rules *map[string]rune) string {
	var sb strings.Builder

	sb.WriteRune(rune(start[0]))
	for i := 0; i < (len(start) - 1); i++ {
		addition, found := (*rules)[start[i:i+2]]
		if found {
			sb.WriteRune(addition)
		}
		sb.WriteRune(rune(start[i+1]))
	}
	return sb.String()
}

func countRunes(polymer string) [26]int {
	ret := [26]int{}
	for _, val := range polymer {
		ret[val-'A']++
	}
	return ret
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		first, rules := readData(input)
		for i := 0; i < 10; i++ {
			next := apply(first, &rules)
			first = next
		}
		counts := countRunes(first)

		minCount := math.MaxInt32
		maxCount := math.MinInt32
		for _, val := range counts {
			if val != 0 {
				if val < minCount {
					minCount = val
				}
				if val > maxCount {
					maxCount = val
				}
			}
		}
		return maxCount - minCount
	})
}
