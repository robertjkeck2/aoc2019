package main

import (
	"fmt"

	"github.com/robertjkeck2/aoc2019/intcode"
)

func main() {
	filename := "opcodes.csv"
	input := int64(2)
	output := intcode.Run(filename, []int64{input})
	fmt.Println(output)
}
