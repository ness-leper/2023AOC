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

func CalibrateP2(value string) (int, error) {
	var first string
	var last string

	words := [10]string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}

	for i := 0; i < len(value); i++ {
		temp := string(value[i])

		if _, err := strconv.Atoi(temp); err == nil {
			if first == "" {
				first = temp
			}
			last = temp
		} else {
			// See if the current letter is the beginning of a word number
			// If yes, is the rest of the word there?
			// If yes, do the first and last setting
			// fmt.Println("Not a number", temp)
			for li := 0; li < len(words); li++ {
				currL := words[li]
				if string(currL[0]) != temp {
					continue
				}
				if i+len(words[li]) > len(value) {
					continue
				}
				var newCheck string
				for ch := 0; ch < len(words[li]); ch++ {
					newCheck = newCheck + string(value[i+ch])
				}
				if newCheck == string(words[li]) {
					if first == "" {
						first = strconv.Itoa(li + 1)
					}
					last = strconv.Itoa(li + 1)
				}
			}
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
		calibrated, err := CalibrateP2(file[i])
		if err != nil {
			fmt.Println(file[i], err)
		}
		sum = sum + calibrated
	}

	fmt.Println(sum)
}
