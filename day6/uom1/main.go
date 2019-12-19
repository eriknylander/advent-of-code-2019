package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type orbiterBody struct {
	name     string
	parent   *orbiterBody
	children []*orbiterBody
}


func main() {
	file, err := os.Open("./orbiters.txt")
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

	nc := []orbiterBody {}
	for _, b := range allOrbiters {
		if b.parent  != nil {
			nc = append(nc, *b)
		}
	}

	sum := 0
	for _, o := range nc {
		sum+=orbitCount(o)
	}

	fmt.Println(sum)
}

func orbitCount(o orbiterBody) int {
	c := 1
	p := o.parent
	for p.parent != nil {
		c++
		p = p.parent
	}

	return c
}
