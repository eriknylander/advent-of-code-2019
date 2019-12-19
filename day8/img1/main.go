package main

import (
	"fmt"
	"sort"
	"strings"
	"io/ioutil"
	"os"
)

//var input = "123456789012"
//var input = "123400123456"

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

	width := 25
	height := 6


	layers := [][]string{}
	i := 0
	for i < len(ia) {
		current := i
		i = (i+width*height)
		//fmt.Println(ia[current:i])
		layers = append(layers, ia[current:i])
	}

	sort.SliceStable(layers, func(i, j int) bool {
		return countChars(layers[i], "0") < countChars(layers[j], "0")
	})

	fmt.Println(countChars(layers[0], "1") * countChars(layers[0],"2"))
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
