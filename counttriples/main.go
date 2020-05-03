package main

import (
	"fmt"
)

func countTriplets(arr []int64, r int64) int64 {
	m2 := make(map[int64]int, 0)
	m3 := make(map[int64]int, 0)

	var count int
	for _, item := range arr {
		if _, ok := m3[item]; ok {
			count += m3[item]
		}
		if _, ok := m2[item]; ok {
			m3[item*r] += m2[item]
		}
		m2[item*r] += 1
	}
	fmt.Println("maps2", m2, "maps3", m3)
	return int64(count)
}

func main() {
	jumps := countTriplets([]int64{1, 2, 2, 4}, 2)
	fmt.Println("jumps", jumps)
}
