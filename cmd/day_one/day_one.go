package main

import (
	dayone "github.com/alexis-moins/advent-of-code-2023/internal/day_one"
	"github.com/alexis-moins/advent-of-code-2023/internal/utils"
)

func main() {
	file, scanner := utils.GetInput("day_one")
	defer file.Close()

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += dayone.GetNumberInLine(line)
	}

	println(sum)
}
