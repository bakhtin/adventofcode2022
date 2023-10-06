package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Instruction struct {
	instructionName string
	instructionArg  int
}

func isInterestingCycle(cycle int) bool {
	if (cycle+20)%40 == 0 || cycle == 20 {
		return true
	}
	return false
}

func main() {
	var instructions []Instruction

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			instructionName := line[0:4]
			instructionArg := 0
			if len(line) > 4 {
				instructionArg, err = strconv.Atoi(line[5:])
			}
			if err != nil {
				log.Fatal(err)
			}
			instructions = append(instructions, Instruction{instructionName, instructionArg})
		}
	}
	f.Close()

	regX := 1
	cycle := 0
	signalStrength := 0

	for i := 0; i < len(instructions); i++ {
		instructionName, instructionArg := instructions[i].instructionName, instructions[i].instructionArg
		switch instructionName {
		case "addx":
			for instructionCycles := 0; instructionCycles < 2; instructionCycles++ {
				cycle += 1
				if isInterestingCycle(cycle) {
					signalStrength += cycle * regX
				}
				if instructionCycles == 1 {
					regX += instructionArg
				}
			}
		case "noop":
			cycle += 1
			if isInterestingCycle(cycle) {
				signalStrength += cycle * regX
			}
		}
	}
	log.Printf("Total signal strength: %d", signalStrength)
}
