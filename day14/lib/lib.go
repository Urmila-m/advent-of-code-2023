package lib

import (
	"bufio"
	"log"
	"math"
	"os"
)

func ParseFile(filePath string) [][]rune {
	file, fErr := os.Open(filePath)
	if fErr != nil {
		log.Fatal(fErr)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	lines := make([][]rune, 0)
	for scanner.Scan() {
		line := make([]rune, 0)
		for _, char := range scanner.Text() {
			line = append(line, char)
		}
		lines = append(lines, line)
	}
	return lines
}

func RollToNorth(filePath string) [][]rune {
	lines := ParseFile(filePath)
	numRows := len(lines)
	numCols := len(lines[0])
	for i := 0; i < numCols; i++ {
		for j := 0; j < numRows; j++ {
			if lines[j][i] == '.' {
				for k := int(math.Min(float64(numRows-1), float64(j+1))); k < numRows; k++ {
					if lines[k][i] == 'O' {
						lines[j][i] = 'O'
						lines[k][i] = '.'
						break
					} else if lines[k][i] == '#' {
						break
					}
				}
			}
		}
	}
	return lines
}

func CalcNorthLoad(filePath string) int {
	totalLoad := 0
	allLines := RollToNorth(filePath)
	totalLines := len(allLines)
	for i, line := range allLines {
		for _, char := range line {
			if char == 'O' {
				totalLoad += totalLines - i
			}
		}
	}
	return totalLoad
}
