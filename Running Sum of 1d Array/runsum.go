package leetcode

func runningSum(nums []int) []int {
	narr := []int{}
	var sum int
	for _, val := range nums {
		sum = sum + val
		narr = append(narr, sum)
	}
	return narr
}
