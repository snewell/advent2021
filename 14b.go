package main

import (
	"bufio"
	"io"
	"math"
	"regexp"

	"github.com/snewell/advent2021/internal/aoc"
)

type pair struct {
	first  rune
	second rune
}

type cacheEntry struct {
	pair
	round int
}

type lookupMap map[pair]rune
type cacheMap map[cacheEntry][26]uint64

func readData(input io.Reader) (string, lookupMap) {
	insertionPattern, _ := regexp.Compile(`^([A-Z]{2}) -> ([A-Z])$`)

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	first := scanner.Text()
	scanner.Scan() // burn the blank
	rules := lookupMap{}
	for scanner.Scan() {
		subs := insertionPattern.FindStringSubmatch(scanner.Text())
		p := pair{first: rune(subs[1][0]), second: rune(subs[1][1])}
		rules[p] = rune(subs[2][0])
	}

	return first, rules
}

func apply(first rune, second rune, rounds int, rules *lookupMap, cache *cacheMap) [26]uint64 {
	p := pair{first: first, second: second}
	ce := cacheEntry{pair: p, round: rounds}
	c, found := (*cache)[ce]
	if found {
		return c
	}
	ret := [26]uint64{}
	addition, found := (*rules)[p]
	if found {
		ret[addition-'A']++
		if rounds > 1 {
			a := apply(first, addition, rounds-1, rules, cache)
			b := apply(addition, second, rounds-1, rules, cache)

			for index := range a {
				ret[index] += a[index] + b[index]
			}
		}
	}
	(*cache)[ce] = ret
	return ret
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		first, rules := readData(input)
		cache := cacheMap{}
		counts := [26]uint64{}
		for i := 0; i < (len(first) - 1); i++ {
			current := first[i : i+2]
			pairCount := apply(rune(current[0]), rune(current[1]), 40, &rules, &cache)
			for index := range pairCount {
				counts[index] += pairCount[index]
			}
		}
		for _, val := range first {
			counts[val-'A']++
		}

		minCount := uint64(math.MaxUint64)
		maxCount := uint64(0)
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
