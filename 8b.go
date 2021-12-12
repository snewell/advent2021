package main

import (
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

type panelSet map[rune]struct{}

func makeSet(input string) panelSet {
	ret := panelSet{}
	for _, val := range input {
		ret[val] = struct{}{}
	}
	return ret
}

func getIntersections(first panelSet, second panelSet) []rune {
	intersection := []rune{}
	for fk := range first {
		_, found := second[fk]
		if found {
			intersection = append(intersection, fk)
		}
	}
	return intersection
}

func findIntersectionCount(key panelSet, candidates []panelSet, target int) int {
	for index, candidate := range candidates {
		intersection := getIntersections(key, candidate)
		if len(intersection) == target {
			return index
		}
	}
	return -1
}

func findSet(target panelSet, count int, candidates []panelSet) (panelSet, []panelSet) {
	index := findIntersectionCount(target, candidates, count)
	if index == -1 {
		panic("No candidate")
	}
	ret := candidates[index]
	candidates[index] = candidates[len(candidates)-1]
	candidates = candidates[:len(candidates)-1]

	return ret, candidates
}

func buildSets(inputs []string) [10]panelSet {
	// 2: 1
	// 3: 7
	// 4: 4
	// 5: 2, 3, 5
	// 6: 0, 6, 9
	// 7: 8
	candidates := [6][]panelSet{}

	for _, input := range inputs {
		index := len(input) - 2
		candidates[index] = append(candidates[index], makeSet(input))
	}

	ret := [10]panelSet{}
	// easy ones
	ret[1] = candidates[0][0]
	ret[7] = candidates[1][0]
	ret[4] = candidates[2][0]
	ret[8] = candidates[5][0]

	// 6 panel numbers
	ret[9], candidates[4] = findSet(ret[4], 4, candidates[4])
	ret[0], candidates[4] = findSet(ret[7], 3, candidates[4])
	ret[6] = candidates[4][0]

	// 3 panel numbers
	ret[3], candidates[3] = findSet(ret[1], 2, candidates[3])
	ret[5], candidates[3] = findSet(ret[4], 3, candidates[3])
	ret[2] = candidates[3][0]

	return ret
}

var (
	lookupMap = map[rune]int{
		'a': 1 << 0,
		'b': 1 << 1,
		'c': 1 << 2,
		'd': 1 << 3,
		'e': 1 << 4,
		'f': 1 << 5,
		'g': 1 << 6,
	}
)

func makeBitKeys(panelValues [10]panelSet) map[int]int {
	ret := map[int]int{}
	for index, numBits := range panelValues {
		var key int
		for lookup := range numBits {
			key |= lookupMap[lookup]
		}
		ret[key] = index
	}
	return ret
}

func makeDigit(bitKeys map[int]int, digitValue string) int {
	var ret int
	for _, bit := range digitValue {
		ret |= lookupMap[bit]
	}
	return bitKeys[ret]
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		total := 0
		figureItOut := func(data aoc.SevenPanelInfo) {
			panelSets := buildSets(data.Inputs)
			bitKeys := makeBitKeys(panelSets)

			result := makeDigit(bitKeys, data.Outputs[0]) * 1000
			result += makeDigit(bitKeys, data.Outputs[1]) * 100
			result += makeDigit(bitKeys, data.Outputs[2]) * 10
			result += makeDigit(bitKeys, data.Outputs[3])
			total += result
		}
		aoc.LoadSevenPanelInfo(input, figureItOut)
		return total
	})
}
