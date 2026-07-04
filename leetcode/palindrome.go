package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	txt := strconv.Itoa(x)
	txtLen := len(txt)

	for i, _ := range txt {
		if txt[i] != txt[txtLen-i-1] {
			return false
		}
	}

	return true
}

// func main() {
// 	a := []int{4, 6, 1, 3}
// 	fmt.Println(sort.Ints(a))
// 	fmt.Println(isPalindrome(12231))
// }
