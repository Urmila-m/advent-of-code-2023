package lib

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ParseFile(filePath string) (string, map[string]map[string]string) {
	file, fErr := os.Open(filePath)

	if fErr != nil {
		log.Fatal(fErr)
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(file)

	instructions := ""
	network := make(map[string]map[string]string)
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		instructions = scanner.Text()
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			equalSplit := strings.Split(line, " = ")
			value := equalSplit[0]

			commaSplit := strings.Split(strings.Trim(equalSplit[1], "()"), ", ")
			left := commaSplit[0]
			right := commaSplit[1]
			network[value] = map[string]string{"left": left, "right": right}
		}
	}
	return instructions, network
}

type Path struct {
	Source     string
	Dest       string
	NoOfSteps  int
	Trajectory []string
}

func FindNumOfSteps(instructions string, network map[string]map[string]string, from string, dest string) Path {
	start := from
	steps := 0
	path := make([]string, 0)
	for {
		if string(instructions[steps%len(instructions)]) == "L" {
			steps++
			next := network[from]["left"]
			path = append(path, next)
			if next == dest {
				break
			} else {
				from = next
			}
		} else {
			steps++
			next := network[from]["right"]
			path = append(path, next)
			if next == dest {
				break
			} else {
				from = next
			}
		}
	}
	return Path{Source: start, Dest: dest, NoOfSteps: steps, Trajectory: path}
}

func FindAllPaths(instructions string, network map[string]map[string]string) []Path {
	allSrcToDst := make([]Path, 0)
	for node, _ := range network {
		if string(node[2]) == "A" {
			//find number of steps
			from := node
			start := node
			path := make([]string, 0)
			dest := ""
			steps := 0
			for {
				if string(instructions[steps%len(instructions)]) == "L" {
					steps++
					next := network[from]["left"]
					path = append(path, next)
					if string(next[2]) == "Z" {
						dest = next
						break
					} else {
						from = next
					}
				} else {
					steps++
					next := network[from]["right"]
					path = append(path, next)
					if string(next[2]) == "Z" {
						dest = next
						break
					} else {
						from = next
					}
				}
			}
			if dest != "" {
				allSrcToDst = append(allSrcToDst, Path{Source: start, Dest: dest, NoOfSteps: steps, Trajectory: path})
			}
		}
	}
	return allSrcToDst
}

func FindGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// FindLCM find The Least Common Multiple (LCM) via GCD
func FindLCM(numbers []int) int {
	result := numbers[0] * numbers[1] / FindGCD(numbers[0], numbers[1])

	remainingNum := numbers[2:]
	for i := 0; i < len(remainingNum); i++ {

		result = FindLCM([]int{result, remainingNum[i]})
	}
	return result
}
