package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/alexis-moins/advent-of-code-2023/internal/utils"
)

func main() {
	file, scanner := utils.GetInput("day-one")
	defer file.Close()

	sum := 0

	for scanner.Scan() {
		digits := []rune{}

		for _, char := range scanner.Text() {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}

		format := fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1])

		number, err := strconv.Atoi(format)

		if err != nil {
			log.Fatal(err)
		}

		sum += number
	}

	println(sum)
}
