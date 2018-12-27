package main

import (
	"fmt"
	"math"
)

type JobQueue []int32

func (q *JobQueue) Slice(n int) {
}

type ScheduleJobs map[int32][]int32

func main() {
	jobs := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	curNodes := []int32{1, 2, 3, 4, 5, 6}
	board := make(ScheduleJobs, 0)
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

	nextNodes := []int32{1, 2, 3, 5, 8}
	fmt.Println("current", curNodes, board)
	m := rebalanceJobs(board, nextNodes)
	fmt.Println("next", nextNodes, m)
}

func rebalanceJobs(board ScheduleJobs, nextNodes []int32) ScheduleJobs {
	totalJobs := 0
	curNodes := make(map[int32]int32, 0)
	for node, jobs := range board {
		totalJobs = totalJobs + len(jobs)
		curNodes[node] = node
	}

	jobNeedReAssign := make([]int32, 0)
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
		}
	}

	newNodes := make([]int32, 0)
	for _, node := range nextNodes {
		if _, ok := curNodes[node]; !ok {
			newNodes = append(newNodes, node)
		}
	}

	totalNodes := len(curNodes) + len(newNodes)
	max := int(math.Ceil(float64(totalJobs) / float64(totalNodes)))
	// fmt.Println(max, totalJobs, curNodes)
	rem := totalJobs % totalNodes

	for node, jobs := range board {
		l := len(jobs)
		if l > max {
			s := max - 1
			jobNeedReAssign = append(jobNeedReAssign, jobs[s:]...)
			newJobs := append(jobs[:s], jobs[s:]...)
			board[node] = newJobs
		}
	}

	cur := 0
	for i := 0; i < len(newNodes); i++ {
		next := cur + max
		if i > rem-1 {
			next = next - 1
		}
		board[newNodes[i]] = jobNeedReAssign[cur:next]
		cur = next
	}

	return board
}
