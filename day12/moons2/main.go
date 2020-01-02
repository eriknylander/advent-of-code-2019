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

func (m moon) equals(b moon) bool {
	return m.x == b.x && m.y == b.y && m.z == b.z && m.xVel == b.xVel && m.yVel == b.yVel && m.zVel == b.zVel
}

func (m moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=< x=%d, y=%d, z=%d>", m.x, m.y, m.z, m.xVel, m.yVel, m.zVel)
}

func (m moon) calcEnergy() float64 {
	pot := math.Abs(float64(m.x)) + math.Abs(float64(m.y)) + math.Abs(float64(m.z))
	vel := math.Abs(float64(m.xVel)) + math.Abs(float64(m.yVel)) + math.Abs(float64(m.zVel))
	return pot * vel
}

func getGravity(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	}
	return 0
}

func main() {
	moons := []moon{
		{x: -1, y: 0, z: 2},
		{x: 2, y: -10, z: -7},
		{x: 4, y: -8, z: 8},
		{x: 3, y: 5, z: -1},
	}

	// moons := []moon{
	// 	{x: -8, y: -10, z: 0},
	// 	{x: 5, y: 5, z: 10},
	// 	{x: 2, y: -7, z: 3},
	// 	{x: 9, y: -8, z: -3},
	// }

	// moons := []moon{
	// 	{x: 14, y: 2, z: 8},
	// 	{x: 7, y: 4, z: 10},
	// 	{x: 1, y: 17, z: 16},
	// 	{x: -4, y: -1, z: 1},
	// }

	current := make([]moon, len(moons))
	_ = copy(current, moons)

	counter := 0
	steps := []int{}

	m1Found := false
	m2Found := false
	m3Found := false
	m4Found := false
	for {
		current = moveOneStep(current)
		counter++

		if moons[0].equals(current[0]) && !m1Found {
			m1Found = true
			steps = append(steps, counter)
			fmt.Println(steps)
		}

		if moons[1].equals(current[1]) && !m2Found {
			m2Found = true
			steps = append(steps, counter)
			fmt.Println(steps)
		}

		if moons[2].equals(current[2]) && !m3Found {
			m3Found = true
			steps = append(steps, counter)
			fmt.Println(steps)
		}

		if moons[3].equals(current[3]) && !m4Found {
			m4Found = true
			steps = append(steps, counter)
			fmt.Println(steps)
		}

		if m1Found && m2Found && m3Found && m4Found {
			break
		}
	}

	fmt.Println(steps)
}

func moveOneStep(state []moon) []moon {
	current := make([]moon, len(state))
	_ = copy(current, state)

	newStep := []moon{}
	for _, m := range current {
		nm := m
		for _, cm := range current {
			nm.xVel += getGravity(m.x, cm.x)
			nm.yVel += getGravity(m.y, cm.y)
			nm.zVel += getGravity(m.z, cm.z)
		}

		nm.x += nm.xVel
		nm.y += nm.yVel
		nm.z += nm.zVel
		newStep = append(newStep, nm)
	}
	return newStep
}
