package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func calculateOutcome(p1 string, p2 string) (int, error) {
	xBonusScore := 1
	yBonusScore := 2
	zBonusScore := 3
	winScore := 6
	drawScore := 3
	lostScore := 0
	if p1 == "A" { // Rock
		if p2 == "X" {
			return drawScore + xBonusScore, nil
		} else if p2 == "Y" {
			return winScore + yBonusScore, nil
		} else if p2 == "Z" {
			return lostScore + zBonusScore, nil
		}
	} else if p1 == "B" { // Paper
		if p2 == "X" {
			return lostScore + xBonusScore, nil
		} else if p2 == "Y" {
			return drawScore + yBonusScore, nil
		} else if p2 == "Z" {
			return winScore + zBonusScore, nil
		}
	} else if p1 == "C" { // Scissors
		if p2 == "X" {
			return winScore + xBonusScore, nil
		} else if p2 == "Y" {
			return lostScore + yBonusScore, nil
		} else if p2 == "Z" {
			return drawScore + zBonusScore, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("Unknown round data provided. p1: %s, p2: %s", p1, p2))
}

func main() {
	totalSum := 0

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			p1Played := line[0]
			p2Played := line[2]
			roundOutcomeSum, err := calculateOutcome(string(p1Played), string(p2Played))
			if err != nil {
				log.Fatal(err)
			}
			totalSum += roundOutcomeSum
		}
	}
	f.Close()
	fmt.Println(totalSum)
}
