package main

import (
	"fmt"

	daythree "github.com/alexis-moins/advent-of-code-2023/internal/day_three"
	"github.com/alexis-moins/advent-of-code-2023/internal/utils"
)

func main() {
	file, scanner := utils.GetInput("day_three")
	defer file.Close()

    engine := daythree.ParseEngineSchematic(scanner)

    sum := 0

    for _, position := range engine.PotentialGears {
        parts := engine.FindAdjacentEngineParts(position)

        fmt.Printf("parts: %v\n", parts)

        if len(parts) == 2 {
            sum += parts[0] * parts[1]
        }
    }

	fmt.Printf("sum: %v\n", sum)
}
