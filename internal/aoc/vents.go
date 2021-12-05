package aoc

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	First  Point
	Second Point
}

var (
	linePattern *regexp.Regexp
)

func min(first int, second int) int {
	if first < second {
		return first
	}
	return second
}

func max(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

func FindHorizontalOverlap(first Line, second Line) []Point {
	if first.First.Y == second.First.Y {
		// we have some overlap
		leftX := max(first.First.X, second.First.X)
		rightX := min(first.Second.X, second.Second.X)

		if leftX <= rightX {
			ret := make([]Point, rightX-leftX+1)
			for i := 0; i < (rightX - leftX + 1); i++ {
				ret[i] = Point{
					X: leftX + i,
					Y: first.First.Y,
				}
			}
			return ret
		}
	}
	return []Point{}
}

func FindVerticalOverlap(first Line, second Line) []Point {
	if first.First.X == second.First.X {
		// we have some overlap
		leftY := max(first.First.Y, second.First.Y)
		rightY := min(first.Second.Y, second.Second.Y)

		if leftY <= rightY {
			ret := make([]Point, rightY-leftY+1)
			for i := 0; i < (rightY - leftY + 1); i++ {
				ret[i] = Point{
					X: first.First.X,
					Y: leftY + i,
				}
			}
			return ret
		}
	}
	return []Point{}
}

func extractPoint(x string, y string) (Point, error) {
	realX, err := strconv.Atoi(x)
	if err != nil {
		return Point{}, err
	}
	realY, err := strconv.Atoi(y)
	if err != nil {
		return Point{}, err
	}
	return Point{
		X: realX,
		Y: realY,
	}, nil
}

func ReadLines(input io.Reader, pred func(Line) bool) ([]Line, error) {
	ret := []Line{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		subs := linePattern.FindStringSubmatch(scanner.Text())
		// fmt.Printf("subs = %v", subs)
		first, err := extractPoint(subs[1], subs[2])
		if err != nil {
			return nil, err
		}
		second, err := extractPoint(subs[3], subs[4])
		if err != nil {
			return nil, err
		}
		line := Line{
			First:  first,
			Second: second,
		}
		if pred(line) {
			ret = append(ret, line)
		}
	}
	return ret, nil
}

func init() {
	var err error
	linePattern, err = regexp.Compile(`^([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)$`)
	if err != nil {
		panic(err)
	}
}
