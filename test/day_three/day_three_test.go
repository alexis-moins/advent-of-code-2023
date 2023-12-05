package daythree_test

import (
	"testing"

	daythree "github.com/alexis-moins/advent-of-code-2023/internal/day_three"
	utils "github.com/alexis-moins/advent-of-code-2023/test/utils"
)

func TestGetNumberInLine(t *testing.T) {
	file, scanner := utils.GetInput("day_three")
	defer file.Close()

	engineSchematic := daythree.ParseEngineSchematic(scanner)
	engineParts := engineSchematic.GetEngineParts()

	got := len(engineParts)
	want := 9

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

    got = 0
    for _, enginePart := range engineParts {
        got += enginePart
    }

    want = 4363

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
