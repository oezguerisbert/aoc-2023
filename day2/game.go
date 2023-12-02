package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Color string
const (
  Red Color = "red"
  Blue Color = "blue"
  Green Color = "green"
)

type Cube struct {
  color Color
  amount int
}

type  Reveal struct {
  cubes []Cube
}

type Game struct {
  gameId string
  reveals []Reveal
}

type GameCollection struct {
	games []Game
}

func (g Game) FindGameIdForReveals(
  Cubes []Cube,
) []string {
  var gameIds map[string]bool = make(map[string]bool)

  loop:
  for _, reveal := range g.reveals {
    for _, cube := range Cubes {
      for _, revealCube := range reveal.cubes {
        if cube.color == revealCube.color && revealCube.amount > cube.amount {
          // fmt.Println("Found a cube that is too big, skipping", cube, revealCube)
          gameIds[g.gameId] = false
          continue loop
        }
        if cube.color == revealCube.color && revealCube.amount <= cube.amount {
          if _, found := gameIds[g.gameId]; !found {
            gameIds[g.gameId] = true
          }
        }
      }
    }
  }

  var gameIdsStringArray []string
  for gameId, found := range gameIds {
    if found {
      gameIdsStringArray = append(gameIdsStringArray, gameId)
    }
  }

  return gameIdsStringArray
}

func (gc GameCollection) GetPowerOfReveals() map[string]int {
	var gamePower map[string]int = make(map[string]int)
	for _, game := range gc.games {
		reveals := game.GetFewestReveal()
		for _, reveal := range reveals.cubes {
			if _, found := gamePower[game.gameId]; !found {
				gamePower[game.gameId] = 1
			}
			gamePower[game.gameId] *= reveal.amount
		}
	}
	return gamePower
}

func (gc GameCollection) SumOfPowerOfReveals() int {
	gamePower := gc.GetPowerOfReveals()
	sum := 0
	for _, power := range gamePower {
		sum += power
	}
	return sum
}


func (g Game) GetFewestReveal() Reveal {
	var fewestReveal Reveal
	redCube := Cube{Red, 0}
	blueCube := Cube{Blue, 0}
	greenCube := Cube{Green, 0}

	for _, reveal := range g.reveals {
		for _, cube := range reveal.cubes {
			if cube.color == Red && redCube.amount < cube.amount {
				redCube.amount = cube.amount
			}
			if cube.color == Blue && blueCube.amount < cube.amount {
				blueCube.amount = cube.amount
			}
			if cube.color == Green && greenCube.amount < cube.amount {
				greenCube.amount = cube.amount
			}
		}
	}
	fewestReveal.cubes = append(fewestReveal.cubes, redCube)
	fewestReveal.cubes = append(fewestReveal.cubes, blueCube)
	fewestReveal.cubes = append(fewestReveal.cubes, greenCube)

	return fewestReveal
}

func CreateGame(line string) Game {
  gameIdString := strings.Split(line, ": ")[0]
	gamereveals := strings.Split(line, ": ")[1]
	gameId := strings.Split(gameIdString, " ")[1]
	gameReveals := strings.Split(gamereveals, "; ")
	var reveals []Reveal
	for _, reveal := range gameReveals {
		// fmt.Println(reveal)
		gameCubes := strings.Split(reveal, ", ")
		var cubes []Cube
		for _, cube := range gameCubes {
			amount := strings.Split(cube, " ")[0]
			color := strings.Split(cube, " ")[1]
			colorColor := Color(color)
			amountInt, err := strconv.Atoi(amount)
			if err != nil {
				fmt.Println("Error converting string to int")
				os.Exit(1)
			}
			cubes = append(cubes, Cube{colorColor, amountInt})
			reveals = append(reveals, Reveal{
				cubes,
			})
		}
	}
    
  return Game {
		gameId, 
		reveals,
	}
}