import sys
import math
import io

def add(program, src1, src2):
    return program[src1] + program[src2]

def mul(program, src1, src2):
    return program[src1] * program[src2]

def runProgram(program):
    ip = 0
    while program[ip] != 99:
        op = program[ip]
        sourceAddr1 = program[ip+1]
        sourceAddr2 = program[ip+2]
        dstAddr = program[ip+3]
        if len(program) <= dstAddr:
            return 0
        if op == 1:
            sum = add(program, sourceAddr1, sourceAddr2)
            program[dstAddr] = sum
        if op == 2:
            prod = mul(program, sourceAddr1, sourceAddr2)
            program[dstAddr] = prod
        ip = ip+4
    return program[0]

def resetMemory(base):
    a = []
    for i in range(0, len(base)):
        a.append(base[i])
    return a

baseProgram = [1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,19,6,23,2,23,6,27,2,6,27,31,2,13,31,35,1,9,35,39,2,10,39,43,1,6,43,47,1,13,47,51,2,6,51,55,2,55,6,59,1,59,5,63,2,9,63,67,1,5,67,71,2,10,71,75,1,6,75,79,1,79,5,83,2,83,10,87,1,9,87,91,1,5,91,95,1,95,6,99,2,10,99,103,1,5,103,107,1,107,6,111,1,5,111,115,2,115,6,119,1,119,6,123,1,123,10,127,1,127,13,131,1,131,2,135,1,135,5,0,99,2,14,0,0]
#baseProgram = [1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,19,6,23,2,23,6,27,2,6,27,31,2,13,31,35,1,9,35,39,2,10,39,43,1,6,43,47,1,13,47,51,2,6,51,55,2,55,6,59,1,59,5,63,2,9,63,67,1,5,67,71,2,10,71,75,1,6,75,79,1,79,5,83,2,83,10,87,1,9,87,91,1,5,91,95,1,95,6,99,2,10,99,103,1,5,103,107,1,107,6,111,1,5,111,115,2,115,6,119,1,119,6,123,1,123,10,127,1,127,13,131,1,131,2,135,1,135,5,0,99,2,14,0,0]
#baseProgram = [1,0,0,0,99]
#baseProgram = [2,3,0,3,99]
#baseProgram = [2,4,4,5,99,0]
#baseProgram = [1,1,1,4,99,5,6,0,99]

for noun in range(0,100):
    for verb in range (0, 100):
        print("Noun ", noun)
        print("Verb ", verb)
        p = resetMemory(baseProgram)
        p[1] = noun
        p[2] = verb
        result = runProgram(p)
        print(result)
        if result == 19690720:
            sys.exit(0)

