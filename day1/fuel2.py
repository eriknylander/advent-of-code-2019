import sys
import math
import io

def calcFuel(mass):
    return math.floor(mass / 3) - 2

def calcFuelRecursive(mass):
    f = calcFuel(mass)
    if f <= 0:
        return 0
    else:
        return f + calcFuelRecursive(f)


f = open("./fuel.txt", "r")

sum = 0
for moduleMass in f:
    massInt = int(moduleMass)
    fuel = calcFuelRecursive(massInt)
    sum = sum+fuel

print(sum)