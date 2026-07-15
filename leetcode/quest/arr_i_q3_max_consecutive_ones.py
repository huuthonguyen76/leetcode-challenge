# https://leetcode.com/problems/max-consecutive-ones/

from typing import List


class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        max_count = 0
        cur_count = 0

        for num in nums:
            if num != 1:
                cur_count = 0
                continue

            cur_count += 1
            if cur_count > max_count:
                max_count = cur_count

        return max_count

print(Solution().findMaxConsecutiveOnes([1,1,0,1,1,1]))
