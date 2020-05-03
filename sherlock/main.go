package main

import (
	"fmt"
	"math"
)

// Complete the jumpingOnClouds function below.
// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	middle := int(math.Ceil(float64(len(s))/2.0) - 1)
	runes := []rune(s)
	anagramCount := 0
	// a word can't start after the middle point
	var word []rune
	for startIndex := 0; startIndex < middle; startIndex++ {
		// We only want to prepend a letter if
		// an angram streak is broken or a new
		// intent needs to be started
		if len(word) == 0 {
			word = append(word, runes[startIndex])
		}

		// To keep discovery going
		endIndex := 0
		run := true
		for run == true {
			endIndex += 1
			result := FindReverse(runes, startIndex, endIndex)
			if result == 0 {
				run = false
			} else {
				anagramCount += 1
			}
		}
		if len(word) <= 1 {
			break
		}
		_, word = word[0], word[1:]
	}
	return int32(anagramCount)
}

func FindReverse(runes []rune, startIndex, endIndex int) int {
	//for i := len(runes) - 1; i >= 0; i-- {
	//        index := 0
	//        backIndex := i
	//        streak := 0
	//        if word[index] == runes[i] {
	//                streak++
	//                for index < len(word) {
	//                        backIndex--
	//                        index++
	//                        if len(word) == streak {
	//                                // If is the same word it doesn't count
	//                                if startIndex == backIndex && endIndex == index {
	//                                        return 0
	//                                }
	//                                return 1
	//                        }
	//                        if word[index] == runes[backIndex] {
	//                                streak++
	//                        } else {
	//                                return 0
	//                        }
	//                }
	//        }
	//}

	word := runes[startIndex:endIndex]

	reversed := Reverse(runes)
	middle := int(math.Ceil(float64(len(runes))/2.0) - 1)
	for i := 0; i < middle; i++ {
		if runes[i] == reversed[i] {
			index := 0
			tempCount := 0
			for _, l := range word {
				if runes[i+index] == l {
					tempCount += 1
					if tempCount == len(word) {
						return 1
					}
				}
				index++
			}
		}
	}
	return 0
}

func Reverse(runes []rune) []rune {
	var reverseHalf []rune
	for i := len(runes) - 1; i >= 0; i-- {
		reverseHalf = append([]rune{runes[i]}, reverseHalf...)
	}
	return reverseHalf
}

func main() {
	jumps := sherlockAndAnagrams("ifailuhkqq")
	fmt.Println("jumps", jumps)

	jumps = sherlockAndAnagrams("abba")
	fmt.Println("jumps", jumps)
}
