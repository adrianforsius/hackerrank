package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	// We want minimum swaps do we need to save the intermidate states to calculate this?
	// For each swap we want at least one of the two get to its final position or so
	// that we can make something like a three way swap, say
	// We could swap 4 numbers, 1 - 2, 3 - 4, 1 - 3, 2 - 4. To optmize this we could
	// given [1,2,3,4] we would get [2,1,4,3] -> [4,3,2,1]
	// this could be done i two swaps, what if we want something that I suspsect can
	// be done in 3 steps like [4,2,3,1], or it could be done in 1 swap, Ok so something
	// that we can do in even swaps seems like we can check if we can do with less even
	// swaps and odd swaps we can do in less odd swaps lets try this out.
	// can we split this problem into smaller pieces? can we isolate part of the array?
	// do we need the whole context?
	// How do we start? can we loop from both ends?
	// split and concor, that is get two halfs upper and lower half

	a, count := sort(arr, 0)
	fmt.Println("array", a, "count", count)
	return int32(count)
}

func middleSorted(arr []int32) int {
	middle := int(math.Ceil(float64(len(arr)) / float64(2)))
	if len(arr)%2 == 1 {
		for i := 0; i < len(arr); i++ {
			if int(arr[i]) == middle {
				arr[i], arr[middle-1] = arr[middle-1], arr[i]
				return 1
			}
		}
	}
	return 0
}

func sort(arr []int32, count int) ([]int32, int) {
	count += middleSorted(arr)
	if len(arr) == 2 || len(arr) == 3 {
		if arr[0] > arr[len(arr)-1] {
			//fmt.Println("add count min")
			//fmt.Println("lowset", arr[0], "highest", arr[len(arr)-1])
			count += 1
			arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]
			fmt.Println("set block", arr)
		}
		return arr, count
	} else {
		//fmt.Println("arr", arr)
		mid := math.RoundToEven(float64(len(arr)) / 2.0)
		high := arr[int(math.Floor(float64(mid))):]
		low := arr[:int(math.Ceil(float64(mid)))]
		swaps := make([][]int32, 0)
		highCount := 0
		for i := 0; i < len(low); i++ {
			if low[i] > int32(mid) {
				for highCount < len(high) {
					if high[highCount] <= int32(mid) {
						count += 1
						swaps = append(swaps, []int32{int32(i), int32(highCount)})
						highCount++
						fmt.Println("swap", arr)
						break
					}
					highCount++
				}
			}
		}
		for _, pair := range swaps {
			low[pair[0]], high[pair[1]] = high[pair[1]], low[pair[0]]
		}
		//fmt.Println("low", low, "high", high)
		lowSorted, lowCount := sort(low, 0)
		highSorted, highCount := sort(high, 0)
		//fmt.Println("low", lowSorted, "high", highSorted)
		//fmt.Println("low", lowCount, "high", highCount)
		return append(lowSorted, highSorted...), highCount + lowCount + count
		// loop high and and swap to get on right side
	}
	return nil, 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
