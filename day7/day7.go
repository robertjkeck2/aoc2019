package main

import (
	"fmt"

	"github.com/robertjkeck2/aoc2019/intcode"
)

func main() {
	filename := "opcodes.csv"
	allowedPhases := []int64{5, 6, 7, 8, 9}
	maxSignal := findMaxSignal(filename, allowedPhases)
	fmt.Println(maxSignal)
}

func runAmplifiers(filename string, phases []int64) (output int64) {
	input := int64(0)
	for i := int64(0); i < 5; i++ {
		output := intcode.Run(filename, []int64{phases[i], input})
		input = output[0]
	}
	return input
}

func findMaxSignal(filename string, allowedPhased []int64) (maxSignal int64) {
	permuts := permutation(allowedPhased)
	for _, permutation := range permuts {
		output := runAmplifiers(filename, permutation)
		if output > maxSignal {
			maxSignal = output
		}
	}
	return maxSignal
}

func permutation(xs []int64) (permuts [][]int64) {
	var rc func([]int64, int64)
	rc = func(a []int64, k int64) {
		if k == int64(len(a)) {
			permuts = append(permuts, append([]int64{}, a...))
		} else {
			for i := k; i < int64(len(xs)); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)
	return permuts
}
