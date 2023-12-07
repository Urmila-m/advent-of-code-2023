package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseFile(fileName string) map[string][]int {
	file, fErr := os.Open(fileName)

	if fErr != nil {
		log.Fatal(fErr)
	}

	defer func(f *os.File) {
		closeErr := f.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	parsed := make(map[string][]int)
	for scanner.Scan() {
		line := scanner.Text()
		colonSplit := strings.Split(line, ":")
		attrName := colonSplit[0]
		samplesStr := strings.Fields(strings.TrimSpace(colonSplit[1]))
		samples := make([]int, 0)
		for _, sampleStr := range samplesStr {
			if sample, convertErr := strconv.Atoi(sampleStr); convertErr != nil {
				log.Fatal(convertErr)
			} else {
				samples = append(samples, sample)
			}
		}
		parsed[attrName] = samples
	}
	return parsed
}

func MapTimeDistance(parsed map[string][]int) map[int]int {
	mapped := make(map[int]int)
	for i, time := range parsed["Time"] {
		mapped[time] = parsed["Distance"][i]
	}
	return mapped
}

func FindDistance(chargingTime, totalTime int) int {
	return (totalTime - chargingTime) * chargingTime
}

func FindNumWaysToBeat(totalTime, distanceRecord int) int {
	numWaysToBeat := 0
	midPoint := totalTime / 2
	for i := midPoint; i < totalTime; i++ {
		distance := FindDistance(i, totalTime)
		if distance > distanceRecord {
			numWaysToBeat++
		} else {
			break
		}
	}

	for i := midPoint - 1; i > 0; i-- {
		distance := FindDistance(i, totalTime)
		if distance > distanceRecord {
			numWaysToBeat++
		} else {
			break
		}
	}
	return numWaysToBeat
}

func ParseFile2(fileName string) (totalTime, distanceRecord int) {
	file, fErr := os.Open(fileName)

	if fErr != nil {
		log.Fatal(fErr)
	}

	defer func(f *os.File) {
		closeErr := f.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	parsed := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()

		colonSplit := strings.Split(line, ":")
		attrName := colonSplit[0]
		attrValueStr := ""
		samplesStr := strings.Fields(strings.TrimSpace(colonSplit[1]))
		for _, sampleStr := range samplesStr {
			attrValueStr = fmt.Sprintf("%s%s", attrValueStr, sampleStr)
		}
		attrValue, convertErr := strconv.Atoi(attrValueStr)
		if convertErr != nil {
			log.Fatal(convertErr)
		}
		parsed[attrName] = attrValue
	}
	return parsed["Time"], parsed["Distance"]
}
