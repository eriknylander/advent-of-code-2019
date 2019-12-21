package main

import (
	"fmt"
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
	id          int
	memory      []int
	ip          int
	rbs         int
	input       chan int
	readInput   func()
	output      chan int
	writeOutput func()
	finished    func()
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
		p.readInput()
		dst := p.getWriteAddress(ins.paramOneMode, mem[ip+1])
		mem[dst] = <-p.input
		//fmt.Printf("Process %d: finished reading input\n", p.id)
	case output:
		//fmt.Printf("Process %d: start writing output\n", p.id)
		c := p.getValue(ins.paramOneMode, mem[ip+1])
		//fmt.Printf("Process %d: finished reading input\n", p.id)
		p.output <- c
		p.writeOutput()
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

//var program = []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

//var program = []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}

//var program = []int{104, 1125899906842624, 99}

var program = []int{1102, 34463338, 34463338, 63, 1007, 63, 34463338, 63, 1005, 63, 53, 1102, 3, 1, 1000, 109, 988, 209, 12, 9, 1000, 209, 6, 209, 3, 203, 0, 1008, 1000, 1, 63, 1005, 63, 65, 1008, 1000, 2, 63, 1005, 63, 904, 1008, 1000, 0, 63, 1005, 63, 58, 4, 25, 104, 0, 99, 4, 0, 104, 0, 99, 4, 17, 104, 0, 99, 0, 0, 1102, 1, 21, 1008, 1101, 427, 0, 1028, 1102, 23, 1, 1012, 1101, 32, 0, 1009, 1101, 37, 0, 1007, 1102, 1, 892, 1023, 1102, 27, 1, 1004, 1102, 1, 38, 1013, 1102, 1, 20, 1005, 1101, 0, 29, 1001, 1101, 0, 22, 1015, 1102, 1, 35, 1003, 1101, 0, 39, 1016, 1102, 34, 1, 1011, 1101, 899, 0, 1022, 1102, 195, 1, 1024, 1101, 36, 0, 1014, 1101, 0, 24, 1000, 1102, 1, 31, 1006, 1101, 0, 28, 1017, 1101, 422, 0, 1029, 1102, 1, 33, 1019, 1102, 1, 26, 1018, 1102, 1, 0, 1020, 1102, 25, 1, 1002, 1102, 712, 1, 1027, 1101, 0, 190, 1025, 1101, 0, 715, 1026, 1102, 1, 1, 1021, 1101, 30, 0, 1010, 109, 30, 2105, 1, -6, 4, 187, 1106, 0, 199, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -19, 1206, 10, 211, 1106, 0, 217, 4, 205, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -13, 1202, 8, 1, 63, 1008, 63, 28, 63, 1005, 63, 241, 1001, 64, 1, 64, 1106, 0, 243, 4, 223, 1002, 64, 2, 64, 109, 8, 1201, -2, 0, 63, 1008, 63, 29, 63, 1005, 63, 263, 1105, 1, 269, 4, 249, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -9, 2101, 0, 3, 63, 1008, 63, 24, 63, 1005, 63, 295, 4, 275, 1001, 64, 1, 64, 1106, 0, 295, 1002, 64, 2, 64, 109, 12, 2107, 31, 0, 63, 1005, 63, 317, 4, 301, 1001, 64, 1, 64, 1106, 0, 317, 1002, 64, 2, 64, 109, 7, 21101, 40, 0, 0, 1008, 1016, 43, 63, 1005, 63, 341, 1001, 64, 1, 64, 1106, 0, 343, 4, 323, 1002, 64, 2, 64, 109, -14, 1208, -1, 31, 63, 1005, 63, 363, 1001, 64, 1, 64, 1106, 0, 365, 4, 349, 1002, 64, 2, 64, 109, 9, 1208, -6, 20, 63, 1005, 63, 387, 4, 371, 1001, 64, 1, 64, 1105, 1, 387, 1002, 64, 2, 64, 109, 2, 2102, 1, -7, 63, 1008, 63, 31, 63, 1005, 63, 413, 4, 393, 1001, 64, 1, 64, 1106, 0, 413, 1002, 64, 2, 64, 109, 21, 2106, 0, -6, 4, 419, 1106, 0, 431, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -25, 2108, 35, -6, 63, 1005, 63, 449, 4, 437, 1106, 0, 453, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 3, 21107, 41, 42, 0, 1005, 1012, 471, 4, 459, 1105, 1, 475, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 7, 21108, 42, 39, -2, 1005, 1017, 495, 1001, 64, 1, 64, 1105, 1, 497, 4, 481, 1002, 64, 2, 64, 109, -8, 1206, 9, 515, 4, 503, 1001, 64, 1, 64, 1106, 0, 515, 1002, 64, 2, 64, 109, 4, 1205, 6, 529, 4, 521, 1105, 1, 533, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -8, 2107, 26, -5, 63, 1005, 63, 553, 1001, 64, 1, 64, 1106, 0, 555, 4, 539, 1002, 64, 2, 64, 109, -6, 2102, 1, 1, 63, 1008, 63, 26, 63, 1005, 63, 575, 1105, 1, 581, 4, 561, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 10, 2101, 0, -8, 63, 1008, 63, 37, 63, 1005, 63, 601, 1105, 1, 607, 4, 587, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -19, 1207, 8, 23, 63, 1005, 63, 627, 1001, 64, 1, 64, 1106, 0, 629, 4, 613, 1002, 64, 2, 64, 109, 18, 21101, 43, 0, 3, 1008, 1013, 43, 63, 1005, 63, 655, 4, 635, 1001, 64, 1, 64, 1106, 0, 655, 1002, 64, 2, 64, 109, -16, 1207, 6, 25, 63, 1005, 63, 677, 4, 661, 1001, 64, 1, 64, 1106, 0, 677, 1002, 64, 2, 64, 109, 25, 21102, 44, 1, -4, 1008, 1015, 44, 63, 1005, 63, 703, 4, 683, 1001, 64, 1, 64, 1106, 0, 703, 1002, 64, 2, 64, 109, 17, 2106, 0, -9, 1106, 0, 721, 4, 709, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -16, 1205, 0, 737, 1001, 64, 1, 64, 1105, 1, 739, 4, 727, 1002, 64, 2, 64, 109, -12, 21107, 45, 44, 5, 1005, 1013, 759, 1001, 64, 1, 64, 1106, 0, 761, 4, 745, 1002, 64, 2, 64, 109, 4, 1201, -8, 0, 63, 1008, 63, 27, 63, 1005, 63, 783, 4, 767, 1106, 0, 787, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -16, 2108, 25, 4, 63, 1005, 63, 803, 1105, 1, 809, 4, 793, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 27, 21102, 46, 1, -5, 1008, 1018, 43, 63, 1005, 63, 829, 1106, 0, 835, 4, 815, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -27, 1202, 8, 1, 63, 1008, 63, 27, 63, 1005, 63, 857, 4, 841, 1105, 1, 861, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 23, 21108, 47, 47, -2, 1005, 1017, 883, 4, 867, 1001, 64, 1, 64, 1106, 0, 883, 1002, 64, 2, 64, 109, -1, 2105, 1, 5, 1001, 64, 1, 64, 1106, 0, 901, 4, 889, 4, 64, 99, 21102, 1, 27, 1, 21102, 915, 1, 0, 1105, 1, 922, 21201, 1, 29589, 1, 204, 1, 99, 109, 3, 1207, -2, 3, 63, 1005, 63, 964, 21201, -2, -1, 1, 21102, 1, 942, 0, 1106, 0, 922, 21202, 1, 1, -1, 21201, -2, -3, 1, 21102, 957, 1, 0, 1105, 1, 922, 22201, 1, -1, -2, 1106, 0, 968, 21202, -2, 1, -2, 109, -3, 2106, 0, 0}

//var program = []int{1102, 34463338, 34463338, 63, 1007, 63, 34463338, 63, 1005, 63, 53, 1101, 3, 0, 1000, 109, 988, 209, 12, 9, 1000, 209, 6, 209, 3, 203, 0, 1008, 1000, 1, 63, 1005, 63, 65, 1008, 1000, 2, 63, 1005, 63, 904, 1008, 1000, 0, 63, 1005, 63, 58, 4, 25, 104, 0, 99, 4, 0, 104, 0, 99, 4, 17, 104, 0, 99, 0, 0, 1101, 0, 708, 1029, 1101, 1, 0, 1021, 1102, 38, 1, 1015, 1101, 25, 0, 1004, 1101, 21, 0, 1018, 1102, 1, 34, 1016, 1101, 0, 713, 1028, 1101, 735, 0, 1024, 1101, 31, 0, 1003, 1102, 1, 24, 1010, 1101, 20, 0, 1011, 1101, 0, 27, 1005, 1102, 726, 1, 1025, 1101, 426, 0, 1027, 1101, 0, 777, 1022, 1102, 1, 32, 1001, 1101, 37, 0, 1009, 1101, 429, 0, 1026, 1102, 1, 36, 1019, 1101, 0, 0, 1020, 1101, 0, 30, 1012, 1101, 0, 770, 1023, 1101, 0, 35, 1014, 1101, 0, 33, 1007, 1102, 23, 1, 1002, 1101, 0, 28, 1017, 1102, 1, 22, 1013, 1102, 39, 1, 1006, 1101, 0, 26, 1000, 1101, 29, 0, 1008, 109, 6, 2102, 1, -1, 63, 1008, 63, 27, 63, 1005, 63, 203, 4, 187, 1106, 0, 207, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -15, 2108, 26, 9, 63, 1005, 63, 225, 4, 213, 1106, 0, 229, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 9, 21101, 40, 0, 10, 1008, 1010, 40, 63, 1005, 63, 251, 4, 235, 1106, 0, 255, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 11, 21108, 41, 40, 0, 1005, 1011, 271, 1106, 0, 277, 4, 261, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -7, 1207, 3, 32, 63, 1005, 63, 297, 1001, 64, 1, 64, 1105, 1, 299, 4, 283, 1002, 64, 2, 64, 109, 3, 1201, -1, 0, 63, 1008, 63, 42, 63, 1005, 63, 323, 1001, 64, 1, 64, 1105, 1, 325, 4, 305, 1002, 64, 2, 64, 109, 2, 2102, 1, -7, 63, 1008, 63, 24, 63, 1005, 63, 345, 1106, 0, 351, 4, 331, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -6, 21107, 42, 43, 8, 1005, 1011, 369, 4, 357, 1106, 0, 373, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -7, 2108, 30, 7, 63, 1005, 63, 393, 1001, 64, 1, 64, 1106, 0, 395, 4, 379, 1002, 64, 2, 64, 109, 18, 21108, 43, 43, -3, 1005, 1011, 413, 4, 401, 1106, 0, 417, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 17, 2106, 0, -4, 1105, 1, 435, 4, 423, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -29, 2107, 26, 2, 63, 1005, 63, 451, 1105, 1, 457, 4, 441, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 20, 1206, -2, 471, 4, 463, 1105, 1, 475, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -9, 1205, 8, 489, 4, 481, 1105, 1, 493, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -12, 1202, -1, 1, 63, 1008, 63, 26, 63, 1005, 63, 515, 4, 499, 1105, 1, 519, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 25, 1205, -6, 531, 1106, 0, 537, 4, 525, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -31, 1208, 8, 31, 63, 1005, 63, 555, 4, 543, 1106, 0, 559, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 13, 1207, 1, 38, 63, 1005, 63, 577, 4, 565, 1106, 0, 581, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 4, 21101, 44, 0, 1, 1008, 1013, 47, 63, 1005, 63, 605, 1001, 64, 1, 64, 1106, 0, 607, 4, 587, 1002, 64, 2, 64, 109, -6, 2107, 38, 0, 63, 1005, 63, 629, 4, 613, 1001, 64, 1, 64, 1106, 0, 629, 1002, 64, 2, 64, 109, 13, 21102, 45, 1, -7, 1008, 1012, 45, 63, 1005, 63, 655, 4, 635, 1001, 64, 1, 64, 1105, 1, 655, 1002, 64, 2, 64, 109, 9, 1206, -7, 667, 1106, 0, 673, 4, 661, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -27, 2101, 0, 7, 63, 1008, 63, 29, 63, 1005, 63, 699, 4, 679, 1001, 64, 1, 64, 1106, 0, 699, 1002, 64, 2, 64, 109, 17, 2106, 0, 10, 4, 705, 1106, 0, 717, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 14, 2105, 1, -8, 4, 723, 1001, 64, 1, 64, 1106, 0, 735, 1002, 64, 2, 64, 109, -21, 1202, -8, 1, 63, 1008, 63, 34, 63, 1005, 63, 755, 1105, 1, 761, 4, 741, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 18, 2105, 1, -6, 1001, 64, 1, 64, 1106, 0, 779, 4, 767, 1002, 64, 2, 64, 109, -15, 1201, -6, 0, 63, 1008, 63, 29, 63, 1005, 63, 801, 4, 785, 1105, 1, 805, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -14, 1208, 0, 24, 63, 1005, 63, 825, 1001, 64, 1, 64, 1106, 0, 827, 4, 811, 1002, 64, 2, 64, 109, 15, 21102, 46, 1, -2, 1008, 1013, 49, 63, 1005, 63, 847, 1106, 0, 853, 4, 833, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -17, 2101, 0, 2, 63, 1008, 63, 23, 63, 1005, 63, 873, 1106, 0, 879, 4, 859, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 16, 21107, 47, 46, 2, 1005, 1016, 899, 1001, 64, 1, 64, 1105, 1, 901, 4, 885, 4, 64, 99, 21101, 0, 27, 1, 21101, 0, 915, 0, 1106, 0, 922, 21201, 1, 55486, 1, 204, 1, 99, 109, 3, 1207, -2, 3, 63, 1005, 63, 964, 21201, -2, -1, 1, 21102, 942, 1, 0, 1105, 1, 922, 22102, 1, 1, -1, 21201, -2, -3, 1, 21101, 0, 957, 0, 1105, 1, 922, 22201, 1, -1, -2, 1105, 1, 968, 22101, 0, -2, -2, 109, -3, 2106, 0, 0}

//var program = []int{109, 2000, 109, 19, 204, -34, 99}

func main() {
	fullMem := program
	for i := 1; i < 3000; i++ {
		fullMem = append(fullMem, 0)
	}
	p := process{
		memory:   copyProgram(fullMem),
		finished: func() {},
		output:   make(chan int, 10),
		input:    make(chan int, 10),
	}
	p.writeOutput = func() {
		o := <-p.output
		fmt.Printf("ip: %d, Output: %d\n", p.ip, o)
	}
	p.readInput = func() {
		p.input <- 1
	}

	p.Run()
}
