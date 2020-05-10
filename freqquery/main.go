package main

import (
	"fmt"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	// Loop over each action and change the structure accordingly
	// This looks like events could we possibly check if two events cancel each other out?
	// at each print there is a break, how about we split the query into parts of printing
	//var orig [][]int32
	//copy(orig, queries)

	split := make([][][]int32, 0)
	split = append(split, make([][]int32, 0))
	index := 0
	for i, item := range queries {
		split[index] = append(split[index], item)
		// remove last iteration if last iteration doesn't end with print
		if len(queries)-1 == i {
			if !isPrint(item) {
				split = split[:len(split)-1]
			}
			continue
		}
		if isPrint(item) {
			index++
			split = append(split, make([][]int32, 0))
		}
	}
	insertMap := make(map[int32]int32, 0)
	result := make([]int32, 0)
	for _, stack := range split {
		for _, part := range stack {
			if isInsert(part) {
				if _, ok := insertMap[part[1]]; !ok {
					insertMap[part[1]] = 1
				} else {
					insertMap[part[1]] += 1
				}
			}
			if isDelete(part) {
				if _, ok := insertMap[part[1]]; ok {
					insertMap[part[1]] -= 1
				}
			}
		}
		check := false
		for _, entry := range insertMap {
			if entry <= 0 {
				continue
			}
			if entry == stack[len(stack)-1][1] {
				check = true
				break
			}
		}
		if check == false {
			result = append(result, 0)
		} else {
			result = append(result, 1)
		}
	}
	return result
}

func isInsert(s []int32) bool {
	return s[0] == 1
}
func isDelete(s []int32) bool {
	return s[0] == 2
}
func isPrint(s []int32) bool {
	return s[0] == 3
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
