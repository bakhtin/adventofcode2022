package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// [T]     [D]         [L]
// [R]     [S] [G]     [P]         [H]
// [G]     [H] [W]     [R] [L]     [P]
// [W]     [G] [F] [H] [S] [M]     [L]
// [Q]     [V] [B] [J] [H] [N] [R] [N]
// [M] [R] [R] [P] [M] [T] [H] [Q] [C]
// [F] [F] [Z] [H] [S] [Z] [T] [D] [S]
// [P] [H] [P] [Q] [P] [M] [P] [F] [D]
//  1   2   3   4   5   6   7   8   9

func ParseMove(move string) (map[string]int, error) {
	s := strings.Split(move, " ")
	if s[0] != "move" {
		return make(map[string]int), errors.New(fmt.Sprintf("Illegal move: %s", move))
	}
	count, _ := strconv.Atoi(s[1])
	stackFrom, _ := strconv.Atoi(s[3])
	stackTo, _ := strconv.Atoi(s[5])
	parsedMove := make(map[string]int)
	parsedMove["count"] = count
	parsedMove["stackFrom"] = stackFrom - 1
	parsedMove["stackTo"] = stackTo - 1
	return parsedMove, nil
}

func main() {
	stacks := [][]byte{
		{'P', 'F', 'M', 'Q', 'W', 'G', 'R', 'T'},
		{'H', 'F', 'R'},
		{'P', 'Z', 'R', 'V', 'G', 'H', 'S', 'D'},
		{'Q', 'H', 'P', 'B', 'F', 'W', 'G'},
		{'P', 'S', 'M', 'J', 'H'},
		{'M', 'Z', 'T', 'H', 'S', 'R', 'P', 'L'},
		{'P', 'T', 'H', 'N', 'M', 'L'},
		{'F', 'D', 'Q', 'R'},
		{'D', 'S', 'C', 'N', 'L', 'P', 'H'},
	}

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "move") {
			parsedMove, err := ParseMove(line)
			if err != nil {
				log.Fatal(err)
			}
			stackFrom := parsedMove["stackFrom"]
			stackTo := parsedMove["stackTo"]
			count := parsedMove["count"]
			movedStackReverse := stacks[stackFrom][len(stacks[stackFrom])-count:]
			for i := 0; i < len(movedStackReverse); i++ {
				stacks[stackTo] = append(stacks[stackTo], movedStackReverse[i])
			}
			stacks[stackFrom] = stacks[stackFrom][:len(stacks[stackFrom])-count]
		}
	}
	f.Close()
	for _, peek := range stacks {
		fmt.Printf("%s", string(peek[len(peek)-1]))
	}
}
