package main

import (
	"bufio"
	"fmt"
	"math"
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
	x                int
	y                int
	asteroidsInSight []asteroid
}

func main() {
	system := buildSystem("../input.txt")
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

	for _, a := range asteroidsInSystem {
		for _, b := range asteroidsInSystem {
			angle, distance := getAngleAndDistance(a, b)
		}
	}
}

func getAngleAndDistance(a, b asteroid) (float64, float64) {
	angle := (math.Atan2(float64(b.y-a.y), float64(b.x-a.x)) * 180) / math.Pi
	distance := math.Sqrt(math.Pow(float64(b.y-a.y), 2) + math.Pow(float64(b.x-a.x), 2))

	n := angle + 90
	if n < 0 {
		n += 360
	}

	return n, distance
}
