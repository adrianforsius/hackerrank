package main

import (
	"fmt"
	"strings"
)

// ---------- To much memory -----------
//repeats := int(n) / len(s)
//rest := int(n) % len(s)
//if rest > 0 {
//        repeats += 1
//}
//long := strings.Repeat(s, int(repeats))
//sliced := long[:n]
//return int64(strings.Count(sliced, "a"))

// ---------- To slow -----------
//runes := []rune(s)
//var total int64
//for i := 0; i < int(n); i++ {
//        index := i % len(s)
//        val := runes[index]
//        if val == 'a' {
//                total += 1
//        }
//}
//return total

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	rep := int(n) / len(s)
	rest := int(n) % len(s)
	as := strings.Count(s, "a")
	count := as * int(rep)
	if rest > 0 {
		count += strings.Count(s[:rest], "a")
	}
	return int64(count)

}

func main() {
	out := repeatedString("aba", 10)
	fmt.Println("string", out)
}
