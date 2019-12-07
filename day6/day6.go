package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type orbitalObject struct {
	id    string
	count int
}

type orbits map[string]orbitalObject

func main() {
	filename := "orbits.txt"
	allOrbits := orbits{}
	loadOrbits(filename, &allOrbits)
	totalOrbits := countAllOrbits(allOrbits)
	fmt.Println(totalOrbits)
	path1, len1 := traceOrbitalPath(allOrbits, "YOU")
	path2, len2 := traceOrbitalPath(allOrbits, "SAN")
	numTransfers := countOrbitalTransfers(path1, path2, len1, len2)
	fmt.Println(numTransfers)
}

func loadOrbits(filename string, allOrbits *orbits) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		orbit := scanner.Text()
		buildOrbitalPaths(orbit, allOrbits)
		if err != nil {
			break
		}
	}
	return
}

func buildOrbitalPaths(orbit string, allOrbits *orbits) {
	objects := strings.Split(orbit, ")")
	(*allOrbits)[objects[1]] = orbitalObject{objects[0], 0}
}

func incrementCount(allOrbits *orbits, objectID string) (count int) {
	object, ok := (*allOrbits)[objectID]
	if !ok {
		return 0
	}
	if object.count == 0 {
		object.count = 1 + incrementCount(allOrbits, object.id)
		(*allOrbits)[objectID] = object
	}
	return object.count
}

func countAllOrbits(allOrbits orbits) (totalPaths int) {
	for objectID := range allOrbits {
		totalPaths += incrementCount(&allOrbits, objectID)
	}
	return totalPaths
}

func traceOrbitalPath(allOrbits orbits, objectID string) (path []orbitalObject, length int) {
	object, ok := allOrbits[objectID]
	for ok {
		path = append(path, object)
		object, ok = allOrbits[object.id]
	}
	length = len(path)
	return path, length
}

func countOrbitalTransfers(path1, path2 []orbitalObject, len1, len2 int) (transfers int) {
	for len1 > 0 && len2 > 0 {
		len1--
		len2--
		if path1[len1].id != path2[len2].id {
			transfers = len1 + len2 + 2
			return transfers
		}
	}
	return
}
