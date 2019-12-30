package main

import (
	"fmt"
	"math"
)

type moon struct {
	x    int
	y    int
	z    int
	xVel int
	yVel int
	zVel int
}

func (m moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=< x=%d, y=%d, z=%d>", m.x, m.y, m.z, m.xVel, m.yVel, m.zVel)
}

func (m moon) calcEnergy() float64 {
	pot := math.Abs(float64(m.x)) + math.Abs(float64(m.y)) + math.Abs(float64(m.z))
	vel := math.Abs(float64(m.xVel)) + math.Abs(float64(m.yVel)) + math.Abs(float64(m.zVel))
	return pot * vel
}
func main() {
	// moons := []moon{
	// 	{x: -1, y: 0, z: 2},
	// 	{x: 2, y: -10, z: -7},
	// 	{x: 4, y: -8, z: 8},
	// 	{x: 3, y: 5, z: -1},
	// }

	// moons := []moon{
	// 	{x: -8, y: -10, z: 0},
	// 	{x: 5, y: 5, z: 10},
	// 	{x: 2, y: -7, z: 3},
	// 	{x: 9, y: -8, z: -3},
	// }

	moons := []moon{
		{x: 14, y: 2, z: 8},
		{x: 7, y: 4, z: 10},
		{x: 1, y: 17, z: 16},
		{x: -4, y: -1, z: 1},
	}

	n := 1000
	current := moons
	for i := 0; i < n; i++ {
		newMoons := []moon{}
		for _, m := range current {
			nm := m
			for _, cm := range current {
				if m.x < cm.x {
					nm.xVel++
				} else if m.x > cm.x {
					nm.xVel--
				}

				if m.y < cm.y {
					nm.yVel++
				} else if m.y > cm.y {
					nm.yVel--
				}

				if m.z < cm.z {
					nm.zVel++
				} else if m.z > cm.z {
					nm.zVel--
				}
			}

			nm.x += nm.xVel
			nm.y += nm.yVel
			nm.z += nm.zVel
			newMoons = append(newMoons, nm)
		}
		current = newMoons
	}

	totalEnergy := 0.0
	for _, c := range current {
		fmt.Println(c)
		totalEnergy += c.calcEnergy()
	}
	fmt.Printf("Total energy: %v\n", totalEnergy)
}
