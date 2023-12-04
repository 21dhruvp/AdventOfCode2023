package main

import (
	"bufio"
	"fmt"
	"strconv"
	"log"
	"os"
	"strings"
)

func GetInput() []string {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func main() {
    sum := 0
    input := GetInput();
    for _, s := range input {
        gameNum := s[strings.Index(s, " ") + 1 : strings.Index(s, ":")]
        rounds := strings.Split(s[strings.Index(s, ":") + 2:], "; ")
        redMax, greenMax, blueMax := 0, 0, 0
        for _, round := range rounds {
            cubes := strings.Split(round, ", ")
            for _, cube := range cubes {
                splitCube := strings.Split(cube, " ")
                num, _ := strconv.Atoi(splitCube[0])
                color := splitCube[1]

                if color == "red" && num > redMax {
                    redMax = num
                } else if color == "green" && num > greenMax {
                    greenMax = num
                } else if color == "blue" && num > blueMax {
                    blueMax = num
                }
            }
        }
        fmt.Println("Game " + gameNum + " maxes: red: " + strconv.Itoa(redMax) + " green: " + strconv.Itoa(greenMax) + " blue: " + strconv.Itoa(greenMax))
        sum += redMax * greenMax * blueMax
    }
        fmt.Println("********************Solution********************")
        fmt.Println(sum)
}
