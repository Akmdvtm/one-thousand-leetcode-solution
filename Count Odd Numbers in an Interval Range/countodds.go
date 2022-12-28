package leetcode

func countOdds(low int, high int) int {
	var sum int
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			sum++
		}
	}
	return sum
}
