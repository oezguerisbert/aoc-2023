package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func Part2() {
  fmt.Println("---------------")
  fmt.Println("Day 1 - Part 2:")
  file, err := os.Open("inputs/1/part1-prod.txt")
  if err != nil {
    fmt.Println("Error reading file")
    os.Exit(1)
  }
  digits := map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "zero": 0,
  }
  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)
  sum := 0
  for fileScanner.Scan() {
    line := fileScanner.Text()
    collection := [][]int{}
    // find the index of each digit in the line, by walking through the line
    for idx, char := range line {
      // convert the character to a string
      charString := fmt.Sprintf("%c", char)
      // check if the character is a digit
      if _, err := strconv.Atoi(charString); err == nil {
        // if it is a digit, convert it to an int
        number, err := strconv.Atoi(charString)
        if err != nil {
          fmt.Println("Error converting string to int")
          os.Exit(1)
        }
        // add the index and the number to the collection
        collection = append(collection, []int{idx, number})
      }
    }
    // same thing for the words
    for word, number := range digits {
      // find all th indexes of the word in the line
      indexes := []int{}
      for idx := range line {
        // check if the word is a prefix of the line
        // get the part of the line starting at the index `idx`
        p := line[idx:]
        // check if the word is a prefix of the part of the line
        if strings.HasPrefix(p, word) {
          // if it is, add the index to the list of indexes
          indexes = append(indexes, idx)
        }
      }
      // add the indexes and the number to the collection
      for _, idx := range indexes {
        // fmt.Println(line, idx, number)
        collection = append(collection, []int{idx, number})
      }
      
    }
    // sort by index
    for i := 0; i < len(collection); i++ {
      for j := i + 1; j < len(collection); j++ {
        // is the index of the first element greater than the index of the second element?
        if collection[i][0] > collection[j][0] {
          // if it is, swap the elements
          collection[i], collection[j] = collection[j], collection[i] // swap
        }
      }
    }
    // fmt.Println(line, collection)
    firstNumber := collection[0][1] // get the first number
    lastNumber := collection[len(collection) - 1][1] // get the last number
    // fmt.Println(firstNumber, lastNumber)
    newNumberString := fmt.Sprintf("%d%d", firstNumber, lastNumber) // concatenate the first and last number
    newNumber, err := strconv.Atoi(newNumberString) // convert the string to an int
    if err != nil {
      fmt.Println("Error converting string to int")
      os.Exit(1)
    }
    sum += newNumber // add the new number to the sum
  }
  fmt.Println("Sum", sum) // print the sum
}