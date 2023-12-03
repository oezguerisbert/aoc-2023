package day3

import (
	"bufio"
	"fmt"
	"os"
)

func containsRange(positions []Position, position []Position) bool {
	for _, p := range positions {
		if p.x == position[0].x && p.y == position[0].y {
			return true
		}
	}
	return false
}

func Part1() {
	fmt.Println("---------------")
	fmt.Println("Day 3 - Part 1:")
	f := "part1-prod"
	// f = "test2"
	file, err := os.Open(fmt.Sprintf("inputs/3/%s.txt", f))
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	file.Close()
	game := ParseGame(lines)
	sum := 0
	visitedNumbers := map[int][]Position{}
	for _, s := range game.GetAllSymbols() {
		numbers := s.numbers
		for _, n := range numbers {
			nV := n.value
			if _, ok := visitedNumbers[nV]; !ok {
				if visitedNumbers[nV] == nil {
					visitedNumbers[nV] = []Position{}
				}
			}
			if !containsRange(visitedNumbers[nV], []Position{n.start_position, n.end_position}) {
				visitedNumbers[nV] = append(visitedNumbers[nV], n.start_position, n.end_position)
				sum += nV
			}
		}
	}
	fmt.Println("  Sum: ", sum)

}
