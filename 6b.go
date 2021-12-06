package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/snewell/advent2021/internal/aoc"
)

func readInput(input io.Reader) map[int]int {
	ret := map[int]int{}
	var data string
	fmt.Fscan(input, &data)
	values := strings.Split(data, ",")

	for _, value := range values {
		realValue, _ := strconv.Atoi(value)
		ret[realValue]++
	}

	return ret
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		counts := readInput(input)
		for i := 0; i < 256; i++ {
			next := map[int]int{}
			birthers, found := counts[0]
			delete(counts, 0)

			for days, count := range counts {
				next[days-1] = count
			}
			if found {
				next[8] = birthers
				next[6] += birthers
			}

			counts = next
		}

		total := 0
		for _, count := range counts {
			total += count
		}
		return total
	})
}
