package day3

import (
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	x int
	y int
}

type Result struct {
	value          int
	start_position Position
	end_position   Position
}

type Symbol struct {
	value   string
	numbers []Result
}

type Node struct {
	value    string
	position Position
}

type Game struct {
	matrix [][]Node
}

func ParseGame(lines []string) Game {
	var matrix [][]Node
	for x, line := range lines {
		var row []Node
		for y, symbol := range line {
			row = append(row, Node{string(symbol), Position{x, y}})
		}
		matrix = append(matrix, row)
	}
	return Game{matrix}
}

func (g Game) GetNode(position Position) (Node, error) {
	if position.x < 0 || position.y < 0 {
		return Node{}, fmt.Errorf("Position out of bounds")
	}
	if position.x >= len(g.matrix) || position.y >= len(g.matrix[position.x]) {
		return Node{}, fmt.Errorf("Position out of bounds")
	}
	return g.matrix[position.x][position.y], nil
}

func conainsResult(results []Result, result Result) bool {
	for _, r := range results {
		if r.value == result.value && r.start_position.x == result.start_position.x && r.start_position.y == result.start_position.y && r.end_position.x == result.end_position.x && r.end_position.y == result.end_position.y {
			return true
		}
	}
	return false
}

func (g Game) GetAllSymbols() []Symbol {
	var symbols []Symbol
	for _, row := range g.matrix {
		for _, node := range row {
			isNumber := node.value >= "0" && node.value <= "9"
			if !isNumber && node.value != "." {
				nodes := g.GetNodesNextToSymbol(node)
				numbers := []Result{}
				for _, node := range nodes {
					n, start, end := g.GetFullNumberInRow(node)
					if !conainsResult(numbers, Result{n, start, end}) {
						numbers = append(numbers, Result{n, start, end})
					}
				}
				symbols = append(symbols, Symbol{
					node.value,
					numbers})
			}
		}
	}
	return symbols
}

func (g Game) GetNodesNextToSymbol(symbol Node) []Node {
	var numbers []Node
	positions := []Position{
		{symbol.position.x - 1, symbol.position.y},     // up
		{symbol.position.x - 1, symbol.position.y - 1}, // up left
		{symbol.position.x - 1, symbol.position.y + 1}, // up right
		{symbol.position.x + 1, symbol.position.y},     // down
		{symbol.position.x + 1, symbol.position.y - 1}, // down left
		{symbol.position.x + 1, symbol.position.y + 1}, // down right
		{symbol.position.x, symbol.position.y - 1},     // left
		{symbol.position.x, symbol.position.y + 1},     // right
	}
	for _, position := range positions {
		node, err := g.GetNode(position)
		if err != nil {
			continue
		}
		if node.value >= "0" && node.value <= "9" {
			numbers = append(numbers, node)
		}
	}
	return numbers
}

func containsPosition(positions []Position, position Position) bool {
	for _, pos := range positions {
		if pos.x == position.x && pos.y == position.y {
			return true
		}
	}
	return false
}

func (g Game) GetFullNumberInRow(number Node) (int, Position, Position) {
	visitedPositions := []Position{}
	row := g.matrix[number.position.x]
	numbersString := []Node{}
	numbersStringLeft := []Node{}
	numbersStringRight := []Node{}
	for i := number.position.y - 1; i >= 0; i-- {
		node := row[i]
		isNumber := node.value >= "0" && node.value <= "9"
		if isNumber && !containsPosition(visitedPositions, node.position) {
			numbersStringLeft = append(numbersStringLeft, node)
			visitedPositions = append(visitedPositions, node.position)
		} else {
			break
		}
	}
	for i := number.position.y + 1; i < len(row); i++ {
		node := row[i]
		if node.value >= "0" && node.value <= "9" && !containsPosition(visitedPositions, node.position) {
			numbersStringRight = append(numbersStringRight, node)
			visitedPositions = append(visitedPositions, node.position)
		} else {
			break
		}
	}

	numbersString = append(numbersString, numbersStringLeft...)
	numbersString = append(numbersString, numbersStringRight...)
	numbersString = append(numbersString, number)
	numberString := ""

	for i := 0; i < len(numbersString); i++ {
		for j := i + 1; j < len(numbersString); j++ {
			if numbersString[i].position.y > numbersString[j].position.y {
				temp := numbersString[i]
				numbersString[i] = numbersString[j]
				numbersString[j] = temp
			}
		}
	}
	for _, node := range numbersString {
		numberString += node.value
	}
	nString, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println("Error converting string to int")
		os.Exit(1)
	}
	startPos := numbersString[0].position
	endPos := numbersString[len(numbersString)-1].position
	return nString, startPos, endPos
}
