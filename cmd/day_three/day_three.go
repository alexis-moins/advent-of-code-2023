package main

import (
	"fmt"

	daythree "github.com/alexis-moins/advent-of-code-2023/internal/day_three"
	"github.com/alexis-moins/advent-of-code-2023/internal/utils"
)

func main() {
	file, scanner := utils.GetInput("day_three")
	defer file.Close()

	schematic := daythree.ParseEngineSchematic(scanner)
	engineParts := schematic.GetEngineParts()

	sum := 0
	for _, enginePart := range engineParts {
		sum += enginePart
	}

	fmt.Printf("sum: %v\n", sum)
}
