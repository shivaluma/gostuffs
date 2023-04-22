package dynamic_programming

import "github.com/shivaluma/gostuffs/math"

func LongestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = math.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n][m]
}
