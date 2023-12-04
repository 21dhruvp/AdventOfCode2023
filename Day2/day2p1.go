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
    MAXRED, MAXGREEN, MAXBLUE := 12, 13, 14
    input := GetInput();
    for _, s := range input {
        gameNum, _ := strconv.Atoi(s[strings.Index(s, " ") + 1 : strings.Index(s, ":")])
        rounds := strings.Split(s[strings.Index(s, ":") + 2:], "; ")
        validGame := true
        for _, round := range rounds {
            currRed, currGreen, currBlue := 0, 0, 0
            cubes := strings.Split(round, ", ")
            for _, cube := range cubes {
                splitCube := strings.Split(cube, " ")
                num, _ := strconv.Atoi(splitCube[0])
                color := splitCube[1]

                if color == "red" {
                    currRed += num
                } else if color == "green" {
                    currGreen += num
                } else if color == "blue" {
                    currBlue += num
                }
            }
            if currRed > MAXRED || currGreen > MAXGREEN || currBlue > MAXBLUE {
                validGame = false
            }
        }
        if validGame {
            sum += gameNum
        }
    }
        fmt.Println("********************Solution********************")
        fmt.Println(sum)
}
