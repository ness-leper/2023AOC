package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func CheckCoordinate(engine [][]string, row int, column int) bool {
	output := false

	colM1 := engine[row][column]
  fmt.Println(colM1)
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(colM1) && colM1 != "." {
		output = true
	}

	return output
}

func P1Adjacent(engine [][]string, row int, column int) bool {
	output := false

	// column - 1
	if column-1 >= 0 && !output {
    output = CheckCoordinate(engine, row, column - 1)
	}
	// column + 1
	if column+1 <= len(engine[row][column])-1 && !output {
    output = CheckCoordinate(engine, row, column + 1)
	}

	// row - 1 [ column - 1 ]
	if row - 1 >= 0 && !output {
    output = CheckCoordinate(engine, row - 1, column)
	}

	// row + 1 [ column - 1 ]
	if row + 1 <= len(engine)-1 && !output {
    output = CheckCoordinate(engine, row + 1, column)
	}

	// row - 1 [ column - 1 ]
	if row-1 >= 0 && column - 1 >= 0 && !output {
    output = CheckCoordinate(engine, row - 1, column - 1)
	}

  // row -1 [ column + 1]
  if row -1 >=0 && (column + 1) < len(engine[row][column]) && !output {
    output = CheckCoordinate(engine, row -1, column + 1)
  }

	// row + 1 [ column - 1 ]
	if row+1 < len(engine) && (column - 1) >= 0 && !output {
    output = CheckCoordinate(engine, row + 1, column - 1)
	}

  // row +1 [ column + 1]
  if row + 1 < len(engine) && (column + 1) < len(engine[row][column]) && !output {
    output = CheckCoordinate(engine, row + 1, column + 1)
  }

	return output
}

func SolveP1(file []string) {
	// Need to read in each line into a character array
	var engine [][]string
	for i := 0; i < len(file); i++ {
		var output []string
		for li := 0; li < len(file[i]); li++ {
			output = append(output, string(file[i][li]))
		}
		engine = append(engine, output)
	}

	var foundNumbers []int
	// for i := 0; i < len(engine); i++ {
	for i := 0; i < len(file); i++ {
		var makeNumber string
		enginePart := false
		for ni := 0; ni < len(engine[i]); ni++ {
			re := regexp.MustCompile(`^\d+$`)
			if re.MatchString(engine[i][ni]) {
				makeNumber = makeNumber + engine[i][ni]
				if !enginePart {
					enginePart = P1Adjacent(engine, i, ni)
				}
			}
			if engine[i][ni] == "." && len(makeNumber) > 0 {
				fmt.Println(makeNumber, enginePart)
				if enginePart {
					converted, err := strconv.Atoi(makeNumber)
					if err == nil {
						foundNumbers = append(foundNumbers, converted)
					}
				}
				enginePart = false
				makeNumber = ""
			}
		}
	} // end of identification loop

	// Stored in an array of arrays
	// Find each number in the array (0-9)
	// Check if any of the adjacent cells (master array +/- 1) same array beginning-1/end+1) have a symbol (not .)
	sum := 0
	for i := 0; i < len(foundNumbers); i++ {
		sum = sum + foundNumbers[i]
	}

	fmt.Println(sum)
}

func main() {
	file, err := ReadFile("puzzletest.txt")
	if err != nil {
		fmt.Println(err)
	}

	SolveP1(file)
}
