package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/snewell/advent2021/internal/aoc"
)

type sevenPanelInfo struct {
	inputs  []string
	outputs []string
}

func readInput(input io.Reader) []sevenPanelInfo {
	scanner := bufio.NewScanner(input)

	ret := []sevenPanelInfo{}
	for scanner.Scan() {
		sides := strings.Split(scanner.Text(), "|")
		data := sevenPanelInfo{
			inputs:  strings.Split(sides[0], " "),
			outputs: strings.Split(sides[1], " "),
		}
		data.inputs = data.inputs[:len(data.inputs)-1]
		data.outputs = data.outputs[1:]
		ret = append(ret, data)
	}
	return ret
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		notes := readInput(input)
		count := 0
		for _, data := range notes {
			for _, inputData := range data.outputs {
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
		return count
	})
}
