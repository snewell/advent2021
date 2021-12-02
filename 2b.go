package main

import (
	"fmt"
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		position := 0
		depth := 0
		aim := 0

		for {
			var direction string
			var value int
			_, err := fmt.Fscan(input, &direction, &value)
			if err != nil {
				if err == io.EOF {
					return position * depth
				}
				panic(err)
			}
			switch direction {
			case "forward":
				position += value
				depth += value * aim

			case "up":
				aim -= value

			case "down":
				aim += value

			default:
				panic(fmt.Sprintf("Unexpected command: %v", direction))
			}
		}
	})
}
