package daytwo

import (
	"log"
	"strconv"
	"strings"
)

var maxNumberOfCubes Round = Round{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// Round represents the number of cubes that were taken out of the bag
// grouped by the color of the cubes (red, green or blue).
type Round = map[string]int

type Game struct {
	Id     int
	rounds []Round
}

// IsValid
func (self Game) IsValid() bool {
	for _, round := range self.rounds {
		for color, numberOfCubes := range round {
			maximum, ok := maxNumberOfCubes[color]

			if !ok {
				log.Fatal("color not found")
			}

			if numberOfCubes > maximum {
				return false
			}
		}
	}
	return true
}

// parseRound returns a Round which is a map of colors (red, green or blue) and the number of.
func parseRound(round string) Round {
	subsets := strings.Split(strings.TrimSpace(round), ",")

	cubes := Round{}

	for _, subset := range subsets {
		info := strings.Split(strings.TrimSpace(subset), " ")

		if len(info) != 2 {
			log.Fatal("too many fields in info")
		}

		number, err := strconv.Atoi(info[0])

		if err != nil {
			log.Fatal(err)
		}

		if _, ok := cubes[info[1]]; !ok {
			cubes[info[1]] = 0
		}

		cubes[info[1]] += number
	}

	return cubes
}

func ParseGame(line string) Game {
	splitLine := strings.Split(line, ":")

	if len(splitLine) != 2 {
		log.Fatal("too many fields in game")
	}

	toConvert := strings.Replace(splitLine[0], "Game ", "", 1)
	id, err := strconv.Atoi(toConvert)

	if err != nil {
		log.Fatal(err)
	}

	rounds := []Round{}

	for _, round := range strings.Split(splitLine[1], ";") {
		rounds = append(rounds, parseRound(round))
	}

	return Game{
		Id:     id,
		rounds: rounds,
	}
}
