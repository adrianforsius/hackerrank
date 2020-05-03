package main

import (
	"fmt"
)

// Complete the jumpingOnClouds function below.
// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	reverse := Reverse(s)
	for i := 0; i < len(s); i++ {

	}
	runes := []rune(s)
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

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Complete the jumpingOnClouds function below.
// Complete the sherlockAndAnagrams function below.
//func sherlockAndAnagrams(s string) int32 {
//        middle := int(math.Ceil(float64(len(s))/2.0) - 1)
//        runes := []rune(s)
//        anagramCount := 0
//        // a word can't start after the middle point
//        var word []rune
//        for startIndex := 0; startIndex < middle; startIndex++ {
//                // We only want to prepend a letter if
//                // an angram streak is broken or a new
//                // intent needs to be started
//                if len(word) == 0 {
//                        word = append(word, runes[startIndex])
//                }

//                // To keep discovery going
//                endIndex := 0
//                run := true
//                for run == true {
//                        endIndex += 1
//                        result := FindReverse(runes, startIndex, endIndex)
//                        if result == 0 {
//                                run = false
//                        } else {
//                                anagramCount += 1
//                        }
//                }
//                if len(word) <= 1 {
//                        break
//                }
//                _, word = word[0], word[1:]
//        }
//        return int32(anagramCount)
//}

//func FindReverse(runes []rune, startIndex, endIndex int) int {
//        word := runes[startIndex:endIndex]
//        if len(runes) == len(word) {
//                return 0
//        }

//        middle := int(math.Ceil(float64(len(runes))/2.0) - 1)
//        reversed := Reverse(runes)
//        for i := 0; i < len(runes); i++ {
//                // We don't start a new search if it starts after half
//                if i >= middle {
//                        return 0
//                }
//                // find first letter in word to start
//                if word[0] == reversed[i] {
//                        index := 0
//                        tempCount := 0
//                        // go letter by letter
//                        for _, l := range word {
//                                // go letter by letter
//                                if runes[i+index] == l {
//                                        tempCount += 1
//                                        if tempCount == len(word) {
//                                                return 1
//                                        }
//                                } else {
//                                        // if there isn't a match streak is broken
//                                        break
//                                }
//                                index++
//                        }
//                }
//        }
//        return 0
//}

//func Reverse(runes []rune) []rune {
//        var reverseHalf []rune
//        for i := len(runes) - 1; i >= 0; i-- {
//                reverseHalf = append([]rune{runes[i]}, reverseHalf...)
//        }
//        return reverseHalf
//}

func main() {
	jumps := sherlockAndAnagrams("ifailuhkqq")
	fmt.Println("jumps", jumps)

	jumps = sherlockAndAnagrams("abba")
	fmt.Println("jumps", jumps)
}
