package day2

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
  file, err := os.Open("inputs/2/part1-prod.txt")
  if err != nil {
    fmt.Println("Error reading file")
    os.Exit(1)
  }
  fmt.Println("---------------")
  fmt.Println("Day 2 - Part 2:")
	fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)
  var games []Game
  for fileScanner.Scan() {
    game := fileScanner.Text()
    aGame := CreateGame(game)
    games = append(games, aGame)
  }
	gameCollection := GameCollection{games}
	fmt.Println("Sum of power of reveals:", gameCollection.SumOfPowerOfReveals())
}