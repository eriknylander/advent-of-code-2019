package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	youID   = "YOU"
	santaID = "SAN"
)

type orbiterBody struct {
	name     string
	parent   *orbiterBody
	children []*orbiterBody
}


func main() {
	file, err := os.Open("./santa.txt")
	if err != nil {
		panic(fmt.Errorf("error: %s\n", err))
	}
	defer file.Close()

	allOrbiters := map[string]*orbiterBody{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pn := strings.Split(scanner.Text(), ")")[0]
		on := strings.Split(scanner.Text(), ")")[1]

		parent := allOrbiters[pn]
		orbiter := allOrbiters[on]

		if parent == nil {
			parent = &orbiterBody{name: pn, parent: nil, children: []*orbiterBody{}}
		}

		if orbiter == nil {
			orbiter = &orbiterBody{
				name:   on,
				parent: parent,
			}
		} else {
			orbiter.parent = parent
		}

		parent.children = append(parent.children, orbiter)
		allOrbiters[on] = orbiter
		allOrbiters[pn] = parent

	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error: %s\n", err))
	}

	santa := allOrbiters[santaID]
	you := allOrbiters[youID]

	santaStepsMap := make(map[string]int)
	santaTotalSteps := 0
	santaParent := santa.parent
	for santaParent != nil {
		santaStepsMap[santaParent.name]= santaTotalSteps
		santaParent = santaParent.parent
		santaTotalSteps++
	}

	youSteps := 0
	santaSteps := 0
	youParent := you.parent
	for youParent != nil {
		if s, onSantasPath := santaStepsMap[youParent.name]; onSantasPath {
			santaSteps = s
			break
		}
		youParent = youParent.parent
		youSteps++
	}
	fmt.Println(youSteps+ santaSteps)
}

func santaInOrbit(b orbiterBody) bool {
	if b.name == santaID {
		return true
	}

	if len(b.children) == 0 {
		return false
	}

	foundSanta := false
	for _, c := range b.children {
		foundSanta = santaInOrbit(*c)
		if foundSanta {
			break
		}
	}

	return foundSanta
}
