package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	curNodes := []int{1, 2, 3, 4, 5, 6}
	board := make(map[int][]int, 0)
	max := int(math.Ceil(float64(len(jobs)) / float64(len(curNodes))))
	rem := len(jobs) % len(curNodes)
	cur := 0
	for i, node := range curNodes {
		next := cur + max
		if i > rem-1 {
			next = next - 1
		}
		board[node] = jobs[cur:next]
		cur = next
	}

	nextNodes := []int{1, 2, 3, 5, 6}
	fmt.Println("current", curNodes, board)
	m := rebalanceJobs(board, nextNodes)
	fmt.Println("next", nextNodes, m)
}

func rebalanceJobs(board map[int][]int, nextNodes []int) map[int][]int {
	totalJobs := 0
	curNodes := make(map[int]int, 0)
	delNodes := make(map[int]int, 0)
	newNodes := make([]int, 0)
	reassignJobs := make([]int, 0)

	// keyNodes use to sort nodes
	keyNodes := make([]int, 0)

	for node, jobs := range board {
		totalJobs += len(jobs)

		if in := inSlice(nextNodes, node); in {
			keyNodes = append(keyNodes, node)
			curNodes[node] = node
		} else {
			reassignJobs = append(reassignJobs, jobs...)
			delNodes[node] = node
		}
	}

	for _, node := range nextNodes {
		if _, ok := curNodes[node]; !ok {
			newNodes = append(newNodes, node)
		}
	}

	totalNodes := len(curNodes) + len(newNodes)
	min := int(math.Floor(float64(totalJobs) / float64(totalNodes)))
	rem := totalJobs % totalNodes

	sort.Sort(sort.IntSlice(keyNodes))

	// reduce job of last node
	orem := rem
	for _, node := range keyNodes {
		jobs := board[node]
		l := len(jobs)
		if l > min && orem >= 0 {
			if orem > 0 {
				orem--
			} else if orem == 0 {
				reassignJobs = append(reassignJobs, jobs[min:]...)
				newJobs := removeElements(jobs, min, l)
				board[node] = newJobs
			}
		}
	}

	for node, jobs := range board {
		// del node
		if _, ok := delNodes[node]; ok {
			delete(board, node)
			continue
		}
		// add job to node empty
		if len(jobs) == 0 {
			board[node] = append(board[node], reassignJobs[0:min]...)
			reassignJobs = removeElements(reassignJobs, 0, min)
		}
	}

	// add job to new node
	for i := 0; i < len(newNodes); i++ {
		board[newNodes[i]] = append(board[newNodes[i]], reassignJobs[0:min]...)
		reassignJobs = removeElements(reassignJobs, 0, min)
	}

	// add more job to current node
	nrem := rem
	for _, node := range keyNodes {
		jobs := board[node]
		addMore := false
		if len(jobs) == min && len(reassignJobs) > 0 && nrem > 0 {
			addMore = true
			nrem--
		} else if len(jobs) < min && len(reassignJobs) > 0 {
			addMore = true
		}

		if addMore {
			board[node] = append(board[node], reassignJobs[0])
			reassignJobs = removeElements(reassignJobs, 0, 1)
		}
	}

	return board
}

func inSlice(m []int, k int) bool {
	for _, j := range m {
		if j == k {
			return true
		}
	}
	return false
}

func removeElements(s []int, from int, to int) []int {
	if from > len(s) || from > to {
		return s
	}
	return append(s[:from], s[to:]...)
}
