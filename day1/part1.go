package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1(file *os.File) {
  fileScanner := bufio.NewScanner(file)
  sum := 0
  fileScanner.Split(bufio.ScanLines)
  allowed := []rune{'0','1','2','3','4','5','6','7','8','9'}
  theCombNumber := ""
  for fileScanner.Scan() {
    for _, c := range fileScanner.Text() {
      // is c a digit?
      for _, d := range allowed {
        if c == d {
          theCombNumber += string(c)
        }
      }
    }
    firtsDigit := theCombNumber[0]
    lastDigit := theCombNumber[len(theCombNumber)-1]
    // convert the string to an int
    numberThing := fmt.Sprintf("%s%s", string(firtsDigit), string(lastDigit))
    i, err := strconv.Atoi(numberThing)
    if err != nil {
      fmt.Println("Error converting string to int")
      os.Exit(1)
    }
    sum += i
    theCombNumber = ""
  }
  fmt.Println("Sum", sum)
}