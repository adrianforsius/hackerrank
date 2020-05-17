package main

import (
	"fmt"
)

// Complete the maximumToys function below.

func maximumToys(prices []int32, k int32) int32 {
	sorted := make([]int32, 0)
	for _, p := range prices {
		// Filter overpriced toys
		if p > k {
			continue
		}

		sorted = insert(sorted, p)
	}

	var sum int32
	var count int
	fmt.Println(sorted)
	for i := len(sorted) - 1; i >= 0; i-- {
		//fmt.Println(sorted)
		//for i, s := range sorted {
		sum += sorted[i]
		fmt.Println("value", sorted[i], "sum", sum)
		count++
		if sum >= k {
			return int32(count)
		}
	}
	return int32(count)
}

func insert(list []int32, val int32) []int32 {
	for i := 0; i < len(list); i++ {
		if val < list[i] {
			return append(list[:i], append([]int32{val}, list[i:]...)...)
		}
	}
	return append(list, val)
}

func main() {
	// We can sort and then loop again over from the start until max amount is reach
	// this will however take n**2 + n assuming we sort with something like bubblesort
	// insert sort seems like a great choice since we will do 2*n
	toys := maximumToys([]int32{1, 12, 5, 111, 200, 1000, 10, 10, 2, 2, 2, 0, 0}, 50)
	fmt.Println("jumps", toys)
}
