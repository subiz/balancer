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
	for i := 0; i < len(curNodes); i++ {
		next := cur + max
		if i > rem-1 {
			next = next - 1
		}
		board[curNodes[i]] = jobs[cur:next]
		cur = next
	}

	nextNodes := []int{1, 2, 3, 5, 8}
	fmt.Println("current", curNodes, board)
	m := rebalanceJobs(board, nextNodes)
	fmt.Println("next", nextNodes, m)
}

func rebalanceJobs(board map[int][]int, nextNodes []int) map[int][]int {
	totalJobs := 0
	curNodes := make(map[int]int, 0)
	keyNodes := make([]int, 0)

	for node, jobs := range board {
		totalJobs = totalJobs + len(jobs)
		curNodes[node] = node
	}

	jobNeedReAssign := make([]int, 0)
	for _, node := range curNodes {
		existed := false
		for _, n := range nextNodes {
			if node == n {
				existed = true
				break
			}
		}
		if !existed {
			jobNeedReAssign = append(jobNeedReAssign, board[node]...)
			delete(board, node)
			delete(curNodes, node)
		} else {
			keyNodes = append(keyNodes, node)
		}
	}

	newNodes := make([]int, 0)
	for _, node := range nextNodes {
		if _, ok := curNodes[node]; !ok {
			newNodes = append(newNodes, node)
		}
	}

	totalNodes := len(curNodes) + len(newNodes)
	min := int(math.Floor(float64(totalJobs) / float64(totalNodes)))

	// reduce job of prior last node
	rem := totalJobs % totalNodes
	sort.Sort(sort.IntSlice(keyNodes))
	for _, node := range keyNodes {
		jobs := board[node]
		l := len(jobs)
		if l > min && rem >= 0 {
			if rem > 0 {
				rem--
			} else if rem == 0 {
				jobNeedReAssign = append(jobNeedReAssign, jobs[min:]...)
				newJobs := removeElements(jobs, min, l)
				board[node] = newJobs
			}
		}
	}

	// add job to current empty node
	for node, jobs := range board {
		if len(jobs) == 0 {
			board[node] = append(board[node], jobNeedReAssign[0:min]...)
			jobNeedReAssign = removeElements(jobNeedReAssign, 0, min)
		}
	}

	// add job to new node
	rem = totalJobs % totalNodes
	for i := 0; i < len(newNodes); i++ {
		board[newNodes[i]] = append(board[newNodes[i]], jobNeedReAssign[0:min]...)
		jobNeedReAssign = removeElements(jobNeedReAssign, 0, min)
	}

	// add job to current node
	rem = totalJobs % totalNodes
	sort.Sort(sort.IntSlice(keyNodes))
	for _, node := range keyNodes {
		jobs := board[node]
		if len(jobs) == min && len(jobNeedReAssign) > 0 && rem > 0 {
			board[node] = append(board[node], jobNeedReAssign[0])
			jobNeedReAssign = removeElements(jobNeedReAssign, 0, 1)
			rem--
		} else if len(jobs) < min && len(jobNeedReAssign) > 0 {
			board[node] = append(board[node], jobNeedReAssign[0])
			jobNeedReAssign = removeElements(jobNeedReAssign, 0, 1)
		}
	}

	return board
}

func removeElements(s []int, from int, to int) []int {
	if from > len(s) || from > to {
		return s
	}
	return append(s[:from], s[to:]...)
}
