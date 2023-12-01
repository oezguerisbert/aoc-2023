package main

import (
	"aoc-2023/learning-go/day1"
	"fmt"
	"os"
)


func main() {
  // get the file name from the command line
  // if no file name is given, print and error and exit
  if len(os.Args) < 2 {
    fmt.Println("Please provide a file name")
    os.Exit(1)
  }
  filename := os.Args[1]
  // get the file content as a list of strings, one per line
  file, err := os.Open(filename)
  file2, err2 := os.Open(filename)
  if err != nil {
    fmt.Println("Error reading file")
    os.Exit(1)
  }
  if err2 != nil {
    fmt.Println("Error reading file")
    os.Exit(1)
  }
  day1.Part1(file)
  day1.Part2(file2)
}


