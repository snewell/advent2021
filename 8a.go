package main

import (
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		count := 0
		countFunc := func(data aoc.SevenPanelInfo) {
			for _, inputData := range data.Outputs {
				switch len(inputData) {
				case 2: // 1
					count++

				case 4: // 4
					count++

				case 3: // 7
					count++

				case 7: // 8
					count++
				}
			}
		}
		aoc.LoadSevenPanelInfo(input, countFunc)
		return count
	})
}
