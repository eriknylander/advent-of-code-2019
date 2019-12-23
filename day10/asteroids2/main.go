package main

import (
	"bufio"
	"fmt"
	"math"
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
	x int
	y int
}

type relatedAsteroid struct {
	a        asteroid
	distance float64
	angle    float64
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

	type candidate struct {
		asteroid  asteroid
		theOthers []relatedAsteroid
		inSight   int
	}

	candiates := []candidate{}

	for _, a := range asteroidsInSystem {
		others := []relatedAsteroid{}
		otherMap := map[float64]bool{}
		for _, b := range asteroidsInSystem {
			if a.x == b.x && a.y == b.y {
				continue
			}
			angle, distance := getAngleAndDistance(a, b)
			others = append(others, relatedAsteroid{a: b, distance: distance, angle: angle})
			otherMap[angle] = true
		}
		sort.SliceStable(others, func(i, j int) bool {
			iAst := others[i]
			jAst := others[j]
			if iAst.angle == jAst.angle {
				return iAst.distance < jAst.distance
			}

			return iAst.angle < jAst.angle
		})

		candiates = append(candiates, candidate{asteroid: a, theOthers: others, inSight: len(otherMap)})
	}

	sort.SliceStable(candiates, func(i, j int) bool {
		return candiates[i].inSight > candiates[j].inSight
	})
	best := candiates[0]

	lastAngle := -1.0
	lap := 0
	laps := make([][]relatedAsteroid, len(best.theOthers))
	for _, a := range best.theOthers {
		if a.angle != lastAngle {
			if lap != 0 {
				lap = 0
			}
		} else {
			lap++
		}

		laps[lap] = append(laps[lap], a)
		lastAngle = a.angle
	}

	all := []relatedAsteroid{}
	for _, l := range laps {
		all = append(all, l...)
	}
	twoHundreth := all[199].a
	fmt.Printf("x: %d, y: %d, asteroids in sight: %d\n", best.asteroid.x, best.asteroid.y, best.inSight)
	fmt.Printf("200th vaporized: x:%d, y:%d, answer: %d\n", twoHundreth.x, twoHundreth.y, (100*twoHundreth.x)+twoHundreth.y)
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
