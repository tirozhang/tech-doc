package _070_Climbing_Stairs

var hash = make(map[int]int)

// 自顶向下
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	if v, ok := hash[n]; ok {
		return v
	}
	hash[n] = climbStairs(n-1) + climbStairs(n-2)
	return hash[n]
}

// 自底向上
func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
