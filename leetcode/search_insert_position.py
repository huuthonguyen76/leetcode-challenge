from typing import List

class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        def __search_recursive(left: int, right: int):
            if left == right:
                return left

            if left + 1 == right:
                if nums[left] == target:
                    return left + 1
                elif nums[right] == target:
                    return right
                else:
                    return left + 1

            middle = (left + right) // 2
            if nums[middle] < target:
                return __search_recursive(middle, right)
            elif nums[middle] > target:
                return __search_recursive(left, middle)
            else:
                return middle

        if nums[0] == target:
            return 0
        
        if nums[-1] == target:
            return len(nums) - 1
        
        if nums[0] > target:
            return 0
        
        if nums[-1] < target:
            return len(nums)
        
        return __search_recursive(0, len(nums) - 1)


nums = [1, 3, 5, 6]
target = 2
print(Solution().searchInsert(nums, target))
