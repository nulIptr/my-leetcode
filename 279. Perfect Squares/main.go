package main

import "math"

func numSquares(n int) int {

	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		for j := 0; i+j*j <= n; j++ {
			dp[i+j*j] = min(dp[i+j*j], dp[i]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
