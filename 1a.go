package main

import (
	"fmt"
	"io"
	"math"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) {
		count := 0

		previous := math.MaxInt32
		for {
			var current int
			_, err := fmt.Fscan(input, &current)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			if current > previous {
				count++
			}
			previous = current
		}

		fmt.Printf("%v\n", count)
	})
}
