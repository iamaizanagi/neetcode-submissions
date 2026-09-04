class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix = 1
        postfix = 1

        final_arr = []

        for i in range(1, len(nums)+1):
            final_arr.append(prefix)
            prefix = prefix * nums[i-1]

        for j in range(len(nums)-1,-1,-1):
            final_arr[j] = postfix * final_arr[j]
            postfix = postfix * nums[j]

        return final_arr