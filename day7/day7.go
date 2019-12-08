package main

import (
	"fmt"

	"github.com/robertjkeck2/aoc2019/intcode"
)

type ampInput struct {
	phase int64
	input int64
}

func main() {
	filename := "testOpcodes.csv"
	output := intcode.Run(filename)
	fmt.Println(output)
}
