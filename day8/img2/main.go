package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Errorf("error: %s\n", err))
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	//input := "0222112222120000"
	input := string(b)

	ia := strings.Split(input, "")

	width := 25
	height := 6

	layers := [][][]string{}
	i := 0
	for i < len(ia) {
		layer := [][]string{}
		for j := 0; j < height; j++ {
			current := i
			i = i + width
			layer = append(layer, ia[current:i])
		}
		layers = append(layers, layer)
	}

	finalImage := make([][]string, height)
	for i := range finalImage {
		finalImage[i] = make([]string, width)
		for j := range finalImage[i] {
			finalImage[i][j] = ","
		}
	}

	for _, l := range layers {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				val := l[i][j]
				fiVal := finalImage[i][j]
				if fiVal == "," {
					switch val {
					case "0":
						finalImage[i][j] = "#"
					case "1":
						finalImage[i][j] = "."
					}
				}
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf(finalImage[i][j])
		}
		fmt.Printf("\n")
	}
}

func countChars(i []string, c string) int {
	sum := 0
	for _, v := range i {
		if v == c {
			sum++
		}
	}

	return sum
}
