package _21_Best_Time_To_Buy_And_Sell_Stock

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	left, right, maxP := 0, 0, 0
	for right < len(prices) {
		if prices[left] > prices[right] {
			left = right
		}
		if maxP < prices[right]-prices[left] {
			maxP = prices[right] - prices[left]
		}
		right++
	}
	return maxP
}

// 解法一 模拟 DP
func maxProfitDp(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min, maxP := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxP {
			maxP = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxP
}

// 解法二 单调栈
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	stack, res := []int{prices[0]}, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			index := len(stack) - 1
			for ; index >= 0; index-- {
				if stack[index] < prices[i] {
					break
				}
			}
			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}
		res = max(res, stack[len(stack)-1]-stack[0])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
