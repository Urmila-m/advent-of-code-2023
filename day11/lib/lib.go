package lib

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
)

type Point struct {
	x int
	y int
}

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

	matrix := make([][]rune, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		runesInLine := make([]rune, 0)
		hasGalaxy := false
		for _, char := range line {
			if !hasGalaxy && char == '#' {
				hasGalaxy = true
			}
			runesInLine = append(runesInLine, char)
		}
		matrix = append(matrix, runesInLine)
		if !hasGalaxy {
			matrix = append(matrix, runesInLine)
		}
	}
	return matrix
}

func AddColumnForEmpty(matrix [][]rune) [][]rune {
	transposeMatrix := FindMatrixTranspose(matrix)
	emptyRowsInTranspose := make([]int, 0)
	for i, line := range transposeMatrix {
		if !slices.Contains(line, '#') {
			emptyRowsInTranspose = append(emptyRowsInTranspose, i)
		}
	}

	for i, index := range emptyRowsInTranspose {
		transposeMatrix = slices.Insert(transposeMatrix, index+1, transposeMatrix[index])
		for j := i + 1; j < len(emptyRowsInTranspose); j++ {
			emptyRowsInTranspose[j] = emptyRowsInTranspose[j] + 1
		}
	}

	return FindMatrixTranspose(transposeMatrix)
}

func FindMatrixTranspose(matrix [][]rune) [][]rune {
	transposeMatrix := make([][]rune, 0)

	for i := 0; i < len(matrix[0]); i++ {
		tempRow := make([]rune, 0)
		for j := 0; j < len(matrix); j++ {
			tempRow = append(tempRow, matrix[j][i])
		}
		transposeMatrix = append(transposeMatrix, tempRow)
	}
	return transposeMatrix
}

func FindAllGalaxies(filePath string) []Point {
	effectiveMatrix := AddColumnForEmpty(ParseFile(filePath))

	allGalaxies := make([]Point, 0)

	for i, row := range effectiveMatrix {
		for j, elem := range row {
			if elem == '#' {
				allGalaxies = append(allGalaxies, Point{x: i, y: j})
			}
		}
	}
	return allGalaxies
}

func FindDistance(src, dst Point) int {
	rowDistance := math.Abs(float64(src.x) - float64(dst.x))
	colDistance := math.Abs(float64(src.y) - float64(dst.y))
	return int(rowDistance + colDistance)
}

func FindSumOfGalaxyDistances(filePath string) int {
	distanceSum := 0
	allGalaxies := FindAllGalaxies(filePath)

	for i := 0; i < len(allGalaxies); i++ {
		for j := i + 1; j < len(allGalaxies); j++ {
			distanceSum += FindDistance(allGalaxies[i], allGalaxies[j])
		}
	}
	return distanceSum
}

func ParseFile2(filePath string) ([][]rune, []int, []int) {
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

	matrix := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	emptyRows := make([]int, 0)
	rowNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		runesInLine := make([]rune, 0)
		hasGalaxy := false
		for _, char := range line {
			if !hasGalaxy && char == '#' {
				hasGalaxy = true
			}
			runesInLine = append(runesInLine, char)
		}
		matrix = append(matrix, runesInLine)
		if !hasGalaxy {
			emptyRows = append(emptyRows, rowNum)
		}
		rowNum++
	}

	emptyCols := make([]int, 0)
	for i := 0; i < len(matrix[0]); i++ {
		hasGalaxy := false
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			emptyCols = append(emptyCols, i)
		}
	}
	return matrix, emptyRows, emptyCols
}

func FindAllGalaxies2(filePath string) ([]Point, []int, []int) {
	matrix, emptyRows, emptyCols := ParseFile2(filePath)

	allGalaxies := make([]Point, 0)

	for i, row := range matrix {
		for j, elem := range row {
			if elem == '#' {
				allGalaxies = append(allGalaxies, Point{x: i, y: j})
			}
		}
	}
	return allGalaxies, emptyRows, emptyCols
}

func FindDistance2(src, dst Point, emptyRows []int, emptyCols []int, expandRatio int) int {
	rowDistance := 0
	for i := src.x + 1; i <= dst.x; i++ {
		if slices.Contains(emptyRows, i) {
			rowDistance += expandRatio
		} else {
			rowDistance += 1
		}
	}
	colDistance := 0
	for i := int(math.Min(float64(src.y), float64(dst.y))) + 1; i <= int(math.Max(float64(dst.y), float64(src.y))); i++ {
		if slices.Contains(emptyCols, i) {
			colDistance += expandRatio
		} else {
			colDistance += 1
		}
	}
	return rowDistance + colDistance
}

func FindSumDistance2(filePath string, expandRatio int) int {
	sum := 0
	allGalaxies, emptyRows, emptyCols := FindAllGalaxies2(filePath)
	for i := 0; i < len(allGalaxies); i++ {
		for j := i + 1; j < len(allGalaxies); j++ {
			sum += FindDistance2(allGalaxies[i], allGalaxies[j], emptyRows, emptyCols, expandRatio)
		}
	}
	return sum
}
