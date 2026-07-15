# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/

from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        left_profits = len(prices) * [0]
        right_profits = len(prices) * [0]

        left_min = prices[0]
        right_max = prices[-1]

        for idx, _ in enumerate(prices):
            r_idx = len(prices) - 1 - idx

            left_profits[idx] = prices[idx] - left_min
            right_profits[r_idx] = right_max - prices[r_idx]

            if idx > 0 and left_profits[idx] < left_profits[idx - 1]:
                left_profits[idx] = left_profits[idx - 1]

            if r_idx < len(prices) - 1 and right_profits[r_idx] < right_profits[r_idx + 1]:
                right_profits[r_idx] = right_profits[r_idx + 1]

            if prices[idx] < left_min:
                left_min = prices[idx]
            
            if prices[r_idx] > right_max:
                right_max = prices[r_idx]

        # print(left_profits)
        # print(right_profits)
        total = 0
        for idx, _ in enumerate(prices):
            if idx != len(prices) - 1:
                total = max(total, (left_profits[idx] + right_profits[idx + 1]))
            else:
                total = max(total, (left_profits[idx]))

        return total

print(Solution().maxProfit([3,3,5,0,0,3,1,4]))
print(Solution().maxProfit([1,2,3,4,5]))
print(Solution().maxProfit([2, 1, 4]))
print(Solution().maxProfit([6,1,3,2,4,7]))
print(Solution().maxProfit([1,2,4,2,5,7,2,4,9,0]))
print(Solution().maxProfit([7,6,4,3,1]))
