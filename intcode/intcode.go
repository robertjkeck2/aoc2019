package intcode

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Run takes Intcode instructions via a CSV file and output the result
// Credit to /u/idolstar/ on reddit for the initial intcode emulator
func Run(filename string, input []int64) (output []int64) {
	originalOpcodes := loadOpcodes(filename)
	runIntcodeInstructions(originalOpcodes, input, &output)
	return output
}

func loadOpcodes(filename string) (opcodes [1000000]int64) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	opcode, _ := reader.Read()

	for i, value := range opcode {
		opcodeInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			break
		}
		opcodes[i] = opcodeInt
	}
	return opcodes
}

func runIntcodeInstructions(opcodes [1000000]int64, input []int64, output *[]int64) error {
	relBase := int64(0)
	for ptr := int64(0); ptr < int64(len(opcodes)); {
		opcode := opcodes[ptr] % 100
		offset := int64(0)
		if opcodes[ptr]/10000 == 2 {
			offset = relBase
		}
		switch opcode {
		case 1:
			// ADD
			opcodes[opcodes[ptr+3]+offset] = evaluateParameters(opcodes, ptr, 1, relBase) + evaluateParameters(opcodes, ptr, 2, relBase)
			ptr += 4
		case 2:
			// MULTIPLY
			opcodes[opcodes[ptr+3]+offset] = evaluateParameters(opcodes, ptr, 1, relBase) * evaluateParameters(opcodes, ptr, 2, relBase)
			ptr += 4
		case 3:
			// INPUT
			if opcodes[ptr]/100 == 2 {
				opcodes[relBase+opcodes[ptr+1]] = input[0]
			} else {
				opcodes[opcodes[ptr+1]] = input[0]
			}
			input = input[1:]
			ptr += 2
		case 4:
			// OUTPUT
			*output = append(*output, evaluateParameters(opcodes, ptr, 1, relBase))
			ptr += 2
		case 5:
			// JUMP if TRUE
			if evaluateParameters(opcodes, ptr, 1, relBase) == 0 {
				ptr += 3
			} else {
				ptr = evaluateParameters(opcodes, ptr, 2, relBase)
			}
		case 6:
			// JUMP if FALSE
			if evaluateParameters(opcodes, ptr, 1, relBase) == 0 {
				ptr = evaluateParameters(opcodes, ptr, 2, relBase)
			} else {
				ptr += 3
			}
		case 7:
			// LESS THAN
			if evaluateParameters(opcodes, ptr, 1, relBase) < evaluateParameters(opcodes, ptr, 2, relBase) {
				opcodes[opcodes[ptr+3]+offset] = 1
			} else {
				opcodes[opcodes[ptr+3]+offset] = 0
			}
			ptr += 4
		case 8:
			// EQUALS
			if evaluateParameters(opcodes, ptr, 1, relBase) == evaluateParameters(opcodes, ptr, 2, relBase) {
				opcodes[opcodes[ptr+3]+offset] = 1
			} else {
				opcodes[opcodes[ptr+3]+offset] = 0
			}
			ptr += 4
		case 9:
			// RELATIVE BASE SHIFT
			relBase += evaluateParameters(opcodes, ptr, 1, relBase)
			ptr += 2
		case 99:
			// HALT
			return nil
		default:
			return fmt.Errorf("Unexpected opcode: %d", opcodes[ptr])
		}
	}
	return fmt.Errorf("Ran out of program without halt")
}

func evaluateParameters(opcodes [1000000]int64, ptr int64, parameter int64, relBase int64) int64 {
	j := int64(10)
	for i := int64(0); i < parameter; i++ {
		j *= 10
	}
	parameterMode := (opcodes[ptr] / j) % 10

	switch parameterMode {
	case 0:
		// Position mode (return value at the position of parameter)
		return opcodes[opcodes[ptr+parameter]]
	case 1:
		// Immediate mode (return value of parameter)
		return opcodes[ptr+parameter]
	case 2:
		// Relative mode (return value at position of relative base)
		return opcodes[relBase+opcodes[ptr+parameter]]
	default:
		panic(fmt.Errorf("Unexpected parameter mode %d for opcode %d at position %d", evaluateParameters, opcodes[ptr], ptr, relBase))
	}
}
