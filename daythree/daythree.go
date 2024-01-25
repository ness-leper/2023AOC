package main

import (
	"fmt"
  "readfile"
  "findnumbers"
)


func main() {
	file, err := readfile.ReadFile("puzzletest.txt")
	if err != nil {
		fmt.Println(err)
	}

  // find all the numbers in the engine
  findnumbers.FindNumbers(file)
}
