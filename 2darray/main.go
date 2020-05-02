package main

import (
	"fmt"
)

func hourglassSum(arr [][]int32) int32 {
	var biggest int32
	for row := 0; row < len(arr)-2; row++ {
		for col := 0; col < len(arr[0])-2; col++ {
			row1 := arr[row][col : col+3]
			//row2 := arr[row+1][col : col+2]
			row3 := arr[row+2][col : col+3]

			var sum int32
			for _, item := range row1 {
				sum += item
			}
			for _, item := range row3 {
				sum += item
			}
			sum += arr[row+1][col+1]
			if col == 0 && row == 0 {
				biggest = sum
			}
			if sum > biggest {
				biggest = sum
			}
		}
	}
	return biggest
}

func main() {
	sum := hourglassSum([][]int32{
		{
			1, 1, 1, 0, 0, 0,
		},
		{
			0, 1, 0, 0, 0, 0,
		},
		{
			1, 1, 1, 0, 0, 0,
		},
		{
			0, 0, 2, 4, 4, 0,
		},
		{
			0, 0, 0, 2, 0, 0,
		},
		{
			0, 0, 1, 2, 4, 0,
		},
	})
	fmt.Println("sum", sum)
}
