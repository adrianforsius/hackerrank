package main

import (
	"fmt"
)

// Complete the jumpingOnClouds function below.
func twoStrings(s1 string, s2 string) string {
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)
	for _, item1 := range []rune(s1) {
		s1Map[item1] = 0
	}
	for _, item2 := range []rune(s2) {
		s2Map[item2] = 0
	}
	for i, _ := range s1Map {
		if _, ok := s2Map[i]; ok {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	jumps := twoStrings("hello", "world")
	fmt.Println("jumps", jumps)
}
