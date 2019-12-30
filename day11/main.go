package main

import (
	"fmt"
	"strings"
)

const (
	positionMode  = 0
	immediateMode = 1
	relativeMode  = 2

	add         = 1
	mul         = 2
	read        = 3
	output      = 4
	jumpIfTrue  = 5
	jumpIfFalse = 6
	lessThan    = 7
	equals      = 8
	adjustRBS   = 9

	cancel = 99
)

type instruction struct {
	opCode         int
	paramOneMode   int
	paramTwoMode   int
	paramThreeMode int
	len            int
}

func newInstruction(o int) instruction {
	var opCode int
	var paramOneMode int
	var paramTwoMode int
	var paramThreeMode int

	paramThreeMode = o / 10000

	paramTwoMode = (o - paramThreeMode*10000) / 1000

	paramOneMode = (o - (paramThreeMode * 10000) - (paramTwoMode * 1000)) / 100

	opCode = o % 100

	return instruction{
		opCode:         opCode,
		paramOneMode:   paramOneMode,
		paramTwoMode:   paramTwoMode,
		paramThreeMode: paramThreeMode,
		len:            insLen(opCode),
	}
}

func insLen(opCode int) int {
	switch opCode {
	case 1, 2, 7, 8:
		return 4
	case 3, 4, 9:
		return 2
	case 5, 6:
		return 3
	default:
		return 1
	}
}

func copyProgram(program []int) []int {
	m := []int{}
	for _, i := range program {
		m = append(m, i)
	}

	return m
}

type process struct {
	id       int
	memory   []int
	ip       int
	rbs      int
	input    chan int
	output   chan int
	finished func()
}

func (p *process) Run() {
	fmt.Printf("Running Process %d\n", p.id)
	for p.memory[p.ip] != cancel {
		//	fmt.Printf("ip: %d\n", ip)
		ins := newInstruction(p.memory[p.ip])
		//	fmt.Printf("instruction: %v\n", mem[ip:(ip+len)])
		p.processInstruction(ins)
		if ins.opCode != jumpIfTrue && ins.opCode != jumpIfFalse {
			p.ip += ins.len
		}
	}

	p.finished()
}

func (p *process) processInstruction(ins instruction) {
	mem := p.memory
	ip := p.ip
	switch ins.opCode {
	case add:
		val1 := p.getValue(ins.paramOneMode, mem[ip+1])
		val2 := p.getValue(ins.paramTwoMode, mem[ip+2])
		dst := p.getWriteAddress(ins.paramThreeMode, mem[ip+3])
		sum := val1 + val2
		mem[dst] = sum
	case mul:
		val1 := p.getValue(ins.paramOneMode, mem[ip+1])
		val2 := p.getValue(ins.paramTwoMode, mem[ip+2])
		dst := p.getWriteAddress(ins.paramThreeMode, mem[ip+3])
		prod := val1 * val2
		mem[dst] = prod
	case read:
		//fmt.Printf("Process %d: start reading input\n", p.id)
		dst := p.getWriteAddress(ins.paramOneMode, mem[ip+1])
		mem[dst] = <-p.input
		//fmt.Printf("Process %d: finished reading input\n", p.id)
	case output:
		//fmt.Printf("Process %d: start writing output\n", p.id)
		c := p.getValue(ins.paramOneMode, mem[ip+1])
		//fmt.Printf("Process %d: finished reading input\n", p.id)
		p.output <- c
	case jumpIfTrue:
		if p.getValue(ins.paramOneMode, mem[ip+1]) != 0 {
			p.ip = p.getValue(ins.paramTwoMode, mem[ip+2])
		} else {
			p.ip += ins.len
		}
	case jumpIfFalse:
		if p.getValue(ins.paramOneMode, mem[ip+1]) == 0 {
			p.ip = p.getValue(ins.paramTwoMode, mem[ip+2])
		} else {
			p.ip += ins.len
		}
	case lessThan:
		val1 := p.getValue(ins.paramOneMode, mem[ip+1])
		val2 := p.getValue(ins.paramTwoMode, mem[ip+2])
		dst := p.getWriteAddress(ins.paramThreeMode, mem[ip+3])
		if val1 < val2 {
			mem[dst] = 1
		} else {
			mem[dst] = 0
		}
	case equals:
		val1 := p.getValue(ins.paramOneMode, mem[ip+1])
		val2 := p.getValue(ins.paramTwoMode, mem[ip+2])
		dst := p.getWriteAddress(ins.paramThreeMode, mem[ip+3])
		if val1 == val2 {
			mem[dst] = 1
		} else {
			mem[dst] = 0
		}
	case adjustRBS:
		val1 := p.getValue(ins.paramOneMode, mem[ip+1])
		p.rbs += val1
	}
}

func (p *process) getValue(mode, param int) int {
	mem := p.memory
	switch mode {
	case positionMode:
		return mem[param]
	case immediateMode:
		return param
	case relativeMode:
		return mem[p.rbs+param]
	}

	return 0
}

func (p *process) getWriteAddress(mode, param int) int {
	switch mode {
	case positionMode:
		return param
	case relativeMode:
		return p.rbs + param
	default:
		panic("Invalid mode")
	}
}

