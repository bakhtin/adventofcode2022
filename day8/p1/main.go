package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	visibleTreesCount := 0
	var trees [][]int

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			var rowTrees []int
			for _, c := range line {
				rowTrees = append(rowTrees, int(c-'0'))
			}
			trees = append(trees, rowTrees)
		}
	}
	f.Close()

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			// A tree on the edge is always visible
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(trees[i])-1 {
				visibleTreesCount += 1
			} else {
				leftTrees := append([]int{}, trees[i][:j]...)
				rightTrees := append([]int{}, trees[i][j+1:]...)
				topTrees := make([]int, len(trees[:i]))
				bottomTrees := make([]int, len(trees[i+1:]))
				for k, tree := range trees[:i] {
					topTrees[k] = tree[j]
				}
				for k, tree := range trees[i+1:] {
					bottomTrees[k] = tree[j]
				}
				sort.Ints(leftTrees)
				sort.Ints(rightTrees)
				sort.Ints(topTrees)
				sort.Ints(bottomTrees)
				if trees[i][j] > leftTrees[len(leftTrees)-1] ||
					trees[i][j] > rightTrees[len(rightTrees)-1] ||
					trees[i][j] > topTrees[len(topTrees)-1] ||
					trees[i][j] > bottomTrees[len(bottomTrees)-1] {
					visibleTreesCount += 1
				}
			}
		}
	}
	fmt.Println(visibleTreesCount)
}
