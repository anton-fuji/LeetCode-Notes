// Problem: 69. Sqrt(x)
// Difficulty: Easy
// Tags: Binary Search
// URL: https://leetcode.com/problems/sqrtx/
// Date: 2026-04-13

package sqrtx

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x
	res := 0

	for l <= r {
		mid := l + (r-l)/2

		if mid <= x/mid {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
