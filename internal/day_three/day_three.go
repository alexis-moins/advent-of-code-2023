package daythree

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

// GEAR_CHAR represents a gear in the engine schematic.
const GEAR_CHAR rune = '*'

type Position struct {
	x, y int
}

// Range represents a range from start to end (included)
type Range struct {
	start, end, y int
}

func (self Range) IsAdjacentTo(position Position) bool {
	return self.y+1 >= position.y && position.y >= self.y-1 && (position.x == self.start-1 || position.x == self.end)
}

type EngineSchematic struct {
	grid           []string
	parts          map[Range]string
	PotentialGears []Position
}

func ParseEngineSchematic(scanner *bufio.Scanner) EngineSchematic {
	grid := []string{}

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	engineSchematic := EngineSchematic{
		grid:           grid,
		parts:          map[Range]string{},
		PotentialGears: []Position{},
	}

	engineSchematic.parseGrid()
	return engineSchematic
}

func (self *EngineSchematic) parseGrid() {
	for y, line := range self.grid {
		index := -1
		// TODO: need to parse number if it is the end of the line

		for x, char := range line {
			fmt.Printf("char: %s\n", string(char))

			if unicode.IsDigit(char) {
				if index == -1 {
					index = x
				}

				if x == len(line)-1 {
					value := line[index:]
					self.parts[Range{start: index, end: x, y: y}] = value

					index = -1
				}
			} else if char == GEAR_CHAR {
				self.PotentialGears = append(self.PotentialGears, Position{x, y})
                index = -1
			} else if index != -1 {
				value := line[index:x]
				self.parts[Range{start: index, end: x - 1, y: y}] = value

				index = -1
			}
		}
	}
}

func (self EngineSchematic) FindAdjacentEngineParts(position Position) []int {
	parts := []int{}

	for _range, value := range self.parts {
        fmt.Printf("value: %v\n", value)
		if _range.IsAdjacentTo(position) {
			part, err := strconv.Atoi(value)

			if err != nil {
				log.Fatalln(err)
			}

			parts = append(parts, part)
		}
	}

	return parts
}
