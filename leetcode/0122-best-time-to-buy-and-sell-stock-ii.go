/**
* 122. Best Time to Buy and Sell Stock II (easy)
**/

package main

// one pass
func maxProfit(prices []int) int {
	ret := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1]
		}
	}

	return ret
}

// peak valley appraoch
func maxProfit2(prices []int) int {
	n := len(prices) - 1
	i, ret := 0, 0

	for i < n {
		for i < n && prices[i] >= prices[i+1] {
			i++
		}

		valley := prices[i]

		for i < n && prices[i] <= prices[i+1] {
			i++
		}

		peak := prices[i]

		ret += peak - valley
	}

	return ret
}
