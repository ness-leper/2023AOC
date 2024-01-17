package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
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

func main(){
	file, err := ReadFile("puzzletest.txt")
	if err != nil {
		fmt.Println(err)
	}

  // Need to read in each line into a character array
  // Stored in an array of arrays
  // Find each number in the array (0-9)
  // Check if any of the adjacent cells (master array +/- 1) same array beginning-1/end+1) have a symbol (not .)


  for i := 0; i < len(file); i++ {
    fmt.Println(file[i])
  }
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
