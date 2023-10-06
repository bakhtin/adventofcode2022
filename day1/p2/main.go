package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type IntHeap []int // max heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	maxCaloriesSum := 0
	maxCaloriesCandidateSum := 0
	topCalorieHolders := &IntHeap{}
	heap.Init(topCalorieHolders)

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			maxCaloriesCandidateSum += int(calories)
		} else {
			heap.Push(topCalorieHolders, maxCaloriesCandidateSum)
			maxCaloriesCandidateSum = 0
		}
	}
	f.Close()
	for i := 0; i < 3; i++ {
		maxCaloriesSum += heap.Pop(topCalorieHolders).(int)
	}
	fmt.Println(maxCaloriesSum)
}
