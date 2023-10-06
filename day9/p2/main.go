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

type Rope struct {
	knots [10]Coordinates
}

func isAdjacent(headPos, tailPos Coordinates) bool {
	if math.Abs(float64(headPos.x-tailPos.x)) <= 1.0 && math.Abs(float64(headPos.y-tailPos.y)) <= 1.0 {
		return true
	}
	return false
}

func isDiagonalMove(headPos, tailPos Coordinates) bool {
	if (math.Abs(float64(headPos.x-tailPos.x)) == 1.0 && math.Abs(float64(headPos.y-tailPos.y)) > 1.0) ||
		(math.Abs(float64(headPos.x-tailPos.x)) > 1.0 && math.Abs(float64(headPos.y-tailPos.y)) == 1.0) {
		return true
	}
	return false
}

func saveCoordinates(visitedPos map[Coordinates]int, tailPos Coordinates) {
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

	elvesRope := Rope{}
	tailVisitedCoordinates := make(map[Coordinates]int)
	saveCoordinates(tailVisitedCoordinates, elvesRope.knots[len(elvesRope.knots)-1])
	for i := 0; i < len(moves); i++ {
		moveDirection, moveDistance := moves[i].direction, moves[i].distance
		switch moveDirection {
		case 'R':
			for j := 0; j < moveDistance; j++ {
				elvesRope.knots[0].x += 1
				for knot := 1; knot < len(elvesRope.knots); knot++ {
					if !isAdjacent(elvesRope.knots[knot], elvesRope.knots[knot-1]) {
						if isDiagonalMove(elvesRope.knots[knot], elvesRope.knots[knot-1]) {
							yInc := 1
							if elvesRope.knots[knot-1].y-elvesRope.knots[knot].y < 0 {
								yInc = -1
							}
							elvesRope.knots[knot].y += yInc
						}
						elvesRope.knots[knot].x += 1
					}
					if knot == len(elvesRope.knots)-1 {
						saveCoordinates(tailVisitedCoordinates, elvesRope.knots[knot])
					}
				}
			}
		case 'L':
			for j := 0; j < moveDistance; j++ {
				elvesRope.knots[0].x -= 1
				for knot := 1; knot < len(elvesRope.knots); knot++ {
					if !isAdjacent(elvesRope.knots[knot], elvesRope.knots[knot-1]) {
						if isDiagonalMove(elvesRope.knots[knot], elvesRope.knots[knot-1]) {
							yInc := 1
							if elvesRope.knots[knot-1].y-elvesRope.knots[knot].y < 0 {
								yInc = -1
							}
							elvesRope.knots[knot].y += yInc
						}
						elvesRope.knots[knot].x -= 1
					}
					if knot == len(elvesRope.knots)-1 {
						saveCoordinates(tailVisitedCoordinates, elvesRope.knots[knot])
					}
				}
			}
		case 'D':
			for j := 0; j < moveDistance; j++ {
				elvesRope.knots[0].y -= 1
				for knot := 1; knot < len(elvesRope.knots); knot++ {
					if !isAdjacent(elvesRope.knots[knot], elvesRope.knots[knot-1]) {
						if isDiagonalMove(elvesRope.knots[knot], elvesRope.knots[knot-1]) {
							xInc := 1
							if elvesRope.knots[knot-1].x-elvesRope.knots[knot].x < 0 {
								xInc = -1
							}
							elvesRope.knots[knot].x += xInc
						}
						elvesRope.knots[knot].y -= 1
					}
					if knot == len(elvesRope.knots)-1 {
						saveCoordinates(tailVisitedCoordinates, elvesRope.knots[knot])
					}
				}
			}
		case 'U':
			for j := 0; j < moveDistance; j++ {
				elvesRope.knots[0].y += 1
				for knot := 1; knot < len(elvesRope.knots); knot++ {
					if !isAdjacent(elvesRope.knots[knot], elvesRope.knots[knot-1]) {
						if isDiagonalMove(elvesRope.knots[knot], elvesRope.knots[knot-1]) {
							xInc := 1
							if elvesRope.knots[knot-1].x-elvesRope.knots[knot].x < 0 {
								xInc = -1
							}
							elvesRope.knots[knot].x += xInc
						}
						elvesRope.knots[knot].y += 1
					}
					if knot == len(elvesRope.knots)-1 {
						saveCoordinates(tailVisitedCoordinates, elvesRope.knots[knot])
					}
				}
			}
		}
	}
	log.Println("Visited times", len(tailVisitedCoordinates))
}
