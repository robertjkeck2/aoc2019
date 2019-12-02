package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "opcodes.csv"
	opcodes := loadOpcodes(filename)
	for i := 0; i < len(opcodes); i = i + 4 {
		value := evaluateOpcode(opcodes[i], opcodes[i+1], opcodes[i+2], opcodes)
		if value == -1 {
			break
		} else {
			opcodes[opcodes[i+3]] = value
		}
	}
	fmt.Println(opcodes[0])
}

func loadOpcodes(filename string) (opcodes []int64) {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	opcode, _ := reader.Read()

	for _, value := range opcode {
		opcodeInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			break
		}
		opcodes = append(opcodes, opcodeInt)
	}
	return opcodes
}

func evaluateOpcode(op, pos1, pos2 int64, opcodes []int64) (value int64) {
	switch op {
	case 1:
		value = opcodes[pos1] + opcodes[pos2]
	case 2:
		value = opcodes[pos1] * opcodes[pos2]
	case 99:
		value = -1
	}
	return value
}
