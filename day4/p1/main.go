package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func makeRange(min int, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func RangeWithin(rangeA []int, rangeB []int) bool {
	if rangeA[0] >= rangeB[0] && rangeA[len(rangeA)-1] <= rangeB[len(rangeB)-1] || // rangeA fully within rangeB
		rangeA[0] <= rangeB[0] && rangeA[len(rangeA)-1] >= rangeB[len(rangeB)-1] { // rangeB fully within rangeA
		return true
	}
	return false
}

func main() {
	rangesFullyWithin := 0

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			ranges := strings.Split(line, ",")
			rangeAStr := strings.Split(ranges[0], "-")
			rangeBStr := strings.Split(ranges[1], "-")
			rangeAStart, err := strconv.ParseInt(rangeAStr[0], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			rangeAEnd, err := strconv.ParseInt(rangeAStr[1], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			rangeBStart, err := strconv.ParseInt(rangeBStr[0], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			rangeBEnd, err := strconv.ParseInt(rangeBStr[1], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			rangeA := makeRange(int(rangeAStart), int(rangeAEnd))
			rangeB := makeRange(int(rangeBStart), int(rangeBEnd))
			if RangeWithin(rangeA, rangeB) {
				rangesFullyWithin += 1
			}
		}
	}
	f.Close()
	fmt.Println(rangesFullyWithin)
}
