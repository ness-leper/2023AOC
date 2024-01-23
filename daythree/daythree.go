package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Part struct {
  Number string
  StartAt int
  EndAt int
}

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

func LoadEngine(file []string) [][]string {
	var engine [][]string
	for i := 0; i < len(file); i++ {
		var output []string
		for li := 0; li < len(file[i]); li++ {
			output = append(output, string(file[i][li]))
		}
		engine = append(engine, output)
	}

  return engine
}

func LocateNumbers(engine []string)  {
  var foundNumbers []Part
  newNumber := true 
  for i := 0; i < len(engine); i++ {
    var found Part
    re := regexp.MustCompile(`^\d+$`)
    if re.MatchString(engine[i]) {
      found.Number = found.Number + engine[i]
      if found.StartAt == 0 {
        found.StartAt = i
      }
      found.EndAt = i
    }
    if newNumber {
      newNumber = false
    }
    if engine[i] == "." {
      newNumber = true
      foundNumbers = append(foundNumbers, found)
    }
  }

  fmt.Println(foundNumbers)
}

func SolveP1(file []string) {
	// Need to read in each line into a character array
  engine := LoadEngine(file)

  // re := regexp.MustCompile(`^\d+$`)
  // if re.MatchString(engine[i][ni]) {
  // }

	// var foundNumbers []Part

  for i := 0; i < len(engine); i++ {
    LocateNumbers(engine[i])
    // foundNumbers = append(foundNumbers, LocateNumbers(engine[i]))
  }
}

func main() {
	file, err := ReadFile("puzzletest.txt")
	if err != nil {
		fmt.Println(err)
	}

	SolveP1(file)
}
