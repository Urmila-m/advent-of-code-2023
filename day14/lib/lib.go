package lib

import (
	"bufio"
	"fmt"
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

func RollToNorth(lines [][]rune) [][]rune {
	numRows := len(lines)
	numCols := len(lines[0])
	for i := 0; i < numCols; i++ {
		for j := 0; j < numRows-1; j++ {
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

func RollToSouth(lines [][]rune) [][]rune {
	numRows := len(lines)
	numCols := len(lines[0])
	for i := 0; i < numCols; i++ {
		for j := numRows - 1; j > 0; j-- {
			if lines[j][i] == '.' {
				for k := int(math.Max(0, float64(j-1))); k >= 0; k-- {
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

func RollToWest(lines [][]rune) [][]rune {
	numRows := len(lines)
	numCols := len(lines[0])
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if lines[i][j] == '.' {
				for k := int(math.Min(float64(j+1), float64(numCols-1))); k < numCols; k++ {
					if lines[i][k] == 'O' {
						lines[i][j] = 'O'
						lines[i][k] = '.'
						break
					} else if lines[i][k] == '#' {
						break
					}
				}
			}
		}
	}
	return lines
}

func RollToEast(lines [][]rune) [][]rune {
	numRows := len(lines)
	numCols := len(lines[0])
	for i := 0; i < numRows; i++ {
		for j := numCols - 1; j > 0; j-- {
			if lines[i][j] == '.' {
				for k := int(math.Max(0, float64(j-1))); k >= 0; k-- {
					if lines[i][k] == 'O' {
						lines[i][j] = 'O'
						lines[i][k] = '.'
						break
					} else if lines[i][k] == '#' {
						break
					}
				}
			}
		}
	}
	return lines
}

func DisplayMatrix(matrix [][]rune) {
	for _, line := range matrix {
		fmt.Print("{\t")
		for _, char := range line {
			fmt.Printf("'%c',\t", char)
		}
		fmt.Print("},\n")
	}
}

func CalcNorthLoad(matrix [][]rune) int {
	totalLoad := 0
	totalLines := len(matrix)
	for i, line := range matrix {
		for _, char := range line {
			if char == 'O' {
				totalLoad += totalLines - i
			}
		}
	}
	return totalLoad
}

func RollACycle(matrix [][]rune) [][]rune {
	return RollToEast(RollToSouth(RollToWest(RollToNorth(matrix))))
}

func DeepCloneMatrix(matrix [][]rune) [][]rune {
	clone := make([][]rune, 0)
	for _, line := range matrix {
		cloneLine := make([]rune, 0)
		cloneLine = append(cloneLine, line...)
		clone = append(clone, cloneLine)
	}
	return clone
}

func FindCycleRepeatPosition(matrix [][]rune) (repeatsAfter int, repeatStartPoint int, allNonRepeatMatrices [][][]rune) {
	// find after how many cycles, the pattern goes back to the initial state
	matrices := [][][]rune{
		matrix,
	}
	for i := 1; i <= 1000000000; i++ {
		rolled := RollACycle(DeepCloneMatrix(matrices[i-1]))
		for j, mat := range matrices {
			if isMatrixEqual(mat, rolled) {
				return i, j, matrices
			}
		}
		matrices = append(matrices, rolled)
	}
	return -1, -1, matrices
}

func isMatrixEqual(mat1, mat2 [][]rune) bool {
	for i := 0; i < len(mat1); i++ {
		for j := 0; j < len(mat1[0]); j++ {
			if mat1[i][j] != mat2[i][j] {
				return false
			}
		}
	}
	return true
}

func FindNorthLoadAfterSpin(n int, filePath string) int {
	matrix := ParseFile(filePath)
	repeatsAfter, repeatStartPoint, allNonRepeatMatrices := FindCycleRepeatPosition(matrix)
	if repeatsAfter != -1 {
		repeatPeriod := repeatsAfter - repeatStartPoint
		/* The pattern shows arithemetic series nature of
		f(n) = repeatStartPoint + repeatPeriod * n
		where
		f(n) = number of cycles
		n = {0, 1, 2, ...}

		For example, in the test case, repeatStartPoint = 4, repeatPeriod = 7.
		So, the same pattern repeats at 4, 4 + 7 * 1 , 4 + 7 * 2, ...

		From this, if the number of cycles given in the question falls under any of the above checkpoints, allNonRepeatingMatrices[4] (allNonRepeatingMatrices[repeatStartPoint]) should give the last state after all the rolling.

		For the number of cycles to co-incide with the checkpoints, n must be an integer. For n to be an integer, (f(n) - repeatStartPoint) must be divisible by repeatPeriod i.e.
		if (f(n) - repeatStartPoint)%repeatPeriod == 0, allNonRepeatingMatrices[repeatStartPoint] should give the last state.
		Similarly, if (f(n) - repeatStartPoint)%RepeatPeriod == 1, allNonRepeatingMatrices[repeatStartPoint+1] should give the last state.

		In general, allNonRepeatingMatrices[repeatStartPoint+((f(n) - repeatStartPoint)%RepeatPeriod)] should give the last state.
		*/
		lastState := allNonRepeatMatrices[repeatStartPoint+((n-repeatStartPoint)%repeatPeriod)]
		return CalcNorthLoad(lastState)
	}
	return -1
}
