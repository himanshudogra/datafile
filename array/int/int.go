// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetInt reads an int number from the file
func GetInt(fileName string) ([5]int, error) {
	var numbers [5]int
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			return numbers, err
		}
		i++
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
