package aoc

import (
	"testing"
)

func comparePointOverlaps(t *testing.T, actual []Point, expected []Point) {
	if len(actual) != len(expected) {
		t.Fatalf("Unexpected number of points (%v)", actual)
	}
	for index := range actual {
		if actual[index] != expected[index] {
			t.Fatalf("Unexpected point: %v", actual[index])
		}
	}
}

func TestMismatchedY(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 20, Y: 10},
	}
	second := Line{
		First:  Point{X: 10, Y: 20},
		Second: Point{X: 20, Y: 20},
	}

	overlap := FindHorizontalOverlap(first, second)
	if len(overlap) != 0 {
		t.Fatalf("Unexpected overlap (%v)", overlap)
	}
}

func TestNoHorizontalOverlap(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 20, Y: 10},
	}
	second := Line{
		First:  Point{X: 30, Y: 10},
		Second: Point{X: 40, Y: 10},
	}

	overlap := FindHorizontalOverlap(first, second)
	if len(overlap) != 0 {
		t.Fatalf("Unexpected overlap (%v)", overlap)
	}
}

func TestRightOverlap(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 20, Y: 10},
	}
	second := Line{
		First:  Point{X: 15, Y: 10},
		Second: Point{X: 25, Y: 10},
	}

	overlap := FindHorizontalOverlap(first, second)
	expected := []Point{
		Point{X: 15, Y: 10},
		Point{X: 16, Y: 10},
		Point{X: 17, Y: 10},
		Point{X: 18, Y: 10},
		Point{X: 19, Y: 10},
		Point{X: 20, Y: 10},
	}
	comparePointOverlaps(t, overlap, expected)
}

func TestLeftOverlap(t *testing.T) {
	first := Line{
		First:  Point{X: 15, Y: 10},
		Second: Point{X: 25, Y: 10},
	}
	second := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 20, Y: 10},
	}

	overlap := FindHorizontalOverlap(first, second)
	expected := []Point{
		Point{X: 15, Y: 10},
		Point{X: 16, Y: 10},
		Point{X: 17, Y: 10},
		Point{X: 18, Y: 10},
		Point{X: 19, Y: 10},
		Point{X: 20, Y: 10},
	}
	comparePointOverlaps(t, overlap, expected)
}

func TestSingleHorizontalOverlap(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 20, Y: 10},
	}
	second := Line{
		First:  Point{X: 20, Y: 10},
		Second: Point{X: 25, Y: 10},
	}

	overlap := FindHorizontalOverlap(first, second)
	expected := []Point{
		Point{X: 20, Y: 10},
	}
	comparePointOverlaps(t, overlap, expected)
}

func TestMismatchedX(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 10, Y: 20},
	}
	second := Line{
		First:  Point{X: 20, Y: 10},
		Second: Point{X: 20, Y: 20},
	}

	overlap := FindVerticalOverlap(first, second)
	if len(overlap) != 0 {
		t.Fatalf("Unexpected overlap (%v)", overlap)
	}
}

func TestNoVerticalOverlap(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 10, Y: 20},
	}
	second := Line{
		First:  Point{X: 10, Y: 30},
		Second: Point{X: 10, Y: 40},
	}

	overlap := FindVerticalOverlap(first, second)
	if len(overlap) != 0 {
		t.Fatalf("Unexpected overlap (%v)", overlap)
	}
}

func TestTopOverlap(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 10, Y: 20},
	}
	second := Line{
		First:  Point{X: 10, Y: 15},
		Second: Point{X: 10, Y: 25},
	}

	overlap := FindVerticalOverlap(first, second)
	expected := []Point{
		Point{X: 10, Y: 15},
		Point{X: 10, Y: 16},
		Point{X: 10, Y: 17},
		Point{X: 10, Y: 18},
		Point{X: 10, Y: 19},
		Point{X: 10, Y: 20},
	}
	comparePointOverlaps(t, overlap, expected)
}

func TestBottomOverlap(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 15},
		Second: Point{X: 10, Y: 25},
	}
	second := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 10, Y: 20},
	}

	overlap := FindVerticalOverlap(first, second)
	expected := []Point{
		Point{X: 10, Y: 15},
		Point{X: 10, Y: 16},
		Point{X: 10, Y: 17},
		Point{X: 10, Y: 18},
		Point{X: 10, Y: 19},
		Point{X: 10, Y: 20},
	}
	comparePointOverlaps(t, overlap, expected)
}

func TestSingleVerticalOverlap(t *testing.T) {
	first := Line{
		First:  Point{X: 10, Y: 10},
		Second: Point{X: 10, Y: 20},
	}
	second := Line{
		First:  Point{X: 10, Y: 20},
		Second: Point{X: 10, Y: 30},
	}

	overlap := FindVerticalOverlap(first, second)
	expected := []Point{
		Point{X: 10, Y: 20},
	}
	comparePointOverlaps(t, overlap, expected)
}
