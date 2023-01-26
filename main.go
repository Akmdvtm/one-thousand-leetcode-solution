package main

import "fmt"

func main() {
	var s = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// s = [0,0,1,1,1,2,2,3,3,4]
	fmt.Println(removeDuplicates(s))
	fmt.Println(s)
}

func removeDuplicates(nums []int) int {
	var counter int
	var last int

	for i, val := range nums {
		if i == 0 {
			last = val
			continue
		} else if val == last {
			counter++
			nums[i] = -99
		}
		last = val
	}

	return counter
}
