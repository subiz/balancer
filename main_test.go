package main

import (
	"sort"
	"testing"
)

func compareIntArr(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	sort.Ints(a1)
	sort.Ints(a2)

	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func compareM(c1, c2 map[int][]int) bool {
	if len(c1) != len(c2) {
		return false
	}

	for k, v1 := range c1 {
		v2, ok := c2[k]
		if !ok {
			return false
		}

		if !compareIntArr(v1, v2) {
			return false
		}
	}
	return true
}

func TestBalance(t *testing.T) {
	tcs := []struct {
		desc     string
		config   map[int][]int
		nodes    []int
		expected map[int][]int
	}{{
		desc: "test 1, remove node",
		config: map[int][]int{
			1: []int{1, 2},
			2: []int{3, 4},
			3: []int{5, 6},
			4: []int{7, 8},
		},
		nodes: []int{1, 2, 3},
		expected: map[int][]int{
			1: []int{1, 2, 7},
			2: []int{3, 4, 8},
			3: []int{5, 6},
		},
	}, {
		desc: "test 2, add node",
		config: map[int][]int{
			1: []int{1, 2},
			2: []int{3, 4},
			3: []int{5, 6},
			4: []int{7, 8},
		},
		nodes: []int{1, 2, 3},
		expected: map[int][]int{
			1: []int{1, 2},
			2: []int{3, 4},
			3: []int{5, 6},
			4: []int{7},
			5: []int{8},
		},
	}}

	for _, tc := range tcs {
		out := rebalanceJobs(tc.config, tc.nodes)
		if !compareM(out, tc.expected) {
			t.Fatalf("in test '%s', expect %v, got %v", tc.desc, tc.expected, out)
		}
	}
}
