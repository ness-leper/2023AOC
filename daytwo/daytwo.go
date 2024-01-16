package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func SolveP1(line string) (int, error) {
	game := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	valid := true
	sumGame := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	currGame := strings.Split(line, ":")
	actions := strings.Split(currGame[1], "; ")
  fmt.Println(actions)

  // ##################################
  // There's got to be a better method to break down this string...
	for i := 0; i < len(actions); i++ {
		trimmed := strings.TrimSpace(actions[i])
		pulls := strings.Split(trimmed, ", ")
		for pi := 0; pi < len(pulls); pi++ {
			cp := pulls[pi]
			curr := strings.Split(cp, " ")
			val, err := strconv.Atoi(curr[0])
			if err == nil {
				sumGame[curr[1]] = sumGame[curr[1]] + val
			}
		}
    // ###################
    // Is there better way than 3 blocks?
    if sumGame["red"] > game["red"] {
      valid = false
    }

    if sumGame["green"] > game["green"] {
      valid = false
    }

    if sumGame["blue"] > game["blue"] {
      valid = false
    }
    // ###################
    sumGame["red"] = 0
    sumGame["green"] = 0
    sumGame["blue"] = 0
	}
  // ##################################

  gameID := strings.Split(strings.TrimSpace(currGame[0]), " ")
  fmt.Println(gameID[1], sumGame, valid)
	if valid {
		if val, err := strconv.Atoi(gameID[1]); err == nil {
      return val, nil 
    } else {
      return 0, err
    }
	}

	return 0, nil
}

func main() {
	file, err := ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}
	var sum int
	for i := 0; i < len(file); i++ {
		temp, err := SolveP1(file[i])
		if err != nil {
			fmt.Println(file[i], err)
		}
    fmt.Println(temp)
		sum = sum + temp
	}

	fmt.Println("Final: ", sum)
}
