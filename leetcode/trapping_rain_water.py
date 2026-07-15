# https://leetcode.com/problems/trapping-rain-water/

from typing import List

class Solution:
    def trap(self, height: List[int]) -> int:
        l_height = height

        l, r = 0, len(l_height) - 1
        total_water = 0
        left_max, right_max = l_height[l], l_height[r]
        while l < r:
            if l_height[l] < l_height[r]:
                l += 1

                if l_height[l] > left_max:
                    left_max = l_height[l]
                else:
                    if min(left_max, right_max) > l_height[l]:
                        total_water += min(left_max, right_max) - l_height[l]

            else:
                r -= 1

                if l_height[r] > right_max:
                    right_max = l_height[r]
                else:
                    if min(left_max, right_max) > l_height[r]:
                        total_water += min(left_max, right_max) - l_height[r]
        
        return total_water


# print(Solution().trap([0,1,0,2,1,0,1,3,2,1,2,1]))
print(Solution().trap([4,2,0,3,2,5]))
