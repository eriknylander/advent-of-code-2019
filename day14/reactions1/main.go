package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type elementQuantity struct {
	element  string
	quantity int
}

type reaction struct {
	inputs []elementQuantity
	result elementQuantity
}

type conversion struct {
	ores elementQuantity
	em   elementQuantity
}

func parseElementQuantity(r string) elementQuantity {
	r = strings.Trim(r, " ")

	element := strings.Split(r, " ")[1]
	quantity, err := strconv.Atoi(strings.Split(r, " ")[0])
	if err != nil {
		panic(err)
	}

	return elementQuantity{element: element, quantity: quantity}
}

func readInputFile() []reaction {
	reactions := []reaction{}
	file, err := os.Open("./example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		outputStr := strings.Split(line, "=>")[1]
		inputsStr := strings.Split(line, "=>")[0]

		result := parseElementQuantity(outputStr)

		inputs := []elementQuantity{}
		for _, inputStr := range strings.Split(inputsStr, ",") {
			inputs = append(inputs, parseElementQuantity(inputStr))
		}

		reactions = append(reactions, reaction{
			inputs: inputs,
			result: result,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return reactions
}

func main() {
	reactions := readInputFile()
	conversionTable := make(map[string]conversion)

	// Find initial conversions
	for _, r := range reactions {
		if len(r.inputs) == 1 && strings.EqualFold(r.inputs[0].element, "ORE") {
			conversionTable[r.result.element] = conversion{ores: r.inputs[0], em: r.result}
		}
	}

	conversionTableFinished := false
	for !conversionTableFinished {
		for _, r := range reactions {
			if _, ok := conversionTable[r.result.element]; !ok {
				ores := 0
				for _, ie := range r.inputs {
					if ce, ok := conversionTable[ie.element]; ok {
						o := ce.ores.quantity
						for o/ie.quantity < 1 {
							o += ce.ores.quantity
						}
						ores += o
					} else {
						break
					}
				}
				conversionTable[r.result.element] = conversion{
					ores: elementQuantity{element: "ORE", quantity: ores},
					em:   r.result,
				}
			}
		}
	}

	fmt.Println(conversionTable)
}
