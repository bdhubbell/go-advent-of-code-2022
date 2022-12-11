package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var myPoints int

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		opponentChoice := values[0]
		myChoice := values[1]
		
		if opponentChoice == "A" {

			// +3 for choosing scissors & +0 for lose
			if myChoice == "X" {
				myPoints += 3
			}

			// +1 for choosing rock & +3 for draw
			if myChoice == "Y" {
				myPoints += 4
			}

			// +2 for choosing paper & +6 for win
			if myChoice == "Z" {
				myPoints += 8
			}

		} else if opponentChoice == "B" {

			// +1 for choosing rock & +0 for lose
			if myChoice == "X" {
				myPoints += 1
			}

			// +2 for choosing paper & +3 for draw
			if myChoice == "Y" {
				myPoints += 5
			}

			// +3 for choosing scissors & +6 for win
			if myChoice == "Z" {
				myPoints += 9
			}

		} else if opponentChoice == "C" {

			// +2 for choosing paper & +0 for lose
			if myChoice == "X" {
				myPoints += 2
			}

			// +3 for choosing scissors & +3 for draw
			if myChoice == "Y" {
				myPoints += 6
			}

			// +1 for choosing rock & +6 for win
			if myChoice == "Z" {
				myPoints += 7
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(myPoints)
}