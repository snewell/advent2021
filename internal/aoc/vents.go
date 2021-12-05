package aoc

import (
	"bufio"
	_ "fmt"
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
