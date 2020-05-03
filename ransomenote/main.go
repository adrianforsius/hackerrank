package main

import "fmt"

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	words := magazineMap(magazine)
	for _, word := range note {
		if _, ok := words[word]; ok {
			fmt.Println("No")
			return
		}
		if words[word] == 0 {
			fmt.Println("No")
			return
		}
		words[word] -= 1
	}
	fmt.Println("yes")
}
func magazineMap(words []string) map[string]int {
	wordsMap := make(map[string]int)
	for _, word := range words {
		if _, ok := wordsMap[word]; !ok {
			wordsMap[word] = 1
		} else {
			wordsMap[word] += 1
		}
	}
	return wordsMap
}

func main() {
	checkMagazine([]string{"give", "me", "one", "grand", "today", "night"}, []string{"give", "one", "grand", "today"})
	//fmt.Println("jumps", jumps)
}
