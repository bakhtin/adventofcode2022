package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

type Coordinates struct {
	x int
	y int
}

type Move struct {
	direction byte
	distance  int
}

func isAdjacent(headPos, tailPos Coordinates) bool {
	if math.Abs(float64(headPos.x-tailPos.x)) <= 1.0 && math.Abs(float64(headPos.y-tailPos.y)) <= 1.0 {
		return true
	}
	return false
}

func isDiagonalMove(headPos, tailPos Coordinates) bool {
	if math.Abs(float64(headPos.x-tailPos.x)) >= 1.0 || math.Abs(float64(headPos.y-tailPos.y)) >= 1.0 {
		return true
	}
	return false
}

func visitCoordinates(visitedPos map[Coordinates]int, tailPos Coordinates) {
	_, ok := visitedPos[tailPos]
	if ok {
		visitedPos[tailPos] += 1
	} else {
		visitedPos[tailPos] = 1
	}
}

func main() {
	var moves []Move

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {

			moveDirection := line[0]
			moveDistance, err := strconv.Atoi(line[2:])
			if err != nil {
				log.Fatal(err)
			}
			moves = append(moves, Move{moveDirection, moveDistance})
		}
	}
	f.Close()

	headPos := Coordinates{0, 0}
	tailPos := Coordinates{0, 0}
	tailVisitedCoordinates := make(map[Coordinates]int)
	visitCoordinates(tailVisitedCoordinates, tailPos)
	for i := 0; i < len(moves); i++ {
		moveDirection, moveDistance := moves[i].direction, moves[i].distance
		switch moveDirection {
		case 'R':
			for j := 0; j < moveDistance; j++ {
				headPos.x += 1
				if !isAdjacent(headPos, tailPos) {
					if isDiagonalMove(headPos, tailPos) {
						tailPos.y = headPos.y
					}
					tailPos.x += 1
				}
				visitCoordinates(tailVisitedCoordinates, tailPos)
			}
		case 'L':
			for j := 0; j < moveDistance; j++ {
				headPos.x -= 1
				if !isAdjacent(headPos, tailPos) {
					if isDiagonalMove(headPos, tailPos) {
						tailPos.y = headPos.y
					}
					tailPos.x -= 1
				}
				visitCoordinates(tailVisitedCoordinates, tailPos)
			}
		case 'D':
			for j := 0; j < moveDistance; j++ {
				headPos.y -= 1
				if !isAdjacent(headPos, tailPos) {
					if isDiagonalMove(headPos, tailPos) {
						tailPos.x = headPos.x
					}
					tailPos.y -= 1
				}
				visitCoordinates(tailVisitedCoordinates, tailPos)
			}
		case 'U':
			for j := 0; j < moveDistance; j++ {
				headPos.y += 1
				if !isAdjacent(headPos, tailPos) {
					if isDiagonalMove(headPos, tailPos) {
						tailPos.x = headPos.x
					}
					tailPos.y += 1
				}
				visitCoordinates(tailVisitedCoordinates, tailPos)
			}
		}
	}
	log.Println("Visited times", len(tailVisitedCoordinates))
}
