// Problem: 70. Climbing Stairs
// Difficulty: Easy
// Tags: Fibonacci
// URL: https://leetcode.com/problems/climbing-stairs/
// Date: 2026-04-14

package climbing_stairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
