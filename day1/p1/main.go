package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	maxCaloriesSum := 0
	maxCaloriesCandidateSum := 0

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			maxCaloriesCandidateSum += int(calories)
		} else {
			if maxCaloriesCandidateSum > maxCaloriesSum {
				maxCaloriesSum = maxCaloriesCandidateSum
			}
			maxCaloriesCandidateSum = 0
		}
	}
	f.Close()
	fmt.Println(maxCaloriesSum)
}
