package main

import "github.com/alexis-moins/advent-of-code-2023/internal/utils"

func main() {
	file, scanner := utils.GetInput("day-one")
	defer file.Close()

	for scanner.Scan() {
		println(scanner.Text())
	}
}
