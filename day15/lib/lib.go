package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindHash(input string) int {
	currentValue := 0
	for _, char := range input {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func ParseFile(filePath string) string {
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

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func FindHashSum(initSequence string) int {
	sum := 0
	individualSteps := strings.Split(initSequence, ",")
	for _, step := range individualSteps {
		sum += FindHash(step)
	}
	return sum
}

func FindResultingConfiguration(initSequence string) map[int][]map[string]int {
	resultConfig := make(map[int][]map[string]int)
	/*
		resultConfig format:
			{
				boxNo: [{label: focalLength}]
			}
		example:
			{
				1: [{"rn": 1}, {"cm": 2}]
				3: [{"pc": 4}]
			}
	*/
	for _, seq := range strings.Split(initSequence, ",") {
		operator := FindOperatorInSeq(seq)
		operatorSplit := strings.Split(seq, operator)
		label := operatorSplit[0]
		boxNo := FindHash(label)
		focalLength := -1
		if operator == "=" {
			if num, convertErr := strconv.Atoi(string(operatorSplit[1][0])); convertErr == nil {
				focalLength = num
			}
		}

		_, exists := resultConfig[boxNo]
		if !exists {
			resultConfig[boxNo] = make([]map[string]int, 0)
		}
		resultConfig[boxNo] = PerformOperation(operator, resultConfig[boxNo], exists, label, focalLength)
	}
	return resultConfig
}

func FindOperatorInSeq(seq string) string {
	if strings.Contains(seq, "=") {
		return "="
	} else if strings.Contains(seq, "-") {
		return "-"
	}
	return ""
}

func PerformOperation(operator string, box []map[string]int, hasLens bool, label string, focalLength int) []map[string]int {
	if operator == "-" {
		if !hasLens {
			return box
		}
		if index, exists := SearchLabelInBox(label, box); exists {
			return append(box[:index], box[index+1:]...)
		}
	} else if operator == "=" {
		if !hasLens {
			return append(box, map[string]int{label: focalLength})
		}
		if index, exists := SearchLabelInBox(label, box); exists {
			box[index][label] = focalLength
			return box
		} else {
			return append(box, map[string]int{label: focalLength})
		}
	}
	return box
}

func SearchLabelInBox(label string, box []map[string]int) (int, bool) {
	for i, lens := range box {
		if _, exists := lens[label]; exists {
			return i, exists
		}
	}
	return -1, false
}

func CalcFocusingPower(boxNo int, slotNumber int, focalLength int) int {
	return (boxNo + 1) * slotNumber * focalLength
}

func SumAllFocusingPower(resultConfig map[int][]map[string]int) int {
	sum := 0
	for boxNo, box := range resultConfig {
		for i, lens := range box {
			for _, focalLength := range lens {
				sum += CalcFocusingPower(boxNo, i+1, focalLength)
			}
		}
	}
	return sum
}