type Position struct {
	x int
	y int
}

var program = []int{3, 8, 1005, 8, 310, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1002, 8, 1, 28, 1, 105, 11, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 55, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 76, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 98, 1, 1004, 7, 10, 1006, 0, 60, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 127, 2, 1102, 4, 10, 1, 1108, 7, 10, 2, 1102, 4, 10, 2, 101, 18, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 166, 1006, 0, 28, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 190, 1006, 0, 91, 1, 1108, 5, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1002, 8, 1, 220, 1, 1009, 14, 10, 2, 1103, 19, 10, 2, 1102, 9, 10, 2, 1007, 4, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 258, 2, 3, 0, 10, 1006, 0, 4, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1001, 8, 0, 286, 1006, 0, 82, 101, 1, 9, 9, 1007, 9, 1057, 10, 1005, 10, 15, 99, 109, 632, 104, 0, 104, 1, 21102, 1, 838479487636, 1, 21102, 327, 1, 0, 1106, 0, 431, 21102, 1, 932813579156, 1, 21102, 1, 338, 0, 1106, 0, 431, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21101, 0, 179318033447, 1, 21101, 385, 0, 0, 1105, 1, 431, 21101, 248037678275, 0, 1, 21101, 0, 396, 0, 1105, 1, 431, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 0, 709496558348, 1, 21102, 419, 1, 0, 1105, 1, 431, 21101, 825544561408, 0, 1, 21101, 0, 430, 0, 1106, 0, 431, 99, 109, 2, 22101, 0, -1, 1, 21101, 40, 0, 2, 21102, 462, 1, 3, 21101, 0, 452, 0, 1106, 0, 495, 109, -2, 2105, 1, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 457, 458, 473, 4, 0, 1001, 457, 1, 457, 108, 4, 457, 10, 1006, 10, 489, 1101, 0, 0, 457, 109, -2, 2106, 0, 0, 0, 109, 4, 2101, 0, -1, 494, 1207, -3, 0, 10, 1006, 10, 512, 21101, 0, 0, -3, 22101, 0, -3, 1, 22101, 0, -2, 2, 21101, 1, 0, 3, 21102, 531, 1, 0, 1105, 1, 536, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 559, 2207, -4, -2, 10, 1006, 10, 559, 22101, 0, -4, -4, 1106, 0, 627, 21202, -4, 1, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21102, 578, 1, 0, 1105, 1, 536, 22101, 0, 1, -4, 21101, 1, 0, -1, 2207, -4, -2, 10, 1006, 10, 597, 21102, 0, 1, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 619, 21201, -1, 0, 1, 21102, 1, 619, 0, 105, 1, 494, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2106, 0, 0}

func main() {
	inputCh := make(chan int, 1)
	outputCh := make(chan int, 2)

	done := make(chan int, 1)

	fullMem := program
	for i := 1; i < 3000; i++ {
		fullMem = append(fullMem, 0)
	}
	brainz := process{
		input:  inputCh,
		output: outputCh,
		memory: copyProgram(fullMem),
		finished: func() {
			done <- 1
		},
	}

	go brainz.Run()

	surface := map[Position]int{Position{x: 0, y: 0}: 1}

	go func() {
		x, y, direction := 0, 0, 0
		for {
			currentPos := Position{x: x, y: y}
			current := 0
			if val, ok := surface[currentPos]; ok {
				current = val
			}
			inputCh <- current

			color := <-outputCh
			turn := <-outputCh

			surface[currentPos] = color

			direction = rotate(turn, direction)
			xDiff, yDiff := getSteps(direction)
			x += xDiff
			y += yDiff
		}
	}()

	<-done

	xMin := 1000
	xMax := -1000
	yMin := 1000
	yMax := -1000

	for pos := range surface {
		if pos.x < xMin {
			xMin = pos.x
		}

		if pos.x > xMax {
			xMax = pos.x
		}

		if pos.y < yMin {
			yMin = pos.y
		}

		if pos.y > yMax {
			yMax = pos.y
		}
	}

	fmt.Printf("xMin: %d\n", xMin)
	fmt.Printf("xMax: %d\n", xMax)
	fmt.Printf("yMin: %d\n", yMin)
	fmt.Printf("yMax: %d\n", yMax)

	panels := make([][]string, yMax+1)
	for i := range panels {
		panels[i] = make([]string, xMax+1)
	}

	for pos, val := range surface {
		ch := " "
		if val == 1 {
			ch = "#"
		}
		panels[pos.y][pos.x] = ch
	}

	for j := 0; j <= yMax; j++ {
		fmt.Println(strings.Join(panels[j], " "))
	}

}

func rotate(turn, direction int) int {
	if turn == 0 { // turn lef
		return (direction + 270) % 360
	} else {
		return (direction + 90) % 360
	}
}

func getSteps(direction int) (int, int) {
	if direction == 0 {
		return 0, -1
	} else if direction == 90 {
		return 1, 0
	} else if direction == 180 {
		return 0, 1
	} else {
		return -1, 0
	}
}
