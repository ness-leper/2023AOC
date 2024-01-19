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

func SolveP1(file []string) {
	// Need to read in each line into a character array
	var engine [][]string
	for i := 0; i < len(file); i++ {
		var output []string
		for li := 0; li < len(file[i]); li++ {
			output = append(output, string(file[i][li]))
		}
		engine = append(engine, output) }

  var foundNumbers []int
	for i := 0; i < len(engine); i++ {
		var makeNumber string

    // All this was done. Getting rid of it.
    // Hate this. Need to come up with a better method.
    posStart := struct {
      Found bool
      Position int
    }{
      Found: false,
      Position: 0,
    }
    posEnd := 0
		for ni := 0; ni < len(engine[i]); ni++ {
			re := regexp.MustCompile(`^\d+$`)
			if re.MatchString(engine[i][ni]) {
        if posStart.Found == false {
          posStart.Found = true
          posStart.Position = ni
        }
        posEnd = ni
				makeNumber = makeNumber + engine[i][ni]
			}

      if !re.MatchString(engine[i][ni]) {
				// TODO:
				// Need to validate if the found number is good to add.
				if len(makeNumber) > 0 {
					// Check the cell to the left
          found := false
          if posStart.Position - 1 >= 0{
            if engine[i][posStart.Position - 1] != "." {
              found = true
            }
          }
					// Check the cell to the right
          if posEnd + 1 <= len(engine[i]) {
            if engine[i][posEnd + 1] != "." {
              found = true
            }
          }
					// Check the row-1 cell[x-1...y+1]
          if i - 1 >= 0  && posStart.Position - 1 >= 0 && posEnd + 1 <= len(engine[i])-1 {
            for ch := posStart.Position - 1; ch < posEnd+1; ch++ {
              if engine[i-1][ch] != "." {
                found = true
              }
            }
          }
					// Check the row+1 cell[x-1...y+1]
          fmt.Println(makeNumber, posStart.Position, posEnd)
          checkPos := 0
          if posStart.Position - 1 >= 0 {
            checkPos = posStart.Position
          }
          if i + 1 <= len(engine)-1  && posEnd + 1 <= len(engine[i])-1 {
            fmt.Println(engine[i+1])
            for ch := checkPos; ch < posEnd+1; ch++ {
              fmt.Println(ch,found, engine[i+1][ch])
              if engine[i+1][ch] != "." {
                found = true
              }
            }
          }
          fmt.Println(found)

          if found {
            converted, err := strconv.Atoi(makeNumber)
            if err == nil {
              foundNumbers = append(foundNumbers, converted)
            }
          }
					makeNumber = ""
          posStart.Found = false
          posStart.Position = 0
          posEnd = 0
				} // end of appending the Make Number
			} // End of the Else
		} // End of the For loop

	} // end of identification loop

	// Stored in an array of arrays
	// Find each number in the array (0-9)
	// Check if any of the adjacent cells (master array +/- 1) same array beginning-1/end+1) have a symbol (not .)
  sum := 0
  fmt.Println(foundNumbers)
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
