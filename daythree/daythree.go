package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	for i := 0; i < len(engine); i++ {
		var makeNumber string
		var foundNumbers []int
		for ni := 0; ni < len(engine[i]); ni++ {
			re := regexp.MustCompile(`^\d+$`)
			if re.MatchString(engine[i][ni]) {
				makeNumber = makeNumber + engine[i][ni]
			} else {
        // TODO:
        // Need to validate if the found number is good to add.
				fmt.Println(makeNumber)
        makeNumber = ""
			}
		}
		fmt.Println(engine[i])
		fmt.Println(foundNumbers)
	}

	// Stored in an array of arrays
	// Find each number in the array (0-9)
	// Check if any of the adjacent cells (master array +/- 1) same array beginning-1/end+1) have a symbol (not .)
}

func main() {
	file, err := ReadFile("puzzletest.txt")
	if err != nil {
		fmt.Println(err)
	}

	SolveP1(file)
	// var sum int
	// for i := 0; i < len(file); i++ {
	// 	temp, err := SolveP2(file[i])
	// 	if err != nil {
	// 		fmt.Println(file[i], err)
	// 	}
	// 	sum = sum + temp
	// }

	// fmt.Println("Final: ", sum)

}
