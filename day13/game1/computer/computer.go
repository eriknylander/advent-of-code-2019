package computer

import "fmt"

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

type Process struct {
	id       int
	memory   []int
	ip       int
	rbs      int
	input    chan int
	output   chan int
	finished func()
}

func NewProcess(id int, mem []int, input, output chan int, finished func()) Process {
	return Process{
		id:       id,
		memory:   copyProgram(mem),
		input:    input,
		output:   output,
		finished: finished,
	}
}

func (p *Process) Run() {
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

func (p *Process) processInstruction(ins instruction) {
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

func (p *Process) getValue(mode, param int) int {
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

func (p *Process) getWriteAddress(mode, param int) int {
	switch mode {
	case positionMode:
		return param
	case relativeMode:
		return p.rbs + param
	default:
		panic("Invalid mode")
	}
}
