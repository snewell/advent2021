package aoc

import (
	"bufio"
	"io"
)

func convertRawHeightMap(line string) []int {
	ret := make([]int, len(line))
	for index := range line {
		ret[index] = int(line[index] - '0')
	}
	return ret
}

func LoadBasins(input io.Reader) [][]int {
	ret := [][]int{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ret = append(ret, convertRawHeightMap(scanner.Text()))
	}
	return ret
}

func isLowPoint(heightMap [][]int, row int, col int) bool {
	current := heightMap[row][col]
	// horizontal checks
	if row != 0 {
		if current >= heightMap[row-1][col] {
			return false
		}
	}
	if row != len(heightMap)-1 {
		if current >= heightMap[row+1][col] {
			return false
		}
	}
	// vertical checks
	if col != 0 {
		if current >= heightMap[row][col-1] {
			return false
		}
	}
	// fmt.Printf("col = %v len = %v\n", col, len(heightMap))
	if col != len(heightMap[row])-1 {
		if current >= heightMap[row][col+1] {
			return false
		}
	}

	return true
}

func FindLowBasinPoints(heightMap [][]int, handler func(int, int)) {
	for row := range heightMap {
		for col := range heightMap[row] {
			if isLowPoint(heightMap, row, col) {
				handler(row, col)
			}
		}
	}
}
