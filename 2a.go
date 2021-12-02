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

			case "up":
				depth -= value

			case "down":
				depth += value

			default:
				panic(fmt.Sprintf("Unexpected command: %v", direction))
			}
		}
	})
}
