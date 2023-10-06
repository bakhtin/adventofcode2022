package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func HasRepetitions(l []int) bool {
	for i := 0; i < len(l); i++ {
		for j := len(l) - 1; j > i; j-- {
			if l[i] == l[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	patternPos := 0
	slidingWindow := []int{}

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			for pos, c := range line {
				slidingWindow = append(slidingWindow, int(c))
				if pos >= 3 {
					if !HasRepetitions(slidingWindow) {
						patternPos = pos
						break
					}
					slidingWindow = slidingWindow[1:]
				}
			}
		}
	}
	f.Close()
	fmt.Println(patternPos + 1)
}
