package lib

import (
	"log"
	"strconv"
	"strings"
)

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	GameId   int
	CubeSets []CubeSet
}

func IsGameValid(game Game, referenceCubeSet CubeSet) bool {
	for _, cubeSet := range game.CubeSets {
		if !(cubeSet.Red <= referenceCubeSet.Red && cubeSet.Blue <= referenceCubeSet.Blue && cubeSet.Green <= referenceCubeSet.Green) {
			return false
		}
	}
	return true
}

func ParseGameFromLine(line string) Game {
	splitByColon := strings.Split(line, ":")
	gameId, parseErr := strconv.Atoi(strings.Split(splitByColon[0], " ")[1])
	if parseErr != nil {
		log.Fatal(parseErr)
	}
	var cubeSets []CubeSet
	cubeSetsInfo := strings.Split(splitByColon[1], ";")
	cubeColors := [3]string{"Red", "Green", "Blue"}
	for _, cubeSetInfo := range cubeSetsInfo {
		cubeSetMap := make(map[string]int)
		for _, cube := range strings.Split(cubeSetInfo, ",") {
			for _, cubeColor := range cubeColors {
				if strings.Contains(cube, strings.ToLower(cubeColor)) {
					count, err := strconv.Atoi(strings.Split(cube, " ")[1])
					if err != nil {
						log.Fatal(err)
					}
					cubeSetMap[cubeColor] = count
				}
			}

		}
		cubeSet := CubeSet{
			Red:   cubeSetMap["Red"],
			Blue:  cubeSetMap["Blue"],
			Green: cubeSetMap["Green"],
		}
		cubeSets = append(cubeSets, cubeSet)
	}
	game := Game{
		GameId:   gameId,
		CubeSets: cubeSets,
	}
	return game
}

func FindFewestCubesPossible(game Game) CubeSet {
	fewestPossibleCubeSet := CubeSet{0, 0, 0}
	for _, cubeSet := range game.CubeSets {
		if fewestPossibleCubeSet.Red < cubeSet.Red {
			fewestPossibleCubeSet.Red = cubeSet.Red
		}
		if fewestPossibleCubeSet.Green < cubeSet.Green {
			fewestPossibleCubeSet.Green = cubeSet.Green
		}
		if fewestPossibleCubeSet.Blue < cubeSet.Blue {
			fewestPossibleCubeSet.Blue = cubeSet.Blue
		}
	}
	return fewestPossibleCubeSet
}

func FindCubeSetPower(cubeSet CubeSet) int {
	return cubeSet.Red * cubeSet.Green * cubeSet.Blue
}
