package main

import (
	"bufio"
	"fmt"
        "strconv"
        "strings"
	"log"
	"os"
	"unicode"
)

func first(n int, _ error) int {
    return n
}

func main() {
    file, err := os.Open("input.txt")
    if(err != nil) {
        log.Fatal(err)
    }

    read := bufio.NewReader(file)

    sum := 0
    currDigit := 0
    var digits [9]string = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    line, err := read.ReadString('\n')

    for err == nil {
        earliestInd, latestInd := len(line), 0
        earliest, latest := 0, 0
        for i:=0; i < 9; i+=1 {
            early := strings.Index(line, digits[i])   
            last := strings.LastIndex(line, digits[i])      
            if(early < earliestInd && early != -1) {
                earliestInd = early
                earliest = i+1
            }
            if (last > latestInd && last != -1) {
                latestInd = last
                latest = i+1
            }
            //fmt.Println("-------------------------------------")
            //fmt.Println("earliest: " + strconv.Itoa(earliest) + ", earliestInd: " + strconv.Itoa(earliestInd) + ", latest: " + strconv.Itoa(latest) + ", latestInd: " + strconv.Itoa(latestInd))
        }

        for i:=0; i < len(line[:earliestInd]); i+=1 {
            if unicode.IsDigit(rune(line[i])) {
                currDigit += int(first(strconv.Atoi(string(line[i]))))
                break;
            }
        }
        if(currDigit == 0) {
            currDigit += earliest
        }
        currDigit *= 10
        foundLater := false
        var end int
        if (latest == 0) {
            if earliest == 0 {
                end = 0
            } else {
                end = earliestInd + len(digits[earliest-1])
            }
        } else {
            end = latestInd + len(digits[latest-1])-1
        }
        fmt.Print("end: ")
        fmt.Println(end)
        for i:= len(line)-4; i >= end; i-=1 {
            fmt.Println("current: " + string(line[i]))
            if unicode.IsDigit(rune(line[i])) {
                currDigit += int(first(strconv.Atoi(string(line[i]))))
                foundLater = true
                break;
            }
        }
        if(!foundLater) {
            currDigit += latest
        } /*else if (latest == 0) {
            currDigit += (currDigit/10)
        }*/
        fmt.Println(currDigit)
        sum += currDigit
        currDigit = 0
        line, err = read.ReadString('\n')
    }
    fmt.Println(sum)
}
