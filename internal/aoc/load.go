package aoc

import (
	"fmt"
	"io"
	"os"
)

func Run(prog func(io.Reader) interface{}) {
	if len(os.Args) == 0 {
		prog(os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		result := prog(f)
		fmt.Printf("%v\n", result)
	}
}
