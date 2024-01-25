package main

import (
	"fmt"
  "readfile"
)


func main() {
	file, err := readfile.ReadFile("puzzletest.txt")
	if err != nil {
		fmt.Println(err)
	}


  fmt.Println(file)
}
