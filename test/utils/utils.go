package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Return a file and a scanner created from the file.
// The file is returned so that the user can manually close it.
func GetInput(name string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fmt.Sprintf("../data/%s.txt", name))

	if err != nil {
		log.Fatal(err)
	}

	return file, bufio.NewScanner(file)
}
