package leetcode

import "math"

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := 1; i < len(memo); i++ {
		memo[i] = -2
	}
	return coinChangeDP(coins, amount, memo)
}

func coinChangeDP(coins []int, amount int, memo []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if memo[amount] != -2 {
		return memo[amount]
	}
	res := math.MaxInt
	for _, coin := range coins {
		subRes := coinChangeDP(coins, amount-coin, memo)
		if subRes == -1 {
			continue
		}
		if res > subRes+1 {
			res = subRes + 1
		}
	}
	if res == math.MaxInt {
		memo[amount] = -1
		return -1
	}
	memo[amount] = res
	return res
}
