package medium

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	var l1, l2 = len(text1), len(text2)
	dp := [][]int{}
	for i := 0; i < l1+1; i++ {
		dp = append(dp, make([]int, l2+1))
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	// for i := 0; i < l1+1; i++ {
	// 	for j := 0; j < l2+1; j++ {
	// 		fmt.Printf("%d", dp[i][j])
	// 	}
	// 	fmt.Println()
	// }
	return dp[l1][l2]
}
