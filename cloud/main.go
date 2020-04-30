package main

import (
	"fmt"
	//. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

//func TestOrder(t *testing.T) {
//        RegisterFailHandler(Fail)
//        RunSpecs(t, "One suite")
//}

//var _ = Describe("as a user with random input", func() {
//        Describe("and last value odd", func() {
//                It("should return a bigger value than a even run", func() {
//                        //Expect(true).To(Equal(false))
//                })
//        })
//})

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	// Jump two if we can
	// otherwise jump one
	jumps := make([]int, 0)

	i := 0
	fmt.Println(c, i)
	for 2 < len(c) {
		step2 := c[2]
		if step2 != 1 {
			i += 2
			c = c[2:]
		} else {
			i += 1
			c = c[1:]
		}
		fmt.Println("step", c, i)
		jumps = append(jumps, i)
	}
	return int32(len(jumps))
}

func main() {
	jumps := jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0})
	fmt.Println("jumps", jumps)

	jumps = jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0})
	fmt.Println("jumps", jumps)
}
