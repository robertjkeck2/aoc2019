package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pwRange struct {
	low, high int64
}

func main() {
	input := os.Args[1]
	pw := parseRange(input)
	numPossibilities := countPossibilities(pw)
	fmt.Println(numPossibilities)
}

func parseRange(in string) (pw pwRange) {
	splitInput := strings.Split(in, "-")
	low, _ := strconv.ParseInt(splitInput[0], 10, 64)
	high, _ := strconv.ParseInt(splitInput[1], 10, 64)
	pw = pwRange{low, high}
	return
}

func countPossibilities(pw pwRange) (count int) {
	count = 0
	for i := pw.low; i < pw.high; i++ {
		if checkCritera(int(i)) {
			count++
		}
	}
	return
}

func checkCritera(value int) (valid bool) {
	valid = true
	strValue := strconv.Itoa(value)
	counts := make(map[int]int)
	if value-100000 < 0 {
		valid = false
		return
	}
	if strValue[0] != strValue[1] && strValue[1] != strValue[2] && strValue[2] != strValue[3] && strValue[3] != strValue[4] && strValue[4] != strValue[5] {
		valid = false
		return
	}
	if strValue[0] > strValue[1] || strValue[1] > strValue[2] || strValue[2] > strValue[3] || strValue[3] > strValue[4] || strValue[4] > strValue[5] {
		valid = false
		return
	}
	for j := 48; j < 58; j++ {
		for i := 1; i < 6; i++ {
			if int(strValue[i]) == j && strValue[i-1] == strValue[i] {
				counts[j]++
			}
		}
	}
	if len(counts) == 1 {
		for _, val := range counts {
			if val > 1 {
				valid = false
				return
			}
		}
	}
	return
}
