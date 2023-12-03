package dayone

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

var stringToRune = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func GetNumberInLine(line string) int {
	digits := []rune{}
	letters := ""

	for _, char := range line {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
			letters = ""
		} else {
			letters += string(char)

			for key, value := range stringToRune {
				if strings.HasSuffix(letters, key) {
					digits = append(digits, value)
					letters = ""

					break
				}
			}
		}
	}

	format := string(digits[0]) + string(digits[len(digits)-1])
	number, err := strconv.Atoi(format)

	if err != nil {
		log.Fatal(err)
	}

	return number
}
