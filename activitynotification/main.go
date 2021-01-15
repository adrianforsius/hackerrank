package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	days := int(d)
	var count int32
	var median int32
	sorted := expenditure[0:days]
	//sorted = qsort(sorted)
	sorted = sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	for i := days; i < len(expenditure); i++ {
		if days%2 == 0 {
			median = (sorted[int(math.Floor(float64(days/2)))] + sorted[int(math.Ceil(float64(days/2)))]) / 2
		} else {
			median = sorted[days/2]
		}
		val := expenditure[i]
		fmt.Println("val", expenditure[i])
		fmt.Println("sorted", sorted)
		fmt.Println("median", median)
		if expenditure[i] >= median*2 {
			count++
		}
		sorted = sorted[1:]
		e := sort.Search(len(sorted), func(e int) bool { return expenditure[i] < sorted[e] })
		if e < len(sorted) {
			sorted = append(sorted[:e], append([]int32{val}, sorted[e:]...)...)
		} else {
			sorted = append(sorted, val)
		}
	}
	return count
}

func qsort(a []int32) []int32 {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}

func main() {
	jumps := activityNotifications([]int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5)
	fmt.Println("jumps", jumps)
}
