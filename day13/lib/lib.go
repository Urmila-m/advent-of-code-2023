package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ParseFile(filePath string) [][]string {
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
	chunks := make([][]string, 0)
	chunk := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			chunks = append(chunks, chunk)
			chunk = make([]string, 0)
		} else {
			chunk = append(chunk, line)
		}
	}
	chunks = append(chunks, chunk)
	return chunks
}

func FindHorizontalMirror(chunk []string) (allMirrorsX []int, atLeastOneMirrorExists bool) {
	allMirrorsX = make([]int, 0)
	atLeastOneMirrorExists = false
	for i := 0; i < len(chunk)-1; i++ {
		if chunk[i] == chunk[i+1] {
			exists := true
			j := 0
			for (i+1+j) < len(chunk) && (i-j) >= 0 {
				upper := chunk[i-j]
				lower := chunk[i+1+j]
				if upper != lower {
					exists = false
					break
				}
				j++
			}
			if exists {
				atLeastOneMirrorExists = true
				mirrorX := i + 1
				allMirrorsX = append(allMirrorsX, mirrorX)
			}
		}
	}
	return allMirrorsX, atLeastOneMirrorExists
}

func FindMatrixTranspose(matrix []string) []string {
	transposeMatrix := make([]string, 0)

	for i := 0; i < len(matrix[0]); i++ {
		tempRow := ""
		for j := 0; j < len(matrix); j++ {
			tempRow = fmt.Sprintf("%s%c", tempRow, matrix[j][i])
		}
		transposeMatrix = append(transposeMatrix, tempRow)
	}
	return transposeMatrix
}

func FindVerticalMirror(chunk []string) (mirrorsY []int, exists bool) {
	transposeMatrix := FindMatrixTranspose(chunk)
	return FindHorizontalMirror(transposeMatrix)
}

func SummarizeNotes(filePath string) int {
	sum := 0
	for _, chunk := range ParseFile(filePath) {
		if mirrorsX, exists := FindHorizontalMirror(chunk); exists {
			for _, mirrorX := range mirrorsX {
				sum += 100 * mirrorX
			}
		} else {
			mirrorsY, verticalExists := FindVerticalMirror(chunk)
			if verticalExists {
				for _, mirrorY := range mirrorsY {
					sum += mirrorY
				}
			}
		}
	}
	return sum
}

func CompareString(str1 string, str2 string) []int {
	noMatch := make([]int, 0)
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			noMatch = append(noMatch, i)
		}
	}
	return noMatch
}

func FindSmudgeHorizontalMirror(chunk []string) (mirrorXWithSmudge int, existsWithSmudge bool) {
	smudgeFound := false
	mirrorXWithSmudge = 0
	prevMirrorsX, prevExists := FindHorizontalMirror(chunk)
	for i := 0; i < len(chunk)-1; i++ {
		for j := i + 1; j < len(chunk); j++ {
			noMatches := CompareString(chunk[i], chunk[j])
			if len(noMatches) == 1 {
				tmp := chunk[i]
				chunk[i] = chunk[j]
				mirrorsX, exists := FindHorizontalMirror(chunk)
				if exists {
					if prevExists {
						for _, mirrorX := range mirrorsX {
							if mirrorX != prevMirrorsX[0] {
								mirrorXWithSmudge = mirrorX
							}
						}
					} else {
						mirrorXWithSmudge = mirrorsX[0]
					}
					smudgeFound = true
					break
				} else {
					chunk[i] = tmp
				}
			}
		}
		if smudgeFound {
			break
		}
	}
	return mirrorXWithSmudge, smudgeFound
}

func FindSmudgeVerticalMirror(chunk []string) (mirrorYWithSmudge int, existsWithSmudge bool) {
	transposeMatrix := FindMatrixTranspose(chunk)
	return FindSmudgeHorizontalMirror(transposeMatrix)
}

func SummarizeNotesWithSmudge(filePath string) int {
	sum := 0
	for _, chunk := range ParseFile(filePath) {
		if mirrorX, exists := FindSmudgeHorizontalMirror(chunk); exists {
			sum += 100 * mirrorX
		} else {
			mirrorY, verticalExists := FindSmudgeVerticalMirror(chunk)
			if verticalExists {
				if verticalExists {
					sum += mirrorY
				}
			}
		}
	}
	return sum
}
