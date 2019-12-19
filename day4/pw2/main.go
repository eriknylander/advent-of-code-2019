package main

import (
	"fmt"
	"strconv"
)

func main() {
	potential := 0
	for i:=178416; i <= 676461; i++{
		if meetsRequirement(i) {
			potential++
		}
	}

	fmt.Println(potential)
}

func meetsRequirement(i int) bool {
	if !decreases(i) {
		return false
	}

	if !hasTwoAdjacent(i) {
		return false
	}

	return true
}

func hasTwoAdjacent(num int) bool {
	numS := strconv.Itoa(num)
	numCountMap := make(map[uint8]int)
	for i := 0; i < len(numS); i++ {
		numCountMap[numS[i]]++
	}

	hasDouble := false
	for _, c := range numCountMap {
		if c == 2 {
			hasDouble = true
		}
	}

	return hasDouble
}

func decreases(num int) bool {
	numS := strconv.Itoa(num)
	for i := 1; i < len(numS); i++ {
		if numS[i] < numS[i-1] {
			return false
		}
	}

	return true
}

