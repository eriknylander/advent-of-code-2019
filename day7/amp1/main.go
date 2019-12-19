package main

import (
	"fmt"
	"sort"
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

var ip = 0

//var program = []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}
//var program = []int{3,23,3,24,1002,24,10,24,1002,23,-1,23, 101,5,23,23,1,24,23,23,4,23,99,0,0}
//var program = []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}
var program = []int{3,8,1001,8,10,8,105,1,0,0,21,42,55,64,77,94,175,256,337,418,99999,3,9,102,4,9,9,1001,9,5,9,102,2,9,9,101,3,9,9,4,9,99,3,9,102,2,9,9,101,5,9,9,4,9,99,3,9,1002,9,4,9,4,9,99,3,9,102,4,9,9,101,5,9,9,4,9,99,3,9,102,5,9,9,1001,9,3,9,1002,9,5,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99}
var mem = []int{}

var inputCh chan int
var outputCh chan int

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
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					for m := 0; m < 5; m++ {
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

func main() {
	inputCh = make(chan int, 10)
	outputCh = make(chan int, 10)

	allPhaseSettings := getAllPhaseSettings()
	thrustSignals := []int{}

	for _, ps := range allPhaseSettings {
		outputCh <- 0
		for i := 0; i < 5; i++ {
			ip = 0
			mem = copyProgram(program)
			inputCh <- ps[i]
			o := <- outputCh
			inputCh <- o
			runProgram()
		}
		thrustSignals = append(thrustSignals, <- outputCh)
	}

	sort.Ints(thrustSignals)
	fmt.Println(thrustSignals[len(thrustSignals)-1])
}

func runProgram() {
	for mem[ip] != cancel {
		//	fmt.Printf("ip: %d\n", ip)
		ins := newInstruction(mem[ip])
		//	fmt.Printf("instruction: %v\n", mem[ip:(ip+len)])
		process(ins)
		if ins.opCode != jumpIfTrue && ins.opCode != jumpIfFalse {
			ip += ins.len
		}
	}
}



func process(ins instruction) {
	switch ins.opCode {
	case add:
		val1 := getValue(ins.paramOneMode, mem[ip+1])
		val2 := getValue(ins.paramTwoMode, mem[ip+2])
		dst := getValue(positionMode, ip+3)
		sum := val1 + val2
		mem[dst] = sum
	case mul:
		val1 := getValue(ins.paramOneMode, mem[ip+1])
		val2 := getValue(ins.paramTwoMode, mem[ip+2])
		dst := getValue(positionMode, ip+3)
		prod := val1 * val2
		mem[dst] = prod
	case read:
		mem[mem[ip+1]] = <- inputCh
	case output:
		c :=  getValue(ins.paramOneMode, mem[ip+1])
		outputCh <-c
	case jumpIfTrue:
		if getValue(ins.paramOneMode, mem[ip+1]) != 0 {
			ip = getValue(ins.paramTwoMode, mem[ip+2])
		} else {
			ip += ins.len
		}
	case jumpIfFalse:
		if getValue(ins.paramOneMode, mem[ip+1]) == 0 {
			ip = getValue(ins.paramTwoMode, mem[ip+2])
		} else {
			ip += ins.len
		}
	case lessThan:
		val1 := getValue(ins.paramOneMode, mem[ip+1])
		val2 := getValue(ins.paramTwoMode, mem[ip+2])
		dst := getValue(positionMode, ip+3)
		if val1 < val2{
			mem[dst] = 1
		} else {
			mem[dst] = 0
		}
	case equals:
		val1 := getValue(ins.paramOneMode, mem[ip+1])
		val2 := getValue(ins.paramTwoMode, mem[ip+2])
		dst := getValue(positionMode, ip+3)
		if val1 == val2{
			mem[dst] = 1
		} else {
			mem[dst] = 0
		}
	}
}

func getValue(mode, param int) int {
	switch mode {
	case positionMode:
		return mem[param]
	case immediateMode:
		return param
	}

	return 0
}
