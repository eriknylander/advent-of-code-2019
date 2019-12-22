package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// )

// func buildSystem(path string) [][]string {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		panic(fmt.Errorf("error: %s\n", err))
// 	}
// 	defer file.Close()

// 	system := [][]string{}

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		system = append(system, strings.Split(scanner.Text(), ""))
// 	}

// 	return system
// }

// type asteroid struct {
// 	x                int
// 	y                int
// 	asteroidsInSight []asteroid
// }

// func main() {
// 	system := buildSystem("../example5.txt")
// 	asteroidsInSystem := []asteroid{}
// 	for y := 0; y < len(system); y++ {
// 		for x := 0; x < len(system[y]); x++ {
// 			// if system[y][x] == "#" {
// 			// 	asteroidsInSystem = append(asteroidsInSystem, asteroid{
// 			// 		x:                x,
// 			// 		y:                y,
// 			// 		asteroidsInSight: asteroidsInSight(x, y, system),
// 			// 	})
// 			// }
// 		}
// 	}

// 	sort.Slice(asteroidsInSystem, func(i, j int) bool {
// 		return len(asteroidsInSystem[i].asteroidsInSight) > len(asteroidsInSystem[j].asteroidsInSight)
// 	})

// 	a := asteroidsInSight(11, 13, system)
// 	fmt.Println(a[0])
// 	fmt.Println(a[1])
// 	fmt.Println(a[2])
// 	fmt.Println(a[9])

// 	// best := asteroidsInSystem[0]
// 	// fmt.Printf("%d,%d, %d\n", best.x, best.y, len(best.asteroidsInSight))

// 	// for _, a := range asteroidsInSystem {
// 	// 	fmt.Printf("%d,%d, %d\n", a.x, a.y, len(a.asteroidsInSight))
// 	// }
// }

// func asteroidsInSight(x, y int, system [][]string) []asteroid {
// 	a := []asteroid{}

// 	for j := y - 1; j >= 0; j-- {
// 		if system[j][x] == "#" {
// 			a = append(a, asteroid{
// 				x: x,
// 				y: j,
// 			})
// 			break
// 		}
// 	}

// 	for i := x + 1; i < len(system[0]); i++ {
// 		for j := 0; j < y; j++ {
// 			lines := map[float64]bool{}
// 			if system[j][i] == "#" {
// 				if j != y {
// 					dy := float64(j - y)
// 					dx := float64(i - x)
// 					k := dy / dx
// 					if _, exists := lines[k]; !exists {
// 						lines[k] = true
// 						a = append(a, asteroid{
// 							x: i,
// 							y: j,
// 						})
// 					}
// 					continue
// 				}
// 			}
// 		}
// 	}

// 	for i := x + 1; i < len(system[0]); i++ {
// 		if system[y][i] == "#" {
// 			a = append(a, asteroid{
// 				x: i,
// 				y: y,
// 			})
// 			break
// 		}
// 	}

// 	for j := y + 1; j < len(system); j++ {
// 		for i := len(system[0]) - 1; i > x; i-- {
// 			lines := map[float64]bool{}
// 			if system[j][i] == "#" {
// 				if j != y {
// 					dy := float64(j - y)
// 					dx := float64(i - x)
// 					k := dy / dx
// 					if _, exists := lines[k]; !exists {
// 						lines[k] = true
// 						a = append(a, asteroid{
// 							x: i,
// 							y: j,
// 						})
// 					}
// 					continue
// 				}
// 			}
// 		}
// 	}

// 	for j := y + 1; j < len(system); j++ {
// 		if system[j][x] == "#" {
// 			a = append(a, asteroid{
// 				x: x,
// 				y: j,
// 			})
// 			break
// 		}
// 	}

// 	for i := x - 1; i >= 0; i-- {
// 		for j := len(system) - 1; j > y; j-- {
// 			lines := map[float64]bool{}
// 			if system[j][i] == "#" {
// 				if j != y {
// 					dy := float64(j - y)
// 					dx := float64(i - x)
// 					k := dy / dx
// 					if _, exists := lines[k]; !exists {
// 						lines[k] = true
// 						a = append(a, asteroid{
// 							x: i,
// 							y: j,
// 						})
// 					}
// 					continue
// 				}
// 			}
// 		}
// 	}

// 	for i := x - 1; i >= 0; i-- {
// 		if system[y][i] == "#" {
// 			a = append(a, asteroid{
// 				x: i,
// 				y: y,
// 			})
// 			break
// 		}
// 	}

// 	for j := y - 1; j >= 0; j-- {
// 		for i := 0; i < x; i++ {
// 			lines := map[float64]bool{}
// 			if system[j][i] == "#" {
// 				if j != y {
// 					dy := float64(j - y)
// 					dx := float64(i - x)
// 					k := dy / dx
// 					if _, exists := lines[k]; !exists {
// 						lines[k] = true
// 						a = append(a, asteroid{
// 							x: i,
// 							y: j,
// 						})
// 					}
// 					continue
// 				}
// 			}
// 		}
// 	}

// 	return a
// }
