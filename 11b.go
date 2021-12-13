package main

import (
	"bufio"
	"io"

	"github.com/snewell/advent2021/internal/aoc"
)

type chargeMap [][]int

func readCharges(input io.Reader) chargeMap {
	ret := chargeMap{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := make([]int, len(scanner.Text()))
		for index, val := range scanner.Text() {
			row[index] = int(val - '0')
		}
		ret = append(ret, row)
	}
	return ret
}

type point struct {
	row int
	col int
}

func updateNeighbors(row int, col int, next *chargeMap, flashCache *map[point]struct{}) {
	check := func(row int, col int) {
		if (*next)[row][col] > 9 {
			p := point{row: row, col: col}
			_, found := (*flashCache)[p]
			if !found {
				(*flashCache)[p] = struct{}{}
				updateNeighbors(row, col, next, flashCache)
			}
		}
	}
	// left side
	if row > 0 {
		if col > 0 {
			(*next)[row-1][col-1]++
			check(row-1, col-1)
		}
		(*next)[row-1][col]++
		check(row-1, col)
		if col < (len((*next)[row]) - 1) {
			(*next)[row-1][col+1]++
			check(row-1, col+1)
		}
	}

	// right side
	if row < (len(*next) - 1) {
		if col > 0 {
			(*next)[row+1][col-1]++
			check(row+1, col-1)
		}
		(*next)[row+1][col]++
		check(row+1, col)
		if col < (len((*next)[row]) - 1) {
			(*next)[row+1][col+1]++
			check(row+1, col+1)
		}
	}

	// top
	if col > 0 {
		(*next)[row][col-1]++
		check(row, col-1)
	}

	// bottom
	if col < (len((*next)[row]) - 1) {
		(*next)[row][col+1]++
		check(row, col+1)
	}
}

func cycle(current chargeMap) (int, chargeMap) {
	next := current
	// first cycle
	for row, _ := range next {
		for col, _ := range next[row] {
			next[row][col]++
		}
	}

	flashCache := map[point]struct{}{}
	for row, _ := range next {
		for col, _ := range next[row] {
			if next[row][col] > 9 {
				p := point{row: row, col: col}
				_, found := flashCache[p]
				if !found {
					flashCache[p] = struct{}{}
					updateNeighbors(row, col, &next, &flashCache)
				}
			}
		}
	}

	for p := range flashCache {
		next[p.row][p.col] = 0
	}
	return len(flashCache), next
}

func main() {
	aoc.Run(func(input io.Reader) interface{} {
		charges := readCharges(input)
		for i := 0; ; i++ {
			flashes, next := cycle(charges)
			if flashes == 100 {
				return i + 1
			}
			charges = next
		}
	})
}
