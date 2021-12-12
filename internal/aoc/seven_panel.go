package aoc

import (
	"bufio"
	"io"
	"strings"
)

type SevenPanelInfo struct {
	Inputs  []string
	Outputs []string
}

func LoadSevenPanelInfo(input io.Reader, processFunc func(SevenPanelInfo)) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		sides := strings.Split(scanner.Text(), "|")
		data := SevenPanelInfo{
			Inputs:  strings.Split(sides[0], " "),
			Outputs: strings.Split(sides[1], " "),
		}
		data.Inputs = data.Inputs[:len(data.Inputs)-1]
		data.Outputs = data.Outputs[1:]
		processFunc(data)
	}
}
