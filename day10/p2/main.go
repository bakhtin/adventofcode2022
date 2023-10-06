package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Instruction struct {
	instructionName string
	instructionArg  int
}

func isSpriteOverlapsPixel(spritePos, pixelPos int) bool {
	if spritePos == pixelPos || spritePos-1 == pixelPos || spritePos+1 == pixelPos {
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
	const screenWidth = 40
	pixels := [screenWidth]byte{}

	for i := 0; i < len(instructions); i++ {
		instructionName, instructionArg := instructions[i].instructionName, instructions[i].instructionArg
		switch instructionName {
		case "addx":
			for instructionCycles := 0; instructionCycles < 2; instructionCycles++ {
				cycle += 1
				if (cycle-1)%screenWidth == 0 {
					fmt.Printf("%s\n", string(pixels[:]))
				}
				if isSpriteOverlapsPixel(regX, (cycle-1)%screenWidth) {
					pixels[(cycle-1)%screenWidth] = '#'
				} else {
					pixels[(cycle-1)%screenWidth] = '.'
				}
				if instructionCycles == 1 {
					regX += instructionArg
				}
			}
		case "noop":
			cycle += 1
			if (cycle-1)%screenWidth == 0 {
				fmt.Printf("%s\n", string(pixels[:]))
			}
			if isSpriteOverlapsPixel(regX, (cycle-1)%screenWidth) {
				pixels[(cycle-1)%screenWidth] = '#'
			} else {
				pixels[(cycle-1)%screenWidth] = '.'
			}
		}
	}
	fmt.Print(string(pixels[:]))
}
