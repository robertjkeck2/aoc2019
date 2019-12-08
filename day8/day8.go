package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	imgHeight := 6
	imgWidth := 25
	filename := "image.txt"
	imageData := loadImageData(filename, imgHeight, imgWidth)
	minLayer := findFewestZeros(imageData)
	checksum := numOnesTimesTwos(imageData[minLayer])
	fmt.Println(checksum)
	decodedImage := decodeImage(imageData)
	fmt.Println(decodedImage)
	displayImage(decodedImage, imgHeight, imgWidth)
}

func loadImageData(filename string, imgHeight int, imgWidth int) (layers map[int][]int) {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var allRunes []rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allRunes = []rune(scanner.Text())
		if err != nil {
			break
		}
	}

	layerSize := imgHeight * imgWidth
	layers = map[int][]int{}
	for i := 0; i < len(allRunes); i = i + layerSize {
		var layer []int
		for j := 0; j < layerSize; j++ {
			layer = append(layer, int(allRunes[i+j])-'0')
		}
		layers[i/layerSize] = layer
	}

	return layers
}

func findFewestZeros(layers map[int][]int) int {
	minZeros := len(layers[0])
	minLayer := 0
	for index, layer := range layers {
		numZeros := 0
		for _, val := range layer {
			if val == 0 {
				numZeros++
			}
		}
		if numZeros < minZeros {
			minZeros = numZeros
			minLayer = index
		}
	}
	return minLayer
}

func numOnesTimesTwos(layer []int) int {
	numOnes, numTwos := 0, 0
	for _, val := range layer {
		switch val {
		case 1:
			numOnes++
		case 2:
			numTwos++
		}
	}
	return numOnes * numTwos
}

func decodeImage(layers map[int][]int) (decodedImage [150]int) {
	for i := 0; i < len(layers[0]); i++ {
		for j := 0; j < len(layers); j++ {
			if layers[j][i] != 2 && decodedImage[i] == 0 {
				decodedImage[i] = layers[j][i]
				break
			}
		}
	}
	return decodedImage
}

func displayImage(decodedImage [150]int, imgHeight int, imgWidth int) {
	for i, pixelVal := range decodedImage {
		if i%imgWidth == 0 {
			fmt.Println()
		}
		if pixelVal == 0 {
			fmt.Print(" ")
		} else {
			fmt.Print("█")
		}
	}
}
