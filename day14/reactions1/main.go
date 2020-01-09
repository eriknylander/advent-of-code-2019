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

type rootNode struct {
	edges []weightedNode
}

type weightedNode struct {
	weight  int
	element string
	edges   []weightedNode
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

func readInputFile() map[string]reaction {
	reactions := make(map[string]reaction)
	file, err := os.Open("./example2.txt")
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

		reactions[result.element] = reaction{
			inputs: inputs,
			result: result,
		}
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

	fuelNode := rootNode{}
	fuelNode.edges = buildReactionTree(reactions["FUEL"].inputs, reactions, conversionTable)

	fmt.Println(conversionTable)
}

func buildReactionTree(inputs []elementQuantity, reactions map[string]reaction, conversionTable map[string]conversion) []weightedNode {
	nodes := []weightedNode{}
	for _, in := range inputs {
		if _, ok := conversionTable[in.element]; ok {
			nodes = append(nodes, weightedNode{
				weight:  in.quantity,
				element: in.element,
			})
		} else {
			r := reactions[in.element]
			nodes = append(nodes, weightedNode{
				weight:  in.quantity,
				element: in.element,
				edges:   buildReactionTree(r.inputs, reactions, conversionTable),
			})
		}
	}
	return nodes
}
