package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

type pathPoint struct {
	X, Y, steps int
}

func main() {
	filename := "wires.csv"
	wire1, wire2 := loadWirePaths(filename)
	points1 := drawPath(wire1)
	points2 := drawPath(wire2)
	intersections := findIntersections(points1, points2)
	fewestSteps := findMinStepIntersection(intersections)
	fmt.Println(fewestSteps)
	fmt.Println(calculateMinSteps(fewestSteps))
}

func loadWirePaths(filename string) (wire1, wire2 []string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	wire1, _ = reader.Read()
	wire2, _ = reader.Read()

	return wire1, wire2
}

func drawPath(wire []string) (points []pathPoint) {
	x := 0
	y := 0
	steps := 0
	for _, instr := range wire {
		direction := string([]rune(instr)[0])
		count, _ := strconv.ParseInt(instr[1:], 10, 64)
		switch direction {
		case "U":
			for i := 0; i < int(count); i++ {
				points = append(points, pathPoint{x, y, steps})
				y++
				steps++
			}
		case "D":
			for i := 0; i < int(count); i++ {
				points = append(points, pathPoint{x, y, steps})
				y--
				steps++
			}
		case "L":
			for i := 0; i < int(count); i++ {
				points = append(points, pathPoint{x, y, steps})
				x--
				steps++
			}
		case "R":
			for i := 0; i < int(count); i++ {
				points = append(points, pathPoint{x, y, steps})
				x++
				steps++
			}
		}
	}
	return points
}

func findIntersections(points1, points2 []pathPoint) (intersections [][]pathPoint) {
	for _, point1 := range points1 {
		for _, point2 := range points2 {
			match := []pathPoint{}
			if point1.X == point2.X && point1.Y == point2.Y {
				match = append(match, point1)
				match = append(match, point2)
				intersections = append(intersections, match)
			}
		}
	}
	return intersections
}

func findClosestIntersection(intersections [][]pathPoint) (closestIntersection pathPoint) {
	minPath := pathPoint{}
	minDist := 10000
	for i, path := range intersections[0] {
		if i != 0 {
			manhattanDist := calculateManhattanDist(path)
			if manhattanDist < minDist {
				minPath = path
				minDist = manhattanDist
			}
		}
	}
	return minPath
}

func findMinStepIntersection(intersections [][]pathPoint) (fewestSteps []pathPoint) {
	minPath := []pathPoint{}
	minSteps := 100000
	for i, path := range intersections {
		if i != 0 {
			steps := calculateMinSteps(path)
			if steps < minSteps {
				minPath = path
				minSteps = steps
			}
		}
	}
	return minPath
}

func calculateManhattanDist(path pathPoint) (dist int) {
	dist = int(math.Abs(float64(path.X)) + math.Abs(float64(path.Y)))
	return dist
}

func calculateMinSteps(path []pathPoint) (steps int) {
	steps = path[0].steps + path[1].steps
	return steps
}
