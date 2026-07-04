# https://leetcode.com/problems/shuffle-the-array/?envType=problem-list-v2&envId=dsa-linear-shoal-array-i

from typing import List


class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        l_result = []
        cursor = len(nums) // 2
        for i in range(0, cursor, 1):
            l_result.append(nums[i]) 
            l_result.append(nums[cursor + i])

        return l_result


print(Solution().shuffle([2,5,1,3,4,7], 3))
 