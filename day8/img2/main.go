package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
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

	input := string(b)

	ia := strings.Split(input,"")

	width := 2
	height := 2


	layers := [][][]string{}
	i := 0
	for i < len(ia) {
		layer := [][]string{}
		for j := 0; j < height; j++ {
			current := i
			i = i+width
			layer = append(layer, ia[current:i])
		}
		layers = append(layers, layer)
	}

	finalImage :=[][]string{}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			finalImage[i][j] = "2"
		}
	}

	for _, l := range layers {
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				val := l[i][j]
				fiVal := finalImage[i][j]
				if fiVal == "2" {
					fiVal = val
				}
			}
		}
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Printf(finalImage[i][j])
		}
		fmt.Printf("\n")
	}
}

func countChars(i []string, c string) int {
	sum := 0
	for _, v := range i {
		if v == c{
			sum++
		}
	}

	return sum
}

