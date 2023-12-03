package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	fmt.Println("---------------")
	fmt.Println("Day 3 - Part 2:")
	f := "part1-prod"
	// f = "test"
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
	for _, s := range game.GetAllSymbols() {
		if s.value != "*" {
			continue
		}
		if len(s.numbers) == 2 {
			sum += s.numbers[0].value * s.numbers[1].value
		}
	}
	fmt.Println("  Sum: ", sum)
}
