package main

import (
	"fmt"
	"math"
)

// Complete the jumpingOnClouds function below.
// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	tripleMap := make([][]int, 0)
	for i := 0; i < len(arr)-2; i++ {
		for e := 0; e < len(arr); e++ {
			item := arr[e]
			if int(math.Log(float64(item))/math.Log(float64(r)))%1 == 0 {
				//fmt.Println("r", r, "log result", math.Log(float64(item))/math.Log(float64(r)), "item", item)
				tripleMap = availableMaps(tripleMap)
				tripleMap = addToTripple(tripleMap, r, item)
			}
		}
	}
	fmt.Println(tripleMap)
	return int64(len(fullMaps(tripleMap)))
}

//func insert(item int64, index int64, [][]int) [][]int {
//}
func addToTripple(tripleMap [][]int, r, item int64) [][]int {
	base := math.Log(float64(item)) / math.Log(float64(r))
	for i, list := range tripleMap {
		// Get the last item and each list and make sure its one lower
		//fmt.Println("last item", list[len(list)-1], "base next", math.Log(float64(item))/math.Log(float64(r)))
		if float64(list[len(list)-1]+1) == base {
			//fmt.Println("add to tripple", tripleMap, "r", r, "item", item, "base", base)
			//fmt.Println(list)
			tripleMap[i] = append(list, int(base))
			return tripleMap
		}
	}
	return append(tripleMap, []int{int(base)})
}

func availableMaps(tripleMap [][]int) [][]int {
	filtered := make([][]int, 0)
	for _, list := range tripleMap {
		if len(list) < 3 {
			filtered = append(filtered, list)
		}
	}
	return filtered
}

func fullMaps(tripleMap [][]int) [][]int {
	filtered := make([][]int, 0)
	for _, list := range tripleMap {
		if len(list) == 3 {
			filtered = append(filtered, list)
		}
	}
	return filtered
}

func main() {
	jumps := countTriplets([]int64{1, 2, 2, 4}, 2)
	fmt.Println("jumps", jumps)
}
