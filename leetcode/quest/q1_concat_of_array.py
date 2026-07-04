# https://leetcode.com/problems/concatenation-of-array/?envType=problem-list-v2&envId=dsa-linear-shoal-array-i

from typing import List


class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        l_result = [0] * len(nums) * 2
        cursor = 0
        for idx, result in enumerate[int](l_result):
            if cursor > len(nums) - 1:
                cursor = 0

            l_result[idx] = nums[cursor]
            cursor += 1
        
        return l_result

print(Solution().getConcatenation([1, 2, 1]))
