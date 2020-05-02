package main

import (
	"fmt"
	"math"
)

//moves := 0
//moved := make(map[int32]int)
//for i := len(q) - 1; i >= 1; i-- {
//        num, ok := moved[q[i]]
//        if !ok {
//                moved[q[i]] = 0
//        }

//        if num < 2 {
//                moves += 1
//                moved[q[i]] += 1
//                q[i-1], q[i] = q[i], q[i-1]
//        }
//}
//fmt.Println(moves)
// Complete the minimumBribes function below.
//func minimumBribes(q []int32) {
//        var bribes int32
//        for i, item := range q {
//                if int(item)-2 < i {
//                        fmt.Println("Too chaotic")
//                        return
//                }
//                for check := i - 1; check > 0; check-- {
//                        if q[check] > q[i] {
//                                bribes += 1
//                        }
//                }
//        }
//        fmt.Println(bribes)
//}
func minimumBribes(q []int32) {
	var bribes int32
	for i, item := range q {
		if i < int(item)-3 {
			fmt.Println("Too chaotic")
			return
		}
		for check := i - 1; check >= int(float64(math.Max(int(q[i])-2), float64(0))); check-- {
			if q[check] > q[i] {
				bribes += 1
			}
		}
	}
	fmt.Println(bribes)
}

func main() {
	minimumBribes([]int32{2, 1, 3, 5, 6, 4})
	//fmt.Println("jumps", jumps)
}

//[]int32{3,4,5,6,7,8,9,1,2}
//[]int32{1,2,3,4,5,6,7,8,9}
