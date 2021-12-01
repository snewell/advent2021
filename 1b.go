package main

import (
	"fmt"
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) {
		count := 0

		windows := []int{0, 0, 0}
		for seed := 1; seed < 4; seed++ {
			var current int
			fmt.Fscan(input, &current)
			for i := 0; i < seed; i++ {
				windows[i] += current
			}
		}

		previous := windows[0]
		for {
			for i := 0; i < 2; i++ {
				windows[i] = windows[i+1]
			}
			windows[2] = 0

			var current int
			_, err := fmt.Fscan(input, &current)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			for i := 0; i < 3; i++ {
				windows[i] += current
			}
			if windows[0] > previous {
				count++
			}
			previous = windows[0]
		}

		fmt.Printf("%v\n", count)
	})
}
