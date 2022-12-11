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

	// X = Rock (1)
	// Y = Paper (2)
	// Z = Siccors (3)

	// L = 0
	// W = 6
	// T = 3

	for scanner.Scan() {

		values := strings.Split(scanner.Text(), " ")
		opponentChoice := values[0]
		myChoice := values[1]

		if myChoice == "X" {

			myPoints += 1

			if opponentChoice == "A" {
				myPoints += 3
			}
			
			if opponentChoice == "C" {
				myPoints += 6
			}

		} else if myChoice == "Y" {
			
			myPoints += 2

			if opponentChoice == "B" {
				myPoints += 3
			}

			if opponentChoice == "A" {
				myPoints += 6
			}

		} else if myChoice == "Z" {

			myPoints += 3

			if opponentChoice == "C" {
				myPoints += 3
			}

			if opponentChoice == "B" {
				myPoints += 6
			}

		}
		
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(myPoints)
}