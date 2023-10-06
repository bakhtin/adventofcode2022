package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func p2Turn(p1 string, expectedOutcome string) (string, error) {
	if p1 == "A" {
		if expectedOutcome == "X" {
			return "Z", nil
		} else if expectedOutcome == "Y" {
			return "X", nil
		} else if expectedOutcome == "Z" {
			return "Y", nil
		}
	} else if p1 == "B" {
		if expectedOutcome == "X" {
			return "X", nil
		} else if expectedOutcome == "Y" {
			return "Y", nil
		} else if expectedOutcome == "Z" {
			return "Z", nil
		}
	} else if p1 == "C" {
		if expectedOutcome == "X" {
			return "Y", nil
		} else if expectedOutcome == "Y" {
			return "Z", nil
		} else if expectedOutcome == "Z" {
			return "X", nil
		}
	}
	return "", errors.New(fmt.Sprintf("Unknown round data provided. p1: %s, expectedOutcome: %s", p1, expectedOutcome))
}

func calculateOutcomeSum(p1 string, p2 string) (int, error) {
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
			expectedOutcome := line[2]
			expectedP2Turn, err := p2Turn(string(p1Played), string(expectedOutcome))
			roundOutcomeSum, err := calculateOutcomeSum(string(p1Played), string(expectedP2Turn))
			if err != nil {
				log.Fatal(err)
			}
			totalSum += roundOutcomeSum
		}
	}
	f.Close()
	fmt.Println(totalSum)
}
