package leetcode

import "strconv"

func isPalindrome(x int) bool {
	runes := []rune(strconv.Itoa(x))
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes) == strconv.Itoa(x)
}

//v2

// func isPalindrome(x int) bool {
// 	runes := []rune(strconv.Itoa(x))
// 	res := []rune{}
// 	for i := len(runes) - 1; i >= 0; i-- {
// 		res = append(res, runes[i])
// 	}
// 	return string(res) == strconv.Itoa(x)
// }
