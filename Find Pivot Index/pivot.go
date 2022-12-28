package pivot

func pivotIndex(nums []int) int {
	var sum, lsum, i int
	if len(nums) == 0 {
		return -1
	}

	for _, val := range nums {
		sum += val
	}

	for i = 0; i < len(nums); i++ {
		if lsum == (sum - nums[i]) {
			return i
		} else {
			lsum += nums[i]
			sum -= nums[i]
		}
	}
	return -1
}
