package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func Part1() {

  file, err := os.Open("inputs/2/part1-prod.txt")
  if err != nil {
    fmt.Println("Error reading file")
    os.Exit(1)
  }
  fmt.Println("---------------")
  fmt.Println("Day 2 - Part 1:")
  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)
  var games []Game
  for fileScanner.Scan() {
    game := fileScanner.Text()
    aGame := CreateGame(game)
    games = append(games, aGame)
  }
  var foundGameIds []string
  for _, game := range games {
    foundGameIds = append(foundGameIds, game.FindGameIdForReveals([]Cube{
      Cube{Red, 12},
      Cube{Green, 13},
      Cube{Blue, 14},
    })...)
  }
  sum := 0
  for _, foundGameId := range foundGameIds {
    gameIdInt, err := strconv.Atoi(foundGameId)
    if err != nil {
      fmt.Println("Error converting string to int")
      os.Exit(1)
    }
    sum += gameIdInt
  }
  fmt.Println("Sum of gameIds:", sum)
}