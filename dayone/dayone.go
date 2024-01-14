package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(file string) ([]string, error) {
  f, err := os.Open(file)
  var errReturn error
  if err != nil {
    errReturn = err
  }

  defer f.Close()

  scanner := bufio.NewScanner(f)

  var lines[] string
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
  file, err := ReadFile("puzzle.txt")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(file)
}
