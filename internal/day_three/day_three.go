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

// func (self EngineSchematic) getAdjacentSlices(x, y int) []string {
// 	slices := []string{}

// 	startIndex := max(x-1, 0)
// 	endIndex := min(xEnd+1, len(self.grid[y])-1)

// 	if xStart != startIndex {
// 		leftChar := self.grid[y][startIndex]
// 		slices = append(slices, string(leftChar))
// 	}

// 	if y > 0 {
// 		// Not on the top row
// 		slices = append(slices, self.grid[y-1][startIndex:endIndex])
// 	}

// 	rightChar := self.grid[y][xEnd]
// 	slices = append(slices, string(rightChar))

// 	if y < len(self.grid)-1 {
// 		// Not on bottom row
// 		slices = append(slices, self.grid[y+1][startIndex:endIndex])
// 	}

// 	return slices
// }

// var re = regexp.MustCompile(`[\.\d]`)

// func (self EngineSchematic) IsEnginePart(x, y int) bool {
// 	slices := self.getAdjacentSlices(self.index, x, y)
// 	fmt.Printf("slices: %v\n", slices)

// 	for _, slice := range slices {
// 		if specialChars := re.ReplaceAll([]byte(slice), []byte("")); len(specialChars) > 0 {
// 			fmt.Printf("specialCharacters: %s\n", string(specialChars))
// 			return true
// 		}
// 	}

// 	return false
// }

// // mustParseDigit returns true if the engine should parse the digits that were found.
// func (self EngineSchematic) mustParseDigit(char rune, x int, line string) bool {
// 	return self.isParsingNumber() && (!unicode.IsDigit(char) || self.isLastCharacterOfLine(x, line))
// }

// func (self EngineSchematic) isLastCharacterOfLine(x int, line string) bool {
// 	return x == len(line)-1
// }

// func (self EngineSchematic) isParsingNumber() bool {
// 	return self.index != -1
// }

// func (self EngineSchematic) GetEngineParts() []int {
// 	partNumbers := []int{}

// 	for y, line := range self.grid {

// 		for x, char := range line {
// 			fmt.Printf("char: %s\n", string(char))
// 			if unicode.IsDigit(char) && !self.isParsingNumber() {
// 				self.index = x
// 			} else if self.mustParseDigit(char, x, line) {
// 				if self.IsEnginePart(x, y) {
// 					var toConvert string

// 					if unicode.IsDigit(char) && self.isLastCharacterOfLine(x, line) {
// 						toConvert = self.grid[y][self.index:]
// 					} else {
// 						toConvert = self.grid[y][self.index:x]
// 					}
// 					fmt.Printf("number: %v\n", toConvert)

// 					number, err := strconv.Atoi(toConvert)

// 					if err != nil {
// 						log.Fatal(err)
// 					}

// 					partNumbers = append(partNumbers, number)
// 				}

// 				self.index = -1
// 			}
// 		}
// 	}

// 	return partNumbers
// }
