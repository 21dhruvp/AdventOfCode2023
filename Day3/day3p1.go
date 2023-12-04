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
    input := GetInput()

}
