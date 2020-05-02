package main

import (
	"fmt"
)

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	size := len(a)
	rest := int(d) % size
	part2, part1 := a[:rest], a[rest:]
	return append(part1, part2...)
}

func main() {
	jumps := rotLeft([]int32{1, 2, 3, 4, 5}, 4)
	fmt.Println("jumps", jumps)
}
