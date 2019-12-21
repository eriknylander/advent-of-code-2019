package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buildSystem(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("error: %s\n", err))
	}
	defer file.Close()

	system := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		system = append(system, strings.Split(scanner.Text(), ""))
	}

	return system
}

type asteroid struct {
	x          int
	y          int
	observable int
}

func main() {
	system := buildSystem("./example1.txt")
	// potentialAsteroids := []asteroid{}
	// for y := 0; y < len(system); y++ {
	// 	for x := 0; x < len(system[y]); x++ {
	// 		if system[y][x] == "#" {
	// 			potentialAsteroids = append(potentialAsteroids, asteroid{
	// 				x:          x,
	// 				y:          y,
	// 				observable: asteroidsInSight(x, y, system),
	// 			})
	// 		}
	// 	}
	// }

	a := asteroidsInSight(4, 2, system)
	fmt.Println(a)
}

func asteroidsInSight(x, y int, system [][]string) int {
	inSight := 0
	inSight += lookRight(x, y, system)
	inSight += lookLeft(x, y, system)
	inSight += lookUp(x, y, system)
	inSight += lookDown(x, y, system)
	return inSight
}

func lookRight(x, y int, system [][]string) int {
	for i := x - 1; i >= 0; i-- {
		if system[y][i] == "#" {
			return 1
		}
	}
	return 0
}

func lookLeft(x, y int, system [][]string) int {
	for i := x + 1; i < len(system[y]); i++ {
		if system[y][i] == "#" {
			return 1
		}
	}
	return 0
}

func lookUp(x, y int, system [][]string) int {
	for i := y - 1; i >= 0; i-- {
		if system[i][x] == "#" {
			return 1
		}
	}
	return 0
}

func lookDown(x, y int, system [][]string) int {
	for i := y + 1; i < len(system); i++ {
		if system[i][x] == "#" {
			return 1
		}
	}
	return 0
}
