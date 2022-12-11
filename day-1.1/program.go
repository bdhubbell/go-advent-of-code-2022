package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var calorieSums []int
	var runningSum int

	for scanner.Scan() {
		lineValue := scanner.Text()

		if (lineValue == "") {
			calorieSums = append(calorieSums, runningSum);
			runningSum = 0

			continue;
		}

		calorieValue, err := strconv.Atoi(scanner.Text())
		runningSum += calorieValue

		if err != nil {
			log.Fatal(err)
		}
	}

	calorieSums = append(calorieSums, runningSum);

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calorieSums)))


	fmt.Println(calorieSums[0])
}