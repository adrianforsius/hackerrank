package main

import (
	"fmt"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	// Loop over each action and change the structure accordingly
	// This looks like events could we possibly check if two events cancel each other out?
	// at each print there is a break, how about we split the query into parts of printing

	statusMap := make(map[int32]int32, 0)
	insertMap := make(map[int32]int32, 0)
	result := make([]int32, 0)
	for _, part := range queries {
		action, value := part[0], part[1]
		if isInsert(action) {
			if _, ok := statusMap[insertMap[value]]; !ok {
				statusMap[insertMap[value]] = 1
			} else if statusMap[insertMap[value]] > 0 {
				statusMap[insertMap[value]] -= 1
			}
			if _, ok := insertMap[value]; !ok {
				insertMap[value] = 1
			} else {
				insertMap[value] += 1
			}
			if _, ok := statusMap[insertMap[value]]; !ok {
				statusMap[insertMap[value]] = 1
			} else {
				statusMap[insertMap[value]] += 1
			}
			statusMap[statusMap[insertMap[value]]] += 1
		} else if isDelete(action) {
			if statusMap[insertMap[value]] > 0 {
				statusMap[insertMap[value]] -= 1
			}
			if _, ok := insertMap[value]; !ok {
				insertMap[value] = 1
			} else if insertMap[value] > 0 {
				insertMap[value] -= 1
			}
			if insertMap[value] > 0 {
				statusMap[insertMap[value]] += 1
			}
		} else if isPrint(action) {
			if val, ok := statusMap[value]; !ok {
				result = append(result, 0)
			} else {
				if val > 0 {
					result = append(result, 1)
				} else {
					result = append(result, 0)
				}
			}
		}
	}
	return result
}

func isInsert(s int32) bool {
	return s == 1
}
func isDelete(s int32) bool {
	return s == 2
}
func isPrint(s int32) bool {
	return s == 3
}

func main() {
	jumps := freqQuery([][]int32{
		//{1, 5},
		//{1, 6},
		//{3, 2},
		//{1, 10},
		//{1, 10},
		//{1, 6},
		//{2, 5},
		//{3, 2},
		{1, 3},
		{2, 3},
		{3, 2},
		{1, 4},
		{1, 5},
		{1, 5},
		{1, 4},
		{3, 2},
		{2, 4},
		{3, 2},
	})
	fmt.Println("jumps", jumps)
}
