// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetInt function reads an int from each line of a file.
func GetInt(fileName string) ([]int, error) {
	var numbers []int

	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
