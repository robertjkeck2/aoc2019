package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type param struct {
	key, value, mode int64
}

func main() {
	filename := "opcodes.csv"
	originalOpcodes := loadOpcodes(filename)
	runProgram(originalOpcodes)
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

func parseOpcode(opcodes []int64, i int) (op int64, params []param) {
	opcode := opcodes[i]
	if opcode == 99 {
		op = 99
		return op, params
	}
	opcodeStr := strconv.Itoa(int(opcode))
	op = int64(opcodeStr[len(opcodeStr)-1]) - 48
	rangeOffset, keyOffset, modeOffset := 2, 2, 0
	if len(opcodeStr) > 2 {
		if op == 1 || op == 2 || op == 5 || op == 6 || op == 7 || op == 8 {
			params = append(params, param{3, opcodes[i+3], 0})
			if len(opcodeStr) < 4 {
				params = append(params, param{2, opcodes[i+2], 0})
			}
		}
		for j := range opcodeStr[:len(opcodeStr)-rangeOffset] {
			key := int64(len(opcodeStr) - keyOffset - j)
			mode := int64(opcodeStr[j+modeOffset]) - 48
			params = append(params, param{key, opcodes[int64(i)+key], int64(mode)})
		}
	} else {
		if op < 3 || op > 4 {
			params = append(params, param{3, opcodes[i+3], 0})
			params = append(params, param{2, opcodes[i+2], 0})
			params = append(params, param{1, opcodes[i+1], 0})
		} else {
			params = append(params, param{1, opcodes[i+1], 0})
		}
	}
	return op, params
}

func evaluateOpcode(opcodes *[]int64, op int64, params []param, i int) (iNew int) {
	switch op {
	case 1:
		iNew = i + 4
		param1 := params[1]
		param2 := params[2]
		var param1Val, param2Val int64
		if param1.mode == 1 {
			param1Val = param1.value
		} else {
			param1Val = (*opcodes)[param1.value]
		}
		if param2.mode == 1 {
			param2Val = param2.value
		} else {
			param2Val = (*opcodes)[param2.value]
		}
		(*opcodes)[params[0].value] = param1Val + param2Val
	case 2:
		iNew = i + 4
		param1 := params[1]
		param2 := params[2]
		var param1Val, param2Val int64
		if param1.mode == 1 {
			param1Val = param1.value
		} else {
			param1Val = (*opcodes)[param1.value]
		}
		if param2.mode == 1 {
			param2Val = param2.value
		} else {
			param2Val = (*opcodes)[param2.value]
		}
		(*opcodes)[params[0].value] = param1Val * param2Val
	case 3:
		iNew = i + 2
		address := params[0].value
		reader := bufio.NewReader(os.Stdin)
		inputVal, _ := reader.ReadString('\n')
		inputValInt, _ := strconv.Atoi(inputVal[0:1])
		(*opcodes)[address] = int64(inputValInt)
	case 4:
		iNew = i + 2
		address := params[0].value
		if params[0].mode == 1 {
			fmt.Printf("CHECK - %d\n---\n", params[0].value)
		} else {
			fmt.Printf("CHECK - %d\n---\n", (*opcodes)[address])
		}
	case 5:
		param1 := params[1]
		param2 := params[2]
		var param1Val, param2Val int64
		if param1.mode == 1 {
			param1Val = param1.value
		} else {
			param1Val = (*opcodes)[param1.value]
		}
		if param2.mode == 1 {
			param2Val = param2.value
		} else {
			param2Val = (*opcodes)[param2.value]
		}
		if param1Val != 0 {
			iNew = int(param2Val)
		} else {
			iNew = i + 3
		}
	case 6:
		param1 := params[1]
		param2 := params[2]
		var param1Val, param2Val int64
		if param1.mode == 1 {
			param1Val = param1.value
		} else {
			param1Val = (*opcodes)[param1.value]
		}
		if param2.mode == 1 {
			param2Val = param2.value
		} else {
			param2Val = (*opcodes)[param2.value]
		}
		if param1Val == 0 {
			iNew = int(param2Val)
		} else {
			iNew = i + 3
		}
	case 7:
		iNew = i + 2
		param1 := params[1]
		param2 := params[2]
		var param1Val, param2Val int64
		if param1.mode == 1 {
			param1Val = param1.value
		} else {
			param1Val = (*opcodes)[param1.value]
		}
		if param2.mode == 1 {
			param2Val = param2.value
		} else {
			param2Val = (*opcodes)[param2.value]
		}
		if param1Val < param2Val {
			(*opcodes)[params[0].value] = 1
		} else {
			(*opcodes)[params[0].value] = 0
		}
	case 8:
		iNew = i + 2
		param1 := params[1]
		param2 := params[2]
		var param1Val, param2Val int64
		if param1.mode == 1 {
			param1Val = param1.value
		} else {
			param1Val = (*opcodes)[param1.value]
		}
		if param2.mode == 1 {
			param2Val = param2.value
		} else {
			param2Val = (*opcodes)[param2.value]
		}
		if param1Val > param2Val {
			(*opcodes)[params[0].value] = 1
		} else {
			(*opcodes)[params[0].value] = 0
		}
	case 99:
		iNew = -1
	}
	return iNew
}

func runProgram(originalOpcodes []int64) (output int64) {
	opcodes := make([]int64, len(originalOpcodes))
	copy(opcodes, originalOpcodes)
	i := 0
	for i >= 0 {
		opcode, params := parseOpcode(opcodes, i)
		iNew := evaluateOpcode(&opcodes, opcode, params, i)
		i = iNew
	}
	output = opcodes[0]
	return output
}
