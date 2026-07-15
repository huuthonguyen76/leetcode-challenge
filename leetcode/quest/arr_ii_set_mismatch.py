# https://leetcode.com/problems/set-mismatch/?envType=problem-list-v2&envId=dsa-linear-shoal-array-ii

from typing import List


class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        for idx, num in enumerate(nums):
            if idx == 0:
                continue

            if nums[idx - 1] != nums[idx] - 1:
                return (nums[idx - 1], nums[idx] + 1)


print(Solution().findErrorNums([1, 1]))
