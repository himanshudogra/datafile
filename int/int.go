// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// GetInt reads an int number from the file
func GetInt(fileName string) ([5]int, error) {
	var numbers [5]int
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		i++
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil

}
