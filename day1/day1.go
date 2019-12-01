package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var totalFuel int64
	filename := "mass.txt"
	masses := loadMasses(filename)
	for _, mass := range masses {
		totalFuel += calculateFuel(mass)
	}
	fmt.Println(totalFuel)
}

func loadMasses(filename string) []int64 {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			break
		}
		lines = append(lines, mass)
	}
	return lines
}

func calculateFuel(mass int64) int64 {
	fuel := int64(mass/3) - 2
	return fuel
}
