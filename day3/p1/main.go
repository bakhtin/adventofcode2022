package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SliceIntersectionHash[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func ItemPriority(item byte) int {
	priority := 0
	if strings.ToLower(string(item)) == string(item) {
		priority = int(item) - 96
	} else {
		priority = int(item) - 64 + 26
	}

	return priority
}

func main() {
	prioritySum := 0

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			rucksackItemsCount := len(line)
			compartmentAItems := line[0 : rucksackItemsCount/2]
			compartmentBItems := line[rucksackItemsCount/2:]
			commonItems := SliceIntersectionHash([]byte(compartmentAItems), []byte(compartmentBItems))
			prioritySum += ItemPriority(commonItems[0])
		}
	}
	f.Close()
	fmt.Println(prioritySum)
}
