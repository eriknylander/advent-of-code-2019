import sys
import math
import io

def calcFuel(mass):
    return math.floor(int(moduleMass) / 3) - 2

f = open("./fuel.txt", "r")

sum = 0
for moduleMass in f:
    fuel = calcFuel(moduleMass)
    sum = sum+fuel

print(sum)