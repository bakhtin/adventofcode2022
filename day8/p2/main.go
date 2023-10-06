package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func TreeScenicScore(trees []int, treeHeight int) int {
	treeScenicScore := 0
	for _, tree := range trees {
		if tree < treeHeight {
			treeScenicScore += 1
		} else {
			treeScenicScore += 1
			break
		}
	}
	return treeScenicScore
}

func main() {
	bestTreeScenicScore := 0
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
			if i != 0 && i != len(trees)-1 && j != 0 && j != len(trees[i])-1 {
				leftTrees := append([]int{}, trees[i][:j]...)
				rightTrees := append([]int{}, trees[i][j+1:]...)
				topTrees := make([]int, len(trees[:i]))
				bottomTrees := make([]int, len(trees[i+1:]))
				leftTreesScenicScore := 0
				rightTreesScenicScore := 0
				topTreesScenicScore := 0
				bottomTreesScenicScore := 0
				for k, tree := range trees[:i] {
					topTrees[k] = tree[j]
				}
				for k, tree := range trees[i+1:] {
					bottomTrees[k] = tree[j]
				}
				// left trees scenic score
				// reverse before calculating the score
				for k, l := 0, len(leftTrees)-1; k < l; k, l = k+1, l-1 {
					leftTrees[k], leftTrees[l] = leftTrees[l], leftTrees[k]
				}
				leftTreesScenicScore = TreeScenicScore(leftTrees, trees[i][j])
				// right trees scenic score
				rightTreesScenicScore = TreeScenicScore(rightTrees, trees[i][j])
				// top trees scenic score
				// reverse before calculating the score
				for k, l := 0, len(topTrees)-1; k < l; k, l = k+1, l-1 {
					topTrees[k], topTrees[l] = topTrees[l], topTrees[k]
				}
				topTreesScenicScore = TreeScenicScore(topTrees, trees[i][j])
				// bottom trees scenic score
				bottomTreesScenicScore = TreeScenicScore(bottomTrees, trees[i][j])
				currentBestTreesScenicScore := leftTreesScenicScore * rightTreesScenicScore * topTreesScenicScore * bottomTreesScenicScore
				if currentBestTreesScenicScore > bestTreeScenicScore {
					bestTreeScenicScore = currentBestTreesScenicScore
				}
			}
		}
	}
	fmt.Println(bestTreeScenicScore)
}
