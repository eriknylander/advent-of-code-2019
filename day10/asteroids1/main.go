package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	x                int
	y                int
	asteroidsInSight int
}

func main() {
	system := buildSystem("./input.txt")
	asteroidsInSystem := []asteroid{}

	for y := 0; y < len(system); y++ {
		for x := 0; x < len(system[y]); x++ {
			if system[y][x] == "#" {
				asteroidsInSystem = append(asteroidsInSystem, asteroid{
					x: x,
					y: y,
				})
			}
		}
	}

	for i, asteroid := range asteroidsInSystem {
		diag := lookDiagonal(asteroid, asteroidsInSystem)
		hv := lookHorisontalAndVertical(asteroid, system)
		fmt.Printf("x: %d, y: %d, observable: diag: %d, hv: %d\n", asteroid.x, asteroid.y, diag, hv)
		asteroidsInSystem[i].asteroidsInSight = diag + hv
	}

	sort.Slice(asteroidsInSystem, func(i, j int) bool {
		return asteroidsInSystem[i].asteroidsInSight > asteroidsInSystem[j].asteroidsInSight
	})

	best := asteroidsInSystem[0]
	fmt.Printf("%d,%d, %d\n", best.x, best.y, best.asteroidsInSight)
}

func lookDiagonal(a asteroid, asteroids []asteroid) int {
	type relation struct {
		x bool
		y bool
		k float64
	}
	lineMap := map[relation]bool{}
	for _, o := range asteroids {
		if a.x == o.x && a.y == o.y {
			continue
		}

		dx := float64(a.x - o.x)
		if dx == 0 {
			continue
		}

		dy := float64(a.y - o.y)
		if dy == 0 {
			continue
		}

		k := dy / dx
		key := relation{k: k}
		if dx > 0 {
			key.x = true
		}

		if dy > 0 {
			key.y = true
		}

		if _, exists := lineMap[key]; !exists {
			lineMap[key] = true
		}
	}
	return len(lineMap)
}

func lookHorisontalAndVertical(a asteroid, system [][]string) int {
	inSight := 0
	inSight += lookRight(a.x, a.y, system)
	inSight += lookLeft(a.x, a.y, system)
	inSight += lookUp(a.x, a.y, system)
	inSight += lookDown(a.x, a.y, system)
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
