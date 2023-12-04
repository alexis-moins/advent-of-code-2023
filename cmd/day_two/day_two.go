package main

import (
	daytwo "github.com/alexis-moins/advent-of-code-2023/internal/day_two"
	"github.com/alexis-moins/advent-of-code-2023/internal/utils"
)

func main() {
	file, scanner := utils.GetInput("day_two")
	defer file.Close()

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		game := daytwo.ParseGame(line)

		if game.IsValid() {
			sum += game.Id
		}
	}

	println(sum)
}
