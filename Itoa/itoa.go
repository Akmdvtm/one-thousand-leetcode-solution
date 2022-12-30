package leetcode

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}

	sign := 1

	index := 0
	n := len(s) - 1

	if s[index] == '-' {
		sign = -1
		index++
	} else if s[index] == '+' {
		sign = 1
		index++
	}
	result := 0
	for index <= n {
		digit := s[index] - '0'

		if !(digit >= 0 && digit <= 9) {
			break
		}

		if result > math.MaxInt32/10 || result == math.MaxInt32/10 && digit > math.MaxInt32%10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = (result * 10) + int(digit)
		index++
	}

	return sign * result
}
