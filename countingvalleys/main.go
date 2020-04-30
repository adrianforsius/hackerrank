package main

import (
	"fmt"
)

// Complete the jumpingOnClouds function below.
func countingValleys(n int32, s string) int32 {
	altitude := 0
	valleys := 0
	runes := []rune(s)
	for _, unit := range runes {
		if unit == 'U' {
			altitude += 1
			if altitude == 0 {
				valleys += 1
			}
		} else {
			altitude -= 1
		}
	}
	return int32(valleys)
}

func main() {
	valley := countingValleys(8, "UDDDUDUU")
	fmt.Println("valley", valley)
}
