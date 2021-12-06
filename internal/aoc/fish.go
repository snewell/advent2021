package aoc

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type FishPeriods [9]int

func ReadFishPeriods(input io.Reader) FishPeriods {
	ret := FishPeriods{}
	var data string
	fmt.Fscan(input, &data)
	values := strings.Split(data, ",")

	for _, value := range values {
		realValue, _ := strconv.Atoi(value)
		ret[realValue]++
	}

	return ret
}

func PassDays(periods FishPeriods, count int) FishPeriods {
	for i := 0; i < count; i++ {
		birthers := periods[0]

		for index := range periods[:8] {
			periods[index] = periods[index+1]
		}
		periods[8] = birthers
		periods[6] += birthers
	}
	return periods
}
