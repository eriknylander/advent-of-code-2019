package main

import (
	"fmt"
	"sort"
	"sync"
)

const (
	positionMode = 0
	immediateMode = 1

	add = 1
	mul = 2
	read = 3
	output = 4
	jumpIfTrue = 5
	jumpIfFalse = 6
	lessThan = 7
	equals = 8

	cancel = 99
)


//var program = []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}
//var program = []int{3,23,3,24,1002,24,10,24,1002,23,-1,23, 101,5,23,23,1,24,23,23,4,23,99,0,0}
//var program = []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}
var program = []int{3,8,1001,8,10,8,105,1,0,0,21,42,55,64,77,94,175,256,337,418,99999,3,9,102,4,9,9,1001,9,5,9,102,2,9,9,101,3,9,9,4,9,99,3,9,102,2,9,9,101,5,9,9,4,9,99,3,9,1002,9,4,9,4,9,99,3,9,102,4,9,9,101,5,9,9,4,9,99,3,9,102,5,9,9,1001,9,3,9,1002,9,5,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99}
//var program  = []int{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5}
//var program = []int{3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10}


type instruction struct {
	opCode       int
	paramOneMode int
	paramTwoMode int
	len int
}

func newInstruction(o int) instruction {
	var opCode int
	var paramOneMode int
	var paramTwoMode int

	if o > 999 {
		paramTwoMode = 1
	}

	if (o - (1000 * paramTwoMode)) > 99 {
		paramOneMode = 1
	}

	opCode = o % 100

	return instruction{
		opCode:       opCode,
		paramOneMode: paramOneMode,
		paramTwoMode: paramTwoMode,
		len: insLen(opCode),
	}
}

func insLen(opCode int) int {
	switch opCode {
	case 1, 2, 7, 8:
		return 4
	case 3,4:
		return 2
	case 5,6:
		return 3
	default:
		return 1
	}
}

func getAllPhaseSettings() [][]int {
	allPhaseSettings := [][]int{}
	for i := 5; i < 10; i++ {
		for j := 5; j < 10; j++ {
			for k := 5; k < 10; k++ {
				for l := 5; l < 10; l++ {
					for m := 5; m < 10; m++ {
						phaseSettings := []int{i,j,k,l,m}
						if allUnique(phaseSettings) {
							allPhaseSettings = append(allPhaseSettings, phaseSettings)
						}
					}
				}
			}
		}
	}

	return allPhaseSettings
}

func allUnique(phaseSettings []int) bool {
	hitMap := map[int]int{}
	for _, p := range phaseSettings {
		hitMap[p] ++
	}

	for _, c := range hitMap {
		if c > 1 {return false}
	}
	return true
}

func copyProgram(program []int) []int {
	m := []int{}
	for _, i := range program{
		m = append(m, i)
	}

	return m
}

type process struct {
	id int
	memory []int
	ip int
	input chan int
	output chan int
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


func main() {
	allPhaseSettings := getAllPhaseSettings()
	//allPhaseSettings := [][]int{{9,7,8,5,6}}
	thrustSignals := []int{}

	for _, ps := range allPhaseSettings {
		wg := sync.WaitGroup{}

		processes := []process{{
			id: 0,
			memory:   copyProgram(program),
			ip:       0,
			output:   make(chan int, 10),
			finished: func() {
				fmt.Printf("Process %d done\n", 0)
				wg.Done()
			},
		}}

		for i := 1; i < 5; i++ {
			processes = append(processes, process{
				id: i,
				memory:   copyProgram(program),
				ip:       0,
				input:    processes[i-1].output,
				output:   make(chan int, 10),
				finished: func() {
					fmt.Printf("Process %d done\n", i)
					wg.Done()
				},
			})
		}

		processes[0].input = processes[4].output


		for i := 0; i < 5; i++ {
			ic := i
			processes[i].input <- ps[i]
			wg.Add(1)
			go processes[ic].Run()
		}
		processes[0].input <- 0


		wg.Wait()
		thrustSignals = append(thrustSignals, <- processes[4].output)
	}

	sort.Ints(thrustSignals)
	fmt.Println(thrustSignals[len(thrustSignals)-1])
}


func (p *process) processInstruction(ins instruction) {
	mem := p.memory
	ip := p.ip
	switch ins.opCode {
	case add:
		val1 := p.getValue(ins.paramOneMode, mem[ip+1])
		val2 := p.getValue(ins.paramTwoMode, mem[ip+2])
		dst := p.getValue(positionMode, ip+3)
		sum := val1 + val2
		mem[dst] = sum
	case mul:
		val1 := p.getValue(ins.paramOneMode, mem[ip+1])
		val2 := p.getValue(ins.paramTwoMode, mem[ip+2])
		dst := p.getValue(positionMode, ip+3)
		prod := val1 * val2
		mem[dst] = prod
	case read:
		fmt.Printf("Process %d: start reading input\n", p.id)
		mem[mem[ip+1]] = <- p.input
		fmt.Printf("Process %d: finished reading input\n", p.id)
	case output:
		fmt.Printf("Process %d: start writing output\n", p.id)
		c :=  p.getValue(ins.paramOneMode, mem[ip+1])
		fmt.Printf("Process %d: finished reading input\n", p.id)
		p.output <-c
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
		dst := p.getValue(positionMode, ip+3)
		if val1 < val2{
			mem[dst] = 1
		} else {
			mem[dst] = 0
		}
	case equals:
		val1 := p.getValue(ins.paramOneMode, mem[ip+1])
		val2 := p.getValue(ins.paramTwoMode, mem[ip+2])
		dst := p.getValue(positionMode, ip+3)
		if val1 == val2{
			mem[dst] = 1
		} else {
			mem[dst] = 0
		}
	}
}

func (p *process) getValue(mode, param int) int {
	mem := p.memory
	switch mode {
	case positionMode:
		return mem[param]
	case immediateMode:
		return param
	}

	return 0
}
