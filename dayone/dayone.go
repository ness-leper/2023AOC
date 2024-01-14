package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(file string) ([]string, error) {
	f, err := os.Open(file)
	var errReturn error
	if err != nil {
		errReturn = err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		errReturn = err
	}

	f.Close()

	return lines, errReturn
}

func CalibrateP1(value string) (int, error) {
	var first string
	var last string

	for i := 0; i < len(value); i++ {
		temp := string(value[i])

		if _, err := strconv.Atoi(temp); err == nil {
			if first == "" {
				first = temp
			}
			last = temp
		}
	}

	val, err := strconv.Atoi(first + last)
	return val, err
}

func main() {
	file, err := ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}

  var sum int
	for i := 0; i < len(file); i++ {
		calibrated, err := CalibrateP1(file[i])
		if err != nil {
			fmt.Println(err)
		}
    sum = sum + calibrated
	}

  fmt.Println(sum)
}
